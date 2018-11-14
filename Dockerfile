FROM golang:1.11.0-alpine3.8 as builder
RUN adduser -D -g '' gouser
COPY code/ $GOPATH/src/github.com/Drewan-Tech/coin_and_purse_ledger_service/
WORKDIR $GOPATH/src/github.com/Drewan-Tech/coin_and_purse_ledger_service/cmd/srv
RUN go get -d -v
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go test -c ./home -o /go/bin/tests/home.test
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/main

FROM scratch
LABEL maintainer="Benton Drew <benton.s.drew@drewantech.com>"
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --chown=gouser --from=builder /go/bin/ /service/
USER gouser
WORKDIR /service/
EXPOSE 80
ENTRYPOINT ["/service/main"]
