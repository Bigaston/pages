package mainsite

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pages/utils"
)

func ListSites(c *fiber.Ctx) error {
	if !utils.IsMainSite(c) {
		return c.Next()
	}

	return c.Render("index", fiber.Map{})
}
