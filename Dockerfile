FROM golang:1.26 AS builder

ARG VERSION=dev

WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags "-X sigs.k8s.io/external-dns/pkg/apis/externaldns.Version=${VERSION} -w -s" \
    -o external-dns .

FROM gcr.io/distroless/static:nonroot

COPY --from=builder /workspace/external-dns /external-dns

ENTRYPOINT ["/external-dns"]
