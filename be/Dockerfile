# Stage 1: Build the Go binary
FROM golang:1.22.6 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Force Go to build a static Linux binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/psp-dashboard-be cmd/main.go

# Stage 2: Minimal container to run the binary
FROM scratch

# Copy binary from builder
COPY --from=builder /app/bin/psp-dashboard-be /bin/psp-dashboard-be

# Set binary as entrypoint
ENTRYPOINT ["/bin/psp-dashboard-be"]
