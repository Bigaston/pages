package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pages/controllers/deployed"
	"github.com/pages/utils"
)

func Init(app *fiber.App) {
	// Deployed Site
	app.Post("/~upload", deployed.UploadSite)
	app.Use("/", deployed.ServeSite)

	// Main Site
	app.Get("/", defaultPath)
}

func defaultPath(c *fiber.Ctx) error {
	if !utils.IsMainSite(c) {
		return c.Next()
	}

	return c.SendString("Tu es à la racine du site la !")
}
