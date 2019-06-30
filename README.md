# Simple Web Server
Adds a golang web server into the docker infrastructure and exposes it on localhost:8080.

## Requirements
* Docker installed and running

## Running
To start the project from root

1. `docker-compose build`
2. `docker-compose up`

One new endpoint will be accessible.

* `localhost:8080`

## What's going on?
This branch builds on the `base` branch and adds a Golang web server. The web server is using the [Iris framework](https://iris-go.com/) to serve HTML to the client. A simple Golang set up is used and the web server kept simple to highlight issues without a reverse proxy which will also be explored later when a game server is added.

### Folder Structure
This branch has two additional folders, `cmd` & `scripts`. `cmd` is used to hold the Golang applications entry points. The code in here should be lightweight angod anything that might be used by multiple applications should be separated into packages to be included. The `scripts` folder holds the entrypoint for the web server container, this will be dicussed below.

### Golang Configuration
The project is set up with 'Go Modules' which manages the dependencies for the web server. For more information on the package management system in Go the docs are avaliable on the [Golang website](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more). The module system creates two files in the root of the project `go.mod` & `go.sum`.

`go.mod` holds all the version numbers for each package that is included. It is human readable so that anyone can have a look at all the included packages including indirect packages.

`go.sum` holds all the checksum information for packages in the project. It also holds checksums for packages that have been in the project in the past and have been removed.

### Iris
The Iris web framework will be used in the web server and throughout the examples in this repository. You can find out more information on the [Iris website](https://iris-go.com/).

### Scripts
The `scripts` folder is a choice which could be considered strange. The `webserver.entrypoint.sh` just starts the server as is, this could be done from the `webserver.Dockerfile`. The reasoning behind this choice is to be able to extend the start up in later tutorials, a common use case would be to wait for database connectivity and then push in sample data via this script for initialisation. For the moment this is not implemented but will be used in later turorials.

## Issues
There is a big problem with this set up when you want to locally test scalability. The port is shared to localhost, so you can only have one server running at a time. This means you can not run `docker-compose up --scale webserver=2`.

## Conclusion
This sets up the basic web server using Iris which can be tested by connecting to `localhost:8080`, however it also demonstrates the problem with local development with multiple web servers. In the next stage we will use the Traefik and Consul to solve this issue.

## Diff from 0_base
[Diff](https://github.com/nullorvoid/goatscale/compare/0_base...1_simple_webserver)
