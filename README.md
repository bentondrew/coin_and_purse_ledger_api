# coin_and_purse_ledger_service
Service which provides the ledger functions.

Repo structure resources:
- [Go best practices](https://peter.bourgon.org/go-best-practices-2016/#repository-structure)
Server structure resources:
- [Go server components](https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831)
DB Mocking:
- [Adding DB to web application](https://www.sohamkamani.com/blog/2017/10/18/golang-adding-database-to-web-application/)


## Docker image build command

### Local build
```Bash
docker build -t coin_and_purse_ledger_service:0.0.27 .
```

## Image run command

```Bash
docker run --rm --network appnet -p 80:8080 coin_and_purse_ledger_service:0.0.27
```
