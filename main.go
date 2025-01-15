package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/pages/controllers"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	controllers.Init(app)

	log.Fatal(app.Listen(":3000"))
}
