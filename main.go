package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/pages/config"
	"github.com/pages/controllers"
)

func main() {
	config.ReadConfig("./config.toml")

	app := fiber.New()
	app.Use(logger.New())

	controllers.Init(app)

	log.Fatal(app.Listen(":3000"))
}

// 	tlsConfig := tls.Config{
// 		GetCertificate: func(helloInfo *tls.ClientHelloInfo) *tls.Certificate, error{

// 		},
// 	}
// 	listener, err := tls.Listen("tcp", ":443", &tlsConfig)

// 	if err != nil {
// 		panic(err)
// 	}

// 	go httpApp()

// 	log.Fatal(app.Listener(listener))
// }

// func httpApp() {
// 	app := fiber.New()

// 	app.Use(logger.New())

// 	app.Get("/", func(c *fiber.Ctx) error {
// 		return c.Redirect(fmt.Sprintf("https://%s%s", c.GetReqHeaders()["Host"][0], c.Path()))
// 	})

// 	log.Fatal(app.Listen(":80"))
// }
