package main

import (
	// Input flags for command line configuration
	"flag"

	// Iris webserver
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

// By default our webserver will be exposed on 8080 on any network interface
var addr = flag.String("addr", "0.0.0.0:8080", "http service address")

func main() {
	// If we want to set an address from the command line we parse it here
	flag.Parse()

	// Create a new Iris web server
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	// Set a GET route that returns some simple HTML
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})

	// Run the server, this will hold open the application
	app.Run(iris.Addr(*addr), iris.WithoutServerError(iris.ErrServerClosed))
}
