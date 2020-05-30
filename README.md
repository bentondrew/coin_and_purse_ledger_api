# coin_and_purse_ledger_service
Service which provides the ledger functions.

## Rust tips
### Create new cargo package with rust docker image
```bash
docker run --rm -v "$PWD":/usr/src -w /usr/src --user "$(id -u)":"$(id -g)" -e USER=$USER rust:1.43 cargo new ledger_api
```

## Docker instructions
### Build image
```bash
docker build -t ledger-api:v0.1.0 .
```

### Run container
```bash
docker run --rm ledger-api:v0.1.0
```
