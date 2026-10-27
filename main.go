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
	flag.StringVar(&zone, "zone", "is1a", "Zone name")
	flag.IntVar(&secretVersion, "secret-version", 0, "Secret version, 0 for the latest")
	flag.Usage = usage
	flag.Parse()

	if name == "" {
		usage()
		os.Exit(2)
	}

	equal, err := run(name, zone, secretVersion)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(2)
	}

	if !equal {
		fmt.Printf("%s : With Differences\n", name)
		os.Exit(1)
	}

	fmt.Printf("%s : No Differences\n", name)
}

func usage() {
	fmt.Fprintf(os.Stderr, "sakura-secrets-diff %s (commit %s, built %s)\n\n", version, commit, date)
	fmt.Fprintf(os.Stderr, "Usage: sakura-secrets-diff -name <secret-name> [options]\n\n")
	fmt.Fprintf(os.Stderr, "Compares a secret value from stdin with the value in Secret Manager.\n\nOptions:\n")
	flag.PrintDefaults()
}

func run(name, zone string, secretVersion int) (bool, error) {
	localValue, err := io.ReadAll(os.Stdin)
	if err != nil {
		return false, fmt.Errorf("failed to read stdin: %w", err)
	}

	if len(localValue) == 0 {
		return false, fmt.Errorf("no input provided via stdin")
	}

	vaultID, err := LoadVaultID()
	if err != nil {
		return false, err
	}

	op, err := NewSecretOp(zone, vaultID)
	if err != nil {
		return false, err
	}

	remoteValue, err := GetSecret(op, name, secretVersion)
	if err != nil {
		return false, fmt.Errorf("failed to get secret from API: %w", err)
	}

	return CompareSecrets(localValue, remoteValue), nil
}
