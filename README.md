# coin_and_purse_ledger_service
Service which provides the ledger functions.

Repo structure resources:
- [Go best practices](https://peter.bourgon.org/go-best-practices-2016/#repository-structure)


## Docker image build command

### Local build
```Bash
docker build -t coin_and_purse_ledger_service:0.0.2 .
```

## Image run command

```Bash
docker run coin_and_purse_ledger_service:0.0.2
```

### Expected output
Should print `Hello world!`
