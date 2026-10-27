package main

import (
	"testing"
)

func TestCompareSecrets(t *testing.T) {
	tests := []struct {
		name      string
		local     []byte
		remote    string
		wantEqual bool
	}{
		{
			name:      "exact match",
			local:     []byte("my-secret-value"),
			remote:    "my-secret-value",
			wantEqual: true,
		},
		{
			name:      "different values",
			local:     []byte("my-secret-value"),
			remote:    "different-value",
			wantEqual: false,
		},
		{
			name:      "different size",
			local:     []byte("short"),
			remote:    "much longer value",
			wantEqual: false,
		},
		{
			name:      "multiline match",
			local:     []byte("line1\nline2\nline3"),
			remote:    "line1\nline2\nline3",
			wantEqual: true,
		},
		{
			name:      "multiline different",
			local:     []byte("line1\nline2\nline3"),
			remote:    "line1\nline2\nline4",
			wantEqual: false,
		},
		{
			name:      "trailing newline difference",
			local:     []byte("value\n"),
			remote:    "value",
			wantEqual: false,
		},
		{
			name:      "both empty",
			local:     []byte(""),
			remote:    "",
			wantEqual: true,
		},
		{
			name:      "empty local",
			local:     []byte(""),
			remote:    "value",
			wantEqual: false,
		},
		{
			name:      "empty remote",
			local:     []byte("value"),
			remote:    "",
			wantEqual: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareSecrets(tt.local, tt.remote); got != tt.wantEqual {
				t.Errorf("CompareSecrets() = %v, want %v", got, tt.wantEqual)
			}
		})
	}
}
