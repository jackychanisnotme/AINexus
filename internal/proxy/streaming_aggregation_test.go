package proxy

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/lich0821/ccNexus/internal/config"
)

func TestHandleStreamingAsNonStreamingAggregatesOpenAIChatChunks(t *testing.T) {
	rawSSE := strings.Join([]string{
		`data: {"id":"chatcmpl-1","object":"chat.completion.chunk","created":1710000000,"model":"gpt-test","choices":[{"index":0,"delta":{"role":"assistant","content":"hello "},"finish_reason":null}]}`,
		"",
		`data: {"id":"chatcmpl-1","object":"chat.completion.chunk","created":1710000000,"model":"gpt-test","choices":[{"index":0,"delta":{"content":"world","tool_calls":[{"index":0,"id":"call_1","type":"function","function":{"name":"lookup","arguments":"{\"q\""}}]},"finish_reason":null}]}`,
		"",
		`data: {"id":"chatcmpl-1","object":"chat.completion.chunk","created":1710000000,"model":"gpt-test","choices":[{"index":0,"delta":{"tool_calls":[{"index":0,"function":{"arguments":":\"codex\"}"}}]},"finish_reason":"tool_calls"}]}`,
		"",
		`data: {"id":"chatcmpl-1","object":"chat.completion.chunk","created":1710000000,"model":"gpt-test","choices":[],"usage":{"prompt_tokens":11,"completion_tokens":7,"total_tokens":18}}`,
		"",
		"data: [DONE]",
		"",
	}, "\n")

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": []string{"text/event-stream"}},
		Body:       io.NopCloser(strings.NewReader(rawSSE)),
	}
	rec := httptest.NewRecorder()
	p := &Proxy{}

	in, out, text, err := p.handleStreamingAsNonStreaming(rec, resp, config.Endpoint{Name: "OpenAI"}, &passthroughResponseTransformer{}, 0)
	if err != nil {
		t.Fatalf("handleStreamingAsNonStreaming failed: %v", err)
	}
	if in != 11 || out != 7 {
		t.Fatalf("expected usage in=11 out=7, got in=%d out=%d", in, out)
	}
	if text != "hello world" {
		t.Fatalf("unexpected output text: %q", text)
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &payload); err != nil {
		t.Fatalf("response is not json: %v", err)
	}
	if payload["object"] != "chat.completion" {
		t.Fatalf("expected chat.completion object, got %#v", payload["object"])
	}

	choices := payload["choices"].([]interface{})
	message := choices[0].(map[string]interface{})["message"].(map[string]interface{})
	if message["content"] != "hello world" {
		t.Fatalf("unexpected message content: %#v", message["content"])
	}
	toolCalls := message["tool_calls"].([]interface{})
	function := toolCalls[0].(map[string]interface{})["function"].(map[string]interface{})
	if function["name"] != "lookup" || function["arguments"] != `{"q":"codex"}` {
		t.Fatalf("unexpected tool call function: %#v", function)
	}
}
