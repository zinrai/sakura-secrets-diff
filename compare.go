package main

import (
	"bytes"
	"fmt"
)

func CompareSecrets(local []byte, remote string, secretName string) *ComparisonResult {
	if len(local) != len(remote) {
		return &ComparisonResult{
			SecretName: secretName,
			IsEqual:    false,
		}
	}

	isEqual := bytes.Equal(local, []byte(remote))

	return &ComparisonResult{
		SecretName: secretName,
		IsEqual:    isEqual,
	}
}

func (r *ComparisonResult) String() string {
	if r.IsEqual {
		return fmt.Sprintf("%s : No Differences", r.SecretName)
	}
	return fmt.Sprintf("%s : With Differences", r.SecretName)
}
