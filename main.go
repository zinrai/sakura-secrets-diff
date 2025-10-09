package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var (
		resourceID string
		name       string
		zone       string
		version    int
	)

	flag.StringVar(&resourceID, "resource-id", "", "Vault resource ID (required)")
	flag.StringVar(&name, "name", "", "Secret name (required)")
	flag.StringVar(&zone, "zone", "is1a", "Zone name (default: is1a)")
	flag.IntVar(&version, "version", 0, "Secret version (default: 0 = latest)")
	flag.Parse()

	if err := run(resourceID, name, zone, version); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(2)
	}
}

func run(resourceID, name, zone string, version int) error {
	// Validate required parameters
	if resourceID == "" {
		return fmt.Errorf("-resource-id is required")
	}
	if name == "" {
		return fmt.Errorf("-name is required")
	}

	// Read local value from stdin
	localValue, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("failed to read stdin: %w", err)
	}

	if len(localValue) == 0 {
		return fmt.Errorf("no input provided via stdin")
	}

	// Load configuration from environment variables
	config, err := LoadConfig(zone, resourceID)
	if err != nil {
		return err
	}

	// Create API client
	client := NewClient(config)

	// Get remote value from Secret Manager
	remoteValue, err := client.GetSecret(name, version)
	if err != nil {
		return fmt.Errorf("failed to get secret from API: %w", err)
	}

	// Compare secrets
	result := CompareSecrets(localValue, remoteValue, name)

	// Output result
	fmt.Println(result.String())

	// Set exit code based on comparison result
	if !result.IsEqual {
		os.Exit(1)
	}

	return nil
}
