# Build image from golang to make sure we have everything to compile
FROM golang:alpine AS build-golang
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

# Build golang application
RUN go build -o ./bin/main -ldflags "-s -w" ./cmd/webserver/

# Build image from node to do the webpack compile to static
FROM node:alpine AS build-node
WORKDIR /app

# Prepare environment to pull dependencies
RUN apk update

# Copy webpack config and build components
COPY ./client .
# Install npm and run webpack build to create the client
RUN npm install
RUN npm run build

# Build runtime image
FROM alpine:latest
WORKDIR /app
# Copy golang server
COPY --from=build-golang /app/bin .
COPY --from=build-golang /app/scripts ./scripts
# Copy client
COPY --from=build-node /app/dist ./public

# Execute the scripts to start the server
ENTRYPOINT [ "./scripts/webserver.entrypoint.sh" ]
