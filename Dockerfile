FROM golang:1.20 as builder

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy files
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /diglett-amd64-linux

FROM alpine:latest
COPY --from=builder /diglett-amd64-linux /diglett

# Make sure gin run in release mode
ENV GIN_MODE=release

ENTRYPOINT ["/diglett"]