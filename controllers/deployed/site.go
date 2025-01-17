package deployed

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/pages/utils"
)

func ServeSite(c *fiber.Ctx) error {
	if !utils.IsDeployedSite(c) {
		return c.Next()
	}

	if strings.Contains(c.Path(), ".git") {
		return c.SendStatus(http.StatusForbidden)
	}

	route := c.Path()

	path := filepath.Join(utils.DataDirectory, utils.GetDomain(c), route)

	fmt.Println(path)

	if file, err := os.Stat(path); err == nil && !file.IsDir() {
		return c.SendFile(path)
	}

	index_path := filepath.Join(path, "index.html")

	if _, err := os.Stat(index_path); err == nil {
		return c.SendFile(index_path)
	}

	return c.SendStatus(http.StatusNotFound)
}
