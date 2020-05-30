FROM rust:1.43 as build
WORKDIR /ledger_api
COPY ledger_api/ /ledger_api
RUN cargo build --release

FROM gcr.io/distroless/cc
COPY --from=build /ledger_api/target/release/ledger_api /
ENTRYPOINT ["./ledger_api"]
