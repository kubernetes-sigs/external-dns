# Build stage
FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
COPY go.tool.mod go.tool.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o external-dns .

# Runtime stage
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /app/external-dns /bin/external-dns
USER 65532:65532
ENTRYPOINT ["/bin/external-dns"]
