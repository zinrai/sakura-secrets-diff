package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var (
		name          string
		zone          string
		secretVersion int
	)

	flag.StringVar(&name, "name", "", "Secret name (required)")
	flag.StringVar(&zone, "zone", "is1a", "Zone name (default: is1a)")
	flag.IntVar(&secretVersion, "version", 0, "Secret version (default: 0 = latest)")
	flag.Usage = usage
	flag.Parse()

	if name == "" {
		usage()
		os.Exit(2)
	}

	if err := run(name, zone, secretVersion); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(2)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "sakura-secrets-diff %s (commit %s, built %s)\n\n", version, commit, date)
	fmt.Fprintf(os.Stderr, "Usage: sakura-secrets-diff -name <secret-name> [options]\n\n")
	fmt.Fprintf(os.Stderr, "Compares a secret value from stdin with the value in Secret Manager.\n\nOptions:\n")
	flag.PrintDefaults()
}

func run(name, zone string, secretVersion int) error {
	// Read local value from stdin
	localValue, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("failed to read stdin: %w", err)
	}

	if len(localValue) == 0 {
		return fmt.Errorf("no input provided via stdin")
	}

	// Resolve the Vault resource ID from environment variables
	vaultID, err := LoadVaultID()
	if err != nil {
		return err
	}

	// Create API client
	op, err := NewSecretOp(zone, vaultID)
	if err != nil {
		return err
	}

	// Get remote value from Secret Manager
	remoteValue, err := GetSecret(op, name, secretVersion)
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
