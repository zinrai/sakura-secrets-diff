package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sacloud/saclient-go"
	secretmanager "github.com/sacloud/secretmanager-api-go"
	v1 "github.com/sacloud/secretmanager-api-go/apis/v1"
)

// NewSecretOp creates a Secret Manager API client bound to the specified Vault.
// Credentials are resolved by saclient-go from environment variables,
// supporting both static API keys and service principals.
func NewSecretOp(zone, vaultID string) (secretmanager.SecretAPI, error) {
	endpoint := fmt.Sprintf("SAKURA_ENDPOINTS_SECRETMANAGER=https://secure.sakura.ad.jp/cloud/zone/%s/api/cloud/1.1", zone)

	var sc saclient.Client
	if err := sc.SetEnviron(append(os.Environ(), endpoint)); err != nil {
		return nil, fmt.Errorf("failed to configure saclient: %w", err)
	}
	if err := sc.Populate(); err != nil {
		return nil, fmt.Errorf("failed to configure saclient: %w", err)
	}

	client, err := secretmanager.NewClient(&sc)
	if err != nil {
		return nil, fmt.Errorf("failed to create Secret Manager client: %w", err)
	}

	return secretmanager.NewSecretOp(client, vaultID), nil
}

// GetSecret retrieves a secret value from the Vault using the unveil API.
// version 0 means the latest version.
func GetSecret(op secretmanager.SecretAPI, secretName string, version int) (string, error) {
	req := v1.Unveil{
		Name: secretName,
	}
	if version != 0 {
		req.Version = v1.NewOptNilInt(version)
	}

	res, err := op.Unveil(context.Background(), req)
	if err != nil {
		return "", err
	}

	return res.Value, nil
}
