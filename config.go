package main

import (
	"fmt"
	"os"
)

func LoadVaultID() (string, error) {
	if v := os.Getenv("SAKURA_SECRETS_ID"); v != "" {
		return v, nil
	}
	return "", fmt.Errorf("required environment variable not set: SAKURA_SECRETS_ID")
}
