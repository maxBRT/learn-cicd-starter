package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := []struct {
		name          string
		headerKey     string
		headerValue   string
		expected      string
		errorExpected bool
	}{
		{
			name:          "valid header",
			headerKey:     "Authorization",
			headerValue:   "ApiKeyhelloworld",
			expected:      "helloworld",
			errorExpected: false,
		},
		{
			name:          "invalid key",
			headerKey:     "Jello",
			headerValue:   "hello-world",
			expected:      "",
			errorExpected: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h := http.Header{}
			h.Add(tc.headerKey, tc.headerValue)
			got, err := GetAPIKey(h)
			if tc.errorExpected {
				if err == nil {
					t.Errorf("TestGetApiKey() error = %v, wantErr %v", err, tc.errorExpected)
					return
				}
			} else {
				if got != tc.expected {
					t.Errorf("TestGetApiKey() got = %v, want %v", got, tc.expected)
				}
			}
		})
	}
}
