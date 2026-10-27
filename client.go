package main

import (
	"context"
	"fmt"
	"os"

	secretmanager "github.com/sacloud/sacloud-sdk-go/api/secretmanager"
	"github.com/sacloud/sacloud-sdk-go/common/saclient"
)

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

const latestSecretVersion = 0

func GetSecret(op secretmanager.SecretAPI, secretName string, version int) (string, error) {
	req := secretmanager.UnveilParams{
		Name: secretName,
	}
	if version != latestSecretVersion {
		req.Version = &version
	}

	res, err := op.Unveil(context.Background(), req)
	if err != nil {
		return "", err
	}

	return res.Value, nil
}
