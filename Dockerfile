# Stage 1: Build
FROM golang:1.25.1-alpine AS builder 

WORKDIR /app
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

# Copy the entire backend directory
COPY . .

# CHANGE: Name the output binary 'app-binary' to avoid 
# conflicting with your 'server' folder
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app-binary ./server/main.go

# Stage 2: Final (Distroless)
FROM gcr.io/distroless/base-debian12
WORKDIR /

# Copy the binary from the root of the builder
COPY --from=builder /app-binary .

EXPOSE 8084
USER nonroot:nonroot

# Run the uniquely named binary
ENTRYPOINT [ "./app-binary" ]