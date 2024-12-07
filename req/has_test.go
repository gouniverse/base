package req

import (
	"net/http"
	"strings"
	"testing"
)

func TestHas(t *testing.T) {
	tests := []struct {
		name   string
		method string
		url    string
		body   string
		key    string
		want   bool
	}{
		// ... (add test cases from previous responses)
		{
			name:   "GET and POST parameters exist",
			method: "POST",
			url:    "http://example.com/",
			body:   "key=value",
			key:    "key",
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.url, strings.NewReader(tt.body))
			if err != nil {
				t.Fatalf("Failed  to create request: %v", err)
			}

			if tt.method == "POST" {
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			}

			if got := Has(req, tt.key); got != tt.want {
				t.Errorf("Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasGet(t *testing.T) {
	tests := []struct {
		name string
		url  string
		key  string
		want bool
	}{
		{
			name: "GET parameter exists",
			url:  "http://example.com/?key=value",
			key:  "key",
			want: true,
		},
		{
			name: "GET parameter does not exist",
			url:  "http://example.com/",
			key:  "key",
			want: false,
		},
		{
			name: "GET parameter is empty",
			url:  "http://example.com/?key=",
			key:  "key",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.url, nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			if got := HasGet(req, tt.key); got != tt.want {
				t.Errorf("HasGet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasPost(t *testing.T) {
	tests := []struct {
		name string
		url  string
		body string
		key  string
		want bool
	}{
		{
			name: "POST parameter exists",
			url:  "http://example.com/",
			body: "key=value",
			key:  "key",
			want: true,
		},
		{
			name: "POST parameter does not exist",
			url:  "http://example.com/",
			body: "",
			key:  "key",
			want: false,
		},
		{
			name: "POST parameter is empty",
			url:  "http://example.com/",
			body: "key=",
			key:  "key",
			want: true,
		},
		{
			name: "POST parameter parsing error",
			url:  "http://example.com/",
			body: "invalid_data",
			key:  "key",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", tt.url, strings.NewReader(tt.body))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			if got := HasPost(req, tt.key); got != tt.want {
				t.Errorf("HasPost() = %v, want %v", got, tt.want)
			}
		})
	}
}
