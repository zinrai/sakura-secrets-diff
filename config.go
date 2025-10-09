package main

import (
	"fmt"
	"os"
)

// LoadConfig loads configuration from environment variables
func LoadConfig(zone, resourceID string) (*Config, error) {
	accessToken := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")

	if accessToken == "" {
		return nil, fmt.Errorf("SAKURACLOUD_ACCESS_TOKEN environment variable is required")
	}
	if accessTokenSecret == "" {
		return nil, fmt.Errorf("SAKURACLOUD_ACCESS_TOKEN_SECRET environment variable is required")
	}

	return &Config{
		AccessToken:       accessToken,
		AccessTokenSecret: accessTokenSecret,
		Zone:              zone,
		ResourceID:        resourceID,
	}, nil
}
