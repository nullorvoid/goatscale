# Messages between web servers
WIP

## Requirements
* Docker installed and running

## Running
To start the project from root

1. `docker-compose build`
2. `docker-compose up --scale webserver=2`

Endpoints accessible

 * `sessions.localhost`
 * `webserver.localhost`
 * `consul.localhost`
 * `traefik.localhost`

New endpoints

 * `pubsub.localhost`

## What's going on?

## Issues

## Conclusion

## Diff from 4_sessions_webserver
[Diff](https://github.com/nullorvoid/goatscale/compare/4_sessions_webserver...5_messages_webserver)
