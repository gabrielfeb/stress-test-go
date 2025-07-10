# Estágio 1: Build
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/load-tester ./cmd/cli

FROM scratch
WORKDIR /app
COPY --from=builder /app/load-tester .
ENTRYPOINT ["/app/load-tester"]