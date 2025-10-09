package main

import (
	"testing"
)

func TestCompareSecrets(t *testing.T) {
	tests := []struct {
		name       string
		local      []byte
		remote     string
		secretName string
		wantEqual  bool
	}{
		{
			name:       "exact match",
			local:      []byte("my-secret-value"),
			remote:     "my-secret-value",
			secretName: "test-secret",
			wantEqual:  true,
		},
		{
			name:       "different values",
			local:      []byte("my-secret-value"),
			remote:     "different-value",
			secretName: "test-secret",
			wantEqual:  false,
		},
		{
			name:       "different size",
			local:      []byte("short"),
			remote:     "much longer value",
			secretName: "test-secret",
			wantEqual:  false,
		},
		{
			name:       "multiline match",
			local:      []byte("line1\nline2\nline3"),
			remote:     "line1\nline2\nline3",
			secretName: "test-secret",
			wantEqual:  true,
		},
		{
			name:       "multiline different",
			local:      []byte("line1\nline2\nline3"),
			remote:     "line1\nline2\nline4",
			secretName: "test-secret",
			wantEqual:  false,
		},
		{
			name:       "trailing newline difference",
			local:      []byte("value\n"),
			remote:     "value",
			secretName: "test-secret",
			wantEqual:  false,
		},
		{
			name:       "both empty",
			local:      []byte(""),
			remote:     "",
			secretName: "test-secret",
			wantEqual:  true,
		},
		{
			name:       "empty local",
			local:      []byte(""),
			remote:     "value",
			secretName: "test-secret",
			wantEqual:  false,
		},
		{
			name:       "empty remote",
			local:      []byte("value"),
			remote:     "",
			secretName: "test-secret",
			wantEqual:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CompareSecrets(tt.local, tt.remote, tt.secretName)

			if result.IsEqual != tt.wantEqual {
				t.Errorf("CompareSecrets() IsEqual = %v, want %v", result.IsEqual, tt.wantEqual)
			}

			if result.SecretName != tt.secretName {
				t.Errorf("CompareSecrets() SecretName = %v, want %v", result.SecretName, tt.secretName)
			}
		})
	}
}

func TestComparisonResult_String(t *testing.T) {
	tests := []struct {
		name       string
		result     *ComparisonResult
		wantOutput string
	}{
		{
			name: "no differences",
			result: &ComparisonResult{
				SecretName: "my-secret",
				IsEqual:    true,
			},
			wantOutput: "my-secret : No Differences",
		},
		{
			name: "with differences",
			result: &ComparisonResult{
				SecretName: "my-secret",
				IsEqual:    false,
			},
			wantOutput: "my-secret : With Differences",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.result.String()
			if got != tt.wantOutput {
				t.Errorf("ComparisonResult.String() = %q, want %q", got, tt.wantOutput)
			}
		})
	}
}
