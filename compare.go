package main

import (
	"bytes"
	"fmt"
)

// CompareSecrets compares local and remote secret values
func CompareSecrets(local []byte, remote string, secretName string) *ComparisonResult {
	// Quick size check for early return
	if len(local) != len(remote) {
		return &ComparisonResult{
			SecretName: secretName,
			IsEqual:    false,
		}
	}

	// Byte-by-byte comparison
	isEqual := bytes.Equal(local, []byte(remote))

	return &ComparisonResult{
		SecretName: secretName,
		IsEqual:    isEqual,
	}
}

// String returns a human-readable representation of the comparison result
func (r *ComparisonResult) String() string {
	if r.IsEqual {
		return fmt.Sprintf("%s : No Differences", r.SecretName)
	}
	return fmt.Sprintf("%s : With Differences", r.SecretName)
}
