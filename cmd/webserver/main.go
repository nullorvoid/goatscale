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
	"github.com/nullorvoid/goatscale/lib/chat"
	"github.com/nullorvoid/goatscale/lib/consulapi"
	"github.com/nullorvoid/goatscale/lib/pubsubapi"
)

// By default our webserver will be exposed on 8080 on any network interface
var addr = flag.String("addr", "0.0.0.0:8080", "http service address")

// User bind struct
type User struct {
	ID string `json:"userId"`
}

// Message bind struct
type Message struct {
	Message string `json:"message"`
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

	// Create new pub sub client
	pubsub, err := pubsubapi.NewPubSubClient("pubsub:6379")

	if err != nil {
		app.Logger().Fatal("Error creating pub sub client: ", err)
	}

	// Register to chat
	chat, err := chat.NewChatClient(pubsub)

	if err != nil {
		app.Logger().Fatal("Error creating chat client: ", err)
	}

	// Background task for the reading of chat messages
	go func() {
		for {
			msg, err := chat.GetNextMessage()
			if err != nil {
				app.Logger().Fatal("Error getting next message from chat: ", err)
			}

			// Message logged in future can be used with other modules
			app.Logger().Info("Chat Handler: " + msg)
		}
	}()

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
		s.Set("authenticated", true)

		app.Logger().Info("Login Called with user id: ", user.ID)

		ctx.StatusCode(iris.StatusOK)
	})

	// Api for message which does a broadcast over the chat channel to all servers
	app.Post("/message", func(ctx iris.Context) {
		// Check if user is authenticated
		if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
			ctx.StatusCode(iris.StatusForbidden)
			return
		}

		var msg Message
		ctx.ReadJSON(&msg)

		// Send the message received to chat
		err := chat.SendMessage(msg.Message)

		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
		} else {
			ctx.StatusCode(iris.StatusOK)
		}
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
