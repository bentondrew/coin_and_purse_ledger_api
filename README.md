# coin_and_purse_ledger_service
Service which provides the ledger functions.

## Rust Docker Image
https://hub.docker.com/_/rust/

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

## Docker helpers
### Remove dangling images [source](https://stackoverflow.com/a/33913711)
```bash
docker rmi $(docker images --filter "dangling=true" -q --no-trunc)
```

## Potential Rust Web Service Frameworks
https://actix.rs/docs/installation/
