package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/pages/config"
)

func IsDeployedSite(c *fiber.Ctx) bool {
	domain := c.Get("Host", config.Config.Global.BaseHostname)

	fmt.Println(domain)

	return domain != config.Config.Global.BaseHostname
}

func IsMainSite(c *fiber.Ctx) bool {
	domain := c.Get("Host", config.Config.Global.BaseHostname)

	fmt.Println(domain)

	return domain == config.Config.Global.BaseHostname
}

func GetDomain(c *fiber.Ctx) string {
	return c.Get("Host", config.Config.Global.BaseHostname)
}

const TempDirectory = "./temp"
const DataDirectory = "./data"

var ImportantFiles []string = []string{".git"}
