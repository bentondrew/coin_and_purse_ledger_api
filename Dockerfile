FROM golang:1.11.5-alpine3.8 as builder
RUN apk add --no-cache git && \
    adduser -D -g '' gouser
COPY app/ $GOPATH/src/github.com/Drewan-Tech/coin_and_purse_ledger_service/app/
WORKDIR $GOPATH/src/github.com/Drewan-Tech/coin_and_purse_ledger_service/app
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go test github.com/Drewan-Tech/coin_and_purse_ledger_service/app/api
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/main

FROM scratch
LABEL maintainer="Benton Drew <benton.s.drew@drewantech.com>"
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --chown=gouser --from=builder /go/bin/ /service/
USER gouser
WORKDIR /service/
EXPOSE 8080
ENV DB_HOST ledgerdb
ENV DB_PORT 5432
ENV DB_USER ledgerservice
ENV DB_PASS 12345
ENV DB_DATABASE ledger
ENTRYPOINT ["/service/main"]
