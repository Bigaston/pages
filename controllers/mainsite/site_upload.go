package mainsite

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pages/utils"
)

func UploadSite(c *fiber.Ctx) error {
	if !utils.IsMainSite(c) {
		return c.Next()
	}

	siteName := c.Params("site")

	responseText, err := utils.UploadSite(siteName, c)

	if err != nil {
		return err
	}

	return c.SendString(responseText)
}
