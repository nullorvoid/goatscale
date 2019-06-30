# Scalable Web Server
This configuration solves the scalability issue of the '1_simple_webserver' by having the web server add itself to the consul catalogue, and accessed via Traefik.

## Requirements
* Docker installed and running

## Running
To start the project from root

1. `docker-compose build`
2. `docker-compose up --scale webserver=2`

Endpoints accessible

 * `consul.localhost`
 * `traefik.localhost`

New endpoints

* `webserver.localhost`

## What's going on?
In this branch we have added the Consul client to a new package found in the `lib` folder. We connect to Consul when the server starts and register our presence. During this registration we also add specific `labels` for Traefik to use when reading the Consul catelgue, these labels `traefik.backend` specifies the name of the service and `traefik.enabled` to expose the web server via Traefik.

Traefik will then use a load balance algorithm to direct requests via the two registered servers that were started with `--scale webserver=2`, by default this algorithm is simple round robin.

### Folder Structure
The `lib` folder has been added to the project and copied into the build in the `webserver.Dockerfile`. This folder will serve as the general purpose code folder, holding modules that will be used in the applications. For the moment it just has a wrapper to the consul api.

## Issues
There are some issues with the setup of the webserver, Consul and Traefik. The biggest issue is that no health status for the nodes has been written, this means that if servers crash or get shutdown these changes will not be seen in Traefik. For the purpose of this repo which is an introduction to the basic components of a game server architecture it is suitable, however this is far from production code, and would need a lot more checks and APIs to get to that stage.

## Conclusion
In this branch Consul has been added and the `lib` folder to hold the code for the API. It demonstrates how requests can be routed over multiple web servers and how we can register with Consul and have Traefik read this information and route requests accordingly.

## Diff from 1_simple_webserver
[Diff](https://github.com/nullorvoid/goatscale/compare/1_simple_webserver...2_scalable_webserver)
