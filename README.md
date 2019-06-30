# base
This is the starting point for most people and will introduce the base of the infrastructure that will allow us to start creating our backend services.

## Requirements
* Docker installed and running

## Running
To start the project from root

1. `docker-compose build`
2. `docker-compose up`

Two URLs will now be accessible to you.

* `traefik.localhost`
* `consul.localhost`

## What's going on?
In this branch there are two applications running, Traefik and Consul. Both of these are being run in separate containers managed by Docker, which are defined in a Docker Compose file. Docker compose is used to manage multiple containers at the same time as well as other parts of docker, for instance network connections.

Covering Docker and Compose is out of scope for this tutorial however if you want to read about it see the links below

* [Docker](https://www.docker.com/)
* [Docker Compose](https://docs.docker.com/compose/)

Running in Docker is Traefik which has two distinct jobs, and Consul which for the moment is used for one distinct job.

### Traefik
Traefik is acting as the [Reverse Proxy](https://en.wikipedia.org/wiki/Reverse_proxy) & [Load Balancer](https://en.wikipedia.org/wiki/Load_balancing_(computing)).

#### Reverse Proxy
In simple terms a reverse proxy allows the separation of internal infrastructure topology and how the services are accessed. There are two key pieces of configuration that make this work in the local environment.

The first is to tell docker to expose port 80 on the Traefik container and bind it to localhost port 80. This allows requests to be sent to Traefik. This can be found in the `docker-compose.yml` line 10.

The second part is to add an endpoint in Traefik on the same port, so that Traefik knows what port incoming requests will be made on. This can be found in `config/traefik.toml` line 2.

Now that localhost:80 is directing to our traefik container in docker, and traefik listening on this port, requests will be able to be processed. The next step is to configure Traefik to send these requests to the correct places.

Docker-compose allows us to set up configuration labels, these labels can be read by services to set the service configuration. In this case we are using the labels to tell Traefik about the services that were started. Both Traefik and Consul have accessible UIs, these services are being routed to using the labels defined in `docker-compose.yml` on lines 17 and 29.

The UI for Traefik is defined in `config/traefik.toml` line 7, and the UI for Consul automatically exposed when the service starts.

#### Load Balancer
The load balancer feature of Traefik is out of the box routing that automatically sends requests in a round robin fashion between instances of the same service. This will be explored later in the tutorials.

### Consul
Consul is used for [Service Discovery](https://en.wikipedia.org/wiki/Service_discovery) this will be explored later in the tutorials, at this stage it is possible to access the Consul catalogue of services and see them.

Traefik is automatically connecting to this catalogue based on configuration found in `config/traefik.toml` line 16 and registering the services found there.

## Conclusion
This branch introduced how a user can access the infrastructure and be redirected to services transparently. It has also set up two other key services with configuration based in `docker-compose.yml` and `config/traefik.toml`.
