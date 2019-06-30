# Web server client
In this branch the web server is now serving HTML that is being built and bundled with webpack.

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

## What's going on?
This branch has two distinct changes to the code base. The first is we are adding a `client` folder and making this build with webpack, the second part is adding the build step to the `webserver.Dockerfile` and having the web server serve these files.

### Client
The `client` folder contains the standard node set up including `package.json`, and `package-lock.json` with a `src` folder holding the Typescript files and a single HTML page. There is also a `tsconfig.json` because the project is Typescript and will be transpiled, and the `webpack.config.js` because webpack needs to know what and how to compile the bundle with.

Inside the `src` folder `index.ts` builds the page dynamically using very simple HMTL. There are two other points going on in this file, the first is a simple UUID generation function, and the second the user identification storage.

The UUID generation is very simple and is not a true versioned UUID but is good enough for this demo repository. In a production scenario or for integration with other services, a package that can generate a standardised UUID should be used.

For user identification the web page stores this generated UUID into `window.localStorage`. This can be reset by the user but is normally fine for demos and examples. In production a login system should be incorporated with username and password, or integration with a different OAuth based platform. The code that uses this is in the function `getUserId` in `index.ts`.

`webapiclient.ts` adds the ability to call into the api that are exposed on the web server. It is fairly generic handler that passes any data that is returned by the API and makes sure we have the correct passing http code.

### Web Server
In the web server a new API has been added and the server is now serving the `index.html` file and not raw HTML string as previous branches have done. The new API is just a simple login test that we will be expanding with session management in a later branch. The format is defined at the beginning of the file as `type User struct` and parsed as JSON inside the route handler.

### Build process
The build process has been updated to a two step build. The first for the web server and the second for the client using webpack. Both the builds are done on alpine images of the respective technologies and then copied into an alpine only distribution. In a production game assets are very important to consider when making game builds and a lot of customisation is normally done to reduce the load on the servers and host bundles of assets on CDNs. This is outside the scope of this tutorial but is something to consider when releasing a production game.

The client is a node / typescript project which is built by webpack and output to a folder called `dist`. It is not required to be done outside docker as the `webserver.Dockerfile` will build this and copy the correct files into a folder to be distributed.

## Issues
In the diff it can be seen that `depends_on` has been added to the Traefik service description. This was to prevent containers coming up in the wrong order. This only reduces the problems encountered when spawning infrastructure it does **NOT** solve it. The `depends_on` flag specifies that one container relies on the other to have started, not that one container is running a service that another relies on. This distinction should be noted as `depends_on` only guarantees that the container is up and running, not that any service running on it is.

To fix this issue scripts that are run to test and wait for other container services to be up could be written. At this stage it would make the tutorial complex and diverge from the intent of these basic branches. In future branches this may become necessary and at that time it will be introduced. For the moment it is good enough that this problem has been displayed.

## Conclusion
Now that the base application is set up and APIs have been proven to work the application can start to take shape. This finalises the base infrastructure with all the basic configuration required to access functionality on the web server. In the next branches user identification will be added in the form of sessions, and a message passing system to allow users to talk to each other regardless of what server they send their messages to.

## Diff from 2_scalable_webserver
[Diff](https://github.com/nullorvoid/goatscale/compare/2_scalable_webserver...3_client_webserver)
