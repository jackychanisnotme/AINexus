package config

import "testing"

func TestNormalizeThinkingEffortPreservesProviderDefault(t *testing.T) {
	tests := map[string]string{
		"":        "",
		" ":       "",
		"default": "",
		"auto":    "",
		"inherit": "",
		"off":     "off",
		"low":     "low",
		"medium":  "medium",
		"high":    "high",
		"xhigh":   "xhigh",
		"invalid": "off",
	}

	for input, want := range tests {
		if got := NormalizeThinkingEffort(input); got != want {
			t.Fatalf("NormalizeThinkingEffort(%q) = %q, want %q", input, got, want)
		}
	}
}
