package proxy

import (
	"bufio"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/lich0821/ccNexus/internal/config"
	"github.com/lich0821/ccNexus/internal/logger"
)

// TestOpenAIResponsesStreamHeartbeatIsResponseCreated verifies that ccNexus
// sends a response.created event as the initial keep-alive for OpenAI Responses
// API streaming clients (Hermes / Python SDK openai>=1.0).
//
// The Python SDK's ResponseStreamState._create_initial_response() hard-requires
// that the very first parsed event has type=="response.created"; any other type
// raises RuntimeError which cancels the connection. SSE comments (": ...") and
// other event types (response.in_progress, etc.) all trigger this error.
func TestOpenAIResponsesStreamHeartbeatIsResponseCreated(t *testing.T) {
	logger.GetLogger().Clear()
	logger.GetLogger().SetMinLevel(logger.DEBUG)

	upstreamReady := make(chan struct{}, 1)
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		flusher, _ := w.(http.Flusher)
		upstreamReady <- struct{}{}
		// Slow upstream: silent for longer than one heartbeat interval.
		time.Sleep(200 * time.Millisecond)
		_, _ = w.Write([]byte(strings.Join([]string{
			`data: {"type":"response.output_text.delta","delta":"hello"}`,
			"",
			`data: {"type":"response.completed","response":{"id":"r1","object":"response","status":"completed","usage":{"input_tokens":3,"output_tokens":4,"total_tokens":7},"output":[{"type":"message","role":"assistant","content":[{"type":"output_text","text":"hello"}]}]}}`,
			"",
			"data: [DONE]",
			"",
		}, "\n")))
		if flusher != nil {
			flusher.Flush()
		}
	}))
	defer upstream.Close()

	endpoint := failoverPolicyTestEndpoint("Primary", upstream.URL)
	p := newFailoverPolicyTestProxy([]config.Endpoint{endpoint}, upstream.Client())
	p.streamHeartbeatInterval = 30 * time.Millisecond

	proxySrv := httptest.NewServer(http.HandlerFunc(p.handleProxy))
	defer proxySrv.Close()

	firstDataCh := make(chan string, 1)
	errCh := make(chan error, 1)

	go func() {
		resp, err := proxySrv.Client().Post(
			proxySrv.URL+"/v1/responses",
			"application/json",
			strings.NewReader(`{"model":"gpt-5.5","stream":true,"input":"hi"}`),
		)
		if err != nil {
			errCh <- err
			return
		}
		defer resp.Body.Close()
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "data:") {
				firstDataCh <- line
				return
			}
		}
		firstDataCh <- ""
	}()

	<-upstreamReady

	select {
	case err := <-errCh:
		t.Fatalf("request failed: %v", err)
	case line := <-firstDataCh:
		if line == "" {
			t.Fatal("no data: event received; Python SDK would raise RuntimeError and cancel")
		}
		jsonPart := strings.TrimSpace(strings.TrimPrefix(line, "data:"))
		if !strings.Contains(jsonPart, `"type":"response.created"`) {
			t.Fatalf("Python SDK requires first event type==response.created, got: %q", line)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("no data: event within 100ms; heartbeat was only SSE comments which Python SDK skips")
	}
}
