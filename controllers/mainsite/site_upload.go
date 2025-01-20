package mainsite

import (
	"fmt"

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

func UploadSiteWithName(c *fiber.Ctx) error {
	if !utils.IsMainSite(c) {
		return c.Next()
	}

	subdomain := c.FormValue("subdomain")
	domain := c.FormValue("domain")

	siteName := subdomain + domain

	fmt.Println(siteName)

	responseText, err := utils.UploadSite(siteName, c)

	if err != nil {
		return err
	}

	return c.SendString(responseText)
}
