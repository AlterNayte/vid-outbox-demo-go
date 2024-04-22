# Build stage
FROM golang:1.20 AS builder
ENV GOOS linux
ENV CGO_ENABLED 0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main ./cmd/server.go

# Runtime stage
FROM gcr.io/distroless/static AS production

WORKDIR /app

COPY --from=builder /app/main /app/main

EXPOSE 1323

CMD ["/app/main"]