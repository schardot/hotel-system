# Build stage
FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go application
RUN go build -o go-app .

# Final stage
FROM golang:1.20

WORKDIR /app

COPY --from=builder /app/go-app /app/go-app
COPY --from=builder /app .

CMD ["./go-app"]
