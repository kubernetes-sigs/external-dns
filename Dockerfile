FROM alpine:3.5

COPY build/linux-amd64/external-dns /external-dns

ENTRYPOINT ["/external-dns"]
