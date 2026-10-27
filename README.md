# sakura-secrets-diff

A command-line tool to compare a secret value from stdin with a secret stored in Sakura Cloud Secret Manager. It reports only whether the two differ and never outputs either value.

## Usage

```bash
$ printf '%s' "$value" | sakura-secrets-diff -name <secret-name>
```

The value is read from stdin rather than from an argument, so it can be piped from any command that decrypts or renders secrets without appearing in the process list. Run the tool once per secret.

The comparison is byte by byte. `echo` appends a newline, which makes the values differ unless the stored secret ends with one too.

`sakura-secrets-diff -h` lists the options.

### Output and exit codes

```
{secret-name} : No Differences
{secret-name} : With Differences
```

- `0`: No differences
- `1`: With differences
- `2`: Error (invalid usage, API failure, authentication error, etc.)

## Configuration

Set the Vault resource ID:

```bash
$ export SAKURA_SECRETS_ID="your-vault-resource-id"
```

Set either static API keys:

```bash
$ export SAKURA_ACCESS_TOKEN="your-access-token"
$ export SAKURA_ACCESS_TOKEN_SECRET="your-access-token-secret"
```

or service principal credentials:

```bash
$ export SAKURA_SERVICE_PRINCIPAL_ID="your-service-principal-id"
$ export SAKURA_SERVICE_PRINCIPAL_KEY_KID="your-key-id"
$ export SAKURA_PRIVATE_KEY_PATH="/path/to/private-key.pem"
```

## Requires

- A Sakura Cloud account with access to Secret Manager
- API credentials: static API keys or a service principal

## License

This project is licensed under the [MIT License](./LICENSE).
