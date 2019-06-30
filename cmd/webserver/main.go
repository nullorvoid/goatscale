package main

import (
	// Standard golang packages
	"errors"
	"flag"
	"net"
	"time"

	// Iris webserver
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/sessions/sessiondb/redis"
	"github.com/kataras/iris/sessions/sessiondb/redis/service"

	// Custom packages
	"github.com/nullorvoid/goatscale/lib/consulapi"
)

// By default our webserver will be exposed on 8080 on any network interface
var addr = flag.String("addr", "0.0.0.0:8080", "http service address")

// User bind struct
type User struct {
	ID string `json:"userId"`
}

func main() {
	// If we want to set an address from the command line we parse it here
	flag.Parse()

	db := redis.New(service.Config{
		Network:     "tcp",
		Addr:        "sessions:6379",
		IdleTimeout: time.Duration(5) * time.Minute,
	})

	// Close the database connection if the application closed
	defer db.Close()

	sess := sessions.New(sessions.Config{
		Cookie:       "sessionscookieid",
		Expires:      45 * time.Minute,
		AllowReclaim: true,
	})

	sess.UseDatabase(db)

	// Create a new Iris web server
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	// Get IP to register to consul with
	ip, err := getIP()

	if err != nil {
		app.Logger().Fatal("Error getting ip: ", err)
	}

	err = registerToConsul(ip)

	if err != nil {
		app.Logger().Fatal("Error creating consul client: ", err)
	}

	// Register routes for our application
	app.RegisterView(iris.HTML("./public", ".html"))

	// Get basic single page view
	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	// Login route to check user can call functional APIs
	app.Post("/login", func(ctx iris.Context) {
		var user User
		ctx.ReadJSON(&user)

		s := sess.Start(ctx)
		s.Set("userData", user)

		app.Logger().Info("Login Called with user id: ", user.ID)

		ctx.StatusCode(iris.StatusOK)
	})

	// Serve pages from public
	assetHandler := app.StaticHandler("./public", false, false)
	app.SPA(assetHandler)

	// Run the server, this will hold open the application
	app.Run(iris.Addr(*addr), iris.WithoutServerError(iris.ErrServerClosed))
}

func getIP() (net.IP, error) {
	ifaces, err := net.Interfaces()

	if err != nil {
		return nil, err
	}

	// This is currently a hack which I'm not happy with...
	for _, i := range ifaces {
		addrs, _ := i.Addrs()

		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					return ipnet.IP, nil
				}
			}
		}
	}

	return nil, errors.New("no ip addresses found")
}

func registerToConsul(ip net.IP) error {
	consul, err := consulapi.NewConsulClient("consul:8500")

	if err != nil {
		return err
	}

	return consul.Register("webserver", ip.String(), 8080)
}
