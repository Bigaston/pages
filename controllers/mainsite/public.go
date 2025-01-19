package mainsite

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/pages/utils"
)

func ServePublic(c *fiber.Ctx) error {
	if !utils.IsMainSite(c) {
		return c.Next()
	}

	route := strings.Replace(c.Path(), "/public", "", -1)

	path := filepath.Join("layouts/public", route)

	fmt.Println(path)

	if file, err := os.Stat(path); err == nil && !file.IsDir() {
		return c.SendFile(path)
	}

	return c.SendStatus(http.StatusNotFound)
}
