package convert

import "testing"

func TestParseSSEDataWithAndWithoutSpace(t *testing.T) {
	cases := []struct {
		name      string
		input     string
		wantEvent string
		wantData  string
	}{
		{"with space", "event: message\ndata: {\"a\":1}", "message", "{\"a\":1}"},
		{"without space", "event:message\ndata:{\"a\":1}", "message", "{\"a\":1}"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ev, data := parseSSE([]byte(tc.input))
			if ev != tc.wantEvent {
				t.Errorf("event = %q, want %q", ev, tc.wantEvent)
			}
			if data != tc.wantData {
				t.Errorf("data = %q, want %q", data, tc.wantData)
			}
		})
	}
}

func TestFilterNonResponsesStreamEventDropsChatChunk(t *testing.T) {
	chatChunk := []byte("data: {\"id\":\"chatcmpl-x\",\"object\":\"chat.completion.chunk\",\"choices\":[{\"delta\":{\"role\":\"assistant\"}}]}\n\n")
	if got := FilterNonResponsesStreamEvent(chatChunk); got != nil {
		t.Fatalf("expected chat.completion.chunk to be filtered, got %q", got)
	}
}

func TestFilterNonResponsesStreamEventKeepsResponsesEvent(t *testing.T) {
	created := []byte("data: {\"type\":\"response.created\",\"sequence_number\":0,\"response\":{}}\n\n")
	if got := FilterNonResponsesStreamEvent(created); string(got) != string(created) {
		t.Fatalf("expected response.created to pass through unchanged")
	}
}

func TestFilterNonResponsesStreamEventKeepsDone(t *testing.T) {
	done := []byte("data: [DONE]\n\n")
	if got := FilterNonResponsesStreamEvent(done); string(got) != string(done) {
		t.Fatalf("expected [DONE] to pass through")
	}
}
