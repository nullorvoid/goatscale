# Build image from golang to make sure we have everything to compile
FROM golang:alpine AS build-env
WORKDIR /app

# Prepare environment to pull dependencies
RUN apk update
RUN apk add git

# Copy go mod and sum files
COPY ./go.* ./

# Get dependancies which will be cached if the mod files don't change
RUN go mod download

# Copy webserver entry, scripts for running the server
COPY ./cmd/webserver ./cmd/webserver
COPY ./scripts ./scripts
COPY ./lib ./lib

# Build application
RUN go build -o ./bin/main -ldflags "-s -w" ./cmd/webserver/

# Build runtime image
FROM alpine:latest
WORKDIR /app
COPY --from=build-env /app/bin .
COPY --from=build-env /app/scripts ./scripts

# Execute the scripts to start the server
ENTRYPOINT [ "./scripts/webserver.entrypoint.sh" ]
