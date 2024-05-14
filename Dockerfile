FROM golang:alpine as builder

WORKDIR /app/server

COPY server/go.mod server/go.sum ./

RUN go mod download

COPY server/ .

RUN go build -ldflags="-s -w" -o main ./cmd

FROM debian:bookworm-slim

RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server/main /app/server/main

EXPOSE 8080

CMD ["/app/server/main"]

