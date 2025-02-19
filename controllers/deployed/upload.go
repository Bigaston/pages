package deployed

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/pages/utils"
)

func UploadSite(c *fiber.Ctx) error {
	if !utils.IsDeployedSite(c) {
		return c.Next()
	}

	subDomain := c.Subdomains()

	if len(subDomain) == 0 {
		return c.SendStatus(http.StatusBadRequest)
	}

	siteName := utils.GetDomain(c)

	responseText, err := utils.UploadSite(siteName, c)

	if err != nil {
		return err
	}

	return c.SendString(responseText)
}
