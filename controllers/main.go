package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pages/controllers/deployed"
	"github.com/pages/controllers/mainsite"
)

func Init(app *fiber.App) {
	// Deployed Site
	app.Post("/~upload", deployed.UploadSite)
	app.Use("/", deployed.ServeSite)

	// Main Site
	app.Post("/site/upload", mainsite.UploadSiteWithName)
	app.Post("/site/:site/upload", mainsite.UploadSite)
	app.Get("/site/:site", mainsite.SiteEdit)
	app.Get("/dashboard", mainsite.Dashboard)
	app.Use("/public", mainsite.ServePublic)
}
