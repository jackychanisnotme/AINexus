package proxy

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/lich0821/ccNexus/internal/config"
)

func TestHandleProxyRejectsInvalidJSONBeforeEndpointAttempt(t *testing.T) {
	upstreamHits := 0
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upstreamHits++
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":"resp-upstream","usage":{"input_tokens":1,"output_tokens":1},"output":[]}`))
	}))
	defer upstream.Close()

	p := newFailoverPolicyTestProxy([]config.Endpoint{
		failoverPolicyTestEndpoint("Primary", upstream.URL),
	}, upstream.Client())

	tests := []struct {
		name string
		body string
	}{
		{name: "empty body", body: ""},
		{name: "malformed json", body: `{"model":"gpt-5.5"`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			upstreamHits = 0
			req := httptest.NewRequest(http.MethodPost, "/v1/responses", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			p.handleProxy(rec, req)

			if rec.Code != http.StatusBadRequest {
				t.Fatalf("expected status 400, got %d body=%q", rec.Code, rec.Body.String())
			}
			if upstreamHits != 0 {
				t.Fatalf("expected invalid request to skip upstream endpoints, got hits=%d", upstreamHits)
			}
			if !strings.Contains(rec.Body.String(), "invalid_request_error") {
				t.Fatalf("expected structured invalid request response, got %q", rec.Body.String())
			}
		})
	}
}
