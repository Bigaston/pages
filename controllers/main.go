package controllers

import "github.com/gofiber/fiber/v2"

const tempDirectory = "./temp"
const dataDirectory = "./data"

func Init(app *fiber.App) {
	app.Post("/~upload/:site", uploadSite)

	app.Use("/~site/:site", serveSite)
}
