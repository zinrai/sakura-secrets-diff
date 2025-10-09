package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	apiBaseURL = "https://secure.sakura.ad.jp/cloud/zone"
	apiPath    = "/api/cloud/1.1/secretmanager"
)

// SakuraClient represents the Sakura Cloud API client
type SakuraClient struct {
	Config     *Config
	HTTPClient *http.Client
}

// NewClient creates a new SakuraClient
func NewClient(config *Config) *SakuraClient {
	return &SakuraClient{
		Config: config,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetSecret retrieves a secret from the specified vault using the unveil API
func (c *SakuraClient) GetSecret(secretName string, version int) (string, error) {
	url := fmt.Sprintf("%s/%s%s/vaults/%s/secrets/unveil",
		apiBaseURL,
		c.Config.Zone,
		apiPath,
		c.Config.ResourceID,
	)

	reqBody := UnveilRequest{
		Secret: SecretRequest{
			Name:    secretName,
			Version: version,
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.Config.AccessToken, c.Config.AccessTokenSecret)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	var unveilResp UnveilResponse
	if err := json.Unmarshal(body, &unveilResp); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	return unveilResp.Secret.Value, nil
}
