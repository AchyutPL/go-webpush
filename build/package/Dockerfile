# syntax=docker/dockerfile:1

FROM golang:1.22 AS builder

# Set destination for COPY
WORKDIR /app

RUN ls

# Download Go modules
COPY .env go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY . .

# Build
RUN make build

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose


FROM alpine:3.14 AS production

# Add certificates
WORKDIR /app

COPY .env .
COPY --from=builder /app/bin/main .
# Expose port
EXPOSE 8080

# Run
ENTRYPOINT ./main