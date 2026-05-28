# Stage 1: Build
FROM golang:1.22-alpine AS builder

RUN apk add --no-cache git ca-certificates build-base

WORKDIR /src

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build evilginx2
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /build/evilginx .

# Stage 2: Runtime
FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata libcap

# Create evilginx user
RUN adduser -D -h /home/evilginx evilginx

WORKDIR /home/evilginx

# Copy the binary
COPY --from=builder /build/evilginx /usr/local/bin/evilginx

# Copy phishlets and redirectors
COPY --from=builder /src/phishlets /home/evilginx/phishlets
COPY --from=builder /src/redirectors /home/evilginx/redirectors

# Create necessary directories
RUN mkdir -p /home/evilginx/.evilginx && \
    chown -R evilginx:evilginx /home/evilginx

# Allow binding to privileged ports
RUN setcap 'cap_net_bind_service=+ep' /usr/local/bin/evilginx

USER evilginx

EXPOSE 53 80 443 5000

VOLUME ["/home/evilginx/.evilginx"]

ENTRYPOINT ["evilginx"]
CMD ["-p", "/home/evilginx/phishlets", "-t", "/home/evilginx/redirectors"]
