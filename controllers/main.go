package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/pages/controllers/deployed"
	"github.com/pages/utils"
)

func Init(app *fiber.App) {
	// Deployed Site
	app.Post("/~upload", deployed.ServeSite)
	app.Use("/", deployed.UploadSite)

	// Main Site
	app.Get("/", defaultPath)
}

func defaultPath(c *fiber.Ctx) error {
	if !utils.IsMainSite(c) {
		return c.Next()
	}

	fmt.Println("Ouais Ouais Default Path")

	return c.SendString("Tu es à la racine du site la !")
}
