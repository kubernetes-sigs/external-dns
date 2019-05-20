FROM golang:1.12.5 as builder
WORKDIR /external-dns
COPY . .
RUN make build

FROM gcr.io/distroless/static
COPY --from=builder /external-dns/build/external-dns /external-dns
ENTRYPOINT ["./external-dns"]
