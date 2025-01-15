package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func serveSite(c *fiber.Ctx) error {
	if strings.Contains(c.Path(), ".git") {
		return c.SendStatus(http.StatusForbidden)
	}

	route := strings.Replace(c.Path(), "/~site", "", -1)

	path := filepath.Join(dataDirectory, route)

	if file, err := os.Stat(path); err == nil && !file.IsDir() {
		return c.SendFile(path)
	}

	index_path := filepath.Join(path, "index.html")

	if _, err := os.Stat(index_path); err == nil {
		return c.SendFile(index_path)
	}

	return c.SendStatus(http.StatusNotFound)
}
