# Added Redis for sessions
Redis has been added to manage sessions in the web server.  Extra Iris packages have been added to interact with Redis.

## Requirements
* Docker installed and running

## Running
To start the project from root

1. `docker-compose build`
2. `docker-compose up --scale webserver=2`

Endpoints accessible

 * `webserver.localhost`
 * `consul.localhost`
 * `traefik.localhost`

New endpoints

 * `sessions.localhost`

## What's going on?
This branch enables session storage in Iris using Redis. To let Iris communicate with Redis we use the [github.com/gomodule/redigo package](https://github.com/gomodule/redigo). The package in Iris that is used to manage sessions is [github.com/kataras/iris/sessions/sessiondb/redis/service](https://github.com/kataras/iris/tree/master/sessions/sessiondb/redis).

[Redis Commander](https://github.com/joeferner/redis-commander) has been added to view the data in Redis for debug use.

In the `docker-compose.yml` the two services called `sessions` and `sessions-viewer` have been added. The `sessions` service uses the Alpine version of Redis while `sessions-viewer` is based on a Node Alpine distribution. Both of these are lightweight images and it will not matter if multiple services spawn their own individual Redis instances in future branches.

The code added to `cmd/webserver/main.go` is quite simple as Iris does most of the heavy lifting.

The important commands related to session are inside the login route.

```go
app.Post("/login", func(ctx iris.Context) {
	var user User
	ctx.ReadJSON(&user)

	s := sess.Start(ctx)
	s.Set("userData", user)

	app.Logger().Info("Login Called with user id: ", user.ID)

	ctx.StatusCode(iris.StatusOK)
})
```

- `sess.Start(...)` creates a new session key in Redis
- `s.Set(...)` creates a new key for `userData` in Redis

It is also worth noting the configuration of the redis store does have other options than the ones displayed. Documentation can be found [here](https://godoc.org/github.com/gomodule/redigo/redis).

## Issues
The connection to Redis does not attempt to reconnect in the event of failure.  Additionally the address is hardcoded, not configurable.  This is fine for the sake of demonstration but should be noted.

## Conclusion
This is a simple addition to the code base with no real side effects. It shows how to configure `docker-compose` to add Redis and Redis Commander and have them talk to each other, and it shows how to connect and establish a session with `Iris`. The next branch will show how to validate a user has a session when calling an API that sends a message between servers.

## Diff from 3_client_webserver
[Diff](https://github.com/nullorvoid/goatscale/compare/3_client_webserver...4_sessions_webserver)
