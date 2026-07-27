# sakura-secrets-diff

A command-line tool to compare a secret value from stdin with a secret stored in Sakura Cloud Secret Manager.

## Features

- Compares secret values between stdin and Sakura Cloud Secret Manager
- Reads secret values from stdin
- Reports only whether differences exist or not
- Never outputs actual secret values

## Requirements

- Sakura Cloud account with Secret Manager access
- Valid API credentials (static API keys or a service principal)

## Configuration

Set the Vault resource ID:

```bash
$ export SAKURA_SECRETS_ID="your-vault-resource-id"
```

API credentials are resolved by [saclient-go](https://github.com/sacloud/saclient-go). Set either static API keys:

```bash
$ export SAKURA_ACCESS_TOKEN="your-access-token"
$ export SAKURA_ACCESS_TOKEN_SECRET="your-access-token-secret"
```

or service principal credentials:

```bash
$ export SAKURA_SERVICE_PRINCIPAL_ID="your-service-principal-id"
$ export SAKURA_SERVICE_PRINCIPAL_KEY_ID="your-key-id"
$ export SAKURA_PRIVATE_KEY_PATH="/path/to/private-key.pem"
```

## Usage

### Basic Usage

```bash
$ echo "$value" | sakura-secrets-diff -name <secret-name>
```

### With sops and yq

Use case for integrating [sops](https://github.com/getsops/sops) and [yq](https://github.com/mikefarah/yq).

```bash
$ sops -d secrets.yaml | yq -r 'to_entries[] | "\(.key)\t\(.value)"' | \
while IFS=$'\t' read -r key value; do
  echo "$value" | sakura-secrets-diff -name "$key"
done
```

### Options

- `-name` (required): Secret name
- `-zone` (optional, default: `is1a`): Zone name
- `-version` (optional, default: `0`): Secret version (0 = latest)

### Exit Codes

- `0`: No differences (secrets match)
- `1`: With differences (secrets don't match)
- `2`: Error (API failure, authentication error, etc.)

## Output Format

The tool outputs one of the following:

```
{secret-name} : No Differences
```

or

```
{secret-name} : With Differences
```

The actual secret values are never displayed.

## Examples

### Check if a secret matches

```bash
$ echo "my-secret-value" | sakura-secrets-diff -name my-secret
# Output: my-secret : No Differences
# Exit code: 0
```

### Suppress output (use exit code only)

```bash
$ echo "$value" | sakura-secrets-diff -name my-secret >/dev/null
$ echo $?  # 0, 1, or 2
```

## Security

- Secret values are read from stdin (not command-line arguments)
- Secret values are never written to stdout
- API credentials are read from environment variables
- All comparisons are done in memory

## Notes

This tool performs byte-by-byte comparison. If you get "With Differences" unexpectedly, check for trailing newlines in your input.

## License

This project is licensed under the [MIT License](./LICENSE).
