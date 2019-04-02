# coin_and_purse_ledger_api
Service which provides the ledger API.

Repo structure resources:
- [Go best practices](https://peter.bourgon.org/go-best-practices-2016/#repository-structure)
Server structure resources:
- [Go server components](https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831)
DB Mocking:
- [Adding DB to web application](https://www.sohamkamani.com/blog/2017/10/18/golang-adding-database-to-web-application/)
HTTP in Go:
- [What to know about HTTP in Go](https://scene-si.org/2017/09/27/things-to-know-about-http-in-go/)


## Docker image build command

### Local build
```Bash
docker build -t coin_and_purse_ledger_api:0.1.0 .
```

## Image run command
This should primarily be deployed with the docker-compose file in the
[app repo.](https://github.com/Drewan-Tech/coin_and_purse_app)

Following is an example of a direct run command.
* Note: This requires a previously deployed network _appnet_.
```Bash
docker run --rm --network appnet -p 80:8080 coin_and_purse_ledger_api:0.1.0
```
