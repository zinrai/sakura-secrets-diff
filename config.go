package main

import (
	"fmt"
	"os"
)

// LoadConfig loads configuration from environment variables
func LoadConfig(zone string) (*Config, error) {
	accessToken := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")
	secretsID := os.Getenv("SAKURACLOUD_SECRETS_ID")

	var missing []string
	if accessToken == "" {
		missing = append(missing, "SAKURACLOUD_ACCESS_TOKEN")
	}
	if accessTokenSecret == "" {
		missing = append(missing, "SAKURACLOUD_ACCESS_TOKEN_SECRET")
	}
	if secretsID == "" {
		missing = append(missing, "SAKURACLOUD_SECRETS_ID")
	}

	if len(missing) > 0 {
		return nil, fmt.Errorf("required environment variables not set: %v", missing)
	}

	return &Config{
		AccessToken:       accessToken,
		AccessTokenSecret: accessTokenSecret,
		Zone:              zone,
		ResourceID:        secretsID,
	}, nil
}
