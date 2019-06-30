# goatscale
<img align="left" src="/logo/goatscale.png">

goatscale is a repository to introduce infrastructure to developers who wish to learn the basics of designing a scalable game backend. The code language of choice is golang and you will need to know at least how to read this for later examples.

Each branch in this repository tries to explain a small portion of what you will need to know, building on these in digestible steps.

It is assumed that you are familiar with Docker and Docker Compose which are required to run each of the branches. In the future this will be expanded on when required for the more complex examples that are closer to a production situation. However to get started you should be able to follow the code with basic knowledge.

The requirements are described in the readme for each branch and how to get it up and running. However for the start you will need Docker installed, if you do not have this or are not sure what it is, you can find more information on the [docker site](https://www.docker.com/).

# Branches

## Phase 1
Phase 1 introduces the web server and basic infrastructure requirements to make this service scale. It finishes with a small chat application where a user can send a message from one client and receive it on all other connected clients.

### 0_base
The base branch is the starting point for the simple examples. If you are looking for a place to start this is the place.

### 1_simple_webserver
This adds a simple Iris web server to the docker infrastructure and demonstrates the problems that are encountered when running local development without using the server configuration in the base branch.

### 2_scalable_webserver
This branch fixes the issues that are encountered in '1_simple_webserver' by using Traefik and Consul.

### 3_client_webserver
This adds a simple client app which is built by Webpack and served by Iris as a single page application.

### 4_sessions_webserver
This branch adds Redis as a session store using Iris' built in session management.

### 5_messages_webserver
This branch adds another Redis instance as a pub/sub channel allowing messages to be sent between the web servers running in the infrastructure.

### 6_websockets_webserver
Currently WIP.

## WIP: Phase 2
Phase 2 adds the game server and sets up the web server to route clients into new game instances.

## WIP: Phase 3
Phase 3 adds logging, metrics, and KPI to visualise game data.

# WIP
* Aside from features this repo needs TESTS, and integration with Travis.ci and coveralls.

# Special thanks!
Shout out to [@evertras](https://github.com/Evertras) [@cainmartin](https://github.com/cainmartin) and [@jrouault](https://github.com/jrouault) for taking the time to review and give great feedback.

# Credits
Created my free logo at [LogoMakr](https://LogoMakr.com)