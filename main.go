package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pages/controllers"
)

func main() {
	app := fiber.New()

	controllers.Init(app)

	log.Fatal(app.Listen(":3000"))
}
