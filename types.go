package main

// Config holds the application configuration
type Config struct {
	AccessToken       string
	AccessTokenSecret string
	Zone              string
	ResourceID        string
}

// UnveilRequest represents the request body for the unveil API
type UnveilRequest struct {
	Secret SecretRequest `json:"Secret"`
}

// SecretRequest represents the secret request parameters
type SecretRequest struct {
	Name    string `json:"Name"`
	Version int    `json:"Version"`
}

// UnveilResponse represents the response from the unveil API
type UnveilResponse struct {
	Secret SecretResponse `json:"Secret"`
}

// SecretResponse represents the secret response data
type SecretResponse struct {
	Name    string `json:"Name"`
	Version int    `json:"Version"`
	Value   string `json:"Value"`
}

// ComparisonResult represents the result of a secret comparison
type ComparisonResult struct {
	SecretName string
	IsEqual    bool
}
