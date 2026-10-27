package main

import (
	"testing"
)

func TestCompareSecretsReportsDifference(t *testing.T) {
	tests := []struct {
		name   string
		local  []byte
		remote string
	}{
		{
			name:   "same length, different content",
			local:  []byte("my-secret-value"),
			remote: "different-value",
		},
		{
			name:   "trailing newline only on one side",
			local:  []byte("value\n"),
			remote: "value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if CompareSecrets(tt.local, tt.remote) {
				t.Errorf("CompareSecrets(%q, %q) = true, want false", tt.local, tt.remote)
			}
		})
	}
}
